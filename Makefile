genuser:
	sh scripts/user.sh

genproject:
	sh scripts/project.sh

genmember:
	sh scripts/member.sh

genimage:
	sh scripts/image.sh

runuser:
	go run cmd/app/user-service/main.go

runproject:
	go run cmd/app/project-service/main.go

runmember:
	go run cmd/app/member-service/main.go

runimage:
	go run cmd/app/image-service/main.go

test:
	go test -v ./... --cover

.PHONY: genuser genproject genmember runuser runproject runmember test genimage runimage