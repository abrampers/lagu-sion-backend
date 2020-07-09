package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// TODO: On prod, sslmode must be enabled
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=ls_client dbname=lagu_sion password=kepadaallahbripuji sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Book{}, &Song{}, &Verse{})

	// song := &Song{
	// 	Number: 1,
	// 	Title:  "Ke Pada Allah Bri Puji",
	// 	Verses: []Verse{
	// 		Verse{Contents: "Ke hahahah\nhahahahaah"},
	// 		Verse{Contents: "Ke hahahahh\nahahahaah"},
	// 		Verse{Contents: "Ke hahahahha\nhahahaah"},
	// 	},
	// 	Reff: Verse{Contents: "TestReff"},
	// 	Book: Book{ShortName: "LS", LongName: "Lagu Sion"},
	// }

	// var song Song
	//
	// file, _ := ioutil.ReadFile("books/lagu-sion/1.json")
	// if err := json.Unmarshal(file, &song); err != nil {
	// 	panic(err)
	// }
	//
	// db.Create(&song)

	file, _ := ioutil.ReadFile("books/lagu-sion/1.yaml")

	json, err := yaml.YAMLToJSON(file)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(json))
}
