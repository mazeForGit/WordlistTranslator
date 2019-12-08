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
func ExecuteLongRunningTaskOnRequest() {
	
	for {
		//PrintMemUsage()
		time.Sleep(2 * time.Second)
		if GlobalConfig.RequestExecution {
			
			readNextWordFromLink()
			
			if (GlobalConfig.WordToStartWithNext == GlobalConfig.WordToStartWith) {
				break
			}
		}
	}
}
func readNextWordFromLink() {

	var s Status
	var nextWordLink string = ""
	var err error
	var client *http.Client
	var resp *http.Response
	var req *http.Request
	var body, payload []byte
	var contentType string
	var word string = ""
	var wordtype string = ""
	var wordbreak string = ""
	var nextWord string = ""
	var requestUrl string = ""
	var wrd Word
	var t string
	var p *strings.Reader	
	var doc *goquery.Document
	
			nextWordLink = GlobalConfig.WordListExtractorUrl + "/" + GlobalConfig.WordToStartWithNext
			//fmt.Println("ReadNextWordFromLink = " + nextWordLink)

			resp, err = http.Get(nextWordLink)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			
			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			
			//fmt.Println(".. OnResponse")
			contentType = resp.Header.Get("Content-Type")
			//fmt.Println("Content-Type=", contentType)
				
			if strings.HasPrefix(contentType, "text") {
					
				t = string(body)
				
				//fmt.Println(t)
				// from: https://stackoverflow.com/questions/44441665/how-to-extract-only-text-from-html-in-golang
				p = strings.NewReader(t)
				doc, _ = goquery.NewDocumentFromReader(p)

				word = ""
				wordtype = ""
				wordbreak = ""
				nextWord = ""
							
				// find the items
				doc.Find(".dwdswb-ft .dwdswb-ft-lemmaansatz").Each(func(i int, s *goquery.Selection) {
					// for each item found
							
					word = s.Text()
					word = strings.TrimSpace(word)
					//fmt.Println("word = ",word)
				})
						
				doc.Find(".dwdswb-ft .dwdswb-ft-blocks .dwdswb-ft-block").Each(func(i int, s *goquery.Selection) { 
					//fmt.Println("s = ",s.Text())
							
					grammar := s.Find(".dwdswb-ft-blocklabel")
					//fmt.Println("grammar = ",grammar.Text())
					if (grammar.Text() == "Grammatik") {
						wordtype = s.Find(".dwdswb-ft-blocktext").Text()
						wordtype = strings.TrimSpace(wordtype)
						//fmt.Println("wordtype = ",wordtype)
					}
					if (grammar.Text() == "Worttrennung") {
						wordbreak = s.Find(".dwdswb-ft-blocktext").Text()
						wordbreak = strings.TrimSpace(wordbreak)
						//fmt.Println("wordbreak = ",wordbreak)
					}
				})
						
				doc.Find(".table-condensed td").Each(func(i int, s *goquery.Selection) {
					// for each item found
							
					//fmt.Println("s = ",s.Text())
					style, _ := s.Attr("style")
					//fmt.Println("style = ",style)
					if style == "text-align:right" {
						firstRow := s.Find("a").First()
						nextWord = firstRow.Text()
						nextWord = strings.TrimSpace(nextWord)
						//fmt.Println("firstRow = nextWord = ",firstRow.Text())
						link, _ := firstRow.Attr("href")
						nextWordLink = strings.TrimSpace(link)
						//fmt.Println("link = ",link)

						i = strings.LastIndexAny(nextWordLink, "/")
						GlobalConfig.WordToStartWithNext = nextWordLink[i+1:]
						//fmt.Println("WordToStartWithNext = ", GlobalConfig.WordToStartWithNext)
					}
				})
				GlobalConfig.CountWordsRead++
				//fmt.Println("word=",word,",wordtype=",wordtype,",wordbreak=",wordbreak,",nextWord=",nextWord)
				if strings.HasPrefix(wordtype, "Adjektiv") {
					GlobalConfig.CountWordsDetected++
					//fmt.Println("word=",word,",wordtype=",wordtype,",wordbreak=",wordbreak,",nextWord=",nextWord)
					//fmt.Println("CountWordsRead=",GlobalConfig.CountWordsRead,",CountWordsDetected=",GlobalConfig.CountWordsDetected,",CountWordsRequested=",GlobalConfig.CountWordsRequested,",CountWordsInserted=",GlobalConfig.CountWordsInserted)
		
					// do something with the result
					GlobalConfig.CountWordsLookup++
					requestUrl = GlobalConfig.WordListUrl + "/words?name=" + word
					fmt.Println("connect to wordlist and find out if word exists using = " + requestUrl)
					resp, err = http.Get(requestUrl)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					
					defer resp.Body.Close()
					body, err = ioutil.ReadAll(resp.Body)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					//fmt.Println("body = " + string(body))

					s.Text = ""
					json.Unmarshal(body, &s)
					if s.Text == "not found name = " + word {
						GlobalConfig.CountWordsRequested++
						fmt.Println("connect to wordlist and add word = " + word)

						wrd = Word{Id: 0, Name: word, New: true, Occurance: 0, Tests: nil}
						payload, err = json.Marshal(wrd)
						//fmt.Println("payload = " + string(payload))

						requestUrl = GlobalConfig.WordListUrl + "/words"
						req, err = http.NewRequest("POST", requestUrl, bytes.NewBuffer(payload))
						if err != nil {
							fmt.Println(err.Error())
							return
						}
						
						req.Header.Set("Content-Type", "application/json")
						client = &http.Client{}
						resp, err = client.Do(req)
						if err != nil {
							fmt.Println(err.Error())
							return
						}

						defer resp.Body.Close()
						body, err = ioutil.ReadAll(resp.Body)
						if err != nil {
							fmt.Println(err.Error())
							return
						}
						//fmt.Println("body = " + string(body))

						s.Text = ""
						json.Unmarshal(body, &s)
						if s.Text == "entity added" {
							GlobalConfig.CountWordsInserted++
						} else {
							fmt.Println("unexpected .. something went wrong")	
						}
					}
				}
			}
}