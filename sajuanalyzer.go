package urmy_handler

import (
	//	"container/ring"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	//	"reflect"
)

type sajuAnalyzer struct {
	Sibsung    []Sibsung
	Sib2Unsung []Chungan_Unsung
}

type SaJuAnalyzer interface {
	Evaluate_GoonbHab(host Person, opponent Person) (string, string, string, string)
	Find_GoongHab(*PersonSaju, *PersonSaju, []Chungan, []Jiji, []Sibsung, []Chungan_Unsung) (Person, Person)
	GetAnalyzerTable() *sajuAnalyzer
}

type Sibsung struct {
	Prop      string
	Comp_Prop []Compare_Prop
}

type Compare_Prop struct {
	Comp_Prop string
	Title     string
}

type Person struct {
	LoginID string
	Chun    []Chungan
	Ji      []Jiji
	Result  []Result_record
}

type Result_record struct {
	ChunGanHab  ChunGanHab
	ChunGanGeok ChunGanGeok
	Sibsung     SibsungResult
	YukHab      YukHab
	SamHab      SamHab
	BanHab      BanHab
	BangHab     BangHab
	AmHab       AmHab
	MyeongAmHab MyeongAmHab
	WonJin      WonJin
	GuiMoon     GuiMoon
	Hyung       Hyung
	Choong      Choong
	Pa          Pa
	Hae         Hae
	GyeokGak    GyeokGak
	IpMyo       IpMyo
}

type IpMyo struct {
	Exist   bool
	WhichJi int
}

type ChunGanHab struct {
	Exist     bool
	GabGi     int
	ElGyeong  int
	ByeongSin int
	JeongIm   int
	MuGye     int
}

type ChunGanGeok struct {
	Exist bool

	GabGyeong    int
	ElSin        int
	ByeongIm     int
	JeongHye     int
	MuGab        int
	GiEl         int
	GyeongByeong int
	SinJeong     int
	ImMu         int
	GyeGi        int
}

type SibsungResult struct {
	Exist   bool
	Sibsung int
}

type YukHab struct {
	Exist  bool
	InHye  int
	MyoSul int
	JinYu  int
	SaSin  int
	OMi    int
	JaChuk int
}

type SamHab struct {
	Exist    bool
	InOSul   int
	HaeMyoMi int
	SinJaJin int
	SaYuChuk int
}

type BanHab struct {
	Exist  bool
	InO    int
	InSul  int
	OSul   int
	HaeMyo int
	MyoMi  int
	HaeMi  int
	SinJa  int
	JaJin  int
	SinJin int
	SaYu   int
	YuChuk int
	SaChuk int
}

type BangHab struct {
	Exist     bool
	InMyoJin  int
	SaOMi     int
	SinYuSul  int
	HaeJaChuk int
}

type AmHab struct {
	Exist bool
	AmHab int
}

type MyeongAmHab struct {
	Exist     bool
	GabGi     int
	ElGyeong  int
	ByeongSin int
	JeongIm   int
	MuGye     int
}

type WonJin struct {
	Exist  bool
	InYu   int
	MyoSin int
	JinHae int
	SaSul  int
	OChuk  int
	JaMi   int
}

type GuiMoon struct {
	Exist bool
	InMi  int
	JaYu  int
}

type Hyung struct {
	Exist   bool
	InSa    int
	JaMyo   int
	JinJin  int
	SaSin   int
	OO      int
	YuYu    int
	SulMi   int
	ChukSul int
	HaeHae  int
}

type Choong struct {
	Exist  bool
	InSin  int
	MyoYu  int
	JinSul int
	SaHae  int
	JaO    int
	ChukMi int
}

type Pa struct {
	Exist   bool
	InHae   int
	MyoO    int
	JinChuk int
	SaSin   int
	SulMi   int
	JaYu    int
}

type Hae struct {
	Exist  bool
	InSa   int
	MyoJin int
	OChuk  int
	JaMi   int
	HaeSin int
	YuSul  int
}

type GyeokGak struct {
	Exist   bool
	InJa    int
	MyoSa   int
	MyoChuk int
	JinO    int
	OSin    int
	MiYu    int
	YuHae   int
	SulJa   int
}

func person_chungan_input(a []string, b []Chungan, d []Jiji, e []Sibsung, f []Chungan_Unsung) Person {
	var c Person
	c.Chun = make([]Chungan, 6)
	c.Ji = make([]Jiji, 6)
	c.Result = make([]Result_record, 6)
	for i := 0; i < len(a); i++ {
		if i < 6 {
			for j := 0; j < 10; j++ {
				if a[i] == b[j].Title {
					c.Chun[i] = b[j]

					break
				}
			}
		} else {
			for j := 0; j < 12; j++ {
				if a[i] == d[j].Title {
					c.Ji[i-len(a)/2] = d[j]

					break
				}
			}
		}
	}

	for i := 0; i < 4; i++ {
		switch c.Ji[i].Title {
		case "진":
			for j := 0; j < 4; j++ {
				if c.Chun[j].Title == "신" || c.Chun[j].Title == "임" {
					c.Chun[j].Properties.IpMyo.Exist = true
					c.Chun[j].Properties.IpMyo.WhichJi += (i + 1)
				}
			}
		case "술":
			for j := 0; j < 4; j++ {
				if c.Chun[j].Title == "병" || c.Chun[j].Title == "무" || c.Chun[j].Title == "을" {
					c.Chun[j].Properties.IpMyo.Exist = true
					c.Chun[j].Properties.IpMyo.WhichJi += (i + 1)
				}
			}
		case "축":
			for j := 0; j < 4; j++ {
				if c.Chun[j].Title == "경" || c.Chun[j].Title == "기" || c.Chun[j].Title == "정" {
					c.Chun[j].Properties.IpMyo.Exist = true
					c.Chun[j].Properties.IpMyo.WhichJi += (i + 1)
				}
			}
		case "미":
			for j := 0; j < 4; j++ {
				if c.Chun[j].Title == "갑" || c.Chun[j].Title == "해" {
					c.Chun[j].Properties.IpMyo.Exist = true
					c.Chun[j].Properties.IpMyo.WhichJi += (i + 1)
				}
			}
		}
	}

	for i := 0; i < 5; i++ { //십성 (본관을 넣어야함)
		if c.Chun[2].Properties.Prop == e[i].Prop {
			for j := 0; j < 6; j++ {
				for k := 0; k < 5; k++ {
					if c.Chun[j].Properties.Prop == e[i].Comp_Prop[k].Comp_Prop {
						//c.Chun[j].Properties.Sibsung = "본관"
						c.Chun[j].Properties.Sibsung = e[i].Comp_Prop[k].Title
					}

					if c.Ji[j].Properties.Prop == e[i].Comp_Prop[k].Comp_Prop {
						c.Ji[j].Properties.Sibsung = e[i].Comp_Prop[k].Title
					}

				}
			}
			break
		}
	}

	//십이운성
	for i := 0; i < len(a)/2; i++ {
		for j := 0; j < 10; j++ {
			if c.Chun[i].Title == f[j].Title {
				for k := 0; k < 12; k++ {
					if c.Ji[i].Title == f[j].Properties[k].Jiji_char {
						c.Chun[i].Properties.Unsung_by_Jiji.level = f[j].Properties[k].Level
						c.Chun[i].Properties.Unsung_by_Jiji.Unsung_title = f[j].Properties[k].Prop
					}
				}
			}

			if c.Chun[2].Title == f[j].Title {
				for k := 0; k < 12; k++ {
					if c.Ji[i].Title == f[j].Properties[k].Jiji_char {
						c.Chun[i].Properties.Unsung_Me.level = f[j].Properties[k].Level
						c.Chun[i].Properties.Unsung_Me.Unsung_title = f[j].Properties[k].Prop
					}
				}
			}
		}
	}

	//창고

	for i := 0; i < len(a)/2; i++ {
		c.Ji[i].Properties.ChangGo.Exist = false
		if c.Ji[i].Properties.Prop == "earth" {
			switch c.Chun[2].Properties.Prop {
			case "tree":
				switch c.Ji[i].Properties.Jijanggans[1].Chungan_char.Properties.Prop {
				case "tree":
					c.Ji[i].Properties.ChangGo.YangInGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "fire":
					c.Ji[i].Properties.ChangGo.SikSangGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "iron":
					c.Ji[i].Properties.ChangGo.GwanGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "water":
					c.Ji[i].Properties.ChangGo.InSungGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				}
			case "fire":
				switch c.Ji[i].Properties.Jijanggans[1].Chungan_char.Properties.Prop {
				case "tree":
					c.Ji[i].Properties.ChangGo.InSungGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "fire":
					c.Ji[i].Properties.ChangGo.YangInGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "iron":
					c.Ji[i].Properties.ChangGo.JaeGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "water":
					c.Ji[i].Properties.ChangGo.GwanGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				}
			case "earth":
				switch c.Ji[i].Properties.Jijanggans[1].Chungan_char.Properties.Prop {
				case "tree":
					c.Ji[i].Properties.ChangGo.GwanGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "fire":
					c.Ji[i].Properties.ChangGo.InSungGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "iron":
					c.Ji[i].Properties.ChangGo.SikSangGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "water":
					c.Ji[i].Properties.ChangGo.JaeGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				}
			case "iron":
				switch c.Ji[i].Properties.Jijanggans[1].Chungan_char.Properties.Prop {
				case "tree":
					c.Ji[i].Properties.ChangGo.JaeGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "fire":
					c.Ji[i].Properties.ChangGo.GwanGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "iron":
					c.Ji[i].Properties.ChangGo.YangInGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "water":
					c.Ji[i].Properties.ChangGo.SikSangGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				}
			case "water":
				switch c.Ji[i].Properties.Jijanggans[1].Chungan_char.Properties.Prop {
				case "tree":
					c.Ji[i].Properties.ChangGo.SikSangGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "fire":
					c.Ji[i].Properties.ChangGo.JaeGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "iron":
					c.Ji[i].Properties.ChangGo.InSungGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				case "water":
					c.Ji[i].Properties.ChangGo.YangInGo += 1
					c.Ji[i].Properties.ChangGo.Exist = true
				}
			}
		}
	}
	return c
}

func Umyang_balance(a Person, b Person) {
	count_plus := 0
	count_minus := 0

	for i := 0; i < 5; i++ {
		if a.Chun[i].Properties.Umyang == 1 {
			count_plus++
		} else if a.Chun[i].Properties.Umyang == 0 {
			count_minus++
		}

		if b.Chun[i].Properties.Umyang == 1 {
			count_plus++
		} else if b.Chun[i].Properties.Umyang == 0 {
			count_minus++
		}

		if a.Ji[i].Properties.Umyang == 1 {
			count_plus++
		} else if a.Ji[i].Properties.Umyang == 0 {
			count_minus++
		}

		if b.Ji[i].Properties.Umyang == 1 {
			count_plus++
		} else if b.Ji[i].Properties.Umyang == 0 {
			count_minus++
		}

	}

}

func Ohang_balance(a Person, b Person) {
	count_tree := 0
	count_fire := 0
	count_earth := 0
	count_iron := 0
	count_water := 0

	for i := 0; i < 4; i++ {
		if a.Chun[i].Properties.Prop == "tree" {
			count_tree++
		} else if a.Chun[i].Properties.Prop == "fire" {
			count_fire++
		} else if a.Chun[i].Properties.Prop == "earth" {
			count_earth++
		} else if a.Chun[i].Properties.Prop == "iron" {
			count_iron++
		} else if a.Chun[i].Properties.Prop == "water" {
			count_water++
		}

		if b.Chun[i].Properties.Prop == "tree" {
			count_tree++
		} else if b.Chun[i].Properties.Prop == "fire" {
			count_fire++
		} else if b.Chun[i].Properties.Prop == "earth" {
			count_earth++
		} else if b.Chun[i].Properties.Prop == "iron" {
			count_iron++
		} else if b.Chun[i].Properties.Prop == "water" {
			count_water++
		}

		if a.Ji[i].Properties.Prop == "tree" {
			count_tree++
		} else if a.Ji[i].Properties.Prop == "fire" {
			count_fire++
		} else if a.Ji[i].Properties.Prop == "earth" {
			count_earth++
		} else if a.Ji[i].Properties.Prop == "iron" {
			count_iron++
		} else if a.Ji[i].Properties.Prop == "water" {
			count_water++
		}

		if b.Ji[i].Properties.Prop == "tree" {
			count_tree++
		} else if b.Ji[i].Properties.Prop == "fire" {
			count_fire++
		} else if b.Ji[i].Properties.Prop == "earth" {
			count_earth++
		} else if b.Ji[i].Properties.Prop == "iron" {
			count_iron++
		} else if b.Ji[i].Properties.Prop == "water" {
			count_water++
		}
	}
	fmt.Println(count_tree, count_fire, count_earth, count_iron, count_water)

}

func chungan_hab(a Person) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			a.Result[i].ChunGanHab.Exist = false
			a.Result[j].ChunGanHab.Exist = false
			if math.Abs(float64(a.Chun[i].Id-a.Chun[j].Id)) == 5 {
				switch {
				case a.Chun[i].Title == "갑" || a.Chun[i].Title == "기":
					a.Result[i].ChunGanHab.Exist = true
					a.Result[j].ChunGanHab.Exist = true
					a.Result[i].ChunGanHab.GabGi += 1
					a.Result[j].ChunGanHab.GabGi += 1
				case a.Chun[i].Title == "을" || a.Chun[i].Title == "경":
					a.Result[i].ChunGanHab.Exist = true
					a.Result[j].ChunGanHab.Exist = true
					a.Result[i].ChunGanHab.ElGyeong += 1
					a.Result[j].ChunGanHab.ElGyeong += 1
				case a.Chun[i].Title == "병" || a.Chun[i].Title == "신":
					a.Result[i].ChunGanHab.Exist = true
					a.Result[j].ChunGanHab.Exist = true
					a.Result[i].ChunGanHab.ByeongSin += 1
					a.Result[j].ChunGanHab.ByeongSin += 1
				case a.Chun[i].Title == "정" || a.Chun[i].Title == "임":
					a.Result[i].ChunGanHab.Exist = true
					a.Result[j].ChunGanHab.Exist = true
					a.Result[i].ChunGanHab.JeongIm += 1
					a.Result[j].ChunGanHab.JeongIm += 1
				case a.Chun[i].Title == "무" || a.Chun[i].Title == "계":
					a.Result[i].ChunGanHab.Exist = true
					a.Result[j].ChunGanHab.Exist = true
					a.Result[i].ChunGanHab.MuGye += 1
					a.Result[j].ChunGanHab.MuGye += 1

				}
			}
		}
	}
}

func Find_Chungan_hab(a Person, b Person) (Person, Person) {
	for i := 1; i < 3; i++ {
		a.Result[i].ChunGanHab.Exist = false
		b.Result[i].ChunGanHab.Exist = false
		if math.Abs(float64(a.Chun[i].Id-b.Chun[i].Id)) == 5 {
			switch {
			case a.Chun[i].Title == "갑" || a.Chun[i].Title == "기":
				a.Result[i].ChunGanHab.GabGi += 1
				b.Result[i].ChunGanHab.GabGi += 1
				a.Result[i].ChunGanHab.Exist = true
				b.Result[i].ChunGanHab.Exist = true

			case a.Chun[i].Title == "을" || a.Chun[i].Title == "경":
				a.Result[i].ChunGanHab.ElGyeong += 1
				b.Result[i].ChunGanHab.ElGyeong += 1
				a.Result[i].ChunGanHab.Exist = true
				b.Result[i].ChunGanHab.Exist = true
			case a.Chun[i].Title == "병" || a.Chun[i].Title == "신":
				a.Result[i].ChunGanHab.ByeongSin += 1
				b.Result[i].ChunGanHab.ByeongSin += 1
				a.Result[i].ChunGanHab.Exist = true
				b.Result[i].ChunGanHab.Exist = true
			case a.Chun[i].Title == "정" || a.Chun[i].Title == "임":
				a.Result[i].ChunGanHab.JeongIm += 1
				b.Result[i].ChunGanHab.JeongIm += 1
				a.Result[i].ChunGanHab.Exist = true
				b.Result[i].ChunGanHab.Exist = true
			case a.Chun[i].Title == "무" || a.Chun[i].Title == "계":
				a.Result[i].ChunGanHab.MuGye += 1
				b.Result[i].ChunGanHab.MuGye += 1
				a.Result[i].ChunGanHab.Exist = true
				b.Result[i].ChunGanHab.Exist = true
			}
		}
	}
	return a, b
}

func chungan_geok(a Person) Person {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			a.Result[i].ChunGanGeok.Exist = false
			a.Result[j].ChunGanGeok.Exist = false
			if i != j {
				if math.Abs(float64(a.Chun[i].Id-a.Chun[j].Id)) == 6 {
					switch {
					case a.Chun[i].Title == "갑":
						a.Result[i].ChunGanGeok.Exist = true
						a.Result[j].ChunGanGeok.Exist = true
						a.Result[i].ChunGanGeok.GabGyeong -= 1
						a.Result[j].ChunGanGeok.GabGyeong += 1
					case a.Chun[i].Title == "을":
						a.Result[i].ChunGanGeok.Exist = true
						a.Result[j].ChunGanGeok.Exist = true
						a.Result[i].ChunGanGeok.ElSin -= 1
						a.Result[j].ChunGanGeok.ElSin += 1
					case a.Chun[i].Title == "병":
						a.Result[i].ChunGanGeok.Exist = true
						a.Result[j].ChunGanGeok.Exist = true
						a.Result[i].ChunGanGeok.ByeongIm -= 1
						a.Result[j].ChunGanGeok.ByeongIm += 1
					case a.Chun[i].Title == "정":
						a.Result[i].ChunGanGeok.Exist = true
						a.Result[j].ChunGanGeok.Exist = true
						a.Result[i].ChunGanGeok.JeongHye -= 1
						a.Result[j].ChunGanGeok.JeongHye += 1
					case a.Chun[i].Title == "무":
						a.Result[i].ChunGanGeok.Exist = true
						a.Result[j].ChunGanGeok.Exist = true
						a.Result[i].ChunGanGeok.MuGab -= 1
						a.Result[j].ChunGanGeok.MuGab += 1
					case a.Chun[i].Title == "기":
						a.Result[i].ChunGanGeok.Exist = true
						a.Result[j].ChunGanGeok.Exist = true
						a.Result[i].ChunGanGeok.GiEl -= 1
						a.Result[j].ChunGanGeok.GiEl += 1
					case a.Chun[i].Title == "경":
						a.Result[i].ChunGanGeok.Exist = true
						a.Result[j].ChunGanGeok.Exist = true
						a.Result[i].ChunGanGeok.GyeongByeong -= 1
						a.Result[j].ChunGanGeok.GyeongByeong += 1
					case a.Chun[i].Title == "신":
						a.Result[i].ChunGanGeok.Exist = true
						a.Result[j].ChunGanGeok.Exist = true
						a.Result[i].ChunGanGeok.SinJeong -= 1
						a.Result[j].ChunGanGeok.SinJeong += 1
					case a.Chun[i].Title == "임":
						a.Result[i].ChunGanGeok.Exist = true
						a.Result[j].ChunGanGeok.Exist = true
						a.Result[i].ChunGanGeok.ImMu -= 1
						a.Result[j].ChunGanGeok.ImMu += 1
					case a.Chun[i].Title == "계":
						a.Result[i].ChunGanGeok.Exist = true
						a.Result[j].ChunGanGeok.Exist = true
						a.Result[i].ChunGanGeok.GyeGi -= 1
						a.Result[j].ChunGanGeok.GyeGi += 1
					}
				}
			}
		}
	}
	return a
}

