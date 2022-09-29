.PHONY: frontend-install frontend-build build install 

frontend-install:
	cd src/frontend && yarn

frontend-build:
	cd src/frontend && yarn build
	rm -Rf dist/frontend/*.map


build:
	go build -o dist/nodebook .

# install: 
# 	go install .
# 	@echo "nodebook built and installed."
