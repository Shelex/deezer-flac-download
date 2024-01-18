package deezer

type ResSongInfoArtist struct {
	ArtId             string      `json:"ART_ID"`
	RoleId            string      `json:"ROLE_ID"`
	ArtistsSongsOrder string      `json:"ARTISTS_SONGS_ORDER"`
	ArtName           string      `json:"ART_NAME"`
	ArtistIsDummy     bool        `json:"ARTIST_IS_DUMMY"`
	ArtPicture        string      `json:"ART_PICTURE"`
	Rank              string      `json:"RANK"`
	Locales           interface{} `json:"LOCALES"`
	Type              string      `json:"__TYPE__"`
}

type ResSongInfoMedia struct {
	Type string `json:"TYPE"`
	Href string `json:"HREF"`
}

type ResSongInfoRights struct {
	StreamAdsAvailable bool   `json:"STREAM_ADS_AVAILABLE"`
	StreamAds          string `json:"STREAM_ADS"`
	StreamSubAvailable bool   `json:"STREAM_SUB_AVAILABLE"`
	StreamSub          string `json:"STREAM_SUB"`
}

type ResSongInfoContributors struct {
	MainArtist     []string `json:"main_artist"`
	Composer       []string `json:"composer"`
	Featuring      []string `json:"featuring"`
	Narrator       []string `json:"narrator"`
	MusicPublisher []string `json:"music publisher"`
}

type ResSongInfoExplicitTrackContent struct {
	ExplicitLyricsStatus int `json:"EXPLICIT_LYRICS_STATUS"`
	ExplicitCoverStatus  int `json:"EXPLICIT_COVER_STATUS"`
}

type ResSongInfoAvailableCountries struct {
	StreamAds     []string      `json:"STREAM_ADS"`
	StreamSubOnly []interface{} `json:"STREAM_SUB_ONLY"`
}

type ResSongInfoData struct {
	SngId                string                          `json:"SNG_ID"`
	ProductTrackId       string                          `json:"PRODUCT_TRACK_ID"`
	UploadId             int                             `json:"UPLOAD_ID"`
	SngTitle             string                          `json:"SNG_TITLE"`
	ArtId                string                          `json:"ART_ID"`
	ProviderId           string                          `json:"PROVIDER_ID"`
	ArtName              string                          `json:"ART_NAME"`
	ArtistIsDummy        bool                            `json:"ARTIST_IS_DUMMY"`
	Artists              []ResSongInfoArtist             `json:"ARTISTS"`
	AlbId                string                          `json:"ALB_ID"`
	AlbTitle             string                          `json:"ALB_TITLE"`
	Type                 int                             `json:"TYPE"`
	Md5Origin            string                          `json:"MD5_ORIGIN"`
	Video                bool                            `json:"VIDEO"`
	Duration             string                          `json:"DURATION"`
	AlbPicture           string                          `json:"ALB_PICTURE"`
	ArtPicture           string                          `json:"ART_PICTURE"`
	RankSng              string                          `json:"RANK_SNG"`
	FilesizeAac64        string                          `json:"FILESIZE_AAC_64"`
	FilesizeMp364        string                          `json:"FILESIZE_MP3_64"`
	FilesizeMp3128       string                          `json:"FILESIZE_MP3_128"`
	FilesizeMp3256       string                          `json:"FILESIZE_MP3_256"`
	FilesizeMp3320       string                          `json:"FILESIZE_MP3_320"`
	FilesizeFlac         string                          `json:"FILESIZE_FLAC"`
	Filesize             string                          `json:"FILESIZE"`
	Gain                 string                          `json:"GAIN"`
	MediaVersion         string                          `json:"MEDIA_VERSION"`
	DiskNumber           string                          `json:"DISK_NUMBER"`
	TrackNumber          string                          `json:"TRACK_NUMBER"`
	TrackToken           string                          `json:"TRACK_TOKEN"`
	TrackTokenExpire     int                             `json:"TRACK_TOKEN_EXPIRE"`
	Version              string                          `json:"VERSION"`
	Media                []ResSongInfoMedia              `json:"MEDIA"`
	ExplicitLyrics       string                          `json:"EXPLICIT_LYRICS"`
	Rights               ResSongInfoRights               `json:"RIGHTS"`
	Isrc                 string                          `json:"ISRC"`
	HierarchicalTitle    string                          `json:"HIERARCHICAL_TITLE"`
	SngContributors      ResSongInfoContributors         `json:"SNG_CONTRIBUTORS"`
	LyricsId             int                             `json:"LYRICS_ID"`
	ExplicitTrackContent ResSongInfoExplicitTrackContent `json:"EXPLICIT_TRACK_CONTENT"`
	Copyright            string                          `json:"COPYRIGHT"`
	PhysicalReleaseDate  string                          `json:"PHYSICAL_RELEASE_DATE"`
	SMod                 int                             `json:"S_MOD"`
	SPremium             int                             `json:"S_PREMIUM"`
	DateStartPremium     string                          `json:"DATE_START_PREMIUM"`
	DateStart            string                          `json:"DATE_START"`
	Status               int                             `json:"STATUS"`
	UserId               int                             `json:"USER_ID"`
	URLRewriting         string                          `json:"URL_REWRITING"`
	SngStatus            string                          `json:"SNG_STATUS"`
	AvailableCountries   ResSongInfoAvailableCountries   `json:"AVAILABLE_COUNTRIES"`
	UpdateDate           string                          `json:"UPDATE_DATE"`
	Type0                string                          `json:"__TYPE__"`
	DigitalReleaseDate   string                          `json:"DIGITAL_RELEASE_DATE"`
}

