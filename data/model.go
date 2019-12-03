package data

import (
	"errors"
	"fmt"
	"sort"
)

var GlobalWordList WordList

type SorterWordByName []Word

func (a SorterWordByName) Len() int           { return len(a) }
func (a SorterWordByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SorterWordByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type SorterTestByName []Test

func (a SorterTestByName) Len() int           { return len(a) }
func (a SorterTestByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SorterTestByName) Less(i, j int) bool { 
	if a[i].Name < a[j].Name {
       return true
    }
    if a[i].Name > a[j].Name {
       return false
    }
    return a[i].Category < a[j].Category
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func Clear(wl WordList, w bool, t bool) WordList {
	
	if w {
		wl.LastUsedId = 0
		wl.Count = 0
		wl.Words = nil
	}
	if t {
		wl.Tests = nil
	}
	
	return wl
}
func RebuildTestIndex(wl WordList) WordList {

	GlobalWordList.Tests = nil
	
	for i := 0; i < len(wl.Words); i++ {
	
		// tests & category
		for j := 0; j < len(wl.Words[i].Tests); j++ {
			if !wl.containsTestCategory(wl.Words[i].Tests[j].Name, wl.Words[i].Tests[j].Category) {
				t := Test {
					Name: wl.Words[i].Tests[j].Name,
					Category: wl.Words[i].Tests[j].Category, 
				}
				newTests := append(wl.Tests, t)
				wl.Tests = newTests
			}
		}
    }
	sort.Sort(SorterTestByName(wl.Tests))
	
	return wl
}
func RebuildWordAndTestIndex(wl WordList) WordList {

	wl.LastUsedId = 0
	for i := 0; i < len(wl.Words); i++ {
		wl.LastUsedId++
        wl.Words[i].Id = wl.LastUsedId 
		wl.Words[i].Occurance = 0
		wl.Words[i].New = false
		
		// tests & category
		for j := 0; j < len(wl.Words[i].Tests); j++ {
			if !wl.containsTestCategory(wl.Words[i].Tests[j].Name, wl.Words[i].Tests[j].Category) {
				t := Test {
					Name: wl.Words[i].Tests[j].Name,
					Category: wl.Words[i].Tests[j].Category, 
				}
				newTests := append(wl.Tests, t)
				wl.Tests = newTests
			}
		}
    }
	wl.Count = len(wl.Words)
	
	return wl
}
func GetWordsListAsCsv() string {
	var list []string
	for i := 0; i < len(GlobalWordList.Words); i++ {
		list = append(list, GlobalWordList.Words[i].Name)
	}
	sort.Strings(list)

	var result string
	result = ""
	for i := 0; i < len(list); i++ {
		result += list[i] + ","
	}

	return result
}
func GetTestsList(name string) []string {
	var list []string
	for i := 0; i < len(GlobalWordList.Tests); i++ {
		if GlobalWordList.Tests[i].Name == name {
			list = append(list, GlobalWordList.Tests[i].Category)
		}
	}
	sort.Strings(list)

	return list
}
func AddWordToList(wl WordList, name string) (WordList, error) {
	fmt.Println("data/model.AddWordToList")
	fmt.Println("name=", name)
	fmt.Println("wl.containsWord(name)=", wl.containsWord(name))
	if wl.containsWord(name) {
		return wl, errors.New("already exists")
	}
	fmt.Println("before append wl=", wl)
	wl = appendWord(wl, name)
	fmt.Println("after append wl=", wl)
	return wl, nil
}