package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/tnyeanderson/protonvpn-servers/pvpn"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var outputFormats = []string{"table", "json"}

type featureFlag struct {
	property string
	value    int
}

var featureFlags = map[string]featureFlag{
	"secure-core": featureFlag{"Secure Core", pvpn.FeatureSecureCore},
	"tor":         featureFlag{"Tor", pvpn.FeatureTor},
	"p2p":         featureFlag{"P2P", pvpn.FeatureP2P},
	"streaming":   featureFlag{"Streaming", pvpn.FeatureStreaming},
	"ipv6":        featureFlag{"IPv6", pvpn.FeatureIPV6},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "protonvpn-servers",
	Short:   "Fetch and filter through the list of ProtonVPN servers.",
	Version: "v0.0.1",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Raw output
		if raw, err := cmd.Flags().GetBool("raw"); err == nil && raw {
			out, err := pvpn.FetchJSON()
			if err != nil {
				return err
			}
			fmt.Println(string(out))
			os.Exit(0)
		}

		// Structured output
		format, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		if !slices.Contains(outputFormats, format) {
			return fmt.Errorf("invalid output format: %s", format)
		}

		servers, err := pvpn.FetchLogicalServers()
		if err != nil {
			return err
		}

		filters := parseFilterFlags(cmd.Flags())

		servers.Filter(filters).Sort()

		switch format {
		case "json":
			return servers.PrintJSON(os.Stdout)
		case "table":
			return servers.PrintTable(os.Stdout)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Bool("raw", false, "provide the upstream server list directly")
	rootCmd.Flags().StringP("output", "o", "table", fmt.Sprintf("output format [%s]", strings.Join(outputFormats, "/")))
	rootCmd.Flags().String("entry-country", "", "filter by EntryCountry")
	rootCmd.Flags().String("exit-country", "", "filter by ExitCountry")
	rootCmd.Flags().String("city", "", "filter by City")
	rootCmd.Flags().Int("tier", 0, "filter by Tier")
	rootCmd.Flags().Int("max-load", 0, "filter by Load")
	rootCmd.Flags().Float64("max-score", 0, "filter by Score")
	rootCmd.Flags().Int("status", 0, "filter by Status")
	rootCmd.Flags().Bool("exact-features", false, "feature set must match exactly, instead of inclusively")
	for name, f := range featureFlags {
		rootCmd.Flags().Bool(name, false, "filter by feature: "+f.property)
	}
}

func parseFilterFlags(flags *pflag.FlagSet) *pvpn.Filters {
	f := &pvpn.Filters{}
	features := 0

	flags.Visit(func(flag *pflag.Flag) {
		switch flag.Name {
		case "entry-country":
			if v, err := flags.GetString(flag.Name); err == nil {
				f.EntryCountry(v)
			}
		case "exit-country":
			if v, err := flags.GetString(flag.Name); err == nil {
				f.ExitCountry(v)
			}
		case "city":
			if v, err := flags.GetString(flag.Name); err == nil {
				f.City(v)
			}
		case "tier":
			if v, err := flags.GetInt(flag.Name); err == nil {
				f.Tier(v)
			}
		case "max-load":
			if v, err := flags.GetInt(flag.Name); err == nil {
				f.MaxLoad(v)
			}
		case "max-score":
			if v, err := flags.GetFloat64(flag.Name); err == nil {
				f.MaxScore(v)
			}
		case "status":
			if v, err := flags.GetInt(flag.Name); err == nil {
				f.Status(v)
			}
		}

		if featureFlag, ok := featureFlags[flag.Name]; ok {
			if v, err := flags.GetBool(flag.Name); err == nil && v {
				features = features | featureFlag.value
			}
		}
	})

	if v, err := flags.GetBool("exact-features"); err == nil && v {
		f.ExactFeatures(features)
	} else {
		if features != 0 {
			f.IncludesFeatures(features)
		}
	}

	return f
}
