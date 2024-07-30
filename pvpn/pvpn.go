// pvpn finds and filters ProtonVPN endpoints
package pvpn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"text/tabwriter"
)

type response struct {
	LogicalServers LogicalServers
}

// Server represents a physical VPN server, which gets grouped into a
// LogicalServer.
type Server struct {
	EntryIP string
	ExitIP  string
	Domain  string
	Label   string
	Status  int
}

// LogicalServer represents a logical set of physical VPN servers.
type LogicalServer struct {
	Name         string
	EntryCountry string
	ExitCountry  string
	HostCountry  string
	Region       string
	City         string
	Domain       string
	Tier         int
	Features     int
	Status       int
	Load         int
	Score        float64
	Servers      []Server
}

// LogicalServers is a slice of LogicalServer items that can be filtered,
// sorted, and printed.
type LogicalServers []LogicalServer

// Filter returns a new LogicalServers slice containing only the items which
// match the provided Filters.
func (l LogicalServers) Filter(f *Filters) LogicalServers {
	return slices.DeleteFunc(l, func(ls LogicalServer) bool { return !f.Match(ls) })
}

// Sort modifies the LogicalServers slice in place, sorting by Score in
// ascending order. Lower scores are better. The slice is returned for
// convenience.
func (l LogicalServers) Sort() LogicalServers {
	slices.SortStableFunc(l, func(a, b LogicalServer) int {
		d := a.Score - b.Score
		if d > 0 {
			return 1
		}
		if d < 0 {
			return -1
		}
		return 0
	})
	return l
}

// PrintJSON outputs the LogicalServers in JSON format.
func (l LogicalServers) PrintJSON(out io.Writer) error {
	b, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return err
	}
	if _, err := out.Write(b); err != nil {
		return err
	}
	return nil
}

// PrintJSON outputs the LogicalServers as a table.
func (l LogicalServers) PrintTable(out io.Writer) error {
	w := tabwriter.NewWriter(out, 0, 2, 2, ' ', 0)
	w.Write([]byte(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\n", "SCORE", "LOAD", "DOMAIN", "ENTRY IP", "EXIT IP", "LABEL")))
	for _, ls := range l {
		for _, s := range ls.Servers {
			line := fmt.Sprintf("%.4v\t%v\t%v\t%v\t%v\t%v\n", ls.Score, ls.Load, ls.Domain, s.EntryIP, s.ExitIP, s.Label)
			if _, err := w.Write([]byte(line)); err != nil {
				return err
			}
		}
	}
	return w.Flush()
}

// FetchJSON retrieves the list of API servers from upstream and returns the
// resulting JSON.
func FetchJSON() ([]byte, error) {
	res, err := http.Get(VPNServersEndpoint)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(res.Body)
}

// FetchLogicalServers retrieves the list of API servers from upstream and
// returns the unmarshaled LogicalServers.
func FetchLogicalServers() (LogicalServers, error) {
	body, err := FetchJSON()
	if err != nil {
		return nil, err
	}
	r := &response{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}
	return r.LogicalServers, nil
}
