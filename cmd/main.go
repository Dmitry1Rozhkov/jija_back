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

	_ = json.Unmarshal([]byte(atmsFile), &atmsData)

	officesFile, err := os.ReadFile("offices_rich.json")
	if err != nil {
		log.Panic(err)
	}

	_ = json.Unmarshal([]byte(officesFile), &offices)

	// init config
	appConfig, err := config.Init("./")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(len(atmsData.Atms), len(offices))
	// init handler
	handlers := handler.New(appConfig, atmsData.Atms, offices)

	// init server
	s := server.NewServer(appConfig, handlers)

	if err = s.GinRouter.Run(fmt.Sprintf("%s:%s", appConfig.HTTP.Host, appConfig.HTTP.Port)); err != nil {
		log.Fatal(err)
	}

}
