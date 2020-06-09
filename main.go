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
		str += fmt.Sprintf("Razrade(glavno): {{[%v]}}", r.StringRAZRADE())
	}
	return fmt.Sprintf(" {Naziv: %v, Razrade [%v]} ", l.Naziv, str)
}

type Razrade struct {
	Tipovi []Tipovi `json:"tipovi"`
	Ponude []int    `json:"ponude"`
}

func (l *Razrade) StringRAZRADE() string {
	str := ""
	for _, m := range l.Tipovi {
		str += fmt.Sprintf("Tip: [%v], ", m.Naziv)
	}
	for _, m := range l.Ponude {
		str += fmt.Sprintf("Ponuda: [%v]", m)
	}
	return str
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

func (m *masterPonude) svePonude() string {
	str := ""
	for _, f := range m.slavePonude {
		str += fmt.Sprintf("Broj: [%v], id: [%v], naziv: [%v], Vrijeme: [%v], Tecajevi: [%v],TvKanal: [%v], ImaStatistiku: [%v]", f.Broj, f.id, f.Naziv, f.Vrijeme, f.stringTecajevi(), f.TvKanal, f.ImaStatistiku) // vanjski tecajevi
	}
	return str
}
func (s *slavePonude) stringTecajevi() string {
	str := ""
	for _, i := range s.Tecajevi {
		str += fmt.Sprintf("[ Naziv tecaja: %v, tecaj: %v ]", i.Naziv, i.Tecaj)
	}
	return str
}

type Tecajevi struct {
	Tecaj float64 `json:"tecaj"`
	Naziv string  `json:"naziv"`
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

func main() {

	var dLige masterLige
	var dPonude []slavePonude
	var svePonude masterPonude

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
	fmt.Println("Sve LIGE (stringer):\n")
	fmt.Println(dLige.StringMASLige() + "\n\n")
	fmt.Println("main PONUDE:\n")
	fmt.Println(dPonude)
	fmt.Println("Sve PONUDE (stringer):\n\n\n")
	fmt.Println(svePonude.svePonude())
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
	f, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)
	plyLog(players)
	var csvPlayer SLCCsvPlayer
	fmt.Print("\n\n DOLJE DODIN\n\n")
	fmt.Println(csvPlayer.String())
}

type SLCCsvPlayer struct {
	CsvPlayer []CsvPlayer
}

type CsvPlayer struct {
	ID          string `json:id`
	FrName      string `json:"firstname"`
	LaName      string `json:"lastname"`
	Email       string `json:"email"`
	Tip         string `json:"tip"`
	Status      string `json:"status"`
	Saldo       string `json:"saldo"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phonenumber"`
}

func (p *SLCCsvPlayer) String() string {
	str := ""
	for _, r := range p.CsvPlayer {
		currentTime := time.Now()
		date := currentTime.Format("2006-01-02 15:04:05")
		str += date + "  --  info  --  ID: " + r.ID + ", firstname: " + r.FrName + ", lastname: " + r.LaName + ", email: " + r.Email + ", tip: " + r.Tip + ", status: " + r.Status + ", saldo: " + r.Saldo + ", country: " + r.Country + ", phonenumber: " + r.PhoneNumber + "\n"
	}
	return str
}

func plyLog(ply []CsvPlayer) {
	re := ""
	for _, p := range ply {
		s := p.ID + " " + p.FrName + " " + p.LaName + " " + p.Email + " " + p.Tip + " " + p.Status + " " + p.Saldo + " " + p.Country + " " + p.PhoneNumber
		currentTime := time.Now()
		date := currentTime.Format("2006-01-02 15:04:05")
		re += date + "  --  info  --  " + s + "\n"
		//treba ispisat u file pomocu fmt. paketa, sutra cemo
	}

}
func (p *SLCCsvPlayer) plyLogStringer() string {
	re := ""
	for _, p := range p.CsvPlayer {
		s := p.ID + " " + p.FrName + " " + p.LaName + " " + p.Email + " " + p.Tip + " " + p.Status + " " + p.Saldo + " " + p.Country + " " + p.PhoneNumber
		currentTime := time.Now()
		date := currentTime.Format("2006-01-02 15:04:05")
		re += date + "  --  info  --  " + s + "\n"
		//treba ispisat u file pomocu fmt. paketa, sutra cemo
	}
	return re
}

func errChk(msg string, err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(msg)
	}
}

/* treba rjesit sve ponude preko stringera ispis - eventualno pitat dodu
trea rjesiz ispis players u .log file
rjesit sta ne valja*/
