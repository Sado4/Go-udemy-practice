package main

import "fmt"

func main() {
	var m = map[string]int{"a": 100, "b": 200}
	fmt.Println(m)

	m2 := map[string]int{"a": 100, "b": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}

	/*
		var m3 map[string]int
		m3["pc"] = 5000
		fmt.Println(m3)
	*/

	m4 := make(map[int]string)
	m4[1] = "US"
	m4[2] = "JAPAN"
	m4[3] = "CHINA"
	fmt.Println(m4)

	s := m["a"]
	fmt.Println(s)

	s1 := m["c"]
	fmt.Println(s1)

	s3, ok := m["b"]

	if !ok {
		fmt.Println("error")
	}

	s4, ok := m["c"]
	//_, ok = m["b"]

	m["b"] = 300

	m["c"] = 500

	delete(m, "c")

	fmt.Println(len(m))

	for k, v := range m {
		fmt.Println(k, v)
	}

	m5 := map[int][]string{
		1: {"A"},
		2: {"B", "C"},
		3: {"D", "E"},
	}

	m6 := map[int]map[float64]string{
		1: {3.14: "円周率"},
	}

}
