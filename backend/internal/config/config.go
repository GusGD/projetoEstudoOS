package config

import (
	"os"
	"strconv"
)

// Config representa a configuração global do sistema
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	RabbitMQ RabbitMQConfig
	JWT      JWTConfig
}

// ServerConfig configurações do servidor
type ServerConfig struct {
	Port string
	Mode string
}

// DatabaseConfig configurações do banco de dados
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// RedisConfig configurações do Redis
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// RabbitMQConfig configurações do RabbitMQ
type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

// JWTConfig configurações do JWT
type JWTConfig struct {
	Secret     string
	Expiration int // em horas
}

var instance *Config

// GetInstance retorna a instância única da configuração (Singleton)
func GetInstance() *Config {
	if instance == nil {
		instance = &Config{
			Server: ServerConfig{
				Port: getEnv("PORT", "8080"),
				Mode: getEnv("GIN_MODE", "debug"),
			},
			Database: DatabaseConfig{
				Host:     getEnv("DB_HOST", "localhost"),
				Port:     getEnv("DB_PORT", "5432"),
				User:     getEnv("DB_USER", "postgres"),
				Password: getEnv("DB_PASSWORD", ""),
				DBName:   getEnv("DB_NAME", "ordem_servicos"),
				SSLMode:  getEnv("DB_SSLMODE", "disable"),
			},
			Redis: RedisConfig{
				Host:     getEnv("REDIS_HOST", "localhost"),
				Port:     getEnv("REDIS_PORT", "6379"),
				Password: getEnv("REDIS_PASSWORD", ""),
				DB:       getEnvAsInt("REDIS_DB", 0),
			},
			RabbitMQ: RabbitMQConfig{
				Host:     getEnv("RABBITMQ_HOST", "localhost"),
				Port:     getEnv("RABBITMQ_PORT", "5672"),
				User:     getEnv("RABBITMQ_USER", "guest"),
				Password: getEnv("RABBITMQ_PASSWORD", "guest"),
			},
			JWT: JWTConfig{
				Secret:     getEnv("JWT_SECRET", "your-secret-key"),
				Expiration: getEnvAsInt("JWT_EXPIRATION", 24),
			},
		}
	}
	return instance
}

// getEnv obtém variável de ambiente com valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt obtém variável de ambiente como inteiro com valor padrão
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
