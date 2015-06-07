package main

import (
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/qiniu/log"
)

// WriteTemp writes out a temporary file at /dev/shm with a random prefix,
// and returns the file name.
func (s *server) WriteTemp(name string, data []byte) (string, error) {
	id := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, id); err != nil {
		panic(err) // This shouldn't happen
	}
	fileName := "/dev/shm/" + hex.EncodeToString(id) + "-" + name
	err := ioutil.WriteFile(fileName, data, FileMode.ModeTemporary)
	if err != nil {
		return nil, err
	}
	return fileName, nil
}
