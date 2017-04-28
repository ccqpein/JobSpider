package main

import (
	. "fmt"
	. "github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	. "strings"
)

var BaseUrls = []string{
	"https://www.indeed.com/jobs?q=%[1]s&l=%[2]s",
}

type TagStructures []string

func GetSearchPages(baseUrls, keyWords []string) []*Document {
	var results []*Document
	for _, url := range baseUrls {
		doc, err := NewDocument(Sprintf(url, keyWords[0], keyWords[1]))
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, doc)
	}
	return results
}

func GetAllNodes(s *Document, tagStructures TagStructures) *Selection {
	var temp *Selection
	for _, v := range tagStructures {
		temp = s.Find(v)
	}
	return temp
}

/*
func FilterSelectionsDate(s *Selection, days int) {

	s.FilterFunction(func(_ int, s *Selection) bool {
		if day, _ := strconv.Atoi(Split(s.Find(".date").Text(), " ")[0]); day <= days {
			return true
		} else {
			return false
		}
	}).Each(func(_ int, s *Selection) {
		title := s.Find(".jobtitle").Text()
		link, ok := s.Find(".jobtitle").Find("a").Attr("href")
		Println(title)
		Println(ok, link)
	})
}*/

func GetTitleAndLink(s *Selection) {

	s.Each(func(_ int, s *Selection) {
		title := s.Find(".jobtitle").Text()
		link, ok := s.Find(".jobtitle").Find("a").Attr("href")
		Println(title)
		Println(ok, link)
	})
}

func main() {
	testkeyW := []string{"haskell", ""}
	testtagStrc := []string{".row.result"}

	a := GetSearchPages(BaseUrls, testkeyW)

	Println("haha")

	FilterSelectionsDate(GetAllNodes(a[0], testtagStrc), 9)
}
