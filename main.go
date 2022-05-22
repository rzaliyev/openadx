package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/mxmCherry/openrtb/v16/native1/request"
	"github.com/mxmCherry/openrtb/v16/openrtb2"
)

func main() {

	file, err := ioutil.ReadFile("native_app.json")
	if err != nil {
		log.Fatal(err)
	}
	var bidRequest openrtb2.BidRequest
	err = json.Unmarshal(file, &bidRequest)
	if err != nil {
		log.Fatal(err)
	}

	rawNative := bidRequest.Imp[0].Native.Request

	native := struct {
		Native request.Request
	}{}
	err = json.Unmarshal([]byte(rawNative), &native)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Request ID: %s\n", bidRequest.ID)
	fmt.Printf("Maximum latency: %+v\n", time.Millisecond*time.Duration(bidRequest.TMax))
	fmt.Printf("AuctionType: %s\n", AuctionType(bidRequest.AT))
	fmt.Printf("Allowed currencies: %v\n", bidRequest.Cur)
	fmt.Printf("White list of buyer seats: %v\n", bidRequest.WSeat)
	fmt.Printf("Block list of buyer seats: %v\n", bidRequest.BSeat)
	fmt.Printf("Number of impressions: %d\n", len(bidRequest.Imp))
}
