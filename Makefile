genproto:
	sh scripts/image.sh
	sh scripts/member.sh
	sh scripts/project.sh
	sh scripts/token.sh
	sh scripts/user.sh

runuser:
	go run cmd/app/user-service/main.go

runproject:
	go run cmd/app/project-service/main.go

runmember:
	go run cmd/app/member-service/main.go

runimage:
	go run cmd/app/image-service/main.go

runtoken:
	go run cmd/app/token-service/main.go

test:
	go test -v ./... --cover

.PHONY: genproto runuser runproject runmember runimage runtoken test