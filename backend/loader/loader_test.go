package loader

import (
	"backend/models"
	"os"
	"path/filepath"
	"testing"
)

// mockFileSetup sets up a temporary directory and creates mock JSON files for testing
func mockFileSetup() (string, error) {
	dir, err := os.MkdirTemp("", "test")
	if err != nil {
		return "", err
	}
	files := []struct {
		name    string
		content []byte
	}{
		{"test1.json", []byte(`{"group_id":1,"endTime":"2024-01-22T18:00:56.584447635Z","meetingId":"testMeetingId","offsetSeconds":"0433.02","speaker":"alex@aircover.ai","startTime":"2024-01-22T18:00:56.584447635Z","text":"Hello, world","timestampMs":1705946456584}`)},
	}
	for _, file := range files {
		if err := os.WriteFile(filepath.Join(dir, file.name), file.content, 0644); err != nil {
			return "", err
		}
	}
	return dir, nil
}

func TestLoadUtterances(t *testing.T) {
	dir, err := mockFileSetup()
	if err != nil {
		t.Fatalf("Failed to set up mock file environment: %v", err)
	}
	defer os.RemoveAll(dir) // clean up

	utterances, err := LoadUtterances(dir)
	if err != nil {
		t.Fatalf("Failed to load utterances: %v", err)
	}
	if len(utterances) != 1 {
		t.Errorf("Expected 1 utterance, got %d", len(utterances))
	}

	expected := models.Utterance{
		GroupId:       1,
		EndTime:       "2024-01-22T18:00:56.584447635Z",
		MeetingId:     "testMeetingId",
		OffsetSeconds: "0433.02",
		Speaker:       "alex@aircover.ai",
		StartTime:     "2024-01-22T18:00:56.584447635Z",
		Text:          "Hello, world",
		TimestampMs:   1705946456584,
	}
	if utterances[0] != expected {
		t.Errorf("Unexpected utterance data: got %+v, want %+v", utterances[0], expected)
	}
}
