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
type FilterFunc func(*Selection, TagStructures, int) bool

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

func DuringTheseDays(s *Selection, tagStructures TagStructures, less int) Bool {
	date := s.Find(tagStructures[0]).Text()
	date = Split(date, " ")

	daysago, _ := strconv.Atoi(date)

	if daysago <= less {
		return true
	}
}

func FilterSelections(s *Selection, f FilterFunc) []*Selection {

}

func main() {
	testkeyW := []string{"haskell", ""}
	testtagStrc := []string{".row.result"}

	a := GetSearchPages(BaseUrls, testkeyW)
	GetAllNodes(a[0], testtagStrc).Each(func(i int, s *Selection) {
		Println(s.Text())
	})
}