func Find_Chungan_Geok(a Person, b Person) (Person, Person) {
	for i := 1; i < 3; i++ {
		a.Result[i].ChunGanGeok.Exist = false
		b.Result[i].ChunGanGeok.Exist = false
		if math.Abs(float64(a.Chun[i].Id-b.Chun[i].Id)) == 6 { //내가 극 당하는 것
			switch {
			case a.Chun[i].Title == "갑":
				a.Result[i].ChunGanGeok.Exist = true

				a.Result[i].ChunGanGeok.GabGyeong -= 1

			case a.Chun[i].Title == "을":
				a.Result[i].ChunGanGeok.Exist = true

				a.Result[i].ChunGanGeok.ElSin -= 1

			case a.Chun[i].Title == "병":
				a.Result[i].ChunGanGeok.Exist = true

				a.Result[i].ChunGanGeok.ByeongIm -= 1

			case a.Chun[i].Title == "정":
				a.Result[i].ChunGanGeok.Exist = true

				a.Result[i].ChunGanGeok.JeongHye -= 1

			case a.Chun[i].Title == "무":
				a.Result[i].ChunGanGeok.Exist = true

				a.Result[i].ChunGanGeok.MuGab -= 1

			case a.Chun[i].Title == "기":
				a.Result[i].ChunGanGeok.Exist = true

				a.Result[i].ChunGanGeok.GiEl -= 1

			case a.Chun[i].Title == "경":
				a.Result[i].ChunGanGeok.Exist = true

				a.Result[i].ChunGanGeok.GyeongByeong -= 1

			case a.Chun[i].Title == "신":
				a.Result[i].ChunGanGeok.Exist = true

				a.Result[i].ChunGanGeok.SinJeong -= 1

			case a.Chun[i].Title == "임":
				a.Result[i].ChunGanGeok.Exist = true

				a.Result[i].ChunGanGeok.ImMu -= 1

			case a.Chun[i].Title == "계":
				a.Result[i].ChunGanGeok.Exist = true

				a.Result[i].ChunGanGeok.GyeGi -= 1

			}
		}

		if math.Abs(float64(a.Chun[i].Id-b.Chun[i].Id)) == 4 { // 상대방이 극 당하는 것
			switch {
			case a.Chun[i].Title == "갑":

				b.Result[i].ChunGanGeok.Exist = true

				b.Result[i].ChunGanGeok.GabGyeong -= 1
			case a.Chun[i].Title == "을":

				b.Result[i].ChunGanGeok.Exist = true

				b.Result[i].ChunGanGeok.ElSin -= 1
			case a.Chun[i].Title == "병":

				b.Result[i].ChunGanGeok.Exist = true

				b.Result[i].ChunGanGeok.ByeongIm -= 1
			case a.Chun[i].Title == "정":

				b.Result[i].ChunGanGeok.Exist = true

				b.Result[i].ChunGanGeok.JeongHye -= 1
			case a.Chun[i].Title == "무":

				b.Result[i].ChunGanGeok.Exist = true

				b.Result[i].ChunGanGeok.MuGab -= 1
			case a.Chun[i].Title == "기":

				b.Result[i].ChunGanGeok.Exist = true

				b.Result[i].ChunGanGeok.GiEl -= 1
			case a.Chun[i].Title == "경":

				b.Result[i].ChunGanGeok.Exist = true

				b.Result[i].ChunGanGeok.GyeongByeong -= 1
			case a.Chun[i].Title == "신":

				b.Result[i].ChunGanGeok.Exist = true

				b.Result[i].ChunGanGeok.SinJeong -= 1
			case a.Chun[i].Title == "임":

				b.Result[i].ChunGanGeok.Exist = true

				b.Result[i].ChunGanGeok.ImMu -= 1
			case a.Chun[i].Title == "계":

				b.Result[i].ChunGanGeok.Exist = true

				b.Result[i].ChunGanGeok.GyeGi -= 1
			}
		}
	}
	return a, b
}

func Find_Ipmyo(a, Person, b Person) (Person, Person) {
	for i := 1; i < 3; i++ {
		a.Result[i].IpMyo.Exist = false
		b.Result[i].IpMyo.Exist = false
		a.Result[i].IpMyo.WhichJi = 0
		b.Result[i].IpMyo.WhichJi = 0
		switch a.Ji[i].Title {
		case "진":
			if a.Chun[i].Title == "신" || a.Chun[i].Title == "임" {
				a.Result[i].IpMyo.Exist = true
				a.Result[i].IpMyo.WhichJi += 1
			} else if b.Chun[i].Title == "신" || b.Chun[i].Title == "임" {
				b.Result[i].IpMyo.Exist = true
				b.Result[i].IpMyo.WhichJi += 1
			}
		case "술":
			if a.Chun[i].Title == "병" || a.Chun[i].Title == "무" || a.Chun[i].Title == "을" {
				a.Result[i].IpMyo.Exist = true
				a.Result[i].IpMyo.WhichJi += 1
			} else if b.Chun[i].Title == "병" || b.Chun[i].Title == "무" || b.Chun[i].Title == "을" {
				b.Result[i].IpMyo.Exist = true
				b.Result[i].IpMyo.WhichJi += 1
			}
		case "축":
			if a.Chun[i].Title == "경" || a.Chun[i].Title == "기" || a.Chun[i].Title == "정" {
				a.Result[i].IpMyo.Exist = true
				a.Result[i].IpMyo.WhichJi += 1
			} else if b.Chun[i].Title == "경" || b.Chun[i].Title == "기" || b.Chun[i].Title == "정" {
				b.Result[i].IpMyo.Exist = true
				b.Result[i].IpMyo.WhichJi += 1
			}
		case "미":
			if a.Chun[i].Title == "갑" || a.Chun[i].Title == "해" {
				a.Result[i].IpMyo.Exist = true
				a.Result[i].IpMyo.WhichJi += 1
			} else if b.Chun[i].Title == "갑" || b.Chun[i].Title == "해" {
				b.Result[i].IpMyo.Exist = true
				b.Result[i].IpMyo.WhichJi += 1
			}
		}

		switch b.Ji[i].Title {
		case "진":
			if a.Chun[i].Title == "신" || a.Chun[i].Title == "임" {
				a.Result[i].IpMyo.Exist = true
				a.Result[i].IpMyo.WhichJi += 2
			} else if b.Chun[i].Title == "신" || b.Chun[i].Title == "임" {
				b.Result[i].IpMyo.Exist = true
				b.Result[i].IpMyo.WhichJi += 2
			}
		case "술":

			if a.Chun[i].Title == "병" || a.Chun[i].Title == "무" || a.Chun[i].Title == "을" {
				a.Result[i].IpMyo.Exist = true
				a.Result[i].IpMyo.WhichJi += 2
			} else if b.Chun[i].Title == "병" || b.Chun[i].Title == "무" || b.Chun[i].Title == "을" {
				b.Result[i].IpMyo.Exist = true
				b.Result[i].IpMyo.WhichJi += 2
			}

		case "축":
			if a.Chun[i].Title == "경" || a.Chun[i].Title == "기" || a.Chun[i].Title == "정" {
				a.Result[i].IpMyo.Exist = true
				a.Result[i].IpMyo.WhichJi += 2
			} else if b.Chun[i].Title == "경" || b.Chun[i].Title == "기" || b.Chun[i].Title == "정" {
				b.Result[i].IpMyo.Exist = true
				b.Result[i].IpMyo.WhichJi += 2
			}
		case "미":
			if a.Chun[i].Title == "갑" || a.Chun[i].Title == "해" {
				a.Result[i].IpMyo.Exist = true
				a.Result[i].IpMyo.WhichJi += 2
			} else if b.Chun[i].Title == "갑" || b.Chun[i].Title == "해" {
				b.Result[i].IpMyo.Exist = true
				b.Result[i].IpMyo.WhichJi += 2
			}
		}
	}
	return a, b
}

func Find_Sibsung_Goonghab(a Person, b Person, e []Sibsung) (Person, Person) {
	for i := 0; i < 5; i++ {
		if a.Chun[1].Properties.Prop == e[i].Prop {
			for k := 0; k < 5; k++ {
				if b.Chun[1].Properties.Prop == e[i].Comp_Prop[k].Comp_Prop {
					a.Chun[1].Properties.Sibsung = e[i].Comp_Prop[k].Title
				}
			}
		}

		if a.Chun[2].Properties.Prop == e[i].Prop {
			for k := 0; k < 5; k++ {
				if b.Chun[2].Properties.Prop == e[i].Comp_Prop[k].Comp_Prop {
					a.Chun[2].Properties.Sibsung = e[i].Comp_Prop[k].Title
				}
			}
		}

		if b.Chun[1].Properties.Prop == e[i].Prop {
			for k := 0; k < 5; k++ {
				if a.Chun[1].Properties.Prop == e[i].Comp_Prop[k].Comp_Prop {
					b.Chun[1].Properties.Sibsung = e[i].Comp_Prop[k].Title
				}
			}
		}

		if b.Chun[2].Properties.Prop == e[i].Prop {
			for k := 0; k < 5; k++ {
				if a.Chun[2].Properties.Prop == e[i].Comp_Prop[k].Comp_Prop {
					b.Chun[2].Properties.Sibsung = e[i].Comp_Prop[k].Title
				}
			}
		}
	}
	a.Result[1].Sibsung.Exist = false
	b.Result[2].Sibsung.Exist = false
	a.Result[1].Sibsung.Exist = false
	b.Result[2].Sibsung.Exist = false
	switch a.Chun[2].Properties.Sibsung {
	case "비겁":
		switch b.Chun[2].Properties.Sibsung {
		case "식상":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang {
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang {
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 2
				} else {
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 4
				}
			} else {
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang {
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 2
				} else {
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 4
				}
			}
		case "재성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang {
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang {
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -3
				} else {
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -1
				}
			} else {
				b.Result[2].Sibsung.Exist = true
				b.Result[2].Sibsung.Sibsung += -4
			}
		case "관성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang {
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang {
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -4
				} else {
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -1
				}
			} else {
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang {
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -2
				} else {
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -4
				}
			}
		case "인성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //비견
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
				} else { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 4
				}
			} else { //겁재
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
				} else { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 4
				}
			}
		}
	case "식상":
		switch b.Chun[2].Properties.Sibsung {
		case "비겁":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang {
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang {
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
				} else {
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 4
				}
			} else {
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang {
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
				} else {
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 4
				}
			}
		case "재성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang {
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang {
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 3
				} else {
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 3
				}
			} else {
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang {
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 3
				} else {
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 3
				}
			}

		case "관성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 4
				} else { //정
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -4
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { // 정
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -4
				} else { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 1
				}
			}

		case "인성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -3
				} else { //정
					if a.Chun[2].Properties.Umyang == 1 {
						a.Result[2].Sibsung.Exist = true
						a.Result[2].Sibsung.Sibsung += -3
					} else {
						a.Result[2].Sibsung.Exist = true
						b.Result[2].Sibsung.Exist = true
						a.Result[2].Sibsung.Sibsung += 2
						b.Result[2].Sibsung.Sibsung += 2
					}
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { // 정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -4
				} else { //편
					if a.Chun[2].Properties.Umyang == 1 {
						a.Result[2].Sibsung.Exist = true
						b.Result[2].Sibsung.Exist = true
						a.Result[2].Sibsung.Sibsung += 2
						b.Result[2].Sibsung.Sibsung += 2
					} else {
						a.Result[2].Sibsung.Exist = true
						a.Result[2].Sibsung.Sibsung += -4
					}
				}
			}
		}
	case "재성":
		switch b.Chun[2].Properties.Sibsung {
		case "비겁":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -3
				} else { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -4
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -4
				} else { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -1
				}
			}
		case "식상":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 3
				} else { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 3
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 3
				} else { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 3
				}
			}
		case "관성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -3
				} else { //정
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 3
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					b.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 3
				} else { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -3
				}
			}
		case "인성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -2
				} else { //정
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -2
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					b.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -2
				} else { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -2
				}
			}
		}
	case "관성":
		switch b.Chun[2].Properties.Sibsung {
		case "비겁":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -4
				} else { //정
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -4
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -2
				} else { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -1
				}
			}
		case "식상":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -4
				} else { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 1
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -4
				} else { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 1
				}
			}
		case "재성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -3
				} else { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -3
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 3
				} else { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 3
				}
			}
		case "인성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					a.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
					b.Result[2].Sibsung.Sibsung += 2
				} else { //정
					a.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
					b.Result[2].Sibsung.Sibsung += 2
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					a.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
					b.Result[2].Sibsung.Sibsung += 2
				} else { //편
					a.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
					b.Result[2].Sibsung.Sibsung += 2
				}
			}
		}
	case "인성":
		switch b.Chun[2].Properties.Sibsung {
		case "비겁":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 1
				} else { //정
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 3
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 1
				} else { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += 3
				}
			}
		case "식상":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -4
				} else { //정
					if a.Chun[2].Properties.Umyang == 1 {
						a.Result[2].Sibsung.Exist = true
						b.Result[2].Sibsung.Exist = true
						a.Result[2].Sibsung.Sibsung += 2
						b.Result[2].Sibsung.Sibsung += 2
					} else {
						b.Result[2].Sibsung.Exist = true
						b.Result[2].Sibsung.Sibsung += -3
					}
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { // 정
					b.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Sibsung += -4
				} else { //편
					if a.Chun[2].Properties.Umyang != 1 {
						a.Result[2].Sibsung.Exist = true
						b.Result[2].Sibsung.Exist = true
						a.Result[2].Sibsung.Sibsung += 2
						b.Result[2].Sibsung.Sibsung += 2
					} else {
						b.Result[2].Sibsung.Exist = true
						b.Result[2].Sibsung.Sibsung += -3
					}
				}
			}
		case "재성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -2
				} else { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -2
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -2
				} else { //편
					a.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += -2
				}
			}
		case "관성":
			if a.Chun[1].Properties.Umyang == a.Chun[2].Properties.Umyang { //편
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //편
					a.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
					b.Result[2].Sibsung.Sibsung += 2
				} else { //정
					a.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
					b.Result[2].Sibsung.Sibsung += 2
				}
			} else { //정
				if a.Chun[2].Properties.Umyang == b.Chun[2].Properties.Umyang { //정
					a.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
					b.Result[2].Sibsung.Sibsung += 2
				} else { //편
					a.Result[2].Sibsung.Exist = true
					b.Result[2].Sibsung.Exist = true
					a.Result[2].Sibsung.Sibsung += 2
					b.Result[2].Sibsung.Sibsung += 2
				}
			}
		}
	}
	return a, b
}

func Find_Unsung_Goonghab(a Person, b Person, f []Chungan_Unsung) (Person, Person) {
	for j := 0; j < 10; j++ {
		if a.Chun[1].Title == f[j].Title {
			for k := 0; k < 12; k++ {
				if b.Ji[1].Title == f[j].Properties[k].Jiji_char {
					a.Chun[1].Properties.Unsung_Me.level = f[j].Properties[k].Level
					a.Chun[1].Properties.Unsung_Me.Unsung_title = f[j].Properties[k].Prop
				}
			}
		}
		if b.Chun[1].Title == f[j].Title {
			for k := 0; k < 12; k++ {
				if a.Ji[1].Title == f[j].Properties[k].Jiji_char {
					b.Chun[1].Properties.Unsung_Me.level = f[j].Properties[k].Level
					b.Chun[1].Properties.Unsung_Me.Unsung_title = f[j].Properties[k].Prop
				}
			}
		}
	}

	for j := 0; j < 10; j++ {
		if a.Chun[2].Title == f[j].Title {
			for k := 0; k < 12; k++ {
				if b.Ji[2].Title == f[j].Properties[k].Jiji_char {
					a.Chun[2].Properties.Unsung_Me.level = f[j].Properties[k].Level
					a.Chun[2].Properties.Unsung_Me.Unsung_title = f[j].Properties[k].Prop
				}
			}
		}
		if b.Chun[2].Title == f[j].Title {
			for k := 0; k < 12; k++ {
				if a.Ji[2].Title == f[j].Properties[k].Jiji_char {
					b.Chun[2].Properties.Unsung_Me.level = f[j].Properties[k].Level
					b.Chun[2].Properties.Unsung_Me.Unsung_title = f[j].Properties[k].Prop
				}
			}
		}
	}
	return a, b
}

