package main

import (
	"botuy1/updates"
	"time"
)

func main() {
	botid := "5757526655:AAEDpfyVrZvBZZsdULg8lktpCmwVoMxSaV0"
	for {
		t := time.NewTicker(time.Second * 1)
		<-t.C
		var getter updates.GetUpdates
		getter.Offset = 3
		getter.Get(botid, "getUpdates")

	}
	//newsapi.FetchNews("fr")

}
