package bcdl

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type Album struct {
	Title  string             `json:"title"`
	Artist string             `json:"artist"`
	Cover  string             `json:"album"`
	URL    string             `json:"link"`
	Songs  types.Slice[*Song] `json:"songs"`
}

func (a *Album) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}
