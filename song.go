package bcdl

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

type Song struct {
	Title  string `json:"title"`
	Album  string `json:"album"`
	Artist string `json:"artist"`
	Cover  string `json:"cover"`
	URL    string `json:"url"`
	Link   string `json:"link"`
}

func (s *Song) String() string {
	toString, _ := fjson.MarshalToString(s)
	return toString
}
