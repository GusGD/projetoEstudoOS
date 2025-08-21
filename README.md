# Sistema de Ordem de Serviços

Sistema completo para gestão de ordens de serviço em prefeituras, desenvolvido com Go (Backend) e Vue 3 (Frontend).

## 🏗️ Arquitetura

### Backend (Go)
- **Padrões implementados**: Repository, Service Layer, Event Sourcing, Observer, State Pattern
- **Banco de dados**: PostgreSQL com Event Sourcing para auditoria completa
- **Cache**: Redis para performance
- **Mensageria**: RabbitMQ para comunicação assíncrona
- **Autenticação**: JWT com RBAC (Role-Based Access Control)

### Frontend (Vue 3)
- **Padrões implementados**: Component-based, Pinia Store, Composition API
- **Roteamento**: Vue Router com layouts separados
- **Estado**: Gerenciamento centralizado com Pinia
- **UI**: Componentes reutilizáveis organizados por domínio

## 🚀 Estrutura do Projeto

```
OrdemServicos/
├── backend/                 # Backend em Go
│   ├── cmd/server/         # Ponto de entrada da aplicação
│   ├── internal/           # Código interno do sistema
│   │   ├── auth/           # Autenticação e RBAC
│   │   ├── os/             # Lógica de Ordem de Serviço
│   │   ├── users/          # Gestão de usuários
│   │   ├── notifications/  # Sistema de notificações
│   │   ├── config/         # Configurações (Singleton)
│   │   ├── shared/         # Código compartilhado entre domínios
│   │   └── infrastructure/ # Camada de infraestrutura
│   ├── pkg/                # Bibliotecas utilitárias
│   ├── migrations/         # Migrations do PostgreSQL
│   └── scripts/            # Scripts de automação
└── frontend/               # Frontend em Vue 3
    ├── src/
    │   ├── components/     # Componentes organizados por domínio
    │   ├── store/          # Gerenciamento de estado (Pinia)
    │   ├── api/            # Serviços de API (Facade Pattern)
    │   ├── composables/    # Hooks reutilizáveis
    │   └── layouts/        # Layouts da aplicação
    └── public/             # Arquivos estáticos
```

## 🛠️ Tecnologias

### Backend
- **Go 1.21+** - Linguagem principal
- **Gin** - Framework web
- **PostgreSQL** - Banco de dados principal
- **Redis** - Cache e sessões
- **RabbitMQ** - Mensageria
- **JWT** - Autenticação

### Frontend
- **Vue 3** - Framework JavaScript
- **Pinia** - Gerenciamento de estado
- **Vue Router** - Roteamento
- **Axios** - Cliente HTTP
- **Vite** - Build tool

## 📋 Pré-requisitos

- Go 1.21+
- Node.js 18+
- PostgreSQL 14+
- Redis 6+
- RabbitMQ 3.8+

## 🚀 Instalação

### Backend

1. **Configurar variáveis de ambiente**
   ```bash
   cd backend
   cp env.example .env
   # Editar .env com suas configurações
   ```

2. **Instalar dependências**
   ```bash
   go mod tidy
   ```

3. **Executar migrations**
   ```bash
   # Conectar ao PostgreSQL e executar:
   psql -d ordem_servicos -f migrations/001_create_initial_tables.sql
   ```

4. **Executar aplicação**
   ```bash
   go run cmd/server/main.go
   ```

### Frontend

1. **Configurar variáveis de ambiente**
   ```bash
   cd frontend
   cp env.example .env
   # Editar .env com suas configurações
   ```

2. **Instalar dependências**
   ```bash
   npm install
   ```

3. **Executar em desenvolvimento**
   ```bash
   npm run dev
   ```
## 📊 Funcionalidades Principais

### Gestão de OS
- ✅ Criação e edição de ordens de serviço
- ✅ Controle de status com histórico completo
- ✅ Sistema de prioridades (1-5)
- ✅ Atribuição de responsáveis
- ✅ Sistema de tags para categorização
- ✅ Event Sourcing para auditoria completa

### Usuários e Permissões
- ✅ Sistema de autenticação JWT
- ✅ RBAC (Role-Based Access Control)
- ✅ Perfis customizáveis por prefeitura
- ✅ Controle granular de permissões

### Notificações
- ✅ Sistema de notificações em tempo real
- ✅ Múltiplos canais (email, SMS, push)
- ✅ Histórico de notificações

### Relatórios e Analytics
- ✅ Dashboard com métricas em tempo real
- ✅ Relatórios exportáveis (PDF, Excel)
- ✅ Filtros avançados por status, prioridade, responsável

## 🧪 Testes

### Backend
```bash
cd backend
go test ./...
```

### Frontend
```bash
cd frontend
npm run test
```

## 📦 Deploy

### Docker (Recomendado)
```bash
docker-compose up -d
```

### Manual
1. Build do backend: `go build -o bin/server cmd/server/main.go`
2. Build do frontend: `npm run build`
3. Configurar servidor web (nginx/apache)
4. Configurar SSL/TLS

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.


## 🗺️ Roadmap

- [ ] API GraphQL
- [ ] Sistema de backup automático
- [ ] Integração com sistemas externos (ERPs)
- [ ] App mobile (React Native)
- [ ] Machine Learning para priorização automática
- [ ] Sistema de gamificação para funcionários
