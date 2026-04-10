package go_amazon_paapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-amazon-paapi/auth"
	"go-amazon-paapi/locale"
	"go-amazon-paapi/models"
	"go-amazon-paapi/operations"
	"io"
	"net/http"
)

// ErrRateLimitExceeded is returned when Amazon APIs respond with an HTTP 429 code.
// SDK users can use errors.Is(err, ErrRateLimitExceeded) to intercept it.
var ErrRateLimitExceeded = errors.New("rate limit exceeded (HTTP 429)")

const defaultBaseURL = "https://creatorsapi.amazon"

// --- 1. FUNCTIONAL OPTIONS ---

type clientConfig struct {
	locale     locale.Locale
	httpClient *http.Client
}

// Option defines the signature for functional options.
type Option func(*clientConfig)

// WithMarketplace sets the target locale for the client.
func WithMarketplace(loc locale.Locale) Option {
	return func(c *clientConfig) {
		c.locale = loc
	}
}

// WithHttpClient allows overriding the default HTTP client (e.g., for proxies or custom timeouts).
func WithHttpClient(hc *http.Client) Option {
	return func(c *clientConfig) {
		c.httpClient = hc
	}
}

// --- 2. BUILDER PATTERN ---

// ClientBuilder is the structure returned by New().
type ClientBuilder struct {
	cfg clientConfig
}

// New initializes the builder with default settings.
func New(opts ...Option) *ClientBuilder {
	// Set safe default values
	b := &ClientBuilder{
		cfg: clientConfig{
			httpClient: &http.Client{},
			locale:     locale.UnitedStates, // Default fallback
		},
	}

	// Apply the options passed to New()
	for _, opt := range opts {
		opt(&b.cfg)
	}

	return b
}

// CreateClient finalizes client creation by instantiating the authenticator and injecting credentials.
// It also allows passing extra options (like WithHttpClient) at the last moment.
// It returns an error if the initial authentication fails.
func (b *ClientBuilder) CreateClient(partnerTag, clientID, clientSecret string, extraOpts ...Option) (*Client, error) {
	for _, opt := range extraOpts {
		opt(&b.cfg)
	}

	authEndpoint := auth.DefaultLwAEndpoint
	switch b.cfg.locale.Region {
	case locale.RegionEU:
		authEndpoint = "https://api.amazon.co.uk/auth/o2/token"
	case locale.RegionFE:
		authEndpoint = "https://api.amazon.co.jp/auth/o2/token"
	}

	authCfg := auth.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   b.cfg.httpClient,
		LwAEndpoint:  authEndpoint,
	}

	authenticator := auth.NewAuthenticator(authCfg)

	// Perform initial authentication check
	if _, err := authenticator.GetToken(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to initialize client, authentication error: %w", err)
	}

	return &Client{
		authenticator: authenticator,
		httpClient:    b.cfg.httpClient,
		baseURL:       defaultBaseURL,
		targetLocale:  b.cfg.locale,
		partnerTag:    partnerTag,
	}, nil
}

// --- 3. MAIN CLIENT ---

type Client struct {
	authenticator *auth.Authenticator
	httpClient    *http.Client
	baseURL       string
	targetLocale  locale.Locale
	partnerTag    string
}

// --- 4. FLUENT OPERATIONS ---

// GetItemsOp handles the fluent interface for the GetItems operation.
type GetItemsOp struct {
	client *Client
	req    models.GetItemsRequest
}

// GetItems returns a new builder for the GetItems operation.
func (c *Client) GetItems() *GetItemsOp {
	return &GetItemsOp{
		client: c,
		req:    models.GetItemsRequest{},
	}
}

// WithItemIds sets the item identifiers (max 10).
func (op *GetItemsOp) WithItemIds(ids []string) *GetItemsOp {
	op.req.ItemIds = ids
	return op
}

// WithItemIdType sets the type of item identifier (e.g., ASIN).
func (op *GetItemsOp) WithItemIdType(t models.ItemIdType) *GetItemsOp {
	op.req.ItemIdType = t
	return op
}

