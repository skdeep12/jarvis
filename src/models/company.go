package models

// Company has basic information about a company
type Company struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// BasicRatiosCompany has fundamental ratios for a company
type BasicRatiosCompany struct {
	PE string `json:"PE"`
	RoE string `json:"RoE"`
	EPS string `json:"EPS"`
}
