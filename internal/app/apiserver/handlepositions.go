package apiserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type positionsSummary struct {
	Keyword  string `json:"keyword"`
	Position int    `json:"position"`
	URL      string `json:"url"`
	Volume   int    `json:"volume"`
	Results  int    `json:"results"`
	Updated  string `json:"updated"`
}

type positions struct {
	Positions []positionsSummary
}

func (s *APIServer) HandlePositions(w http.ResponseWriter, r *http.Request) {
	posits := positions{}

	domain := r.URL.Query().Get("domain")

	db, err := sql.Open("postgres", s.config.Database.DatabaseURL)

	row, err := db.Query("SELECT keyword, position, url, volume, results, updated FROM positions WHERE domain = $1 ORDER BY volume", domain)

	defer row.Close()

	for row.Next() {
		pos := positionsSummary{}
		err = row.Scan(
			&pos.Keyword,
			&pos.Position,
			&pos.URL,
			&pos.Volume,
			&pos.Results,
			&pos.Updated,
		)
		if err != nil {
			fmt.Print(err)
		}
		posits.Positions = append(posits.Positions, pos)
	}

	out, err := json.Marshal(posits)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}
