// Code generated by goa v3.0.4, DO NOT EDIT.
//
// api gRPC client CLI support package
//
// Command:
// $ goa gen calcsvc/design

package cli

import (
	auctionc "calcsvc/gen/grpc/auction/client"
	userc "calcsvc/gen/grpc/user/client"
	"flag"
	"fmt"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `auction (get-auction-product-list-by-status|get-auction-product-detail)
user (list|show|add|remove|rate|multi-add|multi-update)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` auction get--auction--product--list--by--status --message '{
      "auction_status": 2,
      "current_page": 1,
      "page_size": 30
   }'` + "\n" +
		os.Args[0] + ` user list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		auctionFlags = flag.NewFlagSet("auction", flag.ContinueOnError)

		auctionGetAuctionProductListByStatusFlags       = flag.NewFlagSet("get-auction-product-list-by-status", flag.ExitOnError)
		auctionGetAuctionProductListByStatusMessageFlag = auctionGetAuctionProductListByStatusFlags.String("message", "", "")

		auctionGetAuctionProductDetailFlags       = flag.NewFlagSet("get-auction-product-detail", flag.ExitOnError)
		auctionGetAuctionProductDetailMessageFlag = auctionGetAuctionProductDetailFlags.String("message", "", "")

		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		userShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		userShowMessageFlag = userShowFlags.String("message", "", "")
		userShowViewFlag    = userShowFlags.String("view", "", "")

		userAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		userAddMessageFlag = userAddFlags.String("message", "", "")

		userRemoveFlags       = flag.NewFlagSet("remove", flag.ExitOnError)
		userRemoveMessageFlag = userRemoveFlags.String("message", "", "")

		userRateFlags       = flag.NewFlagSet("rate", flag.ExitOnError)
		userRateMessageFlag = userRateFlags.String("message", "", "")

		userMultiAddFlags       = flag.NewFlagSet("multi-add", flag.ExitOnError)
		userMultiAddMessageFlag = userMultiAddFlags.String("message", "", "")

		userMultiUpdateFlags       = flag.NewFlagSet("multi-update", flag.ExitOnError)
		userMultiUpdateMessageFlag = userMultiUpdateFlags.String("message", "", "")
	)
	auctionFlags.Usage = auctionUsage
	auctionGetAuctionProductListByStatusFlags.Usage = auctionGetAuctionProductListByStatusUsage
	auctionGetAuctionProductDetailFlags.Usage = auctionGetAuctionProductDetailUsage

	userFlags.Usage = userUsage
	userListFlags.Usage = userListUsage
	userShowFlags.Usage = userShowUsage
	userAddFlags.Usage = userAddUsage
	userRemoveFlags.Usage = userRemoveUsage
	userRateFlags.Usage = userRateUsage
	userMultiAddFlags.Usage = userMultiAddUsage
	userMultiUpdateFlags.Usage = userMultiUpdateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "auction":
			svcf = auctionFlags
		case "user":
			svcf = userFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "auction":
			switch epn {
			case "get-auction-product-list-by-status":
				epf = auctionGetAuctionProductListByStatusFlags

			case "get-auction-product-detail":
				epf = auctionGetAuctionProductDetailFlags

			}

		case "user":
			switch epn {
			case "list":
				epf = userListFlags

			case "show":
				epf = userShowFlags

			case "add":
				epf = userAddFlags

			case "remove":
				epf = userRemoveFlags

			case "rate":
				epf = userRateFlags

			case "multi-add":
				epf = userMultiAddFlags

			case "multi-update":
				epf = userMultiUpdateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "auction":
			c := auctionc.NewClient(cc, opts...)
			switch epn {
			case "get-auction-product-list-by-status":
				endpoint = c.GetAuctionProductListByStatus()
				data, err = auctionc.BuildGetAuctionProductListByStatusPayload(*auctionGetAuctionProductListByStatusMessageFlag)
			case "get-auction-product-detail":
				endpoint = c.GetAuctionProductDetail()
				data, err = auctionc.BuildGetAuctionProductDetailPayload(*auctionGetAuctionProductDetailMessageFlag)
			}
		case "user":
			c := userc.NewClient(cc, opts...)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = userc.BuildShowPayload(*userShowMessageFlag, *userShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = userc.BuildAddPayload(*userAddMessageFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = userc.BuildRemovePayload(*userRemoveMessageFlag)
			case "rate":
				endpoint = c.Rate()
				data, err = userc.BuildRatePayload(*userRateMessageFlag)
			case "multi-add":
				endpoint = c.MultiAdd()
				data, err = userc.BuildMultiAddPayload(*userMultiAddMessageFlag)
			case "multi-update":
				endpoint = c.MultiUpdate()
				data, err = userc.BuildMultiUpdatePayload(*userMultiUpdateMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// auctionUsage displays the usage of the auction command and its subcommands.
func auctionUsage() {
	fmt.Fprintf(os.Stderr, `The auction service retrieves bottles given a set of criteria.
