create-volumes:
	docker volume create --name=postgres_bambino_volume

remove-volumes:
	docker rm postgres_bambino_volume

docker-compose:
	docker-compose -f docker-compose.yaml up

run:
	go run bambino.go

build:
	go build bambino.go