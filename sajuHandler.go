package urmy_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"time"
)

type SajuTable struct {
	Chungan []Chungan
	Jiji    []Jiji
	Manse   []YearGanji
	Julgys  []YearJulgy
	SaeUn   []SaeUn
}

type Chungan struct {
	Id         int
	Title      string
	Properties Property_Chun
}

type Property_Chun struct {
	Umyang         int
	Prop           string
	Sibsung        string
	IpMyo          IpMyo
	Unsung_Me      Unsung_Me
	Unsung_by_Jiji Unsung_by_Jiji
}

type Unsung_Me struct {
	level        int
	Unsung_title string
}

type Unsung_by_Jiji struct {
	level        int
	Unsung_title string
}

type Chungan_Unsung struct {
	Title      string
	Properties []Property_Unsung
}

type Property_Unsung struct {
	Level     int
	Jiji_char string
	Prop      string
}

type Jiji struct { //HP넣기
	Id         int
	Title      string
	Properties Property_Ji
}

type Property_Ji struct {
	Umyang     int
	Prop       string
	Sibsung    string
	Hae        int
	Go         int
	Ji         int
	ChangGo    ChangGo
	Jijanggans []Jijanggan
}

type ChangGo struct {
	Exist     bool
	InSungGo  int
	YangInGo  int
	JaeGo     int
	GwanGo    int
	SikSangGo int
}

type Jijanggan struct {
	Chungan_char Chungan_Jijanggan
	Value_len    int
}

type Chungan_Jijanggan struct {
	Id         int
	Title      string
	Properties Property_Chun_Jijanggan
}

type Property_Chun_Jijanggan struct {
	Umyang int
	Prop   string
}

type YearGanji struct {
	WhichYear     int
	Chungan_Title string
	Jiji_Title    string
	MonthGanji    []MonthGanji
}

type MonthGanji struct {
	WhichMonth    int
	Chungan_Title string
	Jiji_Title    string
	Day_Ganji     []Day_Ganji
}

type Day_Ganji struct {
	WhichDay      int
	Chungan_Title string
	Jiji_Title    string
	Time_Ganji    []Time_Ganji
}

type Time_Ganji struct {
	Chungan_Title string
	Jiji_Title    string
}

type SaJuPalJa struct {
	YearChun  string
	YearJi    string
	MonthChun string
	MonthJi   string
	DayChun   string
	DayJi     string
	TimeChun  string
	TimeJi    string
	DaeUnChun string
	DaeUnJi   string
	SaeUnChun string
	SaeUnJi   string
}

type YearJulgy struct {
	Year        int
	MonthJulgys []MonthJulgy
}

type MonthJulgy struct {
	Month     int
	DayJulgys []DayJulgy
}

type DayJulgy struct {
	Day   int
	Title string
}

type SaeUn struct {
	Year        int
	SaeUnGanjis SaeUnGanji
}

type SaeUnGanji struct {
	Chun string
	Ji   string
}

type SaJuHandler interface {
	ExtractSaju(saju *Saju) *SaJuPalJa
	ExtractDaeUnSaeUn(saju *Saju, palja *SaJuPalJa) *SaJuPalJa
	GetSajuTable() *SajuTable
}

func (s *SajuTable) GetSajuTable() *SajuTable {
	var sajutable SajuTable
	sajutable.Chungan = s.Chungan
	sajutable.Jiji = s.Jiji
	sajutable.Julgys = s.Julgys
	sajutable.Manse = s.Manse
	return &sajutable
}