// WithResources specifies the types of values to return. It appends to the existing resources.
func (op *GetItemsOp) WithResources(res []models.Resource) *GetItemsOp {
	op.req.Resources = append(op.req.Resources, res...)
	return op
}

// WithResourceGroups specifies the types of values to return using predefined groups. It appends to the existing resources.
func (op *GetItemsOp) WithResourceGroups(groups ...models.ResourceGroup) *GetItemsOp {
	op.req.Resources = append(op.req.Resources, models.GetResourcesForGroups(groups...)...)
	return op
}

// WithCondition filters offers by condition type (e.g., New, Used, Any).
func (op *GetItemsOp) WithCondition(c models.Condition) *GetItemsOp {
	op.req.Condition = c
	return op
}

// WithCurrencyOfPreference sets the currency for price information.
func (op *GetItemsOp) WithCurrencyOfPreference(curr string) *GetItemsOp {
	op.req.CurrencyOfPreference = curr
	return op
}

// WithLanguagesOfPreference sets the languages for item information.
func (op *GetItemsOp) WithLanguagesOfPreference(langs []string) *GetItemsOp {
	op.req.LanguagesOfPreference = langs
	return op
}

// Execute performs the API call.
func (op *GetItemsOp) Execute(ctx context.Context) (*models.GetItemsResponse, error) {
	return operations.GetItems(ctx, op.client, &op.req)
}

// --- SearchItems Fluent Interface ---

// SearchItemsOp handles the fluent interface for the SearchItems operation.
type SearchItemsOp struct {
	client *Client
	req    models.SearchItemsRequest
}

// SearchItems returns a new builder for the SearchItems operation.
func (c *Client) SearchItems() *SearchItemsOp {
	return &SearchItemsOp{
		client: c,
		req:    models.SearchItemsRequest{},
	}
}

func (op *SearchItemsOp) WithKeywords(k string) *SearchItemsOp {
	op.req.Keywords = k
	return op
}

func (op *SearchItemsOp) WithBrowseNodeId(id string) *SearchItemsOp {
	op.req.BrowseNodeId = id
	return op
}

func (op *SearchItemsOp) WithSearchIndex(idx string) *SearchItemsOp {
	op.req.SearchIndex = idx
	return op
}

func (op *SearchItemsOp) WithItemCount(count int) *SearchItemsOp {
	op.req.ItemCount = count
	return op
}

func (op *SearchItemsOp) WithItemPage(page int) *SearchItemsOp {
	op.req.ItemPage = page
	return op
}

func (op *SearchItemsOp) WithSortBy(sort string) *SearchItemsOp {
	op.req.SortBy = sort
	return op
}

func (op *SearchItemsOp) WithResources(res []models.Resource) *SearchItemsOp {
	op.req.Resources = append(op.req.Resources, res...)
	return op
}

// WithResourceGroups specifies the types of values to return using predefined groups.
func (op *SearchItemsOp) WithResourceGroups(groups ...models.ResourceGroup) *SearchItemsOp {
	op.req.Resources = append(op.req.Resources, models.GetResourcesForGroups(groups...)...)
	return op
}

func (op *SearchItemsOp) WithCondition(c models.Condition) *SearchItemsOp {
	op.req.Condition = c
	return op
}

func (op *SearchItemsOp) WithCurrencyOfPreference(curr string) *SearchItemsOp {
	op.req.CurrencyOfPreference = curr
	return op
}

func (op *SearchItemsOp) WithLanguagesOfPreference(langs []string) *SearchItemsOp {
	op.req.LanguagesOfPreference = langs
	return op
}

func (op *SearchItemsOp) Execute(ctx context.Context) (*models.SearchItemsResponse, error) {
	return operations.SearchItems(ctx, op.client, &op.req)
}

// --- GetVariations Fluent Interface ---

// GetVariationsOp handles the fluent interface for the GetVariations operation.
type GetVariationsOp struct {
	client *Client
	req    models.GetVariationsRequest
}

// GetVariations returns a new builder for the GetVariations operation.
func (c *Client) GetVariations() *GetVariationsOp {
	return &GetVariationsOp{
		client: c,
		req:    models.GetVariationsRequest{},
	}
}

