package main

import "fmt"

func main() {
	var ledd1 int16
	var ledd2 int16

	fmt.Println("Skriv inn den første verdien du vil addere:")
	fmt.Scanln(&ledd1)
	/* Oppgave C - må behandle ikke-heltall og for store tall på et eller annet vis:
	if [error - vet enda ikke korleis en skal derivere dette]
	for [error] = 1 {
		fmt.Println("Hmm, det er fint om du begrenser tallverdien til et heltall (innenfor |32000|) - verdier utenfor int16 vil ikke kunne kalkuleres. Så la oss ta dette steget på nytt.\n\nSkriv inn første verdi du vil legge sammen:"
	fmt.Scanln(&ledd1)
	if !error, sett error til 0
	}
	 */
	fmt.Println("Skriv inn den andre verdien du vil addere:")
	fmt.Scanln(&ledd2)
	/* Oppgave C - må behandle ikke-heltall, for store tall på et eller annet vis:
	if [error - vet enda ikke korleis en skal derivere dette]
	for [error] = 1 {
		fmt.Println("Hmm, det er fint om du begrenser tallverdien til et heltall (innenfor |32000|) - verdier utenfor int16 vil ikke kunne kalkuleres. Så la oss ta dette steget på nytt.\n\nSkriv inn andre verdi du vil legge sammen:"
	fmt.Scanln(&ledd1)
	if !error, sett error til 0
	}
	 */

	fmt.Println("Summen av tallene er:" )
	fmt.Print(funkB(ledd1, ledd2))}

func funkB(x int16, y int16) int16 {
	return x + y
}