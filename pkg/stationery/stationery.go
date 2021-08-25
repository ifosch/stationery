package stationery

import (
	"fmt"

	"google.golang.org/api/drive/v3"
)

// Service is a description of a valid service related to a document store, like Google Drive.
// It is intended for make testing easier.
// Methods supported:
//  - `GetFiles`: based on a string parameter to specify a query, it returns a list of files metadata, or an error.
type Service interface {
	GetFiles(string) ([]*drive.File, error)
	ExportFile(*drive.File, string) (string, error)
}

// ExportHTML gets an HTML version of the specified file's content.
func ExportHTML(svc Service, file *drive.File) (content string, err error) {
	return svc.ExportFile(file, "text/html")
}

// BuildFileList creates a printable output from a list of files.
func BuildFileList(r []*drive.File) (output string) {
	for _, file := range r {
		output += fmt.Sprintf(" - %v (%v)\n", file.Name, file.Id)
	}

	return
}

// GetFiles uses the `Service` to get a list of `drive.File`s, based on the `q` parameter.
// It returns the list, or whatever error coming from the `Service`, or if the list is empty.
func GetFiles(svc Service, q string) ([]*drive.File, error) {
	r, err := svc.GetFiles(q)
	if err != nil {
		return nil, err
	}

	if len(r) < 1 {
		return nil, fmt.Errorf("no files found")
	}

	return r, nil
}
