package Data

/*
	This file will holds the words that shouldn't be on chat or send as a data to the server
	It is better to not change edit this file and instead make a words list on config/words.l
	There words are default for each server, and you should not remove these.
*/

var LoadedWordList *[]string

var defaultWords []string = []string{
	"321",
	"nice",
	"yes",
}

func LoadWordsFromConfig() {
	// Load Words From Config File
	LoadedWordList = &[]string{"config"}
	*LoadedWordList = append(*LoadedWordList, defaultWords...)
}

func GetAllWords() *[]string {
	return LoadedWordList
}
