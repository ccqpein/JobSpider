package Spiders

import (
	. "fmt"
	. "github.com/PuerkitoBio/goquery"
	"log"
	//	. "strconv"
	"os/exec"
	//. "strings"
)

type TagStructures []string
type Job struct {
	Title string
	Link  string
	Date  []string
}

var cmd *exec.Cmd = exec.Command("open", "-a", "Safari")

func getSearchPages(keyWords, location []string, baseUrls string, a chan *Document) {
	for _, l := range location {
		for _, w := range keyWords {
			//Println(w, l)
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

func OpenJobPage(job Job) {
	err := exec.Command("open", job.Link).Start()
	if err != nil {
		log.Fatal(err)
	}
}

/*
func FilterJobTitle(job Job, exceptWords *[]string) bool {
	titleWords := Split(job.Title, " ")
	return false
}
*/
