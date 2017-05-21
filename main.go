package main

import (
	. "./Spiders"
	//. "github.com/PuerkitoBio/goquery"
)

func main() {
	testkeyW := []string{"clojure", "python", "golang", "lisp"}
	testLocation := []string{"Washington%2C+DC", "Chicago%2C+IL", "Boston%2C+MA", "Maryland", "Pennsylvania", "New+York+State", "Ohio", "Illinois", "Indianapolis%2C+IN"}

	//testLocation := []string{"Washington%2C+DC"}
	//testkeyW := []string{"clojure"}

	a := make(chan bool)

	go IndeedFlow(testkeyW, testLocation, a)
	go DiceFlow(testkeyW, testLocation, a)

	for i := 0; i < 2; i++ {
		<-a
	}
}
