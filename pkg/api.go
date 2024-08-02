package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GameDataCard struct {
	Id                   uint16    `json:"id"`
	Name                 string    `json:"name"`
	Image                string    `json:"image"`
	TypeId               uint16    `json:"objective_type_id"`
	RequredCopies        uint8     `json:"count"`
	VictoryPoints        uint8     `json:"vp"`
	AffiliationId        *uint16   `json:"affiliation_id"`
	TraitId              *uint16   `json:"trait_id"`
	RankIds              []*uint16 `json:"rank_ids"`
	RequiredCharacterIds []*uint16 `json:"required_character_ids"`
}

type GameData struct {
	Cards []GameDataCard `json:"cards"`
}

func FetchGameData() (*GameData, error) {
	resp, err := http.Get("https://app.knightmodels.com/gamedata")
	if err != nil {
		return nil, fmt.Errorf("FetchGameData: HTTP request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("FetchGameData: Failed to read response body: %w", err)
	}

	var gameData GameData
	err = json.Unmarshal(body, &gameData)
	if err != nil {
		return nil, fmt.Errorf("FetchGameData: Failed to unmarshal JSON response: %w", err)
	}

	return &gameData, nil
}
