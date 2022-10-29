package updates

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (getter *GetUpdates) Get(botid string, method string) *Update {
	url := fmt.Sprintf("https://api.telegram.org/bot%v/%v", botid, method)
	fmt.Println(url)
	content, err := http.Get(url)
	//read := make([]byte, 9000)
	//n, _ := content.Body.Read(read)
	//fmt.Printf("\n%s\n", read[:n])
	var updates *Update
	if err != nil {
		return nil
	}
	decoder := json.NewDecoder(content.Body)
	err = decoder.Decode(&updates)
	if err != nil {
		return nil
	}
	fmt.Printf("%v\n", len(updates.Result))
	fmt.Printf("%v\n", updates.Result[len(updates.Result)-1])
	return updates
}
