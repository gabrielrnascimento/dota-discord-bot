APP_NAME = "dota2-discord-bot"
BUILD_DIR = "build"

# Go commands
GO = go
GO_BUILD = $(GO) build
GO_CLEAN = $(GO) clean
GO_RUN = $(GO) run
GO_TEST = $(GO) test
GO_LINT = golangci-lint run

# Run the program
run:
	@echo "Running the program..."
	@$(GO_RUN) cmd/bot/main.go