# Obligatorisk oppgave 4 #

#### Gruppemedlemmer ####
* Knut Andreas Aas 
* Knut Henæs 
* Sondre Lømsland
* Vemund H. Jernsletten

## Oppgave ##
Oppgaven er et systemutviklingsprosjekt. Oppgaven skal inneholde en systemspesifikasjon som beskriver samfunnets og/eller brukeres nytteverdi av applikasjonen. I tillegg skal dere utrede, velge og dokumentere systemets arkitektur.  Systemarkitekturen viser til alle nodene, som er involvert i systemet og beskriver deres funksjon.  

Eksempler på prosjekt:

Finne “alt” om et spesifikt tema:
* Nyheter
* Twittermeldinger
* Facebook profiler

Finne spesifikke objekter basert på lokasjon og levere en liste / informasjon til brukeren
følge med på “realtids-data” fra noen databaser på WWW; levere en værmelding fra flere værdatabaser
levere produktinformasjon
Dere velger altså selv hva applikasjonen skal gjøre, men plattformen den utvikles til skal være web. Dere skal ved hjelp av pakker fra oblig 3 eller andre aktuelle pakker, visualisere applikasjonen i et brukergrensesnitt.

Programmet skal hente data fra relevante APIer som skal benyttes til noe fornuftig. (Eksempelvis vil værdata være interessant i en værapplikasjon)  Denne dataen, som i oblig 3, trenger ikke å lagres på en fysisk disk, og det holder derfor å kun lagre den i RAM.

Basert på verdien på dataen dere mottar fra APIer skal dere generere meningsfylte tilbakemeldinger til sluttbrukeren. Et realistisk eksempel på dette kan være at dersom temperaturen er over 30 grader i Kristiansand får brukeren følgende respons fra serveren: «Finn frem solbriller, i dag blir det varmt!»

Programmet deres skal vise meningsfulle data til brukeren, dvs. ikke data i rå JSON-format, men strukturerte data, som “Location: Kristiansand, Temperature high: 5 degrees Celsius, Temperature low: 0 degrees Celsius” osv. Formatet kan dere velge selv, så lenge det er brukbart. Med brukbart menes det blant annet at brukernes opplevelse av å bruke systemet oppfattes som intuitivt og enkelt.

Fordi dette er et prosjekt forventes det også at applikasjonen har flere eller større funksjonaliteter enn å bare hente og returnere en spesifikk type data uten mål og mening. Eksempler på dette:

Parkeringsapplikasjon som henter live data og lager et kart i GUI med markører som representerer parkeringsplasser.
Værapplikasjon som returner værdato for en oppgitt data fra brukeren, og avhengig av verdiene på dataen foreslår hvilke klær brukeren burde benytte.

I tillegg skal dere utvikle enhetstester for å kvalitetssikre at kildekoden oppfører seg pålitelig under forventede og uforventede betingelser. Det er fordelaktig at samtlige enheter har en eller flere enhetstester. Defensiv programmering skal anvendes der det er fornuftig, for å øke robustheten på applikasjonen.

Oppgaven vil bli bedømt ut fra løsning i henhold til oppgavetekst og anvendelse av konsepter og teori fra forelesninger og tidligere innleveringer; f.eks. feilhåndtering og korrekt håndtering av prosessignalring, beskrivelse av hardware/software og nettverksegenskaper.
