package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Использование:")
		fmt.Println("  go run main.go add questions.txt      - Добавить вопросы")
		fmt.Println("  go run main.go check new_questions.txt - Проверить вопросы")
		return
	}

	command := os.Args[1]
	filePath := os.Args[2]

	db, err := InitDB("questions.db")
	if err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}
	defer db.Close()

	switch command {
	case "add":
		count, err := AddQuestionsFromFile(filePath, db)
		if err != nil {
			log.Fatal("Ошибка при добавлении:", err)
		}
		fmt.Printf("✅ Добавлено %d вопросов\n", count)

	case "check":
		existing, newOnes, err := CheckQuestions(filePath, db)
		if err != nil {
			log.Fatal("Ошибка при проверке:", err)
		}

		fmt.Println("\n✅ Уже есть в базе:")
		for _, q := range existing {
			fmt.Println("- " + q)
		}

		fmt.Println("\n➕ Новые вопросы:")
		for _, q := range newOnes {
			fmt.Println("- " + q)
		}

		fmt.Printf("\nИтого: %d новых, %d уже есть\n", len(newOnes), len(existing))

	default:
		fmt.Println("Неизвестная команда:", command)
	}
}
