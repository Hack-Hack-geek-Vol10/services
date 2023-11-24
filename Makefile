genuser:
	sh scripts/user.sh

genproject:
	sh scripts/project.sh

runuser:
	go run cmd/app/user-service/main.go

runproject:
	go run cmd/app/project-service/main.go

test:
	go test -v ./... --cover

.PHONY: genuser genproject runuser runproject test