Usage:
    %s [globalflags] auction COMMAND [flags]

COMMAND:
    get-auction-product-list-by-status: 获取拍卖列表
    get-auction-product-detail: 拍卖详情

Additional help:
    %s auction COMMAND --help
`, os.Args[0], os.Args[0])
}
func auctionGetAuctionProductListByStatusUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] auction get-auction-product-list-by-status -message JSON

获取拍卖列表
    -message JSON: 

Example:
    `+os.Args[0]+` auction get--auction--product--list--by--status --message '{
      "auction_status": 2,
      "current_page": 1,
      "page_size": 30
   }'
`, os.Args[0])
}

func auctionGetAuctionProductDetailUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] auction get-auction-product-detail -message JSON

拍卖详情
    -message JSON: 

Example:
    `+os.Args[0]+` auction get--auction--product--detail --message '{
      "id": "Doloremque dolorem dignissimos consequatur ad officia."
   }'
`, os.Args[0])
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `The storage service makes it possible to view, add or remove wine bottles.
Usage:
    %s [globalflags] user COMMAND [flags]

COMMAND:
    list: List all stored bottles
    show: Show bottle by ID
    add: Add new bottle and return its ID.
    remove: Remove bottle from storage
    rate: Rate bottles by IDs
    multi-add: Add n number of bottles and return their IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be added.
    multi-update: Update bottles with the given IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be updated. The IDs in the query parameter is mapped to each part in the request.

Additional help:
    %s user COMMAND --help
`, os.Args[0], os.Args[0])
}
func userListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user list

List all stored bottles

Example:
    `+os.Args[0]+` user list
`, os.Args[0])
}

func userShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user show -message JSON -view STRING

Show bottle by ID
    -message JSON: 
    -view STRING: 

Example:
    `+os.Args[0]+` user show --message '{
      "id": "Dolor mollitia est vel."
   }' --view "default"
`, os.Args[0])
}

func userAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user add -message JSON

Add new bottle and return its ID.
    -message JSON: 

Example:
    `+os.Args[0]+` user add --message '{
      "composition": [
         {
            "percentage": 26,
            "varietal": "Syrah"
         },
         {
            "percentage": 26,
            "varietal": "Syrah"
         },
         {
            "percentage": 26,
            "varietal": "Syrah"
         }
      ],
      "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
      "name": "Blue\'s Cuvee",
      "rating": 1,
      "vintage": 1952,
      "winery": {
         "country": "USA",
         "name": "Longoria",
         "region": "Central Coast, California",
         "url": "http://www.longoriawine.com/"
      }
   }'
`, os.Args[0])
}

func userRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user remove -message JSON

Remove bottle from storage
    -message JSON: 

Example:
    `+os.Args[0]+` user remove --message '{
      "id": "Numquam dolor esse ullam nisi."
   }'
`, os.Args[0])
}

func userRateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user rate -message JSON

Rate bottles by IDs
    -message JSON: 

Example:
    `+os.Args[0]+` user rate --message '{
      "field": {
         "3001836966": {
            "field": [
               "Harum perferendis est saepe.",
               "Accusantium molestiae voluptatibus.",
               "Provident quibusdam aperiam."
            ]
         }
      }
   }'
`, os.Args[0])
}

func userMultiAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user multi-add -message JSON

Add n number of bottles and return their IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be added.
    -message JSON: 

Example:
    `+os.Args[0]+` user multi--add --message '{
      "field": [
         {
            "composition": [
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 1978,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 1978,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 1978,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         }
      ]
   }'
`, os.Args[0])
}

func userMultiUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user multi-update -message JSON

Update bottles with the given IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be updated. The IDs in the query parameter is mapped to each part in the request.
    -message JSON: 

Example:
    `+os.Args[0]+` user multi--update --message '{
      "bottles": [
         {
            "composition": [
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 1978,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 1978,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 26,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 1978,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         }
      ],
      "ids": [
         "Adipisci quibusdam et cupiditate voluptas.",
         "Saepe vel quas laboriosam et modi.",
         "Ducimus non expedita.",
         "Et dolores."
      ]
   }'
`, os.Args[0])
}
