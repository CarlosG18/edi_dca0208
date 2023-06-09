package main

import (
	//"fmt"
	//"main/packs/sort"
	"main/packs/hash"
	//"math/rand"
	//"time"
)

func main() {
	table := hash.Init(10)
	table.Add(1, "carlos")
	table.Add(2, "edu")
	table.Add(11, "jurema")
	table.Add(32, "valdir")
	//tupla_carlos := table.Get(1)
	table.Print_table()
	table.Remove(11)
	table.Print_table()
	//table.Add(11, "jurema")
	table.Print_table()
	//temKey := table.Search(11)
	//fmt.Println("tem a key 11? ", temKey)
}
	
	/*arvore := tree.NewNode(20)
	arvore = arvore.Add(29)
	arvore = arvore.Add(35)
	arvore = arvore.Add(51)
	arvore = arvore.Add(13)
	arvore = arvore.Add(25)
	arvore = arvore.Add(18)
	arvore = arvore.Add(19)
	arvore.PrintPre()

	
		v := []int{10, 7, 3, 8, 15, 20}
		t := tree.CreateBst(v)
		u := tree.BstNode{}
		u.Add(10)
		u.Add(7)
		u.Add(3)
		u.Add(8)
		u.Add(20)
		t.Remove(3)
		t.Remove(20)
		u.Remove(3)
		fmt.Println("size da arvore = ", t.Size())
		t.PrintIn()
		fmt.Println("size da arvore = ", u.Size())
		u.PrintIn()

	 testes algortitmos de ordenação
	vetor := FullArrayDecres(10000)
	fmt.Println("vetor = ", vetor)
	fmt.Println()

	//BubbleSort
	start := time.Now()
	vetor_bubble := sort.BubbleSort(vetor)
	end := time.Since(start)
	fmt.Println("vetor ordenado com BubbleSort = ", vetor_bubble)
	fmt.Println()

	//SelectionSort
	start1 := time.Now()
	vetor_selection := sort.SelectionSort(vetor)
	end1 := time.Since(start1)
	fmt.Println("vetor ordenado com SelectionSort = ", vetor_selection)

	//InsertionSort
	start2 := time.Now()
	vetor_insertion := sort.InsertionSort(vetor)
	end2 := time.Since(start2)
	fmt.Println("vetor ordenado com InsertionSort = ", vetor_insertion)

	//MergeSort
	start3 := time.Now()
	vetor_merge := sort.MergeSort(vetor,len(vetor))
	end3 := time.Since(start3)
	fmt.Println("vetor ordenado com MergeSort = ", vetor_merge)

	vetor_quick := vetor

	//QuickSort
	start4 := time.Now()
	sort.QuickSort(vetor_quick,0,len(vetor_quick)-1)
	end4 := time.Since(start4)
	fmt.Println("vetor ordenado com QuickSort = ", vetor_quick)

	//CountingSort
	start5 := time.Now()
	vetor_couting := sort.CountingSort(vetor)
	end5 := time.Since(start5)
	fmt.Println("vetor ordenado com CountingSort = ", vetor_couting)
	fmt.Println()

	fmt.Println("tempo gasto com BubbleSort O(n²) = ", end)
	fmt.Println("tempo gasto com SelectionSort O(n²) = ", end1)
	fmt.Println("tempo gasto com InsertionSort O(n²) = ", end2)
	fmt.Println("tempo gasto com MergeSort O(logn) = ", end3)
	fmt.Println("tempo gasto com QuickSort O(nlogn) = ", end4)
	fmt.Println("tempo gasto com CountingSort O(n) = ", end5)


	// fazendo testes em arvores
	// vetor := []string{"c","a","r","l","o","s"}
	// t := tree.CreateBst_str(vetor)
	// // bst_verify := t.IsBst()

	// if t != nil{
	//   t.PrintPos()
	//   // fmt.Println(bst_verify)
	// }else{
	//   fmt.Println(t)
	// }
}


func Add_numeros(qtd int, intervalo int) {
	tree := tree.BstNode{}
	vetor_numeros := FullArrayRandom(qtd, intervalo)

	for i := 0; i < len(vetor_numeros); i++ {
		tree.Add(vetor_numeros[i])
	}
	TestTree(&tree, vetor_numeros)
	fmt.Print("print In = ")
	tree.PrintIn()
	fmt.Println()
	fmt.Print("print Pre = ")
	tree.PrintPre()
	fmt.Println()
	fmt.Print("print Pos = ")
	tree.PrintPos()
	fmt.Println()
	max := tree.Max()
	min := tree.Min()
	fmt.Println("max = ", max, " min = ", min)
	fmt.Println()

}

func TestTree(tree *tree.BstNode, vetor []int) {
	fmt.Println("testando a busca na árvore = ", vetor)

	for i := 0; i < 10; i++ {
		busca := tree.Search(i)
		fmt.Println("procurando o ", i, "achou? ", busca)
	}
	fmt.Println()
}

func FullArrayRandom(tam int, intervalo int) []int {
	vetor := make([]int, tam)

	for i := 0; i < tam; i++ {
		vetor[i] = rand.Intn(intervalo)
	}
	return vetor
}

// func FullArraySemRepet(tam int) []int{

// }

func FullArrayCres(tam int) []int {
	vetor := make([]int, tam)

	for i := 0; i < tam; i++ {
		vetor[i] = i
	}
	return vetor
}

func FullArrayDecres(tam int) []int {
	vetor := make([]int, tam)

	for i := 0; i < tam; i++ {
		vetor[i] = (tam - 1) - i
	}
	return vetor
}*/