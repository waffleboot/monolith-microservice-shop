package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	orders "monolith-microservice-shop/pkg/orders/interfaces/public/http"
)

func main() {
	req := orders.PostOrderRequest{
		ProductID: "1",
		Address: orders.PostOrderAddress{
			Name:     "n/a",
			Street:   "n/a",
			PostCode: "n/a",
			Country:  "Russia",
			City:     "Moscow",
		},
	}
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("http://localhost:8080/orders", "application/json", bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		fmt.Println(string(data))
		return
	}
	var result orders.PostOrdersResponse
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.OrderID)
}
