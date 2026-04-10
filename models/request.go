package models

type APIRequest interface {
	SetPartnerTag(tag string)
	SetMarketplace(marketplace string)
}

// BaseRequest contains common fields for PA-API requests
type BaseRequest struct {
	Marketplace string `json:"marketplace,omitempty"`
	PartnerTag  string `json:"partnerTag"`
	PartnerType string `json:"partnerType,omitempty"`
}

func (r *BaseRequest) SetPartnerTag(tag string) {
	r.PartnerTag = tag
}

func (r *BaseRequest) SetMarketplace(m string) {
	r.Marketplace = m
}

type GetItemsRequest struct {
	BaseRequest
	Condition             Condition  `json:"condition,omitempty"`
	CurrencyOfPreference  string     `json:"currencyOfPreference,omitempty"`
	ItemIdType            ItemIdType `json:"itemIdType,omitempty"`
	ItemIds               []string   `json:"itemIds"` // Required (max 10)
	LanguagesOfPreference []string   `json:"languagesOfPreference,omitempty"`
	Resources             []Resource `json:"resources,omitempty"`
}

type SearchItemsRequest struct {
	BaseRequest
	BrowseNodeId          string     `json:"browseNodeId,omitempty"`
	Condition             Condition  `json:"condition,omitempty"`
	CurrencyOfPreference  string     `json:"currencyOfPreference,omitempty"`
	ItemCount             int        `json:"itemCount,omitempty"`
	ItemPage              int        `json:"itemPage,omitempty"`
	Keywords              string     `json:"keywords,omitempty"`
	LanguagesOfPreference []string   `json:"languagesOfPreference,omitempty"`
	Resources             []Resource `json:"resources,omitempty"`
	SearchIndex           string     `json:"searchIndex,omitempty"`
	SortBy                string     `json:"sortBy,omitempty"`
}

type GetVariationsRequest struct {
	BaseRequest
	ASIN                  string     `json:"asin"` // Required
	Condition             Condition  `json:"condition,omitempty"`
	CurrencyOfPreference  string     `json:"currencyOfPreference,omitempty"`
	LanguagesOfPreference []string   `json:"languagesOfPreference,omitempty"`
	Resources             []Resource `json:"resources,omitempty"`
}

type GetBrowseNodesRequest struct {
	BaseRequest
	BrowseNodeIds         []string   `json:"browseNodeIds"` // Required
	LanguagesOfPreference []string   `json:"languagesOfPreference,omitempty"`
	Resources             []Resource `json:"resources,omitempty"`
}
