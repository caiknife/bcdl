package bcdl

import (
	"regexp"
)

var (
	songRegexp  = regexp.MustCompile(`https://\w+\.bandcamp\.com/track/\w+`)
	albumRegexp = regexp.MustCompile(`https://\w+\.bandcamp\.com/album/\w+`)
)

func IsAlbumLink(link string) bool {
	return albumRegexp.MatchString(link)
}

func IsSongLink(link string) bool {
	return songRegexp.MatchString(link)
}
