// Code generated by goa v3.0.4, DO NOT EDIT.
//
// api HTTP client CLI support package
//
// Command:
// $ goa gen calcsvc/design

package cli

import (
	auctionc "calcsvc/gen/http/auction/client"
	userc "calcsvc/gen/http/user/client"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `auction (pick|get)
user (list|show|add|remove|rate|multi-add|multi-update)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` auction pick --body '{
      "name": "Blue\'s Cuvee",
      "varietal": [
         "pinot noir",
         "merlot",
         "cabernet franc"
      ],
      "winery": "longoria"
   }'` + "\n" +
		os.Args[0] + ` user list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	userMultiAddEncoderFn userc.UserMultiAddEncoderFunc,
	userMultiUpdateEncoderFn userc.UserMultiUpdateEncoderFunc,
) (goa.Endpoint, interface{}, error) {
	var (
		auctionFlags = flag.NewFlagSet("auction", flag.ContinueOnError)

		auctionPickFlags    = flag.NewFlagSet("pick", flag.ExitOnError)
		auctionPickBodyFlag = auctionPickFlags.String("body", "REQUIRED", "")

		auctionGetFlags = flag.NewFlagSet("get", flag.ExitOnError)

		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		userShowFlags    = flag.NewFlagSet("show", flag.ExitOnError)
		userShowIDFlag   = userShowFlags.String("id", "REQUIRED", "ID of bottle to show")
		userShowViewFlag = userShowFlags.String("view", "", "")

		userAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		userAddBodyFlag = userAddFlags.String("body", "REQUIRED", "")

		userRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		userRemoveIDFlag = userRemoveFlags.String("id", "REQUIRED", "ID of bottle to remove")

		userRateFlags = flag.NewFlagSet("rate", flag.ExitOnError)
		userRatePFlag = userRateFlags.String("p", "REQUIRED", "map[uint32][]string is the payload type of the user service rate method.")

		userMultiAddFlags    = flag.NewFlagSet("multi-add", flag.ExitOnError)
		userMultiAddBodyFlag = userMultiAddFlags.String("body", "REQUIRED", "")

		userMultiUpdateFlags    = flag.NewFlagSet("multi-update", flag.ExitOnError)
		userMultiUpdateBodyFlag = userMultiUpdateFlags.String("body", "REQUIRED", "")
		userMultiUpdateIdsFlag  = userMultiUpdateFlags.String("ids", "REQUIRED", "")
	)
	auctionFlags.Usage = auctionUsage
	auctionPickFlags.Usage = auctionPickUsage
	auctionGetFlags.Usage = auctionGetUsage

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
			case "pick":
				epf = auctionPickFlags

			case "get":
				epf = auctionGetFlags

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
			c := auctionc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "pick":
				endpoint = c.Pick()
				data, err = auctionc.BuildPickPayload(*auctionPickBodyFlag)
			case "get":
				endpoint = c.Get()
				data = nil
			}
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = userc.BuildShowPayload(*userShowIDFlag, *userShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = userc.BuildAddPayload(*userAddBodyFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = userc.BuildRemovePayload(*userRemoveIDFlag)
			case "rate":
				endpoint = c.Rate()
				var err error
				var val map[uint32][]string
				err = json.Unmarshal([]byte(*userRatePFlag), &val)
				data = val
				if err != nil {
					return nil, nil, fmt.Errorf("invalid JSON for userRatePFlag, example of valid JSON:\n%s", "'{\n      \"\\u003cuint32 Value\\u003e\": [\n         \"Laboriosam consequatur delectus doloribus.\",\n         \"Est mollitia.\",\n         \"Voluptas ex enim.\",\n         \"Est explicabo eveniet dolore.\"\n      ]\n   }'")
				}
			case "multi-add":
				endpoint = c.MultiAdd(userMultiAddEncoderFn)
				data, err = userc.BuildMultiAddPayload(*userMultiAddBodyFlag)
			case "multi-update":
				endpoint = c.MultiUpdate(userMultiUpdateEncoderFn)
				data, err = userc.BuildMultiUpdatePayload(*userMultiUpdateBodyFlag, *userMultiUpdateIdsFlag)
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
    pick: Pick implements pick.
    get: get

Additional help:
    %s auction COMMAND --help
`, os.Args[0], os.Args[0])
}
func auctionPickUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] auction pick -body JSON

