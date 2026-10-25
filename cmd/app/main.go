package main

import (
	"fmt"
	"order-flow/internal/api"
)

func main() {
	fmt.Println("Order-flow service is starting")

	api.HandleRequest()
}
