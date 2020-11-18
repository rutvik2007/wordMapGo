package main

import "fmt"

//ChannelMap hola
type ChannelMap struct {
	wordCounts   map[string]int
	askCountChan chan string
	getCountChan chan int

	addWordChan chan string

	askReduceMapChan chan ReduceInfoStruct
	getReduceMapChan chan ReduceInfoStruct

	killChannel chan int
}

//ReduceInfoStruct packages information to reduce count map into a struct
type ReduceInfoStruct struct {
	functor ReduceFunc
	word    string
	count   int
}

// Listen hola
func (cm ChannelMap) Listen() {
	var redstruct ReduceInfoStruct
	var word string
	for {
		select {
		case word = <-cm.askCountChan: ///rutvik
			cm.getCountChan <- cm.wordCounts[word]
		case word = <-cm.addWordChan: /// alejandro //prolly bad syntax for assigning to word variable
			//...

		case redstruct = <-cm.askReduceMapChan: ///rutvik  //change variable
			for word, count := range cm.wordCounts {
				redstruct.word, redstruct.count = redstruct.functor(redstruct.word, redstruct.count, word, count)
			}
			cm.getReduceMapChan <- redstruct

		case <-cm.killChannel: ///alejandro
			//...
		}
	}

}

//Stop hola
func (cm ChannelMap) Stop() {
	return
}

// Reduce hola
func (cm ChannelMap) Reduce(functor ReduceFunc, accumStr string, accumInt int) (string, int) {
	var redstruct ReduceInfoStruct
	redstruct.functor = functor
	redstruct.word = accumStr
	redstruct.count = accumInt
	cm.askReduceMapChan <- redstruct
	redstruct = <-cm.getReduceMapChan
	return redstruct.word, redstruct.count
}

// AddWord hola
func (cm ChannelMap) AddWord(word string) {

	return

}

// GetCount hola
func (cm ChannelMap) GetCount(word string) int {
	cm.askCountChan <- word
	return <-cm.getCountChan
}

//NewChannelMap returns a new channelMap
func NewChannelMap() *ChannelMap {
	var newChanMap ChannelMap
	newChanMap.wordCounts = make(map[string]int)
	newChanMap.askCountChan = make(chan string, ASK_BUFFER_SIZE)
	newChanMap.getCountChan = make(chan int, ASK_BUFFER_SIZE)
	newChanMap.addWordChan = make(chan string, ADD_BUFFER_SIZE)
	newChanMap.askReduceMapChan = make(chan ReduceInfoStruct, ASK_BUFFER_SIZE)
	newChanMap.getReduceMapChan = make(chan ReduceInfoStruct, ASK_BUFFER_SIZE)
	newChanMap.killChannel = make(chan int)
	return &newChanMap
}

//NewLockingMap returns a new ChannelMap
func NewLockingMap() *ChannelMap {

	return NewChannelMap()
}

func main() {
	fmt.Printf("hello world\n")
}
