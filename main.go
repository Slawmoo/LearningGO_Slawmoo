package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

type SLCCsvPlayer struct {
	CsvPlayer []CsvPlayer
}

type CsvPlayer struct {
	ID          string `json:ID`
	FrName      string `json:"firstname"`
	LaName      string `json:"lastname"`
	Email       string `json:"email"`
	Tip         string `json:"tip"`
	Status      string `json:"status"`
	Saldo       string `json:"saldo"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phonenumber"`
}

func main() {
	/*
		var dLige masterLige
		var dPonude []slavePonude
		var posPonude masterPonude

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
		fmt.Println(posPonude.Ponude())
	*/
	csvFile, err1 := os.Open("players.csv")
	errChk("opening", err1)
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	var players []CsvPlayer
	for {
		line, err2 := reader.Read()
		if err2 == io.EOF {
			break
		} else if err2 != nil {
			errChk("reading", err2)
		}
		players = append(players, CsvPlayer{
			ID:          line[0],
			FrName:      line[1],
			LaName:      line[2],
			Email:       line[3],
			Tip:         line[4],
			Status:      line[5],
			Saldo:       line[6],
			Country:     line[7],
			PhoneNumber: line[8],
		})
	}
	//playersJson, err3 := json.Marshal(players)
	//errChk("out", err3)
	fmt.Println(players)
	fmt.Println("\n")
	f, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	log.SetOutput(f)
	plyLog(players)
	// plyLogStringer()
}

func plyLog(ply []CsvPlayer) {
	for _, p := range ply {
		s := p.ID + " " + p.FrName + " " + p.LaName + " " + p.Email + " " + p.Tip + " " + p.Status + " " + p.Saldo + " " + p.Country + " " + p.PhoneNumber
		currentTime := time.Now()
		date := currentTime.Format("2006-01-02 15:04:05")
		log.Printf("%v  --  info  --  %v\n", date, s)
	}
}
func (c *SLCCsvPlayer) plyLogStringer() {
	for _, p := range c.CsvPlayer {
		s := p.ID + " " + p.FrName + " " + p.LaName + " " + p.Email + " " + p.Tip + " " + p.Status + " " + p.Saldo + " " + p.Country + " " + p.PhoneNumber
		currentTime := time.Now()
		date := currentTime.Format("2006-01-02 15:04:05")
		log.Printf("%v  --  info  --  %v\n", date, s)
	}
}

func errChk(msg string, err error) {
	if err != nil {
		log.Fatal(err)
		fmt.Println(msg)
	}
}

// LOG TIME FORMAT
type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02 15:04:05") + string(bytes) + "\n")
}

//--------------------------------------
