package main

import (
	"fmt"
	_ "kalista/config"
	_ "kalista/utils/sources"
	"kalista/utils/struct2json"
)

func main() {
	//routers.Start()

	amain()
}

type My struct {
	A int `json:"a"`
}

type myStruct struct {
	M   My `json:"m" noroot:"true"`
	*My `noroot:"true"`

	B *int
}

func amain() {

	input := myStruct{
		My: &My{
			A: 1,
		},
	}

	ans, _ := struct2json.Convert(input)

	fmt.Println(ans)
}
