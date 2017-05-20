package main

import (
	. "./Spiders"
	. "github.com/PuerkitoBio/goquery"
)

func main() {
	testkeyW := []string{"clojure", "python", "golang", "lisp"}
	testLocation := []string{"Washington%2C+DC", "Chicago%2C+IL", "Boston%2C+MA", "Maryland", "Pennsylvania", "New+York+State"}

	a := make(chan *Document)

	IndeedFlow(testkeyW, testLocation, a)
}
