package data

import (
	"errors"
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
    Type  string	`json:"type"`
    LastUsedId  int	`json:"lastusedid"`
    Count  int	`json:"count"`
	Tests []Test	`json:"tests"`
	Words []Word	`json:"words"`
}

func (wl WordList) containsTestCategory(test string, category string) bool {
	for i := 0; i < len(wl.Tests); i++ {
        if wl.Tests[i].Name == test && wl.Tests[i].Category == category {
			return true
        }
    }
    return false
}
func (wl WordList) containsWord(name string) bool {
	for i := 0; i < len(wl.Words); i++ {
        if wl.Words[i].Name == name {
            return true
        }
    }
    return false
}
func clearNewAndOccurance(wl WordList) WordList {
	for i := 0; i < len(wl.Words); i++ {
        wl.Words[i].New = false
        wl.Words[i].Occurance = 0
    }
	
	return wl
}
func (wl WordList) DeleteWordById(id int) (WordList, error) {
    
	if id < 1 || id > wl.LastUsedId {
		return wl, errors.New("not found") 
	} else {
		for i := 0; i < len(wl.Words); i++ {
			if wl.Words[i].Id == id {
			
				// Remove the element at index i from wl
				copy(wl.Words[i:], wl.Words[i+1:]) 		// Shift a[i+1:] left one index.
				wl.Words = wl.Words[:len(wl.Words)-1]  	// Truncate slice.
				wl.Count = len(wl.Words)
				return wl, nil
			}
		}
	}
    return wl, errors.New("not found")
}
func (wl WordList) GetWordById(id int) (Word, error) {
    var w Word
	
	if id < 1 || id > wl.LastUsedId {
		return w, errors.New("not found") 
	} else {
		for i := 0; i < len(wl.Words); i++ {
			if wl.Words[i].Id == id {
				w = wl.Words[i]
				return w, nil
			}
		}
	}
    return w, errors.New("not found")
}
func (wl WordList) GetWordByName(name string) (Word, error) {
    var w Word
	
	for i := 0; i < len(wl.Words); i++ {
		if wl.Words[i].Name == name {
			w = wl.Words[i]
			return w, nil
		}
	}
	
    return w, errors.New("not found")
}
func appendWord(wl WordList, name string) WordList {
	var foundWord bool = false
	for i := 0; i < len(wl.Words); i++ {
		if wl.Words[i].Name == name {
			foundWord = true
			return wl
        }
    }
	if !foundWord {	
		wl.LastUsedId++
		w := Word {
			Id: wl.LastUsedId,
			Name: name,
			Occurance: 0,
			New: true,
		}
		wl.Words = append(wl.Words, w)
		wl.Count = len(wl.Words)
	}
	return wl
}
func appendWordIncludingTest(wl WordList, name string, test string, category string) WordList {
	var foundWord bool = false
	for i := 0; i < len(wl.Words); i++ {
		if wl.Words[i].Name == name {
			foundWord = true
			if test != "" && category != "" {
				var foundTest bool = false
				for j := 0; j < len(wl.Words[i].Tests); j++ {
					if  wl.Words[i].Tests[j].Name == test && wl.Words[i].Tests[j].Category == category {
						foundTest = true
						return wl
					}
				}
				if !foundTest {		
					t := Test {
						Name: test,
						Category: category, 
					}
					wl.Words[i].Tests = append(wl.Words[i].Tests, t)
					return wl
				}
			}
        }
    }
	if !foundWord {	
		wl.LastUsedId++
		w := Word {
			Id: len(wl.Words),
			Name: name,
			Occurance: 0,
			New: true,
		}
		if test != "" && category != "" {			
			t := Test {
				Name: test,
				Category: category, 
			}
			w.Tests = append(w.Tests, t)
		}
		wl.Words = append(wl.Words, w)
		wl.Count = len(wl.Words)
	}
    return wl
}
