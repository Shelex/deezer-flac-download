package deezer

import "encoding/json"

type ResSongInfoArtist struct {
	ArtName string `json:"ART_NAME"`
}

type ResSongInfoContributors struct {
	Composer  []string `json:"composer"`
	Featuring []string `json:"featuring"`
}

// contributors could be array or object
type SongInfoContributors struct {
	Contributors []*ResSongInfoContributors
	Contributor  *ResSongInfoContributors
}

func (sic *SongInfoContributors) UnmarshalJSON(data []byte) error {
	if data[0] == '[' {
		return json.Unmarshal(data, &sic.Contributors)
	}
	return json.Unmarshal(data, &sic.Contributor)
}

type ResSongInfoData struct {
	SngId               string               `json:"SNG_ID"`
	SngTitle            string               `json:"SNG_TITLE"`
	ArtName             string               `json:"ART_NAME"`
	Artists             []ResSongInfoArtist  `json:"ARTISTS"`
	AlbId               string               `json:"ALB_ID"`
	AlbTitle            string               `json:"ALB_TITLE"`
	FilesizeAac64       string               `json:"FILESIZE_AAC_64"`
	FilesizeMp364       string               `json:"FILESIZE_MP3_64"`
	FilesizeMp3128      string               `json:"FILESIZE_MP3_128"`
	FilesizeMp3256      string               `json:"FILESIZE_MP3_256"`
	FilesizeMp3320      string               `json:"FILESIZE_MP3_320"`
	FilesizeFlac        string               `json:"FILESIZE_FLAC"`
	Filesize            string               `json:"FILESIZE"`
	DiskNumber          string               `json:"DISK_NUMBER"`
	TrackNumber         string               `json:"TRACK_NUMBER"`
	TrackToken          string               `json:"TRACK_TOKEN"`
	Version             string               `json:"VERSION"`
	Isrc                string               `json:"ISRC"`
	SngContributors     SongInfoContributors `json:"SNG_CONTRIBUTORS"`
	Copyright           string               `json:"COPYRIGHT"`
	PhysicalReleaseDate string               `json:"PHYSICAL_RELEASE_DATE"`
}

type ResSongInfoIsrcData struct {
	ArtName  string `json:"ART_NAME"`
	AlbId    string `json:"ALB_ID"`
	AlbTitle string `json:"ALB_TITLE"`
}

type ResSongInfoIsrc struct {
	Data  []ResSongInfoIsrcData `json:"data"`
	Count int                   `json:"count"`
	Total int                   `json:"total"`
}

type ResSongInfoRelatedAlbumsData struct {
	ArtName  string `json:"ART_NAME"`
	AlbId    string `json:"ALB_ID"`
	AlbTitle string `json:"ALB_TITLE"`
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
			Format  string `json:"format"`
			Sources []struct {
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
		Data  []ResSongInfoData `json:"data"`
		Count int               `json:"count"`
		Total int               `json:"total"`
	} `json:"SONGS"`
}

type ResAlbumGenres struct {
	Data []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type ResAlbumContributor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ResAlbumArtist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ResAlbumTracks struct {
	Data []struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Artist struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"artist"`
		Album struct {
			ID      int    `json:"id"`
			Title   string `json:"title"`
			CoverXl string `json:"cover_xl"`
		} `json:"album"`
	} `json:"data"`
}

type ResAlbum struct {
	ID           int                   `json:"id"`
	Title        string                `json:"title"`
	CoverXl      string                `json:"cover_xl"`
	Contributors []ResAlbumContributor `json:"contributors"`
	Artist       ResAlbumArtist        `json:"artist"`
}
