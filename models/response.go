package models

// APIError represents a single error returned by the API.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetItemsResponse is the response for GetItems.
type GetItemsResponse struct {
	Errors      []APIError   `json:"errors,omitempty"`
	ItemsResult *ItemsResult `json:"itemsResult,omitempty"`
	ItemResults *ItemsResult `json:"itemResults,omitempty"`
}

// Result is a helper method to return the ItemsResult regardless of whether Amazon used "itemsResult" or "itemResults" in the JSON response.
func (r *GetItemsResponse) Result() *ItemsResult {
	if r.ItemsResult != nil {
		return r.ItemsResult
	}
	return r.ItemResults
}

// SearchItemsResponse is the response for SearchItems.
type SearchItemsResponse struct {
	Errors       []APIError    `json:"errors,omitempty"`
	SearchResult *SearchResult `json:"searchResult,omitempty"`
}

// GetVariationsResponse is the response for GetVariations.
type GetVariationsResponse struct {
	Errors           []APIError        `json:"errors,omitempty"`
	VariationsResult *VariationsResult `json:"variationsResult,omitempty"`
}

// GetBrowseNodesResponse is the response for GetBrowseNodes.
type GetBrowseNodesResponse struct {
	Errors            []APIError         `json:"errors,omitempty"`
	BrowseNodesResult *BrowseNodesResult `json:"browseNodesResult,omitempty"`
}

type ItemsResult struct {
	Items []Item `json:"items"`
}

type SearchResult struct {
	Items             []Item             `json:"items"`
	TotalResultCount  int64              `json:"totalResultCount,omitempty"`
	SearchURL         string             `json:"searchURL,omitempty"`
	SearchRefinements *SearchRefinements `json:"searchRefinements,omitempty"`
}

// --- SearchRefinements Section ---

type SearchRefinements struct {
	BrowseNode       *Refinement  `json:"browseNode,omitempty"`
	OtherRefinements []Refinement `json:"otherRefinements,omitempty"`
	SearchIndex      *Refinement  `json:"searchIndex,omitempty"`
}

type Refinement struct {
	Id          string          `json:"id"`
	DisplayName string          `json:"displayName"`
	Bins        []RefinementBin `json:"bins"`
}

type RefinementBin struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type VariationsResult struct {
	Items            []Item            `json:"items"`
	VariationSummary *VariationSummary `json:"variationSummary,omitempty"`
}

type VariationSummary struct {
	PageCount           int64                  `json:"pageCount,omitempty"`
	Price               *VariationSummaryPrice `json:"price,omitempty"`
	VariationCount      int64                  `json:"variationCount,omitempty"`
	VariationDimensions []VariationDimension   `json:"variationDimensions,omitempty"`
}

type VariationSummaryPrice struct {
	HighestPrice *Money `json:"highestPrice,omitempty"`
	LowestPrice  *Money `json:"lowestPrice,omitempty"`
}

type VariationDimension struct {
	DisplayName string   `json:"displayName,omitempty"`
	Locale      string   `json:"locale,omitempty"`
	Name        string   `json:"name,omitempty"`
	Values      []string `json:"values,omitempty"`
}

type BrowseNodesResult struct {
	BrowseNodes []BrowseNode `json:"browseNodes"`
}

type BrowseNodeInfo struct {
	BrowseNodes      []BrowseNode      `json:"browseNodes,omitempty"`
	WebsiteSalesRank *WebsiteSalesRank `json:"websiteSalesRank,omitempty"`
}

type WebsiteSalesRank struct {
	ContextFreeName string `json:"contextFreeName,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`
	Id              string `json:"id,omitempty"`
	SalesRank       int    `json:"salesRank,omitempty"`
}

type BrowseNode struct {
	Id              string       `json:"id"`
	DisplayName     string       `json:"displayName,omitempty"`
	ContextFreeName string       `json:"contextFreeName,omitempty"`
	IsRoot          bool         `json:"isRoot,omitempty"`
	SalesRank       int          `json:"salesRank,omitempty"`
	Ancestor        *BrowseNode  `json:"ancestor,omitempty"`
	Children        []BrowseNode `json:"children,omitempty"`
}

