package main

import "fmt" 

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex {
		40.68, -74.33,
	}
	m["key2"] = Vertex {
		21.33,22.12,
		}
	m["key3"] = Vertex {
			11.21,33.33,
		}
	v,ok := m["key3"]
	if ok {
		fmt.Println("value exists for key3 in the map as " ,v) 
	}
	delete(m,"key3")
	v,ok = m["key3"]
	if !ok {
		fmt.Println("value doesn't exists for key3 in the map")
	}
	fmt.Println(m["Bell Labs"])
}
