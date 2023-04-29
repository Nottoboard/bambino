create-volumes:
	docker volume create --name=fileserver_db

remove-volumes:
	docker rm fileserver_db

docker-compose:
	docker-compose -f docker-compose.yaml up

run:
	go run bambino.go

build:
	go build bambino.go