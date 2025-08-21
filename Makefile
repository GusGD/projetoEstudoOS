# Makefile para Sistema de Ordem de Serviços
# Comandos para facilitar o desenvolvimento

.PHONY: help install build run test clean docker-up docker-down docker-build

# Variáveis
BACKEND_DIR = backend
FRONTEND_DIR = frontend
DOCKER_COMPOSE = docker-compose.yml

# Comando padrão
help:
	@echo "Sistema de Ordem de Serviços - Comandos disponíveis:"
	@echo ""
	@echo "🔧 Desenvolvimento:"
	@echo "  make install     - Instalar dependências (backend + frontend)"
	@echo "  make build       - Build do projeto completo"
	@echo "  make run         - Executar em desenvolvimento"
	@echo "  make test        - Executar testes"
	@echo ""
	@echo "🐳 Docker:"
	@echo "  make docker-up   - Subir serviços com Docker"
	@echo "  make docker-down - Parar serviços Docker"
	@echo "  make docker-build- Build das imagens Docker"
	@echo ""
	@echo "🧹 Manutenção:"
	@echo "  make clean       - Limpar arquivos temporários"
	@echo "  make reset       - Reset completo (cuidado!)"
	@echo ""
	@echo "📊 Status:"
	@echo "  make status      - Status dos serviços"
	@echo "  make logs        - Logs dos serviços"

# Instalação de dependências
install: install-backend install-frontend

install-backend:
	@echo "📦 Instalando dependências do Backend..."
	cd $(BACKEND_DIR) && go mod tidy
	@echo "✅ Backend instalado com sucesso!"

install-frontend:
	@echo "📦 Instalando dependências do Frontend..."
	cd $(FRONTEND_DIR) && npm install
	@echo "📝 Configurando variáveis de ambiente..."
	cd $(FRONTEND_DIR) && cp env.example .env
	@echo "✅ Frontend instalado com sucesso!"

# Build do projeto
build: build-backend build-frontend

build-backend:
	@echo "🔨 Build do Backend..."
	cd $(BACKEND_DIR) && go build -o bin/server cmd/server/main.go
	@echo "✅ Backend buildado com sucesso!"

build-frontend:
	@echo "🔨 Build do Frontend..."
	cd $(FRONTEND_DIR) && npm run build
	@echo "✅ Frontend buildado com sucesso!"

# Execução em desenvolvimento
run: run-backend run-frontend

run-backend:
	@echo "🚀 Executando Backend..."
	cd $(BACKEND_DIR) && go run cmd/server/main.go

run-frontend:
	@echo "🚀 Executando Frontend..."
	cd $(FRONTEND_DIR) && npm run dev

# Testes
test: test-backend test-frontend

test-backend:
	@echo "🧪 Executando testes do Backend..."
	cd $(BACKEND_DIR) && go test ./...

test-frontend:
	@echo "🧪 Executando testes do Frontend..."
	cd $(FRONTEND_DIR) && npm run test

# Docker
docker-up:
	@echo "🐳 Subindo serviços Docker..."
	docker-compose -f $(DOCKER_COMPOSE) up -d
	@echo "✅ Serviços Docker iniciados!"
	@echo "🌐 Acesse: http://localhost:3000 (Frontend) | http://localhost:8080 (Backend)"

docker-down:
	@echo "🐳 Parando serviços Docker..."
	docker-compose -f $(DOCKER_COMPOSE) down
	@echo "✅ Serviços Docker parados!"

docker-build:
	@echo "🔨 Build das imagens Docker..."
	docker-compose -f $(DOCKER_COMPOSE) build
	@echo "✅ Imagens Docker buildadas!"

docker-logs:
	@echo "📋 Logs dos serviços Docker..."
	docker-compose -f $(DOCKER_COMPOSE) logs -f

# Status e logs
status:
	@echo "📊 Status dos serviços:"
	docker-compose -f $(DOCKER_COMPOSE) ps

logs:
	@echo "📋 Logs dos serviços:"
	docker-compose -f $(DOCKER_COMPOSE) logs

# Limpeza
clean:
	@echo "🧹 Limpando arquivos temporários..."
	cd $(BACKEND_DIR) && go clean
	cd $(FRONTEND_DIR) && rm -rf node_modules dist
	@echo "✅ Limpeza concluída!"

reset: docker-down clean
	@echo "🔄 Reset completo executado!"
	@echo "💡 Execute 'make install' para reinstalar dependências"

# Desenvolvimento específico
dev-backend:
	@echo "🔧 Modo desenvolvimento Backend..."
	cd $(BACKEND_DIR) && air

dev-frontend:
	@echo "🔧 Modo desenvolvimento Frontend..."
	cd $(FRONTEND_DIR) && npm run dev

# Migrations
migrate:
	@echo "🗄️ Executando migrations..."
	cd $(BACKEND_DIR) && go run cmd/migrate/main.go

# Seed de dados
seed:
	@echo "🌱 Executando seed de dados..."
	cd $(BACKEND_DIR) && go run cmd/seed/main.go

# Linting
lint: lint-backend lint-frontend

lint-backend:
	@echo "🔍 Linting do Backend..."
	cd $(BACKEND_DIR) && golangci-lint run

lint-frontend:
	@echo "🔍 Linting do Frontend..."
	cd $(FRONTEND_DIR) && npm run lint

# Formatação
format: format-backend format-frontend

format-backend:
	@echo "✨ Formatando código Go..."
	cd $(BACKEND_DIR) && go fmt ./...

format-frontend:
	@echo "✨ Formatando código Vue..."
	cd $(FRONTEND_DIR) && npm run format

# Health check
health:
	@echo "🏥 Verificando saúde dos serviços..."
	@echo "Backend: $(shell curl -s http://localhost:8080/api/v1/health || echo "❌ Offline")"
	@echo "Frontend: $(shell curl -s http://localhost:3000 || echo "❌ Offline")"
	@echo "PostgreSQL: $(shell docker exec ordem_servicos_postgres pg_isready -U postgres || echo "❌ Offline")"
	@echo "Redis: $(shell docker exec ordem_servicos_redis redis-cli ping || echo "❌ Offline")"
	@echo "RabbitMQ: $(shell docker exec ordem_servicos_rabbitmq rabbitmq-diagnostics ping || echo "❌ Offline")"

# Backup
backup:
	@echo "💾 Criando backup do banco..."
	docker exec ordem_servicos_postgres pg_dump -U postgres ordem_servicos > backup_$(shell date +%Y%m%d_%H%M%S).sql
	@echo "✅ Backup criado com sucesso!"

# Restore
restore:
	@echo "📥 Restaurando backup..."
	@read -p "Digite o nome do arquivo de backup: " backup_file; \
	docker exec -i ordem_servicos_postgres psql -U postgres ordem_servicos < $$backup_file
	@echo "✅ Backup restaurado com sucesso!"
