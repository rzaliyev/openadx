package main

type AuctionType int

const (
	AuctionTypeFirstPrice AuctionType = iota + 1
	AuctionTypeSecondPricePlus
)

func (a AuctionType) String() string {
	switch a {
	case AuctionTypeFirstPrice:
		return "First Price"
	case AuctionTypeSecondPricePlus:
		return "Second Price Plus"
	}
	return "Unknown"
}
