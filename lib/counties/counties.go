package counties

type County struct {
	StateName     string `toml:"state_name"`
	StateAbbrev   string `toml:"state_abbreviation"`
	CountyName    string `toml:"county_name"`
	AppraiserSite string `toml:"appraiser_site"`
}

type Counties struct {
	Counties []County `toml:"county"`
}
