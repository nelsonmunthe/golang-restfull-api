package role

type GetResponseItem struct {
	ID             uint     `json:"id"`
	IsVerified     bool     `json:"isVerified"`
	IsReadyStock   bool     `json:"isReadStock"`
	ImageURLs      []string `json:"imageUrls"`
	UnitName       string   `json:"unitName"`
	TravelDistance string   `json:"travelDistance"`
	Transmission   string   `json:"transmission"`
	Price          string   `json:"price"`
}
