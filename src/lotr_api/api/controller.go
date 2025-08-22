package api

import (
	"log"
	"lotr_api/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Controller struct {
	Con *gorm.DB
}

func (c *Controller) GetPlayers(ctx fiber.Ctx) error {

	id := ctx.Query("id")
	name := ctx.Query("name")

	if id != "" {
		var player models.Player
		if err := c.Con.First(&player, "id = ?", id).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
		}

		return ctx.Status(fiber.StatusOK).JSON(player)
	}

	if name != "" {
		var player models.Player
		if err := c.Con.First(&player, "name = ?", name).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(player)
	}

	var players []models.Player
	if err := c.Con.Find(&players).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve players"})
	}
	return ctx.Status(fiber.StatusOK).JSON(players)
}

func (c *Controller) CreatePlayer(ctx fiber.Ctx) error {
	payload := models.Player{}
	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	player := models.NewPlayer(payload.Name)

	if err := c.Con.Create(&player).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create player"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(player)
}

func (c *Controller) PutPlayer(ctx fiber.Ctx) error {
	payload := models.Player{}
	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Con.Save(&payload).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create player"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(payload)
}

func (c *Controller) DeletePlayer(ctx fiber.Ctx) error {
	payload := &models.Player{}
	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Con.Delete(&payload).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete player"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"ok": "Player deleted successfully"})
}

func (c *Controller) GetCampaigns(ctx fiber.Ctx) error {

	loader := c.Con.Preload("Players.Player")
	loader = loader.Preload("Players.Heroes.Collection")
	loader = loader.Preload("Players.Deck.Cards.Collection")
	loader = loader.Preload("CenarioCampaigns.CenarioPlayers.Cenario")
	loader = loader.Preload("CenarioCampaigns.CenarioPlayers.Player")
	loader = loader.Preload("CenarioCampaigns.Cenario")
	loader = loader.Preload("DefeatedHeros.Collection")

	id := ctx.Query("id")
	name := ctx.Query("name")

	if id != "" {
		campaign := models.Campaign{}
		if err := loader.First(&campaign, "id = ?", id); err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Campaign not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(campaign)
	}

	if name != "" {
		campaign := models.Campaign{}
		if err := loader.First(&campaign, "name = ?", name); err != nil {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Campaign not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(campaign)
	}

	var campaigns []models.Campaign
	if err := loader.Find(&campaigns).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve campaigns"})
	}
	return ctx.Status(fiber.StatusOK).JSON(campaigns)
}

func (c *Controller) CreateCampaign(ctx fiber.Ctx) error {

	payload := models.Campaign{}

	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	campaign := models.NewCampaign(payload.Name, payload.Players)

	if err := c.Con.Create(&campaign).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create campaign"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(campaign)
}

func (c *Controller) PutCampaign(ctx fiber.Ctx) error {

	payload := &models.Campaign{}

	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	oldCampaign := &models.Campaign{}
	if err := c.Con.First(oldCampaign, "id = ?", payload.Id).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Campaign not found"})
	}

	if err := c.Con.Model(&oldCampaign).Association("DefeatedHeros").Clear(); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete defeated heroes"})
	}

	if err := c.Con.Save(&payload).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update campaign"})
	}

	return ctx.Status(fiber.StatusOK).JSON(payload)
}

func (c *Controller) DeleteCampaign(ctx fiber.Ctx) error {

	payload := &models.Campaign{}
	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Con.Delete(&payload, "id = ?", payload.Id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete campaign"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"ok": "Campaign deleted successfully"})
}

func (c *Controller) GetCollections(ctx fiber.Ctx) error {

	id := ctx.Query("id")

	if id != "" {
		collection := models.Collection{}
		if err := c.Con.First(&collection, "id = ?", id).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Collection not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(collection)
	}

	name := ctx.Query("name")
	if name != "" {
		collection := models.Collection{}
		if err := c.Con.First(&collection, "name = ?", name).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Collection not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(collection)
	}

	var collections []models.Collection
	if err := c.Con.Find(&collections).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve collections"})
	}
	return ctx.Status(fiber.StatusOK).JSON(collections)
}

func (c *Controller) CreateCollection(ctx fiber.Ctx) error {

	payload := models.Collection{}
	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	collection := models.NewCollection(payload.Name, payload.Description)

	if err := c.Con.Create(&collection).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create collection"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(collection)
}

