package main

import (
	"encoding/json"
	"fmt"

	"github.com/chkkchy/go-playground/constant"
)

func main() {
	gender := constant.GetGender(1)
	fmt.Println(gender.ID, gender.Value)

	r := constant.ToData("gender", gender)
	b, _ := json.Marshal(r)
	fmt.Println(string(b))

}
