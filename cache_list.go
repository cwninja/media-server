package main

import "strings"
import "path/filepath"

type EntryStatus struct {
	Key    string `json:"key"`
	Format string `json:"format"`
	Status string `json:"status"`
}

func getCacheList(mediaDir string) (output []EntryStatus, err error) {
	matches, err := filepath.Glob(filepath.Join(mediaDir, "*"))
	if err != nil {
		return
	}

	output = make([]EntryStatus, 0, len(matches))
	for _, v := range matches {
		filename := filepath.Base(v)
		parts := strings.SplitN(filename, ".", 2)
		if len(parts) == 2 {
			key, ext := parts[0], parts[1]
			l := len(output)
			output = output[:l+1]
			if strings.HasSuffix(filename, ".downloading") {
				output[l] = EntryStatus{key, "", "downloading"}
			} else {
				output[l] = EntryStatus{key, ext, "cached"}
			}
		}
	}

	return
}
