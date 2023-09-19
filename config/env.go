package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// for more information, see --> https://github.com/joho/godotenv/issues/43#issuecomment-503183127
const projectDirName = "barcode-generator"

func LoadEnvFile() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("success load env file")
}

func GetDbName(dbName string) string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	return string(rootPath) + dbName
}
