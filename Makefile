migrateup:
	migrate -path cmd/migrations -database "postgresql://postgres:postgres@localhost:5432/vol10?sslmode=disable" -verbose up

migrateup1:
	migrate -path cmd/migrations -database "postgresql://postgres:postgres@localhost:5432/vol10?sslmode=disable" -verbose up 1

migratedown:
	migrate -path cmd/migrations -database "postgresql://postgres:postgres@localhost:5432/vol10?sslmode=disable" -verbose down

migratedown1:
	migrate -path cmd/migrations -database "postgresql://postgres:postgres@localhost:5432/vol10?sslmode=disable" -verbose down 1

run:
	sh cmd/scripts/run_server.sh

test:
	go test -v ./... --cover

.PHONY: migrateup migrateup1 migratedown migratedown1