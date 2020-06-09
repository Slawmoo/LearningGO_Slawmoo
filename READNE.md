Zadatak 003 - Validacije i manipulacija stringom

Extendati strukturu Players dodavajuci polje: smsNotificationValidCroatianNumber.

U to polje, pri ucitavanju podataka upisati broj telefona u internacionalnom obliku bez vodecih nula (primjer: 38531631431.

Extendati i stringera da ispisuje ovaj podataka.

Napraviti validator e-maila koji pri citanju podataka prvo provjeri da li je email validan. Ako nije, ne ucitava se u polje igraci.


Pomoc:
Broj telefona se sastoji od:
leading data (moze biti 00 ili +)
international country prefix (hrvatska: 385, Serbija 381, USA 1, ... Slobodno guglati)
operator prefix:
	- u hr mobilni operatori imaju dvoznamenkasti prefix u rasponu od 90-99
	- u slucaju "national formata mobilnog broja, moblinom prefixu predstoji vodeca 0: primjer: 0981967452)
subscriber number: unique number subscribera: Min 6 digits.

Takoder, broj telefona nema vise od 15 digita u internacionalnom obliku s vodecim nulama


Glossary:
https://www.quora.com/What-is-maximum-and-minimum-length-of-any-mobile-number-across-the-world

https://en.wikipedia.org/wiki/Email_address#Syntax 

--------------------------------------------------------------------
====================================================================
--------------------------------------------------------------------

Zadatak 002 - Rad s fileovima (2020-05-27)

Napraviti log file, tj text file imena log.log u koji se ispisuje sve sto aplikacija treba ispisati. Standardizirati svaki zapis s 3 elementa da izgleda ovako:

{DATETIME} -- {LEVEL} -- {log zapis}

Gdje je DATETIME  datum i vrijeme u formatu: YYYY-MM-DD hh:mm:ss
LEVEL - tip loga, npr: Error, info, ...


Iz filea players.csv ucitati podatke u memoriju (strukturu) i ispisati ih u log ponovno koristeći stringere.

--------------------------------------------------------------------
====================================================================
--------------------------------------------------------------------

Zadatak 001

Osnova aplikacije

Napravi aplikaciju koja će pri startu učitati 2 JSON file-a sa sljedećih URL-ova:
https://www.aeternus.hr/go/lige.json
https://www.aeternus.hr/go/ponude.json
JSON podaci (lige.json i ponude.json) su međusobno povezani preko ID-eva ponude. Podaci u ovom trenutku još uvijek mogu biti spremljeni u memoriji aplikacije.

Na kraju, koristeci String metode (stringere) za pojedine strukture, ispisati dobiveno u log (dovoljno koristiti fmt.Printf, fmt.Println)