type ResSongInfoIsrcData struct {
	ArtName            string            `json:"ART_NAME"`
	ArtId              string            `json:"ART_ID"`
	AlbPicture         string            `json:"ALB_PICTURE"`
	AlbId              string            `json:"ALB_ID"`
	AlbTitle           string            `json:"ALB_TITLE"`
	Duration           string            `json:"DURATION"`
	DigitalReleaseDate string            `json:"DIGITAL_RELEASE_DATE"`
	Rights             ResSongInfoRights `json:"RIGHTS"`
	LyricsId           int               `json:"LYRICS_ID"`
	Type               string            `json:"__TYPE__"`
}

type ResSongInfoIsrc struct {
	Data  []ResSongInfoIsrcData `json:"data"`
	Count int                   `json:"count"`
	Total int                   `json:"total"`
}

type ResSongInfoRelatedAlbumsData struct {
	ArtName            string            `json:"ART_NAME"`
	ArtId              string            `json:"ART_ID"`
	AlbPicture         string            `json:"ALB_PICTURE"`
	AlbId              string            `json:"ALB_ID"`
	AlbTitle           string            `json:"ALB_TITLE"`
	Duration           string            `json:"DURATION"`
	DigitalReleaseDate string            `json:"DIGITAL_RELEASE_DATE"`
	Rights             ResSongInfoRights `json:"RIGHTS"`
	LyricsId           int               `json:"LYRICS_ID"`
	Type               string            `json:"__TYPE__"`
}

type ResSongInfoRelatedAlbums struct {
	Data  []ResSongInfoRelatedAlbumsData `json:"data"`
	Count int                            `json:"count"`
	Total int                            `json:"total"`
}

type ResSongInfo struct {
	Data          ResSongInfoData          `json:"DATA"`
	Isrc          ResSongInfoIsrc          `json:"ISRC"`
	RelatedAlbums ResSongInfoRelatedAlbums `json:"RELATED_ALBUMS"`
}

