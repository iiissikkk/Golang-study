include .env
export

service-run:
	go run main.go

migrate-up:
	migrate -path /home/isik/GolandProjects/awesomeProject/3\ part/migration/ -database ${CONN_STRING} up

migrate-down:
	migrate -path /home/isik/GolandProjects/awesomeProject/3\ part/migration/ -database ${CONN_STRING} down
