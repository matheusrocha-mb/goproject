go-build:
	go build -o rock3 main.go && cp rock3 ~/go/bin && rm rock3

rock-build:
	docker-compose build

rock-start:
	docker-compose up -d

rock-down:
	docker-compose down -v --remove-orphans

go-run:
	go run main.go

docker-run:
	docker run rock3_image

docker-build:
	docker build -t rock3_image .