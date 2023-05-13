package main

import "fmt"

func main() {
	map_ := make(map[string]int)
	map_["k1"] = 1
	map_["k2"] = 2
	fmt.Println("map:", map_, "size:", len(map_))

	value_1 := map_["k1"]
	value_3 := map_["k3"]
	fmt.Println("get k1 and k3 from map_:", value_1, value_3)

	delete(map_, "k2")
	_, is_present := map_["k2"]
	fmt.Println("Is k2 in map_:", is_present)

	map1_ := map[string]int {"a": 1, "b": 2}
	fmt.Println("map1:", map1_)
}