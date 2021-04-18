package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/jasongauvin/zebi-scraper/api/models"
)

func GetFact (c *gin.Context) {
	category := c.Param("name")
	fmt.Println(category)
	
	allFacts := make([]models.Fact, 0)

	// create a new collector
	collector := colly.NewCollector(
		// Visit only domains
        colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
    )

	// On every a element which has href attribute call callback
    collector.OnHTML(".factsList li", func(e *colly.HTMLElement) {
        factId, err := strconv.Atoi(e.Attr("id"))
        if err != nil {
            log.Error("Could not get id")
        }
		// Print id
		fmt.Println("id : ", factId)

        factDesc := e.Text
        fact := models.Fact{
            ID:          factId,
            Description: factDesc,
        }

        allFacts = append(allFacts, fact)
		
    })

	// Before making a request print "Visiting ..."
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	// Start scraping
	if err := collector.Visit("https://www.factretriever.com/" + category + "-facts");
	err != nil {
		log.Error("Visit error ", err)
	}
	if category == " " {
		fmt.Println("Category is empty")
	}

	// writeJSON(allFacts)

}

func writeJSON(data []models.Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("rhinofacts.json", file, 0644)
}