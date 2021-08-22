package stationery

import (
	"fmt"
	"testing"

	"google.golang.org/api/drive/v3"
)

type MockService struct {
	responses map[string][]*drive.File
}

func NewMockService(input string, expected []*drive.File) *MockService {
	return &MockService{
		responses: map[string][]*drive.File{
			input: expected,
		},
	}
}

func (ms *MockService) GetFiles(q string) ([]*drive.File, error) {
	return ms.responses[q], nil
}

func TestGetFiles(t *testing.T) {
	tc := []struct {
		input         string
		expected      []*drive.File
		expectedError error
	}{
		{
			"",
			[]*drive.File{
				{
					Name: "A file",
					Id:   "0000-00000-00000-0000",
				},
				{
					Name: "Another file",
					Id:   "1111-11111-11111-1111",
				},
			},
			nil,
		},
		{
			"name contains 'Another'",
			[]*drive.File{
				{
					Name: "Another file",
					Id:   "1111-11111-11111-1111",
				},
			},
			nil,
		},
		{
			"name contains 'Another'",
			nil,
			fmt.Errorf("no files found"),
		},
	}

	for _, tt := range tc {
		svc := NewMockService(tt.input, tt.expected)

		output, err := GetFiles(svc, tt.input)
		if err != nil && (tt.expectedError == nil || err.Error() != tt.expectedError.Error()) {
			t.Errorf("got unexpected error %v", err)
		}

		if len(output) != len(tt.expected) {
			t.Errorf("got %v, want %v", output, tt.expected)
		}
	}
}

func TestBuildFileList(t *testing.T) {
	tc := []struct {
		input    []*drive.File
		expected string
	}{
		{
			[]*drive.File{},
			"",
		},
		{
			[]*drive.File{
				{
					Name: "A file",
					Id:   "0000-00000-00000-0000",
				},
			},
			` - A file (0000-00000-00000-0000)
`,
		},
		{
			[]*drive.File{
				{
					Name: "A file",
					Id:   "0000-00000-00000-0000",
				},
				{
					Name: "Another file",
					Id:   "1111-11111-11111-1111",
				},
			},
			` - A file (0000-00000-00000-0000)
 - Another file (1111-11111-11111-1111)
`,
		},
	}

	for _, tt := range tc {
		output := BuildFileList(tt.input)

		if output != tt.expected {
			t.Errorf("got %v, want %v", output, tt.expected)
		}
	}
}
