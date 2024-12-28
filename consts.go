package bcdl

import (
	"github.com/caiknife/mp3lister/lib/types"
)

const (
	localProxy = "socks5://127.0.0.1:7897"
)

const (
	ErrInputLinksAreEmpty  types.Error = "请输入下载链接"
	ErrInvalidBandCampLink types.Error = "请输入正确的BandCamp链接"
	ErrInvalidLinkType     types.Error = "该链接不能下载歌曲"
	ErrDownloadLinkIsEmpty types.Error = "下载链接为空"
)
