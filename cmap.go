package main

//ChannelMap to provide concurrent access to the wordCount map using channels
type ChannelMap struct {
	wordCounts map[string]int

	askCountChan chan string
	getCountChan chan int

	addWordChan chan string

	askReduceMapChan chan ReduceInfoStruct
	getReduceMapChan chan ReduceInfoStruct

	killStream chan int
}

//ReduceInfoStruct packages information to reduce count map into a struct
type ReduceInfoStruct struct {
	functor ReduceFunc
	word    string
	count   int
}

// Listen to requests
func (cm ChannelMap) Listen() {
	var (
		word      string
		redStruct ReduceInfoStruct
	)

	for {
		select {
		case word = <-cm.askCountChan: ///rutvik
			cm.getCountChan <- cm.wordCounts[word]

		case word = <-cm.addWordChan: /// alejandro=
			if _, ok := cm.wordCounts[word]; ok {
				cm.wordCounts[word]++
			} else {
				cm.wordCounts[word] = 1
			}

		case redStruct = <-cm.askReduceMapChan: ///rutvik
			for word, count := range cm.wordCounts {
				redStruct.word, redStruct.count = redStruct.functor(redStruct.word, redStruct.count, word, count)
			}
			cm.getReduceMapChan <- redStruct
		case <-cm.killStream: ///alejandro
			break
		}
	}

}

//Stop listening to requests
func (cm ChannelMap) Stop() {
	cm.killStream <- 69420
	return
}

// Reduce the wordCount map
func (cm ChannelMap) Reduce(functor ReduceFunc, accumStr string, accumInt int) (string, int) {
	var redstruct ReduceInfoStruct
	redstruct.functor = functor
	redstruct.word = accumStr
	redstruct.count = accumInt
	cm.askReduceMapChan <- redstruct
	redstruct = <-cm.getReduceMapChan
	return redstruct.word, redstruct.count
}

// AddWord to the wordCount map
func (cm ChannelMap) AddWord(word string) {
	cm.addWordChan <- word
	return
}

// GetCount for word from wordCount map
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
	newChanMap.killStream = make(chan int)
	return &newChanMap
}

//NewLockingMap returns a new ChannelMap
func NewLockingMap() *ChannelMap {
	return NewChannelMap()
}

func min_word(w1 string, c1 int, w2 string, c2 int) (string, int) {
	if c1 < c2 {
		return w1, c1
	}
	return w2, c2
}
