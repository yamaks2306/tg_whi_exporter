package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func GetWebhookInfo(token string) (*TgWebhookInfo, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getWebhookInfo", token)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("No response from request")
		return nil, err
	}

	var result TgWebhookInfo
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
		return nil, err
	}

	return &result, nil
}
