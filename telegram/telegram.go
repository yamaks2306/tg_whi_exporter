package telegram

import (
	"fmt"
	"log"
	"io"
	"encoding/json"
	"net/http"
)

type TgWebhookInfo struct {
	Status bool `json:"ok"`
	Result struct {
		URL                    string `json:"url"`
		Has_custom_certificate bool   `json:"has_custom_certificate"`
		Pending_update_count   int    `json:"pending_update_count"`
		Max_connections        int    `json:"max_connections"`
		IP_address             string `json:"ip_address"`
	} `json:"result"`
}

func GetWebhookInfo(token string) TgWebhookInfo {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getWebhookInfo", token)
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	log.Println(response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("No response from request")
	}

	var result TgWebhookInfo
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}