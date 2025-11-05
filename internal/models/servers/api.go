package servers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	scnorion_nats "github.com/scncore/nats"
)

func GetLatestServerReleaseFromAPI(tmpDir string) (*scnorion_nats.scnorionRelease, error) {
	latestServerReleasePath := filepath.Join(tmpDir, "latest.json")

	if _, err := os.Stat(latestServerReleasePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("latest server releases json file doesn't exist, reason: %v", err)
	}

	data, err := os.ReadFile(latestServerReleasePath)
	if err != nil {
		return nil, err
	}

	r := scnorion_nats.scnorionRelease{}
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
