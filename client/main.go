package main

import (
	"context"
	"fmt"
    pb "github.com/silverliningco/ahri_grpc/proto"
	"log"
    "io"
    "google.golang.org/protobuf/encoding/protojson"//"encoding/json"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client is running")

	opts := grpc.WithInsecure() // no SSL certificates

	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}
	// @doc: The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding
	// function returns.
	defer cc.Close() //

	c := pb.NewSearchServiceClient(cc)


    // Hardcoded json from Client
    myjsonString := []byte(`{"id":0,"location":{"latitude":42.123456,"longitude":-75.456789,"elevation":804},"outdoorDesignConditions":{"weatherStation":"Raleigh-Durham IAP","state":"NC","elevation":436,"latitude":36,"heating99DB":23,"cooling01DB":92,"coincidentWB":76,"DG45RH":51,"DG50RH":44,"DG55RH":37,"dailyRange":"Medium"},"indoorDesignConditions":{"winterIndoorF":70,"summerIndoorF":75,"coolingRH":50},"loadCalculation":{"sensibleBTUH":16043.181,"heatingBTUH":26342.136,"latentBTUH":2334.585},"systemAttributes":{"heatedCooled":{"providesCooling":true,"providesHeating":true},"fuelSource":"Natural Gas","energyDistributionMethod":"Forced air"}}`);

	search := &pb.Search{}
    err = protojson.Unmarshal(myjsonString, search)
    if err != nil {
		log.Fatalf("unmarshal %v", err)
	}

    //fmt.Println("=== Request ===")
    //fmt.Println(search)

	// Loading equipment search
	myResultStream, err := c.EquipmentSearch(context.Background(), search)
   
	if err != nil {
		fmt.Printf("Error happened while getting: %v \n", err)
	}

    fmt.Println("=== Response ===")
    for {
		res, err := myResultStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving results: %v", err)
		}

       // Encode proto.Message in JSON myResults
        b, err := protojson.Marshal(res)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(string(b))
	}

    
  
}
