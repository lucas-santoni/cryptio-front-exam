package portfolio

import (
	"math/big"
	"time"
)

type Asset struct {
	// LogoImage is an optional field.
	LogoImage string `json:"logoImage,omitempty" example:"https://host.org/image.png"`
	Name      string `json:"name" example:"Bitcoin"`
	// Symbol is an optional field.
	Symbol string `json:"symbol,omitempty" example:"btc"`
}

type PortfolioTopAssetEntry struct {
	AssetInformation Asset      `json:"asset"`
	Volume           *big.Float `json:"volume" swaggertype:"number" example:"123.456"`
	UnitPrice        *big.Float `json:"unitPrice" swaggertype:"number" example:"123.456"`
	TotalPrice       *big.Float `json:"totalPrice" swaggertype:"number" example:"123.456"`
	Repartition      *big.Float `json:"repartition" swaggertype:"number" example:"51.23"`
}

type PortfolioHistoryEntry struct {
	Date      time.Time  `json:"date" example:"2021-01-01T00:00:00Z"`
	Valuation *big.Float `json:"valuation" swaggertype:"number" example:"123.456"`
}

func topAssets() ([]PortfolioTopAssetEntry, error) {
	return fakeTopAssets(), nil
}

func history() ([]PortfolioHistoryEntry, error) {
	return fakeHistory(), nil
}
