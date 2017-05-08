package main

import (
	. "fmt"
	. "github.com/PuerkitoBio/goquery"
	"log"
	//"reflect"
	. "strconv"
	. "strings"
)

var BaseUrls = []string{
	"https://www.indeed.com/jobs?q=%[1]s&l=%[2]s",
}

var KeyWords = []string{}

var Location = []string{}

type TagStructures []string

func GetSearchPages(baseUrls, keyWords, location []string, a chan *Document) {
	for _, url := range baseUrls {
		for _, l := range location {
			for _, w := range keyWords {
				doc, err := NewDocument(Sprintf(url, w, l))
				if err != nil {
					log.Fatal(err)
				}
				a <- doc
			}
		}
	}
	close(a)
}

func GetAllNodes(s *Document, tagStructures TagStructures) *Selection {
	temp := s.Find(tagStructures[0])
	for _, v := range tagStructures[1:] {
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
			if i, err := Atoi(dateStr[0]); (i <= 3 && err == nil) || dateStr[1] == "minutes" || dateStr[1] == "hours" {
				title := s.Find(".jobtitle").Text()
				if link, ok := s.Find(".jobtitle").Find("a").Attr("href"); ok {
					Println(title, dateStr)
					Println("https://www.indeed.com" + link)
				}
			}
		}
	})
}

func main() {
	testkeyW := []string{"clojure", "python", "golang", "lisp"}
	testtagStrc := []string{".row.result"}
	testLocation := []string{"Washington%2C+DC", "Chicago%2C+IL", "Boston%2C+MA", "Maryland", "Pennsylvania", "New+York+State"}

	a := make(chan *Document)

	go GetSearchPages(BaseUrls, testkeyW, testLocation, a)

	for doc := range a {
		GetTitleAndLink(GetAllNodes(doc, testtagStrc))
	}

}