func (s *SajuTable) ExtractSaju(saju *Saju) *SaJuPalJa {
	var t SaJuPalJa
	tbirth := time.Date(saju.Year, time.Month(saju.Month), saju.Day, 0, 0, 0, 0, time.UTC)
	yjulgy := time.Date(s.Julgys[saju.Year-s.Manse[0].WhichYear].Year, time.Month(s.Julgys[saju.Year-s.Manse[0].WhichYear].MonthJulgys[1].Month), s.Julgys[saju.Year-s.Manse[0].WhichYear].MonthJulgys[1].DayJulgys[0].Day, 0, 0, 0, 0, time.UTC)
	mjulgy := time.Date(s.Julgys[saju.Year-s.Manse[0].WhichYear].Year, time.Month(s.Julgys[saju.Year-s.Manse[0].WhichYear].MonthJulgys[saju.Month-1].Month), s.Julgys[saju.Year-s.Manse[0].WhichYear].MonthJulgys[saju.Month-1].DayJulgys[0].Day, 0, 0, 0, 0, time.UTC)

	if tbirth.Before(yjulgy) {
		t.YearChun = s.Manse[saju.Year-s.Manse[0].WhichYear-1].Chungan_Title
		t.YearJi = s.Manse[saju.Year-s.Manse[0].WhichYear-1].Jiji_Title
	} else {
		t.YearChun = s.Manse[saju.Year-s.Manse[0].WhichYear].Chungan_Title
		t.YearJi = s.Manse[saju.Year-s.Manse[0].WhichYear].Jiji_Title
	}

	if tbirth.Before(mjulgy) {
		if saju.Month == 1 {
			t.MonthChun = s.Manse[saju.Year-s.Manse[0].WhichYear-1].MonthGanji[11].Chungan_Title
			t.MonthJi = s.Manse[saju.Year-s.Manse[0].WhichYear-1].MonthGanji[11].Jiji_Title
		} else {
			// 월 배열 -1 , 절기 전 -1
			t.MonthChun = s.Manse[saju.Year-s.Manse[0].WhichYear].MonthGanji[saju.Month-2].Chungan_Title
			t.MonthJi = s.Manse[saju.Year-s.Manse[0].WhichYear].MonthGanji[saju.Month-2].Jiji_Title
		}
	} else {
		t.MonthChun = s.Manse[saju.Year-s.Manse[0].WhichYear].MonthGanji[saju.Month-1].Chungan_Title
		t.MonthJi = s.Manse[saju.Year-s.Manse[0].WhichYear].MonthGanji[saju.Month-1].Jiji_Title
	}

	t.DayChun = s.Manse[saju.Year-s.Manse[0].WhichYear].MonthGanji[saju.Month-1].Day_Ganji[saju.Day-1].Chungan_Title
	t.DayJi = s.Manse[saju.Year-s.Manse[0].WhichYear].MonthGanji[saju.Month-1].Day_Ganji[saju.Day-1].Jiji_Title

	indexTime := inTimeSpan(saju.Time)
	t.TimeChun = s.Manse[saju.Year-s.Manse[0].WhichYear].MonthGanji[saju.Month-1].Day_Ganji[saju.Day-1].Time_Ganji[indexTime].Chungan_Title
	t.TimeJi = s.Manse[saju.Year-s.Manse[0].WhichYear].MonthGanji[saju.Month-1].Day_Ganji[saju.Day-1].Time_Ganji[indexTime].Jiji_Title
	return &t
}

func (s *SajuTable) ExtractDaeUnSaeUn(saju *Saju, palja *SaJuPalJa) *SaJuPalJa {

	var index int
	var daeunindex int
	currentdaeunindex := 0
	age := time.Now().Year() - saju.Year
	palja.SaeUnChun = s.SaeUn[time.Now().Year()-1950].SaeUnGanjis.Chun
	palja.SaeUnJi = s.SaeUn[time.Now().Year()-1950].SaeUnGanjis.Ji

	for i := 0; i < 10; i++ {
		if s.Chungan[i].Title == palja.YearChun {
			index = i
			break
		}
	}
	switch saju.Gender {
	case true:
		switch s.Chungan[index].Properties.Umyang { //1991807
		case 1:
			daeunindex = s.forwardDaeunIndex(saju, palja)
		case 0:
			daeunindex = s.reverseDaeunIndex(saju, palja)
		}

	case false:
		switch s.Chungan[index].Properties.Umyang { //1991807
		case 1:
			daeunindex = s.reverseDaeunIndex(saju, palja)
		case 0:
			daeunindex = s.forwardDaeunIndex(saju, palja)
		}
	}

	for i := 0; (age - (i*10 + daeunindex)) >= 0; i++ {
		currentdaeunindex += 1
	}
	fmt.Println(index)
	switch saju.Gender {
	case true:
		switch s.Chungan[index].Properties.Umyang {
		case 1:
			palja = s.forwardDaeunGanji(currentdaeunindex, saju, palja)
		case 0:
			palja = s.reverseDaeunGanji(currentdaeunindex, saju, palja)
		}
	case false:
		switch s.Chungan[index].Properties.Umyang {
		case 1:
			palja = s.reverseDaeunGanji(currentdaeunindex, saju, palja)
		case 0:
			palja = s.forwardDaeunGanji(currentdaeunindex, saju, palja)
		}
	}

	return palja
}

