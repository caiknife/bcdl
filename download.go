package bcdl

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"sync"

	"github.com/bogem/id3v2/v2"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/ncmdl/v2"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
)

func AsyncDownload(songs types.Slice[*DownloadItem], destDir string) error {
	pool, err := ants.NewPool(DefaultPoolSize)
	if err != nil {
		err = errors.WithMessage(err, "ants pool init")
		return err
	}
	defer pool.Release()

	wg := &sync.WaitGroup{}
	songs.ForEach(func(info *DownloadItem, i int) {
		wg.Add(1)
		err := pool.Submit(func() {
			defer wg.Done()
			err := DownloadFile(info.Link, info, destDir)
			if err != nil {
				err = errors.WithMessage(err, "download file")
				ncmdl.AppLogger.Errorln(err)
				return
			}
		})
		if err != nil {
			err = errors.WithMessage(err, "ant pool submit task")
			ncmdl.AppLogger.Errorln(err)
		}
	})
	wg.Wait()
	return nil
}

func DownloadFile(url string, item *DownloadItem, destDir string) error {
	url = strings.TrimSpace(url)
	if url == "" {
		return ErrDownloadLinkIsEmpty
	}
	path := filepath.Join(destDir, item.SavePath())
	if !fileutil.IsExist(path) {
		err := fileutil.CreateDir(path)
		if err != nil {
			err = errors.WithMessage(err, "create dir")
			return err
		}
	}

	mp3File := filepath.Join(destDir, item.SaveFileName())
	if fileutil.IsExist(mp3File) {
		ncmdl.AppLogger.Warnln(item.FileName(), "文件已经存在")
		return nil
	}

	ncmdl.AppLogger.Infoln("开始下载文件", item.FileName())
	err := netutil.DownloadFile(mp3File, url)
	if err != nil {
		err = errors.WithMessage(err, "net download music file")
		return err
	}

	err = WriteTag(mp3File, item)
	if err != nil {
		err = errors.WithMessage(err, "write tag")
		return err
	}
	return nil
}

func WriteTag(filePath string, item *DownloadItem) error {
	ncmdl.AppLogger.Infoln("正在写入标签", item.FileName())
	open, err := id3v2.Open(filePath, id3v2.Options{Parse: true})
	if err != nil {
		err = errors.WithMessage(err, "id3 open file")
		return err
	}
	defer open.Close()

	open.SetDefaultEncoding(id3v2.EncodingUTF8)
	open.SetAlbum(item.Album)
	open.SetTitle(item.Title)
	open.SetArtist(item.AllArtistsTag())
	open.AddTextFrame("TRCK", id3v2.EncodingUTF8, fmt.Sprintf("%2d", item.TrackNum))

	// 专辑封面图片
	response, err := netutil.HttpGet(item.CoverURL())
	if err != nil {
		err = errors.WithMessage(err, "net http get album pic url")
		return err
	}
	defer response.Body.Close()

	pic, err := io.ReadAll(response.Body)
	if err != nil {
		err = errors.WithMessage(err, "read album pic content")
		return err
	}
	cover := id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    "image/jpeg",
		PictureType: id3v2.PTFrontCover,
		Description: item.Album,
		Picture:     pic,
	}
	open.AddAttachedPicture(cover)

	err = open.Save()
	if err != nil {
		err = errors.WithMessage(err, "save id3")
		return err
	}
	return nil
}
