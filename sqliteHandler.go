package urmy_handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	pq "github.com/lib/pq"
)

type sqliteHandler struct {
	db *sqlx.DB
}

type User struct {
	LoginID     string `db:"loginid"`
	Password    string `db:"password"`
	Nickname    string `db:"nickname"`
	Name        string `db:"name"`
	Birthdate   string `db:"birthday"`
	PhoneNumber string `db:"phoneno"`
	Gender      bool   `db:"gender"`
}

type PersonSaju struct {
	LoginID   string `json:"LoginId"`
	YearChun  string `json:"YearChun"`
	YearJi    string `json:"YearJi"`
	MonthChun string `json:"MonthChun"`
	MonthJi   string `json:"MonthJi"`
	DayChun   string `json:"DayChun"`
	DayJi     string `json:"DayJi"`
	TimeChun  string `json:"TimeChun"`
	TimeJi    string `json:"TimeJi"`
	DaeunChun string `json:"DaeunChun"`
	DaeUnJi   string `json:"DaeUnJi"`
	SaeunChun string `json:"SaeunChun"`
	SaeunJi   string `json:"SaeunJi"`
}

type Saju struct {
	Gender bool
	Year   int
	Month  int
	Day    int
	Time   string
}

type DBHandler interface {
	GetUrMyUser(loginId string, password string, userdataexist bool) (bool, *User)
	AddUrMyUser(loginId string, password string, nickname string, name string, phoneNo string, genderstate bool, birthdate string) (*Saju, error)
	AddUrMyAdditionalInfo(loginId string, birthdate string) (*Saju, error)
	InputUrMySaJuInfo(loginid string, sajudata *SaJuPalJa) error
	//InputUrMySaJuDaeSaeUnInfo(loginid string, sajudata *SaJuPalJa)
	GetUrMyFriendList(phoneNo string) (*PersonSaju, error)
	GetMySaju(loginId string) (*PersonSaju, error)
	Close()
}

func NewDBHandler() DBHandler {
	return newSqliteHandler()
}

func (s *sqliteHandler) GetUrMyFriendList(phoneNo string) (*PersonSaju, error) {
	var identifier string
	var friendsaju PersonSaju
	err := s.db.Get(&identifier, "SELECT loginid FROM urmyusers WHERE phoneno=$1", phoneNo)
	if err != nil {
		fmt.Println(err)
	}

	err = s.db.Get(&friendsaju, "SELECT * FROM urmysaju WHERE loginid=$1", identifier)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		if friendsaju.LoginID == "" {
			return nil, nil
		} else {
			return &friendsaju, nil
		}
	}
}

func (s *sqliteHandler) GetUrMyUser(loginId string, password string, userdataexist bool) (bool, *User) {
	var err error
	user := User{}
	if userdataexist {
		err = s.db.Get(&user, "SELECT loginid, password, nickname FROM urmyusers WHERE loginid=$1 AND password=$2", loginId, password)
	} else {
		err = s.db.Get(&user, "SELECT loginid, password, nickname, name, phoneno, birthday, gender FROM urmyusers WHERE loginid=$1 AND password=$2", loginId, password)
	}

	if err != nil {
		fmt.Println(err)
		return false, nil
	} else {
		if userdataexist {
			return true, nil
		} else {
			return true, &user
		}
	}
}

func (s *sqliteHandler) GetMySaju(loginId string) (*PersonSaju, error) {
	var friendsaju PersonSaju
	err := s.db.Get(&friendsaju, "SELECT * FROM urmysaju WHERE loginid=$1", loginId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		if friendsaju.LoginID == "" {
			return nil, nil
		} else {
			return &friendsaju, nil
		}
	}
}

func (s *sqliteHandler) AddUrMyUser(loginId string, password string, nickname string, name string, phoneNo string, genderstate bool, birthdate string) (*Saju, error) {
	birthdateinsert := pq.QuoteLiteral(birthdate)
	rows, err := s.db.Query("INSERT INTO urmysaju (loginId) VALUES ($1)", loginId)
	if err != nil {
		panic(err)
	}

	rows, err = s.db.Query("INSERT INTO urmyusers (loginId, password, nickname, name, phoneNo, gender, birthday, createdat) VALUES ($1, $2, $3, $4, $5, $6, $7, current_timestamp)",
		loginId, password, nickname, name, phoneNo, genderstate, birthdateinsert)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	temp, err := time.Parse("2006-01-02 15:04:05", birthdate)
	if err != nil {
		return nil, err
	}
	year, month, day := temp.Date()
	hour, min, _ := temp.Clock()
	var saju Saju
	saju.Gender = genderstate
	saju.Year, saju.Month, saju.Day = year, int(month), day
	saju.Time = formatTime(hour, min)
	return &saju, nil
}

func (s *sqliteHandler) AddUrMyAdditionalInfo(loginid string, birthdate string) (*Saju, error) {
	birthdateinsert := pq.QuoteLiteral(birthdate)
	rows, err := s.db.Query("UPDATE urmyusers SET birthday=$1 WHERE loginId=$2", birthdateinsert, loginid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	temp, err := time.Parse("2006-01-02 15:04:05", birthdate)
	if err != nil {
		return nil, err
	}
	year, month, day := temp.Date()
	hour, min, _ := temp.Clock()
	var saju Saju
	saju.Year, saju.Month, saju.Day = year, int(month), day
	saju.Time = formatTime(hour, min)
	return &saju, nil
}

func formatTime(hour int, min int) string {
	formattedhour := strconv.Itoa(hour)
	formattedmin := strconv.Itoa(min)
	if formattedhour == "0" {
		formattedhour = "00"
	}
	if formattedmin == "0" {
		formattedmin = "00"
	}
	return formattedhour + ":" + formattedmin

}

func (s *sqliteHandler) InputUrMySaJuInfo(loginid string, sajudata *SaJuPalJa) error {
	rows, err := s.db.Query("UPDATE urmysaju SET yearChun=$1, yearJi=$2, monthChun=$3, monthJi=$4, dayChun=$5, dayJi=$6, timeChun=$7, timeJi=$8, daeunChun=$9, daeunJi=$10, saeunChun=$11, saeunJi=$12 WHERE loginId=$13",
		sajudata.YearChun, sajudata.YearJi, sajudata.MonthChun, sajudata.MonthJi, sajudata.DayChun, sajudata.DayJi, sajudata.TimeChun, sajudata.TimeJi, sajudata.DaeUnChun, sajudata.DaeUnJi, sajudata.SaeUnChun, sajudata.SaeUnJi, loginid)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer rows.Close()
	return nil
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}

/*
const (
	host     = "192.168.10.150"
	port     = 5432
	user     = "koreaogh"
	password = "ogh1898"
	dbname   = "urmydb"
)
*/

const (
	host     = "172.31.210.221"
	port     = 5432
	user     = "koreaogh"
	password = "ogh1898"
	dbname   = "urmydb"
)

func newSqliteHandler() DBHandler {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}

	errPing := database.Ping
	if errPing != nil {
		//panic(errPing)
		fmt.Println(err)
	} else {
		fmt.Println("attached")
	}

	return &sqliteHandler{db: database}
}
