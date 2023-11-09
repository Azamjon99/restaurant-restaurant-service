PWD=$(shell pwd)
SERVICE=restaurant-restaurant-svc
MIGRATION_PATH=${PWD}/src/infrastructure/migrations
PROTOS_PATH=$(PWD)/src/infrastructure/protos

server:
	go run main.go

migration:
	migrate create -ext sql -dir pkg/database/migrations -seq $(table)

migrateup:
	migrate -database "postgres://postgres:postgres@localhost:5432/restaurant_restaurant?sslmode=disable&search_path=public" -path ./src/infrastructure/migrations up

migratedown:
	migrate -database "postgres://postgres:postgres@localhost:5432/restaurant_restaurant?sslmode=disable&search_path=public" -path ./src/infrastructure/migrations down

add-protos-submodule:
	git submodule add https://github.com/Azamjon99/services-proto.git ./src/infrastructure/protos

pull-protos-submodule:
	git submodule update --recursive --remote

gen-restaurant-restaurant:
	protoc \
	--go_out=./src/application/protos \
	--go_opt=paths=import \
	--go-grpc_out=./src/application/protos \
	--go-grpc_opt=paths=import \
	-I=$(PROTOS_PATH)/restaurant_restaurant \
	$(PROTOS_PATH)/restaurant_restaurant/*.proto

docker:
	docker build --rm -t restaurant-restaurant-svc -f ${PWD}/deploy/Dockerfile .

compose-up:
	docker-compose -f ./deploy/docker-compose.yml up

compose-down:
	docker-compose -f ./deploy/docker-compose.yml down