func Find_Banghab(a Person) {
	num1 := 5
	num2 := 5
	num3 := 5
	//인묘진
	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "인" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "묘" {
			num2 = j
		}
	}
	for k := 0; k < 4; k++ {
		if a.Ji[k].Title == "진" {
			num3 = k
		}
	}
	a.Result[num3].BangHab.Exist = false
	a.Result[num2].BangHab.Exist = false
	a.Result[num1].BangHab.Exist = false
	a.Result[4].BangHab.Exist = false
	a.Result[5].BangHab.Exist = false
	if num1 != 5 || num2 != 5 || num3 != 5 {
		switch {
		case num1 == 5 && num2 == 5 && num3 == 5:
			{
				a.Result[num1].BangHab.InMyoJin += 1
				a.Result[num2].BangHab.InMyoJin += 1
				a.Result[num3].BangHab.InMyoJin += 1
				a.Result[num3].BangHab.Exist = true
				a.Result[num2].BangHab.Exist = true
				a.Result[num1].BangHab.Exist = true
			}
		case num1 == 5:
			if num2 == 5 {
				if a.Ji[4].Title == "인" && a.Ji[5].Title == "묘" {
					a.Result[4].BangHab.InMyoJin += 1
					a.Result[5].BangHab.InMyoJin += 1
					a.Result[num3].BangHab.InMyoJin += 1
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
				} else if a.Ji[4].Title == "묘" && a.Ji[5].Title == "인" {
					a.Result[4].BangHab.InMyoJin += 1
					a.Result[5].BangHab.InMyoJin += 1
					a.Result[num3].BangHab.InMyoJin += 1
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
				}

			} else if num3 == 5 {
				if a.Ji[4].Title == "인" && a.Ji[5].Title == "진" {
					a.Result[4].BangHab.InMyoJin += 1
					a.Result[5].BangHab.InMyoJin += 1
					a.Result[num2].BangHab.InMyoJin += 1
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
				} else if a.Ji[4].Title == "진" && a.Ji[5].Title == "인" {
					a.Result[4].BangHab.InMyoJin += 1
					a.Result[5].BangHab.InMyoJin += 1
					a.Result[num2].BangHab.InMyoJin += 1
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
				}

			} else if num2 != 5 && num3 != 5 {
				if a.Ji[4].Title == "인" {
					a.Result[4].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[4].BangHab.InMyoJin += 1
					a.Result[num2].BangHab.InMyoJin += 1
					a.Result[num3].BangHab.InMyoJin += 1
				}
				if a.Ji[5].Title == "인" {
					a.Result[5].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[5].BangHab.InMyoJin += 1
					a.Result[num2].BangHab.InMyoJin += 1
					a.Result[num3].BangHab.InMyoJin += 1
				}

			}
			fallthrough

		case num2 == 5:
			if num3 == 5 {
				if a.Ji[4].Title == "묘" && a.Ji[5].Title == "진" {
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[4].BangHab.InMyoJin += 1
					a.Result[5].BangHab.InMyoJin += 1
					a.Result[num1].BangHab.InMyoJin += 1
				} else if a.Ji[4].Title == "진" && a.Ji[5].Title == "묘" {
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[4].BangHab.InMyoJin += 1
					a.Result[5].BangHab.InMyoJin += 1
					a.Result[num1].BangHab.InMyoJin += 1
				}

			} else if num1 != 5 && num3 != 5 {
				if a.Ji[4].Title == "묘" {
					a.Result[4].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[4].BangHab.InMyoJin += 1
					a.Result[num1].BangHab.InMyoJin += 1
					a.Result[num3].BangHab.InMyoJin += 1
				} else if a.Ji[5].Title == "묘" {
					a.Result[4].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[5].BangHab.InMyoJin += 1
					a.Result[num1].BangHab.InMyoJin += 1
					a.Result[num3].BangHab.InMyoJin += 1
				}

			}
			fallthrough

		case num3 == 5:
			if num1 != 5 && num2 != 5 {
				if a.Ji[4].Title == "진" {
					a.Result[4].BangHab.InMyoJin += 1
					a.Result[num1].BangHab.InMyoJin += 1
					a.Result[num2].BangHab.InMyoJin += 1
					a.Result[num2].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[4].BangHab.Exist = true
				} else if a.Ji[5].Title == "진" {
					a.Result[5].BangHab.InMyoJin += 1
					a.Result[num1].BangHab.InMyoJin += 1
					a.Result[num2].BangHab.InMyoJin += 1
					a.Result[5].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
				}

			}

		}
	}

	//사오미
	num1 = 5
	num2 = 5
	num3 = 5
	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "사" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "오" {
			num2 = j
		}
	}
	for k := 0; k < 4; k++ {
		if a.Ji[k].Title == "미" {
			num3 = k
		}
	}

	if num1 != 5 || num2 != 5 || num3 != 5 {
		switch {
		case num1 == 5 && num2 == 5 && num3 == 5:
			{
				a.Result[num1].BangHab.Exist = true
				a.Result[num3].BangHab.Exist = true
				a.Result[num3].BangHab.Exist = true
				a.Result[num1].BangHab.SaOMi += 1
				a.Result[num2].BangHab.SaOMi += 1
				a.Result[num3].BangHab.SaOMi += 1
			}
		case num1 == 5:
			if num2 == 5 {
				if a.Ji[4].Title == "사" && a.Ji[5].Title == "오" {
					a.Result[4].BangHab.SaOMi += 1
					a.Result[5].BangHab.SaOMi += 1
					a.Result[num3].BangHab.SaOMi += 1
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
				} else if a.Ji[4].Title == "오" && a.Ji[5].Title == "사" {
					a.Result[4].BangHab.SaOMi += 1
					a.Result[5].BangHab.SaOMi += 1
					a.Result[num3].BangHab.SaOMi += 1
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
				}
			} else if num3 == 5 {
				if a.Ji[4].Title == "사" && a.Ji[5].Title == "미" {
					a.Result[4].BangHab.SaOMi += 1
					a.Result[5].BangHab.SaOMi += 1
					a.Result[num2].BangHab.SaOMi += 1
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
				} else if a.Ji[4].Title == "미" && a.Ji[5].Title == "사" {
					a.Result[4].BangHab.SaOMi += 1
					a.Result[5].BangHab.SaOMi += 1
					a.Result[num2].BangHab.SaOMi += 1
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
				}
			} else if num2 != 5 && num3 != 5 {
				if a.Ji[4].Title == "사" {
					a.Result[4].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[4].BangHab.SaOMi += 1
					a.Result[num2].BangHab.SaOMi += 1
					a.Result[num3].BangHab.SaOMi += 1
				}
				if a.Ji[5].Title == "사" {
					a.Result[5].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[5].BangHab.SaOMi += 1
					a.Result[num2].BangHab.SaOMi += 1
					a.Result[num3].BangHab.SaOMi += 1
				}
			}
			fallthrough

		case num2 == 5:
			if num3 == 5 {
				if a.Ji[4].Title == "오" && a.Ji[5].Title == "미" {
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[4].BangHab.SaOMi += 1
					a.Result[5].BangHab.SaOMi += 1
					a.Result[num1].BangHab.SaOMi += 1
				} else if a.Ji[4].Title == "미" && a.Ji[5].Title == "오" {
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[4].BangHab.SaOMi += 1
					a.Result[5].BangHab.SaOMi += 1
					a.Result[num1].BangHab.SaOMi += 1
				}
			} else if num1 != 5 && num3 != 5 {
				if a.Ji[4].Title == "오" {
					a.Result[4].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[4].BangHab.SaOMi += 1
					a.Result[num1].BangHab.SaOMi += 1
					a.Result[num3].BangHab.SaOMi += 1
				} else if a.Ji[5].Title == "오" {
					a.Result[5].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[5].BangHab.SaOMi += 1
					a.Result[num1].BangHab.SaOMi += 1
					a.Result[num3].BangHab.SaOMi += 1
				}
			}
			fallthrough

		case num3 == 5:
			if num1 != 5 && num2 != 5 {
				if a.Ji[4].Title == "미" {
					a.Result[4].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[4].BangHab.SaOMi += 1
					a.Result[num1].BangHab.SaOMi += 1
					a.Result[num2].BangHab.SaOMi += 1
				} else if a.Ji[5].Title == "미" {
					a.Result[5].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[5].BangHab.SaOMi += 1
					a.Result[num1].BangHab.SaOMi += 1
					a.Result[num2].BangHab.SaOMi += 1
				}
			}

		}
	}

	//신유술
	num1 = 5
	num2 = 5
	num3 = 5
	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "신" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "유" {
			num2 = j
		}
	}
	for k := 0; k < 4; k++ {
		if a.Ji[k].Title == "술" {
			num3 = k
		}
	}

	if num1 != 5 || num2 != 5 || num3 != 5 {
		switch {
		case num1 == 5 && num2 == 5 && num3 == 5:
			{
				a.Result[num1].BangHab.Exist = true
				a.Result[num3].BangHab.Exist = true
				a.Result[num3].BangHab.Exist = true
				a.Result[num1].BangHab.SaOMi += 1
				a.Result[num2].BangHab.SaOMi += 1
				a.Result[num3].BangHab.SaOMi += 1
			}
		case num1 == 5:
			if num2 == 5 {
				if a.Ji[4].Title == "신" && a.Ji[5].Title == "유" {
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[4].BangHab.SinYuSul += 1
					a.Result[5].BangHab.SinYuSul += 1
					a.Result[num3].BangHab.SinYuSul += 1
				} else if a.Ji[4].Title == "유" && a.Ji[5].Title == "신" {
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[4].BangHab.SinYuSul += 1
					a.Result[5].BangHab.SinYuSul += 1
					a.Result[num3].BangHab.SinYuSul += 1
				}
			} else if num3 == 5 {
				if a.Ji[4].Title == "신" && a.Ji[5].Title == "술" {
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[4].BangHab.SinYuSul += 1
					a.Result[5].BangHab.SinYuSul += 1
					a.Result[num2].BangHab.SinYuSul += 1
				} else if a.Ji[4].Title == "술" && a.Ji[5].Title == "신" {
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[4].BangHab.SinYuSul += 1
					a.Result[5].BangHab.SinYuSul += 1
					a.Result[num2].BangHab.SinYuSul += 1
				}
			} else if num2 != 5 && num3 != 5 {
				if a.Ji[4].Title == "신" {
					a.Result[4].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[4].BangHab.SinYuSul += 1
					a.Result[num2].BangHab.SinYuSul += 1
					a.Result[num3].BangHab.SinYuSul += 1
				}
				if a.Ji[5].Title == "신" {
					a.Result[5].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[5].BangHab.SinYuSul += 1
					a.Result[num2].BangHab.SinYuSul += 1
					a.Result[num3].BangHab.SinYuSul += 1
				}
			}
			fallthrough

		case num2 == 5:
			if num3 == 5 {
				if a.Ji[4].Title == "유" && a.Ji[5].Title == "술" {
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[4].BangHab.SinYuSul += 1
					a.Result[5].BangHab.SinYuSul += 1
					a.Result[num1].BangHab.SinYuSul += 1
				} else if a.Ji[4].Title == "술" && a.Ji[5].Title == "유" {
					a.Result[4].BangHab.Exist = true
					a.Result[5].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[4].BangHab.SinYuSul += 1
					a.Result[5].BangHab.SinYuSul += 1
					a.Result[num1].BangHab.SinYuSul += 1
				}
			} else if num1 != 5 && num3 != 5 {
				if a.Ji[4].Title == "유" {
					a.Result[4].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[4].BangHab.SinYuSul += 1
					a.Result[num1].BangHab.SinYuSul += 1
					a.Result[num3].BangHab.SinYuSul += 1
				} else if a.Ji[5].Title == "유" {
					a.Result[5].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[5].BangHab.SinYuSul += 1
					a.Result[num1].BangHab.SinYuSul += 1
					a.Result[num3].BangHab.SinYuSul += 1
				}
			}
			fallthrough

		case num3 == 5:
			if num1 != 5 && num2 != 5 {
				if a.Ji[4].Title == "술" {
					a.Result[4].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[4].BangHab.SinYuSul += 1
					a.Result[num1].BangHab.SinYuSul += 1
					a.Result[num2].BangHab.SinYuSul += 1
				} else if a.Ji[5].Title == "술" {
					a.Result[5].BangHab.Exist = true
					a.Result[num1].BangHab.Exist = true
					a.Result[num2].BangHab.Exist = true
					a.Result[5].BangHab.SinYuSul += 1
					a.Result[num1].BangHab.SinYuSul += 1
					a.Result[num2].BangHab.SinYuSul += 1
				}
			}
		}

		//해자축
		num1 = 5
		num2 = 5
		num3 = 5
		for i := 0; i < 4; i++ {
			if a.Ji[i].Title == "해" {
				num1 = i
			}
		}
		for j := 0; j < 4; j++ {
			if a.Ji[j].Title == "자" {
				num2 = j
			}
		}
		for k := 0; k < 4; k++ {
			if a.Ji[k].Title == "축" {
				num3 = k
			}
		}

		if num1 != 5 || num2 != 5 || num3 != 5 {
			switch {
			case num1 == 5 && num2 == 5 && num3 == 5:
				{
					a.Result[num1].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[num3].BangHab.Exist = true
					a.Result[num1].BangHab.SaOMi += 1
					a.Result[num2].BangHab.SaOMi += 1
					a.Result[num3].BangHab.SaOMi += 1
				}
			case num1 == 5:
				if num2 == 5 {
					if a.Ji[4].Title == "해" && a.Ji[5].Title == "자" {
						a.Result[4].BangHab.Exist = true
						a.Result[5].BangHab.Exist = true
						a.Result[num3].BangHab.Exist = true
						a.Result[4].BangHab.HaeJaChuk += 1
						a.Result[5].BangHab.HaeJaChuk += 1
						a.Result[num3].BangHab.HaeJaChuk += 1
					} else if a.Ji[4].Title == "자" && a.Ji[5].Title == "해" {
						a.Result[4].BangHab.Exist = true
						a.Result[5].BangHab.Exist = true
						a.Result[num3].BangHab.Exist = true
						a.Result[4].BangHab.HaeJaChuk += 1
						a.Result[5].BangHab.HaeJaChuk += 1
						a.Result[num3].BangHab.HaeJaChuk += 1
					}
				} else if num3 == 5 {
					if a.Ji[4].Title == "해" && a.Ji[5].Title == "축" {
						a.Result[4].BangHab.Exist = true
						a.Result[5].BangHab.Exist = true
						a.Result[num2].BangHab.Exist = true
						a.Result[4].BangHab.HaeJaChuk += 1
						a.Result[5].BangHab.HaeJaChuk += 1
						a.Result[num2].BangHab.HaeJaChuk += 1
					} else if a.Ji[4].Title == "축" && a.Ji[5].Title == "해" {
						a.Result[4].BangHab.Exist = true
						a.Result[5].BangHab.Exist = true
						a.Result[num2].BangHab.Exist = true
						a.Result[4].BangHab.HaeJaChuk += 1
						a.Result[5].BangHab.HaeJaChuk += 1
						a.Result[num2].BangHab.HaeJaChuk += 1
					}
				} else if num2 != 5 && num3 != 5 {
					if a.Ji[4].Title == "해" {
						a.Result[4].BangHab.Exist = true
						a.Result[num2].BangHab.Exist = true
						a.Result[num3].BangHab.Exist = true
						a.Result[4].BangHab.HaeJaChuk += 1
						a.Result[num2].BangHab.HaeJaChuk += 1
						a.Result[num3].BangHab.HaeJaChuk += 1
					}
					if a.Ji[5].Title == "해" {
						a.Result[5].BangHab.Exist = true
						a.Result[num2].BangHab.Exist = true
						a.Result[num3].BangHab.Exist = true
						a.Result[5].BangHab.HaeJaChuk += 1
						a.Result[num2].BangHab.HaeJaChuk += 1
						a.Result[num3].BangHab.HaeJaChuk += 1
					}
				}
				fallthrough

			case num2 == 5:
				if num3 == 5 {
					if a.Ji[4].Title == "자" && a.Ji[5].Title == "축" {
						a.Result[4].BangHab.Exist = true
						a.Result[5].BangHab.Exist = true
						a.Result[num1].BangHab.Exist = true
						a.Result[4].BangHab.HaeJaChuk += 1
						a.Result[5].BangHab.HaeJaChuk += 1
						a.Result[num1].BangHab.HaeJaChuk += 1
					} else if a.Ji[4].Title == "축" && a.Ji[5].Title == "자" {
						a.Result[4].BangHab.Exist = true
						a.Result[5].BangHab.Exist = true
						a.Result[num1].BangHab.Exist = true
						a.Result[4].BangHab.HaeJaChuk += 1
						a.Result[5].BangHab.HaeJaChuk += 1
						a.Result[num1].BangHab.HaeJaChuk += 1
					}
				} else if num1 != 5 && num3 != 5 {
					if a.Ji[4].Title == "자" {
						a.Result[4].BangHab.Exist = true
						a.Result[num1].BangHab.Exist = true
						a.Result[num3].BangHab.Exist = true
						a.Result[4].BangHab.HaeJaChuk += 1
						a.Result[num1].BangHab.HaeJaChuk += 1
						a.Result[num3].BangHab.HaeJaChuk += 1
					} else if a.Ji[5].Title == "자" {
						a.Result[5].BangHab.Exist = true
						a.Result[num1].BangHab.Exist = true
						a.Result[num3].BangHab.Exist = true
						a.Result[5].BangHab.HaeJaChuk += 1
						a.Result[num1].BangHab.HaeJaChuk += 1
						a.Result[num3].BangHab.HaeJaChuk += 1
					}
				}
				fallthrough

			case num3 == 5:
				if num1 != 5 && num2 != 5 {
					if a.Ji[4].Title == "축" {
						a.Result[4].BangHab.Exist = true
						a.Result[num1].BangHab.Exist = true
						a.Result[num2].BangHab.Exist = true
						a.Result[4].BangHab.HaeJaChuk += 1
						a.Result[num1].BangHab.HaeJaChuk += 1
						a.Result[num2].BangHab.HaeJaChuk += 1
					} else if a.Ji[5].Title == "축" {
						a.Result[5].BangHab.Exist = true
						a.Result[num1].BangHab.Exist = true
						a.Result[num2].BangHab.Exist = true
						a.Result[5].BangHab.HaeJaChuk += 1
						a.Result[num1].BangHab.HaeJaChuk += 1
						a.Result[num2].BangHab.HaeJaChuk += 1
					}
				}
			}
		}
	}
}

func Find_Banghab_Goonghab(a Person, b Person) (Person, Person) {
	for i := 1; i < 3; i++ {
		a.Result[i].BangHab.Exist = false
		b.Result[i].BangHab.Exist = false
		//인묘진
		if a.Ji[i].Title == "인" || a.Ji[i].Title == "묘" || a.Ji[i].Title == "진" {
			if b.Ji[i].Title == "인" || b.Ji[i].Title == "묘" || b.Ji[i].Title == "진" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].BangHab.Exist = true
					b.Result[i].BangHab.Exist = true
					a.Result[i].BangHab.InMyoJin += 1
					b.Result[i].BangHab.InMyoJin += 1
				}
			}
		}

		//사오미
		if a.Ji[i].Title == "사" || a.Ji[i].Title == "오" || a.Ji[i].Title == "미" {
			if b.Ji[i].Title == "사" || b.Ji[i].Title == "오" || b.Ji[i].Title == "미" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].BangHab.Exist = true
					b.Result[i].BangHab.Exist = true
					a.Result[i].BangHab.SaOMi += 1
					b.Result[i].BangHab.SaOMi += 1
				}
			}
		}

		//신유술
		if a.Ji[i].Title == "신" || a.Ji[i].Title == "유" || a.Ji[i].Title == "술" {
			if b.Ji[i].Title == "신" || b.Ji[i].Title == "유" || b.Ji[i].Title == "술" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].BangHab.Exist = true
					b.Result[i].BangHab.Exist = true
					a.Result[i].BangHab.SinYuSul += 1
					b.Result[i].BangHab.SinYuSul += 1
				}
			}
		}

		//해자축
		if a.Ji[i].Title == "해" || a.Ji[i].Title == "자" || a.Ji[i].Title == "축" {
			if b.Ji[i].Title == "해" || b.Ji[i].Title == "자" || b.Ji[i].Title == "축" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].BangHab.Exist = true
					b.Result[i].BangHab.Exist = true
					a.Result[i].BangHab.HaeJaChuk += 1
					b.Result[i].BangHab.HaeJaChuk += 1
				}
			}
		}
	}

	return a, b
}

