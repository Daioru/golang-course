package main

import (
	"fmt"
	"reflect"
)

type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

func main() {
	typ1 := reflect.TypeOf(perms{})
	typ2 := reflect.TypeOf(contact{})
	fmt.Println(typ1.Size(), typ2.Size())
}
