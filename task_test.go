package bcdl

import (
	"testing"
)

func TestTask_Album(t *testing.T) {
	task, err := NewTask(testAlbumLink, OptionTmp(true), OptionProxy(localProxy))
	if err != nil {
		t.Error(err)
		return
	}
	err = task.Download()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTask_Song(t *testing.T) {
	task, err := NewTask(testSongLink, OptionTmp(true), OptionProxy(localProxy))
	if err != nil {
		t.Error(err)
		return
	}
	err = task.Download()
	if err != nil {
		t.Error(err)
		return
	}
}
