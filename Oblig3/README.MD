# Obligatorisk oppgave 3 #

#### Gruppemedlemmer ####
* Knut Andreas Aas 
* Knut Henæs 
* Sondre Lømsland
* Vemund H. Jernsletten


## [Oppgave 1](https://github.com/knuthenaes/Bitmen/tree/master/Oblig3/src/Oppgave%201%20og%20Oppgave%202) ##

*a)* Sett opp en webserver som lytter til port 8080.

*b)* Når klienter aksesserer webserveren på path lik "/" skal klienten få tilbake strengen "Hello, client".
Strengen skal vises i nettleseren.


## [Oppgave 2](https://github.com/knuthenaes/Bitmen/tree/master/Oblig3/src/Oppgave%201%20og%20Oppgave%202) API kall og behandling av JSON og HTML templates ##

*a)* Presenter data på webserveren slik at den er leselig for mennesker når klienter aksesserer stiene /1, /2, /3, /4, /5.

Dataen som skal presenteres skal hentes fra fem ulike APIer, hvor alle returnerer data i JSON format. Dere velger selv hvilke datasett dere benytter. Det er denne dataen som skal presenteres på de ulike stiene på webserveren

*b)* Samtlige stier i oppgave 2 skal rendres til klienter ved bruk av Go templates.

 

## [Oppgave 3](https://github.com/knuthenaes/Bitmen/tree/master/Oblig3/src/Oppgave%203) UDP/TCP server og Internett standarder ##

Implementer et serverprogram i henhold til RFC 865. Serveren skal svare både på UDP og TCP.

FYI: Når serveren lytter etter tilkoblinger på port 17, så fungerer det ved hjelp av kommandoen "telnet localhost 17". Dette gjelder TCP, men når det kommer til UDP så må NETCAT lastes ned. Dette er fordi telnet kun fungerer til TCP, mens NETCAT fungerer til begge deler. Vi pratet med Martin Nenseth, han sa dette var godt nok.