func Find_Samhab(a Person) {
	num1 := 5
	num2 := 5
	num3 := 5
	//신자진
	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "신" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "자" {
			num2 = j
		}
	}
	for k := 0; k < 4; k++ {
		if a.Ji[k].Title == "진" {
			num3 = k
		}
	}

	a.Result[num1].SamHab.Exist = false
	a.Result[num2].SamHab.Exist = false
	a.Result[num3].SamHab.Exist = false
	a.Result[4].SamHab.Exist = false
	a.Result[5].SamHab.Exist = false

	if num1 != 5 || num2 != 5 || num3 != 5 {
		switch {
		case num1 == 5 && num2 == 5 && num3 == 5:
			{
				a.Result[num1].SamHab.Exist = true
				a.Result[num2].SamHab.Exist = true
				a.Result[num3].SamHab.Exist = true
				a.Ji[num1].Properties.Hae += 1
				a.Ji[num2].Properties.Go += 1
				a.Ji[num3].Properties.Ji += 1
				a.Result[num1].SamHab.SinJaJin += 1
				a.Result[num2].SamHab.SinJaJin += 1
				a.Result[num3].SamHab.SinJaJin += 1
			}
		case num1 == 5:
			if num2 == 5 {
				if a.Ji[4].Title == "신" && a.Ji[5].Title == "자" {

					a.Result[4].SamHab.SinJaJin += 1
					a.Result[5].SamHab.SinJaJin += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num3].SamHab.SinJaJin += 1
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
				} else if a.Ji[4].Title == "자" && a.Ji[5].Title == "신" {
					a.Result[4].SamHab.SinJaJin += 1
					a.Result[5].SamHab.SinJaJin += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num3].SamHab.SinJaJin += 1
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
				}

			} else if num3 == 5 {
				if a.Ji[4].Title == "신" && a.Ji[5].Title == "진" {
					a.Result[4].SamHab.SinJaJin += 1
					a.Result[5].SamHab.SinJaJin += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num2].SamHab.SinJaJin += 1
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
				} else if a.Ji[4].Title == "진" && a.Ji[5].Title == "신" {
					a.Result[4].SamHab.SinJaJin += 1
					a.Result[5].SamHab.SinJaJin += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num2].SamHab.SinJaJin += 1
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
				}

			} else if num2 != 5 && num3 != 5 {
				if a.Ji[4].Title == "신" {
					a.Result[4].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.SinJaJin += 1
					a.Ji[num2].Properties.Go += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num2].SamHab.SinJaJin += 1
					a.Result[num3].SamHab.SinJaJin += 1
				} else {
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[5].SamHab.SinJaJin += 1
					a.Ji[num2].Properties.Go += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num2].SamHab.SinJaJin += 1
					a.Result[num3].SamHab.SinJaJin += 1
				}
			}
			fallthrough

		case num2 == 5:
			if num3 == 5 {
				if a.Ji[4].Title == "자" && a.Ji[5].Title == "진" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[4].SamHab.SinJaJin += 1
					a.Result[5].SamHab.SinJaJin += 1
					a.Ji[num1].Properties.Hae += 1
					a.Result[num1].SamHab.SinJaJin += 1
				} else if a.Ji[4].Title == "진" && a.Ji[5].Title == "자" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[4].SamHab.SinJaJin += 1
					a.Result[5].SamHab.SinJaJin += 1
					a.Ji[num1].Properties.Hae += 1
					a.Result[num1].SamHab.SinJaJin += 1
				}
			} else if num1 != 5 && num3 != 5 {
				if a.Ji[4].Title == "자" {
					a.Result[4].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.SinJaJin += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num1].SamHab.SinJaJin += 1
					a.Result[num3].SamHab.SinJaJin += 1
				} else if a.Ji[5].Title == "자" {
					a.Result[4].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[5].SamHab.SinJaJin += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num1].SamHab.SinJaJin += 1
					a.Result[num3].SamHab.SinJaJin += 1
				}
			}
			fallthrough

		case num3 == 5:
			if num1 != 5 && num2 != 5 {
				if a.Ji[4].Title == "진" {
					a.Result[4].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[4].SamHab.SinJaJin += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num1].SamHab.SinJaJin += 1
					a.Result[num2].SamHab.SinJaJin += 1
				} else if a.Ji[5].Title == "진" {
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[5].SamHab.SinJaJin += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num1].SamHab.SinJaJin += 1
					a.Result[num2].SamHab.SinJaJin += 1
				}
			}

		}
	}

	//사유축
	num1 = 5
	num2 = 5
	num3 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "사" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "유" {
			num2 = j
		}
	}
	for k := 0; k < 4; k++ {
		if a.Ji[k].Title == "축" {
			num3 = k
		}
	}

	if num1 != 5 || num2 != 5 || num3 != 5 {
		switch {
		case num1 == 5 && num2 == 5 && num3 == 5:
			{
				a.Result[num1].SamHab.Exist = true
				a.Result[num2].SamHab.Exist = true
				a.Result[num3].SamHab.Exist = true
				a.Ji[num1].Properties.Hae += 1
				a.Ji[num2].Properties.Go += 1
				a.Ji[num3].Properties.Ji += 1
				a.Result[num1].SamHab.SaYuChuk += 1
				a.Result[num2].SamHab.SaYuChuk += 1
				a.Result[num3].SamHab.SaYuChuk += 1
			}
		case num1 == 5:
			if num2 == 5 {
				if a.Ji[4].Title == "사" && a.Ji[5].Title == "유" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.SaYuChuk += 1
					a.Result[5].SamHab.SaYuChuk += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num3].SamHab.SaYuChuk += 1
				} else if a.Ji[4].Title == "유" && a.Ji[5].Title == "사" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.SaYuChuk += 1
					a.Result[5].SamHab.SaYuChuk += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num3].SamHab.SaYuChuk += 1
				}
			} else if num3 == 5 {
				if a.Ji[4].Title == "사" && a.Ji[5].Title == "축" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[4].SamHab.SaYuChuk += 1
					a.Result[5].SamHab.SaYuChuk += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num2].SamHab.SaYuChuk += 1
				} else if a.Ji[4].Title == "축" && a.Ji[5].Title == "사" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[4].SamHab.SaYuChuk += 1
					a.Result[5].SamHab.SaYuChuk += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num2].SamHab.SaYuChuk += 1
				}
			} else if num2 != 5 && num3 != 5 {
				if a.Ji[4].Title == "사" || a.Ji[5].Title == "사" {
					a.Result[4].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.SaYuChuk += 1
					a.Result[5].SamHab.SaYuChuk += 1
					a.Ji[num2].Properties.Go += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num2].SamHab.SaYuChuk += 1
					a.Result[num3].SamHab.SaYuChuk += 1
				}
			}
			fallthrough

		case num2 == 5:
			if num3 == 5 {
				if a.Ji[4].Title == "유" && a.Ji[5].Title == "축" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[4].SamHab.SaYuChuk += 1
					a.Result[5].SamHab.SaYuChuk += 1
					a.Ji[num1].Properties.Hae += 1
					a.Result[num1].SamHab.SaYuChuk += 1
				} else if a.Ji[4].Title == "축" && a.Ji[5].Title == "유" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[4].SamHab.SaYuChuk += 1
					a.Result[5].SamHab.SaYuChuk += 1
					a.Ji[num1].Properties.Hae += 1
					a.Result[num1].SamHab.SaYuChuk += 1
				}
			} else if num1 != 5 && num3 != 5 {
				if a.Ji[4].Title == "유" {
					a.Result[4].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.SaYuChuk += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num1].SamHab.SaYuChuk += 1
					a.Result[num3].SamHab.SaYuChuk += 1
				} else if a.Ji[5].Title == "유" {
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[5].SamHab.SaYuChuk += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num1].SamHab.SaYuChuk += 1
					a.Result[num3].SamHab.SaYuChuk += 1
				}
			}
			fallthrough

		case num3 == 5:
			if num1 != 5 && num2 != 5 {
				if a.Ji[4].Title == "축" {
					a.Result[4].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[4].SamHab.SaYuChuk += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num1].SamHab.SaYuChuk += 1
					a.Result[num2].SamHab.SaYuChuk += 1
				} else if a.Ji[5].Title == "축" {
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[5].SamHab.SaYuChuk += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num1].SamHab.SaYuChuk += 1
					a.Result[num2].SamHab.SaYuChuk += 1
				}
			}

		}
	}

	//해묘미
	num1 = 5
	num2 = 5
	num3 = 5
	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "해" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "묘" {
			num2 = j
		}
	}
	for k := 0; k < 4; k++ {
		if a.Ji[k].Title == "미" {
			num3 = k
		}
	}

	if num1 != 5 || num2 != 5 || num3 != 5 {
		switch {
		case num1 == 5 && num2 == 5 && num3 == 5:
			{

				a.Result[num1].SamHab.Exist = true
				a.Result[num2].SamHab.Exist = true
				a.Result[num3].SamHab.Exist = true
				a.Ji[num1].Properties.Hae += 1
				a.Ji[num2].Properties.Go += 1
				a.Ji[num3].Properties.Ji += 1
				a.Result[num1].SamHab.HaeMyoMi += 1
				a.Result[num2].SamHab.HaeMyoMi += 1
				a.Result[num3].SamHab.HaeMyoMi += 1
			}
		case num1 == 5:
			if num2 == 5 {
				if a.Ji[4].Title == "해" && a.Ji[5].Title == "묘" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.HaeMyoMi += 1
					a.Result[5].SamHab.HaeMyoMi += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num3].SamHab.HaeMyoMi += 1
				} else if a.Ji[4].Title == "묘" && a.Ji[5].Title == "해" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.HaeMyoMi += 1
					a.Result[5].SamHab.HaeMyoMi += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num3].SamHab.HaeMyoMi += 1
				}
			} else if num3 == 5 {
				if a.Ji[4].Title == "해" && a.Ji[5].Title == "미" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[4].SamHab.HaeMyoMi += 1
					a.Result[5].SamHab.HaeMyoMi += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num2].SamHab.HaeMyoMi += 1
				} else if a.Ji[4].Title == "미" && a.Ji[5].Title == "해" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[4].SamHab.HaeMyoMi += 1
					a.Result[5].SamHab.HaeMyoMi += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num2].SamHab.HaeMyoMi += 1
				}
			} else if num2 != 5 && num3 != 5 {
				if a.Ji[4].Title == "해" {
					a.Result[4].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.HaeMyoMi += 1

					a.Ji[num2].Properties.Go += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num2].SamHab.HaeMyoMi += 1
					a.Result[num3].SamHab.HaeMyoMi += 1
				}

				if a.Ji[5].Title == "해" {
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[5].SamHab.HaeMyoMi += 1
					a.Ji[num2].Properties.Go += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num2].SamHab.HaeMyoMi += 1
					a.Result[num3].SamHab.HaeMyoMi += 1
				}
			}
			fallthrough

		case num2 == 5:
			if num3 == 5 {
				if a.Ji[4].Title == "묘" && a.Ji[5].Title == "미" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[4].SamHab.HaeMyoMi += 1
					a.Result[5].SamHab.HaeMyoMi += 1
					a.Ji[num1].Properties.Hae += 1
					a.Result[num1].SamHab.HaeMyoMi += 1
				} else if a.Ji[4].Title == "미" && a.Ji[5].Title == "묘" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[4].SamHab.HaeMyoMi += 1
					a.Result[5].SamHab.HaeMyoMi += 1
					a.Ji[num1].Properties.Hae += 1
					a.Result[num1].SamHab.HaeMyoMi += 1
				}
			} else if num1 != 5 && num3 != 5 {
				if a.Ji[4].Title == "묘" {
					a.Result[4].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.HaeMyoMi += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num1].SamHab.HaeMyoMi += 1
					a.Result[num3].SamHab.HaeMyoMi += 1
				} else if a.Ji[5].Title == "묘" {
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[5].SamHab.HaeMyoMi += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num1].SamHab.HaeMyoMi += 1
					a.Result[num3].SamHab.HaeMyoMi += 1
				}
			}
			fallthrough

		case num3 == 5:
			if num1 != 5 && num2 != 5 {
				if a.Ji[4].Title == "미" {
					a.Result[4].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[4].SamHab.HaeMyoMi += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num1].SamHab.HaeMyoMi += 1
					a.Result[num2].SamHab.HaeMyoMi += 1
				} else if a.Ji[5].Title == "미" {
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[5].SamHab.HaeMyoMi += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num1].SamHab.HaeMyoMi += 1
					a.Result[num2].SamHab.HaeMyoMi += 1
				}
			}

		}

	}

	//인오술
	num1 = 5
	num2 = 5
	num3 = 5
	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "인" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "오" {
			num2 = j
		}
	}
	for k := 0; k < 4; k++ {
		if a.Ji[k].Title == "술" {
			num3 = k
		}
	}

	if num1 != 5 || num2 != 5 || num3 != 5 {
		switch {
		case num1 == 5 && num2 == 5 && num3 == 5:
			{
				a.Result[num1].SamHab.Exist = true
				a.Result[num2].SamHab.Exist = true
				a.Result[num3].SamHab.Exist = true
				a.Ji[num1].Properties.Hae += 1
				a.Ji[num2].Properties.Go += 1
				a.Ji[num3].Properties.Ji += 1
				a.Result[num1].SamHab.InOSul += 1
				a.Result[num2].SamHab.InOSul += 1
				a.Result[num3].SamHab.InOSul += 1
			}
		case num1 == 5:
			if num2 == 5 {
				if a.Ji[4].Title == "인" && a.Ji[5].Title == "오" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.InOSul += 1
					a.Result[5].SamHab.InOSul += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num3].SamHab.InOSul += 1
				} else if a.Ji[4].Title == "오" && a.Ji[5].Title == "인" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.InOSul += 1
					a.Result[5].SamHab.InOSul += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num3].SamHab.InOSul += 1
				}
			} else if num3 == 5 {
				if a.Ji[4].Title == "인" && a.Ji[5].Title == "술" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[4].SamHab.InOSul += 1
					a.Result[5].SamHab.InOSul += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num2].SamHab.InOSul += 1
				} else if a.Ji[4].Title == "술" && a.Ji[5].Title == "인" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[4].SamHab.InOSul += 1
					a.Result[5].SamHab.InOSul += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num2].SamHab.InOSul += 1
				}
			} else if num2 != 5 && num3 != 5 {
				if a.Ji[4].Title == "인" {
					a.Result[4].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.InOSul += 1
					a.Ji[num2].Properties.Go += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num2].SamHab.InOSul += 1
					a.Result[num3].SamHab.InOSul += 1
				}
				if a.Ji[5].Title == "인" {
					a.Result[5].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[5].SamHab.InOSul += 1
					a.Ji[num2].Properties.Go += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num2].SamHab.InOSul += 1
					a.Result[num3].SamHab.InOSul += 1
				}
			}
			fallthrough

		case num2 == 5:
			if num3 == 5 {
				if a.Ji[4].Title == "오" && a.Ji[5].Title == "술" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[4].SamHab.InOSul += 1
					a.Result[5].SamHab.InOSul += 1
					a.Ji[num1].Properties.Hae += 1
					a.Result[num1].SamHab.InOSul += 1
				} else if a.Ji[4].Title == "술" && a.Ji[5].Title == "오" {
					a.Result[4].SamHab.Exist = true
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[4].SamHab.InOSul += 1
					a.Result[5].SamHab.InOSul += 1
					a.Ji[num1].Properties.Hae += 1
					a.Result[num1].SamHab.InOSul += 1
				}
			} else if num1 != 5 && num3 != 5 {
				if a.Ji[4].Title == "오" {
					a.Result[4].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[4].SamHab.InOSul += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num1].SamHab.InOSul += 1
					a.Result[num3].SamHab.InOSul += 1
				} else if a.Ji[5].Title == "오" {
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num3].SamHab.Exist = true
					a.Result[5].SamHab.InOSul += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num3].Properties.Ji += 1
					a.Result[num1].SamHab.InOSul += 1
					a.Result[num3].SamHab.InOSul += 1
				}
			}
			fallthrough

		case num3 == 5:
			if num1 != 5 && num2 != 5 {
				if a.Ji[4].Title == "술" {
					a.Result[4].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[4].SamHab.InOSul += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num1].SamHab.InOSul += 1
					a.Result[num2].SamHab.InOSul += 1
				} else if a.Ji[5].Title == "술" {
					a.Result[5].SamHab.Exist = true
					a.Result[num1].SamHab.Exist = true
					a.Result[num2].SamHab.Exist = true
					a.Result[5].SamHab.InOSul += 1
					a.Ji[num1].Properties.Hae += 1
					a.Ji[num2].Properties.Go += 1
					a.Result[num1].SamHab.InOSul += 1
					a.Result[num2].SamHab.InOSul += 1
				}
			}

		}

	}
}

func Find_Samhab_Goonghab(a Person, b Person) (Person, Person) {
	for i := 1; i < 3; i++ {
		a.Result[i].SamHab.Exist = false
		b.Result[i].SamHab.Exist = false
		//신자진
		if a.Ji[i].Title == "신" || a.Ji[i].Title == "자" || a.Ji[i].Title == "진" {
			if b.Ji[i].Title == "신" || b.Ji[i].Title == "자" || b.Ji[i].Title == "진" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].SamHab.Exist = true
					b.Result[i].SamHab.Exist = true
					a.Result[i].SamHab.SinJaJin += 1
					b.Result[i].SamHab.SinJaJin += 1
				}
			}
		}

		//사유축
		if a.Ji[i].Title == "사" || a.Ji[i].Title == "유" || a.Ji[i].Title == "축" {
			if b.Ji[i].Title == "사" || b.Ji[i].Title == "유" || b.Ji[i].Title == "축" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].SamHab.Exist = true
					b.Result[i].SamHab.Exist = true
					a.Result[i].SamHab.SaYuChuk += 1
					b.Result[i].SamHab.SaYuChuk += 1
				}
			}
		}

		//해묘미
		if a.Ji[i].Title == "해" || a.Ji[i].Title == "묘" || a.Ji[i].Title == "미" {
			if b.Ji[i].Title == "해" || b.Ji[i].Title == "묘" || b.Ji[i].Title == "미" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].SamHab.Exist = true
					b.Result[i].SamHab.Exist = true
					a.Result[i].SamHab.HaeMyoMi += 1
					b.Result[i].SamHab.HaeMyoMi += 1
				}
			}
		}

		//인오술
		if a.Ji[i].Title == "인" || a.Ji[i].Title == "오" || a.Ji[i].Title == "술" {
			if b.Ji[i].Title == "인" || b.Ji[i].Title == "오" || b.Ji[i].Title == "술" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].SamHab.Exist = true
					b.Result[i].SamHab.Exist = true
					a.Result[i].SamHab.InOSul += 1
					b.Result[i].SamHab.InOSul += 1
				}
			}
		}
	}

	return a, b
}

func Find_Yukhab(a Person) {
	//축자
	num1 := 5
	num2 := 5
	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "축" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "자" {
			num2 = j
		}
	}
	a.Result[num1].YukHab.Exist = false
	a.Result[num2].YukHab.Exist = false
	a.Result[4].YukHab.Exist = false
	a.Result[5].YukHab.Exist = false
	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].YukHab.Exist = true
			a.Result[num2].YukHab.Exist = true
			a.Ji[num1].Properties.Prop = "earth"
			a.Ji[num2].Properties.Prop = "earth"
			a.Result[num1].YukHab.JaChuk += 1
			a.Result[num2].YukHab.JaChuk += 1
		case num1 == 5:
			if a.Ji[4].Title == "축" {
				a.Result[4].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[4].YukHab.JaChuk += 1
				a.Ji[num2].Properties.Prop = "earth"
				a.Result[num2].YukHab.JaChuk += 1
			} else if a.Ji[5].Title == "축" {
				a.Result[5].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[5].YukHab.JaChuk += 1
				a.Ji[num2].Properties.Prop = "earth"
				a.Result[num2].YukHab.JaChuk += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "자" {
				a.Result[4].YukHab.Exist = true
				a.Result[num1].YukHab.Exist = true
				a.Result[num1].YukHab.JaChuk += 1
				a.Ji[num1].Properties.Prop = "earth"
				a.Result[4].YukHab.JaChuk += 1
			} else if a.Ji[5].Title == "자" {
				a.Result[5].YukHab.Exist = true
				a.Result[num1].YukHab.Exist = true
				a.Result[num1].YukHab.JaChuk += 1
				a.Ji[num1].Properties.Prop = "earth"
				a.Result[5].YukHab.JaChuk += 1
			}
		}
	}

	//인해
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "인" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "해" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].YukHab.Exist = true
			a.Result[num2].YukHab.Exist = true
			a.Ji[num1].Properties.Prop = "tree"
			a.Ji[num2].Properties.Prop = "tree"
			a.Result[num1].YukHab.InHye += 1
			a.Result[num2].YukHab.InHye += 1
		case num1 == 5:
			if a.Ji[4].Title == "인" {
				a.Result[4].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[4].YukHab.InHye += 1
				a.Ji[num2].Properties.Prop = "tree"
				a.Result[num2].YukHab.InHye += 1
			} else if a.Ji[5].Title == "인" {
				a.Result[5].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[5].YukHab.InHye += 1
				a.Ji[num2].Properties.Prop = "tree"
				a.Result[num2].YukHab.InHye += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "해" {
				a.Result[4].YukHab.Exist = true
				a.Result[num1].YukHab.Exist = true
				a.Result[num1].YukHab.InHye += 1
				a.Ji[num1].Properties.Prop = "tree"
				a.Result[4].YukHab.InHye += 1
			} else if a.Ji[5].Title == "해" {
				a.Result[5].YukHab.Exist = true
				a.Result[num1].YukHab.Exist = true
				a.Result[num1].YukHab.InHye += 1
				a.Ji[num1].Properties.Prop = "tree"
				a.Result[5].YukHab.InHye += 1
			}
		}
	}

	//묘술
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "묘" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "술" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].YukHab.Exist = true
			a.Result[num2].YukHab.Exist = true
			a.Ji[num1].Properties.Prop = "fire"
			a.Ji[num2].Properties.Prop = "fire"
			a.Result[num1].YukHab.MyoSul += 1
			a.Result[num2].YukHab.MyoSul += 1
		case num1 == 5:
			if a.Ji[4].Title == "묘" {
				a.Result[4].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[4].YukHab.MyoSul += 1
				a.Ji[num2].Properties.Prop = "fire"
				a.Result[num2].YukHab.MyoSul += 1
			} else if a.Ji[5].Title == "묘" {
				a.Result[5].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[5].YukHab.MyoSul += 1
				a.Ji[num2].Properties.Prop = "fire"
				a.Result[num2].YukHab.MyoSul += 1
			}
			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "술" {
				a.Result[4].YukHab.Exist = true
				a.Result[num1].YukHab.Exist = true
				a.Result[num1].YukHab.MyoSul += 1
				a.Ji[num1].Properties.Prop = "fire"
				a.Result[4].YukHab.MyoSul += 1
			} else if a.Ji[5].Title == "술" {
				a.Result[5].YukHab.Exist = true
				a.Result[num1].YukHab.Exist = true
				a.Result[num1].YukHab.MyoSul += 1
				a.Ji[num1].Properties.Prop = "fire"
				a.Result[5].YukHab.MyoSul += 1
			}

		}
	}

	//진유
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "진" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "유" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].YukHab.Exist = true
			a.Result[num2].YukHab.Exist = true
			a.Ji[num1].Properties.Prop = "iron"
			a.Ji[num2].Properties.Prop = "iron"
			a.Result[num1].YukHab.JinYu += 1
			a.Result[num2].YukHab.JinYu += 1
		case num1 == 5:
			if a.Ji[4].Title == "진" {
				a.Result[4].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[4].YukHab.JinYu += 1
				a.Ji[num2].Properties.Prop = "iron"
				a.Result[num2].YukHab.JinYu += 1
			} else if a.Ji[5].Title == "진" {
				a.Result[5].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[5].YukHab.JinYu += 1
				a.Ji[num2].Properties.Prop = "iron"
				a.Result[num2].YukHab.JinYu += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "유" {
				a.Result[num1].YukHab.Exist = true
				a.Result[4].YukHab.Exist = true
				a.Result[num1].YukHab.JinYu += 1
				a.Ji[num1].Properties.Prop = "iron"
				a.Result[4].YukHab.JinYu += 1
			} else if a.Ji[5].Title == "유" {
				a.Result[num1].YukHab.Exist = true
				a.Result[5].YukHab.Exist = true
				a.Result[num1].YukHab.JinYu += 1
				a.Ji[num1].Properties.Prop = "iron"
				a.Result[5].YukHab.JinYu += 1
			}
		}
	}

	//사신
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "사" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "신" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].YukHab.Exist = true
			a.Result[num2].YukHab.Exist = true
			a.Ji[num1].Properties.Prop = "water"
			a.Ji[num2].Properties.Prop = "water"
			a.Result[num1].YukHab.SaSin += 1
			a.Result[num2].YukHab.SaSin += 1
		case num1 == 5:
			if a.Ji[4].Title == "사" {
				a.Result[4].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[4].YukHab.SaSin += 1
				a.Ji[num2].Properties.Prop = "water"
				a.Result[num2].YukHab.SaSin += 1
			} else if a.Ji[5].Title == "사" {
				a.Result[5].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[5].YukHab.SaSin += 1
				a.Ji[num2].Properties.Prop = "water"
				a.Result[num2].YukHab.SaSin += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "신" {
				a.Result[num1].YukHab.Exist = true
				a.Result[4].YukHab.Exist = true
				a.Result[num1].YukHab.SaSin += 1
				a.Ji[num1].Properties.Prop = "water"
				a.Result[4].YukHab.SaSin += 1
			} else if a.Ji[5].Title == "신" {
				a.Result[num1].YukHab.Exist = true
				a.Result[5].YukHab.Exist = true
				a.Result[num1].YukHab.SaSin += 1
				a.Ji[num1].Properties.Prop = "water"
				a.Result[5].YukHab.SaSin += 1
			}
		}
	}

	//오미
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "오" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "미" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].YukHab.Exist = true
			a.Result[num2].YukHab.Exist = true
			a.Ji[num1].Properties.Prop = "fire"
			a.Ji[num2].Properties.Prop = "fire"
			a.Result[num1].YukHab.OMi += 1
			a.Result[num2].YukHab.OMi += 1
		case num1 == 5:
			if a.Ji[4].Title == "오" {
				a.Result[4].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[4].YukHab.OMi += 1
				a.Ji[num2].Properties.Prop = "fire"
				a.Result[num2].YukHab.OMi += 1
			} else if a.Ji[5].Title == "오" {
				a.Result[5].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[5].YukHab.OMi += 1
				a.Ji[num2].Properties.Prop = "fire"
				a.Result[num2].YukHab.OMi += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "미" {
				a.Result[4].YukHab.Exist = true
				a.Result[num2].YukHab.Exist = true
				a.Result[num1].YukHab.OMi += 1
				a.Ji[num1].Properties.Prop = "fire"
				a.Result[4].YukHab.OMi += 1
			} else if a.Ji[5].Title == "미" {
				a.Result[num1].YukHab.Exist = true
				a.Result[5].YukHab.Exist = true
				a.Result[num1].YukHab.OMi += 1
				a.Ji[num1].Properties.Prop = "fire"
				a.Result[5].YukHab.OMi += 1
			}
		}
	}

}

