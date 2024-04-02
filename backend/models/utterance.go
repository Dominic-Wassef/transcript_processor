package models

// Utterance represents a segment of speech in a transcript.
type Utterance struct {
	GroupId       int    `json:"group_id"`
	EndTime       string `json:"endTime"`
	MeetingId     string `json:"meetingId"`
	OffsetSeconds string `json:"offsetSeconds"`
	Speaker       string `json:"speaker"`
	StartTime     string `json:"startTime"`
	Text          string `json:"text"`
	TimestampMs   int64  `json:"timestampMs"`
	Timestamp     string `json:"timestamp"`
}

// Transcript aggregates multiple Utterances.
type Transcript struct {
	Utterances []Utterance `json:"utterances"`
}
