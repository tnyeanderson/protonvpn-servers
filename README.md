# protonvpn-servers

CLI to find, filter, and sort available ProtonVPN servers.

Install:

```sh
go install github.com/tnyeanderson/protonvpn-servers@latest
```

Usage:

```
$ protonvpn-servers -h

Fetch and filter through the list of ProtonVPN servers.

Usage:
  protonvpn-servers [flags]

Flags:
      --city string            filter by City
      --entry-country string   filter by EntryCountry
      --exact-features         feature set must match exactly, instead of inclusively
      --exit-country string    filter by ExitCountry
  -h, --help                   help for protonvpn-servers
      --ipv6                   filter by feature: IPv6
      --max-load int           filter by Load
      --max-score float        filter by Score
  -o, --output string          output format [table/json] (default "table")
      --p2p                    filter by feature: P2P
      --raw                    provide the upstream server list directly
      --secure-core            filter by feature: Secure Core
      --status int             filter by Status
      --streaming              filter by feature: Streaming
      --tier int               filter by Tier
      --tor                    filter by feature: Tor
  -v, --version                version for protonvpn-servers
```
