
all:
	go build -o monolith ./cmd/monolith
	go build -o shop ./cmd/microservices/shop
	go build -o orders ./cmd/microservices/orders
	go build -o payments ./cmd/microservices/payments

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
