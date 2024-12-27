package bcdl

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type Album struct {
	Title         string                     `json:"title"`
	Artist        string                     `json:"artist"`
	Cover         string                     `json:"album"`
	URL           string                     `json:"link"`
	DownloadItems types.Slice[*DownloadItem] `json:"download_items"`
}

func (a *Album) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}

type AlbumJSON struct {
	ForTheCurious string `json:"for the curious"`
	Current       struct {
		Audit               int         `json:"audit"`
		Title               string      `json:"title"`
		NewDate             string      `json:"new_date"`
		ModDate             string      `json:"mod_date"`
		PublishDate         string      `json:"publish_date"`
		Private             interface{} `json:"private"`
		Killed              interface{} `json:"killed"`
		DownloadPref        int         `json:"download_pref"`
		RequireEmail        interface{} `json:"require_email"`
		IsSetPrice          interface{} `json:"is_set_price"`
		SetPrice            float64     `json:"set_price"`
		MinimumPrice        float64     `json:"minimum_price"`
		MinimumPriceNonzero float64     `json:"minimum_price_nonzero"`
		RequireEmail0       interface{} `json:"require_email_0"`
		Artist              interface{} `json:"artist"`
		About               string      `json:"about"`
		Credits             interface{} `json:"credits"`
		AutoRepriced        interface{} `json:"auto_repriced"`
		NewDescFormat       int         `json:"new_desc_format"`
		BandId              int         `json:"band_id"`
		SellingBandId       int         `json:"selling_band_id"`
		ArtId               int         `json:"art_id"`
		DownloadDescId      interface{} `json:"download_desc_id"`
		ReleaseDate         string      `json:"release_date"`
		Upc                 interface{} `json:"upc"`
		PurchaseUrl         interface{} `json:"purchase_url"`
		PurchaseTitle       interface{} `json:"purchase_title"`
		FeaturedTrackId     int64       `json:"featured_track_id"`
		Id                  int         `json:"id"`
		Type                string      `json:"type"`
	} `json:"current"`
	PreorderCount        interface{} `json:"preorder_count"`
	HasAudio             bool        `json:"hasAudio"`
	ArtId                int         `json:"art_id"`
	Packages             interface{} `json:"packages"`
	DefaultPrice         float64     `json:"defaultPrice"`
	FreeDownloadPage     interface{} `json:"freeDownloadPage"`
	FREE                 int         `json:"FREE"`
	PAID                 int         `json:"PAID"`
	Artist               string      `json:"artist"`
	ItemType             string      `json:"item_type"`
	Id                   int         `json:"id"`
	LastSubscriptionItem interface{} `json:"last_subscription_item"`
	HasDiscounts         bool        `json:"has_discounts"`
	IsBonus              interface{} `json:"is_bonus"`
	PlayCapData          struct {
		StreamingLimitsEnabled bool `json:"streaming_limits_enabled"`
		StreamingLimit         int  `json:"streaming_limit"`
	} `json:"play_cap_data"`
	IsPurchased                interface{} `json:"is_purchased"`
	ItemsPurchased             interface{} `json:"items_purchased"`
	IsPrivateStream            interface{} `json:"is_private_stream"`
	IsBandMember               interface{} `json:"is_band_member"`
	LicensedVersionIds         interface{} `json:"licensed_version_ids"`
	PackageAssociatedLicenseId interface{} `json:"package_associated_license_id"`
	HasVideo                   interface{} `json:"has_video"`
	TralbumSubscriberOnly      bool        `json:"tralbum_subscriber_only"`
	FeaturedTrackId            int64       `json:"featured_track_id"`
	InitialTrackNum            interface{} `json:"initial_track_num"`
	IsPreorder                 bool        `json:"is_preorder"`
	AlbumIsPreorder            bool        `json:"album_is_preorder"`
	AlbumReleaseDate           string      `json:"album_release_date"`
	Trackinfo                  []struct {
		Id      int64 `json:"id"`
		TrackId int64 `json:"track_id"`
		File    struct {
			Mp3128 string `json:"mp3-128"`
		} `json:"file"`
		Artist            interface{} `json:"artist"`
		Title             string      `json:"title"`
		EncodingsId       int64       `json:"encodings_id"`
		LicenseType       int         `json:"license_type"`
		Private           interface{} `json:"private"`
		TrackNum          int         `json:"track_num"`
		AlbumPreorder     bool        `json:"album_preorder"`
		UnreleasedTrack   bool        `json:"unreleased_track"`
		TitleLink         string      `json:"title_link"`
		HasLyrics         bool        `json:"has_lyrics"`
		HasInfo           bool        `json:"has_info"`
		Streaming         int         `json:"streaming"`
		IsDownloadable    bool        `json:"is_downloadable"`
		HasFreeDownload   interface{} `json:"has_free_download"`
		FreeAlbumDownload bool        `json:"free_album_download"`
		Duration          float64     `json:"duration"`
		Lyrics            interface{} `json:"lyrics"`
		SizeofLyrics      int         `json:"sizeof_lyrics"`
		IsDraft           bool        `json:"is_draft"`
		VideoSourceType   interface{} `json:"video_source_type"`
		VideoSourceId     interface{} `json:"video_source_id"`
		VideoMobileUrl    interface{} `json:"video_mobile_url"`
		VideoPosterUrl    interface{} `json:"video_poster_url"`
		VideoId           interface{} `json:"video_id"`
		VideoCaption      interface{} `json:"video_caption"`
		VideoFeatured     interface{} `json:"video_featured"`
		AltLink           interface{} `json:"alt_link"`
		EncodingError     interface{} `json:"encoding_error"`
		EncodingPending   interface{} `json:"encoding_pending"`
		PlayCount         int         `json:"play_count"`
		IsCapped          bool        `json:"is_capped"`
		TrackLicenseId    interface{} `json:"track_license_id"`
	} `json:"trackinfo"`
	PlayingFrom      string `json:"playing_from"`
	Url              string `json:"url"`
	UseExpandoLyrics bool   `json:"use_expando_lyrics"`
}
