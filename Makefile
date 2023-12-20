dev:
	go run main.go

up:
	docker-compose up -d

down:
	docker-compose down

ammend:
	git add --all && git commit --amend --no-edit && git push origin main -f

.PHONY: dev up down ammend