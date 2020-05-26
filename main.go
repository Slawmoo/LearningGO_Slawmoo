package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type masterLige struct {
	Lige []Lige `json:"lige"`
}

func (l *masterLige) StringMASLige() string {
	str := ""
	for _, lj := range l.Lige {
		str += fmt.Sprintf("Lige: [%v]", lj.StringLIGE())
	}
	return fmt.Sprintf("lige: [%v]", str)
}

type Lige struct {
	Naziv   string    `json:"naziv"`
	Razrade []Razrade `json:"razrade"`
}

func (l *Lige) StringLIGE() string {
	str := ""
	for _, r := range l.Razrade {
		str += fmt.Sprintf("Razrade(glavno): [%v]", r.StringRAZRADE())
	}
	return fmt.Sprintf("naziv: %v, razrade [%v]", l.Naziv, str)
}

type Razrade struct {
	Tipovi []Tipovi `json:"tipovi"`
	Ponude []int    `json:"ponude"`
}

func (l *Razrade) StringRAZRADE() string {
	return fmt.Sprintf("")
}

type Tipovi struct {
	Naziv string `json:"naziv"`
}

type masterPonude struct {
	id          int
	slavePonude []slavePonude
}
type slavePonude struct {
	Broj          string     `json:"broj"`
	id            int        `json:"id"`
	Naziv         string     `json:"naziv"`
	Vrijeme       time.Time  `json:"vrijeme"`
	Tecajevi      []Tecajevi `json:"tecajevi"`
	TvKanal       string     `json:"tv_kanal,omitempty"`
	ImaStatistiku bool       `json:"ima_statistiku,omitempty"`
}

func (s *slavePonude) NazivPonude() string { // samo dodavanje imena ponude
	str := ""
	str += fmt.Sprintf("Ime ponude: [%v]", s.Naziv)
	return str
}
func (m *masterPonude) Ponude() string {
	var masPon masterPonude
	str := ""
	for _, f := range masPon.slavePonude {
		str += fmt.Sprintf("Naziv ponude: [%v] ,Tečaj: [%v]", f.NazivPonude(), f.StringPONUDE()) // vanjski tecajevi
	}
	return str
}
func (s *slavePonude) StringPONUDE() string {
	var slvPon slavePonude
	str := ""
	for _, i := range s.Tecajevi {
		str += fmt.Sprintf("%v. tecaj: [%v]", i, slvPon.StringTECAJI()) //unutarnji tecajevi
	}
	return fmt.Sprintf("2. po redu tečajevi: [%v]", str)
}

type Tecajevi struct {
	Tecaj float64 `json:"tecaj"`
	Naziv string  `json:"naziv"`
}

func (s *slavePonude) StringTECAJI() string {
	for f, j := range s.Tecajevi {
		if f == 1 {
			return fmt.Sprintf("Ime 2. tecaja: [%v], 2-tečaj (vrijednost) : [%v]", j.Naziv, j.Tecaj) // vadenje tecaja
		}
	}
	return ""
}

// zelim ispisat sve nazive ponude, 2. po redu - tecaj(broj) i naziv tecaja
var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return err
	}
	defer r.Body.Close()

	//	fmt.Println(string(body))
	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	var dLige masterLige
	var dPonude masterPonude

	fmt.Println("ziv sams")
	err := getJSON("https://www.aeternus.hr/go/lige.json", &dLige)
	if err != nil {
		fmt.Println(err)
	}
	err1 := getJSON("https://www.aeternus.hr/go/ponude.json", &dPonude)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("main LIGE:\n")
	fmt.Println(dLige)
	fmt.Println("stringer za LIGE:\n")
	fmt.Println(dLige.StringMASLige() + "\n\n\n\n")
	fmt.Println("main PONUDE:\n")
	fmt.Println(dPonude)
	fmt.Println("stringer za PONUDE:\n")
	fmt.Println(dPonude.Ponude())
}
