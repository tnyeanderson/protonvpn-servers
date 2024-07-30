package pvpn

type filterFunc func(l LogicalServer) bool

// Filters represents a set of filters to apply to a LogicalServers slice.
type Filters struct {
	filterFuncs []filterFunc
}

// Match returns true if a LogicalServer passes all of the Filters.
func (f *Filters) Match(l LogicalServer) bool {
	for _, filterFunc := range f.filterFuncs {
		if !filterFunc(l) {
			return false
		}
	}
	return true
}

// EntryCountry adds a filter based on the EntryCountry property of the
// LogicalServer.
func (f *Filters) EntryCountry(v string) *Filters {
	f.filterFuncs = append(f.filterFuncs, func(l LogicalServer) bool {
		return l.EntryCountry == v
	})
	return f
}

// ExitCountry adds a filter based on the ExitCountry property of the
// LogicalServer.
func (f *Filters) ExitCountry(v string) *Filters {
	f.filterFuncs = append(f.filterFuncs, func(l LogicalServer) bool {
		return l.ExitCountry == v
	})
	return f
}

// City adds a filter based on the City property of the LogicalServer.
func (f *Filters) City(v string) *Filters {
	f.filterFuncs = append(f.filterFuncs, func(l LogicalServer) bool {
		return l.City == v
	})
	return f
}

// Tier adds a filter based on the Tier property of the LogicalServer.
func (f *Filters) Tier(v int) *Filters {
	f.filterFuncs = append(f.filterFuncs, func(l LogicalServer) bool {
		return l.Tier == v
	})
	return f
}

// MaxLoad adds a filter based on the Load property of the LogicalServer.
func (f *Filters) MaxLoad(v int) *Filters {
	f.filterFuncs = append(f.filterFuncs, func(l LogicalServer) bool {
		return l.Load <= v
	})
	return f
}

// MaxScore adds a filter based on the Score property of the LogicalServer.
func (f *Filters) MaxScore(v float64) *Filters {
	f.filterFuncs = append(f.filterFuncs, func(l LogicalServer) bool {
		return l.Score <= v
	})
	return f
}

// Status adds a filter based on the Status property of the LogicalServer.
func (f *Filters) Status(v int) *Filters {
	f.filterFuncs = append(f.filterFuncs, func(l LogicalServer) bool {
		return l.Status == v
	})
	return f
}

// IncludesFeatures adds a filter which verifies that the provided features are
// enabled on the LogicalServer.
func (f *Filters) IncludesFeatures(v int) *Filters {
	f.filterFuncs = append(f.filterFuncs, func(l LogicalServer) bool {
		return v&l.Features == v
	})
	return f
}

// ExactFeatures adds a filter which verifies that the provided features are
// exactly the same as those on the LogicalServer.
func (f *Filters) ExactFeatures(v int) *Filters {
	f.filterFuncs = append(f.filterFuncs, func(l LogicalServer) bool {
		return l.Features == v
	})
	return f
}
