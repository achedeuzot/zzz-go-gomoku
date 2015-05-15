################################################################################
#         ____                                                                 #
#   _||__|  |  ______   ______   ______       HOMELESS DIGITAL LOCOMOTIVE      #
#  (        | |  H   | |  D   | |  L   |                                       #
#  /-()---() ~ ()--() ~ ()--() ~ ()--()             ksever & vletarou          #
################################################################################

GOPATH := ${PWD}:${PWD}/_vendor:${GOPATH}
export GOPATH

NAME = gomoku

default: build

build: vet
	go build -v -o ./bin/$(NAME) ./src/$(NAME)

doc:
	godoc -http=:6060 -index

fmt:
	go fmt ./src/...

run: build
	./bin/$(NAME)

test:
	go test ./src/...

profile:
	go build -v -o ./bin/$(NAME) ./src/$(NAME)
	GODEBUG=gctrace=1 ./bin/$(NAME) -prof=1

vet:
	go vet ./src/...

clean:
	rm ./bin/$(NAME)

vendor_clean:
	rm -dvRf ./_vendor/src

vendor_get: vendor_clean
	GOPATH=${PWD}/_vendor go get -d -u -v \
	github.com/davecheney/profile

vendor_update: vendor_get
	rm -rf `find ./_vendor/src -type d -name .git` \
	&& rm -rf `find ./_vendor/src -type d -name .hg` \
	&& rm -rf `find ./_vendor/src -type d -name .bzr` \

.PHONY: build doc fmt lint run test vet clean
