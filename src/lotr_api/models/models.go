package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	Id               string            `json:"id" gorm:"primaryKey"`
	Name             string            `json:"name" gorm:"unique;not null;varchar(100)"`
	Points           int               `json:"points" gorm:"not null"`
	Notes            string            `json:"notes" gorm:"varchar(255)"`
	CreatedAt        time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
	Players          []PlayerToPlay    `json:"players" gorm:"many2many:player_campaign;"`
	CenarioCampaigns []CenarioCampaign `json:"cenario_campaigns" gorm:"foreignKey:CampaignID;references:Id"`
	DefeatedHeros    []Card            `json:"defeated_heroes" gorm:"many2many:campaign_defeated_heroes;"`
}

type Player struct {
	Id        string    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"unique"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type PlayerToPlay struct {
	ID       string `json:"id" gorm:"primaryKey"`
	PlayerID string `json:"player_id" gorm:"not null"`
	Player   Player `json:"player" gorm:"foreignKey:PlayerID"`
	Heroes   []Card `json:"heroes" gorm:"many2many:player_campaign_hero;"`
	DeckID   string `json:"deck_id" gorm:"not null"`
	Deck     Deck   `json:"deck" gorm:"foreignKey:DeckID"`
}

type Collection struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null;unique;varchar(100)"`
	Description string `json:"description" gorm:"not null;varchar(255)"`
}

type Card struct {
	Id           string     `json:"id" gorm:"primaryKey"`
	Number       int        `json:"number" gorm:"not null"`
	CollectionID string     `json:"collection_id" gorm:"not null"`
	Collection   Collection `json:"collection" gorm:"foreignKey:CollectionID"`
	Name         string     `json:"name" gorm:"varchar(100);not null"`
	Description  string     `json:"description" gorm:"varchar(100);not null"`
	Type         string     `json:"type" gorm:"varchar(100);not null"`
	Cost         int        `json:"cost" gorm:"not null"`
}

type Deck struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null;varchar(100);unique"`
	Description string `json:"description" gorm:"not null;varchar(255)"`
	Cards       []Card `json:"cards" gorm:"many2many:deck_cards;"`
}

type Cenario struct {
	Id           string     `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name" gorm:"not null;unique;varchar(100)"`
	Description  string     `json:"description" gorm:"not null;varchar(255)"`
	Order        int        `json:"order" gorm:"not null"`
	CollectionID string     `json:"collection_id" gorm:"not null"`
	Collection   Collection `json:"collection" gorm:"foreignKey:CollectionID"`
}

type CenarioPlayer struct {
	Id                 string  `json:"id" gorm:"primaryKey"`
	CampaignID         string  `json:"campaign_id" gorm:"not null"`
	PlayerID           string  `json:"player_id" gorm:"not null"`
	Player             Player  `json:"player" gorm:"foreignKey:PlayerID"`
	CenarioID          string  `json:"cenario_id" gorm:"not null"`
	Cenario            Cenario `json:"cenario" gorm:"foreignKey:CenarioID"`
	FinalThreat        int     `json:"final_threat" gorm:"not null"`
	ThreatHeroDefeated int     `json:"threat_hero_defeated" gorm:"not null"`
	HeroesDamage       int     `json:"heroes_damage" gorm:"not null"`
	TotalPoints        int     `json:"total" gorm:"not null"`
}

type CenarioCampaign struct {
	Id                 string          `json:"id" gorm:"primaryKey"`
	CampaignID         string          `json:"campaign_id" gorm:"not null"`
	CenarioID          string          `json:"cenario_id" gorm:"not null"`
	Cenario            Cenario         `json:"cenario" gorm:"foreignKey:CenarioID"`
	CenarioPlayers     []CenarioPlayer `json:"cenario_players" gorm:"foreignKey:CampaignID,CenarioID;references:CampaignID,CenarioID"`
	TotalPointsPlayers int             `json:"total_points" gorm:"not null"`
	VictoryPoints      int             `json:"victory_points" gorm:"not null"`
	Rounds             int             `json:"rounds" gorm:"not null"`
	RoundPoints        int             `json:"round_points" gorm:"not null"`
	TotalFinalPoints   int             `json:"total_final_points" gorm:"not null"`
}

func NewPlayer(name string) *Player {
	id := uuid.New().String()
	return &Player{
		Id:   id,
		Name: name,
	}
}

func NewCampaign(name string, players []PlayerToPlay) *Campaign {
	id := uuid.New().String()
	return &Campaign{
		Id:      id,
		Name:    name,
		Players: players,
		Points:  0,
	}
}

func NewCollection(name, description string) *Collection {
	id := uuid.New().String()
	return &Collection{
		Id:          id,
		Name:        name,
		Description: description,
	}
}

func NewCard(number int, collectionID, name, description, typeString string, cost int) *Card {
	id := fmt.Sprintf("%d - %s", number, collectionID)
	return &Card{
		Id:           id,
		Number:       number,
		CollectionID: collectionID,
		Name:         name,
		Description:  description,
		Type:         typeString,
		Cost:         cost,
	}
}

func NewDeck(name, description string, cards []Card) *Deck {
	id := uuid.New().String()
	return &Deck{
		Id:          id,
		Name:        name,
		Description: description,
		Cards:       cards,
	}
}

func NewPlayerToPlay(player Player, heroes []Card, deck Deck) *PlayerToPlay {
	id := uuid.New().String()
	return &PlayerToPlay{
		ID:     id,
		Player: player,
		Heroes: heroes,
		Deck:   deck,
	}
}

func NewCenario(name, description, collectionID string) *Cenario {
	id := uuid.New().String()
	return &Cenario{
		Id:           id,
		Name:         name,
		Description:  description,
		CollectionID: collectionID,
	}
}

func NewCenarioPlayer(campaignID, playerID, cenarioID string, finalThreat, threatHeroDefeated, heroesDamage int) *CenarioPlayer {

	id := fmt.Sprintf("%s-%s-%s", campaignID, playerID, cenarioID)
	totalPoints := finalThreat + threatHeroDefeated + heroesDamage

	return &CenarioPlayer{
		Id:                 id,
		CampaignID:         campaignID,
		PlayerID:           playerID,
		CenarioID:          cenarioID,
		Cenario:            Cenario{Id: cenarioID},
		FinalThreat:        finalThreat,
		ThreatHeroDefeated: threatHeroDefeated,
		HeroesDamage:       heroesDamage,
		TotalPoints:        totalPoints,
	}
}

func NewCenarioCampaign(campaignID, cenarioID string, cenarioPlayers []CenarioPlayer, victoryPoints int, rounds int) *CenarioCampaign {

	totalPointsPlayers := 0
	for _, player := range cenarioPlayers {
		totalPointsPlayers += player.TotalPoints
	}

	totalFinalPoints := totalPointsPlayers - victoryPoints + (rounds * 10)

	id := uuid.New().String()
	return &CenarioCampaign{
		Id:                 id,
		CampaignID:         campaignID,
		CenarioID:          cenarioID,
		Cenario:            Cenario{Id: cenarioID},
		CenarioPlayers:     cenarioPlayers,
		VictoryPoints:      victoryPoints,
		Rounds:             rounds,
		RoundPoints:        rounds * 10,
		TotalPointsPlayers: totalPointsPlayers,
		TotalFinalPoints:   totalFinalPoints,
	}
}
