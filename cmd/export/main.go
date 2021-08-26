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

	if len(q) == 0 {
		log.Fatal("No query specified: this command returns only one document, please add a query returning one single document.")
	}

	r, err := stationery.GetFiles(svc, q)
	if err != nil {
		log.Fatalf("Invalid query: %v", err)
	}

	if len(r) > 1 {
		log.Fatalf("Too many results: query must return only one document, not %v.", len(r))
	}

	content, err := stationery.ExportHTML(svc, r[0])
	if err != nil {
		log.Fatalf("Export failed: %v", err)
	}

	fmt.Println(content)
}
