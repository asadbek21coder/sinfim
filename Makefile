COMPOSE ?= docker compose
LOCAL_COMPOSE := docker-compose.local.yml

.PHONY: help local-up local-down local-down-volumes local-restart local-build local-logs local-ps backend-shell frontend-shell

help:
	@echo "Sinfim.uz local Docker commands"
	@echo ""
	@echo "  make local-up            Build and start frontend, backend, and local services"
	@echo "  make local-down          Stop the local stack"
	@echo "  make local-down-volumes  Stop the local stack and remove local Docker volumes"
	@echo "  make local-restart       Restart the local stack"
	@echo "  make local-build         Build local Docker images"
	@echo "  make local-logs          Follow local stack logs"
	@echo "  make local-ps            Show local stack status"
	@echo "  make backend-shell       Open a shell in the backend container"
	@echo "  make frontend-shell      Open a shell in the frontend container"

local-up:
	$(COMPOSE) -f $(LOCAL_COMPOSE) up -d --build

local-down:
	$(COMPOSE) -f $(LOCAL_COMPOSE) down --remove-orphans

local-down-volumes:
	$(COMPOSE) -f $(LOCAL_COMPOSE) down --remove-orphans --volumes

local-restart: local-down local-up

local-build:
	$(COMPOSE) -f $(LOCAL_COMPOSE) build

local-logs:
	$(COMPOSE) -f $(LOCAL_COMPOSE) logs -f

local-ps:
	$(COMPOSE) -f $(LOCAL_COMPOSE) ps

backend-shell:
	$(COMPOSE) -f $(LOCAL_COMPOSE) exec backend sh

frontend-shell:
	$(COMPOSE) -f $(LOCAL_COMPOSE) exec frontend sh
