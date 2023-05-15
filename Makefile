postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=root -d postgres:12-alpine
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres backend-master
dropdb:
	docker exec -it postgres dropdb backend-master
migrateup:
	/usr/local/bin/migrate -path db/migration -database "postgresql://postgres:root@172.17.0.2:5432/backend-master?sslmode=disable" -verbose up
migratedown:
	/usr/local/bin/migrate -path db/migration -database "postgresql://postgres:root@172.17.0.2:5432/backend-master?sslmode=disable" -verbose down
sqlc:
	sqlc generate
testt:
	clear && go test -v ./...
.PHONY:
	postgres createdb dropdb migrateup migratedown sqlc
