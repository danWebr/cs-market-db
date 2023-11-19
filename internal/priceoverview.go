package data

type PriceOverview struct {
	Success     bool   `json:"success"`
	LowestPrice string `json:"lowest_price,omitempty"`
	Volume      string `json:"volume,omitempty"`
	MedianPrice string `json:"median_price,omitempty"`
}
