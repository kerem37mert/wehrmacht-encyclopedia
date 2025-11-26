package handlers

import (
	"database/sql"
	"net/http"
	"time"
	"wehrmacht-encyclopedia/database"
	"wehrmacht-encyclopedia/models"

	"github.com/labstack/echo/v4"
)

// GetGenerals returns all generals, optionally filtered by branch
func GetGenerals(c echo.Context) error {
	branch := c.QueryParam("branch")
	
	var query string
	var rows *sql.Rows
	var err error
	
	if branch != "" {
		query = "SELECT id, name, branch, rank, birth_date, death_date, biography, photo_url, notable_battles, created_at FROM generals WHERE branch = ? ORDER BY name"
		rows, err = database.DB.Query(query, branch)
	} else {
		query = "SELECT id, name, branch, rank, birth_date, death_date, biography, photo_url, notable_battles, created_at FROM generals ORDER BY name"
		rows, err = database.DB.Query(query)
	}
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()
	
	generals := []models.General{}
	for rows.Next() {
		var g models.General
		err := rows.Scan(&g.ID, &g.Name, &g.Branch, &g.Rank, &g.BirthDate, &g.DeathDate, &g.Biography, &g.PhotoURL, &g.NotableBattles, &g.CreatedAt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		generals = append(generals, g)
	}
	
	return c.JSON(http.StatusOK, generals)
}

// GetGeneral returns a single general by ID
func GetGeneral(c echo.Context) error {
	id := c.Param("id")
	
	var g models.General
	query := "SELECT id, name, branch, rank, birth_date, death_date, biography, photo_url, notable_battles, created_at FROM generals WHERE id = ?"
	err := database.DB.QueryRow(query, id).Scan(&g.ID, &g.Name, &g.Branch, &g.Rank, &g.BirthDate, &g.DeathDate, &g.Biography, &g.PhotoURL, &g.NotableBattles, &g.CreatedAt)
	
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "General not found"})
	}
	
	return c.JSON(http.StatusOK, g)
}

// GetTerms returns all terms, optionally filtered by search query
func GetTerms(c echo.Context) error {
	search := c.QueryParam("search")
	
	var query string
	var rows *sql.Rows
	var err error
	
	if search != "" {
		query = "SELECT id, term, definition, category, created_at FROM terms WHERE term LIKE ? OR definition LIKE ? ORDER BY term"
		searchPattern := "%" + search + "%"
		rows, err = database.DB.Query(query, searchPattern, searchPattern)
	} else {
		query = "SELECT id, term, definition, category, created_at FROM terms ORDER BY term"
		rows, err = database.DB.Query(query)
	}
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()
	
	terms := []models.Term{}
	for rows.Next() {
		var t models.Term
		err := rows.Scan(&t.ID, &t.Term, &t.Definition, &t.Category, &t.CreatedAt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		terms = append(terms, t)
	}
	
	return c.JSON(http.StatusOK, terms)
}

// GetBattles returns all battles
func GetBattles(c echo.Context) error {
	query := "SELECT id, name, date, location, description, participants, outcome, created_at FROM battles ORDER BY date"
	rows, err := database.DB.Query(query)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()
	
	battles := []models.Battle{}
	for rows.Next() {
		var b models.Battle
		err := rows.Scan(&b.ID, &b.Name, &b.Date, &b.Location, &b.Description, &b.Participants, &b.Outcome, &b.CreatedAt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		battles = append(battles, b)
	}
	
	return c.JSON(http.StatusOK, battles)
}

// GetBattle returns a single battle by ID
func GetBattle(c echo.Context) error {
	id := c.Param("id")
	
	var b models.Battle
	query := "SELECT id, name, date, location, description, participants, outcome, created_at FROM battles WHERE id = ?"
	err := database.DB.QueryRow(query, id).Scan(&b.ID, &b.Name, &b.Date, &b.Location, &b.Description, &b.Participants, &b.Outcome, &b.CreatedAt)
	
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Battle not found"})
	}
	
	return c.JSON(http.StatusOK, b)
}

// GetDailyQuote returns a quote that rotates daily
func GetDailyQuote(c echo.Context) error {
	// Get total count of quotes
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM quotes").Scan(&count)
	if err != nil || count == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No quotes available"})
	}
	
	// Calculate which quote to show based on current day
	now := time.Now()
	dayOfYear := now.YearDay()
	quoteIndex := (dayOfYear % count) + 1
	
	query := `
		SELECT q.id, q.general_id, q.quote_text, q.context, q.date, q.created_at, g.name, g.rank
		FROM quotes q
		JOIN generals g ON q.general_id = g.id
		WHERE q.id = ?
	`
	
	var qwg models.QuoteWithGeneral
	err = database.DB.QueryRow(query, quoteIndex).Scan(
		&qwg.ID, &qwg.GeneralID, &qwg.QuoteText, &qwg.Context, &qwg.Date, &qwg.CreatedAt, &qwg.GeneralName, &qwg.Rank,
	)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusOK, qwg)
}

