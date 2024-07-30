package pvpn

import "testing"

func TestMatch(t *testing.T) {
	f := &Filters{}
	s := LogicalServers{
		LogicalServer{EntryCountry: "nonmatch", ExitCountry: "nonmatch"},
		LogicalServer{EntryCountry: "nonmatch", ExitCountry: "match"},
		LogicalServer{EntryCountry: "match", ExitCountry: "nonmatch"},
		LogicalServer{EntryCountry: "match", ExitCountry: "match"},
	}
	f.EntryCountry("match").ExitCountry("match")
	if f.Match(s[0]) {
		t.Fatal("unexpected match")
	}
	if f.Match(s[1]) {
		t.Fatal("unexpected match")
	}
	if f.Match(s[2]) {
		t.Fatal("unexpected match")
	}
	if !f.Match(s[3]) {
		t.Fatal("unexpected non-match")
	}
}

func TestEntryCountry_Filter(t *testing.T) {
	f := &Filters{}
	s := LogicalServers{
		LogicalServer{EntryCountry: "nonmatch"},
		LogicalServer{EntryCountry: "match"},
	}
	f.EntryCountry("match")
	if f.Match(s[0]) {
		t.Fatal("unexpected match")
	}
	if !f.Match(s[1]) {
		t.Fatal("unexpected non-match")
	}
}

func TestExitCountry_Filter(t *testing.T) {
	f := &Filters{}
	s := LogicalServers{
		LogicalServer{ExitCountry: "nonmatch"},
		LogicalServer{ExitCountry: "match"},
	}
	f.ExitCountry("match")
	if f.Match(s[0]) {
		t.Fatal("unexpected match")
	}
	if !f.Match(s[1]) {
		t.Fatal("unexpected non-match")
	}
}

func TestCity_Filter(t *testing.T) {
	f := &Filters{}
	s := LogicalServers{
		LogicalServer{City: "nonmatch"},
		LogicalServer{City: "match"},
	}
	f.City("match")
	if f.Match(s[0]) {
		t.Fatal("unexpected match")
	}
	if !f.Match(s[1]) {
		t.Fatal("unexpected non-match")
	}
}

func TestTier_Filter(t *testing.T) {
	f := &Filters{}
	s := LogicalServers{
		LogicalServer{Tier: 1},
		LogicalServer{Tier: 2},
	}
	f.Tier(2)
	if f.Match(s[0]) {
		t.Fatal("unexpected match")
	}
	if !f.Match(s[1]) {
		t.Fatal("unexpected non-match")
	}
}

func TestMaxLoad_Filter(t *testing.T) {
	f := &Filters{}
	s := LogicalServers{
		LogicalServer{Load: 1},
		LogicalServer{Load: 2},
		LogicalServer{Load: 3},
	}
	f.MaxLoad(2)
	if !f.Match(s[0]) {
		t.Fatal("unexpected non-match")
	}
	if !f.Match(s[1]) {
		t.Fatal("unexpected non-match")
	}
	if f.Match(s[2]) {
		t.Fatal("unexpected match")
	}
}

func TestMaxScore_Filter(t *testing.T) {
	f := &Filters{}
	s := LogicalServers{
		LogicalServer{Score: 1.0},
		LogicalServer{Score: 2.5},
		LogicalServer{Score: 2.6},
	}
	f.MaxScore(2.5)
	if !f.Match(s[0]) {
		t.Fatal("unexpected non-match")
	}
	if !f.Match(s[1]) {
		t.Fatal("unexpected non-match")
	}
	if f.Match(s[2]) {
		t.Fatal("unexpected match")
	}
}

func TestStatus_Filter(t *testing.T) {
	f := &Filters{}
	s := LogicalServers{
		LogicalServer{Status: 1},
		LogicalServer{Status: 2},
	}
	f.Status(2)
	if f.Match(s[0]) {
		t.Fatal("unexpected match")
	}
	if !f.Match(s[1]) {
		t.Fatal("unexpected non-match")
	}
}

func TestIncludesFeatures_Filter(t *testing.T) {
	f := &Filters{}
	s := LogicalServers{
		LogicalServer{Features: FeatureP2P},
		LogicalServer{Features: FeatureTor | FeatureP2P},
		LogicalServer{Features: FeatureTor | FeatureSecureCore},
	}
	f.IncludesFeatures(FeatureP2P)
	if !f.Match(s[0]) {
		t.Fatal("unexpected non-match")
	}
	if !f.Match(s[1]) {
		t.Fatal("unexpected non-match")
	}
	if f.Match(s[2]) {
		t.Fatal("unexpected match")
	}
}

func TestExactFeatures_Filter(t *testing.T) {
	f := &Filters{}
	s := LogicalServers{
		LogicalServer{Features: FeatureTor},
		LogicalServer{Features: FeatureTor | FeatureP2P},
		LogicalServer{Features: FeatureP2P | FeatureSecureCore},
	}
	f.ExactFeatures(FeatureTor | FeatureP2P)
	if f.Match(s[0]) {
		t.Fatal("unexpected match")
	}
	if !f.Match(s[1]) {
		t.Fatal("unexpected non-match")
	}
	if f.Match(s[2]) {
		t.Fatal("unexpected match")
	}
}
