package deezer

import (
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"

	"github.com/sferaggio/deezer-flac-download/config"
	"golang.org/x/crypto/blowfish"
)

func calcBfKey(songId []byte, config config.Configuration) []byte {
	preKey := []byte(config.PreKey)
	songIdHash := md5.Sum(songId)
	songIdMd5 := hex.EncodeToString(songIdHash[:])
	twoBytes := 8 * 2
	key := make([]byte, twoBytes)
	for i := 0; i < 16; i++ {
		key[i] = songIdMd5[i] ^ songIdMd5[i+16] ^ preKey[i]
	}
	return key
}

func blowfishDecrypt(data []byte, key []byte, config config.Configuration) ([]byte, error) {
	iv, err := hex.DecodeString(config.Iv)
	if err != nil {
		return nil, err
	}
	c, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCDecrypter(c, iv)
	res := make([]byte, len(data))
	cbc.CryptBlocks(res, data)
	return res, nil
}
