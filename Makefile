.PHONY: ui
ui:
	cd ui/ && \
		npm run build

build:
	@go build -o bin/muscles -ldflags "-w -s" cmd/web/*
build_min:
	@CGO_ENABLED=0 GOOS=linux go build  -o bin/muscles -ldflags "-w -s" -a -installsuffix cgo cmd/web/*
run:  build
	./bin/muscles

clean:
	rm -rf ./bin
	rm -rf ./ui/public/

all: ui run 
# temp:
# 	go run cmd/web/!(*_test).go

test:
	@go test -v -cover ./...

up:
	sudo docker-compose up -d

down:
	sudo docker-compose down

running:
	sudo docker ps

