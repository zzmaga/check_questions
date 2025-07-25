package main

import (
	"database/sql"
	"os"
	"strings"
)

func CheckQuestions(filePath string, db *sql.DB) (existing []string, newOnes []string, err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	questions := []string{}
	for _, line := range lines {
		q := strings.TrimSpace(line)
		if q != "" {
			questions = append(questions, q)
		}
	}

	// Загружаем все существующие вопросы
	rows, err := db.Query("SELECT question_text FROM questions")
	if err != nil {
		return
	}
	defer rows.Close()

	existingSet := make(map[string]bool)
	for rows.Next() {
		var q string
		rows.Scan(&q)
		existingSet[q] = true
	}

	// Сравнение
	for _, q := range questions {
		if existingSet[q] {
			existing = append(existing, q)
		} else {
			newOnes = append(newOnes, q)
		}
	}

	return
}
