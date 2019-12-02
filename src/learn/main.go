package main

import (
	"fmt"
	"io/ioutil"
	"mapreduce"
	"strings"
	"unicode"
)

// learn go
type KeyValue struct {
	Key   string
	Value string
}

func main() {
	//file, _ := os.Open("a.txt")
	//decoder := json.NewDecoder(file)
	//for {
	//	kv := new(KeyValue)
	//	err := decoder.Decode(&kv)
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println(kv.Key, kv.Value)
	//}
	bytes, _ := ioutil.ReadFile("src/main/pg-grimm.txt")
	mapF("dwqd", string(bytes))
}

func mapF(filename string, contents string) []mapreduce.KeyValue {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	words := strings.FieldsFunc(contents, f)

	var kvs []mapreduce.KeyValue
	for _, word := range words {
		kvs = append(kvs, mapreduce.KeyValue{word, "1"})
	}
	fmt.Println(kvs)
	return kvs
}
