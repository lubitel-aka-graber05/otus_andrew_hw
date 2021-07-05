package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var editString = regexp.MustCompile(`(?m)[0-9]`)

type WorkingWithWords struct {
	countComma         int
	countDot           int
	countedCommaAndDot []int
	tempSl             []string
}

func (w *WorkingWithWords) CountCommaAndDot(s string) []int {
	for _, s1 := range s {
		if s1 == ',' {
			w.countComma++
		}
		if s1 == '.' {
			w.countDot++
		}
	}
	w.countedCommaAndDot = append(w.countedCommaAndDot, w.countComma, w.countDot)
	return w.countedCommaAndDot
}

func (w *WorkingWithWords) AppendCountSort(s string) []string {
	if s == "" {
		return w.tempSl
	}
	w.tempSl = strings.Fields(s)
	countWords := make(map[string]int)
	for _, s := range w.tempSl {
		countWords[s]++
	}
	w.tempSl = nil
	for s, n := range countWords {
		w.tempSl = append(w.tempSl, strconv.Itoa(n)+s)
	}

	sort.Slice(w.tempSl, func(i, j int) bool {
		return w.tempSl[j] < w.tempSl[i]
	})
	w.tempSl = w.tempSl[:10]
	for i := 0; i < len(w.tempSl); i++ {
		w.tempSl[i] = editString.ReplaceAllString(w.tempSl[i], "")
	}
	sort.Strings(w.tempSl[1:3])
	sort.Strings(w.tempSl[3:5])
	sort.Strings(w.tempSl[5:])
	return w.tempSl
}

var hardWork WorkingWithWords

func Count(s string) []int {
	return hardWork.CountCommaAndDot(s)
}

func Top10(s string) []string {
	return hardWork.AppendCountSort(s)
}