func (op *GetVariationsOp) WithASIN(asin string) *GetVariationsOp {
	op.req.ASIN = asin
	return op
}

func (op *GetVariationsOp) WithResources(res []models.Resource) *GetVariationsOp {
	op.req.Resources = append(op.req.Resources, res...)
	return op
}

// WithResourceGroups specifies the types of values to return using predefined groups.
func (op *GetVariationsOp) WithResourceGroups(groups ...models.ResourceGroup) *GetVariationsOp {
	op.req.Resources = append(op.req.Resources, models.GetResourcesForGroups(groups...)...)
	return op
}

func (op *GetVariationsOp) WithCondition(c models.Condition) *GetVariationsOp {
	op.req.Condition = c
	return op
}

func (op *GetVariationsOp) WithCurrencyOfPreference(curr string) *GetVariationsOp {
	op.req.CurrencyOfPreference = curr
	return op
}

func (op *GetVariationsOp) WithLanguagesOfPreference(langs []string) *GetVariationsOp {
	op.req.LanguagesOfPreference = langs
	return op
}

func (op *GetVariationsOp) Execute(ctx context.Context) (*models.GetVariationsResponse, error) {
	return operations.GetVariations(ctx, op.client, &op.req)
}

// --- GetBrowseNodes Fluent Interface ---

// GetBrowseNodesOp handles the fluent interface for the GetBrowseNodes operation.
type GetBrowseNodesOp struct {
	client *Client
	req    models.GetBrowseNodesRequest
}

// GetBrowseNodes returns a new builder for the GetBrowseNodes operation.
func (c *Client) GetBrowseNodes() *GetBrowseNodesOp {
	return &GetBrowseNodesOp{
		client: c,
		req:    models.GetBrowseNodesRequest{},
	}
}

func (op *GetBrowseNodesOp) WithBrowseNodeIds(ids []string) *GetBrowseNodesOp {
	op.req.BrowseNodeIds = ids
	return op
}

func (op *GetBrowseNodesOp) WithResources(res []models.Resource) *GetBrowseNodesOp {
	op.req.Resources = append(op.req.Resources, res...)
	return op
}

// WithResourceGroups specifies the types of values to return using predefined groups.
func (op *GetBrowseNodesOp) WithResourceGroups(groups ...models.ResourceGroup) *GetBrowseNodesOp {
	op.req.Resources = append(op.req.Resources, models.GetResourcesForGroups(groups...)...)
	return op
}

func (op *GetBrowseNodesOp) WithLanguagesOfPreference(langs []string) *GetBrowseNodesOp {
	op.req.LanguagesOfPreference = langs
	return op
}

func (op *GetBrowseNodesOp) Execute(ctx context.Context) (*models.GetBrowseNodesResponse, error) {
	return operations.GetBrowseNodes(ctx, op.client, &op.req)
}

// Do is the core method called by the various operations.
// It handles request serialization, Auth token injection, HTTP request execution,
// generic/throttling error handling, and response deserialization.
func (c *Client) Do(ctx context.Context, method, path string, in interface{}, out interface{}) error {
	// --- AUTOMATIC PAYLOAD INJECTION ---
	// If the input payload supports our interface, inject the missing data
	if reqBody, ok := in.(models.APIRequest); ok {
		reqBody.SetPartnerTag(c.partnerTag)
		reqBody.SetMarketplace(c.targetLocale.Marketplace)
	}

	var buf io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		buf = bytes.NewBuffer(b)
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return fmt.Errorf("failed to create http request: %w", err)
	}

	token, err := c.authenticator.GetToken(ctx)
	if err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}

	// --- SETUP HEADERS ---
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-marketplace", c.targetLocale.Marketplace) // Required by documentation
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http client execute failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		if resp.StatusCode == http.StatusTooManyRequests {
			return ErrRateLimitExceeded
		}
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("amazon api error (status %d): %s", resp.StatusCode, string(body))
	}

	if out != nil && resp.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return fmt.Errorf("failed to decode response body: %w", err)
		}
	}

	return nil
}