// Item contains all information for a single product.
type Item struct {
	ASIN                string               `json:"asin"`
	DetailPageURL       string               `json:"detailPageURL"`
	ParentASIN          string               `json:"parentASIN,omitempty"`
	BrowseNodeInfo      *BrowseNodeInfo      `json:"browseNodeInfo,omitempty"`
	Images              *Images              `json:"images,omitempty"`
	ItemInfo            *ItemInfo            `json:"itemInfo,omitempty"`
	OffersV2            *OffersV2            `json:"offersV2,omitempty"`
	VariationAttributes []VariationAttribute `json:"variationAttributes,omitempty"`
}

type VariationAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// --- OffersV2 Section ---

type OffersV2 struct {
	Listings []OfferListing `json:"listings,omitempty"`
}

type OfferListing struct {
	Availability   *OfferAvailability `json:"availability,omitempty"`
	Condition      *OfferCondition    `json:"condition,omitempty"`
	DealDetails    *DealDetails       `json:"dealDetails,omitempty"`
	IsBuyBoxWinner bool               `json:"isBuyBoxWinner,omitempty"`
	LoyaltyPoints  *LoyaltyPoints     `json:"loyaltyPoints,omitempty"`
	MerchantInfo   *MerchantInfo      `json:"merchantInfo,omitempty"`
	Price          *OfferPrice        `json:"price,omitempty"`
	Type           string             `json:"type,omitempty"`
	ViolatesMAP    bool               `json:"violatesMAP,omitempty"`
}

type OfferAvailability struct {
	MaxOrderQuantity int    `json:"maxOrderQuantity,omitempty"`
	Message          string `json:"message,omitempty"`
	MinOrderQuantity int    `json:"minOrderQuantity,omitempty"`
	Type             string `json:"type,omitempty"`
}

type OfferCondition struct {
	ConditionNote string `json:"conditionNote,omitempty"`
	SubCondition  string `json:"subCondition,omitempty"`
	Value         string `json:"value,omitempty"`
}

type DealDetails struct {
	AccessType                        string `json:"accessType,omitempty"`
	Badge                             string `json:"badge,omitempty"`
	EarlyAccessDurationInMilliseconds int64  `json:"earlyAccessDurationInMilliseconds,omitempty"`
	EndTime                           string `json:"endTime,omitempty"`
	PercentClaimed                    int    `json:"percentClaimed,omitempty"`
	StartTime                         string `json:"startTime,omitempty"`
}

type LoyaltyPoints struct {
	Points int `json:"points,omitempty"`
}

type MerchantInfo struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OfferPrice struct {
	Money        *Money       `json:"money,omitempty"`
	PricePerUnit *Money       `json:"pricePerUnit,omitempty"`
	SavingBasis  *SavingBasis `json:"savingBasis,omitempty"`
	Savings      *Savings     `json:"savings,omitempty"`
}

type SavingBasis struct {
	Money                *Money `json:"money,omitempty"`
	SavingBasisType      string `json:"savingBasisType,omitempty"`
	SavingBasisTypeLabel string `json:"savingBasisTypeLabel,omitempty"`
}

type Savings struct {
	Money      *Money `json:"money,omitempty"`
	Percentage int    `json:"percentage,omitempty"`
}

type Money struct {
	Amount        float64 `json:"amount,omitempty"`
	Currency      string  `json:"currency,omitempty"`
	DisplayAmount string  `json:"displayAmount,omitempty"`
}

// --- Images Section ---

type Images struct {
	Primary  *ImageType  `json:"primary,omitempty"`
	Variants []ImageType `json:"variants,omitempty"`
}

type ImageType struct {
	Small  *ImageSize `json:"small,omitempty"`
	Medium *ImageSize `json:"medium,omitempty"`
	Large  *ImageSize `json:"large,omitempty"`
}

type ImageSize struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// --- ItemInfo Section ---

type ItemInfo struct {
	ByLineInfo      *ByLineInfo      `json:"byLineInfo,omitempty"`
	Classifications *Classifications `json:"classifications,omitempty"`
	ContentInfo     *ContentInfo     `json:"contentInfo,omitempty"`
	ContentRating   *ContentRating   `json:"contentRating,omitempty"`
	ExternalIds     *ExternalIds     `json:"externalIds,omitempty"`
	Features        *MultiTextInfo   `json:"features,omitempty"`
	ManufactureInfo *ManufactureInfo `json:"manufactureInfo,omitempty"`
	ProductInfo     *ProductInfo     `json:"productInfo,omitempty"`
	TechnicalInfo   *TechnicalInfo   `json:"technicalInfo,omitempty"`
	Title           *TextInfo        `json:"title,omitempty"`
	TradeInInfo     *TradeInInfo     `json:"tradeInInfo,omitempty"`
}

