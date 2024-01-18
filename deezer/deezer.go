package deezer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/sferaggio/deezer-flac-download/config"
)

var lastReqTime int64

var REQ_MIN_INTERVAL int64 = 500000000

func makeReq(method, url string, body io.Reader, config config.Configuration) (*http.Response, error) {
	var err error

	tDiff := time.Now().UnixNano() - lastReqTime
	if tDiff < REQ_MIN_INTERVAL {
		time.Sleep(time.Duration(REQ_MIN_INTERVAL-tDiff) * time.Nanosecond)
	}
	lastReqTime = time.Now().UnixNano()

	shortUrl := url
	if len(shortUrl) > 80 {
		shortUrl = shortUrl[:80] + "..."
	}
	log.Printf("%s %s\n", method, shortUrl)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Origin", "https://www.deezer.com")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://www.deezer.com/")
	req.Header.Add("DNT", "1")
	cookie := &http.Cookie{
		Name:  "arl",
		Value: config.Arl,
	}
	req.AddCookie(cookie)

	var res *http.Response
	res, err = http.DefaultClient.Do(req)
	for err != nil {
		log.Print("(network hiccup)")
		res, err = http.DefaultClient.Do(req)
	}
	return res, err
}

func getAlbum(albumId string, config config.Configuration) (ResAlbum, error) {
	url := fmt.Sprintf("https://api.deezer.com/album/%s", albumId)
	res, err := makeReq("GET", url, nil, config)
	if err != nil {
		return ResAlbum{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bytes, _ := io.ReadAll(res.Body)
		log.Println(string(bytes))
		return ResAlbum{}, fmt.Errorf("got status code %d", res.StatusCode)
	}

	var album ResAlbum
	err = json.NewDecoder(res.Body).Decode(&album)
	return album, err
}

func getAlbumSongs(albumId string, config config.Configuration) (ResAlbumInfo, error) {
	url := fmt.Sprintf("https://www.deezer.com/de/album/%s", albumId)

	res, err := makeReq("GET", url, nil, config)
	if err != nil {
		return ResAlbumInfo{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bytes, _ := io.ReadAll(res.Body)
		log.Println(string(bytes))
		return ResAlbumInfo{}, fmt.Errorf("got status code %d", res.StatusCode)
	}

	bytes, _ := io.ReadAll(res.Body)
	s := string(bytes)

	startMarker := `window.__DZR_APP_STATE__ = `
	endMarker := `</script>`
	startIdx := strings.Index(s, startMarker)
	endIdx := strings.Index(s[startIdx:], endMarker)
	sData := s[startIdx+len(startMarker) : startIdx+endIdx]

	var albumInfo ResAlbumInfo
	if err := json.NewDecoder(strings.NewReader(sData)).Decode(&albumInfo); err != nil {
		log.Printf("failed to decode album data")
	}
	return albumInfo, nil
}

func getTracks(trackID string, config config.Configuration) (ResSongInfo, error) {
	url := fmt.Sprintf("https://www.deezer.com/de/track/%s", trackID)

	res, err := makeReq("GET", url, nil, config)
	if err != nil {
		return ResSongInfo{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bytes, _ := io.ReadAll(res.Body)
		log.Println(string(bytes))
		return ResSongInfo{}, fmt.Errorf("got status code %d", res.StatusCode)
	}

	bytes, _ := io.ReadAll(res.Body)
	s := string(bytes)

	startMarker := `window.__DZR_APP_STATE__ = `
	endMarker := `</script>`
	startIdx := strings.Index(s, startMarker)
	endIdx := strings.Index(s[startIdx:], endMarker)
	sData := s[startIdx+len(startMarker) : startIdx+endIdx]

	var songInfo ResSongInfo
	if err := json.NewDecoder(strings.NewReader(sData)).Decode(&songInfo); err != nil {
		return ResSongInfo{}, err
	}
	return songInfo, nil
}

func getSongUrlData(trackToken string, config config.Configuration) (ResSongUrl, error) {
	url := "https://media.deezer.com/v1/get_url"
	bodyJsonStr := fmt.Sprintf(`{"license_token":"%s","media":[{"type":"FULL","formats":[{"cipher":"BF_CBC_STRIPE","format":"FLAC"}]}],"track_tokens":["%s"]}`, config.LicenseToken, trackToken)
	res, err := makeReq("POST", url, bytes.NewBuffer([]byte(bodyJsonStr)), config)
	if err != nil {
		return ResSongUrl{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bytes, _ := io.ReadAll(res.Body)
		log.Println(string(bytes))
		return ResSongUrl{}, fmt.Errorf("got status code %d", res.StatusCode)
	}

	var songUrlData ResSongUrl
	err = json.NewDecoder(res.Body).Decode(&songUrlData)

	if len(songUrlData.Data) == 0 {
		return ResSongUrl{}, fmt.Errorf("got empty Data array when trying to get song URL")
	}

	if len(songUrlData.Data[0].Errors) > 0 {
		return ResSongUrl{}, fmt.Errorf("got error when trying to get song URL: %s", songUrlData.Data[0].Errors[0].Message)
	}
	return songUrlData, err
}

func getSongUrl(songUrlData ResSongUrl) (string, error) {
	if len(songUrlData.Data) == 0 || len(songUrlData.Data[0].Media) == 0 {
		spew.Fprintf(os.Stderr, "Unexpected songUrlData: %+v\n", songUrlData)
		return "", errors.New("no FLAC version available for this song")
	}
	sources := songUrlData.Data[0].Media[0].Sources
	for _, source := range sources {
		if source.Provider == "ak" {
			return source.Url, nil
		}
	}
	return sources[0].Url, nil
}

func getTitle(song ResSongInfoData) string {
	if song.Version == "" {
		return song.SngTitle
	}

	return fmt.Sprintf("%s %s", song.SngTitle, song.Version)
}

func getArtist(song ResSongInfoData) string {
	artistNames := make([]string, 0)
	for _, artist := range song.Artists {
		artistNames = append(artistNames, artist.ArtName)
	}
	sort.Strings(artistNames)
	fullArtist := strings.Join(artistNames, ", ")
	return fullArtist
}

func getComposer(song ResSongInfoData) string {
	if song.SngContributors.Composer == nil {
		return ""
	}

	composers := append([]string{}, song.SngContributors.Composer...)
	return strings.Join(composers, ", ")
}

func getSongPath(song ResSongInfoData, album ResAlbum, config config.Configuration) string {
	trackNum, err := strconv.Atoi(song.TrackNumber)
	cleanArtist := strings.ReplaceAll(album.Artist.Name, "/", "-")
	cleanAlbumTitle := strings.ReplaceAll(song.AlbTitle, "/", "-")
	cleanSongTitle := strings.ReplaceAll(song.SngTitle, "/", "-")
	if err != nil {
		panic(err)
	}
	rawPath := fmt.Sprintf("%s/%s/%s - %s [WEB FLAC]/%02d - %s.flac", config.DestDir,
		cleanArtist, cleanArtist, cleanAlbumTitle, trackNum, cleanSongTitle)
	return path.Clean(rawPath)
}

func ensureSongDirectoryExists(songPath string, coverUrl string) error {
	songDir := path.Dir(songPath)
	_, err := os.Stat(songDir)

	if !errors.Is(err, os.ErrNotExist) {
		return nil
	}

	os.MkdirAll(songDir, os.ModePerm)

	textFilePath := fmt.Sprintf("%s/info.txt", songDir)
	textFileMessage := "Downloaded from Deezer."
	textFileSourceLink := "fork: https://github.com/Shelex/deezer-flac-download"
	textFileOriginalSourceLink := "original: https://github.com/sferaggio/deezer-flac-download"
	textFileContent := strings.Join([]string{textFileMessage, textFileSourceLink, textFileOriginalSourceLink}, "\n")
	textFileData := []byte(textFileContent)
	err = os.WriteFile(textFilePath, textFileData, 0644)
	if err != nil {
		return err
	}

	if len(coverUrl) == 0 {
		log.Println("Skipping cover")
	} else {
		coverFilePath := fmt.Sprintf("%s/cover.jpg", songDir)
		f, err := os.Create(coverFilePath)
		if err != nil {
			return err
		}
		defer f.Close()
		res, err := http.Get(coverUrl)
		if err != nil {
			log.Printf("failed to get cover image: %s", err)
		}
		defer res.Body.Close()
		_, err = io.Copy(f, res.Body)
		if err != nil {
			return err
		}
	}
	return nil
}

func downloadSong(url string, songPath string, songId string, attempt int, config config.Configuration) error {
	var err error

	if attempt >= 10 {
		return fmt.Errorf("giving up downloading song after %d attempts", attempt)
	}

	f, err := os.Create(songPath)
	if err != nil {
		return err
	}
	defer f.Close()

	res, err := makeReq("GET", url, nil, config)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bytes, _ := io.ReadAll(res.Body)
		log.Println(string(bytes))
		return fmt.Errorf("got status code %d", res.StatusCode)
	}

	bfKey := calcBfKey([]byte(songId), config)
	if err != nil {
		return err
	}

	// One in every third 2048 byte block is encrypted
	blockSize := 2048
	buf := make([]byte, blockSize)
	i := 0
	nRead := 0
	totalBytes := 0
	breakNextTime := false

outer_loop:
	for {
		nRead = 0
		for nRead < blockSize {
			nNewRead, err := res.Body.Read(buf[nRead:])
			nRead += nNewRead
			totalBytes += nNewRead
			if breakNextTime {
				break outer_loop
			}
			if err == io.EOF {
				breakNextTime = true
				break
			}
			if err != nil && err != io.EOF {
				log.Printf("Error reading body on i=%d: %s\n", i, err)
				log.Println("Retrying")
				return downloadSong(url, songPath, songId, attempt+1, config)
			}
		}

		isEncrypted := ((i % 3) == 0)
		isWholeBlock := (nRead == blockSize)

		if isEncrypted && isWholeBlock {
			decBuf, err := blowfishDecrypt(buf, bfKey, config)
			if err != nil {
				return fmt.Errorf("error decrypting: %s", err)
			}
			f.Write(decBuf)
		} else {
			f.Write(buf[:nRead])
		}

		i += 1
	}

	log.Printf("Wrote %d bytes: %s", totalBytes, songPath)

	return nil
}

func DownloadAlbums(ids []string, config config.Configuration, logFile *os.File) {
album_loop:
	for idx, albumId := range ids {
		log.Printf("[%03d/%03d] Downloading album %s\n", idx+1, len(ids), albumId)
		albumInfo, err := getAlbumSongs(albumId, config)
		if err != nil {
			log.Fatalf("error getting album songs: %s\n", err)
		}

		album, err := getAlbum(albumId, config)
		if err != nil {
			log.Fatalf("error getting album: %s\n", err)
		}

		for _, song := range albumInfo.Songs.Data {
			songUrlData, err := getSongUrlData(song.TrackToken, config)

			var songUrl string
			if err == nil {
				songUrl, err = getSongUrl(songUrlData)
			}

			if err != nil {
				msg := fmt.Sprintf("error getting URL for song \"%s\" by %s from \"%s\": %s\n",
					song.SngTitle, song.ArtName, song.AlbTitle, err)
				log.Print(msg)
				logFile.Write([]byte(msg))
				log.Print("Album download failed: " + albumId + "\n\n")
				logFile.Write([]byte("Album download failed: " + albumId + "\n"))
				continue album_loop
			}
			songPath := getSongPath(song, album, config)
			songDir := path.Dir(songPath)
			coverFilePath := songDir + "/cover.jpg"

			err = ensureSongDirectoryExists(songPath, album.CoverXl)
			if err != nil {
				log.Fatalf("error preparing directory for song: %s\n", err)
			}
			err = downloadSong(songUrl, songPath, song.SngId, 0, config)
			if err != nil {
				log.Fatalf("error downloading song: %s\n", err)
			}

			err = addTags(song, songPath, album)
			if err != nil {
				log.Fatalf("error adding tags to song: %s\n", err)
			}
			err = addCover(songPath, coverFilePath)
			if err != nil {
				log.Fatalf("error adding cover image to song: %s\n", err)
			}
		}
		log.Print("Album download succeeded: " + albumId + "\n\n")
		logFile.Write([]byte("Album download succeeded: " + albumId + "\n"))
	}
}

func DownloadTracks(ids []string, config config.Configuration, logFile *os.File) {
	for idx, trackID := range ids {
		log.Printf("[%03d/%03d] Downloading track %s\n", idx+1, len(ids), trackID)
		track, err := getTracks(trackID, config)
		// debug
		// str, _ := json.MarshalIndent(track, "", "\t")
		// log.Println(string(str))
		// end debug
		if err != nil {
			log.Fatalf("error getting track: %s\n", err)
		}

		song := track.Data
		songUrlData, err := getSongUrlData(track.Data.TrackToken, config)
		if err != nil {
			log.Fatalf("error getting song url data: %s\n", err)
		}

		relatedAlbum := track.RelatedAlbums.Data[0]

		album, err := getAlbum(relatedAlbum.AlbId, config)
		if err != nil {
			log.Fatalf("error getting album: %s\n", err)
		}

		var songUrl string
		if err == nil {
			songUrl, err = getSongUrl(songUrlData)
		}

		if err != nil {
			msg := fmt.Sprintf("error getting URL for song \"%s\" by %s from \"%s\": %s\n",
				song.SngTitle, song.ArtName, song.AlbTitle, err)
			log.Print(msg)
			logFile.Write([]byte(msg))
			log.Print("Track download failed: " + trackID + "\n\n")
			logFile.Write([]byte("Track download failed: " + trackID + "\n"))
		}
		songPath := getSongPath(song, album, config)
		songDir := path.Dir(songPath)
		coverFilePath := songDir + "/cover.jpg"

		err = ensureSongDirectoryExists(songPath, album.CoverXl)
		if err != nil {
			log.Fatalf("error preparing directory for song: %s\n", err)
		}
		err = downloadSong(songUrl, songPath, song.SngId, 0, config)
		if err != nil {
			log.Fatalf("error downloading song: %s\n", err)
		}

		err = addTags(song, songPath, album)
		if err != nil {
			log.Fatalf("error adding tags to song: %s\n", err)
		}
		err = addCover(songPath, coverFilePath)
		if err != nil {
			log.Fatalf("error adding cover image to song: %s\n", err)
		}
		log.Print("Track download succeeded: " + trackID + "\n\n")
		logFile.Write([]byte("Track download succeeded: " + trackID + "\n"))
	}
}
