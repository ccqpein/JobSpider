package Spiders

import (
	. "fmt"
	. "github.com/PuerkitoBio/goquery"
	"log"
	//"reflect"
	. "strconv"
	. "strings"
)

var tagStrc TagStructures = []string{".row.result"}
var indeedBaseUrls string = "https://www.indeed.com/jobs?q=%[1]s&l=%[2]s"

// location format is Washington%2C+DC

func GetSearchPages(keyWords, location []string, a chan *Document) {
	for _, l := range location {
		for _, w := range keyWords {
			doc, err := NewDocument(Sprintf(indeedBaseUrls, w, l))
			if err != nil {
				log.Fatal(err)
			}
			a <- doc
		}
	}

	close(a)
}

func GetAllNodes(s *Document) *Selection {
	temp := s.Find(tagStrc[0])
	for _, v := range tagStrc[1:] {
		temp = temp.Find(v)
	}
	return temp
}

func GetTitleAndLink(s *Selection) {
	s.Each(func(_ int, s *Selection) {
		//Println(s.Find(".date").Text())
		dateStr := Split(s.Find(".date").Text(), " ")
		if len(dateStr) > 1 {
			//Println(dateStr)
			if i, err := Atoi(dateStr[0]); (i < 3 && err == nil) || dateStr[1] == "minutes" || dateStr[1] == "hours" {
				title := s.Find(".jobtitle").Text()
				if link, ok := s.Find(".jobtitle").Find("a").Attr("href"); ok {
					Println(title, dateStr)
					Println("https://www.indeed.com" + link)
				}
			}
		}
	})
}

func IndeedFlow(keyWords, location []string, a chan *Document) {
	go GetSearchPages(keyWords, location, a)

	for doc := range a {
		GetTitleAndLink(GetAllNodes(doc))
	}
}
