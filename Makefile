
.PHONY: local
local:
	docker-compose -f ./deployments/docker-compose.yml up

.PHONY: local-down
local-down:
	docker-compose -f ./deployments/docker-compose.yml down

