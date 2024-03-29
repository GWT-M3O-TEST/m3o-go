package nft

import (
	"go.m3o.com/client"
)

type Nft interface {
	Asset(*AssetRequest) (*AssetResponse, error)
	Assets(*AssetsRequest) (*AssetsResponse, error)
	Collection(*CollectionRequest) (*CollectionResponse, error)
	Collections(*CollectionsRequest) (*CollectionsResponse, error)
	Create(*CreateRequest) (*CreateResponse, error)
}

func NewNftService(token string) *NftService {
	return &NftService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type NftService struct {
	client *client.Client
}

// Get a single asset by the contract
func (t *NftService) Asset(request *AssetRequest) (*AssetResponse, error) {

	rsp := &AssetResponse{}
	return rsp, t.client.Call("nft", "Asset", request, rsp)

}

// Return a list of assets
func (t *NftService) Assets(request *AssetsRequest) (*AssetsResponse, error) {

	rsp := &AssetsResponse{}
	return rsp, t.client.Call("nft", "Assets", request, rsp)

}

// Get a collection by its slug
func (t *NftService) Collection(request *CollectionRequest) (*CollectionResponse, error) {

	rsp := &CollectionResponse{}
	return rsp, t.client.Call("nft", "Collection", request, rsp)

}

// Get a list of collections
func (t *NftService) Collections(request *CollectionsRequest) (*CollectionsResponse, error) {

	rsp := &CollectionsResponse{}
	return rsp, t.client.Call("nft", "Collections", request, rsp)

}

// Create your own NFT (coming soon)
func (t *NftService) Create(request *CreateRequest) (*CreateResponse, error) {

	rsp := &CreateResponse{}
	return rsp, t.client.Call("nft", "Create", request, rsp)

}

type Asset struct {
	// last time sold
	LastSale *Sale `json:"last_sale,omitempty"`
	// name of the asset
	Name string `json:"name,omitempty"`
	// associated collection
	Collection *Collection `json:"collection,omitempty"`
	// Creator of the NFT
	Creator *User `json:"creator,omitempty"`
	// the image url
	ImageUrl string `json:"image_url,omitempty"`
	// is it a presale
	Presale bool `json:"presale,omitempty"`
	// related description
	Description string `json:"description,omitempty"`
	// the permalink
	Permalink string `json:"permalink,omitempty"`
	// number of sales
	Sales int32 `json:"sales,omitempty"`
	// the token id
	TokenId string `json:"token_id,omitempty"`
	// traits associated with the item
	Traits []map[string]interface{} `json:"traits,omitempty"`
	// asset contract
	Contract *Contract `json:"contract,omitempty"`
	// id of the asset
	Id int32 `json:"id,omitempty"`
	// listing date
	ListingDate string `json:"listing_date,omitempty"`
	// Owner of the NFT
	Owner *User `json:"owner,omitempty"`
}

type AssetRequest struct {
	ContractAddress string `json:"contract_address,omitempty"`
	TokenId         string `json:"token_id,omitempty"`
}

type AssetResponse struct {
	Asset *Asset `json:"asset,omitempty"`
}

type AssetsRequest struct {
	// order "asc" or "desc"
	Order string `json:"order,omitempty"`
	// order by "sale_date", "sale_count", "sale_price", "total_price"
	OrderBy string `json:"order_by,omitempty"`
	// limit to members of a collection by slug name (case sensitive)
	Collection string `json:"collection,omitempty"`
	// A cursor pointing to the page to retrieve
	Cursor string `json:"cursor,omitempty"`
	// limit returned assets
	Limit int32 `json:"limit,omitempty"`
	// DEPRECATED offset for pagination, please use cursor instead
	Offset int32 `json:"offset,omitempty"`
}

type AssetsResponse struct {
	// list of assets
	Assets []Asset `json:"assets,omitempty"`
	// A cursor to be supplied to retrieve the next page of results
	Next string `json:"next,omitempty"`
	// A cursor to be supplied to retrieve the previous page of results
	Previous string `json:"previous,omitempty"`
}

type Collection struct {
	// creation time
	CreatedAt string `json:"created_at,omitempty"`
	// an image for the collection
	ImageUrl string `json:"image_url,omitempty"`
	// payout address for the collection's royalties
	PayoutAddress string `json:"payout_address,omitempty"`
	// the payment tokens accepted for this collection
	PaymentTokens []Token `json:"payment_tokens,omitempty"`
	// the collection's approval status on OpenSea
	SafelistRequestStatus string `json:"safelist_request_status,omitempty"`
	// the fees that get paid out when a sale is made
	SellerFees string `json:"seller_fees,omitempty"`
	// approved editors for this collection
	Editors []string `json:"editors,omitempty"`
	// name of the collection
	Name string `json:"name,omitempty"`
	// a list of the contracts associated with this collection
	PrimaryAssetContracts []Contract `json:"primary_asset_contracts,omitempty"`
	// collection slug
	Slug string `json:"slug,omitempty"`
	// listing of all the trait types available within this collection
	Traits map[string]interface{} `json:"traits,omitempty"`
	// image used in the banner for the collection
	BannerImageUrl string `json:"banner_image_url,omitempty"`
	// description of the collection
	Description string `json:"description,omitempty"`
	// external link to the original website for the collection
	ExternalLink string `json:"external_link,omitempty"`
	// sales statistics associated with the collection
	Stats map[string]interface{} `json:"stats,omitempty"`
}

type CollectionRequest struct {
	Slug string `json:"slug,omitempty"`
}

type CollectionResponse struct {
	Collection *Collection `json:"collection,omitempty"`
}

type CollectionsRequest struct {
	Limit  int32 `json:"limit,omitempty"`
	Offset int32 `json:"offset,omitempty"`
}

type CollectionsResponse struct {
	Collections []Collection `json:"collections,omitempty"`
}

type Contract struct {
	// seller fees
	SellerFees string `json:"seller_fees,omitempty"`
	// related symbol
	Symbol string `json:"symbol,omitempty"`
	// type of contract e.g "semi-fungible"
	Type string `json:"type,omitempty"`
	// ethereum address
	Address string `json:"address,omitempty"`
	// timestamp of creation
	CreatedAt string `json:"created_at,omitempty"`
	// description of contract
	Description string `json:"description,omitempty"`
	// name of contract
	Name string `json:"name,omitempty"`
	// aka "ERC1155"
	Schema string `json:"schema,omitempty"`
	// owner id
	Owner int32 `json:"owner,omitempty"`
	// payout address
	PayoutAddress string `json:"payout_address,omitempty"`
}

type CreateRequest struct {
	// image data
	Image string `json:"image,omitempty"`
	// name of the NFT
	Name string `json:"name,omitempty"`
	// data if not image
	Data string `json:"data,omitempty"`
	// description
	Description string `json:"description,omitempty"`
}

type CreateResponse struct {
	Asset *Asset `json:"asset,omitempty"`
}

type Sale struct {
	CreatedAt      string       `json:"created_at,omitempty"`
	PaymentToken   *Token       `json:"payment_token,omitempty"`
	Quantity       string       `json:"quantity,omitempty"`
	Transaction    *Transaction `json:"transaction,omitempty"`
	AssetDecimals  int32        `json:"asset_decimals,omitempty"`
	EventTimestamp string       `json:"event_timestamp,omitempty"`
	EventType      string       `json:"event_type,omitempty"`
	TotalPrice     string       `json:"total_price,omitempty"`
	AssetTokenId   string       `json:"asset_token_id,omitempty"`
}

type Token struct {
	Address  string `json:"address,omitempty"`
	Decimals int32  `json:"decimals,omitempty"`
	EthPrice string `json:"eth_price,omitempty"`
	Id       int32  `json:"id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	Name     string `json:"name,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	UsdPrice string `json:"usd_price,omitempty"`
}

type Transaction struct {
	Id               int32  `json:"id,omitempty"`
	Timestamp        string `json:"timestamp,omitempty"`
	ToAccount        *User  `json:"to_account,omitempty"`
	TransactionHash  string `json:"transaction_hash,omitempty"`
	TransactionIndex string `json:"transaction_index,omitempty"`
	BlockHash        string `json:"block_hash,omitempty"`
	BlockNumber      string `json:"block_number,omitempty"`
	FromAccount      *User  `json:"from_account,omitempty"`
}

type User struct {
	Username   string `json:"username,omitempty"`
	Address    string `json:"address,omitempty"`
	ProfileUrl string `json:"profile_url,omitempty"`
}
