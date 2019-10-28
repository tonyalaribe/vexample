package vexample

import "time"

// Product struct holds all submitted form data for listings
type Product struct {
	Slug                  string   `json:"slug" bson:"slug,omitempty"`
	Title                 string   `json:"title" bson:"title,omitempty"`
	WhoMadeIt             string   `json:"who_made_it" bson:"who_made_it,omitempty"`
	IsSupply              bool     `json:"is_supply" bson:"is_supply,omitempty"`
	Description           string   `json:"description" bson:"description,omitempty"`
	WhenMade              string   `json:"when_made" bson:"when_made,omitempty"`
	Price                 int      `json:"price" bson:"price,omitempty"`
	Category              string   `json:"category" bson:"category,omitempty"`
	DuplicateCategorySlug string   `json:"duplicate_category" bson:"duplicate_category,omitempty"`
	FormFactor            string   `json:"form_factor" bson:"form_factor,omitempty"`
	InitialQuantity       int      `json:"initial_quantity" bson:"initial_quantity,omitempty"`
	Tags                  []string `json:"tags" bson:"tags,omitempty"`
	Materials             []string `json:"materials" bson:"materials,omitempty"`
	Images                []Image  `json:"images" bson:"images,omitempty"`
	ShopID                string   `json:"shop_id" bson:"shop_id,omitempty"`
	Shop                  string   `json:"shop" bson:"shop,omitempty"`
	City                  string   `json:"city" bson:"city,omitempty"`
	Country               string   `json:"country" bson:"country,omitempty"`
	CitySlug              string   `json:"city_slug" bson:"city_slug,omitempty"`
	Location              struct {
		Type        string    `json:"type" bson:"type,omitempty"`
		Coordinates []float64 `json:"coordinates" bson:"coordinates,omitempty"`
	}
	DateCreated time.Time `json:"date_created" bson:"date_created,omitempty"`
}

type ShopListingsAnalytics struct {
	Active  int `json:"active"`
	SoldOut int `json:"sold_out"`
	Total   int `json:"total"`
}

type Image struct {
	Path              string `json:"path" bson:"path,omitempty"`
	URL               string `json:"url" bson:"url,omitempty"`
	Size              int    `json:"size" bson:"size,omitempty"`
	Name              string `json:"name" bson:"name,omitempty"`
	Height            int    `json:"height" bson:height,omitempty"`
	Width             int    `json:"width" bson:"width,omitempty"`
	Ratio             int    `json:"image_ratio" bson:"image_ratio,omitempty"`
	Base64Placeholder string `json:"base64_placeholder" bson:"base64_placeholder,omitempty"`
	ProminentColor    string `json:"prominent_color" bson:"prominent_color,omitempty"`
}

type Resource struct {
	Data  []Product `json:"data"`
	Query string    `json:"query"`
}

type Tags struct {
	Slug string
	Name string `bson:"_id"`
}
