package main

import (
	"fmt"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}

}

func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}
}

// O(n) search through n number possibilities,

func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

// Delete

func main() {

	page1 := storyPage{"A long time ago in a galaxy far far away", nil}
	page1.addToEnd("You are alone, and you need to find the sacred helmet before the bad guys do")
	page1.addToEnd("You see a troll ahead")

	page1.addAfter("Testing add After")
	// Functions - has a return value, may execute commands as well but
	// Procedures - Has no return value, just executes
	// methods - functions attached to a struct/object/etc
	page1.playStory()

}