type ByLineInfo struct {
	Brand        *TextInfo     `json:"brand,omitempty"`
	Contributors []Contributor `json:"contributors,omitempty"`
	Manufacturer *TextInfo     `json:"manufacturer,omitempty"`
}

type Contributor struct {
	Locale   string `json:"locale"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	RoleType string `json:"roleType"`
}

type Classifications struct {
	Binding      *TextInfo `json:"binding,omitempty"`
	ProductGroup *TextInfo `json:"productGroup,omitempty"`
}

type ContentInfo struct {
	Edition         *TextInfo             `json:"edition,omitempty"`
	Languages       *ContentInfoLanguages `json:"languages,omitempty"`
	PagesCount      *IntInfo              `json:"pagesCount,omitempty"`
	PublicationDate *TextInfo             `json:"publicationDate,omitempty"`
}

type ContentInfoLanguages struct {
	DisplayValues []Language `json:"displayValues,omitempty"`
	Label         string     `json:"label"`
	Locale        string     `json:"locale"`
}

type Language struct {
	DisplayValue string `json:"displayValue"`
	Type         string `json:"type"`
}

type ContentRating struct {
	AudienceRating *TextInfo `json:"audienceRating,omitempty"`
}

type ExternalIds struct {
	EANs  *MultiTextInfo `json:"eans,omitempty"`
	ISBNs *MultiTextInfo `json:"isbns,omitempty"`
	UPCs  *MultiTextInfo `json:"upcs,omitempty"`
}

type ManufactureInfo struct {
	ItemPartNumber *TextInfo `json:"itemPartNumber,omitempty"`
	Model          *TextInfo `json:"model,omitempty"`
	Warranty       *TextInfo `json:"warranty,omitempty"`
}

type ProductInfo struct {
	Color          *TextInfo       `json:"color,omitempty"`
	IsAdultProduct *BoolInfo       `json:"isAdultProduct,omitempty"`
	ItemDimensions *ItemDimensions `json:"itemDimensions,omitempty"`
	ReleaseDate    *TextInfo       `json:"releaseDate,omitempty"`
	Size           *TextInfo       `json:"size,omitempty"`
	UnitCount      *IntInfo        `json:"unitCount,omitempty"`
}

type ItemDimensions struct {
	Height *Dimension `json:"height,omitempty"`
	Length *Dimension `json:"length,omitempty"`
	Weight *Dimension `json:"weight,omitempty"`
	Width  *Dimension `json:"width,omitempty"`
}

type Dimension struct {
	DisplayValue float64 `json:"displayValue"`
	Label        string  `json:"label"`
	Locale       string  `json:"locale"`
	Unit         string  `json:"unit,omitempty"`
}

type TechnicalInfo struct {
	EnergyEfficiencyClass *TextInfo      `json:"energyEfficiencyClass,omitempty"`
	Formats               *MultiTextInfo `json:"formats,omitempty"`
}

type TradeInInfo struct {
	IsEligibleForTradeIn bool   `json:"isEligibleForTradeIn,omitempty"`
	Price                *Price `json:"price,omitempty"`
}

type Price struct {
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	DisplayAmount string  `json:"displayAmount"`
}

// --- Common Types ---

type TextInfo struct {
	DisplayValue string `json:"displayValue"`
	Label        string `json:"label"`
	Locale       string `json:"locale"`
}

type MultiTextInfo struct {
	DisplayValues []string `json:"displayValues"`
	Label         string   `json:"label"`
	Locale        string   `json:"locale"`
}

type IntInfo struct {
	DisplayValue int    `json:"displayValue"`
	Label        string `json:"label"`
	Locale       string `json:"locale"`
}

type BoolInfo struct {
	DisplayValue bool   `json:"displayValue"`
	Label        string `json:"label"`
	Locale       string `json:"locale"`
}
