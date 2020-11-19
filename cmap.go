package main

import "fmt"

//ChannelMap hola
type ChannelMap struct {
	wordCounts   map[string]int

	askCountChan chan string
	getCountChan chan int

	addWordChan  chan string

	askReduceMapChan chan ReduceInfoStruct
	getReduceMapChan chan ReduceInfoStruct

	killChannel chan int
}

type ReduceInfoStruct struct{
	functor ReduceFunc
	word string
	count int
}

// Listen hola
func (cm ChannelMap) Listen() {

	var (
		word string
		redInfo ReduceInfoStruct
	)

	for {
		select {
			case <-askCountChan:	///rutvik
				fmt.Printf("in askCountChan")
				temp_return_val := 679
				getCountChan<-temp_return_val

			case word<-addWordChan:		/// alejandro
				fmt.Printf("in addWordChan, word = %s\n",word)
				if _, ok := cm.wordCounts[word]; ok {
					cm.wordCounts[word]++
				} else {
					cm.wordCounts[word] = 1
				}

			case redInfo<-askReduceMapChan:	///rutvik  //change variable name 'redstruct'
				//...

			case <-killChan:		///alejandro
				fmt.Printf("in the killChan case\n")
				return//...
		}
	}

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

	addWordChan<-word

}

// GetCount hola
func (cm ChannelMap) GetCount(word string) int {
	return 1
}

//NewChannelMap returns a new channelMap
func NewChannelMap() *EmergingMap {
	var newChanMap ChannelMap
	newChanMap.wordCounts = make(map[string]int)
	newchanmap.askCountChan = make(chan string)
	newchanmap.getCountChan = make(chan int)
	newChanMap.addWordChan = make(chan string)
	newChanMap.askReduceMapChan = make(chan ReduceInfoStruct)
	newChanMap.getReduceMapChan = make(chan ReduceInfoStruct)
	return &newChanMap
}
////type ChannelMap struct {
////	wordCounts   map[string]int
////
////	askCountChan chan string
////	getCountChan chan int
////
////	addWordChan  chan string
////
////	askReduceMapChan chan ReduceInfoStruct
////	getReduceManChan chan ReduceInfoStruct
////
////	killChannel chan int
////}

//NewLockingMap returns a new ChannelMap
func NewLockingMap() *EmergingMap {

	return NewChannelMap()
}

///func main() {
///	fmt.Printf("hello world\n")
///}






