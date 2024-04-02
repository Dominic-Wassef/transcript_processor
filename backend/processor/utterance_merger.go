package processor

import (
	"backend/models"
	"errors"
	"strings"
	"time"
	"unicode"
)

// MergeUtterances combines consecutive utterances by the same speaker when appropriate.
func MergeUtterances(utterances []models.Utterance) ([]models.Utterance, error) {
	if len(utterances) == 0 {
		return nil, errors.New("no utterances to process")
	}

	var merged []models.Utterance
	current := utterances[0] // Start with the first utterance

	for i := 1; i < len(utterances); i++ {
		if shouldMergeWithPreviousUtterance(current, utterances[i]) {
			// Update the current utterance's end time and append the text.
			current.EndTime = utterances[i].EndTime
			current.Text += " " + utterances[i].Text
		} else {
			// If not merging, add the current utterance to the results and move on to the next one.
			merged = append(merged, current)
			current = utterances[i]
		}
	}
	// Append the merge utterance.
	merged = append(merged, current)

	return merged, nil
}

// Determines whether two utterances should be merged
func shouldMergeWithPreviousUtterance(prev, current models.Utterance) bool {
	const maxTimeDiffSeconds = 2

	prevEndTime, errPrevEndTime := time.Parse(time.RFC3339, prev.EndTime)
	currentStartTime, errCurrentStartTime := time.Parse(time.RFC3339, current.StartTime)
	if errPrevEndTime != nil || errCurrentStartTime != nil {
		return false
	}

	timeDiff := currentStartTime.Sub(prevEndTime).Seconds()

	prevTime, errPrev := time.Parse(time.RFC3339, prev.Timestamp)
	currentTime, errCurrent := time.Parse(time.RFC3339, current.Timestamp)
	if errPrev != nil || errCurrent != nil {
		return false
	}

	diff := currentTime.Sub(prevTime).Seconds()
	if startsWithLowercase(current.Text) || len(strings.Fields(current.Text)) <= 3 || diff < maxTimeDiffSeconds {
		return true
	}

	lastChar := prev.Text[len(prev.Text)-1]
	if lastChar == ',' || lastChar == '-' {
		return true
	}

	if prev.Speaker == current.Speaker && timeDiff <= maxTimeDiffSeconds {
		return true
	}

	return false
}

// Checks if the provided text string starts with a lowercase letter
func startsWithLowercase(text string) bool {
	if len(text) == 0 {
		return false
	}
	firstCharacter := rune(text[0])
	return unicode.IsLower(firstCharacter)
}
