package pvpn

import (
	"testing"

	"github.com/go-test/deep"
)

func TestFilter(t *testing.T) {
	s0 := LogicalServer{EntryCountry: "CA", City: "Toronto", Tier: 0}
	s1 := LogicalServer{EntryCountry: "CA", City: "Toronto", Tier: 1}
	s2 := LogicalServer{EntryCountry: "CA", City: "Vancouver", Tier: 1}
	s := LogicalServers{s0, s1, s2}
	f := &Filters{}
	f.EntryCountry("CA").Tier(1)
	o := s.Filter(f)
	if len(s) != 3 {
		t.Fatal("filter modified the slice in-place")
	}
	if diff := deep.Equal(o, LogicalServers{s1, s2}); diff != nil {
		t.Fatal(diff)
	}
}

func TestSort(t *testing.T) {
	s0 := LogicalServer{Score: 0.5}
	s1 := LogicalServer{Score: 1}
	s2 := LogicalServer{Score: 1.5}
	s3 := LogicalServer{Score: 2}
	s := LogicalServers{s2, s1, s3, s0}
	o := s.Sort()
	if diff := deep.Equal(s, LogicalServers{s0, s1, s2, s3}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(o, LogicalServers{s0, s1, s2, s3}); diff != nil {
		t.Fatal(diff)
	}
}
