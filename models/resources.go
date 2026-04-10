package models

// Condition represents the condition of the item.
type Condition string

const (
	ConditionAny  Condition = "Any"
	ConditionNew  Condition = "New"
	ConditionUsed Condition = "Used"
)

// ItemIdType represents the identifier type.
type ItemIdType string

const (
	ItemIdTypeASIN ItemIdType = "ASIN"
)

// Resource represents a specific field to request from Amazon.
type Resource string

const (
	// BrowseNodeInfo Resources
	ResourceBrowseNodeInfoBrowseNodes          Resource = "browseNodeInfo.browseNodes"
	ResourceBrowseNodeInfoBrowseNodesAncestor  Resource = "browseNodeInfo.browseNodes.ancestor"
	ResourceBrowseNodeInfoBrowseNodesSalesRank Resource = "browseNodeInfo.browseNodes.salesRank"
	ResourceBrowseNodeInfoWebsiteSalesRank     Resource = "browseNodeInfo.websiteSalesRank"

	// Images Resources
	ResourceImagesPrimarySmall   Resource = "images.primary.small"
	ResourceImagesPrimaryMedium  Resource = "images.primary.medium"
	ResourceImagesPrimaryLarge   Resource = "images.primary.large"
	ResourceImagesVariantsSmall  Resource = "images.variants.small"
	ResourceImagesVariantsMedium Resource = "images.variants.medium"
	ResourceImagesVariantsLarge  Resource = "images.variants.large"

	// ItemInfo Resources
	ResourceItemInfoByLineInfo      Resource = "itemInfo.byLineInfo"
	ResourceItemInfoClassifications Resource = "itemInfo.classifications"
	ResourceItemInfoContentInfo     Resource = "itemInfo.contentInfo"
	ResourceItemInfoContentRating   Resource = "itemInfo.contentRating"
	ResourceItemInfoExternalIds     Resource = "itemInfo.externalIds"
	ResourceItemInfoFeatures        Resource = "itemInfo.features"
	ResourceItemInfoManufactureInfo Resource = "itemInfo.manufactureInfo"
	ResourceItemInfoProductInfo     Resource = "itemInfo.productInfo"
	ResourceItemInfoTechnicalInfo   Resource = "itemInfo.technicalInfo"
	ResourceItemInfoTitle           Resource = "itemInfo.title"
	ResourceItemInfoTradeInInfo     Resource = "itemInfo.tradeInInfo"

	// OffersV2 Resources
	ResourceOffersV2ListingsAvailability   Resource = "offersV2.listings.availability"
	ResourceOffersV2ListingsCondition      Resource = "offersV2.listings.condition"
	ResourceOffersV2ListingsDealDetails    Resource = "offersV2.listings.dealDetails"
	ResourceOffersV2ListingsIsBuyBoxWinner Resource = "offersV2.listings.isBuyBoxWinner"
	ResourceOffersV2ListingsLoyaltyPoints  Resource = "offersV2.listings.loyaltyPoints"
	ResourceOffersV2ListingsMerchantInfo   Resource = "offersV2.listings.merchantInfo"
	ResourceOffersV2ListingsPrice          Resource = "offersV2.listings.price"
	ResourceOffersV2ListingsType           Resource = "offersV2.listings.type"

	// ParentASIN Resource
	ResourceParentASIN Resource = "parentASIN"

	// SearchRefinements Resource
	ResourceSearchRefinements Resource = "searchRefinements"

	// BrowseNodes Operation Resources
	ResourceBrowseNodesAncestor Resource = "browseNodes.ancestor"
	ResourceBrowseNodesChildren Resource = "browseNodes.children"
)

// ResourceGroup represents a collection of resources for easier grouping.
type ResourceGroup int

const (
	ResourceGroupBrowseNodeInfo    ResourceGroup = 1 + iota // BrowseNodeInfo resources
	ResourceGroupImages                                     // Images resources
	ResourceGroupItemInfo                                   // ItemInfo resources
	ResourceGroupOffersV2                                   // OffersV2 resources
	ResourceGroupParentASIN                                 // ParentASIN resource
	ResourceGroupSearchRefinements                          // SearchRefinements resource
	ResourceGroupBrowseNodes                                // BrowseNodes operation resources
)

var (
	// BrowseNodeInfo resource
	ResourcesBrowseNodeInfo = []Resource{
		ResourceBrowseNodeInfoBrowseNodes,
		ResourceBrowseNodeInfoBrowseNodesAncestor,
		ResourceBrowseNodeInfoBrowseNodesSalesRank,
		ResourceBrowseNodeInfoWebsiteSalesRank,
	}

	// Images resource
	ResourcesImages = []Resource{
		ResourceImagesPrimarySmall,
		ResourceImagesPrimaryMedium,
		ResourceImagesPrimaryLarge,
		ResourceImagesVariantsSmall,
		ResourceImagesVariantsMedium,
		ResourceImagesVariantsLarge,
	}

	// ItemInfo resource
	ResourcesItemInfo = []Resource{
		ResourceItemInfoByLineInfo,
		ResourceItemInfoClassifications,
		ResourceItemInfoContentInfo,
		ResourceItemInfoContentRating,
		ResourceItemInfoExternalIds,
		ResourceItemInfoFeatures,
		ResourceItemInfoManufactureInfo,
		ResourceItemInfoProductInfo,
		ResourceItemInfoTechnicalInfo,
		ResourceItemInfoTitle,
		ResourceItemInfoTradeInInfo,
	}

	// OffersV2 resource
	ResourcesOffersV2 = []Resource{
		ResourceOffersV2ListingsAvailability,
		ResourceOffersV2ListingsCondition,
		ResourceOffersV2ListingsDealDetails,
		ResourceOffersV2ListingsIsBuyBoxWinner,
		ResourceOffersV2ListingsLoyaltyPoints,
		ResourceOffersV2ListingsMerchantInfo,
		ResourceOffersV2ListingsPrice,
		ResourceOffersV2ListingsType,
	}

	// ParentASIN resource
	ResourcesParentASIN = []Resource{
		ResourceParentASIN,
	}

	// SearchRefinements resource
	ResourcesSearchRefinements = []Resource{
		ResourceSearchRefinements,
	}

	// BrowseNodes operation resources
	ResourcesBrowseNodes = []Resource{
		ResourceBrowseNodesAncestor,
		ResourceBrowseNodesChildren,
	}

	// ResourcesMap maps groups to their respective resource slices
	ResourcesMap = map[ResourceGroup][]Resource{
		ResourceGroupBrowseNodeInfo:    ResourcesBrowseNodeInfo,
		ResourceGroupImages:            ResourcesImages,
		ResourceGroupItemInfo:          ResourcesItemInfo,
		ResourceGroupOffersV2:          ResourcesOffersV2,
		ResourceGroupParentASIN:        ResourcesParentASIN,
		ResourceGroupSearchRefinements: ResourcesSearchRefinements,
		ResourceGroupBrowseNodes:       ResourcesBrowseNodes,
	}
)

// GetResources returns the associated slice of Resources for the group.
func (g ResourceGroup) GetResources() []Resource {
	return ResourcesMap[g]
}

// GetResourcesForGroups takes multiple ResourceGroups and returns a flattened slice of all their Resources.
func GetResourcesForGroups(groups ...ResourceGroup) []Resource {
	var res []Resource
	for _, group := range groups {
		res = append(res, group.GetResources()...)
	}
	return res
}
