package main

import (
	"bufio"
	"golang-server-android/config"
	"golang-server-android/db"
	"golang-server-android/models"
	"log"
	"os"
	"strings"
)

func main() {
	config.Init()
	db.Init()
	database := db.GetDb()

	words, err := scanFile("cmd/data.txt")
	if err != nil {
		return
	}
	database.CreateInBatches(&words, 100)
}

func scanFile(filePath string) ([]models.Word, error) {
	var result []models.Word
	tempWord := models.Word{Word: "@", WordType: "*", Definition: "| "}
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text() // GET the line string
		if strings.HasPrefix(line, "@") {
			if strings.HasPrefix(tempWord.Definition, "|") {
				tempWord.Definition = tempWord.Definition[1:]
			}
			result = append(result, tempWord)
			tempWord = models.Word{}
			tempWord.Word = line[1:]
			tempWord.Definition = ""
			continue
		}
		if strings.HasPrefix(line, "*") {
			tempWord.WordType = line[1:]
			continue
		}
		if strings.HasPrefix(line, "-") {
			tempWord.Definition = tempWord.Definition + "|" + line[1:]
			continue
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return nil, err
	}
	if strings.HasPrefix(tempWord.Definition, "|") {
		tempWord.Definition = tempWord.Definition[1:]
	}
	result = append(result, tempWord)
	result = result[1:]
	return result, nil
}
