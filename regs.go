package openadx

import "encoding/json"

type ExtReg struct {
	GDPR GDPRState `json:"gdpr,omitempty"`
}

func (r *BidRequest) GetExtRegs() (*ExtReg, error) {
	gdpr := &ExtReg{}
	if err := json.Unmarshal(r.Regs.Ext, gdpr); err != nil {
		return nil, err
	}
	return gdpr, nil
}

func (r *BidRequest) IsSubjectToCOPPA() bool {
	return r.Regs.COPPA == 1
}

func (r *BidRequest) IsSubjectToGDPR() (gdpr GDPRState) {
	gdpr = GDPRStateUnknown

	if r.Regs.GDPR != nil {
		gdpr = GDPRState(*r.Regs.GDPR)
	} else {
		if extGDPR, err := r.GetExtRegs(); err == nil {
			gdpr = extGDPR.GDPR
		}
	}

	switch gdpr {
	case GDPRStateNo:
		return
	case GDPRStateYes:
		return
	default:
		return GDPRStateUnknown
	}
}
