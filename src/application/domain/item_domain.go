package domain

type ItemDomain struct {
	ID                 string   `json:"id"`
	DistributionCenter []string `json:"distribution_center"`
	Processed          bool     `json:"processed"`
}
