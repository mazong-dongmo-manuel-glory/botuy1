package main

import (
	"botuy1/newsapi"
	"botuy1/updates"
	"fmt"
	"time"
)

func main() {
	botid := "5757526655:AAEDpfyVrZvBZZsdULg8lktpCmwVoMxSaV0"
	var chatid = []string{"-1622761550", "-1001735042795", "-1001886047426"}
	cpt := 0
	for {

		var getter updates.GetUpdates
		getter.Offset = 3
		resp := newsapi.FetchNews("fr")
		message_response := fmt.Sprintf("%v\n%v\n%v", resp.Results[cpt].Title, resp.Results[cpt].Link, resp.Results[cpt].Description)
		for _, id := range chatid {
			updates.SendMessage(botid, id, message_response)
		}
		t := time.NewTicker(time.Hour * 4)
		<-t.C
		cpt += 1

	}

}
