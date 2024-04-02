package loader

import (
	"backend/models"
	"encoding/json"
	"os"
	"path/filepath"
)

// LoadUtterances reads and parses utterance files from the given directory.
func LoadUtterances(dir string) ([]models.Utterance, error) {
	var utterances []models.Utterance

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var utterance models.Utterance
			if err := json.Unmarshal(data, &utterance); err != nil {
				return err
			}

			utterances = append(utterances, utterance)
		}
		return nil
	})

	return utterances, err
}