type ResSongUrl struct {
	Data []struct {
		Errors []struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"errors"`
		Media []struct {
			Cipher struct {
				Type string `json:"type"`
			} `json:"cipher"`
			Exp       int    `json:"exp"`
			Format    string `json:"format"`
			MediaType string `json:"media_type"`
			Nbf       int    `json:"nbf"`
			Sources   []struct {
				Provider string `json:"provider"`
				Url      string `json:"url"`
			} `json:"sources"`
		} `json:"media"`
	} `json:"data"`
}

// This struct does not have all the fields that exist in the JSON
// because we only care about SONGS at the moment
type ResAlbumInfo struct {
	Songs struct {
		Data          []ResSongInfoData `json:"data"`
		Count         int               `json:"count"`
		Total         int               `json:"total"`
		FilteredCount int               `json:"filtered_count"`
	} `json:"SONGS"`
}

type ResAlbumGenres struct {
	Data []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
		Type    string `json:"type"`
	} `json:"data"`
}

type ResAlbumContributor struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Link          string `json:"link"`
	Share         string `json:"share"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Radio         bool   `json:"radio"`
	Tracklist     string `json:"tracklist"`
	Type          string `json:"type"`
	Role          string `json:"role"`
}

type ResAlbumArtist struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Tracklist     string `json:"tracklist"`
	Type          string `json:"type"`
}

type ResAlbumTracks struct {
	Data []struct {
		ID                    int    `json:"id"`
		Readable              bool   `json:"readable"`
		Title                 string `json:"title"`
		TitleShort            string `json:"title_short"`
		TitleVersion          string `json:"title_version"`
		Link                  string `json:"link"`
		Duration              int    `json:"duration"`
		Rank                  int    `json:"rank"`
		ExplicitLyrics        bool   `json:"explicit_lyrics"`
		ExplicitContentLyrics int    `json:"explicit_content_lyrics"`
		ExplicitContentCover  int    `json:"explicit_content_cover"`
		Preview               string `json:"preview"`
		Md5Image              string `json:"md5_image"`
		Artist                struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Tracklist string `json:"tracklist"`
			Type      string `json:"type"`
		} `json:"artist"`
		Album struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Cover       string `json:"cover"`
			CoverSmall  string `json:"cover_small"`
			CoverMedium string `json:"cover_medium"`
			CoverBig    string `json:"cover_big"`
			CoverXl     string `json:"cover_xl"`
			Md5Image    string `json:"md5_image"`
			Tracklist   string `json:"tracklist"`
			Type        string `json:"type"`
		} `json:"album"`
		Type string `json:"type"`
	} `json:"data"`
}

type ResAlbum struct {
	ID                    int                   `json:"id"`
	Title                 string                `json:"title"`
	Upc                   string                `json:"upc"`
	Link                  string                `json:"link"`
	Share                 string                `json:"share"`
	Cover                 string                `json:"cover"`
	CoverSmall            string                `json:"cover_small"`
	CoverMedium           string                `json:"cover_medium"`
	CoverBig              string                `json:"cover_big"`
	CoverXl               string                `json:"cover_xl"`
	Md5Image              string                `json:"md5_image"`
	GenreID               int                   `json:"genre_id"`
	Genres                ResAlbumGenres        `json:"genres"`
	Label                 string                `json:"label"`
	NbTracks              int                   `json:"nb_tracks"`
	Duration              int                   `json:"duration"`
	Fans                  int                   `json:"fans"`
	ReleaseDate           string                `json:"release_date"`
	RecordType            string                `json:"record_type"`
	Available             bool                  `json:"available"`
	Tracklist             string                `json:"tracklist"`
	ExplicitLyrics        bool                  `json:"explicit_lyrics"`
	ExplicitContentLyrics int                   `json:"explicit_content_lyrics"`
	ExplicitContentCover  int                   `json:"explicit_content_cover"`
	Contributors          []ResAlbumContributor `json:"contributors"`
	Artist                ResAlbumArtist        `json:"artist"`
	Type                  string                `json:"type"`
	Tracks                ResAlbumTracks        `json:"tracks"`
}
