package main

import (
	"reflect"
	"fmt"
)

func main(){
	var i int = 3
	fmt.Println("type :",reflect.TypeOf(i))
	fmt.Println("value :",reflect.ValueOf(i))

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
