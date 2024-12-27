package bcdl

import (
	"github.com/caiknife/mp3lister/lib/types"
)

const (
	ErrInputLinksAreEmpty  types.Error = "请输入下载链接"
	ErrInvalidBandCampLink types.Error = "请输入正确的BandCamp链接"
	ErrInvalidLinkType     types.Error = "该链接不能下载歌曲"
)
