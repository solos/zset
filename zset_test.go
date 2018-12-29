package zset

import (
	"fmt"
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
	s.Set(87.0, "test3", "")
	s.Set(89.0, "test4", "")

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
	items := s.GetRange(uint64(1), uint64(2))
	fmt.Println(items)
	if items[0].Score != 77.0 {
		t.Error("Key:", "test2")
	}
	items = s.GetRange(uint64(1), uint64(4))
	fmt.Println(items)
	if items[0].Score != 77.0 {
		t.Error("Key:", "test2")
	}
	if items[1].Score != 87.0 {
		t.Error("Key:", "test3")
	}

	if items[2].Score != 89.0 {
		t.Error("Key:", "test4")
	}

}
