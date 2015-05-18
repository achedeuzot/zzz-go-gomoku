################################################################################
#         ____                                                                 #
#   _||__|  |  ______   ______   ______       HOMELESS DIGITAL LOCOMOTIVE      #
#  (        | |  H   | |  D   | |  L   |                                       #
#  /-()---() ~ ()--() ~ ()--() ~ ()--()             ksever & vletarou          #
################################################################################

GOPATH := ${PWD}:${PWD}/_vendor:${GOPATH}
export GOPATH

NAME = gomoku

default: ./bin/$(NAME)

./bin/$(NAME):
	go vet ./src/...
	go build -v -o ./bin/$(NAME) ./src/$(NAME)

re: clean ./bin/$(NAME)

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

clean:
	rm -f ./bin/$(NAME)

vendor_clean:
	rm -dvRf ./_vendor/src

vendor_get: vendor_clean
	GOPATH=${PWD}/_vendor go get -u -v \
	github.com/davecheney/profile \
	github.com/veandco/go-sdl2/sdl \
	github.com/veandco/go-sdl2/sdl_mixer \
	github.com/veandco/go-sdl2/sdl_image \
	github.com/veandco/go-sdl2/sdl_ttf

vendor_update: vendor_get
	rm -rf `find ./_vendor/src -type d -name .git` \
	&& rm -rf `find ./_vendor/src -type d -name .hg` \
	&& rm -rf `find ./_vendor/src -type d -name .bzr` \

.PHONY: build doc fmt lint run test clean vendor_clean vendor_get vendor_update

# This project requires the sdl2 library to be installed
# brew install sdl2{,_image,_ttf,_mixer}