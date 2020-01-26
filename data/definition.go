package data

import (
	//"errors"
	//"fmt"
)

type Status struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
type Test struct {
    Name  string	`json:"name"`
    Category  string	`json:"category"`
}
type Word struct {
	Id int	`json:"id"`
    Name  string	`json:"name"`
	Occurance int	`json:"occurance"`
	New bool	`json:"new"`
	Tests []Test	`json:"tests"`
}
type WordList struct {
	Session SessionStatus	`json:"session"`
	Words []Word	`json:"words"`
	Tests []Test	`json:"tests"`
}
type Config struct {
    RequestExecution  bool	`json:"requestexecution"`
    WordListUrl  string	`json:"wordlisturl"`
    WordListExtractorUrl  string	`json:"wordlistextractorurl"`
    WordToStartWith  string	`json:"wordtostartwith"`
    WordToStartWithNext  string	`json:"wordtostartwithnext"`
	CountWordsRead  int	`json:"countwordsread"`
	CountWordsDetected  int	`json:"countwordsdetected"`
	CountWordsLookup  int	`json:"countwordlookup"`
	CountWordsRequested  int	`json:"countwordrequested"`
    CountWordsInserted  int	`json:"countwordsinserted"`
}

