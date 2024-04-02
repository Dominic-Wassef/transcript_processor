package processor

import (
	"backend/helpers"
	"backend/models"
	"errors"
)

// ProcessUtterances loads, cleans, and merges utterances.
func ProcessUtterances(utterances []models.Utterance) ([]models.Utterance, error) {
	if len(utterances) == 0 {
		return nil, errors.New("no utterances to process")
	}

	// Clean each utterance's text.
	for i := range utterances {
		utterances[i].Text = helpers.CleanText(utterances[i].Text)
		utterances[i].Timestamp = convertMsToRFC3339(utterances[i].TimestampMs)
	}

	// Merge utterances
	mergedUtterances, err := MergeUtterances(utterances)
	if err != nil {
		return nil, err
	}

	return mergedUtterances, nil
}