func (s *SajuTable) forwardDaeunIndex(saju *Saju, palja *SaJuPalJa) int {
	var daeunindex int
	tjulgy := time.Date(saju.Year, time.Month(saju.Month), s.Julgys[saju.Year-1950].MonthJulgys[saju.Month-1].DayJulgys[0].Day, 0, 0, 0, 0, time.UTC)
	tbirth := time.Date(saju.Year, time.Month(saju.Month), saju.Day, 0, 0, 0, 0, time.UTC)

	if tbirth.After(tjulgy) {
		if saju.Month+1 == 13 {
			tjulgy = time.Date(saju.Year+1, time.Month(saju.Month-11), s.Julgys[saju.Year-1949].MonthJulgys[0].DayJulgys[0].Day, 0, 0, 0, 0, time.UTC)
		} else {
			tjulgy = time.Date(saju.Year, time.Month(saju.Month+1), s.Julgys[saju.Year-1950].MonthJulgys[saju.Month].DayJulgys[0].Day, 0, 0, 0, 0, time.UTC)
		}
	} else if tjulgy == tbirth {
		daeunindex = 1
		return daeunindex
	}
	e := tjulgy.Sub(tbirth).Hours() / 72
	daeunindex = int(math.Round(e))
	return daeunindex
}

func (s *SajuTable) reverseDaeunIndex(saju *Saju, palja *SaJuPalJa) int {
	var daeunindex int
	tjulgy := time.Date(saju.Year, time.Month(saju.Month), s.Julgys[saju.Year-1950].MonthJulgys[saju.Month-1].DayJulgys[0].Day, 0, 0, 0, 0, time.UTC)
	tbirth := time.Date(saju.Year, time.Month(saju.Month), saju.Day, 0, 0, 0, 0, time.UTC)
	if tbirth.Before(tjulgy) {
		if saju.Month-1 == 0 {
			tjulgy = time.Date(saju.Year-1, time.Month(12), s.Julgys[saju.Year-1950].MonthJulgys[11].DayJulgys[0].Day, 0, 0, 0, 0, time.UTC)
		} else {
			tjulgy = time.Date(saju.Year, time.Month(saju.Month-1), s.Julgys[saju.Year-1950].MonthJulgys[saju.Month-2].DayJulgys[0].Day, 0, 0, 0, 0, time.UTC)
		}
	} else if tjulgy == tbirth {
		daeunindex = 1
		return daeunindex
	}
	e := tbirth.Sub(tjulgy).Hours() / 72
	daeunindex = int(math.Round(e))
	return daeunindex
}

func (s *SajuTable) forwardDaeunGanji(currentdaeunindex int, saju *Saju, palja *SaJuPalJa) *SaJuPalJa {
	for i := 0; i < 10; i++ {
		if palja.MonthChun == s.Chungan[i].Title {
			if i+currentdaeunindex < 10 {
				palja.DaeUnChun = s.Chungan[i+currentdaeunindex].Title
			} else {
				palja.DaeUnChun = s.Chungan[i+currentdaeunindex-10].Title
			}
		}
	}
	for i := 0; i < 12; i++ {
		if palja.MonthJi == s.Jiji[i].Title {
			if i+currentdaeunindex < 12 {
				palja.DaeUnJi = s.Jiji[i+currentdaeunindex].Title
			} else {
				palja.DaeUnJi = s.Jiji[i+currentdaeunindex-12].Title
			}
		}
	}
	return palja
}

