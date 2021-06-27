package portfolio

import (
	"back/utils"
	"math/big"
	"math/rand"
	"sort"
	"time"
)

var btc = Asset{
	"https://assets.coingecko.com/coins/images/1/small/bitcoin.png",
	"Bitcoin",
	"BTC",
}

var ltc = Asset{
	"https://assets.coingecko.com/coins/images/2/small/litecoin.png",
	"Litecoin",
	"LTC",
}

var eth = Asset{
	"https://assets.coingecko.com/coins/images/279/small/ethereum.png",
	"Ethereum",
	"ETH",
}

var uni = Asset{
	"https://assets.coingecko.com/coins/images/12504/small/uniswap-uni.png",
	"Uniswap",
	"UNI",
}

var aave = Asset{
	"https://assets.coingecko.com/coins/images/12645/small/AAVE.png",
	"Aave",
	"AAVE",
}

var curve = Asset{
	"https://assets.coingecko.com/coins/images/12124/small/Curve.png",
	"Curve DAO Token",
	"CRV",
}

var thor = Asset{
	"https://assets.coingecko.com/coins/images/6595/small/RUNE.png",
	"THORChain",
	"RUNE",
}

var dot = Asset{
	"https://assets.coingecko.com/coins/images/12171/small/aJGBjJFU_400x400.jpg",
	"Polkadot",
	"DOT",
}

var alice = Asset{
	"https://assets.coingecko.com/coins/images/14375/small/alice_logo.jpg",
	"My Neighbor Alice",
	"ALICE",
}

var other = Asset{
	"",
	"Other",
	"",
}

func parseBigFloatOrPanic(s string) *big.Float {
	f := new(big.Float)
	f, _, err := f.Parse(s, 10)
	utils.OrPanic(err)

	return f
}

func fakeTopAssets() []PortfolioTopAssetEntry {
	entries := []PortfolioTopAssetEntry{
		{btc, parseBigFloatOrPanic("138.4123456"), parseBigFloatOrPanic("13.1"), nil, nil},
		{ltc, parseBigFloatOrPanic("1.85024"), parseBigFloatOrPanic("0.000123"), nil, nil},
		{eth, parseBigFloatOrPanic("23.1285"), parseBigFloatOrPanic("1230004"), nil, nil},
		{uni, parseBigFloatOrPanic("0.00134"), parseBigFloatOrPanic("4.4442"), nil, nil},
		{aave, parseBigFloatOrPanic("123552.1234"), parseBigFloatOrPanic("145.0853245123"), nil, nil},
		{curve, parseBigFloatOrPanic("75923.00004"), parseBigFloatOrPanic("1.12399999"), nil, nil},
		{thor, parseBigFloatOrPanic("123.52348"), parseBigFloatOrPanic("1.425629"), nil, nil},
		{dot, parseBigFloatOrPanic("1.1114"), parseBigFloatOrPanic("0.00000001"), nil, nil},
		{alice, parseBigFloatOrPanic("45.2123"), parseBigFloatOrPanic("111.11"), nil, nil},
		{other, parseBigFloatOrPanic("120.1"), parseBigFloatOrPanic("0.00234512358899"), nil, nil},
	}

	total := big.NewFloat(0)
	for i := range entries {
		entries[i].TotalPrice = big.NewFloat(0).Mul(entries[i].Volume, entries[i].UnitPrice)
		total = total.Add(total, entries[i].TotalPrice)
	}

	// Yeah let's range twice...
	for i := range entries {
		ratio := big.NewFloat(0).Quo(entries[i].TotalPrice, total)
		entries[i].Repartition = big.NewFloat(0).Mul(ratio, big.NewFloat(100))
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Repartition.Cmp(entries[j].Repartition) > 0
	})

	return entries
}

func fakeHistory() []PortfolioHistoryEntry {
	entries := make([]PortfolioHistoryEntry, 60)

	startingDate, err := time.Parse("2006/02/01 15:04:05", "2021/01/01 00:00:00")
	utils.OrPanic(err)

	rand.Seed(42)
	min, max := 100., 200.

	for i := range entries {
		entries[i].Date = startingDate
		entries[i].Valuation = big.NewFloat(min + rand.Float64()*(max-min))
		startingDate = startingDate.Add(24 * time.Hour)
	}

	return entries
}
