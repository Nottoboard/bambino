NAME=bambino:v0.0.2

build:
	docker build --tag bandnoticeboard/$(NAME) .

run:
	docker run --add-host=host.docker.internal:host-gateway --env-file ./.env-local bandnoticeboard/$(NAME)

push:
	docker push bandnoticeboard/$(NAME)


create-volumes:
	docker volume create --name=postgres_bambino_volume

remove-volumes:
	docker rm postgres_bambino_volume

docker-compose:
	docker-compose -f docker-compose.yaml up

#run:
#	go run bambino.go
#
#build:
#	go build bambino.go