// Load .env values.
package initializers

import (
	"log"

	"github.com/dotenv-org/godotenvvault"
)

func LoadEnv() {
	err := godotenvvault.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
