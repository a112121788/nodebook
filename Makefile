.PHONY: deps build install deps-frontend build-frontend embed-frontend

deps:
	go install github.com/markbates/pkger/cmd/pkger

embed-frontend:
	pkger

build: embed-frontend
	go build -o dist/nodebook .

install: embed-frontend
	go install .
	@echo "nodebook built and installed."

deps-frontend:
	cd src/frontend && yarn

build-frontend:
	cd src/frontend && yarn build
	rm -Rf dist/frontend/*.map
