.PHONY: build up down
build:
	env GOOS=linux GOARCH=386 go build -o build cmd/main.go
	docker-compose build
up:
	env GOOS=linux GOARCH=386 go build -o build cmd/main.go
	docker-compose up backend && docker-compose rm -fsv
down:
	docker-compose down --volumes

migrate:
	env GOOS=linux GOARCH=386 go build -o build cmd/main.go
	docker-compose run -e MIGRATE=true backend && docker-compose rm -fsv
test:
	env GOOS=linux GOARCH=386 go test -c testing
	docker-compose up test && docker-compose rm -fsv

binary:
	env GOOS=linux GOARCH=386 go build -o build cmd/main.go

test-binary:
	env GOOS=linux GOARCH=386 go test -c testing

clean:
	docker rm -f $(docker ps -a -q)
