build-go:
	go build -o rock3 main.go

docker-build:
	docker build -t rock3_image .

docker-run:
	docker run rock3_image

docker-compose-up:
	docker-compose up dev