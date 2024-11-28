SERVER_MAIN=cmd/server/main.go
GENERATE_MAIN=cmd/generate/main.go
APP_NAME=go-web-template

ifeq ($(OS),Windows_NT)
	BIN_PATH = bin/$(APP_NAME).exe
else
	BIN_PATH = bin/$(APP_NAME)
endif

all: install gen build

gen-code:
	go run cmd/generate/main.go

install:
	go get github.com/swaggo/swag/cmd/swag@v1.16.3
	go install github.com/swaggo/swag/cmd/swag@v1.16.3
	go install github.com/cosmtrek/air@v1.51.0

gen:
	go generate ./...

watch:gen
	go run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "go build --tags dev -o ${BIN_PATH} ${SERVER_MAIN}" --build.bin "${BIN_PATH}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true \
		--screen.clear_on_rebuild true \
		--log.main_only true \

run:
	go run ${SERVER_MAIN}

clean:
	rm -rf bin/*

build:
	go build -o ${BIN_PATH} ${SERVER_MAIN}