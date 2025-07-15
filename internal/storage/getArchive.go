package storage

// Получаем архив по id
func (s *storage) GetArchive(id string) (*archive, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.Archives[id]
	return task, exists
}
