package shuttles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/champlain-api/champ-cli/cmd"
	"github.com/champlain-api/champ-cli/structs"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

func init() {
	shuttlesCommand.AddCommand(createShuttlesCommand)
}

var createShuttlesCommand = &cobra.Command{
	Use:   "create",
	Short: "Create shuttles",
	Run: func(thisCmd *cobra.Command, args []string) {

		var champlainAPIRequest *http.Request

		var champlainShuttles []structs.ChamplainShuttle

		var shuttles []structs.Shuttle

		champlainShuttleRequest, err := http.Get("https://shuttle.champlain.edu/shuttledata")

		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to connect to Champlain's shuttle API: %s\n", err))
		}

		body, _ := io.ReadAll(champlainShuttleRequest.Body)
		jsonErr := json.Unmarshal(body, &champlainShuttles)
		if jsonErr != nil {
			log.Fatalf("Unable to unmarshal Champlain's API response: %s\n", jsonErr)
		}

		for _, oldShuttle := range champlainShuttles {
			if oldShuttle.UnitID == "162498" || oldShuttle.UnitID == "162499" {
				newShuttle := oldShuttle.ConvertShuttle(&oldShuttle)
				shuttles = append(shuttles, *newShuttle)
			}
		}

		// Now send the data to our API
		for _, shuttle := range shuttles {
			b, _ := json.Marshal(shuttle)
			champlainAPIRequest, _ = http.NewRequest(http.MethodPost,
				fmt.Sprintf("%s/shuttles", cmd.APIUrl), bytes.NewBuffer(b))

			champlainAPIRequest.Header.Set("User-Agent", "champlain-api/1.0")

			champlainAPIRequest.Header = http.Header{
				"Content-Type":  {"application/json"},
				"Authorization": {"Bearer " + cmd.APIkey},
			}

			// Send the request and check the response code to see
			// if there were any errors.
			updateShuttleResponse, err := http.DefaultClient.Do(champlainAPIRequest)
			if err != nil {
				log.Fatalf("Unable to send request: %s\n", err)
			}
			updateShuttleBody, _ := io.ReadAll(updateShuttleResponse.Body)

			switch updateShuttleResponse.StatusCode {
			case 200, 201:
				log.Println("Created shuttle")
			case 400, 401, 500:
				log.Printf("Error %d from API: %s", updateShuttleResponse.StatusCode, string(updateShuttleBody))
			default:
				log.Printf("Got a %d from API: %s", updateShuttleResponse.StatusCode, string(updateShuttleBody))
			}

		}

		champlainAPIRequest.Body.Close()
		champlainShuttleRequest.Body.Close()

	},
}
