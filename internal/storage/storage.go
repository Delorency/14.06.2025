package storage

import (
	config "arch/internal"
	"bytes"
	"sync"
	"time"
)

type IStorage interface {
	AddArchive() (*archive, error)
	CompleteArchive(string)
	GetArchive(string) (*archive, bool)
	ArchiveCheck(*archive) error
	AddFileToArchive(*archive, string) error
	AddDownload(*archive) string
	ProcessArchive()
}

type TaskStatus string

var (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted  TaskStatus = "completed"
	StatusFailed     TaskStatus = "failed"
)

type archive struct {
	ID     string
	Status TaskStatus
	Files  []string
	Errors []string

	ZipBuffer *bytes.Buffer
	ZipName   string

	CreatedAt time.Time

	mu sync.Mutex
}

type storage struct {
	Archives       map[string]*archive
	ActiveArchives int
	Cfg            *config.ConfigArchive
	mu             sync.Mutex
}

func NewStorage(cfg *config.ConfigArchive) IStorage {
	return &storage{
		Archives:       make(map[string]*archive),
		Cfg:            cfg,
		ActiveArchives: 0,
	}
}
