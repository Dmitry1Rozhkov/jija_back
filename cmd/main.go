package main

import (
	"encoding/json"
	"fmt"
	"jija_back/internal/config"
	"jija_back/internal/domain"
	"jija_back/internal/handler"
	"jija_back/internal/server"
	"log"
	"os"
)

func main() {
	// init data
	var offices []domain.Office

	var atmsData struct {
		Atms []domain.Atm `json:"atms"`
	}

	atmsFile, err := os.ReadFile("atms_rich.json")
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal([]byte(atmsFile), &atmsData)
	if err != nil {
		log.Panic(err)
	}

	officesFile, err := os.ReadFile("offices_rich.json")
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal([]byte(officesFile), &offices)
	if err != nil {
		log.Panic(err)
	}
	// init config
	appConfig, err := config.Init("./")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(len(atmsData.Atms), len(offices))
	fmt.Println(len(atmsData.Atms), offices)
	// init handler
	handlers := handler.New(appConfig, atmsData.Atms, offices)

	// init server
	s := server.NewServer(appConfig, handlers)

	if err = s.GinRouter.Run(fmt.Sprintf("%s:%s", appConfig.HTTP.Host, appConfig.HTTP.Port)); err != nil {
		log.Fatal(err)
	}

}
