package updates

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (getter *GetUpdates) Get(botid string, method string) *Update {
	url := fmt.Sprintf("https://api.telegram.org/bot%v/%v", botid, method)
	content, err := http.Get(url)
	reader := make([]byte, 10024)
	n, _ := content.Body.Read(reader)
	fmt.Printf("%s\n", reader[:n])
	var update *Update
	if err != nil {
		return nil
	}
	decoder := json.NewDecoder(content.Body)
	err = decoder.Decode(&update)
	if err != nil {
		fmt.Printf("error decoding")
		return nil
	}

	return update
}
