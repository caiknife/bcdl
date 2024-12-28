package bcdl

import (
	"path/filepath"
	"strings"

	"github.com/caiknife/ncmdl/v2"
	"github.com/pkg/errors"
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
	proxy   string
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
	s := NewSong(t.Link)
	if t.proxy != "" {
		err := s.c.SetProxy(t.proxy)
		if err != nil {
			err = errors.WithMessage(err, "set proxy")
			return err
		}
	}
	err := s.Fetch()
	if err != nil {
		err = errors.WithMessage(err, "fetch song")
		return err
	}
	if t.DryRun {
		ncmdl.AppLogger.Infoln(s.Info())
		return nil
	}

	err = AsyncDownload(s.DownloadItems, t.destDir)
	if err != nil {
		err = errors.WithMessage(err, "download song")
		return err
	}

	return nil
}

func (t *Task) downloadAlbum() error {
	a := NewAlbum(t.Link)
	if t.proxy != "" {
		err := a.c.SetProxy(t.proxy)
		if err != nil {
			err = errors.WithMessage(err, "set proxy")
			return err
		}
	}
	err := a.Fetch()
	if err != nil {
		err = errors.WithMessage(err, "fetch album")
		return err
	}
	if t.DryRun {
		ncmdl.AppLogger.Infoln(a.Info())
		return nil
	}

	err = AsyncDownload(a.DownloadItems, t.destDir)
	if err != nil {
		err = errors.WithMessage(err, "download album")
		return err
	}
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

func OptionProxy(proxy string) TaskOption {
	return func(t *Task) {
		t.proxy = strings.TrimSpace(proxy)
	}
}
