
run:
	go run main.go


build:
	go build main.go

hello:
	echo "Hello, world!"

docker:
	docker build -t my-go-app .

docker_run:
	docker run --rm -it my-go-app

run2:
	docker run --rm -it -p 8080:8080 my-go-app

.PHONY: run build hello docker run run2	