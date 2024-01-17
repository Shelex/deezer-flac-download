package deezer

import (
	"os"

	"github.com/go-flac/flacpicture"
	"github.com/go-flac/flacvorbis"
	"github.com/go-flac/go-flac"
)

func extractFlacComment(f *flac.File) (*flacvorbis.MetaDataBlockVorbisComment, int, error) {
	var err error
	var cmt *flacvorbis.MetaDataBlockVorbisComment
	var cmtIdx int
	for idx, meta := range f.Meta {
		if meta.Type == flac.VorbisComment {
			cmt, err = flacvorbis.ParseFromMetaDataBlock(*meta)
			cmtIdx = idx
			if err != nil {
				return nil, 0, err
			}
		}
	}
	return cmt, cmtIdx, nil
}

func addCover(songPath string, coverPath string) error {
	coverData, err := os.ReadFile(coverPath)
	if err != nil {
		return err
	}

	f, err := flac.ParseFile(songPath)
	if err != nil {
		return err
	}

	picture, err := flacpicture.NewFromImageData(flacpicture.PictureTypeFrontCover,
		"Front cover", coverData, "image/jpeg")
	if err != nil {
		return err
	}

	picturemeta := picture.Marshal()
	f.Meta = append(f.Meta, &picturemeta)
	f.Save(songPath)
	return nil
}

func addTags(song ResSongInfoData, path string, album ResAlbum) error {
	var err error

	f, err := flac.ParseFile(path)
	if err != nil {
		return err
	}

	cmts, idx, err := extractFlacComment(f)
	if err != nil {
		return err
	}
	if cmts == nil && idx > 0 {
		cmts = flacvorbis.New()
	}

	title := getTitle(song)
	artist := getArtist(song)
	composer := getComposer(song)

	cmts.Add("TITLE", title)
	cmts.Add("ALBUM", song.AlbTitle)
	cmts.Add("ARTIST", artist)
	cmts.Add("ALBUMARTIST", album.Artist.Name)
	cmts.Add("COMPOSER", composer)
	cmts.Add("TRACKNUMBER", song.TrackNumber)
	cmts.Add("DISCNUMBER", song.DiskNumber)
	cmts.Add("COPYRIGHT", song.Copyright)
	cmts.Add("DATE", song.PhysicalReleaseDate)
	cmts.Add("ISRC", song.Isrc)
	cmtsmeta := cmts.Marshal()
	if idx > 0 {
		f.Meta[idx] = &cmtsmeta
	} else {
		f.Meta = append(f.Meta, &cmtsmeta)
	}

	f.Save(path)

	return nil
}
