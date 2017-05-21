package Spiders

import (
	. "fmt"
	. "github.com/PuerkitoBio/goquery"
	"log"
	//	. "strconv"
	//	. "strings"
)

type TagStructures []string

func getSearchPages(keyWords, location []string, baseUrls string, a chan *Document) {
	for _, l := range location {
		for _, w := range keyWords {
			doc, err := NewDocument(Sprintf(baseUrls, w, l))
			if err != nil {
				log.Fatal(err)
			}
			//Println(Sprintf(baseUrls, w, l))
			a <- doc
		}
	}

	close(a)
}

func getAllNodes(s *Document) *Selection {
	temp := s.Find(indeedTagStrc[0])
	for _, v := range indeedTagStrc[1:] {
		temp = temp.Find(v)
	}
	return temp
}
