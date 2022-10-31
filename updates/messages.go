package updates

import (
	"fmt"
	"net/http"
	"net/url"
)

func SendMessage(botid string, chaid string, text string) {
	form_data := url.Values{
		"chat_id": {chaid},
		"text":    {text},
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", botid)
	_, err := http.PostForm(url, form_data)
	if err != nil {
		fmt.Printf("error send %s\n", chaid)
		return
	}

}
