package housing

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
	cmd.RootCmd.AddCommand(addHousingCommand)
}

var addHousingCommand = &cobra.Command{
	Use:   "housing",
	Short: "Add housing data from Champlain's scraped data",

	Run: func(thisCmd *cobra.Command, args []string) {
		var champlainAPIRequest *http.Request

		var houses []structs.House

		scrapedDataRequest, err := http.Get("https://gist.githubusercontent.com/hayden36/51ca8bd2ab3ca4cad1d3c6f4f3cb134a/raw/b08692d8defc5863d325028aeb222b8ed66c57b9/housing.json")

		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to load scraped data: %s\n", err))
		}

		body, _ := io.ReadAll(scrapedDataRequest.Body)
		jsonErr := json.Unmarshal(body, &houses)
		if jsonErr != nil {
			log.Fatalf("Unable to unmarshal scraped data: %s\n", jsonErr)
		}

		// Now send the data to our API
		for _, house := range houses {

			b, _ := json.Marshal(house)
			if cmd.Verbose {
				log.Printf("Attempting to add %s", house.Name)
			}
			champlainAPIRequest, _ = http.NewRequest(http.MethodPost,
				fmt.Sprintf("%s/housing", cmd.APIUrl), bytes.NewBuffer(b))

			champlainAPIRequest.Header.Set("User-Agent", "champlain-api/1.0")

			champlainAPIRequest.Header = http.Header{
				"Content-Type":  {"application/json"},
				"Accept":        {"application/json"},
				"Authorization": {"Bearer " + cmd.APIkey},
			}

			// Send the request and check the response code to see
			// if there were any errors.
			addHouseRequest, err := http.DefaultClient.Do(champlainAPIRequest)
			if err != nil {
				log.Fatalf("Unable to send request: %s\n", err)
			}
			addHouseBody, _ := io.ReadAll(addHouseRequest.Body)
			switch addHouseRequest.StatusCode {
			case 201, 200:
				if cmd.Verbose {
					log.Printf("Added %s.\n", house.Name)
				}
			default:
				log.Fatalf("Error %d from API:\n %s", addHouseRequest.StatusCode, string(addHouseBody))

			}
		}

		champlainAPIRequest.Body.Close()
		scrapedDataRequest.Body.Close()

	},
}
