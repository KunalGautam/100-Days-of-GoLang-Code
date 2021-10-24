// Reference Article https://medium.com/@felipedutratine/how-to-organize-the-go-struct-in-order-to-save-memory-c78afcf59ec2
package main

import (
	"fmt"
	"unsafe"
)

type nonOptimized struct {
	someInt    bool    // 1 byte
	someFloat  float64 // 8 bytes
	someBool   int32   // 4 bytes
	someInt1   bool    // 1 byte
	someFloat1 float64 // 8 bytes
	someBool1  int32   // 4 bytes
}

type optimizedStruct struct {
	someFloat  float64 // 8 bytes
	someFloat1 float64 // 8 bytes
	someBool   int32   // 4 bytes
	someBool1  int32   // 4 bytes
	someInt    bool    // 1 byte
	someInt1   bool    // 1 byte
}

func main() {
	a := nonOptimized{}
	b := optimizedStruct{}

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))

}
