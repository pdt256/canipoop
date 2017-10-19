export GOPATH=${PWD}/_vendor

default: build

build: compile

compile: linux-compile local-compile

linux-compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/can-i-poop.linux ./src

install-dependencies:
	go get github.com/namsral/flag
	go get github.com/ereyes01/firebase

local-compile:
	go build -a -o bin/can-i-poop.local ./src

local-run:
	bin/can-i-poop.local -config config/local.conf
