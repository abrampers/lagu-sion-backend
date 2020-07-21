package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/abrampers/lagu-sion-backend/models"
	"github.com/ghodss/yaml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func convertToSongs(songs *[]models.Song) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			return nil
		}
		file, _ := ioutil.ReadFile(path)
		json, err := yaml.YAMLToJSON(file)
		if err != nil {
			log.Printf("convertToSongs: YAMLToJSON error %v\n", err)
			return nil
		}

		song, err := models.NewSong(json)
		if err != nil {
			log.Printf("convertToSongs: NewSong error %v\n", err)
			return nil
		}
		*songs = append(*songs, *song)
		return nil
	}
}

func main() {
	// TODO: On prod, sslmode must be enabled
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=ls_client dbname=lagu_sion password=kepadaallahbripuji sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	db.DropTable(&models.Book{}, &models.Song{}, &models.Verse{})
	db.AutoMigrate(&models.Book{}, &models.Song{}, &models.Verse{})

	roots := []string{
		"books/lagu-sion",
		"books/lagu-sion-edisi-lengkap",
	}

	db.Create(&models.Book{Id: 1, ShortName: "LS", LongName: "Lagu Sion"})
	db.Create(&models.Book{Id: 2, ShortName: "LSEL", LongName: "Lagu Sion Edisi Lengkap"})

	var songs []models.Song

	for _, root := range roots {
		err := filepath.Walk(root, convertToSongs(&songs))
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, song := range songs {
		db.Create(&song)
	}
}
