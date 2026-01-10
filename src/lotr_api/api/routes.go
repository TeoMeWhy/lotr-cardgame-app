package api

func AddPlayerRoutes(app *APIServer) {
	playerRoutes := app.App.Group("/players")
	playerRoutes.Get("/", app.Controller.GetPlayers)
	playerRoutes.Post("/", app.Controller.CreatePlayer)
	playerRoutes.Put("/", app.Controller.PutPlayer)
	playerRoutes.Delete("/", app.Controller.DeletePlayer)
}

func AddCampaignRoutes(app *APIServer) {
	campaignRoutes := app.App.Group("/campaigns")
	campaignRoutes.Get("/", app.Controller.GetCampaigns)
	campaignRoutes.Post("/", app.Controller.CreateCampaign)
	campaignRoutes.Put("/", app.Controller.PutCampaign)
	campaignRoutes.Delete("/", app.Controller.DeleteCampaign)
}

func AddCollectionRoutes(app *APIServer) {
	collectionRoutes := app.App.Group("/collections")
	collectionRoutes.Get("/", app.Controller.GetCollections)
	collectionRoutes.Post("/", app.Controller.CreateCollection)
	collectionRoutes.Put("/", app.Controller.PutCollection)
}

func AddCardRoutes(app *APIServer) {
	cardRoutes := app.App.Group("/cards")
	cardRoutes.Get("/", app.Controller.GetCards)
	cardRoutes.Post("/", app.Controller.CreateCard)
	cardRoutes.Put("/", app.Controller.PutCard)
	cardRoutes.Delete("/", app.Controller.DeleteCard)
}

func AddDeckRoutes(app *APIServer) {
	deckRoutes := app.App.Group("/decks")
	deckRoutes.Get("/", app.Controller.GetDecks)
	deckRoutes.Post("/", app.Controller.CreateDeck)
	deckRoutes.Put("/", app.Controller.PutDeck)
	deckRoutes.Delete("/", app.Controller.DeleteDeck)
}

func AddPlayerToPlayRoutes(app *APIServer) {
	playerToPlayRoutes := app.App.Group("/player_to_play")
	playerToPlayRoutes.Get("/", app.Controller.GetPlayerToPlays)
	playerToPlayRoutes.Post("/", app.Controller.CreatePlayerToPlay)
	playerToPlayRoutes.Put("/", app.Controller.PutPlayerToPlay)
}

func AddCenarioRoutes(app *APIServer) {
	cenarioRoutes := app.App.Group("/cenarios")
	cenarioRoutes.Get("/", app.Controller.GetCenarios)
	cenarioRoutes.Post("/", app.Controller.CreateCenario)
	cenarioRoutes.Put("/", app.Controller.PutCenario)
	cenarioRoutes.Delete("/", app.Controller.DeleteCenario)
}

func AddCenarioPlayerRoutes(app *APIServer) {
	cenarioPlayerRoutes := app.App.Group("/cenario_players")
	cenarioPlayerRoutes.Get("/", app.Controller.GetCenarioPlayers)
	cenarioPlayerRoutes.Post("/", app.Controller.CreateCenarioPlayer)
	cenarioPlayerRoutes.Put("/", app.Controller.PutCenarioPlayer)
}

func AddCenarioCampaignRoutes(app *APIServer) {
	cenarioCampaignRoutes := app.App.Group("/cenario_campaigns")
	cenarioCampaignRoutes.Get("/", app.Controller.GetCenarioCampaigns)
	cenarioCampaignRoutes.Post("/", app.Controller.CreateCenarioCampaign)
	cenarioCampaignRoutes.Put("/", app.Controller.PutCenarioCampaign)
	cenarioCampaignRoutes.Delete("/", app.Controller.DeleteCenarioCampaign)
}
