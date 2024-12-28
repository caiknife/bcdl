package bcdl

import (
	"testing"
)

func TestAlbum_Fetch(t *testing.T) {
	a := NewAlbum(testAlbumLink)
	err := a.c.SetProxy(localProxy)
	if err != nil {
		t.Error(err)
		return
	}
	err = a.Fetch()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(a.Info())
}
