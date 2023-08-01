createdb:
	docker exec -it postgres-docker createdb --username=root --owner=root simple_cashier

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_cashier?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_cashier?sslmode=disable" --verbose down

.PHONY: createdb migrateup migratedown