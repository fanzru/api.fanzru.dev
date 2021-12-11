package api_fanzru_dev
import (
	"api.fanzru.dev/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"api.fanzru.dev/routes"
)

func main() {
	e := routes.Init()

	config := config.GetDatabaseConfig()
	_ = mgm.SetDefaultConfig(nil, "myFirstDatabase", options.Client().ApplyURI(config.URL))

	e.Logger.Fatal(e.Start(":5000"))
}

