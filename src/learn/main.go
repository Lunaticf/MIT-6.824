package main

import "fmt"

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
	var m = make(map[string][]string)
	m["abs"] = append(m["abs"], "dwqd")
	m["abs"] = append(m["abs"], "dwqd")
	m["abs"] = append(m["abs"], "dwqd")
	m["abs"] = append(m["abs"], "dw")
	fmt.Println(m)
}
