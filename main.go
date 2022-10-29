package main

import (
	"botuy1/updates"
	"time"
)

func main() {
	botid := "5074588513:AAEOHB2-lHIrt9OI_Buw_3VZKXBJ3i58N5s"
	for {
		t := time.NewTicker(time.Second * 1)
		<-t.C
		var getter updates.GetUpdates
		getter.Offset = 3
		getter.Get(botid, "getUpdates")

	}
}
