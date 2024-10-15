package main

type S struct {
	data string
}

func (s S) Read() string { return s.data }

func (s *S) Write(str string) { s.data = str }

func main() {

	// Map values are not addressable
	// You cannot take the address of values stored in maps directly because map values are not "addressable" in Go.
	// As a result, you cannot call methods with pointer receivers on values stored in maps.
	sVals := map[int]S{
		1: {data: "TEST"},
		2: {"TEST"},
	}

	sVals[1].Read() // This works because Read has a value receiver

	// This doesn't work because Write requires a pointer receiver:
	// sVals[1].Write("test") // Compilation error

	// Storing pointers in maps:
	// If you store pointers in the map, the values are addressable,
	// and you can call methods with both value and pointer receivers.

	sPtrs := map[int]*S{1: {"TEST"}}

	sPtrs[1].Read()        // OK, because Read has a value receiver
	sPtrs[1].Write("test") // OK, because Write has a pointer receiver

}
