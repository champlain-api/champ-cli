package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/champlain-api/champ-cli/structs"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"strconv"
)

func init() {
	rootCmd.AddCommand(shuttleCmd)
	shuttleCmd.Flags().Int8Var(&updateTime, "refresh-time", 5, "specify the time (in seconds) to update the shuttles")
	shuttleCmd.Flags().BoolVar(&createOnly, "create-only", false, "specify whether to only create the shuttles")
	shuttleCmd.MarkFlagsMutuallyExclusive("create-only", "refresh-time")
}

var updateTime int8
var createOnly bool

var shuttleCmd = &cobra.Command{
	Use:   "shuttles",
	Short: "Converts shuttles to the format our API can use. Runs in a loop.",

	Run: func(cmd *cobra.Command, args []string) {

		var champlainShuttles []structs.ChamplainShuttle
		var shuttles []structs.Shuttle

		resp, err := http.Get("https://shuttle.champlain.edu/shuttledata")
		defer resp.Body.Close()

		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to connect to Champlain's shuttle API: %s\n", err))
		}
		body, _ := io.ReadAll(resp.Body)

		// Convert Champlain's champlainShuttles into ChamplainAPI objects
		jsonErr := json.Unmarshal(body, &champlainShuttles)
		if jsonErr != nil {
			log.Fatalf("Unable to unmarshal Champlain's API response: %s\n", jsonErr)
		}

		// Now go through Champlain's champlainShuttles and convert them to the champlainShuttles that we want
		for _, shuttle := range champlainShuttles {
			if shuttle.UnitID == "162498" || shuttle.UnitID == "162499" {

				lat, _ := strconv.ParseFloat(shuttle.Lat, 32)
				lon, _ := strconv.ParseFloat(shuttle.Lon, 32)
				knots, _ := strconv.ParseFloat(shuttle.Knots, 64)
				direction, _ := strconv.ParseInt(shuttle.Direction, 10, 64)
				id, _ := strconv.ParseInt(shuttle.UnitID, 10, 64)

				newShuttle := structs.Shuttle{
					ID:        int(id),
					Lat:       float32(lat),
					Lon:       float32(lon),
					MPH:       int64(float32(knots) * 1.15),
					Direction: direction,
				}
				log.Println(newShuttle)
				shuttles = append(shuttles, newShuttle)
			}

		}

		// Now send the data to our API
		var champlainAPIRequest *http.Request

		for _, shuttle := range shuttles {
			b, _ := json.Marshal(shuttle)
			champlainAPIRequest, _ = http.NewRequest(http.MethodPut,
				fmt.Sprintf("%s/shuttles/%d", APIUrl, shuttle.ID), bytes.NewBuffer(b))
			champlainAPIRequest.Header.Set("User-Agent", "champlain-api/1.0")

			// Set the authorization header for the HTTP Basic authentication.
			champlainAPIRequest.Header = http.Header{
				"Content-Type":  {"application/json"},
				"Authorization": {"Bearer api-1"},
			}

			res, err := http.DefaultClient.Do(champlainAPIRequest) // Send the request
			if err != nil {
				log.Fatalf("Unable to send request: %s\n", err)
			}
			champBody, _ := io.ReadAll(res.Body)
			log.Print(string(champBody))
		}
		champlainAPIRequest.Body.Close()

	},
}
