-- +goose Up
-- Prices are int64 minor units. The first version uses KZT only.
-- updated_at is maintained by application UPDATE statements.
CREATE TABLE users (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text NOT NULL CHECK (btrim(name) <> ''),
    login text NOT NULL UNIQUE,
    api_key_hash text NOT NULL UNIQUE CHECK (btrim(api_key_hash) <> ''),
    password text NOT NULL,
    role text NOT NULL DEFAULT 'CLIENT' CHECK (role IN ('CLIENT', 'SELLER', 'ADMIN')),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE products (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    sku text NOT NULL UNIQUE CHECK (btrim(sku) <> ''),
    name text NOT NULL CHECK (btrim(name) <> ''),
    price bigint NOT NULL CHECK (price >= 0),
    currency text NOT NULL DEFAULT 'KZT' CHECK (currency = 'KZT'),
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX products_active_id_idx ON products (id) WHERE is_active;

CREATE TABLE inventory (
    product_id bigint PRIMARY KEY REFERENCES products (id),
    available_quantity bigint NOT NULL DEFAULT 0 CHECK (available_quantity >= 0),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE orders (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id bigint NOT NULL REFERENCES users (id),
    status text NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'confirmed', 'rejected')),
    total_amount bigint NOT NULL CHECK (total_amount >= 0),
    currency text NOT NULL DEFAULT 'KZT' CHECK (currency = 'KZT'),
    rejection_reason text,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (user_id, id),
    CHECK (
        (status = 'rejected' AND rejection_reason IS NOT NULL AND btrim(rejection_reason) <> '')
        OR (status <> 'rejected' AND rejection_reason IS NULL)
    )
);

CREATE INDEX orders_user_created_idx ON orders (user_id, created_at DESC, id DESC);
CREATE INDEX orders_pending_idx ON orders (created_at, id) WHERE status = 'pending';

CREATE TABLE order_items (
    order_id uuid NOT NULL REFERENCES orders (id),
    product_id bigint NOT NULL REFERENCES products (id),
    quantity bigint NOT NULL CHECK (quantity > 0),
    unit_price bigint NOT NULL CHECK (unit_price >= 0),
    PRIMARY KEY (order_id, product_id)
);

CREATE INDEX order_items_product_idx ON order_items (product_id);

-- A reservation must refer to an actual item of the same order.
-- Matching quantities and atomic stock updates are enforced by the service transaction.
CREATE TABLE reservations (
    order_id uuid NOT NULL,
    product_id bigint NOT NULL,
    quantity bigint NOT NULL CHECK (quantity > 0),
    created_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id, product_id) REFERENCES order_items (order_id, product_id)
);

CREATE TABLE idempotency_records (
    user_id bigint NOT NULL REFERENCES users (id),
    idempotency_key text NOT NULL CHECK (length(btrim(idempotency_key)) BETWEEN 1 AND 255),
    request_fingerprint text NOT NULL CHECK (btrim(request_fingerprint) <> ''),
    order_id uuid NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, idempotency_key),
    FOREIGN KEY (user_id, order_id) REFERENCES orders (user_id, id)
);

CREATE INDEX idempotency_records_order_idx ON idempotency_records (user_id, order_id);

CREATE TABLE outbox_events (
    event_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id uuid NOT NULL REFERENCES orders (id),
    event_type text NOT NULL CHECK (btrim(event_type) <> ''),
    schema_version integer NOT NULL DEFAULT 1 CHECK (schema_version > 0),
    payload jsonb NOT NULL CHECK (jsonb_typeof(payload) = 'object'),
    created_at timestamptz NOT NULL DEFAULT now(),
    published_at timestamptz
);

CREATE INDEX outbox_events_pending_idx ON outbox_events (created_at, event_id)
    WHERE published_at IS NULL;
CREATE INDEX outbox_events_order_idx ON outbox_events (order_id);

-- No FK to outbox: consumed Kafka events may outlive producer-side records.
CREATE TABLE processed_events (
    consumer_name text NOT NULL CHECK (btrim(consumer_name) <> ''),
    event_id uuid NOT NULL,
    processed_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (consumer_name, event_id)
);

CREATE TABLE notifications (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id uuid NOT NULL,
    order_id uuid NOT NULL REFERENCES orders (id),
    channel text NOT NULL CHECK (channel IN ('database', 'webhook')),
    status text NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'sent', 'failed')),
    payload jsonb NOT NULL CHECK (jsonb_typeof(payload) = 'object'),
    attempts integer NOT NULL DEFAULT 0 CHECK (attempts >= 0),
    next_attempt_at timestamptz NOT NULL DEFAULT now(),
    last_error text,
    sent_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (event_id, channel),
    CHECK ((status = 'sent' AND sent_at IS NOT NULL) OR (status <> 'sent' AND sent_at IS NULL))
);

CREATE INDEX notifications_order_created_idx ON notifications (order_id, created_at, id);
CREATE INDEX notifications_pending_idx ON notifications (next_attempt_at, id)
    WHERE status = 'pending';

-- Order totals, immutable purchase prices, legal state transitions and
-- all-or-nothing reservation require application-level transaction logic.
-- +goose Down
DROP TABLE notifications;
DROP TABLE processed_events;
DROP TABLE outbox_events;
DROP TABLE idempotency_records;
DROP TABLE reservations;
DROP TABLE order_items;
DROP TABLE orders;
DROP TABLE inventory;
DROP TABLE products;
DROP TABLE users;
