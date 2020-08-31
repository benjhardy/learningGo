package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
	//

}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {

	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	currentChoice := node.choices
	for currentChoice != nil {
		fmt.Println(currentChoice.cmd, ":", currentChoice.description)
		currentChoice = currentChoice.nextChoice
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	currentChoice := node.choices
	for currentChoice != nil {
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd) {
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Println("Sorry, I didnt understand that.")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func printArray(a []string) {
	for _, e := range a {
		fmt.Print(e)
	}
}
func main() {

	//abc := []string{"a", "b", "c"}

	//abc = append(abc, "d")

	//printArray(abc)

	scanner = bufio.NewScanner(os.Stdin)

	//start := storyNode(text: "first one")
	start := storyNode{text: `
	You are in a large chamber, deep underground. You see three passafes leading out. a north passafe laedsinto darkness. to the soutth, a pass
	age appears to head upward. The eastern passage appears flat and well traveled.	`}

	darkRoom := storyNode{text: "It is pitch black. You cannot see a thing"}

	darkRoomLit := storyNode{text: "The dark passage is now lit by your lantern. You can continue north or head back south"}

	grue := storyNode{text: "While walking around in the darkness you are eaten by a grue"}

	trap := storyNode{text: "You head down the well travelled path when suddenly a trap door opens and you fall into a pit"}

	treasure := storyNode{text: "You arrive at a small chamber, filled with treasure!"}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try to go back souuth", &grue)
	darkRoom.addChoice("O", "Turn on Lantern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	start.play()
	fmt.Println()
	fmt.Println("THE END!")
}
