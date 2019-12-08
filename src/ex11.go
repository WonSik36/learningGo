package main

// collection map

import "fmt"

func main(){
	var idMap map[int]string  // key: int, value: string, Nil map
	idMap = make(map[int]string)
	idMap[901] = "Apple"
    idMap[134] = "Grape"
    idMap[777] = "Tomato"


	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}

	println(idMap[901])
	delete(idMap, 901)
	println(idMap[901])	// Nil -> empty space
	println(tickers["GOOG"])

	val, exists := tickers["MSFT"]
    if !exists {
        println("No MSFT ticker")
    }else{
		println(val)
	}

	for key, val := range tickers {
        fmt.Println(key, val)
    }
}