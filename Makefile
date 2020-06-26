export:
	@while read LINE; do new-item -path . "$LINE"; done < go.env

docker:
	@docker-compose down
	@docker-compose build
	@docker-compose up -d

dockerdown:
	@docker-compose down

build:
	@echo "---- Building Application ----"
	@go build -o go-workshop-bin *.go

run:
	@echo "---- Running Application ----"
	@go run *.go