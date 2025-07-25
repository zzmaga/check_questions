package main

import (
	"database/sql"
	"os"
	"strings"
)

func AddQuestionsFromFile(filePath string, db *sql.DB) (addedCount int, err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		q := strings.TrimSpace(line)
		if q == "" {
			continue
		}

		_, err := db.Exec("INSERT OR IGNORE INTO questions (question_text) VALUES (?)", q)
		if err == nil {
			addedCount++
		}
	}
	return
}
