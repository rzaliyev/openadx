package openadx

import (
	"encoding/json"
	"log"
	"time"

	nreq "github.com/mxmCherry/openrtb/v16/native1/request"
	"github.com/mxmCherry/openrtb/v16/openrtb2"
)

type ImpID string

type Native struct {
	Request nreq.Request `json:"native"`
}

type BidRequest struct {
	openrtb2.BidRequest
	nativeImps map[ImpID]*nreq.Request
}

func NewBidRequest(data []byte) (*BidRequest, error) {
	br := BidRequest{}

	if err := json.Unmarshal(data, &br.BidRequest); err != nil {
		return nil, err
	}

	return &br, nil
}

func (r *BidRequest) Natives() map[ImpID]*nreq.Request {
	if r.nativeImps == nil {
		r.getNatives()
	}
	return r.nativeImps
}

func (r *BidRequest) getNatives() {
	r.nativeImps = make(map[ImpID]*nreq.Request)
	for _, imp := range r.Imp {
		if imp.Native != nil {
			native := Native{}
			if err := json.Unmarshal([]byte(imp.Native.Request), &native); err != nil {
				log.Fatal(err)
			}
			r.nativeImps[ImpID(imp.ID)] = &native.Request
		}
	}
}

func (r *BidRequest) IsTestMode() bool {
	return r.Test == 1
}

func (r *BidRequest) GetOperationMode() OperationMode {
	return OperationMode(r.Test)
}

func (r *BidRequest) GetAuctionType() AuctionType {
	return AuctionType(r.AT)
}

func (r *BidRequest) GetTMax() time.Duration {
	return time.Millisecond * time.Duration(r.TMax)
}

func (r *BidRequest) IsSeatAllowed(seat string) bool {
	return !StrSlice(r.BSeat).ContainsCaseSensitive(seat)
}

func (r *BidRequest) IsRoadBlockingReady() bool {
	return r.AllImps == 1
}

func (r *BidRequest) IsCurrencyAllowed(cur string) bool {
	return StrSlice(r.Cur).Contains(cur)
}

func (r *BidRequest) IsLanguageAllowed(lang string) bool {
	return StrSlice(r.WLang).Contains(lang)
}

func (r *BidRequest) IsCategoryAllowed(cat string) bool {
	return !StrSlice(r.BCat).Contains(cat)
}

func (r *BidRequest) IsAdvDomainAllowed(domain string) bool {
	return !StrSlice(r.BAdv).Contains(domain)
}
