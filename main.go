package main

import (
	"log"
	"os"

	"github.com/sferaggio/deezer-flac-download/config"
	"github.com/sferaggio/deezer-flac-download/deezer"
)

func printUsage() {
	log.Println("deezer-flac-download is a program to freely download Deezer FLAC files.")
	log.Println("")
	log.Println("To download one or more albums:")
	log.Println("\tdeezer-flac-download album <album_id> [<album_id>...]")
	log.Println("")
	log.Println("To download one or more tracks:")
	log.Println("\tdeezer-flac-download track <track_id> [<track_id>...]")
	log.Println("")
	log.Println("See README for full details.")
}

func main() {
	var err error
	log.SetFlags(0)

	if len(os.Args) < 3 {
		printUsage()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	logFilePath := os.TempDir() + "/deezer-flac-download.log"
	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Fatalf("error creating log file %s: %s\n", logFilePath, err)
	}
	defer logFile.Close()

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error reading config file: %s\n", err)
	}

	if command == "album" {
		deezer.DownloadAlbums(args, config, logFile)
	} else if command == "track" {
		deezer.DownloadTracks(args, config, logFile)
	} else {
		printUsage()
		return
	}
}
