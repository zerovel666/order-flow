# OrderFlow

Учебный backend на Go. 

## Окружение

- Go 1.27.0 — версия из go.mod.
- Git.
- Docker с Compose v2+ для контейнерного запуска.
- Команды ниже рассчитаны на PowerShell.

## Получение проекта

```powershell
git clone https://github.com/zerovel666/order-flow.git
cd order-flow
```

Инструкции относятся к версии с текущими изменениями. Локальные изменения станут доступны при клонировании после их публикации.

## Локальный запуск

```powershell
go run ./cmd/app
```

## Форматирование и проверки

```powershell
go fmt ./...
go mod tidy
go test ./...
go vet ./...
```

`go fmt` форматирует код, `go mod tidy` синхронизирует зависимости, `go test` запускает тесты и компилирует пакеты, `go vet` выполняет статические проверки. Сейчас тестов нет: [no test files] ожидаемо. Внешних зависимостей нет, поэтому отсутствие go.sum нормально.

## Сборка

```powershell
go build -o ./bin/order-flow.exe ./cmd/app
./bin/order-flow.exe
```

## Docker

```powershell
docker compose config --quiet
docker compose up --build -d
Invoke-RestMethod http://localhost:8000/health-check
docker compose logs app
docker compose down
```

Перед запуском освободить порт 8000 от локального экземпляра. Compose также публикует PostgreSQL на 5432. Учебные учётные данные в Compose не предназначены для production. Команда down сохраняет данные в named volume.

## Структура и основные понятия

- cmd/app — точка входа.
- internal/api — маршрутизация и HTTP-обработчики.
- internal/utils/responses — JSON-ответы.
- TASK.md — учебное ТЗ.

`go.mod` описывает модуль целиком; директории с Go-файлами образуют пакеты. `internal` ограничивает импорт кодом родительского дерева. Экспортируемые имена начинаются с заглавной буквы. Go компилирует приложение в исполняемый файл; для запуска собранного приложения интерпретатор PHP/Python не требуется.
