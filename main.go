package main

import (
	"go-data-structure/bst"
)

func main() {
	bst := bst.NewBST()
	bst.Insert(60)
	bst.Insert(50)
	bst.Insert(51)
	bst.Insert(49)
	bst.Insert(49)
	bst.Insert(28)
	bst.Insert(39)
	bst.Insert(65)
	bst.Insert(98)
	bst.Insert(1)
	bst.Insert(8)
	bst.Insert(77)
	bst.Insert(100)
	
	println("Quantidade de nós na árvore: ",bst.Count())

	println("Procurar elemento: ",bst.Search(98))

	//remoção de folha


	bst.InOrder()

	println("Remover elemento: ",bst.Remove(8))

	bst.InOrder()

	println("Remover elemento: ",bst.Remove(1))

	bst.InOrder()

	println("Min ",bst.Min().Value)
	println("Max ",bst.Max().Value)
}
