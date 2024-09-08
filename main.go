package main

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func main() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	id := node.Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64   ID: %d\n", id)
	fmt.Printf("String  ID: %s\n", id)
	fmt.Printf("Base2   ID: %s\n", id.Base2())
	fmt.Printf("Base36  ID: %s\n", id.Base36())
	fmt.Printf("Base58  ID: %s\n", id.Base58())
	fmt.Printf("Base64  ID: %s\n", id.Base64())
	fmt.Printf("Bytes    : %#v\n", id.Bytes())
	fmt.Printf("IntBytes : %#v\n", id.IntBytes())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())
	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())
	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	fmt.Printf(" 2nd generation \n")
	fmt.Printf("------------------------------------------------------ \n")

	id = node.Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base36  ID: %s\n", id.Base36())
	fmt.Printf("Base58  ID: %s\n", id.Base58())
	fmt.Printf("Base64 ID: %s\n", id.Base64())
	fmt.Printf("Bytes    : %#v\n", id.Bytes())
	fmt.Printf("IntBytes : %#v\n", id.IntBytes())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())
	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())
	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

}
