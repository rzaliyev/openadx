package openadx

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

type OperationMode int

const (
	LiveMode OperationMode = iota
	TestMode
)

func (o OperationMode) String() string {
	switch o {
	case LiveMode:
		return "Live Mode"
	case TestMode:
		return "Test Mode"
	}
	return "Unknown"
}
