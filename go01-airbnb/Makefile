.PHONY: createdb migrate_up migrate_down

# Go migrate
create_mysql:
	docker run --name go01-mysql -e MYSQL_ROOT_PASSWORD=1234 -p 3307:3306 -d mysql:latest

migrate_up:
	migrate -path migrations -database "mysql://root:1234@tcp(127.0.0.1:3307)/go01-airbnb" up 1

migrate_down:
	migrate -path migrations -database "mysql://root:1234@tcp(127.0.0.1:3307)/go01-airbnb" down 1

# Main
run:
	go run main.go

build:
	go build main.go