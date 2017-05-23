package Spiders

import (
	. "fmt"
	. "github.com/PuerkitoBio/goquery"
	//	"log"
	//"reflect"
	. "strconv"
	. "strings"
)

var indeedTagStrc TagStructures = []string{".row.result"}
var indeedBaseUrls string = "https://www.indeed.com/jobs?q=%[1]s&l=%[2]s"

// location format is Washington%2C+DC

func getIndeedTitleAndLink(s *Selection) []Job {
	var jobs []Job
	s.Each(func(_ int, s *Selection) {
		//Println(s.Find(".date").Text())
		dateStr := Split(s.Find(".date").Text(), " ")
		if len(dateStr) > 1 {
			//Println(dateStr)
			if i, err := Atoi(dateStr[0]); (i < 1 && err == nil) || dateStr[1] == "minutes" || dateStr[1] == "hours" {
				title := s.Find(".jobtitle").Text()
				if link, ok := s.Find(".jobtitle").Find("a").Attr("href"); ok {
					Println(title, dateStr)
					Println("https://www.indeed.com" + link)
					jobs = append(jobs, Job{Title: title, Link: "https://www.indeed.com" + link, Date: dateStr})
				}
			}
		}
	})
	return jobs
}

func IndeedFlow(keyWords, location []string, flag chan bool, jobs chan Job) {
	a := make(chan *Document)

	go getSearchPages(keyWords, location, indeedBaseUrls, a)

	for doc := range a {
		for _, job := range getIndeedTitleAndLink(getAllNodes(doc, indeedTagStrc)) {
			jobs <- job
		}
		//Println(&*doc.Url)
	}
	Println("finish indeed")

	flag <- true
}
