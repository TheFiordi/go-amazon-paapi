# go-amazon-paapi

[![Go Reference](https://pkg.go.dev/badge/github.com/tuo-utente/go-amazon-paapi.svg)](https://pkg.go.dev/github.com/tuo-utente/go-amazon-paapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/tuo-utente/go-amazon-paapi)](https://goreportcard.com/report/github.com/tuo-utente/go-amazon-paapi)

A modern, idiomatic, and type-safe Go SDK for the **Amazon Creators API** (Product Advertising API - PA-API). 

`go-amazon-paapi` abstracts away the complexities of authentication (Login with Amazon token management), marketplace locale routing, and HTTP request handling, providing a beautiful **Fluent Interface** to build your API requests.

## Features

- ✨ **Fluent Builder Interface:** Easily chain methods to construct complex API requests.
- 🛡️ **Type-Safe Models:** Fully mapped structs for Amazon's extensive JSON responses (`ItemInfo`, `OffersV2`, `Images`, `BrowseNodeInfo`, etc.).
- 🔑 **Automated Authentication:** Built-in thread-safe LwA (Login with Amazon) token generation and caching.
- 🌍 **Global Support:** Pre-configured locales for all Amazon Marketplaces (NA, EU, FE).
- 📦 **Resource Groups:** Easily request sets of data using pre-defined Amazon Resource Groups without manually typing string constants.
- 🚀 **Zero External Dependencies:** Built entirely with the Go standard library.

## Installation

```bash
go get github.com/tuo-utente/go-amazon-paapi
```
*(Note: Replace `tuo-utente` with your actual GitHub username or module path).*

## Quick Start

### 1. Initialization

Create a new client by specifying your target Amazon Marketplace, and then inject your Amazon Partner credentials.

```go
import (
    "context"
    "fmt"
    "log"

    gap "github.com/tuo-utente/go-amazon-paapi"
    "github.com/tuo-utente/go-amazon-paapi/locale"
    "github.com/tuo-utente/go-amazon-paapi/models"
)

func main() {
    partnerTag := "your-partner-tag-20"
    clientID := "amzn1.application-oa2-client.xxxxx"
    clientSecret := "your-client-secret"

    // Initialize the client for a specific marketplace (e.g., Italy)
    client, err := gap.New(gap.WithMarketplace(locale.Italy)).
        CreateClient(partnerTag, clientID, clientSecret)
        
    if err != nil {
        log.Fatalf("Failed to initialize client: %v", err)
    }

    // Client is now ready to use!
}
```

### 2. Fetching Items (`GetItems`)

Use the fluent interface to configure your request. The SDK provides helper methods like `WithResourceGroups` to safely request blocks of data (like `ItemInfo` or `OffersV2`).

```go
// ... client initialization ...

ctx := context.Background()

resp, err := client.GetItems().
    WithItemIds([]string{"B08DYH3BJK", "B00KL8SM92"}).
    WithCondition(models.ConditionNew).
    WithResourceGroups(
        models.ResourceGroupItemInfo,
        models.ResourceGroupOffersV2,
        models.ResourceGroupImages,
    ).
    Execute(ctx)

if err != nil {
    log.Fatalf("API Error: %v", err)
}

// Safely access the results (handles Amazon's JSON inconsistencies)
if res := resp.Result(); res != nil {
    for _, item := range res.Items {
        fmt.Printf("ASIN: %s\n", item.ASIN)
        
        // Print Title
        if item.ItemInfo != nil && item.ItemInfo.Title != nil {
            fmt.Printf("Title: %s\n", item.ItemInfo.Title.DisplayValue)
        }
        
        // Print Price from OffersV2
        if item.OffersV2 != nil && len(item.OffersV2.Listings) > 0 {
            price := item.OffersV2.Listings[0].Price
            if price != nil && price.Money != nil {
                fmt.Printf("Price: %s\n", price.Money.DisplayAmount)
            }
        }
        fmt.Println("---")
    }
}
```

## Supported Operations

The SDK currently supports the following Amazon Creators API operations via the Fluent Interface:

- `client.GetItems()...`
- `client.SearchItems()...`
- `client.GetVariations()...`
- `client.GetBrowseNodes()...`

### Example: Searching Items

```go
resp, err := client.SearchItems().
    WithKeywords("Harry Potter").
    WithSearchIndex("Books").
    WithItemCount(5).
    WithResourceGroups(models.ResourceGroupImages, models.ResourceGroupItemInfo).
    Execute(ctx)
```

## Advanced Configuration

You can customize the underlying HTTP client (e.g., to add proxies, custom timeouts, or logging transports) using the Functional Options pattern during initialization:

```go
customHTTPClient := &http.Client{
    Timeout: 15 * time.Second,
}

builder := gap.New(
    gap.WithMarketplace(locale.UnitedStates),
    gap.WithHttpClient(customHTTPClient),
)

client, err := builder.CreateClient(partnerTag, clientID, clientSecret)
```

## Resource Groups vs Single Resources

Amazon APIs require you to specify exactly what data you want back via "Resources".
This SDK provides two ways to do this:

1. **`WithResourceGroups(...)`**: (Recommended) Use predefined groups that match Amazon's logical blocks.
    ```go
    .WithResourceGroups(models.ResourceGroupItemInfo, models.ResourceGroupOffersV2)
    ```

2. **`WithResources(...)`**: Request highly specific, individual fields.
    ```go
    .WithResources([]models.Resource{
        models.ResourceItemInfoTitle,
        models.ResourceImagesPrimaryLarge,
    })
    ```
*(Note: These methods append to the request, so you can safely chain them together).*

## Error Handling

The SDK exposes custom errors to help you handle specific API states, such as rate limiting:

```go
import (
    "errors"
    gap "github.com/tuo-utente/go-amazon-paapi"
)

// ...

_, err := client.GetItems().WithItemIds([]string{"B08DYH3BJK"}).Execute(ctx)
if err != nil {
    if errors.Is(err, gap.ErrRateLimitExceeded) {
        // Handle HTTP 429 Too Many Requests (e.g., apply backoff/retry logic)
        log.Println("Rate limit hit, backing off...")
    } else {
        log.Printf("Unhandled error: %v", err)
    }
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
