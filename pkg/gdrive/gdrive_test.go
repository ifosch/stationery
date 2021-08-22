package gdrive

import (
	"fmt"
	"testing"
)

func TestGetService(t *testing.T) {
	tc := []struct {
		input         string
		expectedError error
	}{
		{
			"",
			fmt.Errorf("no credentials file specified"),
		},
		{
			"/tmp/nofile",
			fmt.Errorf("unable to read client secret file: open /tmp/nofile: no such file or directory"),
		},
		{
			"/tmp/invalid_credentials.json",
			fmt.Errorf("unable to read client secret file: open /tmp/nofile: no such file or directory"),
		},
	}

	for _, tt := range tc {
		svc, err := GetService(tt.input)
		fmt.Println(tt.input)
		if err != nil && (tt.expectedError == nil || err.Error() != tt.expectedError.Error()) {
			t.Errorf("got unexpected error %v", err)
		}

		if svc != nil {
			t.Errorf("returned service should be nil, %v", svc)
		}
	}
}
