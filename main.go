package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gocarina/gocsv"
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
func writeLog(txt string) {
	f, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02 15:04:05")
	if _, err := f.WriteString(date + " -- " + "info" + " -- " + txt + "\n"); err != nil {
		log.Printf("%v -- Error -- %v", date, err)
	}
}

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

type MAScsvPlayers struct {
	sliceCVSply []CsvPlayer
}
type CsvPlayer struct {
	ID          int
	FrName      string
	LaName      string
	Email       string
	Tip         string
	Status      string
	Saldo       int
	Country     string
	PhoneNumber int64
}

func inStruct() {
	in, err := os.Open("players.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	players := []*CsvPlayer{}

	if err := gocsv.UnmarshalFile(in, &players); err != nil {
		panic(err)
	}
	for _, Ply := range players {
		fmt.Println("Hello, ", ply.ID)
	}
}

/*
func CsvInStruct() {
	log.Println("\n\n\n")
	in, err := os.Open("players.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	ply := []*CsvPlayers{}

	if err := gocsv.UnmarshalFile(in, &ply); err != nil {
		panic(err)
	} else {
		fmt.Printf("%v", ply)
		for _, p := range ply {
			s := strconv.Itoa(p.ID) + " " + p.FrName + " " + p.LaName + " " + p.Email + " " + p.Tip + " " + p.Status + " " + strconv.Itoa(p.Saldo) + " " + p.Country + " " + strconv.FormatInt(p.PhoneNumber, 10)
			writeLog(s)
		}
	}
}
*/
/*
func CsvInCode(filePath string) {
	// Load a csv file.
	f, _ := os.Open(filePath)

	// Create a new reader.
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		//fmt.Println(record)
		for value := range record {
			fmt.Printf("  %v\n", record[value])
		}
	}
}*/
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
	CsvInCode("players.csv")

}
