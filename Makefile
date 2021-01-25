git:
	git add .
	git commit -m "$(msg)"
	git push origin master

postgres: 
	docker-compose -f ./.devEnvironment/postgres.yml up -d

start:
	go run main.go

.PHONY: git postgres main