package main

// yep, it's a bit ugly :(
import (
	"log"
	"net/http"
	"os"

	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/common/cmd"
	payments_ipc "github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/payments/interfaces/ipc"
)

func main() {

	log.Println("Starting monolith")

	ctx := cmd.Context()

	paymentsChannel := make(chan payments_ipc.OrderToProcess)

	router, payments := createService(paymentsChannel)

	go payments.Run()

	server := &http.Server{Addr: os.Getenv("SHOP_MONOLITH_BIND_ADDR"), Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Monolith is listening on %s", server.Addr)

	<-ctx.Done()

	log.Println("Closing monolith")

	if err := server.Close(); err != nil {
		panic(err)
	}

	close(paymentsChannel)

	payments.Close()

}