func (s *SajuTable) reverseDaeunGanji(currentdaeunindex int, saju *Saju, palja *SaJuPalJa) *SaJuPalJa {
	for i := 9; i >= 0; i-- {
		if palja.MonthChun == s.Chungan[i].Title {
			if i-currentdaeunindex < 0 {
				palja.DaeUnChun = s.Chungan[i-currentdaeunindex+10].Title
			} else {
				palja.DaeUnChun = s.Chungan[i-currentdaeunindex].Title
			}
		}
	}
	for i := 11; i >= 0; i-- {
		if palja.MonthJi == s.Jiji[i].Title {
			if i-currentdaeunindex < 0 {
				palja.DaeUnJi = s.Jiji[i-currentdaeunindex+12].Title
			} else {
				palja.DaeUnJi = s.Jiji[i-currentdaeunindex].Title
			}
		}
	}
	return palja
}

func inTimeSpan(check string) int {
	//valueofChar := [12]rune{'자', '축', '인', '묘', '진', '사', '오', '미', '신', '유', '술', '해'}
	valueofChar := [13]int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	data := []struct {
		start string
		end   string
	}{
		{"23:29", "24:00"},
		{"00:00", "01:30"},
		{"01:29", "03:30"},
		{"03:29", "05:30"},
		{"05:29", "07:30"},
		{"07:29", "09:30"},
		{"09:29", "11:30"},
		{"11:29", "13:30"},
		{"13:29", "15:30"},
		{"15:29", "17:30"},
		{"17:29", "19:30"},
		{"19:29", "21:30"},
		{"21:29", "23:30"},
	}
	if check == "00:00" {
		return valueofChar[0]
	}

	newLayout := "15:04"
	checked, _ := time.Parse(newLayout, check)

	a := 0
	for _, t := range data {
		start, _ := time.Parse(newLayout, t.start)
		end, _ := time.Parse(newLayout, t.end)

		if checked.After(start) && checked.Before(end) {
			break
		}
		a++
	}
	return valueofChar[a]
}

func NewSaJuHandler() SaJuHandler {
	return newSaJuHandler()
}

func newSaJuHandler() SaJuHandler {

	var data1 []Chungan
	var data2 []Jiji
	var ganjis = make([]YearGanji, 100)
	var julgy = make([]YearJulgy, 74)
	var SaeUn = make([]SaeUn, 100)

	b, err := ioutil.ReadFile("/etc/config/chungan.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	c, err := ioutil.ReadFile("/etc/config/jiji.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	json.Unmarshal(b, &data1)
	json.Unmarshal(c, &data2)

	for i := 0; i < 100; i++ {
		d, err := ioutil.ReadFile("/etc/config/manse/manse" + strconv.Itoa(1950+i) + ".json")
		if err != nil {
			panic(err)
		}
		json.Unmarshal(d, &ganjis[i])
	}
	fmt.Println("Manse Done")

	for i := 0; i < 74; i++ {
		f, err := ioutil.ReadFile("/etc/config/julgy/julgy" + strconv.Itoa(1950+i) + ".json")
		if err != nil {
			panic(err)
		}
		json.Unmarshal(f, &julgy[i])
	}
	fmt.Println("julgy Done")

	g, err := ioutil.ReadFile("/etc/config/saeun/saeun.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(g, &SaeUn)
	fmt.Println("SaeUn Done")

	return &SajuTable{
		Chungan: data1,
		Jiji:    data2,
		Manse:   ganjis,
		Julgys:  julgy,
		SaeUn:   SaeUn,
	}
}
