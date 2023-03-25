package models

import "time"

type Message struct {
	Audio       *Audio       `json:"audio,omitempty"`
	Button      *Button      `json:"button,omitempty"`
	Context     *Context     `json:"context,omitempty"`
	Document    *Document    `json:"document,omitempty"`
	Errors      []Error      `json:"errors,omitempty"`
	From        string       `json:"from"`
	ID          string       `json:"id"`
	Identity    *Identity    `json:"identity,omitempty"`
	Image       *Image       `json:"image,omitempty"`
	Interactive *Interactive `json:"interactive,omitempty"`
	Order       *Order       `json:"order,omitempty"`
	Referral    *Referral    `json:"referral,omitempty"`
	Sticker     *Sticker     `json:"sticker,omitempty"`
	System      *System      `json:"system,omitempty"`
	Text        *Text        `json:"text,omitempty"`
	Timestamp   time.Time    `json:"timestamp"`
	Type        string       `json:"type"`
	Video       *Video       `json:"video,omitempty"`
}

type Audio struct {
	ID       string `json:"id"`
	MimeType string `json:"mime_type"`
}

type Button struct {
	Payload string `json:"payload"`
	Text    string `json:"text"`
}

type Context struct {
	Forwarded           bool             `json:"forwarded"`
	FrequentlyForwarded bool             `json:"frequently_forwarded"`
	From                string           `json:"from"`
	ID                  string           `json:"id"`
	ReferredProduct     *ReferredProduct `json:"referred_product,omitempty"`
}

type ReferredProduct struct {
	CatalogID         string `json:"catalog_id"`
	ProductRetailerID string `json:"product_retailer_id"`
}

type Document struct {
	Caption  string `json:"caption,omitempty"`
	Filename string `json:"filename"`
	Sha256   string `json:"sha256"`
	MimeType string `json:"mime_type"`
	ID       string `json:"id"`
}

type Identity struct {
	Acknowledged     bool   `json:"acknowledged"`
	CreatedTimestamp string `json:"created_timestamp"`
	Hash             string `json:"hash"`
}

type Image struct {
	Caption string `json:"caption,omitempty"`

	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
	MimeType string `json:"mime_type"`
}

type Interactive struct {
	Type *InteractiveType `json:"type,omitempty"`
}

type InteractiveType struct {
	ButtonReply *ButtonReply `json:"button_reply,omitempty"`
	ListReply   *ListReply   `json:"list_reply,omitempty"`
}

type ButtonReply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type ListReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Order struct {
	CatalogID    string        `json:"catalog_id"`
	Text         string        `json:"text"`
	ProductItems []ProductItem `json:"product_items"`
}

type ProductItem struct {
	ProductRetailerID string `json:"product_retailer_id"`
	Quantity          string `json:"quantity"`
	ItemPrice         string `json:"item_price"`
	Currency          string `json:"currency"`
}

type Referral struct {
	SourceURL    string `json:"source_url"`
	SourceType   string `json:"source_type"`
	SourceID     string `json:"source_id"`
	Headline     string `json:"headline"`
	Body         string `json:"body"`
	MediaType    string `json:"media_type"`
	ImageURL     string `json:"image_url,omitempty"`
	VideoURL     string `json:"video_url,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
}

type Sticker struct {
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
	Animated bool   `json:"animated"`
}

type System struct {
	Body     string `json:"body"`
	Identity string `json:"identity"`
	NewWaID  string `json:"new_wa_id,omitempty"`
	WaID     string `json:"wa_id,omitempty"`
	Type     string `json:"type"`
	Customer string `json:"customer"`
}

type Text struct {
	Body string `json:"body"`
}

type Video struct {
	Caption  string `json:"caption,omitempty"`
	Filename string `json:"filename"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
	MimeType string `json:"mime_type"`
}
