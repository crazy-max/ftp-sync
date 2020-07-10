package model

import (
	"encoding/json"
	"time"

	"github.com/hako/durafmt"
)

// Journal holds ftpgrab entries and status
type Journal struct {
	ServerHost string  `json:"-"`
	Entries    []Entry `json:"entries,omitempty"`
	Count      struct {
		Success int `json:"success,omitempty"`
		Error   int `json:"error,omitempty"`
		Skip    int `json:"skip,omitempty"`
	} `json:"count,omitempty"`
	Status   string        `json:"status,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
}

// Entry represents a journal entry
type Entry struct {
	File       string `json:"file,omitempty"`
	StatusType string `json:"status_type,omitempty"`
	StatusText string `json:"status_text,omitempty"`
}

// EntryStatus represents entry status
type EntryStatus string

func (j Journal) MarshalJSON() ([]byte, error) {
	type Alias Journal
	return json.Marshal(&struct {
		Alias
		Duration string `json:"duration,omitempty"`
	}{
		Alias:    (Alias)(j),
		Duration: durafmt.ParseShort(j.Duration).String(),
	})
}
