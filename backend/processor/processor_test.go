package processor

import (
	"backend/models"
	"reflect"
	"testing"
	"time"
)

func TestProcessUtterances(t *testing.T) {
	// Setting a fixed point in time for reproducible tests
	now := time.Date(2024, 4, 2, 15, 0, 0, 0, time.UTC)
	startTimeMs := now.UnixNano() / int64(time.Millisecond)

	// Mock utterances
	mockUtterances := []models.Utterance{
		{
			GroupId:       1,
			EndTime:       now.Add(20 * time.Second).Format(time.RFC3339),
			MeetingId:     "meeting123",
			OffsetSeconds: "0",
			Speaker:       "Alex",
			StartTime:     now.Format(time.RFC3339),
			Text:          "Hello!!",
			TimestampMs:   startTimeMs,
		},
		{
			GroupId:       1,
			EndTime:       now.Add(40 * time.Second).Format(time.RFC3339),
			MeetingId:     "meeting123",
			OffsetSeconds: "20",
			Speaker:       "Alex",
			StartTime:     now.Add(20 * time.Second).Format(time.RFC3339),
			Text:          "It's me again!!",
			TimestampMs:   startTimeMs,
		},
	}

	// Expected output
	expected := []models.Utterance{
		{
			GroupId:       1,
			EndTime:       now.Add(40 * time.Second).Format(time.RFC3339),
			MeetingId:     "meeting123",
			OffsetSeconds: "0",
			Speaker:       "Alex",
			StartTime:     now.Format(time.RFC3339),
			Text:          "Hello! It's me again!",
			TimestampMs:   startTimeMs,
			Timestamp:     convertMsToRFC3339(startTimeMs),
		},
	}

	processed, err := ProcessUtterances(mockUtterances)
	if err != nil {
		t.Fatalf("ProcessUtterances returned an error: %v", err)
	}

	if len(processed) != len(expected) {
		t.Fatalf("Expected %d processed utterance, got %d", len(expected), len(processed))
	}

	for i, exp := range expected {
		if !reflect.DeepEqual(exp, processed[i]) {
			t.Errorf("At index %d, expected processed utterance to be %+v, got %+v", i, exp, processed[i])
		}
	}
}
