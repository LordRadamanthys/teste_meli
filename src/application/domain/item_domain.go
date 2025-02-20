package domain

type ItemDomain struct {
	ID                        string
	PrimaryDistributionCenter string
	DistributionCenter        []string
	Processed                 bool
}
