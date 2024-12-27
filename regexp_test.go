package bcdl

import (
	"testing"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/stretchr/testify/assert"
)

var (
	albumURLs = types.Slice[string]{
		"https://shirttailstompers.bandcamp.com/album/thats-my-kick",
		"https://campusfive.bandcamp.com/album/hummin-to-myself",
	}

	songURLs = types.Slice[string]{
		"https://shirttailstompers.bandcamp.com/track/sweets",
		"https://campusfive.bandcamp.com/track/manhattan",
	}
)

func TestIsAlbumLink(t *testing.T) {
	albumURLs.ForEach(func(s string, i int) {
		assert.True(t, IsAlbumLink(s))
	})
}

func TestIsSongLink(t *testing.T) {
	songURLs.ForEach(func(s string, i int) {
		assert.True(t, IsSongLink(s))
	})
}
