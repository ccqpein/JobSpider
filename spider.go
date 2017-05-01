package main

import (
	. "fmt"
	. "github.com/PuerkitoBio/goquery"
	"log"
	//	"strconv"
	//	. "strings"
)

var BaseUrls = []string{
	"https://www.indeed.com/jobs?q=%[1]s&l=%[2]s",
}

var KeyWords = []string{}

var Location = []string{}

type TagStructures []string

func GetSearchPages(baseUrls, keyWords, location []string) []*Document {
	var results []*Document
	for _, url := range baseUrls {
		for _, l := range location {
			for _, w := range keyWords {
				doc, err := NewDocument(Sprintf(url, w, l))
				if err != nil {
					log.Fatal(err)
				}
				results = append(results, doc)
			}
		}
	}
	return results
}

// start Concurrency
func GetSearchPages(baseUrl, keyWord, location string) *Document {
	doc, err := NewDocument(Sprintf(baseUrl, keyWord, location))
	if err != nil {
		log.Fatal(err)
	}
	return doc
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
		title := s.Find(".jobtitle").Text()
		if link, ok := s.Find(".jobtitle").Find("a").Attr("href"); ok {
			Println(title)
			Println("https://www.indeed.com" + link)
		}
	})
}

func main() {
	testkeyW := []string{"haskell", "python", "golang", "lisp"}
	testtagStrc := []string{".row.result"}
	testLocation := []string{"Washington%2C+DC", "Chicago%2C+IL", "Boston%2C+MA"}

	a := GetSearchPages(BaseUrls, testkeyW, testLocation)

	for _, doc := range a {
		GetTitleAndLink(GetAllNodes(doc, testtagStrc))
	}
}
