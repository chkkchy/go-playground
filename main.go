package main

import (
	"encoding/json"
	"fmt"

	"github.com/chkkchy/go-playground/constant"
)

func main() {
	gender := constant.GetGender(1)
	fmt.Println(gender.ID, gender.Value)

	var list []constant.Dataiface
	for _, v := range constant.GenderList {
		list = append(list, v)
	}
	r := constant.ToData("gender", list)
	b, _ := json.Marshal(r)
	fmt.Println(string(b))

}
