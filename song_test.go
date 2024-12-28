package bcdl

import (
	"testing"
)

func TestSong_Fetch(t *testing.T) {
	s := NewSong(testSongLink)
	err := s.c.SetProxy(localProxy)
	if err != nil {
		t.Error(err)
		return
	}

	err = s.Fetch()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(s.Info())
}
