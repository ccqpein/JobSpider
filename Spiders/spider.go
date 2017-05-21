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
			//Println(Sprintf(baseUrls, w, l))
			doc, err := NewDocument(Sprintf(baseUrls, w, l))
			if err != nil {
				log.Fatal(err)
			}
			//Println("\n ************* \n", Sprintf(baseUrls, w, l), "\n ***************")
			a <- doc
		}
	}

	close(a)
}

func getAllNodes(s *Document, tagStrc TagStructures) *Selection {
	temp := s.Find(tagStrc[0])
	for _, v := range tagStrc[1:] {
		temp = temp.Find(v)
	}
	return temp
}
