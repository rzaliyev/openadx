package openadx

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {

	req := loadBidRequest(t, "native_app.json")

	t.Run("bid request test mode", func(t *testing.T) {
		assertBool(t, req.IsTestMode(), true)
	})

	t.Run("bid request native objects", func(t *testing.T) {
		assertSize(t, len(req.Natives()), 2)
	})

	t.Run("bid request operation mode", func(t *testing.T) {
		got := req.GetOperationMode()
		want := TestMode
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("bid request auction type", func(t *testing.T) {
		got := req.GetAuctionType()
		want := AuctionTypeSecondPricePlus
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("bid request tmax", func(t *testing.T) {
		got := req.GetTMax()
		want := time.Millisecond * 120
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("bid request seat allowed", func(t *testing.T) {
		assertBool(t, req.IsSeatAllowed("4545"), true)
	})

	t.Run("bid request seat not allowed", func(t *testing.T) {
		assertBool(t, req.IsSeatAllowed("1234"), false)
	})

	t.Run("bid request road-blocking", func(t *testing.T) {
		assertBool(t, req.IsRoadBlockingReady(), true)
	})

	t.Run("bid request currency allowed", func(t *testing.T) {
		assertBool(t, req.IsCurrencyAllowed("kzt"), true)
	})

	t.Run("bid request currency not allowed", func(t *testing.T) {
		assertBool(t, req.IsCurrencyAllowed("EUR"), false)
	})

	t.Run("bid request language allowed", func(t *testing.T) {
		assertBool(t, req.IsLanguageAllowed("en"), true)
	})

	t.Run("bid request language not allowed", func(t *testing.T) {
		assertBool(t, req.IsLanguageAllowed("de"), false)
	})

	t.Run("bid request category allowed", func(t *testing.T) {
		assertBool(t, req.IsCategoryAllowed(string(CategoryBusiness)), true)
	})

	t.Run("bid request category not allowed", func(t *testing.T) {
		assertBool(t, req.IsCategoryAllowed(string(CategoryIllegalContent)), false)
	})

	t.Run("bid request advertiser domain allowed", func(t *testing.T) {
		assertBool(t, req.IsAdvDomainAllowed("co.squidapp"), true)
	})

	t.Run("bid request advertiser domain not allowed", func(t *testing.T) {
		assertBool(t, req.IsAdvDomainAllowed("com.example"), false)
	})

	t.Run("bid request final decision maker", func(t *testing.T) {
		got := req.GetFinalDecision()
		want := FinalDecisionUpstreamSource
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("bid request subject to COPPA", func(t *testing.T) {
		assertBool(t, req.IsSubjectToCOPPA(), true)
	})

	t.Run("bid request subject to GDPR", func(t *testing.T) {
		got := req.IsSubjectToGDPR()
		want := GDPRStateYes
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

func assertSize(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertBool(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func loadBidRequest(t testing.TB, filename string) *BidRequest {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("cannot read bid-request json: %v", err)
	}

	bidRequest, err := NewBidRequest(file)
	if err != nil {
		t.Fatalf("cannot unmarshall bid-request json: %v", err)
	}

	return bidRequest
}
