package main

import (
	"log"
	"net/http"

	"github.com/minio/selfupdate"
)

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	/*
		checksum, err := hex.DecodeString("67c37f96f6db52b5741f8e5fa0db2198e4209e638fb5293d80ed26980a2b7e8d")
		if err != nil {
			log.Panic(err)
		}
	*/
	err = selfupdate.Apply(resp.Body, selfupdate.Options{
		Patcher: selfupdate.NewBSDiffPatcher(),
		// Hash:     crypto.SHA256,
		// Checksum: checksum,
	})
	if err != nil {
		// error handling
		log.Panic(err)
	}
	return err
}

func main() {
	defer log.Println("Hello Version 1.0")
	doUpdate("http://localhost:8080/new1")
}
