package zset

import (
	"testing"
)

var s *SortedSet

func init() {
	s = New()
}

func TestNew(t *testing.T) {
	if s == nil {
		t.Failed()
	}
	s.Set(66.0, "test1", "")
	s.Set(77.0, "test2", "")

	rank, score, extra := s.GetRank("test1", false)
	if rank == 0 {
		t.Log("Key:", "test1", "Rank:", rank, "Score:", score, "Extra:", extra)
	} else {
		t.Error("Key:", "test1", "Rank:", rank, "Score:", score, "Extra:", extra)
	}
	rank, score, extra = s.GetRank("test2", false)
	if rank == 1 {
		t.Log("Key:", "test2", "Rank:", rank, "Score:", score, "Extra:", extra)
	} else {
		t.Error("Key:", "test2", "Rank:", rank, "Score:", score, "Extra:", extra)
	}
}

