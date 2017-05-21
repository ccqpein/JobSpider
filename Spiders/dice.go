package Spiders

import (
	. "fmt"
	. "github.com/PuerkitoBio/goquery"
	//	"log"
	. "strconv"
	. "strings"
)

var diceTagStrc TagStructures = []string{".complete-serp-result-div"}
var diceBaseUrls string = "https://www.dice.com/jobs/sort-date-jobs-q-%[1]s-l-%[2]s-jobs"

// location format is Chicago%2C_IL. Chicago%2C+IL is fine too

func getDiceTitleAndLink(s *Selection) {
	s.Each(func(_ int, s *Selection) {
		//Println(s.Find(".posted").Text())
		dateStr := Split(s.Find(".posted").Text(), " ")
		if len(dateStr) > 1 {
			//Println(dateStr)
			if i, err := Atoi(dateStr[0]); (i < 3 && err == nil) || dateStr[1] == "minutes" || dateStr[1] == "hours" {
				title, _ := s.Find(".list-inline").Find("a").Attr("title")
				if link, ok := s.Find(".list-inline").Find("a").Attr("href"); ok {
					Println()
					Println(title, dateStr)
					Println(link)
				}
			}
		}
	})
}

func DiceFlow(keyWords, location []string, flag chan bool) {
	a := make(chan *Document)
	go getSearchPages(keyWords, location, diceBaseUrls, a)

	for doc := range a {
		getDiceTitleAndLink(getAllNodes(doc, diceTagStrc))
		//Println(&*doc.Url)
	}
	//Println("finish Dice")
	flag <- true
}
