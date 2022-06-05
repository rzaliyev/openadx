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

type FinalDecision int

const (
	FinalDecisionExchange FinalDecision = iota
	FinalDecisionUpstreamSource
)

func (fd FinalDecision) String() string {
	switch fd {
	case FinalDecisionExchange:
		return "Final Decision: Exchange"
	case FinalDecisionUpstreamSource:
		return "Final Decision: Upstream Source"
	}
	return "Unknown"
}

type GDPRState int

const (
	GDPRStateNo GDPRState = iota
	GDPRStateYes
	GDPRStateUnknown
)

type ExtRegsGDPR struct {
	GDPR *int8 `json:"gdpr,omitempty"`
}
