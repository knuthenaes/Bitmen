package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	go readInput(c)
	time.Sleep(5 * 1e2) // Anvendes for at koden ikke skal kjøre igjennom koden før brukeren ikke får tastet inn en verdi
	go addUp(c)
	time.Sleep(5 * 1e2)
}

func readInput(c chan int) {

	var x1 int
	var x2 int

	fmt.Println("Skriv inn det første taller du vil addere: ")
	fmt.Scan(&x1)
	fmt.Println("Skriv inn det andre tallet du vil addere: ")
	fmt.Scan(&x2)

	c <- x1 //sender data via Channel
	c <- x2

	res := <-c // mottar resultat fra Channel
	fmt.Println("sum: ", res)
}

func addUp(c chan int) {
	n1, n2 := <-c, <-c // mottar data fra readInput()
	res := (n1 + n2)

	c <- res // sender resultat til readInput()

}
