package getngrokurl

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Tunnel struct {
	PublicURL string `json:"public_url"`
}

type TunnelsResponse struct {
	Tunnels []Tunnel `json:"tunnels"`
}

func GetNgrokURL() (string, error) {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get("http://127.0.0.1:4040/api/tunnels")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result TunnelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if len(result.Tunnels) == 0 {
		return "", errors.New("no tunnels found")
	}
	return result.Tunnels[0].PublicURL, nil
}
