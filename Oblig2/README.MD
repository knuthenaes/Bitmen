Oppgave 1
a) Skriv et Golang program med navn fileinfo.go

Programmet skal returnere detaljert informasjon om en fil i følgende format:

 Information about file <filnavn>:

Size: X in bytes, KB, MB and GB

 Is/Is not a directory

 Is/Is not a regular file

 Has Unix permission bits: -rwxrwxrwx

 Is/Is not append only

 Is/Is not a device file

 

 Is/Is not a symbolic link

b) Lag en kjørbart program (build) av fileinfo.go som kan utføres fra kommandolinje. Programmet skal ta et argument, som skal være et filnavn.

Kommandoet for å utføre programmet i terminal skal se slik ut: fileinfo <filnavn>

Stien til filnavnet kan være enten relativ eller absolutt.


Oppgave 2
a) Skriv programmer med navn filecount.go som skal lese en tekst-fil, skrive ut totalt antall linjer og de fem “runes” (bokstaver/tegn), som forekommer mest i filen samt hvor mange ganger de telles i den medfølgende filen text.txt. Som forrige oppgave skal filen bygges til et kjørbart program og kunne utføres i terminal med kommandoen: filecount  <filnavn>

Resultatet skal være i følgende format:

Information about <filnavn>:

Number of lines in file:  X

Most common runes:

Rune: X , Counts: X
Rune: X , Counts: X
Rune: X , Counts: X
Rune: X , Counts: X
Rune: X , Counts: X

Oppgave 3
a) Skriv et program med navn addup.go bestående av to go funksjoner og kommunisere med hverandre gjennom channels

 Funksjon A:

Lag en go funksjon som leser inn to tall fra terminal (stdin) og sender disse to verdiene til en annen go funksjon.

 Funksjon B:

Mottakende go funksjon skal legge sammen de to tallene og sende resultatet tilbake til funksjon A som skriver ut resultatet til terminal (stdout).

 

b) Skriv to programmer (addtofile.go og sumfromfile.go) som kommuniserer med hverandre gjennom en fil.

 addtofile.go: Leser inn to tall og skriver disse til en fil.

 sumfromfile.go: Leser tall fra fil laget av program A og adderer de to tallene. Summen skrives tilbake til samme fil. Program A leser resultatet fra fil og skriver ut summen til stdout.

 
 c) Beskriv og implementer feilhåndtering på all I/O i oppgave a) og b)
 
d) Implementer håndtering av SIGINT i programmene i a) og b); programmene skal skrive ut en avslutningsmelding dersom de mottar SIGINT før de fullfører naturlig.

e) Bygg filene som i oppgave 1 og 2 til kjørbare filer på ditt operativsystem og legg dem ved i /bin mappen.
