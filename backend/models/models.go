package models

import "time"

// General represents a Wehrmacht general
type General struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Branch         string    `json:"branch"` // Heer, Kriegsmarine, Luftwaffe
	Rank           string    `json:"rank"`
	BirthDate      string    `json:"birth_date"`
	DeathDate      string    `json:"death_date"`
	Biography      string    `json:"biography"`
	PhotoURL       string    `json:"photo_url"`
	NotableBattles string    `json:"notable_battles"`
	CreatedAt      time.Time `json:"created_at"`
}

// Term represents a military term/definition
type Term struct {
	ID         int       `json:"id"`
	Term       string    `json:"term"`
	Definition string    `json:"definition"`
	Category   string    `json:"category"`
	CreatedAt  time.Time `json:"created_at"`
}

// Battle represents a military battle/operation
type Battle struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Date         string    `json:"date"`
	Location     string    `json:"location"`
	Description  string    `json:"description"`
	Participants string    `json:"participants"`
	Outcome      string    `json:"outcome"`
	CreatedAt    time.Time `json:"created_at"`
}

// Quote represents a quote from a general
type Quote struct {
	ID        int       `json:"id"`
	GeneralID int       `json:"general_id"`
	QuoteText string    `json:"quote_text"`
	Context   string    `json:"context"`
	Date      string    `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

// QuoteWithGeneral combines quote with general information
type QuoteWithGeneral struct {
	Quote
	GeneralName string `json:"general_name"`
	Rank        string `json:"rank"`
}
