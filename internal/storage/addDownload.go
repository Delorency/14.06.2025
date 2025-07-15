package storage

import "fmt"

func (*storage) AddDownload(task *archive) string {
	if task.Status == StatusCompleted {
		return fmt.Sprintf("/download/%s", task.ZipName)
	}
	return ""
}
