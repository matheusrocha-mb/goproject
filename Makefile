go-build:
	go build -o rock3 main.go && cp rock3 ~/go/bin && rm rock3

rocky-build:
	docker-compose build

rocky-start:
	docker-compose up -d

rocky-down:
	docker-compose down -v --remove-orphans

go-run:
	go run main.go

docker-run:
	docker run rock3_image

docker-build:
	docker build -t rock3_image .