package menu

import (
	"fmt"
	"projekt2/myGraph"
	"projekt2/testSuite"
)

func MainMenu() {
	var choice int
	var graphIM myGraph.Graph
	var graphPL myGraph.Graph
	var graphDirected bool
	for {
		fmt.Println("Wybierz opcję:")
		fmt.Println("wygeneruj graf - 1")
		fmt.Println("pokaż graf - 2")
		fmt.Println("algorytm prima MST - 3")
		fmt.Println("algorytm kruskala MST - 4")
		fmt.Println("algorytm Dijkstry - 5")
		fmt.Println("algorytm Bellmana-Forda - 6")
		fmt.Println("zapisz graf do pliku - 7")
		fmt.Println("wczytaj graf z pliku - 8")
		fmt.Println("Uruchom testy - 9")
		fmt.Println("wyjście - 0")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			return
		}
		switch choice {
		case 1:
			fmt.Println("Podaj liczbę wierzchołków:")
			var vertices int
			_, err := fmt.Scanln(&vertices)
			if err != nil {
				return
			}
			fmt.Println("Podaj procent połączeń:")
			var percentageConnected int
			_, err = fmt.Scanln(&percentageConnected)
			if err != nil {
				return
			}
			fmt.Println("Czy graf ma być skierowany? (t/n)")
			var directed string
			_, err = fmt.Scanln(&directed)
			if err != nil {
				return
			}
			if directed == "t" {
				graphIM, graphPL = myGraph.GenerateRandomGraph(vertices, percentageConnected, true)
				graphDirected = true
			} else {
				graphIM, graphPL = myGraph.GenerateRandomGraph(vertices, percentageConnected, false)
				graphDirected = false
			}
		case 2:
			fmt.Println("Macierz incydencji:")
			fmt.Println(graphIM.ToString())
			fmt.Println("Lista poprzedników:")
			fmt.Println(graphPL.ToString())
		case 3:
			if graphDirected {
				break
			}
			fmt.Println("Podaj wierzchołek startowy:")
			var startVertex int
			_, err := fmt.Scanln(&startVertex)
			if err != nil {
				return
			}
			mstIM, timeIM := myGraph.Prim(graphPL, startVertex, true)
			mstPL, timePL := myGraph.Prim(graphPL, startVertex, false)
			fmt.Println("Macierz incydencji:")
			fmt.Println(mstIM.ToString())
			fmt.Println("Czas:", timeIM/1000000, "ms")
			fmt.Println("Lista poprzedników:")
			fmt.Println(mstPL.ToString())
			fmt.Println("Czas:", timePL/1000000, "ms")
		case 4:
			if graphDirected {
				break
			}
			mstIM, timeIM := myGraph.Kruskal(graphIM, true)
			mstPL, timePL := myGraph.Kruskal(graphPL, false)
			fmt.Println("Macierz incydencji:")
			fmt.Println(mstIM.ToString())
			fmt.Println("Czas:", timeIM/1000000, "ms")
			fmt.Println("Lista poprzedników:")
			fmt.Println(mstPL.ToString())
			fmt.Println("Czas:", timePL/1000000, "ms")
		case 5:
			if !graphDirected {
				break
			}
			fmt.Println("Podaj wierzchołek startowy:")
			var startVertex int
			_, err := fmt.Scanln(&startVertex)
			if err != nil {
				return
			}
			fmt.Println("Podaj wierzchołek końcowy:")
			var endVertex int
			_, err = fmt.Scanln(&endVertex)
			if err != nil {
				return
			}
			pathIM, timeIM := myGraph.Dijkstra(graphIM, startVertex, endVertex)
			pathPL, timePL := myGraph.Dijkstra(graphPL, startVertex, endVertex)
			fmt.Println("Macierz incydencji:")
			fmt.Println(pathIM.ToString())
			fmt.Println("Czas:", timeIM/1000000, "ms")
			fmt.Println("Lista poprzedników:")
			fmt.Println(pathPL.ToString())
			fmt.Println("Czas:", timePL/1000000, "ms")
		case 6:
			if !graphDirected {
				break
			}
			fmt.Println("Podaj wierzchołek startowy:")
			var startVertex int
			_, err := fmt.Scanln(&startVertex)
			if err != nil {
				return
			}
			fmt.Println("Podaj wierzchołek końcowy:")
			var endVertex int
			_, err = fmt.Scanln(&endVertex)
			if err != nil {
				return
			}
			pathIM, timeIM := myGraph.BellmanFord(graphIM, startVertex, endVertex)
			pathPL, timePL := myGraph.BellmanFord(graphPL, startVertex, endVertex)
			fmt.Println("Macierz incydencji:")
			fmt.Println(pathIM.ToString())
			fmt.Println("Czas:", timeIM/1000000, "ms")
			fmt.Println("Lista poprzedników:")
			fmt.Println(pathPL.ToString())
			fmt.Println("Czas:", timePL/1000000, "ms")
		case 7:
			fmt.Println("Podaj nazwę pliku:")
			var filename string
			_, err := fmt.Scanln(&filename)
			if err != nil {
				return
			}
			err = myGraph.SaveToFile(graphIM, filename)
			if err != nil {
				fmt.Println("Błąd zapisu do pliku:", err)
			}
		case 8:
			fmt.Println("Podaj nazwę pliku:")
			var filename string
			_, err := fmt.Scanln(&filename)
			if err != nil {
				return
			}
			fmt.Println("Czy graf jest skierowany? (t/n)")
			var directed string
			_, err = fmt.Scanln(&directed)
			if err != nil {
				return
			}
			if directed == "t" {
				graphIM, err = myGraph.ReadFromFile(filename, true, true)
				graphPL, err = myGraph.ReadFromFile(filename, true, false)
			} else {
				graphIM, err = myGraph.ReadFromFile(filename, false, true)
				graphPL, err = myGraph.ReadFromFile(filename, false, false)
			}
			if err != nil {
				fmt.Println("Błąd odczytu z pliku:", err)
			}
		case 9:
			testSuite.RunFullTests()
		case 0:
			return
		}
	}
	return
}
