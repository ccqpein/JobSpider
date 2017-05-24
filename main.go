package main

import (
	. "JobSpider/Spiders"
	"os"
	//. "fmt"
)

func main() {
	testkeyW := []string{"clojure", "python", "golang", "lisp"}
	testLocation := []string{"Chicago%2C+IL", "Boston%2C+MA", "Maryland", "Pennsylvania", "New+York+State", "Ohio", "Illinois", "Indianapolis%2C+IN"}

	//testLocation := []string{"Washington%2C+DC"}
	//testkeyW := []string{"clojure", "python"}

	args := os.Args
	a := make(chan bool)
	jobs := make(chan Job)

	go IndeedFlow(testkeyW, testLocation, a, jobs)
	go DiceFlow(testkeyW, testLocation, a, jobs)

	finishP := 0
	for {
		select {
		case job := <-jobs:
			if len(args) > 1 && args[1] == "open" {
				OpenJobPage(job)
			}
			//Println(job.Link)
		case <-a:
			finishP += 1
		}
		if finishP == 2 {
			close(jobs)
			close(a)
			break
		}
	}

}
