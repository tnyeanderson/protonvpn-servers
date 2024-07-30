package pvpn

const (
	VPNServersEndpoint = "https://api.protonmail.ch/vpn/logicals"
)

const (
	TierFree = 0
	TierPlus = 2
	TierPM   = 3
)

const (
	FeatureSecureCore = 1 << iota
	FeatureTor
	FeatureP2P
	FeatureStreaming
	FeatureIPV6
)
