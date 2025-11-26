package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/location"
)

var (
	position   = []float64{-123.115, 49.285}
	INDEX_NAME = "YourPlaceIndexName" // replace with your place index name
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := location.NewFromConfig(cfg)

	resp, err := client.SearchPlaceIndexForPosition(
		context.Background(),
		&location.SearchPlaceIndexForPositionInput{
			IndexName: aws.String(INDEX_NAME),
			Position:  position,
			// Key:        new(string), // Optional
			// Language: new(string),
			//	MaxResults: new(int32), 	// Default value: 50
		},
	)
	if err != nil {
		panic("failed to search place index for position, " + err.Error())
	}

	// Output the result
	if len(resp.Results) == 0 {
		fmt.Println("No address found")
		return
	}

	place := resp.Results[0].Place

	fmt.Println("Address:", *place.Label)
	fmt.Println("Country:", *place.Country)
	fmt.Println("Region:", *place.Region)
	fmt.Println("District:", *place.SubRegion)

	// Ex value:
	// Address: Chopra yoga, 451 Granville St, Vancouver, BC, V6C 1T1, CAN
	// Country: CAN
	// Region: British Columbia
	// District: Greater Vancouver

}
