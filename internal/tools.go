package internal

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

func IsValidExtension(filename string, allowed []string) bool {
	ext := strings.ToLower(filepath.Ext(filename))

	for _, a := range allowed {
		if ext == strings.TrimSpace(a) {
			return true
		}
	}
	return false
}

func DownloadFile(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ошибка получения файла: %d", res.StatusCode)
	}
	return io.ReadAll(res.Body)
}

func CreateZipArchive(files map[string][]byte) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	for filename, data := range files {
		writer, err := zipWriter.Create(filename)
		if err != nil {
			return nil, err
		}
		if _, err := writer.Write(data); err != nil {
			return nil, err
		}
	}

	if err := zipWriter.Close(); err != nil {
		return nil, err
	}

	return buf, nil
}
