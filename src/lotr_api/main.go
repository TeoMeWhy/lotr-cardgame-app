package main

import (
	"lotr_api/api"
	"lotr_api/clients/db"
	"lotr_api/models"
)

func main() {

	dataPath := "../../data/database.db"
	con := db.GetSQliteClient(dataPath)

	con.AutoMigrate(
		&models.Campaign{},
		&models.Player{},
		&models.PlayerToPlay{},
		&models.Collection{},
		&models.Card{},
		&models.Deck{},
		&models.Cenario{},
		&models.CenarioPlayer{},
		&models.CenarioCampaign{},
	)

	server := api.NewServer("0.0.0.0", "3000", con)
	server.Run()
}
