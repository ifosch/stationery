package gdrive

import (
	"context"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Service is an abstraction object for the Google Drive service.
// It is intended to satisfy `docs.Service` interface.
type Service struct {
	service *drive.Service
}

// GetFiles receives a query parameter and returns a list of `drive.File`s or an error.
func (s *Service) GetFiles(q string) (files []*drive.File, err error) {
	fl, err := s.service.Files.List().Fields("nextPageToken, files(id, name)").Q(q).Do()
	if err != nil {
		return nil, err
	}

	return fl.Files, nil
}

// GetService creates a new `drive.Service` from a credentials file.
// The argument is just the path to a credentials file, which are expected to identify a service account.
// If the argument is empty, and on any other error, an error is returned.
func GetService(credentialsFile string) (*Service, error) {
	if credentialsFile == "" {
		return nil, fmt.Errorf("no credentials file specified")
	}

	ctx := context.Background()
	b, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, drive.DriveMetadataReadonlyScope, drive.DriveReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("invalid client secret file: %v", err)
	}
	client := config.Client(ctx)

	svc, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials: %v", err)
	}

	return &Service{
		service: svc,
	}, nil
}
