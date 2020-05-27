
Zadatak 002 - Rad s fileovima (2020-05-27)

Napraviti log file, tj text file imena log.log u koji se ispisuje sve sto aplikacija treba ispisati. Standardizirati svaki zapis s 3 elementa da izgleda ovako:

{DATETIME} -- {LEVEL} -- {log zapis}

Gdje je DATETIME  datum i vrijeme u formatu: YYYY-MM-DD hh:mm:ss
LEVEL - tip loga, npr: Error, info, ...


Iz filea players.csv ucitati podatke u memoriju (strukturu) i ispisati ih u log ponovno koristeći stringere.



Zadatak 001

Osnova aplikacije

Napravi aplikaciju koja će pri startu učitati 2 JSON file-a sa sljedećih URL-ova:
https://www.aeternus.hr/go/lige.json
https://www.aeternus.hr/go/ponude.json
JSON podaci (lige.json i ponude.json) su međusobno povezani preko ID-eva ponude. Podaci u ovom trenutku još uvijek mogu biti spremljeni u memoriji aplikacije.

Na kraju, koristeci String metode (stringere) za pojedine strukture, ispisati dobiveno u log (dovoljno koristiti fmt.Printf, fmt.Println)
