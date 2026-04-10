package locale

// Region represents the geographical macro-area (NA, EU, FE).
type Region string

const (
	RegionNA Region = "NA"
	RegionEU Region = "EU"
	RegionFE Region = "FE"
)

// Locale groups the specific configurations for an Amazon marketplace.
type Locale struct {
	Country     string
	Marketplace string
	Region      Region
}

// Predefined list of all marketplaces supported by the Creators API.
var (
	Australia          = Locale{"Australia", "www.amazon.com.au", RegionFE}
	Belgium            = Locale{"Belgium", "www.amazon.com.be", RegionEU}
	Brazil             = Locale{"Brazil", "www.amazon.com.br", RegionNA}
	Canada             = Locale{"Canada", "www.amazon.ca", RegionNA}
	Egypt              = Locale{"Egypt", "www.amazon.eg", RegionEU}
	France             = Locale{"France", "www.amazon.fr", RegionEU}
	Germany            = Locale{"Germany", "www.amazon.de", RegionEU}
	India              = Locale{"India", "www.amazon.in", RegionEU}
	Ireland            = Locale{"Ireland", "www.amazon.ie", RegionEU}
	Italy              = Locale{"Italy", "www.amazon.it", RegionEU}
	Japan              = Locale{"Japan", "www.amazon.co.jp", RegionFE}
	Mexico             = Locale{"Mexico", "www.amazon.com.mx", RegionNA}
	Netherlands        = Locale{"Netherlands", "www.amazon.nl", RegionEU}
	Poland             = Locale{"Poland", "www.amazon.pl", RegionEU}
	Singapore          = Locale{"Singapore", "www.amazon.sg", RegionFE}
	SaudiArabia        = Locale{"Saudi Arabia", "www.amazon.sa", RegionEU}
	Spain              = Locale{"Spain", "www.amazon.es", RegionEU}
	Sweden             = Locale{"Sweden", "www.amazon.se", RegionEU}
	Turkey             = Locale{"Turkey", "www.amazon.com.tr", RegionEU}
	UnitedArabEmirates = Locale{"United Arab Emirates", "www.amazon.ae", RegionEU}
	UnitedKingdom      = Locale{"United Kingdom", "www.amazon.co.uk", RegionEU}
	UnitedStates       = Locale{"United States", "www.amazon.com", RegionNA}
)
