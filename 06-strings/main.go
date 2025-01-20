// BLOG: https://go.dev/blog/strings
package main

import (
	"fmt"
)

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println()

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

  // [Exercise: Modify the examples above to use a slice of bytes instead of a string. Hint: Use a conversion to create the slice.]
	fmt.Println()
	fmt.Println("Slice of bytes:")
	byteSlice := []byte(sample)

	fmt.Println("Println:")
	fmt.Println(byteSlice)

	fmt.Println("Byte loop:")
	for i := 0; i < len(byteSlice); i++ {
		fmt.Printf("%x ", byteSlice[i])
	}
	fmt.Println()

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", byteSlice)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", byteSlice)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", byteSlice)

  fmt.Println("Printf with %+q:")
  fmt.Printf("%+q\n", byteSlice)

  // [Exercise: Loop over the string using the %q format on each byte. What does the output tell you?]

  fmt.Println()
  fmt.Println("%q en cada byte:")
  for i:= 0; i < len(sample); i++ {
    fmt.Printf("%q", sample[i])
  }

  fmt.Println()
  fmt.Println("byteSlice:")
  for i:= 0; i < len(byteSlice); i++ {
    fmt.Printf("%q", byteSlice[i])
  }
}
