package bcdl

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

type DownloadItem struct {
	Title  string `json:"title"`
	Album  string `json:"album"`
	Artist string `json:"artist"`
	Cover  string `json:"cover"`
	Link   string `json:"link"`
}

func (d *DownloadItem) String() string {
	toString, _ := fjson.MarshalToString(d)
	return toString
}
