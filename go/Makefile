export GOPATH=${PWD}/_vendor

default: build

build: compile

compile: linux-compile mac-compile

linux-compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/can-i-poop.linux ./src

install-dependencies:
	go get github.com/namsral/flag
	go get github.com/ereyes01/firebase

mac-compile:
	go build -a -o bin/can-i-poop.mac ./src

mac-run:
	bin/can-i-poop.mac -config config/local.conf
