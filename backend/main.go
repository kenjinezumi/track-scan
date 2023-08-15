package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"

)

func main() {
	fmt.Println("Hi")

	//     r := gin.Default()

	//     // Define your API routes here
	//     r.GET("/api/data", func(c *gin.Context) {
	//         // Process data or send a response
	//         c.JSON(200, gin.H{"message": "Hello from the backend!"})
	//     })

	// r.Run(":8080") // Run the server on port 8080
	resp, err := http.Get("https://mvfglobal.com/")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	// Call the MetaTagsAnalyser function
	metaTagsResult := services.MetaTagsAnalyser(doc)
	fmt.Println(metaTagsResult)

	cookiesResponse := services.AnalyzeCookies(resp)
	fmt.Println(cookiesResponse)

	knownNetworks, err := utils.LoadDomains("list_domains.txt") // Add your known ad or tracking networks
	if err != nil {
		log.Fatal("Error loading the domains:", err)
	}
	externalResResult := services.ExternalResourcesAnalyser(doc, knownNetworks...)
	fmt.Println(externalResResult)

	cspReport := services.NewCSPInspector().InspectCSP(resp)
	fmt.Println(cspReport)
	for directive, values := range cspReport.Directives {
		fmt.Println(directive, ":", values)
	}

}
