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
		str += fmt.Sprintf("Razrade: [%v]", r.StringRAZRADE())
	}
	return fmt.Sprintf("naziv: %v, razrade [%v]", l.Naziv, str)
}

type Razrade struct {
	Tipovi []Tipovi `json:"tipovi"`
	Ponude []int    `json:"ponude"`
}

func (lraz *Razrade) StringRAZRADE() string {
	return fmt.Sprintf("")
}

type Tipovi struct {
	Naziv string `json:"naziv"`
}

type masterPonude []struct {
	Broj          string     `json:"broj"`
	id            int        `json:"id"`
	Naziv         string     `json:"naziv"`
	Vrijeme       time.Time  `json:"vrijeme"`
	Tecajevi      []Tecajevi `json:"tecajevi"`
	TvKanal       string     `json:"tv_kanal,omitempty"`
	ImaStatistiku bool       `json:"ima_statistiku,omitempty"`
}

/*func (m *[]masterPonude) masPonude() string {

	return fmt.Sprintf("Ponude: [%v]", str)
}*/

func (p *[]masterPonude) StringPONUDE() string {
	str := ""
	for _, f := range p.Tecajevi {
		str += fmt.Sprintf("Tečajevi: ", f.StringTECAJI())
	}
	return fmt.Sprintf("Ponude: [%v]", str)
}

type Tecajevi struct {
	Tecaj float64 `json:"tecaj"`
	Naziv string  `json:"naziv"`
}

func (t *Tecajevi) StringTECAJI() string {
	str := ""
	str += fmt.Sprintf("Tečaj: %v", t.Tecaj)
	return fmt.Sprintf("Ime tecaja: [%v], razrade [%v]", t.Naziv, str)
}

func main() {
	var dLige masterLige
	var dPonude masterPonude

	fmt.Println("ziv sams")
	err := getJSON("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json", &dLige)
	if err != nil {
		fmt.Println(err)
	}
	err1 := getJSON("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json", &dPonude)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("main LIGE:\n")
	fmt.Println(dLige)
	fmt.Println("stringer za LIGE:\n")
	fmt.Println(dLige.StringMASLige())
	fmt.Println("main PONUDE:\n")
	fmt.Println(dPonude)
	fmt.Println("stringer za PONUDE:\n")
	fmt.Println(dPonude.StringPONUDE())
}

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

/*
	getJSON("https://www.dropbox.com/s/wr9vnmt5e1jhwkq/ponude.json?dl=0", dataPonude) // getting data and populating structs
	lige.json = https://www.dropbox.com/s/2kqweiiqf6nbhfz/lige.json?dl=0
 	ponude.json = https://www.dropbox.com/s/wr9vnmt5e1jhwkq/ponude.json?dl=0
	dataLige := new(GlLige)
		dataPonude := new(GlPonude) // structs in variabs


	   https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json
	   https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json
*/
