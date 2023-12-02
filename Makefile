migrateup:
	migrate -path cmd/migrations -database "postgresql://postgres:postgres@localhost:5432/vol10?sslmode=disable" -verbose up

migrateup1:
	migrate -path cmd/migrations -database "postgresql://postgres:postgres@localhost:5432/vol10?sslmode=disable" -verbose up 1

migratedown:
	migrate -path cmd/migrations -database "postgresql://postgres:postgres@localhost:5432/vol10?sslmode=disable" -verbose down

migratedown1:
	migrate -path cmd/migrations -database "postgresql://postgres:postgres@localhost:5432/vol10?sslmode=disable" -verbose down 1

run:
	sh cmd/scripts/start.sh

down:
	docker-compose down

newnetwork:
	docker network create vol10-networks

.PHONY: migrateup migrateup1 migratedown migratedown1 run down newnetwork