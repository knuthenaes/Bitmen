# Systemarkitektur #

Vår webserver henter data i nåtid fra ulike værmeldings-API’er via openweathermap.org. Vi har begrenset værmeldingene til Trondheim, Stavanger og Oslo. Når tjenesten tas i bruk, vil brukeren se resultatet av dataen som webserveren henter fra de ulike API’ene i en nettleser slik at det blir lesbart. Dette fungerer på følgende måte:

![hierarki](https://user-images.githubusercontent.com/35768318/39866180-da325cf4-544f-11e8-84a2-f20258ee5101.png)  

Hierarkiet viser at hver gang en API hentes vil dette vises i webserveren i form av reelle tall i et redesignet miljø tilpasset et brukergrensesnitt. For å forbedre systemets interface henter nettsiden widgets. Når programmet kjøres starter det nederst i hierarkiet ved å hente data fra de ulike API'ene, så hentes widgets og deretter kommer resultat opp i web serveren vår. Dette gjøres ved bruk av en lokal server (localhost:8080).  
