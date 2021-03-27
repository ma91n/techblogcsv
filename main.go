package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/ma91n/hexoreader"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Article struct {
	Title     string
	Date      string
	Tags      []string
	Category  string
	Author    string
	CharCount int64
}

func main() {

	paths := dirwalk("C:\\Users\\manoj\\tech-blog\\source\\_posts")

	var posts []hexoreader.Post

	for _, path := range paths {
		file, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("readFile %s: %v", path, err)
		}

		post, err := hexoreader.New(bytes.NewReader(file)).ReadAll()
		if err != nil {
			log.Fatalf("hexoreader %s: %v", path, err)
		}
		posts = append(posts, post)
	}

	var output [][]string

	for _, p := range posts {
		output = append(output, []string{p.Title, strings.Join(p.Categories, "|"), strings.Join(p.Tags, "|"), p.Date, fmt.Sprint(len(p.Content))})
	}

	w := csv.NewWriter(os.Stdout)
	w.Write([]string{"title","categories","tags","date","char_count"})
	if err := w.WriteAll(output); err != nil {
		log.Fatal(err)
	}

}

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
