package Spiders

import (
	. "fmt"
	. "github.com/PuerkitoBio/goquery"
	"log"
	. "strconv"
	. "strings"
)

var tagStr TagStructures
var diceBaseUrls string = "https://www.dice.com/jobs/sort-date-jobs-q-%[1]s-l-%[2]s"

// location format is Chicago%2C_IL
