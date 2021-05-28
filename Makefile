.PHONY: help
help: ## ドキュメントのヘルプを表示する。
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: up
up: ## app コンテナを起動する。
	docker-compose up

.PHONY: jmeter
jmeter: ## jmeter を単発で実行する。
	docker-compose -f docker-compose-jmeter.yml run --rm jmeter

.PHONY: jmeter-60
jmeter-60: ## jmeter で秒間100リクエストの負荷を60秒間かける。
	docker-compose -f docker-compose-jmeter.yml run --rm jmeter 6000 60
