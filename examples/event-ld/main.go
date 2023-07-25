package main

import (
	"encoding/json"
	"fmt"
	"os"

	infinite "github.com/infinite-industries/infinite-go"
	"github.com/infinite-industries/infinite-go/formats/jsonld"
)

var client *infinite.Client
var event_id string

func init() {
	client = infinite.New()
}

func main() {

	//  set event id from first command line argument
	if len(os.Args) > 1 {
		event_id = os.Args[1]
	} else {
		fmt.Fprintln(os.Stderr, "First argument must be an event ID")
		os.Exit(1)
	}

	// get event from the API
	event, err := client.Events.Get(event_id)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ldEvent := jsonld.NewEventFromApiV1(event)
	prettyPrint(ldEvent)

}

func prettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Printf("%s\n", b)
}
