package main

import (
	. "./Spiders"
	. "github.com/PuerkitoBio/goquery"
)

func main() {
	testkeyW := []string{"clojure", "python", "golang", "lisp"}
	//	testLocation := []string{"Washington%2C+DC", "Chicago%2C+IL", "Boston%2C+MA", "Maryland", "Pennsylvania", "New+York+State", "Ohio", "Illinois", "Indianapolis%2C+IN"}

	testLocation2 := []string{"Chicago%2C_IL"}
	a := make(chan *Document)

	IndeedFlow(testkeyW, testLocation, a)
	DiceFlow(testkeyW, testLocation2, a)
}
