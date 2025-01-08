package help

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Release struct {
	TagName string `json:"tag_name"`
}

func CheckForUpdates() (bool, error) {
	currentVersion := "v1.0.0"

	repo := "nightvalley/tammy"
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to check for updates: %s", resp.Status)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return false, err
	}

	return release.TagName != currentVersion, nil
}
