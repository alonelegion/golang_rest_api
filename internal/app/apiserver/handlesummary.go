package apiserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Count struct {
	Domain        string `json:"first"`
	PositionCount string `json:"position_count"`
}

func (s *APIServer) HandleSummary(w http.ResponseWriter, r *http.Request) {
	var count []Count

	domain := r.URL.Query().Get("domain")

	db, err := sql.Open("postgres", s.config.Database.DatabaseURL)

	row, err := db.Query("SELECT domain, count(domain) FROM positions WHERE domain=$1 GROUP BY domain", domain)
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	for row.Next() {
		var c Count
		err := row.Scan(&c.Domain, &c.PositionCount)
		if err != nil {
			log.Print(err)
		}
		count = append(count, c)
	}

	json.NewEncoder(w).Encode(count)
}
