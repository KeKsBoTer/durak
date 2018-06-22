
app = durak
package = github.com/KeKsBoTer/$(app)

container-name = durak-db
db-dir = /Users/simon/Go/src/github.com/KeKsBoTer/durak/db


all: build mongo-up run

run:
	./$(app)

build:
	go build -o $(app) $(package)/cmd

mongo-up: 
	bash tools/start-db.sh $(container-name) $(db-dir)
	
mongo-down:
	docker stop $(container-name)

.ONESHELL: mongo-up