package bcdl

import (
	"path/filepath"

	"github.com/caiknife/ncmdl/v2"
	"github.com/samber/lo"
)

type LinkType int

const (
	TypeSong LinkType = iota + 1
	TypeAlbum
)

type Task struct {
	Link string   `json:"link"`
	Type LinkType `json:"type"`

	Tmp    bool `json:"tmp"`
	DryRun bool `json:"dry_run"`

	destDir string
}

func NewTask(link string, opts ...TaskOption) (*Task, error) {
	var linkType LinkType
	switch {
	case IsSongLink(link):
		linkType = TypeSong
	case IsAlbumLink(link):
		linkType = TypeAlbum
	default:
		return nil, ErrInvalidBandCampLink
	}
	t := &Task{
		Link: link,
		Type: linkType,
	}
	for _, opt := range opts {
		opt(t)
	}

	t.destDir = ncmdl.Path(lo.Ternary[string](t.Tmp, filepath.Join(".", "tmp"), "."))
	return t, nil
}

func (t *Task) Download() error {
	switch t.Type {
	case TypeSong:
		return t.downloadSong()
	case TypeAlbum:
		return t.downloadAlbum()
	default:
		return ErrInvalidLinkType
	}
}

func (t *Task) downloadSong() error {
	return nil
}

func (t *Task) downloadAlbum() error {
	return nil
}

type TaskOption func(*Task)

func OptionDryRun(dryRun bool) TaskOption {
	return func(t *Task) {
		t.DryRun = dryRun
	}
}

func OptionTmp(tmp bool) TaskOption {
	return func(t *Task) {
		t.Tmp = tmp
	}
}
