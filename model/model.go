package data

import (
	"fmt"
	"strings"
	"time"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"runtime"
)

var GlobalConfig Config
var GlobalWordList WordList

func ExecuteLongRunningTaskOnRequest() {
	
	GlobalConfig.WordListUrl
	GlobalConfig.WordListExtractorUrl = ""https://dict.tu-chemnitz.de"
	GlobalConfig.RequestExecution = true
	
	// load existing wordlist
	ReadGlobalWordlistFromRemote()
	
	for {
		//PrintMemUsage()
		time.Sleep(2 * time.Second)
		
		if GlobalConfig.RequestExecution {
			for i := 0; i < len(GlobalWordList); i++ {
				if GlobalWordList[i].Tests != nil {
					readNextWordFromLink(GlobalConfig.WordListExtractorUrl, GlobalWordList[i].Name)
				}
			}
		}
	}
}
func ReadNextWordFromLink(baseUrl string, nextWord string) {
	fmt.Println("ReadNextWordFromLink ..")
	fmt.Println("baseUrl = " + baseUrl)
	fmt.Println("nextWord = " + nextWord)
	nextWordLink = baseUrl + "/deutsch-englisch/" + nextWord + ".html"
	fmt.Println("nextWordLink = " + nextWordLink)
	
	var listOfGermanWords []string
	var listOfEnglishWords []string

	resp, err := http.Get(nextWordLink)
	check(err)
	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//fmt.Println(".. OnResponse")
	contentType := resp.Header.Get("Content-Type")
	//fmt.Println("Content-Type=", contentType)
		
	if strings.HasPrefix(contentType, "text") {
		var t string
		t = string(body)
		
		//fmt.Println(t)
			
		// from: https://godoc.org/github.com/PuerkitoBio/goquery
		// from: https://stackoverflow.com/questions/44441665/how-to-extract-only-text-from-html-in-golang
		p := strings.NewReader(t)
		doc, _ := goquery.NewDocumentFromReader(p)
		doc.Find(".adj .r").Each(func(i int, s *goquery.Selection) { 
			//fmt.Println("{adj} = ",s.Text())
			s.Find("a").Each(func(i int, s *goquery.Selection) { 
				wtext := s.Text()
				wlink, _ := s.Attr("href")
				if wtext != "" {
					//fmt.Printf("word = '%s' , link = %s\n", wtext, wlink)
					if strings.HasPrefix(wlink, "/deutsch-englisch/") {
						//fmt.Printf("german: word = '%s' , link = %s\n", wtext, wlink)
						if !contains(listOfGermanWords, wtext) {
							listOfGermanWords = append(listOfGermanWords, wtext)
						}
					} else if strings.HasPrefix(wlink, "/english-german/") {
						//fmt.Printf("english: word = '%s' , link = %s\n", wtext, wlink)
						if !contains(listOfEnglishWords, wtext) {
							listOfEnglishWords = append(listOfEnglishWords, wtext)
						}
					}
				}
			})
		})
		DoSomethingWithTheResult(listOfGermanWords, listOfEnglishWords)
	}
}
func DoSomethingWithTheResult(listOfGermanWords []string, listOfEnglishWords []string) {
	fmt.Println("DoSomethingWithTheResult ..")

	fmt.Println("listOfGermanWords ..")
	fmt.Println(listOfGermanWords)
	fmt.Println("listOfEnglishWords ..")
	fmt.Println(listOfEnglishWords)					
}
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
func ReadGlobalWordlistFromRemote() error {
	fmt.Println("ReadGlobalWordlist")
	fmt.Println("have GlobalWordlist.Words = " + strconv.Itoa(len(GlobalWordList.Words)))
	
    var err error
	var resp *http.Response
	var body []byte
	var requestUrl string = ""
	
	// would reduce the number of words
    //requestUrl = GlobalConfig.WordListUrl + "/wordlist?testOnly=true"
	// but here we need all to find out which words without test are used as well
	requestUrl = GlobalConfig.WordListUrl + "/words?format=json"
    fmt.Println("connect to wordlist and get words with tests = " + requestUrl)
    resp, err = http.Get(requestUrl)
    if err != nil {
        return err
    }

    defer resp.Body.Close()
    body, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    
    //json.Unmarshal(body, &GlobalWordList)
	json.Unmarshal(body, &GlobalWordList.Words)
    fmt.Println("got GlobalWordList.Words = " + strconv.Itoa(len(GlobalWordList.Words)))

	return nil
}
func PrintMemUsage() {
	// from: https://golangcode.com/print-the-current-memory-usage/
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    // for info on each, see: https://golang.org/pkg/runtime/#MemStats
    fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
    fmt.Printf(", TotalAlloc = %v MiB", bToMb(m.TotalAlloc))
    fmt.Printf(", Sys = %v MiB", bToMb(m.Sys))
    fmt.Printf(", NumGC = %v\n", m.NumGC)
}
func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}