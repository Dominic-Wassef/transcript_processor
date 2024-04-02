package handlers

import (
	"backend/loader"
	"backend/models"
	"backend/processor"
	"encoding/json"
	"fmt"
	"net/http"
)

// TranscriptHandler handles the HTTP request to process and return a transcript of processed utterances.
func TranscriptHandler(w http.ResponseWriter, r *http.Request) {
	dir := "/Users/dominic/Desktop/golang/transcriptprocessor/utterances"

	// Load utterances from files
	utterances, err := loader.LoadUtterances(dir)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load utterances: %v", err), http.StatusInternalServerError)
		return
	}

	// Process utterances to correct speaker text
	processedUtterances, err := processor.ProcessUtterances(utterances)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to process utterances: %v", err), http.StatusInternalServerError)
		return
	}

	// Send as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(models.Transcript{Utterances: processedUtterances})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode transcript: %v", err), http.StatusInternalServerError)
	}
}
