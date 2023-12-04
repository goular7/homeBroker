package entity

type Inverstor struct {
	ID string
	Name string
	AssetPosition []*InverstorAssetPosition
}

func NewInvestor(id string) *Inverstor {
	return &Inverstor{
		ID:            id,
		AssetPosition: []*InverstorAssetPosition{},
	}
}

func(i *Inverstor) AddAssetPositions(assetPosition *InverstorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

type InverstorAssetPosition struct {
	AssetID string
	Shares int
}