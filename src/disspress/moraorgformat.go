package disspress

import (
	// "bufio"
	// "fmt"
	"io/ioutil"
	// "reflect"
	// "os"
	"strings"
)

// file ...
func file(dir string) []string {
	file, err := ioutil.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), "\n")
}

// remove takes a []string and gives the same []string back with str removed
func removeIfContains(file []string, str string) []string {
	var indexesToDelete []int

	for line, word := range file {
		if strings.Contains(word, str) {
			indexesToDelete = append(indexesToDelete, line)
		}
	}

	for i := len(indexesToDelete) - 1; i >= 0; i-- {
		file = append(file[:indexesToDelete[i]], file[indexesToDelete[i]+1:]...)
		// this deletes! How cool!
	}
	return file
}

// removes section of org file with tag
func removeTag(file []string, tag string) []string {
	var indWithAst, indWithTag []int

	// get Tag lines
	for line, word := range file {
		if strings.Contains(word, ":"+tag+":") {
			indWithTag = append(indWithTag, line)
		}
	}

	// get Asterisk lines
	for line, word := range file {
		if strings.Contains(word, "*") && !strings.Contains(word, "|") {
			indWithAst = append(indWithAst, line)
		}
	}

	var del1, del2 []int

	for i, lAst := range indWithAst {
		if i >= len(indWithAst)-1 {
			break
		}
		for _, lTag := range indWithTag {
			if lTag == lAst {
				del1 = append(del1, lAst)
				del2 = append(del2, indWithAst[i+1])
			}
		}
	}

	for i := len(del1) - 1; i >= 0; i-- {
		// fmt.Print(del1[i], del2[i])
		file = removeLines(file, del1[i], del2[i])
	}

	return file
}

func removeLines(file []string, start, end int) []string {
	// fmt.Println(file[start])
	// fmt.Println(file[end-1])
	for i := end - 1; i >= start; i-- {
		file = append(file[:i], file[i+1:]...)
	}
	// fmt.Println("after", file[start])

	return file
}

// func main() {
// 	file := file("/home/justen/org/story/mora.org")

// 	file = removeIfContains(file, "#")
// 	file = removeIfContains(file, ".png")

// 	file = removeTag(file, "noexport")
// 	file = removeTag(file, "misc")

// 	fmt.Print(file)
// }
