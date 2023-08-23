package main

import (
	"encoding/json"
	"fmt"
	"os"

	infinite "github.com/infinite-industries/infinite-go"
)

var client *infinite.Client
var command string
var subcommand string

func init() {
	client = infinite.New()
}

func main() {

	//  which aspect of the API is being queried - events or venues?
	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		fmt.Fprintln(os.Stderr, "First argument must be \"events\" or \"venues\"")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		subcommand = os.Args[2]
	} else {
		fmt.Fprintln(os.Stderr, "Second argument must be \"get\" or \"currentVerified\"")
		os.Exit(1)
	}

	switch command {

	case "events":
		switch subcommand {
		case "get":
			var id string
			if len(os.Args) > 3 {
				id = os.Args[3]
			} else {
				fmt.Fprintln(os.Stderr, "Third argument must be the event ID")
				os.Exit(1)
			}
			event, err := client.Events.Get(id)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			prettyPrint(event)
		case "currentVerified":
			var tags []string
			if len(os.Args) > 4 {
				tags = os.Args[4:]
			}
			events, err := client.Events.CurrentVerified(tags...)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			prettyPrint(events)
		default:
			fmt.Fprintln(os.Stderr, "Second argument must be \"get\" or \"currentVerified\"")
			os.Exit(1)
		}
	case "venues":
		switch subcommand {
		case "get":
			var id string
			if len(os.Args) > 3 {
				id = os.Args[3]
			} else {
				fmt.Fprintln(os.Stderr, "Third argument must be the venue ID")
				os.Exit(1)
			}
			venue, err := client.Venues.Get(id)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			prettyPrint(venue)
		default:
			fmt.Fprintln(os.Stderr, "Second argument must be \"get\"")
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "First argument must be \"event\" or \"venues\"")
		os.Exit(1)
	}

}

func prettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Printf("%s\n", b)
}
