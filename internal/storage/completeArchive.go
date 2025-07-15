package storage

// Фиксируем выполнение задачи
func (s *storage) CompleteArchive(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.Archives[id]; exists {
		s.ActiveArchives--
	}
}
