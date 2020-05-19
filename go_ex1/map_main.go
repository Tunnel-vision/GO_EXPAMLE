package main

import "fmt"

type Vertext struct {
	Lat, Long float64
}

//var m map[string]Vertext

func map1_main() {
	m := make(map[string]Vertext)
	m["Bell Labs"] = Vertext{
		40.6846,
		-74.414,
	}
	fmt.Println(m["Bell Labs"])
}

func main() {
	map1_main()
}