Pick implements pick.
    -body JSON: 

Example:
    `+os.Args[0]+` auction pick --body '{
      "name": "Blue\'s Cuvee",
      "varietal": [
         "pinot noir",
         "merlot",
         "cabernet franc"
      ],
      "winery": "longoria"
   }'
`, os.Args[0])
}

func auctionGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] auction get

get

Example:
    `+os.Args[0]+` auction get
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
	fmt.Fprintf(os.Stderr, `%s [flags] user show -id STRING -view STRING

Show bottle by ID
    -id STRING: ID of bottle to show
    -view STRING: 

Example:
    `+os.Args[0]+` user show --id "Soluta consequatur nostrum dolore et." --view "default"
`, os.Args[0])
}

func userAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user add -body JSON

Add new bottle and return its ID.
    -body JSON: 

Example:
    `+os.Args[0]+` user add --body '{
      "composition": [
         {
            "percentage": 71,
            "varietal": "Syrah"
         },
         {
            "percentage": 71,
            "varietal": "Syrah"
         },
         {
            "percentage": 71,
            "varietal": "Syrah"
         },
         {
            "percentage": 71,
            "varietal": "Syrah"
         }
      ],
      "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
      "name": "Blue\'s Cuvee",
      "rating": 1,
      "vintage": 1974,
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
	fmt.Fprintf(os.Stderr, `%s [flags] user remove -id STRING

Remove bottle from storage
    -id STRING: ID of bottle to remove

Example:
    `+os.Args[0]+` user remove --id "Animi enim nam nemo."
`, os.Args[0])
}

func userRateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user rate -p JSON

Rate bottles by IDs
    -p JSON: map[uint32][]string is the payload type of the user service rate method.

Example:
    `+os.Args[0]+` user rate --p '{
      "\u003cuint32 Value\u003e": [
         "Laboriosam consequatur delectus doloribus.",
         "Est mollitia.",
         "Voluptas ex enim.",
         "Est explicabo eveniet dolore."
      ]
   }'
`, os.Args[0])
}

func userMultiAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user multi-add -body JSON

Add n number of bottles and return their IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be added.
    -body JSON: 

Example:
    `+os.Args[0]+` user multi--add --body '[
      {
         "composition": [
            {
               "percentage": 71,
               "varietal": "Syrah"
            },
            {
               "percentage": 71,
               "varietal": "Syrah"
            }
         ],
         "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
         "name": "Blue\'s Cuvee",
         "rating": 2,
         "vintage": 1937,
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
               "percentage": 71,
               "varietal": "Syrah"
            },
            {
               "percentage": 71,
               "varietal": "Syrah"
            }
         ],
         "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
         "name": "Blue\'s Cuvee",
         "rating": 2,
         "vintage": 1937,
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
               "percentage": 71,
               "varietal": "Syrah"
            },
            {
               "percentage": 71,
               "varietal": "Syrah"
            }
         ],
         "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
         "name": "Blue\'s Cuvee",
         "rating": 2,
         "vintage": 1937,
         "winery": {
            "country": "USA",
            "name": "Longoria",
            "region": "Central Coast, California",
            "url": "http://www.longoriawine.com/"
         }
      }
   ]'
`, os.Args[0])
}

func userMultiUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user multi-update -body JSON -ids JSON

Update bottles with the given IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be updated. The IDs in the query parameter is mapped to each part in the request.
    -body JSON: 
    -ids JSON: 

Example:
    `+os.Args[0]+` user multi--update --body '{
      "bottles": [
         {
            "composition": [
               {
                  "percentage": 71,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 71,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 1937,
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
                  "percentage": 71,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 71,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 1937,
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
                  "percentage": 71,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 71,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 2,
            "vintage": 1937,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         }
      ]
   }' --ids '[
      "Aut aut.",
      "Fugiat aliquid facere doloremque placeat ea vel.",
      "Incidunt earum dolorem suscipit ut iusto."
   ]'
`, os.Args[0])
}
