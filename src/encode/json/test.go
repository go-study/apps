package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	//must upper-case ,name or pass is wrong
	Name string
	//or use json tag to export lower-case
	Pass string `json:"pass"`
}

func (u *User) getName() string {
	return u.Name
}
func main() {
	var u User = User{Name: "xxxxxxx"}
	u.Pass = "p"
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println(u)
	fmt.Println(string(b))
}
