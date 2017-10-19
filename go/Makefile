export GOPATH=${PWD}/_vendor

default: build

build: local-compile

compile: local-compile linux-compile pi0-compile pi3-compile

linux-compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/can-i-poop.linux ./src

pi0-compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -a -installsuffix cgo -o bin/can-i-poop.pi0 ./src

pi3-compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -installsuffix cgo -o bin/can-i-poop.pi3 ./src

install-dependencies:
	go get github.com/namsral/flag
	go get github.com/ereyes01/firebase

local-compile:
	go build -a -o bin/can-i-poop.local ./src

local-run:
	bin/can-i-poop.local -config config/local.conf