// GetQuotes returns all quotes
func GetQuotes(c echo.Context) error {
	generalID := c.QueryParam("general_id")
	
	var query string
	var rows *sql.Rows
	var err error
	
	if generalID != "" {
		query = `
			SELECT q.id, q.general_id, q.quote_text, q.context, q.date, q.created_at, g.name, g.rank
			FROM quotes q
			JOIN generals g ON q.general_id = g.id
			WHERE q.general_id = ?
			ORDER BY q.date
		`
		rows, err = database.DB.Query(query, generalID)
	} else {
		query = `
			SELECT q.id, q.general_id, q.quote_text, q.context, q.date, q.created_at, g.name, g.rank
			FROM quotes q
			JOIN generals g ON q.general_id = g.id
			ORDER BY q.date
		`
		rows, err = database.DB.Query(query)
	}
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()
	
	quotes := []models.QuoteWithGeneral{}
	for rows.Next() {
		var qwg models.QuoteWithGeneral
		err := rows.Scan(&qwg.ID, &qwg.GeneralID, &qwg.QuoteText, &qwg.Context, &qwg.Date, &qwg.CreatedAt, &qwg.GeneralName, &qwg.Rank)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		quotes = append(quotes, qwg)
	}
	
	return c.JSON(http.StatusOK, quotes)
}

// SearchAll performs a global search across generals, terms, and battles
func SearchAll(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Search query required"})
	}
	
	searchPattern := "%" + query + "%"
	results := make(map[string]interface{})
	
	// Search generals
	generalRows, _ := database.DB.Query(
		"SELECT id, name, branch, rank FROM generals WHERE name LIKE ? OR biography LIKE ? LIMIT 5",
		searchPattern, searchPattern,
	)
	defer generalRows.Close()
	
	generals := []map[string]interface{}{}
	for generalRows.Next() {
		var id int
		var name, branch, rank string
		generalRows.Scan(&id, &name, &branch, &rank)
		generals = append(generals, map[string]interface{}{
			"id": id, "name": name, "branch": branch, "rank": rank,
		})
	}
	results["generals"] = generals
	
	// Search terms
	termRows, _ := database.DB.Query(
		"SELECT id, term, definition FROM terms WHERE term LIKE ? OR definition LIKE ? LIMIT 5",
		searchPattern, searchPattern,
	)
	defer termRows.Close()
	
	terms := []map[string]interface{}{}
	for termRows.Next() {
		var id int
		var term, definition string
		termRows.Scan(&id, &term, &definition)
		terms = append(terms, map[string]interface{}{
			"id": id, "term": term, "definition": definition,
		})
	}
	results["terms"] = terms
	
	// Search battles
	battleRows, _ := database.DB.Query(
		"SELECT id, name, date, location FROM battles WHERE name LIKE ? OR description LIKE ? LIMIT 5",
		searchPattern, searchPattern,
	)
	defer battleRows.Close()
	
	battles := []map[string]interface{}{}
	for battleRows.Next() {
		var id int
		var name, date, location string
		battleRows.Scan(&id, &name, &date, &location)
		battles = append(battles, map[string]interface{}{
			"id": id, "name": name, "date": date, "location": location,
		})
	}
	results["battles"] = battles
	
	return c.JSON(http.StatusOK, results)
}
