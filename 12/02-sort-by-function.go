package main

import (
	"fmt"
	"sort"
)

//In order to sort by a custom function in Go, we need a corresponding type. Here we’ve created a byLength type that is just an alias for the builtin []string type.
type byLength []string

//We implement sort.Interface - Len, Less, and Swap - on our type so we can use the sort package’s generic Sort function. Len and Swap will usually be similar across types and Less will hold the actual custom sorting logic. In our case we want to sort in order of increasing string length, so we use len(str[i]) and len(str[j]) here.
func (str byLength) Len() int {
	return len(str)
}
func (str byLength) Swap(i, j int) {
	str[i], str[j] = str[j], str[i]
}
func (str byLength) Less(i, j int) bool {
	return len(str[i]) < len(str[j])
}

func main() {
	humming := []string{"hmm", "hm", "hmmmmmmmm", "hmmmm"}
	//With all of this in place, we can now implement our custom sort by converting the original humming slice to byLength, and then use sort.Sort on that typed slice.
	sort.Sort(byLength(humming))
	fmt.Println(humming)
}
