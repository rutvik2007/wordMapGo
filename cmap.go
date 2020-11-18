package main

import "fmt"

//ChannelMap hola
type ChannelMap struct {
	wordCounts   map[string]int
	askCountChan chan string
	getCountChan chan int
	addWordChan  chan string
	reduceMapChan 
}

type reduceStruct struct{

}

// Listen hola
func (cm ChannelMap) Listen() {
	return
}

//Stop hola
func (cm ChannelMap) Stop() {
	return
}

// Reduce hola
func (cm ChannelMap) Reduce(functor ReduceFunc, accumStr string, accumInt int) (string, int) {

	return "Hello", 1
}

// AddWord hola
func (cm ChannelMap) AddWord(word string) {

	return

}

// GetCount hola
func (cm ChannelMap) GetCount(word string) int {
	return 1
}

//NewChannelMap returns a new channelMap
func NewChannelMap() *EmergingMap {
	var newChanMap ChannelMap
	newChanMap.wordCounts = make(map[string]int)
	newChanMap.askCountChan = make(chan string)
	newChanMap.getCountChan = make(chan int)
	newChanMap.reduceMapChan = make(chan )
	return &newChanMap
}

//NewLockingMap returns a new ChannelMap
func NewLockingMap() *EmergingMap {

	return NewChannelMap()
}

func main() {
	fmt.Printf("hello world\n")
}
