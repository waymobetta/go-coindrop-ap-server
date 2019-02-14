all: deploy

.PHONY: test
test:
	go test -v ./...

.PHONY: test/auth
test/auth:
	go test -v auth/*.go

.PHONY: start
start:
	@go run cmd/start.go

start/local:
	@(. .env.local && go run cmd/start.go)

start/staging:
	@(. .env.staging && go run cmd/start.go)

start/prod:
	@(. .env.prod && go run cmd/start.go)

.PHONY: start/docs
start/docs:
	@(cd web/documentation && python -m SimpleHTTPServer 8000)

deploy: build compress
	@echo "deploying..\n"
	@MAKE done

.PHONY: build
build:
	@echo "building.."
	@command go build -ldflags "-s -w" -o cmd/go-coindrop-ap-server cmd/start.go

.PHONY: compress
compress:
	@echo "compressing.."
	@command upx cmd/go-coindrop-ap-server

.PHONY: done
done:
	@echo "done"

.PHONY: goa
goa:
	@goagen bootstrap -d github.com/waymobetta/go-coindrop-ap-server/design
	@rm main.go
	@rm paycheck.go
	@rm healthcheck.go
	@MAKE swagger

.PHONY: swagger
swagger:
	@echo "generating swagger spec"
	@(goagen swagger -d github.com/waymobetta/go-coindrop-ap-server/design && cp swagger/swagger.json web/documentation/swagger.json)
