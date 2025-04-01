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
	"time"
)

func init() {
	cmd.RootCmd.AddCommand(updateShuttlesCommand)
	updateShuttlesCommand.Flags().Int8VarP(&updateDelay, "delay", "d", 3, "specify how often (in seconds) this should update the shuttles")
}

var updateDelay int8

var updateShuttlesCommand = &cobra.Command{
	Use:   "shuttles",
	Short: "Update shuttles from Champlain's data",
	Long:  "Runs in a loop to update shuttle data taken from Champlain's API",

	Run: func(thisCmd *cobra.Command, args []string) {
		var champlainAPIRequest *http.Request

		var champlainShuttles []structs.ChamplainShuttle

		var shuttles []structs.Shuttle
		for {
			// Reset the array for the loop
			champlainShuttles = nil
			shuttles = nil

			champlainShuttleRequest, err := http.Get("https://shuttle.champlain.edu/shuttledata")

			if err != nil {
				log.Fatal(fmt.Sprintf("Failed to connect to Champlain's shuttle API: %s\n", err))
			}

			body, _ := io.ReadAll(champlainShuttleRequest.Body)
			jsonErr := json.Unmarshal(body, &champlainShuttles)
			if jsonErr != nil {
				log.Fatalf("Unable to unmarshal Champlain's API response: %s\n", jsonErr)
			}

			// Go through Champlain's shuttles and convert them to ours
			for _, oldShuttle := range champlainShuttles {
				if oldShuttle.UnitID == "162498" || oldShuttle.UnitID == "162499" {
					newShuttle := oldShuttle.ConvertShuttle(&oldShuttle)
					shuttles = append(shuttles, *newShuttle)
				}
			}

			// Now send the data to our API
			for _, shuttle := range shuttles {
				b, _ := json.Marshal(shuttle)
				if cmd.Verbose {
					log.Printf("Attempting to update shuttle with id %d", shuttle.ID)
				}
				champlainAPIRequest, _ = http.NewRequest(http.MethodPut,
					fmt.Sprintf("%s/shuttles/%d", cmd.APIUrl, shuttle.ID), bytes.NewBuffer(b))

				champlainAPIRequest.Header.Set("User-Agent", "champlain-api/1.0")

				champlainAPIRequest.Header = http.Header{
					"Content-Type":  {"application/json"},
					"Accept":        {"application/json"},
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
				case 201, 200:
					if cmd.Verbose {
						log.Printf("Updated shuttle with id %d\n", shuttle.ID)
					}
				default:
					log.Fatalf("Error %d from API:\n %s", updateShuttleResponse.StatusCode, string(updateShuttleBody))

				}
			}

			champlainAPIRequest.Body.Close()
			champlainShuttleRequest.Body.Close()

			time.Sleep(time.Duration(updateDelay) * time.Second)
		}

	},
}
