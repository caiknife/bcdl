package bcdl

import (
	"fmt"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/ncmdl/v2"
	"github.com/gocolly/colly/v2"
	"github.com/pkg/errors"
)

type Song struct {
	Title         string                     `json:"title"`
	Album         string                     `json:"album"`
	Artist        string                     `json:"artist"`
	Cover         string                     `json:"cover"`
	URL           string                     `json:"url"`
	DownloadItems types.Slice[*DownloadItem] `json:"download_items"`

	c *colly.Collector
}

func (s *Song) String() string {
	toString, _ := fjson.MarshalToString(s)
	return toString
}

func NewSong(link string) *Song {
	song := &Song{URL: link, DownloadItems: types.Slice[*DownloadItem]{}}
	song.initCollector()
	return song
}

func (s *Song) Fetch() error {
	if s.c == nil {
		s.c = colly.NewCollector()
	}
	err := s.c.Visit(s.URL)
	if err != nil {
		err = errors.WithMessage(errors.WithStack(err), "visit url error")
		return err
	}
	return nil
}

func (s *Song) Info() string {
	return fmt.Sprintf("艺人[%s] 专辑[%s] 曲目[%s]", s.Artist, s.Album, s.Title)
}

func (s *Song) initCollector() {
	s.c = colly.NewCollector()

	s.c.OnRequest(func(r *colly.Request) {
		ncmdl.AppLogger.Infoln("正在解析单曲", r.URL.String())
	})

	s.c.OnHTML("div#name-section", func(e *colly.HTMLElement) {
		songTitle := e.ChildText("h2.trackTitle")
		songAlbum := e.ChildText("h3.albumTitle span:nth-child(1) a span")
		songArtist := e.ChildText("h3.albumTitle span:nth-child(2)")
		s.Title = songTitle
		s.Album = songAlbum
		s.Artist = songArtist
	})

	s.c.OnHTML("div#tralbumArt", func(e *colly.HTMLElement) {
		songCover := e.ChildAttr("a", "href")
		s.Cover = songCover
	})

	s.c.OnHTML("script[data-tralbum]", func(e *colly.HTMLElement) {
		songData := e.Attr("data-tralbum")
		songJSON := SongJSON{}
		_ = fjson.UnmarshalFromString(songData, &songJSON)
		for _, t := range songJSON.TrackInfo {
			s.DownloadItems = append(s.DownloadItems, &DownloadItem{
				Title:    s.Title,
				Album:    s.Album,
				Artist:   s.Artist,
				Cover:    s.Cover,
				TrackNum: t.TrackNum,
				Link:     t.File.MP3,
			})
		}
	})
}

type SongJSON struct {
	TrackInfo []struct {
		File struct {
			MP3 string `json:"mp3-128"`
		} `json:"file"`
		Title    string `json:"title"`
		TrackNum int    `json:"track_num"`
	} `json:"trackinfo"`
}
