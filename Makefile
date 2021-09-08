
all:
	go build -o client ./cmd/client
	go build -o monolith ./cmd/monolith
	go build -o shop ./cmd/microservices/shop
	go build -o orders ./cmd/microservices/orders
	go build -o payments ./cmd/microservices/payments

http:
	curl -s http://localhost:8080/products | jq .
	curl -s -H "Content-Type: application/json" --request POST --data '{"product_id":"2","address":{"name":"name","street":"street","city":"city","post_code":"123","country":"RU"}}' http://localhost:8080/orders | jq .OrderID
	curl -s http://localhost:8080/orders/4ba0669e-fc66-11eb-b996-c82a142c4d0c/paid | jq .

qa:
    # "Errors unhandled" check is made by errcheck
	gometalinter \
	    --vendor \
	    --deadline=60s \
	    --exclude="composite literal uses unkeyed fields" \
	    --exclude="should have comment or be unexported" \
	    --exclude="Errors unhandled" \
	    ./...
	go-cleanarch

up:
	docker-compose up

docker-test:
	docker-compose exec tests go test -v ./tests/...

docker-test-monolith:
	docker-compose exec tests go test -v -run "/monolith" ./tests/...

docker-test-microservices:
	docker-compose exec tests go test -v -run "/microservices" ./tests/...