func Find_Yukhab_Goonghab(a Person, b Person) (Person, Person) {
	for i := 1; i < 3; i++ {
		a.Result[i].YukHab.Exist = false
		b.Result[i].YukHab.Exist = false
		//축자
		if a.Ji[i].Title == "자" || a.Ji[i].Title == "축" {
			if b.Ji[i].Title == "자" || b.Ji[i].Title == "축" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].YukHab.Exist = true
					b.Result[i].YukHab.Exist = true
					a.Ji[i].Properties.Prop = "earth"
					a.Result[i].YukHab.JaChuk += 1
					b.Ji[i].Properties.Prop = "earth"
					b.Result[i].YukHab.JaChuk += 1
				}
			}
		}

		//인해
		if a.Ji[i].Title == "인" || a.Ji[i].Title == "해" {
			if b.Ji[i].Title == "인" || b.Ji[i].Title == "해" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].YukHab.Exist = true
					b.Result[i].YukHab.Exist = true
					a.Ji[i].Properties.Prop = "tree"
					b.Ji[i].Properties.Prop = "tree"
					a.Result[i].YukHab.InHye += 1
					b.Result[i].YukHab.InHye += 1
				}
			}
		}

		//진유
		if a.Ji[i].Title == "진" || a.Ji[i].Title == "유" {
			if b.Ji[i].Title == "진" || b.Ji[i].Title == "유" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].YukHab.Exist = true
					b.Result[i].YukHab.Exist = true
					a.Ji[i].Properties.Prop = "iron"
					b.Ji[i].Properties.Prop = "iron"
					a.Result[i].YukHab.JinYu += 1
					b.Result[i].YukHab.JinYu += 1
				}
			}
		}

		//사신
		if a.Ji[i].Title == "사" || a.Ji[i].Title == "신" {
			if b.Ji[i].Title == "사" || b.Ji[i].Title == "신" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].YukHab.Exist = true
					b.Result[i].YukHab.Exist = true
					a.Ji[i].Properties.Prop = "water"
					b.Ji[i].Properties.Prop = "water"
					a.Result[i].YukHab.SaSin += 1
					b.Result[i].YukHab.SaSin += 1
				}
			}
		}

		//오미
		if a.Ji[i].Title == "오" || a.Ji[i].Title == "미" {
			if b.Ji[i].Title == "오" || b.Ji[i].Title == "미" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].YukHab.Exist = true
					b.Result[i].YukHab.Exist = true
					a.Ji[i].Properties.Prop = "fire"
					b.Ji[i].Properties.Prop = "fire"
					a.Result[i].YukHab.OMi += 1
					b.Result[i].YukHab.OMi += 1
				}
			}
		}
	}

	return a, b
}

func Find_Hyungsal(a Person) {

	//인사
	num1 := 5
	num2 := 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "인" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "사" {
			num2 = j
		}
	}

	a.Result[num1].Hyung.Exist = false
	a.Result[num2].Hyung.Exist = false

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hyung.Exist = true
			a.Result[num2].Hyung.Exist = true
			a.Result[num1].Hyung.InSa += 1
			a.Result[num2].Hyung.InSa += 1
		case num1 == 5:
			if a.Ji[4].Title == "인" {
				a.Result[4].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[4].Hyung.InSa += 1
				a.Result[num2].Hyung.InSa += 1
			} else if a.Ji[5].Title == "인" {
				a.Result[5].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[5].Hyung.InSa += 1
				a.Result[num2].Hyung.InSa += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "사" {
				a.Result[num1].Hyung.Exist = true
				a.Result[4].Hyung.Exist = true
				a.Result[num1].Hyung.InSa += 1
				a.Result[4].Hyung.InSa += 1
			} else if a.Ji[5].Title == "사" {
				a.Result[num1].Hyung.Exist = true
				a.Result[5].Hyung.Exist = true
				a.Result[num1].Hyung.InSa += 1
				a.Result[5].Hyung.InSa += 1
			}
		}
	}
	//사신
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "사" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "신" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hyung.Exist = true
			a.Result[num2].Hyung.Exist = true
			a.Result[num1].Hyung.SaSin += 1
			a.Result[num2].Hyung.SaSin += 1
		case num1 == 5:
			if a.Ji[4].Title == "사" {
				a.Result[4].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[4].Hyung.SaSin += 1
				a.Result[num2].Hyung.SaSin += 1
			} else if a.Ji[5].Title == "사" {
				a.Result[5].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[5].Hyung.SaSin += 1
				a.Result[num2].Hyung.SaSin += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "신" {
				a.Result[num1].Hyung.Exist = true
				a.Result[4].Hyung.Exist = true
				a.Result[num1].Hyung.SaSin += 1
				a.Result[4].Hyung.SaSin += 1
			} else if a.Ji[5].Title == "신" {
				a.Result[num1].Hyung.Exist = true
				a.Result[5].Hyung.Exist = true
				a.Result[num1].Hyung.SaSin += 1
				a.Result[5].Hyung.SaSin += 1
			}
		}
	}

	//축술
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "축" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "술" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hyung.Exist = true
			a.Result[num2].Hyung.Exist = true
			a.Result[num1].Hyung.ChukSul += 1
			a.Result[num2].Hyung.ChukSul += 1
		case num1 == 5:
			if a.Ji[4].Title == "축" {
				a.Result[4].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[4].Hyung.ChukSul += 1
				a.Result[num2].Hyung.ChukSul += 1
			} else if a.Ji[5].Title == "축" {
				a.Result[5].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[5].Hyung.ChukSul += 1
				a.Result[num2].Hyung.ChukSul += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "술" {
				a.Result[num1].Hyung.Exist = true
				a.Result[4].Hyung.Exist = true
				a.Result[num1].Hyung.ChukSul += 1
				a.Result[4].Hyung.ChukSul += 1
			} else if a.Ji[5].Title == "술" {
				a.Result[num1].Hyung.Exist = true
				a.Result[5].Hyung.Exist = true
				a.Result[num1].Hyung.ChukSul += 1
				a.Result[5].Hyung.ChukSul += 1
			}
		}
	}

	//술미
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "술" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "미" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hyung.Exist = true
			a.Result[num2].Hyung.Exist = true
			a.Result[num1].Hyung.SulMi += 1
			a.Result[num2].Hyung.SulMi += 1
		case num1 == 5:
			if a.Ji[4].Title == "술" {
				a.Result[4].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[4].Hyung.SulMi += 1
				a.Result[num2].Hyung.SulMi += 1
			} else if a.Ji[5].Title == "술" {
				a.Result[5].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[5].Hyung.SulMi += 1
				a.Result[num2].Hyung.SulMi += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "미" {
				a.Result[num1].Hyung.Exist = true
				a.Result[4].Hyung.Exist = true
				a.Result[num1].Hyung.SulMi += 1
				a.Result[4].Hyung.SulMi += 1
			} else if a.Ji[5].Title == "미" {
				a.Result[num1].Hyung.Exist = true
				a.Result[5].Hyung.Exist = true
				a.Result[num1].Hyung.SulMi += 1
				a.Result[5].Hyung.SulMi += 1
			}
		}
	}

	//자묘
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "자" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "묘" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hyung.Exist = true
			a.Result[num2].Hyung.Exist = true
			a.Result[num1].Hyung.JaMyo += 1
			a.Result[num2].Hyung.JaMyo += 1
		case num1 == 5:
			if a.Ji[4].Title == "자" {
				a.Result[4].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[4].Hyung.JaMyo += 1
				a.Result[num2].Hyung.JaMyo += 1
			} else if a.Ji[5].Title == "자" {
				a.Result[5].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[5].Hyung.JaMyo += 1
				a.Result[num2].Hyung.JaMyo += 1
			}
			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "묘" {
				a.Result[4].Hyung.Exist = true
				a.Result[num1].Hyung.Exist = true
				a.Result[num1].Hyung.JaMyo += 1
				a.Result[4].Hyung.JaMyo += 1
			} else if a.Ji[5].Title == "묘" {
				a.Result[5].Hyung.Exist = true
				a.Result[num1].Hyung.Exist = true
				a.Result[num1].Hyung.JaMyo += 1
				a.Result[5].Hyung.JaMyo += 1
			}
		}
	}

	//진진
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "진" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "진" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hyung.Exist = true
			a.Result[num2].Hyung.Exist = true
			a.Result[num1].Hyung.JinJin += 1
			a.Result[num2].Hyung.JinJin += 1
		case num1 == 5:
			if a.Ji[4].Title == "진" {
				a.Result[4].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[4].Hyung.JinJin += 1
				a.Result[num2].Hyung.JinJin += 1
			} else if a.Ji[5].Title == "진" {
				a.Result[5].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[5].Hyung.JinJin += 1
				a.Result[num2].Hyung.JinJin += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "진" {
				a.Result[num1].Hyung.Exist = true
				a.Result[4].Hyung.Exist = true
				a.Result[num1].Hyung.JinJin += 1
				a.Result[4].Hyung.JinJin += 1
			} else if a.Ji[5].Title == "진" {
				a.Result[num1].Hyung.Exist = true
				a.Result[5].Hyung.Exist = true
				a.Result[num1].Hyung.JinJin += 1
				a.Result[5].Hyung.JinJin += 1
			}
		}
	}

	//오오
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "오" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "오" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hyung.Exist = true
			a.Result[num2].Hyung.Exist = true
			a.Result[num1].Hyung.OO += 1
			a.Result[num2].Hyung.OO += 1
		case num1 == 5:
			if a.Ji[4].Title == "오" {
				a.Result[4].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[4].Hyung.OO += 1
				a.Result[num2].Hyung.OO += 1
			} else if a.Ji[5].Title == "오" {
				a.Result[5].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[5].Hyung.OO += 1
				a.Result[num2].Hyung.OO += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "오" {
				a.Result[num1].Hyung.Exist = true
				a.Result[4].Hyung.Exist = true
				a.Result[num1].Hyung.OO += 1
				a.Result[4].Hyung.OO += 1
			} else if a.Ji[5].Title == "오" {
				a.Result[num1].Hyung.Exist = true
				a.Result[5].Hyung.Exist = true
				a.Result[num1].Hyung.OO += 1
				a.Result[5].Hyung.OO += 1
			}
		}
	}

	//유유
	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "유" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "유" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hyung.Exist = true
			a.Result[num2].Hyung.Exist = true
			a.Result[num1].Hyung.YuYu += 1
			a.Result[num2].Hyung.YuYu += 1
		case num1 == 5:
			if a.Ji[4].Title == "유" {
				a.Result[4].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[4].Hyung.YuYu += 1
				a.Result[num2].Hyung.YuYu += 1
			} else if a.Ji[5].Title == "유" {
				a.Result[5].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[5].Hyung.YuYu += 1
				a.Result[num2].Hyung.YuYu += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "유" {
				a.Result[num1].Hyung.Exist = true
				a.Result[4].Hyung.Exist = true
				a.Result[num1].Hyung.YuYu += 1
				a.Result[4].Hyung.YuYu += 1
			} else if a.Ji[5].Title == "유" {
				a.Result[num1].Hyung.Exist = true
				a.Result[5].Hyung.Exist = true
				a.Result[num1].Hyung.YuYu += 1
				a.Result[5].Hyung.YuYu += 1
			}
		}
	}

	//해해
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "해" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "해" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hyung.Exist = true
			a.Result[num2].Hyung.Exist = true
			a.Result[num1].Hyung.HaeHae += 1
			a.Result[num2].Hyung.HaeHae += 1
		case num1 == 5:
			if a.Ji[4].Title == "해" {
				a.Result[4].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[4].Hyung.HaeHae += 1
				a.Result[num2].Hyung.HaeHae += 1
			} else if a.Ji[5].Title == "해" {
				a.Result[5].Hyung.Exist = true
				a.Result[num2].Hyung.Exist = true
				a.Result[5].Hyung.HaeHae += 1
				a.Result[num2].Hyung.HaeHae += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "해" {
				a.Result[num1].Hyung.Exist = true
				a.Result[4].Hyung.Exist = true
				a.Result[num1].Hyung.HaeHae += 1
				a.Result[4].Hyung.HaeHae += 1
			} else if a.Ji[5].Title == "해" {
				a.Result[num1].Hyung.Exist = true
				a.Result[5].Hyung.Exist = true
				a.Result[num1].Hyung.HaeHae += 1
				a.Result[5].Hyung.HaeHae += 1
			}

		}
	}
}

func Find_Hyungsal_Goonghab(a Person, b Person) (Person, Person) {

	for i := 1; i < 3; i++ {
		a.Result[i].Hyung.Exist = false
		b.Result[i].Hyung.Exist = false
		//인사
		if a.Ji[i].Title == "사" || a.Ji[i].Title == "인" {
			if b.Ji[i].Title == "사" || b.Ji[i].Title == "인" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hyung.Exist = true
					b.Result[i].Hyung.Exist = true
					a.Result[i].Hyung.InSa += 1
					b.Result[i].Hyung.InSa += 1
				}
			}
		}

		//사신
		if a.Ji[i].Title == "사" || a.Ji[i].Title == "신" {
			if b.Ji[i].Title == "사" || b.Ji[i].Title == "신" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hyung.Exist = true
					b.Result[i].Hyung.Exist = true
					a.Result[i].Hyung.SaSin += 1
					b.Result[i].Hyung.SaSin += 1
				}
			}
		}

		//축술
		if a.Ji[i].Title == "술" || a.Ji[i].Title == "축" {
			if b.Ji[i].Title == "술" || b.Ji[i].Title == "축" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hyung.Exist = true
					b.Result[i].Hyung.Exist = true
					a.Result[i].Hyung.ChukSul += 1
					b.Result[i].Hyung.ChukSul += 1
				}
			}
		}

		//술미
		if a.Ji[i].Title == "술" || a.Ji[i].Title == "미" {
			if b.Ji[i].Title == "술" || b.Ji[i].Title == "미" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hyung.Exist = true
					b.Result[i].Hyung.Exist = true
					a.Result[i].Hyung.SulMi += 1
					b.Result[i].Hyung.SulMi += 1
				}
			}
		}

		//자묘
		if a.Ji[i].Title == "자" || a.Ji[i].Title == "묘" {
			if b.Ji[i].Title == "자" || b.Ji[i].Title == "묘" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hyung.Exist = true
					b.Result[i].Hyung.Exist = true
					a.Result[i].Hyung.JaMyo += 1
					b.Result[i].Hyung.JaMyo += 1
				}
			}
		}

		//진진
		if a.Ji[i].Title == "진" {
			if b.Ji[i].Title == "진" {
				a.Result[i].Hyung.Exist = true
				b.Result[i].Hyung.Exist = true
				a.Result[i].Hyung.JinJin += 1
				b.Result[i].Hyung.JinJin += 1
			}
		}

		//오오
		if a.Ji[i].Title == "오" {
			if b.Ji[i].Title == "오" {
				a.Result[i].Hyung.Exist = true
				b.Result[i].Hyung.Exist = true
				a.Result[i].Hyung.OO += 1
				b.Result[i].Hyung.OO += 1
			}
		}

		//유유
		if a.Ji[i].Title == "유" {
			if b.Ji[i].Title == "유" {
				a.Result[i].Hyung.Exist = true
				b.Result[i].Hyung.Exist = true
				a.Result[i].Hyung.YuYu += 1
				b.Result[i].Hyung.YuYu += 1
			}
		}

		//해해
		if a.Ji[i].Title == "해" {
			if b.Ji[i].Title == "해" {
				a.Result[i].Hyung.Exist = true
				b.Result[i].Hyung.Exist = true
				a.Result[i].Hyung.HaeHae += 1
				b.Result[i].Hyung.HaeHae += 1
			}
		}
	}

	return a, b
}

func Find_Choongsal(a Person) {
	//인신
	num1 := 5
	num2 := 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "인" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "신" {
			num2 = j
		}
	}
	a.Result[num1].Choong.Exist = false
	a.Result[num2].Choong.Exist = false

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Choong.Exist = true
			a.Result[num2].Choong.Exist = true
			a.Result[num1].Choong.InSin += 1
			a.Result[num2].Choong.InSin += 1
		case num1 == 5:
			if a.Ji[4].Title == "인" {
				a.Result[4].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[4].Choong.InSin += 1
				a.Result[num2].Choong.InSin += 1
			} else if a.Ji[5].Title == "인" {
				a.Result[5].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[5].Choong.InSin += 1
				a.Result[num2].Choong.InSin += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "신" {
				a.Result[num1].Choong.Exist = true
				a.Result[4].Choong.Exist = true
				a.Result[num1].Choong.InSin += 1
				a.Result[4].Choong.InSin += 1
			} else if a.Ji[5].Title == "신" {
				a.Result[num1].Choong.Exist = true
				a.Result[5].Choong.Exist = true
				a.Result[num1].Choong.InSin += 1
				a.Result[5].Choong.InSin += 1
			}
		}
	}
	//묘유
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "묘" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "유" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Choong.Exist = true
			a.Result[num2].Choong.Exist = true
			a.Result[num1].Choong.MyoYu += 1
			a.Result[num2].Choong.MyoYu += 1
		case num1 == 5:
			if a.Ji[4].Title == "묘" {
				a.Result[4].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[4].Choong.MyoYu += 1
				a.Result[num2].Choong.MyoYu += 1
			} else if a.Ji[5].Title == "묘" {
				a.Result[5].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[5].Choong.MyoYu += 1
				a.Result[num2].Choong.MyoYu += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "유" {
				a.Result[num1].Choong.Exist = true
				a.Result[4].Choong.Exist = true
				a.Result[num1].Choong.MyoYu += 1
				a.Result[4].Choong.MyoYu += 1
			} else if a.Ji[5].Title == "유" {
				a.Result[num1].Choong.Exist = true
				a.Result[5].Choong.Exist = true
				a.Result[num1].Choong.MyoYu += 1
				a.Result[5].Choong.MyoYu += 1
			}

		}
	}

	//진술
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "진" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "술" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Choong.Exist = true
			a.Result[num2].Choong.Exist = true
			a.Result[num1].Choong.JinSul += 1
			a.Result[num2].Choong.JinSul += 1
		case num1 == 5:
			if a.Ji[4].Title == "진" {
				a.Result[4].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[4].Choong.JinSul += 1
				a.Result[num2].Choong.JinSul += 1
			} else if a.Ji[5].Title == "진" {
				a.Result[5].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[5].Choong.JinSul += 1
				a.Result[num2].Choong.JinSul += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "술" {
				a.Result[num1].Choong.Exist = true
				a.Result[4].Choong.Exist = true
				a.Result[num1].Choong.JinSul += 1
				a.Result[4].Choong.JinSul += 1
			} else if a.Ji[5].Title == "술" {
				a.Result[num1].Choong.Exist = true
				a.Result[5].Choong.Exist = true
				a.Result[num1].Choong.JinSul += 1
				a.Result[5].Choong.JinSul += 1
			}
		}
	}

	//사해
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "사" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "해" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Choong.Exist = true
			a.Result[num2].Choong.Exist = true
			a.Result[num1].Choong.SaHae += 1
			a.Result[num2].Choong.SaHae += 1
		case num1 == 5:
			if a.Ji[4].Title == "사" {
				a.Result[4].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[4].Choong.SaHae += 1
				a.Result[num2].Choong.SaHae += 1
			} else if a.Ji[5].Title == "사" {
				a.Result[5].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[5].Choong.SaHae += 1
				a.Result[num2].Choong.SaHae += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "해" {
				a.Result[num1].Choong.Exist = true
				a.Result[4].Choong.Exist = true
				a.Result[num1].Choong.SaHae += 1
				a.Result[4].Choong.SaHae += 1
			} else if a.Ji[5].Title == "해" {
				a.Result[num1].Choong.Exist = true
				a.Result[5].Choong.Exist = true
				a.Result[num1].Choong.SaHae += 1
				a.Result[5].Choong.SaHae += 1
			}
		}
	}
	//자오
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "자" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "오" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Choong.Exist = true
			a.Result[num2].Choong.Exist = true
			a.Result[num1].Choong.JaO += 1
			a.Result[num2].Choong.JaO += 1
		case num1 == 5:
			if a.Ji[4].Title == "자" {
				a.Result[4].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[4].Choong.JaO += 1
				a.Result[num2].Choong.JaO += 1
			} else if a.Ji[5].Title == "자" {
				a.Result[5].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[5].Choong.JaO += 1
				a.Result[num2].Choong.JaO += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "오" {
				a.Result[num1].Choong.Exist = true
				a.Result[4].Choong.Exist = true
				a.Result[num1].Choong.JaO += 1
				a.Result[4].Choong.JaO += 1
			} else if a.Ji[5].Title == "오" {
				a.Result[num1].Choong.Exist = true
				a.Result[5].Choong.Exist = true
				a.Result[num1].Choong.JaO += 1
				a.Result[5].Choong.JaO += 1
			}

		}
	}

	//축미
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "축" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "미" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Choong.Exist = true
			a.Result[num2].Choong.Exist = true
			a.Result[num1].Choong.ChukMi += 1
			a.Result[num2].Choong.ChukMi += 1
		case num1 == 5:
			if a.Ji[4].Title == "축" {
				a.Result[4].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[4].Choong.ChukMi += 1
				a.Result[num2].Choong.ChukMi += 1
			} else if a.Ji[5].Title == "축" {
				a.Result[5].Choong.Exist = true
				a.Result[num2].Choong.Exist = true
				a.Result[5].Choong.ChukMi += 1
				a.Result[num2].Choong.ChukMi += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "미" {
				a.Result[num1].Choong.Exist = true
				a.Result[4].Choong.Exist = true
				a.Result[num1].Choong.ChukMi += 1
				a.Result[4].Choong.ChukMi += 1
			} else if a.Ji[5].Title == "미" {
				a.Result[num1].Choong.Exist = true
				a.Result[5].Choong.Exist = true
				a.Result[num1].Choong.ChukMi += 1
				a.Result[5].Choong.ChukMi += 1
			}
		}
	}
}