func (c *Controller) GetCards(ctx fiber.Ctx) error {

	loader := c.Con.Preload("Collection")

	id := ctx.Query("id")
	if id != "" {
		loader = loader.Where("id = ?", id)
	}

	number := ctx.Query("number")
	if number != "" {
		loader = loader.Where("number = ?", number)
	}

	collection := ctx.Query("collection_id")
	if collection != "" {
		loader = loader.Where("collection_id = ?", collection)
	}

	var cards []models.Card
	if err := loader.Find(&cards).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cards"})
	}
	return ctx.Status(fiber.StatusOK).JSON(cards)
}

func (c *Controller) CreateCard(ctx fiber.Ctx) error {

	payload := models.Card{}
	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	card := models.NewCard(payload.Number, payload.CollectionID, payload.Name, payload.Description)
	if err := c.Con.Create(&card).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create card"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(card)
}

func (c *Controller) PutCard(ctx fiber.Ctx) error {

	payload := &models.Card{}
	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Con.Save(payload).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update card"})
	}

	return ctx.Status(fiber.StatusOK).JSON(payload)
}

func (c *Controller) DeleteCard(ctx fiber.Ctx) error {

	payload := &models.Card{}
	if err := ctx.Bind().Body(&payload); err != nil {
		log.Println("Error binding payload:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Con.Delete(&models.Card{}, "id = ?", payload.Id).Error; err != nil {
		log.Println("Error deleting card:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete card"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"ok": "Card deleted successfully"})
}

func (c *Controller) GetDecks(ctx fiber.Ctx) error {

	loader := c.Con.Preload("Cards").Preload("Cards.Collection")

	id := ctx.Query("id")
	if id != "" {
		deck := models.Deck{}
		if err := loader.First(&deck, "id = ?", id).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Deck not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(deck)
	}

	name := ctx.Query("name")
	if name != "" {
		deck := models.Deck{}
		if err := loader.First(&deck, "name = ?", name).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Deck not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(deck)
	}

	var decks []models.Deck
	if err := loader.Find(&decks).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve decks"})
	}
	return ctx.Status(fiber.StatusOK).JSON(decks)
}

func (c *Controller) CreateDeck(ctx fiber.Ctx) error {

	payload := models.Deck{}
	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	deck := models.NewDeck(payload.Name, payload.Description, payload.Cards)

	if err := c.Con.Create(&deck).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create deck"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(deck)
}

func (c *Controller) PutDeck(ctx fiber.Ctx) error {

	payload := models.Deck{}
	if err := ctx.Bind().Body(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Con.Save(&payload).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update deck"})
	}

	return ctx.Status(fiber.StatusOK).JSON(payload)
}

func (c *Controller) DeleteDeck(ctx fiber.Ctx) error {

	payload := &models.Deck{}
	if err := ctx.Bind().Body(&payload); err != nil {
		log.Println("Error binding payload:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Con.Delete(&models.Deck{}, "id = ?", payload.Id).Error; err != nil {
		log.Println("Error deleting deck:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete deck"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"ok": "Deck deleted successfully"})

}

func (c *Controller) GetPlayerToPlays(ctx fiber.Ctx) error {

	loader := c.Con.Preload("Player").Preload("Heroes").Preload("Deck")

	id := ctx.Query("id")
	if id != "" {
		playerToPlay := &models.PlayerToPlay{}
		if err := loader.First(playerToPlay, "id = ?", id).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player to play not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(playerToPlay)
	}

	var playerToPlays []models.PlayerToPlay
	player := ctx.Query("player")
	if player != "" {
		if err := loader.Find(&playerToPlays, "player_id = ?", player).Error; err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve player to plays"})
		}
		return ctx.Status(fiber.StatusOK).JSON(playerToPlays)
	}

	if err := loader.Find(&playerToPlays).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve player to plays"})
	}
	return ctx.Status(fiber.StatusOK).JSON(playerToPlays)
}

func (c *Controller) CreatePlayerToPlay(ctx fiber.Ctx) error {

	payload := &models.PlayerToPlay{}
	if err := ctx.Bind().Body(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	log.Println(payload)

	playerToPlay := models.NewPlayerToPlay(payload.Player, payload.Heroes, payload.Deck)
	if err := c.Con.Create(playerToPlay).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create player to play"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(playerToPlay)
}

func (c *Controller) PutPlayerToPlay(ctx fiber.Ctx) error {

	payload := &models.PlayerToPlay{}
	if err := ctx.Bind().Body(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if payload.ID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input (ID is required)"})
	}

	oldPlayerToPlay := &models.PlayerToPlay{}
	if err := c.Con.First(oldPlayerToPlay, "id = ?", payload.ID).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player to play not found"})
	}

	if err := c.Con.Model(&oldPlayerToPlay).Association("Heroes").Clear(); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update heroes on player to play"})
	}

	if err := c.Con.Save(&payload).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update player to play"})
	}

	return ctx.Status(fiber.StatusOK).JSON(payload)
}

