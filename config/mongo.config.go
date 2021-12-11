package config
import (
	"github.com/joho/godotenv"
	"os"
)

type DatabaseConfig struct {
	URL string
}

func GetDatabaseConfig() DatabaseConfig{
	godotenv.Load(".env")

	return DatabaseConfig{URL: os.Getenv("DB_URL") }

}
