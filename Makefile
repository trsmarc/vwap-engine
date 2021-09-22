#!sh
APP_NAME:=vwap-engine

run:
	go run .

test:
	go test -race -v ./...

docker-build:
	- docker build -t ${APP_NAME} .

docker-run: docker-build
	- docker run -it ${APP_NAME} go run .

docker-test: docker-build
	- docker run -it ${APP_NAME} go test -race -v ./...
