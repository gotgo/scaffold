package main

import "github.com/facebookgo/inject"

func main() {

}

func dependencies() {
	var graph inject.Graph
	err := graph.Provide(
		&inject.Object{Value: "", Name: "myname"},
	)
	if err != nil {
		panic(err)
	}

	if err := graph.Populate(); err != nil {
		panic(err)
	}
}
