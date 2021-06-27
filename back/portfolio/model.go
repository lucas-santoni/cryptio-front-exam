package portfolio

import (
	"math/big"
	"time"
)

type Asset struct {
	LogoImage string `json:"logoImage,omitempty"`
	Name      string `json:"name"`
	Symbol    string `json:"symbol,omitempty"`
}

type PortfolioTopAssetEntry struct {
	AssetInformation Asset      `json:"asset"`
	Volume           *big.Float `json:"volume"`
	UnitPrice        *big.Float `json:"unitPrice"`
	TotalPrice       *big.Float `json:"totalPrice"`
	Repartition      *big.Float `json:"repartition"`
}

type PortfolioHistoryEntry struct {
	Date      time.Time  `json:"date"`
	Valuation *big.Float `json:"valuation"`
}

func topAssets() ([]PortfolioTopAssetEntry, error) {
	return fakeTopAssets(), nil
}

func history() ([]PortfolioHistoryEntry, error) {
	return fakeHistory(), nil
}
