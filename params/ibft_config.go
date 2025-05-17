package params

type IBFTConfig struct {
	BlockPeriodSeconds    uint64 `json:"blockperiodseconds"`
	EpochLength           uint64 `json:"epochlength"`
	RequestTimeoutSeconds uint64 `json:"requesttimeoutseconds"`
	Policy                uint64 `json:"policy"`
	Ceil2Nby3Block        uint64 `json:"ceil2Nby3Block"`
}

var DefaultIBFTConfig = IBFTConfig{
	BlockPeriodSeconds:    3,
	EpochLength:          1800,
	RequestTimeoutSeconds: 15,
	Policy:                1,
	Ceil2Nby3Block:        0,
}
