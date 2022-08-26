build:
	@go build -o twitterlike ./cmd

build-image:
	@docker build -t twitterlike -f cmd/Dockerfile .
