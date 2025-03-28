.DEFAULT_GOAL := help

.PHONY: help build run

help: ## コマンド一覧を表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## コンテナ内に入る
	docker run -it --rm \
		-v $(PWD):/app \
		-w /app \
		golang:1.23.2-alpine sh

sh: ## なんか出てくる
	sh entrypoint.sh