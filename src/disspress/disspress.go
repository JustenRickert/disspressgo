package disspress

import (
	// "crypto"
	"fmt"
	"io/ioutil"
	// "log"
	"math/rand"
	// "os"
	// "os/exec"
	// "bufio"
	"strings"
	"time"
)

func getSplitMoraFile() (splitFile []string) {
	// filename := "/home/justen/org/story/mora.org"
	filename := "/home/justen/GO/textfile"
	file, err := ioutil.ReadFile(filename)

	splitFile = strings.Split(string(file), " ")

	if err != nil {
		panic(err)
	}

	return splitFile
}

func getWordCount(file []string, which string) int {
	wordCount := make(map[string]int)

	for _, word := range file {
		wordCount[word] += 1
	}
	return wordCount[which]
}

// writeWordWithSleep writes characters one at a time with sleepTime sleep time
// between them
func writeWordWithSleep(word string) {
	for _, char := range word {
		fmt.Print(string(char))
		time.Sleep(50 * time.Millisecond)
	}
}

func indexOfAll(myWord string, file []string) []int {
	var indexes []int

	for i, word := range file {

		if lastLetterIs(".", word) {
			word = word[0 : len(word)-1]
		}

		if myWord == word {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func lastLetterIs(this string, inThat string) bool {
	return string(inThat[len(inThat)-1]) == this
}

func firstLetterIs(this string, inThat string) bool {
	return string(inThat[0]) == this
}

func WriteDissociatively(file []string, i int) {
	r := rand.New(rand.NewSource(int64(i))) // use r.Int() for random
	inParens := false

	for {
		word := file[i]
		writeWordWithSleep(word + " ")

		if firstLetterIs("\"", word) {
			inParens = true

		} else if lastLetterIs(".", word) {
			time.Sleep(1000 * time.Millisecond)

		} else if lastLetterIs("\"", word) {
			inParens = false
			time.Sleep(1000 * time.Millisecond)
		}

		if !inParens && r.Int()%14 == 0 {
			indexesNewWord := indexOfAll(word, file)

			if len(word) >= 4 && len(indexesNewWord) > 0 {
				i = indexesNewWord[r.Intn(len(indexesNewWord))]
			}
		}

		i += 1
		if i >= len(file) {
			i = 0
		}

	}
}
