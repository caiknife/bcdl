package bcdl

import (
	"fmt"
	"path/filepath"

	"github.com/caiknife/mp3lister/lib/fjson"
)

type DownloadItem struct {
	Title    string `json:"title"`
	Album    string `json:"album"`
	Artist   string `json:"artist"`
	Cover    string `json:"cover"`
	TrackNum int    `json:"track_num"`
	Link     string `json:"link"`
}

func (d *DownloadItem) CoverURL() string {
	return d.Cover
}

func (d *DownloadItem) AllArtistsTag() string {
	return d.Artist
}

func (d *DownloadItem) String() string {
	toString, _ := fjson.MarshalToString(d)
	return toString
}

func (d *DownloadItem) SavePath() string {
	return filepath.Join(d.Artist, d.Album)
}

func (d *DownloadItem) FileName() string {
	return fmt.Sprintf("%02d - %s - %s.mp3", d.TrackNum, d.Artist, d.Title)
}

func (d *DownloadItem) SaveFileName() string {
	return filepath.Join(d.SavePath(), d.FileName())
}