func Find_Choongsal_Goonghab(a Person, b Person) (Person, Person) {
	for i := 1; i < 3; i++ {
		a.Result[i].Choong.Exist = false
		b.Result[i].Choong.Exist = false
		//인신
		if a.Ji[i].Title == "인" || a.Ji[i].Title == "신" {
			if b.Ji[i].Title == "인" || b.Ji[i].Title == "신" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Choong.Exist = true
					b.Result[i].Choong.Exist = true
					a.Result[i].Choong.InSin += 1
					b.Result[i].Choong.InSin += 1
				}
			}
		}

		//묘유
		if a.Ji[i].Title == "묘" || a.Ji[i].Title == "유" {
			if b.Ji[i].Title == "묘" || b.Ji[i].Title == "유" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Choong.Exist = true
					b.Result[i].Choong.Exist = true
					a.Result[i].Choong.MyoYu += 1
					b.Result[i].Choong.MyoYu += 1
				}
			}
		}

		//진술
		if a.Ji[i].Title == "진" || a.Ji[i].Title == "술" {
			if b.Ji[i].Title == "진" || b.Ji[i].Title == "술" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Choong.Exist = true
					b.Result[i].Choong.Exist = true
					a.Result[i].Choong.JinSul += 1
					b.Result[i].Choong.JinSul += 1
				}
			}
		}

		//사해
		if a.Ji[i].Title == "사" || a.Ji[i].Title == "해" {
			if b.Ji[i].Title == "사" || b.Ji[i].Title == "해" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Choong.Exist = true
					b.Result[i].Choong.Exist = true
					a.Result[i].Choong.SaHae += 1
					b.Result[i].Choong.SaHae += 1
				}
			}
		}

		//자오
		if a.Ji[i].Title == "자" || a.Ji[i].Title == "오" {
			if b.Ji[i].Title == "자" || b.Ji[i].Title == "오" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Choong.Exist = true
					b.Result[i].Choong.Exist = true
					a.Result[i].Choong.JaO += 1
					b.Result[i].Choong.JaO += 1
				}
			}
		}

		//축미
		if a.Ji[i].Title == "축" || a.Ji[i].Title == "미" {
			if b.Ji[i].Title == "축" || b.Ji[i].Title == "미" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Choong.Exist = true
					b.Result[i].Choong.Exist = true
					a.Result[i].Choong.ChukMi += 1
					b.Result[i].Choong.ChukMi += 1
				}
			}
		}

	}

	return a, b
}

func Find_Pasal(a Person) {
	//자유 귀문이 우선
	num1 := 5
	num2 := 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "자" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "유" {
			num2 = j
		}
	}
	a.Result[num1].Pa.Exist = false
	a.Result[num2].Pa.Exist = false

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Pa.Exist = true
			a.Result[num2].Pa.Exist = true
			a.Result[num1].Pa.JaYu += 1
			a.Result[num2].Pa.JaYu += 1
		case num1 == 5:
			if a.Ji[4].Title == "자" {
				a.Result[4].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[4].Pa.JaYu += 1
				a.Result[num2].Pa.JaYu += 1
			} else if a.Ji[5].Title == "자" {
				a.Result[5].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[5].Pa.JaYu += 1
				a.Result[num2].Pa.JaYu += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "유" {
				a.Result[num1].Pa.Exist = true
				a.Result[4].Pa.Exist = true
				a.Result[num1].Pa.JaYu += 1
				a.Result[4].Pa.JaYu += 1
			} else if a.Ji[5].Title == "유" {
				a.Result[num1].Pa.Exist = true
				a.Result[5].Pa.Exist = true
				a.Result[num1].Pa.JaYu += 1
				a.Result[5].Pa.JaYu += 1
			}
		}
	}

	//묘오
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "묘" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "오" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Pa.Exist = true
			a.Result[num2].Pa.Exist = true
			a.Result[num1].Pa.MyoO += 1
			a.Result[num2].Pa.MyoO += 1
		case num1 == 5:
			if a.Ji[4].Title == "묘" {
				a.Result[4].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[4].Pa.MyoO += 1
				a.Result[num2].Pa.MyoO += 1
			} else if a.Ji[5].Title == "묘" {
				a.Result[5].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[5].Pa.MyoO += 1
				a.Result[num2].Pa.MyoO += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "오" {
				a.Result[num1].Pa.Exist = true
				a.Result[4].Pa.Exist = true
				a.Result[num1].Pa.MyoO += 1
				a.Result[4].Pa.MyoO += 1
			} else if a.Ji[5].Title == "오" {
				a.Result[num1].Pa.Exist = true
				a.Result[5].Pa.Exist = true
				a.Result[num1].Pa.MyoO += 1
				a.Result[5].Pa.MyoO += 1
			}

		}
	}

	//사신  형합이 우선
	//사신
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "사" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "신" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Pa.Exist = true
			a.Result[num2].Pa.Exist = true
			a.Result[num1].Pa.SaSin += 1
			a.Result[num2].Pa.SaSin += 1
		case num1 == 5:
			if a.Ji[4].Title == "사" {
				a.Result[4].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[4].Pa.SaSin += 1
				a.Result[num2].Pa.SaSin += 1
			} else if a.Ji[5].Title == "사" {
				a.Result[5].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[5].Pa.SaSin += 1
				a.Result[num2].Pa.SaSin += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "신" {
				a.Result[num1].Pa.Exist = true
				a.Result[4].Pa.Exist = true
				a.Result[num1].Pa.SaSin += 1
				a.Result[4].Pa.SaSin += 1
			} else if a.Ji[5].Title == "신" {
				a.Result[num1].Pa.Exist = true
				a.Result[5].Pa.Exist = true
				a.Result[num1].Pa.SaSin += 1
				a.Result[5].Pa.SaSin += 1
			}

		}
	}

	//축진
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "축" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "진" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Pa.Exist = true
			a.Result[num2].Pa.Exist = true
			a.Result[num1].Pa.JinChuk += 1
			a.Result[num2].Pa.JinChuk += 1
		case num1 == 5:
			if a.Ji[4].Title == "축" {
				a.Result[4].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[4].Pa.JinChuk += 1
				a.Result[num2].Pa.JinChuk += 1
			} else if a.Ji[5].Title == "축" {
				a.Result[5].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[5].Pa.JinChuk += 1
				a.Result[num2].Pa.JinChuk += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "진" {
				a.Result[num1].Pa.Exist = true
				a.Result[4].Pa.Exist = true
				a.Result[num1].Pa.JinChuk += 1
				a.Result[4].Pa.JinChuk += 1
			} else if a.Ji[5].Title == "진" {
				a.Result[num1].Pa.Exist = true
				a.Result[5].Pa.Exist = true
				a.Result[num1].Pa.JinChuk += 1
				a.Result[5].Pa.JinChuk += 1
			}
		}
	}

	//술미 형이 우선
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "술" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "미" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Pa.Exist = true
			a.Result[num2].Pa.Exist = true
			a.Result[num1].Pa.SulMi += 1
			a.Result[num2].Pa.SulMi += 1
		case num1 == 5:
			if a.Ji[4].Title == "술" {
				a.Result[4].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[4].Pa.SulMi += 1
				a.Result[num2].Pa.SulMi += 1
			} else if a.Ji[5].Title == "술" {
				a.Result[5].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[5].Pa.SulMi += 1
				a.Result[num2].Pa.SulMi += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "미" {
				a.Result[num1].Pa.Exist = true
				a.Result[4].Pa.Exist = true
				a.Result[num1].Pa.SulMi += 1
				a.Result[4].Pa.SulMi += 1
			} else if a.Ji[5].Title == "미" {
				a.Result[num1].Pa.Exist = true
				a.Result[5].Pa.Exist = true
				a.Result[num1].Pa.SulMi += 1
				a.Result[5].Pa.SulMi += 1
			}
		}
	}

	//인해 인중 병 손상
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "인" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "해" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Pa.Exist = true
			a.Result[num2].Pa.Exist = true
			a.Result[num1].Pa.InHae += 1
			a.Result[num2].Pa.InHae += 1
		case num1 == 5:
			if a.Ji[4].Title == "인" {
				a.Result[4].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[4].Pa.InHae += 1
				a.Result[num2].Pa.InHae += 1
			} else if a.Ji[5].Title == "인" {
				a.Result[5].Pa.Exist = true
				a.Result[num2].Pa.Exist = true
				a.Result[5].Pa.InHae += 1
				a.Result[num2].Pa.InHae += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "해" {
				a.Result[num1].Pa.Exist = true
				a.Result[4].Pa.Exist = true
				a.Result[num1].Pa.InHae += 1
				a.Result[4].Pa.InHae += 1
			} else if a.Ji[5].Title == "해" {
				a.Result[num1].Pa.Exist = true
				a.Result[5].Pa.Exist = true
				a.Result[num1].Pa.InHae += 1
				a.Result[5].Pa.InHae += 1
			}
		}
	}
}

func Find_Pasal_Goonghab(a Person, b Person) (Person, Person) {

	for i := 1; i < 3; i++ {
		a.Result[i].Pa.Exist = false
		b.Result[i].Pa.Exist = false
		//자유 귀문이 우선
		if a.Ji[i].Title == "자" || a.Ji[i].Title == "유" {
			if b.Ji[i].Title == "자" || b.Ji[i].Title == "유" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Pa.Exist = true
					b.Result[i].Pa.Exist = true
					a.Result[i].Pa.JaYu += 1
					b.Result[i].Pa.JaYu += 1
				}
			}
		}

		//묘오
		if a.Ji[i].Title == "묘" || a.Ji[i].Title == "오" {
			if b.Ji[i].Title == "묘" || b.Ji[i].Title == "오" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Pa.Exist = true
					b.Result[i].Pa.Exist = true
					a.Result[i].Pa.MyoO += 1
					b.Result[i].Pa.MyoO += 1
				}
			}
		}

		//사신  형합이 우선
		//사신
		if a.Ji[i].Title == "사" || a.Ji[i].Title == "신" {
			if b.Ji[i].Title == "사" || b.Ji[i].Title == "신" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Pa.Exist = true
					b.Result[i].Pa.Exist = true
					a.Result[i].Pa.SaSin += 1
					b.Result[i].Pa.SaSin += 1
				}
			}
		}

		//축진
		if a.Ji[i].Title == "사" || a.Ji[i].Title == "신" {
			if b.Ji[i].Title == "사" || b.Ji[i].Title == "신" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Pa.Exist = true
					b.Result[i].Pa.Exist = true
					a.Result[i].Pa.JinChuk += 1
					b.Result[i].Pa.JinChuk += 1
				}
			}
		}

		//술미 형이 우선
		if a.Ji[i].Title == "술" || a.Ji[i].Title == "미" {
			if b.Ji[i].Title == "술" || b.Ji[i].Title == "미" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Pa.Exist = true
					b.Result[i].Pa.Exist = true
					a.Result[i].Pa.SulMi += 1
					b.Result[i].Pa.SulMi += 1
				}
			}
		}

		//인해 인중 병 손상
		if a.Ji[i].Title == "인" || a.Ji[i].Title == "해" {
			if b.Ji[i].Title == "인" || b.Ji[i].Title == "해" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Pa.Exist = true
					b.Result[i].Pa.Exist = true
					a.Result[i].Pa.InHae += 1
					b.Result[i].Pa.InHae += 1
				}
			}
		}
	}

	return a, b
}

func Find_Haesal(a Person) {
	//자미
	num1 := 5
	num2 := 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "자" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "미" {
			num2 = j
		}
	}

	a.Result[num1].Hae.Exist = false
	a.Result[num2].Hae.Exist = false
	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hae.Exist = true
			a.Result[num2].Hae.Exist = true
			a.Result[num1].Hae.JaMi += 1
			a.Result[num2].Hae.JaMi += 1
		case num1 == 5:
			if a.Ji[4].Title == "자" {
				a.Result[4].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[4].Hae.JaMi += 1
				a.Result[num2].Hae.JaMi += 1
			} else if a.Ji[5].Title == "자" {
				a.Result[5].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[5].Hae.JaMi += 1
				a.Result[num2].Hae.JaMi += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "미" {
				a.Result[num1].Hae.Exist = true
				a.Result[4].Hae.Exist = true
				a.Result[num1].Hae.JaMi += 1
				a.Result[4].Hae.JaMi += 1
			} else if a.Ji[5].Title == "미" {
				a.Result[num1].Hae.Exist = true
				a.Result[5].Hae.Exist = true
				a.Result[num1].Hae.JaMi += 1
				a.Result[5].Hae.JaMi += 1
			}
		}
	}

	//축오 탕화중독
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "축" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "오" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hae.Exist = true
			a.Result[num2].Hae.Exist = true
			a.Result[num1].Hae.OChuk += 1
			a.Result[num2].Hae.OChuk += 1
		case num1 == 5:
			if a.Ji[4].Title == "축" {
				a.Result[4].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[4].Hae.OChuk += 1
				a.Result[num2].Hae.OChuk += 1
			} else if a.Ji[5].Title == "축" {
				a.Result[5].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[5].Hae.OChuk += 1
				a.Result[num2].Hae.OChuk += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "오" {
				a.Result[num1].Hae.Exist = true
				a.Result[4].Hae.Exist = true
				a.Result[num1].Hae.OChuk += 1
				a.Result[4].Hae.OChuk += 1
			} else if a.Ji[5].Title == "오" {
				a.Result[num1].Hae.Exist = true
				a.Result[5].Hae.Exist = true
				a.Result[num1].Hae.OChuk += 1
				a.Result[5].Hae.OChuk += 1
			}
		}
	}

	//인사 형이우선
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "인" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "사" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hae.Exist = true
			a.Result[num2].Hae.Exist = true
			a.Result[num1].Hae.InSa += 1
			a.Result[num2].Hae.InSa += 1
		case num1 == 5:
			if a.Ji[4].Title == "인" {
				a.Result[4].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[4].Hae.InSa += 1
				a.Result[num2].Hae.InSa += 1
			} else if a.Ji[5].Title == "인" {
				a.Result[5].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[5].Hae.InSa += 1
				a.Result[num2].Hae.InSa += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "사" {
				a.Result[num1].Hae.Exist = true
				a.Result[4].Hae.Exist = true
				a.Result[num1].Hae.InSa += 1
				a.Result[4].Hae.InSa += 1
			} else if a.Ji[5].Title == "사" {
				a.Result[num1].Hae.Exist = true
				a.Result[5].Hae.Exist = true
				a.Result[num1].Hae.InSa += 1
				a.Result[5].Hae.InSa += 1
			}
		}
	}

	//묘진
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "묘" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "진" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hae.Exist = true
			a.Result[num2].Hae.Exist = true
			a.Result[num1].Hae.MyoJin += 1
			a.Result[num2].Hae.MyoJin += 1
		case num1 == 5:
			if a.Ji[4].Title == "묘" {
				a.Result[4].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[4].Hae.MyoJin += 1
				a.Result[num2].Hae.MyoJin += 1
			} else if a.Ji[5].Title == "묘" {
				a.Result[5].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[5].Hae.MyoJin += 1
				a.Result[num2].Hae.MyoJin += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "진" {
				a.Result[num1].Hae.Exist = true
				a.Result[4].Hae.Exist = true
				a.Result[num1].Hae.MyoJin += 1
				a.Result[4].Hae.MyoJin += 1
			} else if a.Ji[5].Title == "진" {
				a.Result[num1].Hae.Exist = true
				a.Result[5].Hae.Exist = true
				a.Result[num1].Hae.MyoJin += 1
				a.Result[5].Hae.MyoJin += 1
			}
		}
	}

	//신해 신이 손상
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "신" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "해" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hae.Exist = true
			a.Result[num2].Hae.Exist = true
			a.Result[num1].Hae.HaeSin += 1
			a.Result[num2].Hae.HaeSin += 1
		case num1 == 5:
			if a.Ji[4].Title == "신" {
				a.Result[4].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[4].Hae.HaeSin += 1
				a.Result[num2].Hae.HaeSin += 1
			} else if a.Ji[5].Title == "신" {
				a.Result[5].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[5].Hae.HaeSin += 1
				a.Result[num2].Hae.HaeSin += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "해" {
				a.Result[num1].Hae.Exist = true
				a.Result[4].Hae.Exist = true
				a.Result[num1].Hae.HaeSin += 1
				a.Result[4].Hae.HaeSin += 1
			} else if a.Ji[5].Title == "해" {
				a.Result[num1].Hae.Exist = true
				a.Result[5].Hae.Exist = true
				a.Result[num1].Hae.HaeSin += 1
				a.Result[5].Hae.HaeSin += 1
			}
		}
	}

	//유술
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "유" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "술" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].Hae.Exist = true
			a.Result[num2].Hae.Exist = true
			a.Result[num1].Hae.YuSul += 1
			a.Result[num2].Hae.YuSul += 1
		case num1 == 5:
			if a.Ji[4].Title == "유" {
				a.Result[4].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[4].Hae.YuSul += 1
				a.Result[num2].Hae.YuSul += 1
			} else if a.Ji[5].Title == "유" {
				a.Result[5].Hae.Exist = true
				a.Result[num2].Hae.Exist = true
				a.Result[5].Hae.YuSul += 1
				a.Result[num2].Hae.YuSul += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "술" {
				a.Result[num1].Hae.Exist = true
				a.Result[4].Hae.Exist = true
				a.Result[num1].Hae.YuSul += 1
				a.Result[4].Hae.YuSul += 1
			} else if a.Ji[5].Title == "술" {
				a.Result[num1].Hae.Exist = true
				a.Result[5].Hae.Exist = true
				a.Result[num1].Hae.YuSul += 1
				a.Result[5].Hae.YuSul += 1
			}
		}
	}
}

func Find_Haesal_Goonghab(a Person, b Person) (Person, Person) {
	//자미
	for i := 1; i < 3; i++ {
		a.Result[i].Hae.Exist = false
		b.Result[i].Hae.Exist = false
		if a.Ji[i].Title == "자" || a.Ji[i].Title == "미" {
			if b.Ji[i].Title == "자" || b.Ji[i].Title == "미" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hae.Exist = true
					b.Result[i].Hae.Exist = true
					a.Result[i].Hae.JaMi += 1
					b.Result[i].Hae.JaMi += 1
				}
			}
		}

		//축오 탕화중독
		if a.Ji[i].Title == "축" || a.Ji[i].Title == "오" {
			if b.Ji[i].Title == "축" || b.Ji[i].Title == "오" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hae.Exist = true
					b.Result[i].Hae.Exist = true
					a.Result[i].Hae.OChuk += 1
					b.Result[i].Hae.OChuk += 1
				}
			}
		}

		//인사 형이우선
		if a.Ji[i].Title == "인" || a.Ji[i].Title == "사" {
			if b.Ji[i].Title == "인" || b.Ji[i].Title == "사" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hae.Exist = true
					b.Result[i].Hae.Exist = true
					a.Result[i].Hae.InSa += 1
					b.Result[i].Hae.InSa += 1
				}
			}
		}

		//묘진
		if a.Ji[i].Title == "묘" || a.Ji[i].Title == "진" {
			if b.Ji[i].Title == "묘" || b.Ji[i].Title == "진" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hae.Exist = true
					b.Result[i].Hae.Exist = true
					a.Result[i].Hae.MyoJin += 1
					b.Result[i].Hae.MyoJin += 1
				}
			}
		}

		//신해 신이 손상
		if a.Ji[i].Title == "신" || a.Ji[i].Title == "해" {
			if b.Ji[i].Title == "신" || b.Ji[i].Title == "해" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hae.Exist = true
					b.Result[i].Hae.Exist = true
					a.Result[i].Hae.HaeSin += 1
					b.Result[i].Hae.HaeSin += 1
				}
			}
		}

		//유술
		if a.Ji[i].Title == "유" || a.Ji[i].Title == "술" {
			if b.Ji[i].Title == "유" || b.Ji[i].Title == "술" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].Hae.Exist = true
					b.Result[i].Hae.Exist = true
					a.Result[i].Hae.YuSul += 1
					b.Result[i].Hae.YuSul += 1
				}
			}
		}
	}

	return a, b
}

