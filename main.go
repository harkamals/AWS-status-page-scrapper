package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strings"
)

type Services struct {
	Name    string
	Status  string
	Country string
}

func main() {

	locations := map[string]string{"Singapore": "AP", "Ireland": "EU", "London": "EU"}
	services := make([]Services, 0, 300)

	for country, region := range locations {
		c := colly.NewCollector()

		c.OnHTML(fmt.Sprintf("div[id=%s_block] table tbody tr .pad8", region), func(e *colly.HTMLElement) {

			if strings.Contains(e.Text, country) {

				replacer := strings.NewReplacer(" Amazon", "", " AWS", "", " ("+country+")", "")
				name := replacer.Replace(e.Text)

				this := Services{
					Name:    name,
					Country: country,
					Status:  e.DOM.Next().Text(),
				}
				services = append(services, this)
			}

		})
		c.Visit("https://status.aws.amazon.com")
	}

	result, _ := json.MarshalIndent(services, "", "\t")
	os.Stdout.Write(result)

}
