package hw4

import (
	"fmt"
	"reflect"
)

type SomeName struct {
	a int8  // 1 byte
	c int8  // 1 byte
	b int64 // 8 byte
}

type SomeName1 struct {
	FirstField int
	SecondField *bool
	ThirdField string
	FourthField uint64
	FifthField string
	SixthField *byte
}


func structMemOpt() {

	data := SomeName{}
	typeOfStruct := reflect.TypeOf(data)
	n := typeOfStruct.NumField()
	var totalSize int
	for i := 0; i < n; i++ {
		field := typeOfStruct.Field(i)
		fmt.Printf("Type: %v\t, Size: %v\t, Offset: %v\t, Align: %v\n",
			field.Type.String(), field.Type.Size(), field.Offset, field.Type.Align(),
		)
		if i == n - 1 {
			totalSize = int(field.Offset + field.Type.Size())
		}
	}
	fmt.Println(totalSize)
	fmt.Println(typeOfStruct.Size())
}
