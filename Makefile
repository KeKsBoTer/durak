
app = durak
package = github.com/KeKsBoTer/$(app)

container-name = durak-db
db-dir = ./db


all: build run

run:
	./$(app)

build:
	go build -o $(app) $(package)/cmd

deps:
	go get ./...

mongo-up: 
	bash tools/start-db.sh $(container-name) $(db-dir)
	
mongo-down:
	docker stop $(container-name)

.ONESHELL: mongo-up