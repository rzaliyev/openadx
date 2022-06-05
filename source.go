package openadx

func (r *BidRequest) GetFinalDecision() FinalDecision {
	return FinalDecision(r.Source.FD)
}
