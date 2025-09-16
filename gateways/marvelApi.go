package gateways

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/OlaIsaac/tech-talk/config"
)

type APIResponse struct {
	Data struct {
		Results []Character `json:"results"`
	} `json:"data"`
}

type Character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumbnail   struct {
		Path      string `json:"path"`
		Extension string `json:"extension"`
	} `json:"thumbnail"`
}

func GetCharacter(name string) (Character, error) {
	cfg := config.Get()
	client := http.Client{Timeout: 10 * time.Second}

	ts := strconv.FormatInt(time.Now().Unix(), 10)
	hash := generateHash(ts, cfg.Marvel.PrivateKey, cfg.Marvel.PublicKey)

	// Encode the character name for URL safety
	encodedName := url.QueryEscape(name)

	apiURL := fmt.Sprintf("%s/v1/public/characters?name=%s&ts=%s&apikey=%s&hash=%s",
		cfg.Marvel.BaseURL, encodedName, ts, cfg.Marvel.PublicKey, hash)

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return Character{}, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return Character{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Character{}, fmt.Errorf("failed to read response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return Character{}, fmt.Errorf("API returned status %d: %s", res.StatusCode, string(body))
	}

	var apiResponse APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return Character{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(apiResponse.Data.Results) == 0 {
		return Character{}, fmt.Errorf("no character found with name: %s", name)
	}

	return apiResponse.Data.Results[0], nil
}

func generateHash(ts, privateKey, publicKey string) string {
	data := ts + privateKey + publicKey
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
