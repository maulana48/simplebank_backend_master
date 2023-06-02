DB_URL=postgresql://postgres:root@localhost:5432/backend-master?sslmode=disable

DB_URL_LOCAL=postgresql://postgres:root@172.17.0.2:5432/backend-master?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=root -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres backend-master

dropdb:
	docker exec -it postgres dropdb backend-master

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateup_local:
	/usr/local/bin/migrate -path db/migration -database "postgresql://postgres:root@172.17.0.2:5432/backend-master?sslmode=disable" -verbose up

migratedown_local:
	/usr/local/bin/migrate -path db/migration -database "postgresql://postgres:root@172.17.0.2:5432/backend-master?sslmode=disable" -verbose down
	
server:
	clear && go run main.go

server1:
	go run main.go

sqlc:
	sqlc generate

mockdb:
	mockgen -package mockdb -destination db/mock/store.go github.com/maulana48/backend_master_class/simplebank/db/sqlc Store

test:
	go test -v -cover ./...

test_local:
	clear && go test -v ./...

.PHONY:
	postgres createdb dropdb migrateup migratedown sqlc test server mockdb
