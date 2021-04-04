package cmd

import "testing"

func TestChangeFileLocation(t *testing.T) {
	testingTable := []struct{
		name string
		file string
	}{
		{"no content", ""},
		{"index.html file", "index.html"},
		{"valid file", "test.html"},
		{"wrong extension", "index.txt"},
	}

	for _, tc := range testingTable {
		t.Run(tc.name, func(t *testing.T) {
			ChangeFileLocation(tc.file)
		})
	}
}