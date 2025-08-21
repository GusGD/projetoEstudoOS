package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de sistema")
	}

	// Configurar modo do Gin
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inicializar servidor
	server := gin.Default()

	// Configurar rotas
	setupRoutes(server)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando na porta %s", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}

func setupRoutes(server *gin.Engine) {
	// Grupo de rotas da API
	api := server.Group("/api/v1")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "service": "ordem-servicos"})
		})

		// TODO: Implementar rotas dos módulos
		// auth := api.Group("/auth")
		// os := api.Group("/os")
		// users := api.Group("/users")
		// notifications := api.Group("/notifications")
	}
}