func Find_Wonzin(a Person) {
	//인유
	num1 := 5
	num2 := 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "인" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "유" {
			num2 = j
		}
	}
	a.Result[num1].WonJin.Exist = false
	a.Result[num2].WonJin.Exist = false

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].WonJin.Exist = true
			a.Result[num2].WonJin.Exist = true
			a.Result[num1].WonJin.InYu += 1
			a.Result[num2].WonJin.InYu += 1
		case num1 == 5:
			if a.Ji[4].Title == "인" {
				a.Result[4].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[4].WonJin.InYu += 1
				a.Result[num2].WonJin.InYu += 1
			} else if a.Ji[5].Title == "인" {
				a.Result[5].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[5].WonJin.InYu += 1
				a.Result[num2].WonJin.InYu += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "유" {
				a.Result[num1].WonJin.Exist = true
				a.Result[4].WonJin.Exist = true
				a.Result[num1].WonJin.InYu += 1
				a.Result[4].WonJin.InYu += 1
			} else if a.Ji[5].Title == "유" {
				a.Result[num1].WonJin.Exist = true
				a.Result[5].WonJin.Exist = true
				a.Result[num1].WonJin.InYu += 1
				a.Result[5].WonJin.InYu += 1
			}
		}
	}

	//묘신 금목상쟁
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "묘" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "신" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].WonJin.Exist = true
			a.Result[num2].WonJin.Exist = true
			a.Result[num1].WonJin.MyoSin += 1
			a.Result[num2].WonJin.MyoSin += 1
		case num1 == 5:
			if a.Ji[4].Title == "묘" {
				a.Result[4].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[4].WonJin.MyoSin += 1
				a.Result[num2].WonJin.MyoSin += 1
			} else if a.Ji[5].Title == "묘" {
				a.Result[5].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[5].WonJin.MyoSin += 1
				a.Result[num2].WonJin.MyoSin += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "신" {
				a.Result[num1].WonJin.Exist = true
				a.Result[4].WonJin.Exist = true
				a.Result[num1].WonJin.MyoSin += 1
				a.Result[4].WonJin.MyoSin += 1
			} else if a.Ji[5].Title == "신" {
				a.Result[num1].WonJin.Exist = true
				a.Result[5].WonJin.Exist = true
				a.Result[num1].WonJin.MyoSin += 1
				a.Result[5].WonJin.MyoSin += 1
			}
		}
	}

	//진해 입묘형
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "진" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "해" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].WonJin.Exist = true
			a.Result[num2].WonJin.Exist = true
			a.Result[num1].WonJin.JinHae += 1
			a.Result[num2].WonJin.JinHae += 1
		case num1 == 5:
			if a.Ji[4].Title == "진" {
				a.Result[4].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[4].WonJin.JinHae += 1
				a.Result[num2].WonJin.JinHae += 1
			} else if a.Ji[5].Title == "진" {
				a.Result[5].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[5].WonJin.JinHae += 1
				a.Result[num2].WonJin.JinHae += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "해" {
				a.Result[num1].WonJin.Exist = true
				a.Result[4].WonJin.Exist = true
				a.Result[num1].WonJin.JinHae += 1
				a.Result[4].WonJin.JinHae += 1
			} else if a.Ji[5].Title == "해" {
				a.Result[num1].WonJin.Exist = true
				a.Result[5].WonJin.Exist = true
				a.Result[num1].WonJin.JinHae += 1
				a.Result[5].WonJin.JinHae += 1
			}
		}
	}

	//사술 입묘형
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "사" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "술" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].WonJin.Exist = true
			a.Result[num2].WonJin.Exist = true
			a.Result[num1].WonJin.SaSul += 1
			a.Result[num2].WonJin.SaSul += 1
		case num1 == 5:
			if a.Ji[4].Title == "사" {
				a.Result[4].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[4].WonJin.SaSul += 1
				a.Result[num2].WonJin.SaSul += 1
			} else if a.Ji[5].Title == "사" {
				a.Result[5].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[5].WonJin.SaSul += 1
				a.Result[num2].WonJin.SaSul += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "술" {
				a.Result[num1].WonJin.Exist = true
				a.Result[4].WonJin.Exist = true
				a.Result[num1].WonJin.SaSul += 1
				a.Result[4].WonJin.SaSul += 1
			} else if a.Ji[5].Title == "술" {
				a.Result[num1].WonJin.Exist = true
				a.Result[5].WonJin.Exist = true
				a.Result[num1].WonJin.SaSul += 1
				a.Result[5].WonJin.SaSul += 1
			}
		}
	}

	//오축 탕화형
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "오" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "축" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].WonJin.Exist = true
			a.Result[num2].WonJin.Exist = true
			a.Result[num1].WonJin.OChuk += 1
			a.Result[num2].WonJin.OChuk += 1
		case num1 == 5:
			if a.Ji[4].Title == "오" {
				a.Result[4].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[4].WonJin.OChuk += 1
				a.Result[num2].WonJin.OChuk += 1
			} else if a.Ji[5].Title == "오" {
				a.Result[5].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[5].WonJin.OChuk += 1
				a.Result[num2].WonJin.OChuk += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "축" {
				a.Result[num1].WonJin.Exist = true
				a.Result[4].WonJin.Exist = true
				a.Result[num1].WonJin.OChuk += 1
				a.Result[4].WonJin.OChuk += 1
			} else if a.Ji[5].Title == "축" {
				a.Result[num1].WonJin.Exist = true
				a.Result[5].WonJin.Exist = true
				a.Result[num1].WonJin.OChuk += 1
				a.Result[5].WonJin.OChuk += 1
			}
		}
	}

	//자미탕화원진귀문
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "자" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "미" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].WonJin.Exist = true
			a.Result[num2].WonJin.Exist = true
			a.Result[num1].WonJin.JaMi += 1
			a.Result[num2].WonJin.JaMi += 1
		case num1 == 5:
			if a.Ji[4].Title == "자" {
				a.Result[4].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[4].WonJin.JaMi += 1
				a.Result[num2].WonJin.JaMi += 1
			} else if a.Ji[5].Title == "자" {
				a.Result[5].WonJin.Exist = true
				a.Result[num2].WonJin.Exist = true
				a.Result[5].WonJin.JaMi += 1
				a.Result[num2].WonJin.JaMi += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "미" {
				a.Result[num1].WonJin.Exist = true
				a.Result[4].WonJin.Exist = true
				a.Result[num1].WonJin.JaMi += 1
				a.Result[4].WonJin.JaMi += 1
			} else if a.Ji[5].Title == "미" {
				a.Result[num1].WonJin.Exist = true
				a.Result[5].WonJin.Exist = true
				a.Result[num1].WonJin.JaMi += 1
				a.Result[5].WonJin.JaMi += 1
			}
		}
	}
}

func Find_Wonzin_Goonghab(a Person, b Person) (Person, Person) {

	for i := 1; i < 3; i++ {
		a.Result[i].WonJin.Exist = false
		b.Result[i].WonJin.Exist = false
		//인유
		if a.Ji[i].Title == "인" || a.Ji[i].Title == "유" {
			if b.Ji[i].Title == "인" || b.Ji[i].Title == "유" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].WonJin.Exist = true
					b.Result[i].WonJin.Exist = true
					a.Result[i].WonJin.InYu += 1
					b.Result[i].WonJin.InYu += 1
				}
			}
		}

		//묘신 금목상쟁
		if a.Ji[i].Title == "묘" || a.Ji[i].Title == "신" {
			if b.Ji[i].Title == "묘" || b.Ji[i].Title == "신" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].WonJin.Exist = true
					b.Result[i].WonJin.Exist = true
					a.Result[i].WonJin.MyoSin += 1
					b.Result[i].WonJin.MyoSin += 1
				}
			}
		}

		//진해 입묘형
		if a.Ji[i].Title == "진" || a.Ji[i].Title == "해" {
			if b.Ji[i].Title == "진" || b.Ji[i].Title == "해" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].WonJin.Exist = true
					b.Result[i].WonJin.Exist = true
					a.Result[i].WonJin.JinHae += 1
					b.Result[i].WonJin.JinHae += 1
				}
			}
		}

		//사술 입묘형
		if a.Ji[i].Title == "사" || a.Ji[i].Title == "술" {
			if b.Ji[i].Title == "사" || b.Ji[i].Title == "술" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].WonJin.Exist = true
					b.Result[i].WonJin.Exist = true
					a.Result[i].WonJin.SaSul += 1
					b.Result[i].WonJin.SaSul += 1
				}
			}
		}

		//오축 탕화형
		if a.Ji[i].Title == "오" || a.Ji[i].Title == "축" {
			if b.Ji[i].Title == "오" || b.Ji[i].Title == "축" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].WonJin.Exist = true
					b.Result[i].WonJin.Exist = true
					a.Result[i].WonJin.OChuk += 1
					b.Result[i].WonJin.OChuk += 1
				}
			}
		}

		//자미탕화원진귀문
		if a.Ji[2].Title == "자" || a.Ji[2].Title == "미" {
			if b.Ji[2].Title == "자" || b.Ji[2].Title == "미" {
				if a.Ji[2].Title != b.Ji[2].Title {
					a.Result[i].WonJin.Exist = true
					b.Result[i].WonJin.Exist = true
					a.Result[2].WonJin.JaMi += 1
					b.Result[2].WonJin.JaMi += 1
				}
			}
		}
	}

	return a, b
}

func Find_Guimun(a Person) {
	//인미
	num1 := 5
	num2 := 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "인" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "미" {
			num2 = j
		}
	}
	a.Result[num1].GuiMoon.Exist = false
	a.Result[num2].GuiMoon.Exist = false

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].GuiMoon.Exist = true
			a.Result[num2].GuiMoon.Exist = true
			a.Result[num1].GuiMoon.InMi += 1
			a.Result[num2].GuiMoon.InMi += 1
		case num1 == 5:
			if a.Ji[4].Title == "인" {
				a.Result[4].GuiMoon.Exist = true
				a.Result[num2].GuiMoon.Exist = true
				a.Result[4].GuiMoon.InMi += 1
				a.Result[num2].GuiMoon.InMi += 1
			} else if a.Ji[5].Title == "인" {
				a.Result[5].GuiMoon.Exist = true
				a.Result[num2].GuiMoon.Exist = true
				a.Result[5].GuiMoon.InMi += 1
				a.Result[num2].GuiMoon.InMi += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "미" {
				a.Result[num1].GuiMoon.Exist = true
				a.Result[4].GuiMoon.Exist = true
				a.Result[num1].GuiMoon.InMi += 1
				a.Result[4].GuiMoon.InMi += 1
			} else if a.Ji[5].Title == "미" {
				a.Result[num1].GuiMoon.Exist = true
				a.Result[5].GuiMoon.Exist = true
				a.Result[num1].GuiMoon.InMi += 1
				a.Result[5].GuiMoon.InMi += 1
			}
		}
	}

	//자유 음주형
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "자" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "유" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].GuiMoon.Exist = true
			a.Result[num2].GuiMoon.Exist = true
			a.Result[num1].GuiMoon.JaYu += 1
			a.Result[num2].GuiMoon.JaYu += 1
		case num1 == 5:
			if a.Ji[4].Title == "자" {
				a.Result[4].GuiMoon.Exist = true
				a.Result[num2].GuiMoon.Exist = true
				a.Result[4].GuiMoon.JaYu += 1
				a.Result[num2].GuiMoon.JaYu += 1
			} else if a.Ji[5].Title == "자" {
				a.Result[5].GuiMoon.Exist = true
				a.Result[num2].GuiMoon.Exist = true
				a.Result[5].GuiMoon.JaYu += 1
				a.Result[num2].GuiMoon.JaYu += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "유" {
				a.Result[num1].GuiMoon.Exist = true
				a.Result[4].GuiMoon.Exist = true
				a.Result[num1].GuiMoon.JaYu += 1
				a.Result[4].GuiMoon.JaYu += 1
			} else if a.Ji[5].Title == "유" {
				a.Result[num1].GuiMoon.Exist = true
				a.Result[5].GuiMoon.Exist = true
				a.Result[num1].GuiMoon.JaYu += 1
				a.Result[5].GuiMoon.JaYu += 1
			}
		}
	}
}

func Find_Guimun_Goonghab(a Person, b Person) (Person, Person) {
	//인미
	for i := 1; i < 3; i++ {
		a.Result[i].GuiMoon.Exist = false
		b.Result[i].GuiMoon.Exist = false
		if a.Ji[i].Title == "인" || a.Ji[i].Title == "미" {
			if b.Ji[i].Title == "인" || b.Ji[i].Title == "미" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].GuiMoon.Exist = true
					b.Result[i].GuiMoon.Exist = true
					a.Result[i].GuiMoon.InMi += 1
					b.Result[i].GuiMoon.InMi += 1
				}
			}
		}

		//자유 음주형
		if a.Ji[i].Title == "자" || a.Ji[i].Title == "유" {
			if b.Ji[i].Title == "자" || b.Ji[i].Title == "유" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].GuiMoon.Exist = true
					b.Result[i].GuiMoon.Exist = true
					a.Result[i].GuiMoon.JaYu += 1
					b.Result[i].GuiMoon.JaYu += 1
				}
			}
		}
	}

	return a, b
}

func Find_Gyeokgak(a Person) {
	//사묘축
	num1 := 5
	num2 := 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "사" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "묘" {
			num2 = j
		}
	}
	a.Result[num1].GyeokGak.Exist = false
	a.Result[num2].GyeokGak.Exist = false

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].GyeokGak.Exist = true
			a.Result[num2].GyeokGak.Exist = true
			a.Result[num1].GyeokGak.MyoSa += 1
			a.Result[num2].GyeokGak.MyoSa += 1
		case num1 == 5:
			if a.Ji[4].Title == "사" {
				a.Result[4].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[4].GyeokGak.MyoSa += 1
				a.Result[num2].GyeokGak.MyoSa += 1
			} else if a.Ji[5].Title == "사" {
				a.Result[5].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[5].GyeokGak.MyoSa += 1
				a.Result[num2].GyeokGak.MyoSa += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "묘" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[4].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.MyoSa += 1
				a.Result[4].GyeokGak.MyoSa += 1
			} else if a.Ji[5].Title == "묘" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[5].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.MyoSa += 1
				a.Result[5].GyeokGak.MyoSa += 1
			}
		}
	}

	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "묘" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "축" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].GyeokGak.Exist = true
			a.Result[num2].GyeokGak.Exist = true
			a.Result[num1].GyeokGak.MyoChuk += 1
			a.Result[num2].GyeokGak.MyoChuk += 1
		case num1 == 5:
			if a.Ji[4].Title == "묘" {
				a.Result[4].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[4].GyeokGak.MyoChuk += 1
				a.Result[num2].GyeokGak.MyoChuk += 1
			} else if a.Ji[5].Title == "묘" {
				a.Result[5].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[5].GyeokGak.MyoChuk += 1
				a.Result[num2].GyeokGak.MyoChuk += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "축" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[4].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.MyoChuk += 1
				a.Result[4].GyeokGak.MyoChuk += 1
			} else if a.Ji[5].Title == "축" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[5].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.MyoChuk += 1
				a.Result[5].GyeokGak.MyoChuk += 1
			}

		}
	}

	//신오진
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "신" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "오" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].GyeokGak.Exist = true
			a.Result[num2].GyeokGak.Exist = true
			a.Result[num1].GyeokGak.OSin += 1
			a.Result[num2].GyeokGak.OSin += 1
		case num1 == 5:
			if a.Ji[4].Title == "신" {
				a.Result[4].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[4].GyeokGak.OSin += 1
				a.Result[num2].GyeokGak.OSin += 1
			} else if a.Ji[5].Title == "신" {
				a.Result[5].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[5].GyeokGak.OSin += 1
				a.Result[num2].GyeokGak.OSin += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "오" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[4].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.OSin += 1
				a.Result[4].GyeokGak.OSin += 1
			} else if a.Ji[5].Title == "오" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[5].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.OSin += 1
				a.Result[5].GyeokGak.OSin += 1
			}
		}
	}

	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "오" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "진" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].GyeokGak.Exist = true
			a.Result[num2].GyeokGak.Exist = true
			a.Result[num1].GyeokGak.JinO += 1
			a.Result[num2].GyeokGak.JinO += 1
		case num1 == 5:
			if a.Ji[4].Title == "오" {
				a.Result[4].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[4].GyeokGak.JinO += 1
				a.Result[num2].GyeokGak.JinO += 1
			} else if a.Ji[5].Title == "오" {
				a.Result[5].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[5].GyeokGak.JinO += 1
				a.Result[num2].GyeokGak.JinO += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "진" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[4].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.JinO += 1
				a.Result[4].GyeokGak.JinO += 1
			} else if a.Ji[5].Title == "진" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[5].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.JinO += 1
				a.Result[5].GyeokGak.JinO += 1
			}
		}
	}

	//해유미
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "유" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "미" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].GyeokGak.Exist = true
			a.Result[num2].GyeokGak.Exist = true
			a.Result[num1].GyeokGak.MiYu += 1
			a.Result[num2].GyeokGak.MiYu += 1
		case num1 == 5:
			if a.Ji[4].Title == "유" {
				a.Result[4].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[4].GyeokGak.MiYu += 1
				a.Result[num2].GyeokGak.MiYu += 1
			} else if a.Ji[5].Title == "유" {
				a.Result[5].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[5].GyeokGak.MiYu += 1
				a.Result[num2].GyeokGak.MiYu += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "미" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[4].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.MiYu += 1
				a.Result[4].GyeokGak.MiYu += 1
			} else if a.Ji[5].Title == "미" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[5].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.MiYu += 1
				a.Result[5].GyeokGak.MiYu += 1
			}
		}
	}

	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "유" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "해" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].GyeokGak.Exist = true
			a.Result[num2].GyeokGak.Exist = true
			a.Result[num1].GyeokGak.YuHae += 1
			a.Result[num2].GyeokGak.YuHae += 1
		case num1 == 5:
			if a.Ji[4].Title == "유" {
				a.Result[4].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[4].GyeokGak.YuHae += 1
				a.Result[num2].GyeokGak.YuHae += 1
			} else if a.Ji[5].Title == "유" {
				a.Result[5].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[5].GyeokGak.YuHae += 1
				a.Result[num2].GyeokGak.YuHae += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "해" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[4].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.YuHae += 1
				a.Result[4].GyeokGak.YuHae += 1
			} else if a.Ji[5].Title == "해" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[5].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.YuHae += 1
				a.Result[5].GyeokGak.YuHae += 1
			}
		}
	}

	//인자술
	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "자" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "인" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].GyeokGak.Exist = true
			a.Result[num2].GyeokGak.Exist = true
			a.Result[num1].GyeokGak.InJa += 1
			a.Result[num2].GyeokGak.InJa += 1
		case num1 == 5:
			if a.Ji[4].Title == "자" {
				a.Result[4].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[4].GyeokGak.InJa += 1
				a.Result[num2].GyeokGak.InJa += 1
			} else if a.Ji[5].Title == "자" {
				a.Result[5].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[5].GyeokGak.InJa += 1
				a.Result[num2].GyeokGak.InJa += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "인" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[4].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.InJa += 1
				a.Result[4].GyeokGak.InJa += 1
			} else if a.Ji[5].Title == "인" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[5].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.InJa += 1
				a.Result[5].GyeokGak.InJa += 1
			}
		}
	}

	num1 = 5
	num2 = 5

	for i := 0; i < 4; i++ {
		if a.Ji[i].Title == "자" {
			num1 = i
		}
	}
	for j := 0; j < 4; j++ {
		if a.Ji[j].Title == "술" {
			num2 = j
		}
	}

	if num1 != 5 || num2 != 5 {
		switch {
		case num1 != 5 && num2 != 5:
			a.Result[num1].GyeokGak.Exist = true
			a.Result[num2].GyeokGak.Exist = true
			a.Result[num1].GyeokGak.SulJa += 1
			a.Result[num2].GyeokGak.SulJa += 1
		case num1 == 5:
			if a.Ji[4].Title == "자" {
				a.Result[4].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[4].GyeokGak.SulJa += 1
				a.Result[num2].GyeokGak.SulJa += 1
			} else if a.Ji[5].Title == "자" {
				a.Result[5].GyeokGak.Exist = true
				a.Result[num2].GyeokGak.Exist = true
				a.Result[5].GyeokGak.SulJa += 1
				a.Result[num2].GyeokGak.SulJa += 1
			}

			fallthrough
		case num2 == 5:
			if a.Ji[4].Title == "술" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[4].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.SulJa += 1
				a.Result[4].GyeokGak.SulJa += 1
			} else if a.Ji[5].Title == "술" {
				a.Result[num1].GyeokGak.Exist = true
				a.Result[5].GyeokGak.Exist = true
				a.Result[num1].GyeokGak.SulJa += 1
				a.Result[5].GyeokGak.SulJa += 1
			}
		}
	}
}