func (c *Controller) GetCenarios(ctx fiber.Ctx) error {

	id := ctx.Query("id")
	if id != "" {
		cenario := &models.Cenario{}
		if err := c.Con.Preload("Collection").First(cenario, "id = ?", id).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cenario not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(cenario)
	}

	name := ctx.Query("name")
	if name != "" {
		cenario := &models.Cenario{}
		if err := c.Con.Preload("Collection").First("name = ?", name).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cenario not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(cenario)
	}

	cenarios := []models.Cenario{}
	if err := c.Con.Preload("Collection").Order("`order` ASC").Find(&cenarios).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cenarios"})
	}
	return ctx.Status(fiber.StatusOK).JSON(cenarios)
}

func (c *Controller) CreateCenario(ctx fiber.Ctx) error {

	payload := &models.Cenario{}
	if err := ctx.Bind().Body(payload); err != nil {
		log.Println("Error binding payload:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	cenario := models.NewCenario(payload.Name, payload.Description, payload.CollectionID)
	if err := c.Con.Create(cenario).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create cenario"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(cenario)
}

func (c *Controller) DeleteCenario(ctx fiber.Ctx) error {

	payload := &models.Cenario{}
	if err := ctx.Bind().Body(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.Con.Delete(payload, "id = ?", payload.Id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete cenario"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"ok": "Cenario deleted successfully"})
}

func (c *Controller) PutCenario(ctx fiber.Ctx) error {

	payload := &models.Cenario{}
	if err := ctx.Bind().Body(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if payload.Id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input (ID is required)"})
	}

	if err := c.Con.Save(payload).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update cenario"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(payload)
}

func (c *Controller) GetCenarioPlayers(ctx fiber.Ctx) error {

	loader := c.Con.Preload("Cenario")

	id := ctx.Query("id")
	if id != "" {
		cenarioPlayer := &models.CenarioPlayer{}
		if err := loader.First(cenarioPlayer, "id = ?", id).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cenario player not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(cenarioPlayer)
	}

	campaignID := ctx.Query("campaign_id")
	if campaignID != "" {
		loader = loader.Where("campaign_id = ?", campaignID)
	}

	playerID := ctx.Query("player_id")
	if playerID != "" {
		loader = loader.Where("player_id = ?", playerID)
	}

	var cenarioPlayers []models.CenarioPlayer
	if err := loader.Find(&cenarioPlayers).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cenario players"})
	}

	if len(cenarioPlayers) == 0 && (campaignID != "" || playerID != "") {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cenario players not found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(cenarioPlayers)
}

func (c *Controller) CreateCenarioPlayer(ctx fiber.Ctx) error {

	payload := &models.CenarioPlayer{}
	if err := ctx.Bind().Body(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	cenarioPlayer := models.NewCenarioPlayer(
		payload.CampaignID,
		payload.PlayerID,
		payload.CenarioID,
		payload.FinalThreat,
		payload.ThreatHeroDefeated,
		payload.HeroesDamage,
	)

	if err := c.Con.Create(cenarioPlayer).Error; err != nil {
		log.Println("Error creating cenario player:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create cenario player"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(cenarioPlayer)
}

func (c *Controller) PutCenarioPlayer(ctx fiber.Ctx) error {

	payload := &models.CenarioPlayer{}
	if err := ctx.Bind().Body(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if payload.Id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input (ID is required)"})
	}

	cenarioPlayer := models.NewCenarioPlayer(
		payload.CampaignID,
		payload.PlayerID,
		payload.CenarioID,
		payload.FinalThreat,
		payload.ThreatHeroDefeated,
		payload.HeroesDamage,
	)

	cenarioPlayer.Id = payload.Id

	if err := c.Con.Save(&cenarioPlayer).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update cenario player"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"ok": "Cenario player updated successfully"})
}

func (c *Controller) GetCenarioCampaigns(ctx fiber.Ctx) error {

	id := ctx.Query("id")
	if id != "" {
		cenarioCampaign := &models.CenarioCampaign{}
		if err := c.Con.First(cenarioCampaign, "id = ?", id).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cenario campaign not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(cenarioCampaign)
	}

	var cenarioCampaigns []models.CenarioCampaign

	campaignID := ctx.Query("campaign_id")
	if campaignID != "" {
		if err := c.Con.Where(cenarioCampaigns, "campaign_id = ?", campaignID).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cenario campaign not found"})
		}
		return ctx.Status(fiber.StatusOK).JSON(cenarioCampaigns)
	}

	if err := c.Con.Find(&cenarioCampaigns).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cenario campaigns"})
	}
	return ctx.Status(fiber.StatusOK).JSON(cenarioCampaigns)
}

func (c *Controller) CreateCenarioCampaign(ctx fiber.Ctx) error {

	payload := &models.CenarioCampaign{}
	if err := ctx.Bind().Body(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	cenarioPlayers := []models.CenarioPlayer{}
	loaderCenarioPlayers := c.Con.Where("campaign_id = ?", payload.CampaignID)
	loaderCenarioPlayers = loaderCenarioPlayers.Where("cenario_id = ?", payload.CenarioID)
	loaderCenarioPlayers = loaderCenarioPlayers.Find(&cenarioPlayers)
	if err := loaderCenarioPlayers.Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cenario players"})
	}

	cenarioCampaign := models.NewCenarioCampaign(
		payload.CampaignID,
		payload.CenarioID,
		cenarioPlayers,
		payload.VictoryPoints,
		payload.Rounds,
	)

	if err := c.Con.Create(cenarioCampaign).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create cenario campaign"})
	}

	campaign := &models.Campaign{}
	if err := c.Con.First(campaign, "id = ?", payload.CampaignID).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Campaign not found"})
	}

	campaign.Points += cenarioCampaign.TotalFinalPoints
	if err := c.Con.Save(campaign).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update campaign points"})
	}

	return ctx.Status(fiber.StatusOK).JSON(cenarioCampaign)
}

func (c *Controller) PutCenarioCampaign(ctx fiber.Ctx) error {

	payload := &models.CenarioCampaign{}
	if err := ctx.Bind().Body(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	cenarioPlayers := []models.CenarioPlayer{}
	loaderCenarioPlayers := c.Con.Where("campaign_id = ?", payload.CampaignID)
	loaderCenarioPlayers = loaderCenarioPlayers.Where("cenario_id = ?", payload.CenarioID)
	loaderCenarioPlayers = loaderCenarioPlayers.Find(&cenarioPlayers)
	if err := loaderCenarioPlayers.Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cenario players"})
	}

	cenarioCampaign := models.NewCenarioCampaign(
		payload.CampaignID,
		payload.CenarioID,
		cenarioPlayers,
		payload.VictoryPoints,
		payload.Rounds,
	)

	cenarioCampaign.Id = payload.Id

	if err := c.Con.Save(&cenarioCampaign).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create cenario campaign"})
	}

	campaign := &models.Campaign{}
	if err := c.Con.Preload("CenarioCampaigns").First(campaign, "id = ?", payload.CampaignID).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Campaign not found"})
	}

	points := 0
	for _, c := range campaign.CenarioCampaigns {
		points += c.TotalFinalPoints
	}

	campaign.Points = points
	if err := c.Con.Save(campaign).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update campaign points"})
	}

	return ctx.Status(fiber.StatusOK).JSON(cenarioCampaign)
}

func (c *Controller) DeleteCenarioCampaign(ctx fiber.Ctx) error {
	payload := &models.CenarioCampaign{}
	if err := ctx.Bind().Body(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	playersCenarios := []models.CenarioPlayer{}
	if err := c.Con.Delete(&playersCenarios, "campaign_id = ? AND cenario_id = ?", payload.CampaignID, payload.CenarioID).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete cenario players"})
	}

	if err := c.Con.Delete(&payload, "id = ?", payload.Id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete cenario campaign"})
	}

	campaign := &models.Campaign{}
	if err := c.Con.Preload("CenarioCampaigns").First(&campaign, "id = ?", payload.CampaignID).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Campaign not found"})
	}

	points := 0
	for _, c := range campaign.CenarioCampaigns {
		points += c.TotalFinalPoints
	}

	campaign.Points = points
	if err := c.Con.Save(&campaign).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update campaign points"})
	}

	return ctx.Status(fiber.StatusOK).JSON(campaign)
}
