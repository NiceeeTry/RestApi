Image: 
	docker pull postgres
	
Container:
	docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
Migrate:
	migrate create -ext sql -dir ./schema -seq init

Excec:
	docker exec -it todo-db bash
# psql -U postgres