func Find_Gyeokgak_Goonghab(a Person, b Person) (Person, Person) {
	//사묘축
	for i := 1; i < 3; i++ {
		a.Result[i].GyeokGak.Exist = false
		b.Result[i].GyeokGak.Exist = false
		if a.Ji[i].Title == "묘" || a.Ji[i].Title == "사" {
			if b.Ji[i].Title == "묘" || b.Ji[i].Title == "사" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].GyeokGak.Exist = true
					a.Result[i].GyeokGak.MyoSa += 1
					b.Result[i].GyeokGak.Exist = true
					b.Result[i].GyeokGak.MyoSa += 1
				}
			}
		}

		if a.Ji[i].Title == "묘" || a.Ji[i].Title == "축" {
			if b.Ji[i].Title == "묘" || b.Ji[i].Title == "축" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].GyeokGak.Exist = true
					a.Result[i].GyeokGak.MyoChuk += 1
					b.Result[i].GyeokGak.Exist = true
					b.Result[i].GyeokGak.MyoChuk += 1
				}
			}
		}

		//신오진
		if a.Ji[i].Title == "오" || a.Ji[i].Title == "신" {
			if b.Ji[i].Title == "오" || b.Ji[i].Title == "신" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].GyeokGak.Exist = true
					a.Result[i].GyeokGak.MyoChuk += 1
					b.Result[i].GyeokGak.Exist = true
					b.Result[i].GyeokGak.OSin += 1
				}
			}
		}

		if a.Ji[i].Title == "오" || a.Ji[i].Title == "진" {
			if b.Ji[i].Title == "오" || b.Ji[i].Title == "진" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].GyeokGak.Exist = true
					a.Result[i].GyeokGak.MyoChuk += 1
					b.Result[i].GyeokGak.Exist = true
					b.Result[i].GyeokGak.JinO += 1
				}
			}
		}

		//해유미
		if a.Ji[i].Title == "유" || a.Ji[i].Title == "해" {
			if b.Ji[i].Title == "유" || b.Ji[i].Title == "해" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].GyeokGak.Exist = true
					a.Result[i].GyeokGak.MyoChuk += 1
					b.Result[i].GyeokGak.Exist = true
					b.Result[i].GyeokGak.YuHae += 1
				}
			}
		}

		if a.Ji[i].Title == "유" || a.Ji[i].Title == "미" {
			if b.Ji[i].Title == "유" || b.Ji[i].Title == "미" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].GyeokGak.Exist = true
					a.Result[i].GyeokGak.MyoChuk += 1
					b.Result[i].GyeokGak.Exist = true
					b.Result[i].GyeokGak.MiYu += 1
				}
			}
		}

		//인자술
		if a.Ji[i].Title == "자" || a.Ji[i].Title == "인" {
			if b.Ji[i].Title == "자" || b.Ji[i].Title == "인" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].GyeokGak.Exist = true
					a.Result[i].GyeokGak.MyoChuk += 1
					b.Result[i].GyeokGak.Exist = true
					b.Result[i].GyeokGak.InJa += 1
				}
			}
		}

		if a.Ji[i].Title == "자" || a.Ji[i].Title == "술" {
			if b.Ji[i].Title == "자" || b.Ji[i].Title == "술" {
				if a.Ji[i].Title != b.Ji[i].Title {
					a.Result[i].GyeokGak.Exist = true
					a.Result[i].GyeokGak.MyoChuk += 1
					b.Result[i].GyeokGak.Exist = true
					b.Result[i].GyeokGak.SulJa += 1
				}
			}
		}

	}

	return a, b
}

func Fing_AmHab(a Person) Person {

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			a.Result[i].AmHab.Exist = false
			a.Result[j].AmHab.Exist = false
			if i != j {
				switch a.Ji[i].Title {
				case "인":
					switch a.Ji[j].Title {
					case "유":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "술":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "축":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "묘":
					switch a.Ji[j].Title {
					case "사":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "신":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "진":
					switch a.Ji[j].Title {
					case "사":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "사":
					switch a.Ji[j].Title {
					case "묘":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "진":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "자":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "축":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "오":
					switch a.Ji[j].Title {
					case "신":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "해":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "미":
					switch a.Ji[j].Title {
					case "신":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "유":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "신":
					switch a.Ji[j].Title {
					case "묘":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "오":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "미":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "유":
					switch a.Ji[j].Title {
					case "인":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "술":
					switch a.Ji[j].Title {
					case "해":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "해":
					switch a.Ji[j].Title {
					case "오":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "미":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "술":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "자":
					switch a.Ji[j].Title {
					case "인":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "사":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				case "축":
					switch a.Ji[j].Title {
					case "인":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					case "사":
						a.Result[i].AmHab.Exist = true
						a.Result[j].AmHab.Exist = true
						a.Result[i].AmHab.AmHab = 1
						a.Result[j].AmHab.AmHab = 1
					default:
						a.Result[i].AmHab.Exist = false
						a.Result[j].AmHab.Exist = false
					}
				}
			}
		}
	}

	return a
}

func Find_AmHab_Goonghab(a Person, b Person) (Person, Person) {
	for i := 1; i < 3; i++ {
		switch a.Ji[i].Title {
		case "인":
			switch b.Ji[i].Title {
			case "유":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "술":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "축":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "묘":
			switch b.Ji[i].Title {
			case "사":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "신":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "진":
			switch b.Ji[i].Title {
			case "사":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "사":
			switch b.Ji[i].Title {
			case "묘":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "진":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "자":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "축":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "오":
			switch b.Ji[i].Title {
			case "신":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "해":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "미":
			switch b.Ji[i].Title {
			case "신":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "유":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "신":
			switch b.Ji[i].Title {
			case "묘":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "오":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "미":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "유":
			switch b.Ji[i].Title {
			case "인":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "술":
			switch b.Ji[i].Title {
			case "해":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "해":
			switch b.Ji[i].Title {
			case "오":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "미":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "술":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "자":
			switch b.Ji[i].Title {
			case "인":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "사":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		case "축":
			switch b.Ji[i].Title {
			case "인":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			case "사":
				a.Result[i].AmHab.Exist = true
				b.Result[i].AmHab.Exist = true
				a.Result[i].AmHab.AmHab = 1
				b.Result[i].AmHab.AmHab = 1
			default:
				a.Result[i].AmHab.Exist = false
				b.Result[i].AmHab.Exist = false
			}
		}
	}
	return a, b
}

func Find_Characteristics(host Person) Person {
	Find_Banghab(host)
	Find_Samhab(host)
	Find_Yukhab(host)
	Find_Hyungsal(host)
	Find_Choongsal(host)
	Find_Pasal(host)
	Find_Haesal(host)
	Find_Wonzin(host)
	Find_Guimun(host)
	Find_Gyeokgak(host)
	return host
}

func (sa *sajuAnalyzer) Find_GoongHab(mysaju *PersonSaju, friendsaju *PersonSaju, chungan_table []Chungan, jiji_table []Jiji, sibsung_table []Sibsung, unsung_table []Chungan_Unsung) (Person, Person) {
	host_chungan_received := []string{mysaju.YearChun, mysaju.MonthChun, mysaju.DayChun, mysaju.TimeChun, mysaju.DaeunChun, mysaju.SaeunChun, mysaju.YearJi, mysaju.MonthJi, mysaju.DayJi, mysaju.TimeJi, mysaju.DaeUnJi, mysaju.SaeunJi}
	opponent_chungan_received := []string{friendsaju.YearChun, friendsaju.MonthChun, friendsaju.DayChun, friendsaju.TimeChun, friendsaju.DaeunChun, friendsaju.SaeunChun, friendsaju.YearJi, friendsaju.MonthJi, friendsaju.DayJi, friendsaju.TimeJi, friendsaju.DaeUnJi, friendsaju.SaeunJi}

	host := person_chungan_input(host_chungan_received, chungan_table, jiji_table, sibsung_table, unsung_table)
	opponent := person_chungan_input(opponent_chungan_received, chungan_table, jiji_table, sibsung_table, unsung_table)

	go func() { host, opponent = Find_Sibsung_Goonghab(host, opponent, sa.Sibsung) }()
	go func() { host, opponent = Find_Unsung_Goonghab(host, opponent, sa.Sib2Unsung) }()
	go func() { host, opponent = Find_Chungan_hab(host, opponent) }()
	go func() { host, opponent = Find_Chungan_Geok(host, opponent) }()
	go func() { host, opponent = Find_Banghab_Goonghab(host, opponent) }()
	go func() { host, opponent = Find_Samhab_Goonghab(host, opponent) }()
	go func() { host, opponent = Find_Yukhab_Goonghab(host, opponent) }()
	go func() { host, opponent = Find_Hyungsal_Goonghab(host, opponent) }()
	go func() { host, opponent = Find_Choongsal_Goonghab(host, opponent) }()
	go func() { host, opponent = Find_Pasal_Goonghab(host, opponent) }()
	go func() { host, opponent = Find_Haesal_Goonghab(host, opponent) }()
	go func() { host, opponent = Find_Wonzin_Goonghab(host, opponent) }()
	go func() { host, opponent = Find_Guimun_Goonghab(host, opponent) }()
	go func() { host, opponent = Find_Gyeokgak_Goonghab(host, opponent) }()
	go func() { host, opponent = Find_AmHab_Goonghab(host, opponent) }()

	host.LoginID = mysaju.LoginID
	opponent.LoginID = friendsaju.LoginID

	return host, opponent
}

func (sa *sajuAnalyzer) Evaluate_GoonbHab(host Person, opponent Person) (string, string, string, string) {
	var host_grade [4]int
	var opponent_grade [4]int
	var host_description [4]string
	var opponent_description [4]string
	for i := 1; i < 3; i++ {
		if host.Result[i].ChunGanHab.Exist {
			if host.Result[i].ChunGanHab.GabGi == 1 {
				host_description[i], opponent_description[i] = "생각 또는 가치관이 잘 맞다. ", "생각 또는 가치관이 잘 맞다. "
				host_grade[i], opponent_grade[i] = 5, 5
			} else if host.Result[i].ChunGanHab.ElGyeong == 1 {
				host_description[i], opponent_description[i] = "생각 또는 가치관이 잘 맞다. ", "생각 또는 가치관이 잘 맞다. "
				host_grade[i], opponent_grade[i] = 5, 5
			} else if host.Result[i].ChunGanHab.ByeongSin == 1 {
				host_description[i], opponent_description[i] = "생각 또는 가치관이 잘 맞다. ", "생각 또는 가치관이 잘 맞다. "
				host_grade[i], opponent_grade[i] = 5, 5
			} else if host.Result[i].ChunGanHab.JeongIm == 1 {
				host_description[i], opponent_description[i] = "생각 또는 가치관이 잘 맞다. ", "생각 또는 가치관이 잘 맞다. "
				host_grade[i], opponent_grade[i] = 5, 5
			} else if host.Result[i].ChunGanHab.MuGye == 1 {
				host_description[i], opponent_description[i] = "생각 또는 가치관이 잘 맞다. ", "생각 또는 가치관이 잘 맞다. "
				host_grade[i], opponent_grade[i] = 5, 5
			}

			if host.Result[i].IpMyo.Exist || opponent.Result[i].IpMyo.Exist {
				if host.Result[i].IpMyo.WhichJi == 1 && opponent.Result[i].IpMyo.WhichJi == 1 {
					host_description[i] = "생각 또는 가치관이 잘 맞지만, 스스로가 상대방에게 묻힌다는 기분이다. "
					opponent_description[i] = "생각 또는 가치관이 잘 맞지만, 상대방이 나에게 묻힌다는 기분이다. "
				} else if host.Result[i].IpMyo.WhichJi == 2 && opponent.Result[i].IpMyo.WhichJi == 2 {
					host_description[i] = "생각 또는 가치관이 잘 맞지만, 스스로가 상대방에게 묻힌다는 기분이다. "
					opponent_description[i] = "생각 또는 가치관이 잘 맞지만, 상대방이 나에게 묻힌다는 기분이다. "
				}
				host_grade[i] -= 3
				opponent_grade[i] -= 3
			}
		}

		if host.Result[i].ChunGanGeok.Exist || opponent.Result[i].ChunGanGeok.Exist {
			if host.Result[i].ChunGanGeok.Exist {
				host_description[i] = "생각 또는 가치관이 안 맞으며, 상대방이 나의 생각 또는 가치관을 제어하고자 한다. "
				opponent_description[i] = "생각 또는 가치관이 안 맞으며, 내가 상대방의 생각 또는 가치관을 제어하고자 한다. "
				host_grade[i] = -5

				if host.Result[i].IpMyo.Exist || opponent.Result[i].IpMyo.Exist {
					if host.Result[i].IpMyo.WhichJi == 1 && opponent.Result[i].IpMyo.WhichJi == 1 {
						host_description[i] = "생각 또는 가치관이 잘 안 맞지만, 스스로가 상대방을 잘 받아들인다. "
						opponent_description[i] = "생각 또는 가치관이 잘 맞지만, 상대방이 나를 잘 받아들인다. "
						host_grade[i] += 3
					} else if host.Result[i].IpMyo.WhichJi == 2 && opponent.Result[i].IpMyo.WhichJi == 2 {
						host_grade[i] += 3
						host_description[i] = "생각 또는 가치관이 잘 안 맞으며, 스스로가 상대방에게 묻힌다는 기분이다. "
						opponent_description[i] = "생각 또는 가치관이 잘 안 맞으며, 상대방이 나에게 묻힌다는 기분이다. "
					}
				}

			} else if opponent.Result[i].ChunGanGeok.Exist {
				host_description[i] = "생각 또는 가치관이 잘 안 맞으며, 내가 상대방의 생각 또는 가치관을 제어하고자 한다. "
				opponent_description[i] = "생각 또는 가치관이 안 맞으며, 상대방이 나의 생각 또는 가치관을 제어하고자 한다. "
				opponent_grade[i] = -5
				if host.Result[i].IpMyo.Exist || opponent.Result[i].IpMyo.Exist {
					if host.Result[i].IpMyo.WhichJi == 1 && opponent.Result[i].IpMyo.WhichJi == 1 {
						host_description[i] = "생각 또는 가치관이 잘 안 맞지만, 상대방이 나를 잘 받아들인다. "
						opponent_description[i] = "생각 또는 가치관이 잘 맞지만, 스스로가 상대방을 잘 받아들인다. "
						opponent_grade[i] += 3
					} else if host.Result[i].IpMyo.WhichJi == 2 && opponent.Result[i].IpMyo.WhichJi == 2 {
						host_description[i] = "생각 또는 가치관이 잘 안 맞으며, 상대방이 나에게 묻힌다는 기분이다. "
						opponent_description[i] = "생각 또는 가치관이 잘 안 맞으며, 스스로가 상대방에게 묻힌다는 기분이다. "
						opponent_grade[i] += 3
					}
				}
			}
		}

		if host.Result[i].SamHab.Exist {
			host_description[i] += "실질적으로 서로가 합심해서 어떤 일을 했을 때 잘 맞다. "
			opponent_description[i] += "실실적으로 서로가 합심해서 어떤 일을 했을 때 잘 맞다. "
			host_grade[i] += 9
			opponent_grade[i] += 9
		}

		if host.Result[i].YukHab.Exist {
			host_description[i] += "서로가 부부처럼 잘 맞다. "
			opponent_description[i] += "서로가 부부처럼 잘 맞다. "
			host_grade[i] += 10
			opponent_grade[i] += 10
		}

		if host.Result[i].BangHab.Exist {
			host_description[i] += "가족처럼 잘 맞다. "
			opponent_description[i] += "가족처럼 잘 맞다. "
			host_grade[i] += 8
			opponent_grade[i] += 8
		}

		if host.Result[i].Choong.Exist {
			host_description[i] += "충돌이 일어날 수 있다. "
			opponent_description[i] += "충돌이 일어날 수 있다. "
			host_grade[i] -= 7
			opponent_grade[i] -= 7
		}

		if host.Result[i].GuiMoon.Exist {
			if host.Result[i].GuiMoon.InMi == 1 {
				host_description[i] += "한쪽이 상대방에게 집착할 수 있다. "
				opponent_description[i] += "한쪽이 상대방에게 집착할 수 있다. "
				host_grade[i] -= 3
				opponent_grade[i] -= 3
			} else if host.Result[i].GuiMoon.JaYu == 1 {
				host_description[i] += "술을 마시고 해프닝이 일어나듯, 인지하지 못할 때 문제가 생길 수 있다. "
				opponent_description[i] += "술을 마시고 해프닝이 일어나듯, 인지하지 못할 때 문제가 생길 수 있다. "
				host_grade[i] -= 3
				opponent_grade[i] -= 3
			}

		}

		if host.Result[i].GyeokGak.Exist {

		}

		if host.Result[i].Hae.Exist {

		}

		if host.Result[i].Hyung.Exist {
			host_description[i] += "서로가 만나면 어떤 분쟁이자 소소한 다툼거리가 만들어 질 수 있다. "
			opponent_description[i] += "서로가 만나면 어떤 분쟁이자 소소한 다툼거리가 만들어 질 수 있다. "
			host_grade[i] -= 3
			opponent_grade[i] -= 3
		}

		if host.Result[i].Pa.Exist {
			if host.Result[i].YukHab.Exist {

			} else if host.Result[i].Hyung.Exist {

			} else {
				host_description[i] += "서로가 만나면 자잘한 관계가 깨어질 수 있다. "
				opponent_description[i] += "서로가 만나면 자잘한 관계가 깨어질 수 있다. "
				host_grade[i] -= 1
				opponent_grade[i] -= 1
			}
		}

		if host.Result[i].WonJin.Exist {
			if host.Result[i].AmHab.Exist {
				host_description[i] += "서로가 만나면 애증의 관계가 수 있다. "
				opponent_description[i] += "서로가 만나면 애증의 관계가 수 있다. "
				host_grade[i] -= 3
				opponent_grade[i] -= 3
			} else {
				host_description[i] += "서로가 만나면 서로를 원망할 수 있다. "
				opponent_description[i] += "서로가 만나면 서로를 원망할 수 있다. "
				host_grade[i] -= 5
				opponent_grade[i] -= 5
			}
		} else {
			if host.Result[i].AmHab.Exist {
				host_description[i] += "서로가 겉으로 드러나지 않는 속 정이 있을 수 있다. "
				opponent_description[i] += "서로가 겉으로 드러나지 않는 속 정이 있을 수 있다. "
				host_grade[i] += 3
				opponent_grade[i] += 3
			}
		}

	}

	var final_host_grade float64
	var final_opponent_grade float64
	var final_host_description string
	var final_opponent_description string
	for i := 1; i < 3; i++ {
		final_host_grade += float64(host_grade[i])
		final_opponent_grade += float64(opponent_grade[i])
		if host_description[i] != "" {
			//fmt.Println(host_description[i])
			switch i {
			case 1:
				final_host_description += "각자의 " + host_description[i]
				final_opponent_description += "각자의 " + opponent_description[i]
			case 2:
				final_host_description += "서로가 살아가는 세상에서 " + host_description[i]
				final_opponent_description += "서로가 살아가는 세상에서 " + opponent_description[i]
			}
		}

	}

	final_host_grade_string := fmt.Sprintf("%.2f", math.Round(math.Abs(15+final_host_grade)/30*100.0))
	final_opponent_grade_string := fmt.Sprintf("%.2f", math.Round(math.Abs(15+final_opponent_grade)/30*100*100)/100)

	return final_host_grade_string, final_opponent_grade_string, final_host_description, final_opponent_description
}

func (sa *sajuAnalyzer) GetAnalyzerTable() *sajuAnalyzer {
	var table sajuAnalyzer
	table.Sib2Unsung = sa.Sib2Unsung
	table.Sibsung = sa.Sibsung
	return &table
}

func NewSaJuAnalyzer() *sajuAnalyzer {
	return newSaJuAnalyzer()
}

func newSaJuAnalyzer() *sajuAnalyzer {
	d, err := ioutil.ReadFile("/etc/config/sibsung.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	e, err := ioutil.ReadFile("/etc/config/sib2Unsung.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var sibsung_table []Sibsung
	var sib2Unsung_table []Chungan_Unsung

	json.Unmarshal(d, &sibsung_table)
	json.Unmarshal(e, &sib2Unsung_table)
	return &sajuAnalyzer{
		Sibsung:    sibsung_table,
		Sib2Unsung: sib2Unsung_table,
	}
}
