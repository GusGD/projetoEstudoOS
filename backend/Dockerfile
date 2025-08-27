# Multi-stage build para otimização
FROM golang:1.21-alpine AS builder

# Instalar dependências necessárias
RUN apk add --no-cache git ca-certificates tzdata

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

# Imagem final
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk --no-cache add ca-certificates tzdata

# Criar usuário não-root
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Definir diretório de trabalho
WORKDIR /app

# Copiar binário do builder
COPY --from=builder /app/server .

# Copiar migrations
COPY --from=builder /app/migrations ./migrations

# Definir permissões
RUN chown -R appuser:appgroup /app
USER appuser

# Expor porta
EXPOSE 8080

# Comando de execução
CMD ["./server"]
