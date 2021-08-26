package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ifosch/stationery/pkg/gdrive"
	"github.com/ifosch/stationery/pkg/stationery"
)

func main() {
	log.SetFlags(0)
	q := ""
	if len(os.Args[1:]) > 0 {
		q = strings.Join(os.Args[1:], " ")
	}

	var err error

	svc, err := gdrive.GetService(os.Getenv("DRIVE_CREDENTIALS_FILE"))
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	r, err := stationery.GetFiles(svc, q)
	if err != nil {
		log.Fatal(err)
	}

	if q == "" {
		fmt.Printf("Files:\n%v", stationery.BuildFileList(r))
	} else {
		fmt.Printf("Matching files:\n%v", stationery.BuildFileList(r))
	}
}
