package api

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type APIServer struct {
	App        *fiber.App
	Controller *Controller
	Address    string
	Port       string
}

func (api APIServer) Run() {
	api.App.Listen(api.Address + ":" + api.Port)
}

func NewServer(address, port string, con *gorm.DB) *APIServer {
	app := fiber.New()
	controller := &Controller{Con: con}

	api := &APIServer{
		App:        app,
		Controller: controller,
		Address:    address,
		Port:       port,
	}

	AddPlayerRoutes(api)
	AddCampaignRoutes(api)
	AddCollectionRoutes(api)
	AddCardRoutes(api)
	AddDeckRoutes(api)
	AddPlayerToPlayRoutes(api)
	AddCenarioRoutes(api)
	AddCenarioPlayerRoutes(api)
	AddCenarioCampaignRoutes(api)

	return api
}
