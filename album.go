package bcdl

import (
	"fmt"
	"strings"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/ncmdl/v2"
	"github.com/gocolly/colly/v2"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type Album struct {
	Title         string                     `json:"title"`
	Artist        string                     `json:"artist"`
	Cover         string                     `json:"album"`
	URL           string                     `json:"link"`
	DownloadItems types.Slice[*DownloadItem] `json:"download_items"`

	c *colly.Collector
}

func (a *Album) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}

func NewAlbum(link string) *Album {
	album := &Album{URL: link, DownloadItems: types.Slice[*DownloadItem]{}}
	album.initCollector()
	return album
}

func (a *Album) Fetch() error {
	if a.c == nil {
		a.initCollector()
	}
	err := a.c.Visit(a.URL)
	if err != nil {
		err = errors.WithMessage(err, "visit url error")
		return err
	}
	return nil
}

func (a *Album) Info() string {
	songList := lo.Map[*DownloadItem, string](a.DownloadItems, func(item *DownloadItem, index int) string {
		return fmt.Sprintf("%02d-%s", item.TrackNum, item.Title)
	})
	return fmt.Sprintf("艺人[%s] 专辑[%s] 曲目[%s]", a.Artist, a.Title, strings.Join(songList, ", "))
}

func (a *Album) initCollector() {
	a.c = colly.NewCollector()

	a.c.OnRequest(func(r *colly.Request) {
		ncmdl.AppLogger.Infoln("正在解析专辑", r.URL.String())
	})
	a.c.OnHTML("div#name-section", func(e *colly.HTMLElement) {
		albumTitle := e.ChildText("h2.trackTitle")
		albumArtist := e.ChildText("h3 > span > a")
		a.Title = albumTitle
		a.Artist = albumArtist
	})

	a.c.OnHTML("div#tralbumArt", func(e *colly.HTMLElement) {
		albumCover := e.ChildAttr("a", "href")
		a.Cover = albumCover
	})

	a.c.OnHTML("script[data-tralbum]", func(e *colly.HTMLElement) {
		albumData := e.Attr("data-tralbum")
		albumJSON := AlbumJSON{}
		_ = fjson.UnmarshalFromString(albumData, &albumJSON)
		for _, s := range albumJSON.TrackInfo {
			a.DownloadItems = append(a.DownloadItems, &DownloadItem{
				Title:    s.Title,
				Album:    a.Title,
				Artist:   a.Artist,
				Cover:    a.Cover,
				TrackNum: s.TrackNum,
				Link:     s.File.MP3,
			})
		}
	})
}

type AlbumJSON struct {
	TrackInfo []struct {
		File struct {
			MP3 string `json:"mp3-128"`
		} `json:"file"`
		Title    string `json:"title"`
		TrackNum int    `json:"track_num"`
	} `json:"trackinfo"`
}
