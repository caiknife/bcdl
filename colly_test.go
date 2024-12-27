package bcdl

import (
	"testing"

	"github.com/caiknife/ncmdl/v2"
	"github.com/gocolly/colly/v2"
)

const (
	testAlbumLink = "https://shirttailstompers.bandcamp.com/album/thats-my-kick"
	testSongLink  = "https://shirttailstompers.bandcamp.com/track/sweets"
)

var (
	albumColly = colly.NewCollector()
	songColly  = colly.NewCollector()
)

func init() {
	albumColly.OnRequest(func(r *colly.Request) {
		ncmdl.AppLogger.Infoln(r.URL.String())
	})

	albumColly.OnHTML("div#name-section", func(e *colly.HTMLElement) {
		albumTitle := e.ChildText("h2.trackTitle")
		albumArtist := e.ChildText("h3 span a")
		ncmdl.AppLogger.Infoln("Album Title", albumTitle)
		ncmdl.AppLogger.Infoln("Album Artist", albumArtist)
	})

	albumColly.OnHTML("div#tralbumArt", func(e *colly.HTMLElement) {
		albumCover := e.ChildAttr("a", "href")
		ncmdl.AppLogger.Infoln("Album Cover", albumCover)
	})

	albumColly.OnHTML("table#track_table", func(e *colly.HTMLElement) {
		e.ForEach("td.title-col", func(i int, t *colly.HTMLElement) {
			songTitle := t.ChildText("a span")
			songLink := t.ChildAttr("a", "href")
			ncmdl.AppLogger.Infoln("Song", i+1, songTitle, e.Request.AbsoluteURL(songLink))
		})
	})

	// albumColly.OnHTML("script[data-tralbum]", func(e *colly.HTMLElement) {
	// 	ncmdl.AppLogger.Infoln("Album Data", e.Attr("data-tralbum"))
	// })
}

func init() {
	songColly.OnRequest(func(r *colly.Request) {
		ncmdl.AppLogger.Infoln(r.URL.String())
	})
	
	songColly.OnHTML("div#name-section", func(e *colly.HTMLElement) {
		songTitle := e.ChildText("h2.trackTitle")
		ncmdl.AppLogger.Infoln("Song Title", songTitle)
		songAlbum := e.ChildText("h3.albumTitle span:nth-child(1) a span")
		ncmdl.AppLogger.Infoln("Song Album", songAlbum)
		songArtist := e.ChildText("h3.albumTitle span:nth-child(2)")
		ncmdl.AppLogger.Infoln("Song Artist", songArtist)
	})

	songColly.OnHTML("div#tralbumArt", func(e *colly.HTMLElement) {
		songCover := e.ChildAttr("a", "href")
		ncmdl.AppLogger.Infoln("Song Cover", songCover)
	})

	// songColly.OnHTML("script[data-tralbum]", func(e *colly.HTMLElement) {
	// 	ncmdl.AppLogger.Infoln("Song Data", e.Attr("data-tralbum"))
	// })
}

func TestCollyAlbum(t *testing.T) {
	err := albumColly.Visit(testAlbumLink)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCollySong(t *testing.T) {
	err := songColly.Visit(testSongLink)
	if err != nil {
		t.Error(err)
		return
	}
}
