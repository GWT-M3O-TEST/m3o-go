# Nft

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/nft/api](https://m3o.com/nft/api).

Endpoints:

## Create

Create your own NFT (coming soon)


[https://m3o.com/nft/api#Create](https://m3o.com/nft/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Create your own NFT (coming soon)
func CreateAnNft() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Create(&nft.CreateRequest{
		Name: "Guybrush Threepwood",
Description: "The epic monkey island character",
	})
	fmt.Println(rsp, err)
	
}
```
## Collections

Get a list of collections


[https://m3o.com/nft/api#Collections](https://m3o.com/nft/api#Collections)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Get a list of collections
func ListCollections() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Collections(&nft.CollectionsRequest{
		Limit: 1,
	})
	fmt.Println(rsp, err)
	
}
```
## Asset

Get a single asset by the contract


[https://m3o.com/nft/api#Asset](https://m3o.com/nft/api#Asset)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Get a single asset by the contract
func GetAsingleAsset() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Asset(&nft.AssetRequest{
		ContractAddress: "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb",
TokenId: "1",
	})
	fmt.Println(rsp, err)
	
}
```
## Collection

Get a collection by its slug


[https://m3o.com/nft/api#Collection](https://m3o.com/nft/api#Collection)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Get a collection by its slug
func GetAsingleCollection() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Collection(&nft.CollectionRequest{
		Slug: "doodles-official",
	})
	fmt.Println(rsp, err)
	
}
```
## Assets

Return a list of assets


[https://m3o.com/nft/api#Assets](https://m3o.com/nft/api#Assets)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Return a list of assets
func GetAlistOfAssets() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Assets(&nft.AssetsRequest{
		OrderBy: "sale_date",
Limit: 1,
	})
	fmt.Println(rsp, err)
	
}
```
