package main

//
// go run -gcflags='-m -m' ./main.go

import (
	"fmt"
	"github.com/iimos/gorka"
	"github.com/iimos/gorka/alg"
	"github.com/iimos/gorka/gralang"
)

func main() {
	g := gorka.New()
	err := gralang.Parse(g, `
		a -> b c d
		b -> c
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println(g)

	a, _ := g.NodeByLabel("a")
	cc := alg.ClusteringCoefLocal(g, a)
	fmt.Printf("cc = %f\n", cc)
}

func experimentErdosRandomGraph() {
	var p float32
	for ; p < 1.0; p += 0.05 {
		tot := 100
		yes := 0
		for i := 0; i < tot; i++ {
			g, err := gorka.NewRandom(20, p)
			if err != nil {
				panic(err)
			}
			if g.IsConnected() {
				yes++
			}
		}
		fmt.Printf("p=%.2f: connected: %f\n", p, float32(yes)/float32(tot))
	}
}
