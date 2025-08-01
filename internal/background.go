package internal

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)



func generateTodayPath() string {
	basePath := `C:/Users/Administrator/Desktop/winners`
	filenamePattern := "applications_{date}_past_1days/profiles"
	fileExtension := ".csv"

	date := time.Now().Format("2006-01-02")
	filename := strings.Replace(filenamePattern, "{date}", date, 1)

	fullPath := filepath.Join(basePath, filename) + fileExtension
	return fullPath
}

func Run() {
	todayPath := generateTodayPath()

	if _, err := os.Stat(todayPath); os.IsNotExist(err) {
		log.Printf("CSV file not found: %s. Skipping this run.", todayPath)
		return
	}

	Start(todayPath)
}
