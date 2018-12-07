package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Keys struct {
	Keys []int64 `json:"keys"`
}

func main() {
	keybytes, err := ioutil.ReadFile("key.json")
	if err != nil {
		panic(err.Error())
	}
	var keys Keys
	err = json.Unmarshal(keybytes, &keys)
	if err != nil {
		panic(err.Error())
	}
	data, err := ioutil.ReadFile("data.dat")
	if err != nil {
		panic(err.Error())
	}
	dataMod := len(data) % 16
	if dataMod != 0 {
		for i := 0; i < ((len(data)/16)+1)*16; i++ {
			data = append(data, 0)
		}
	}
	fmt.Println(data)
}

func encrypt(left int64, right int64, keys []int64) (int64, int64) {
	for _, k := range keys {
		temp := right ^ f(left, k)
		right = left
		left = temp
	}
	return left, right
}

func decrypt(left int64, right int64, keys []int64) (int64, int64) {
	for i, _ := range keys {
		temp := left ^ f(right, keys[len(keys)-i-1])
		left = right
		right = temp
	}
	return left, right
}

func f(block int64, key int64) int64 {
	return (block * key) % 64
}
