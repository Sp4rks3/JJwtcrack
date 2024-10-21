package core

import (
	"JJwtcrack/cmd"
	"github.com/jessevdk/go-flags"
	"log"
)

var opts cmd.Options

func Run() {

	if _, err := flags.Parse(&opts); err != nil {
		return
	}
	weakKeys, err := readWeakKeysFromFile(opts.List)
	if err != nil {
		log.Fatalf("Error reading weak keys from file: %v", err)
	}

	for _, weakKey := range weakKeys {
		checkWeakKey(opts.JwtToken, weakKey)
	}
}
