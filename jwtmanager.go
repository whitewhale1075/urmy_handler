package urmy_handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"gopkg.in/redis.v5"
)

type AccessDetails struct {
	AccessUuid string
	UserId     string
}

type RefreshDetails struct {
	RefreshUuid string
	UserId      string
}

type jwtredisHandler struct {
	ac *redis.Client
	rc *redis.Client
}

type AccessTokenDetails struct {
	AccessToken string
	AccessUuid  string
	AtExpires   int64
}

type RefreshTokenDetails struct {
	RefreshToken string
	RefreshUuid  string
	RtExpires    int64
}

type JWTHandler interface {
	GernerateAccessJWT(userid string) (*AccessTokenDetails, error)
	GernerateRefreshJWT(userid string) (*RefreshTokenDetails, error)
	ExtractAccessTokenMetadata(r *http.Request) (*AccessDetails, error)
	ExtractRefreshTokenMetadata(r *http.Request) (*RefreshDetails, error)
	FetchAccessAuth(authD *AccessDetails) (string, error)
	FetchRefreshAuth(authD *RefreshDetails) (string, error)
	CreateAccessAuth(userid string, td *AccessTokenDetails) error
	CreateRefreshAuth(userid string, td *RefreshTokenDetails) error
	DeleteAccessAuth(givenAccessUuid string) (int64, error)
}

var mySigningKey = []byte("mysupersecretphrase")

func (rc *jwtredisHandler) GernerateAccessJWT(userid string) (*AccessTokenDetails, error) {
	td := &AccessTokenDetails{}
	td.AtExpires = time.Now().Add(time.Second * 240).Unix()
	u, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	td.AccessUuid = u.String()

	atclaims := jwt.MapClaims{}
	atclaims["authorized"] = true
	atclaims["access_uuid"] = td.AccessUuid
	atclaims["user_id"] = userid
	atclaims["exp"] = td.AtExpires
	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, atclaims)
	td.AccessToken, err = atoken.SignedString(mySigningKey)

	if err != nil {
		return nil, fmt.Errorf("something went wrong: %s", err.Error())
	}

	return td, nil
}

func (rc *jwtredisHandler) GernerateRefreshJWT(userid string) (*RefreshTokenDetails, error) {
	td := &RefreshTokenDetails{}

	td.RtExpires = time.Now().Add(time.Second * 600).Unix()
	u2, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	td.RefreshUuid = u2.String()

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rtoken.SignedString(mySigningKey)

	if err != nil {
		return nil, fmt.Errorf("something went wrong: %s", err.Error())
	}

	return td, nil
}

func ExtractAccessToken(r *http.Request) string {
	bearToken := r.Header.Get("AuthorizationAccess")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 1 {
		return strArr[0]
	}
	return ""
}

func VerifyAccessToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractAccessToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(mySigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractRefreshToken(r *http.Request) string {
	bearToken := r.Header.Get("AuthorizationRefresh")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 1 {
		return strArr[0]
	}
	return ""
}

func VerifyRefreshToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractRefreshToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(mySigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (rc *jwtredisHandler) ExtractAccessTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyAccessToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		//userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		userId := claims["user_id"].(string)
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func (rc *jwtredisHandler) ExtractRefreshTokenMetadata(r *http.Request) (*RefreshDetails, error) {
	token, err := VerifyRefreshToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			return nil, err
		}
		//userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		userId := claims["user_id"].(string)

		return &RefreshDetails{
			RefreshUuid: refreshUuid,
			UserId:      userId,
		}, nil
	}
	return nil, err
}

func (rc *jwtredisHandler) FetchAccessAuth(authD *AccessDetails) (string, error) {
	userid, err := rc.ac.Get(authD.AccessUuid).Result()
	if err != nil {
		return "", fmt.Errorf("fetch access auth : %s", err)
	}
	//userID, _ := strconv.ParseUint(userid, 10, 64)

	return userid, nil
}

func (rc *jwtredisHandler) FetchRefreshAuth(authD *RefreshDetails) (string, error) {
	userid, err := rc.rc.Get(authD.RefreshUuid).Result()
	if err != nil {
		return "", fmt.Errorf("fetch refresh auth : %s", err)
	}
	//userID, _ := strconv.ParseUint(userid, 10, 64)

	userID := userid
	return userID, nil
}

func (rc *jwtredisHandler) CreateAccessAuth(userid string, td *AccessTokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	now := time.Now()

	errAccess := rc.ac.Set(td.AccessUuid, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	return nil
}

func (rc *jwtredisHandler) CreateRefreshAuth(userid string, td *RefreshTokenDetails) error {
	//converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errRefresh := rc.rc.Set(td.RefreshUuid, userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (rc *jwtredisHandler) ResetAccessAuth(authD *AccessDetails) error {
	//converting Unix to UTC(to Time object)
	rt := time.Duration(time.Minute * 15)
	//.AtExpires = time.Now().Add(time.Minute * 15).Unix()

	errRefresh := rc.ac.Expire(authD.AccessUuid, rt).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (rc *jwtredisHandler) DeleteAccessAuth(givenAccessUuid string) (int64, error) {
	accessdeleted, err := rc.ac.Del(givenAccessUuid).Result()
	if err != nil {
		return 0, err
	}
	return accessdeleted, nil
}

func (rc *jwtredisHandler) DeleteRefreshAuth(givenRefreshUuid string) (int64, error) {
	refreshdeleted, err := rc.rc.Del(givenRefreshUuid).Result()
	if err != nil {
		return 0, err
	}
	return refreshdeleted, nil
}

func NewJWTHandler() JWTHandler {
	return jwtRedisHandler()
}

func jwtRedisHandler() JWTHandler {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.10.160:6379",
		//Addr:     "172.31.210.238:6379",
		Password: "qwer1234", // no password set
		DB:       0,          // use default DB
	})
	ping(rdb)

	rdb2 := redis.NewClient(&redis.Options{
		Addr: "192.168.10.161:6379",
		//Addr:     "172.31.210.110:6379",
		Password: "qwer1234", // no password set
		DB:       0,          // use default DB
	})
	ping(rdb2)
	return &jwtredisHandler{ac: rdb, rc: rdb2}
}

func ping(client *redis.Client) error {
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return err

	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}

/*
func (rc *jwtredisHandler) Refresh(r *http.Request) string {
	accesstokenAuth, err := rc.ExtractAccessTokenMetadata(r)
	if err != nil {
		panic(err)
	}
	accessUserid, err := rc.FetchAccessAuth(accesstokenAuth) //15분이 안지났을때
	if err != nil {
		refreshtokenAuth, err := rc.ExtractRefreshTokenMetadata(r)
		if err != nil {
			panic(err)
		} else {
			refreshUserid, err := rc.FetchRefreshAuth(refreshtokenAuth)
			if err != nil {
				return ""
			} else {
				tdac, err := rc.GernerateAccessJWT(accesstokenAuth.UserId)
				if err != nil {
					panic(err)
				}
				rc.CreateAccessAuth(accessUserid, tdac)
				return tdac.AccessToken
			}
		}
	} else {
		//로그인 지속
		rc.ResetAccessAuth(accesstokenAuth)
		return accessuserTokenvalue
	}

			//Delete the previous Refresh Token
			deleted, delErr := rc.DeleteAuth(td.AccessUuid, td.RefreshUuid)
			if delErr != nil || deleted == 0 { //if any goes wrong

				return
			}
			//Create new pairs of refresh and access tokens
			ts, createErr := rc.GernerateJWT(userId)
			if createErr != nil {

				return
			}
			//save the tokens metadata to redis
			saveErr := rc.CreateAuth(userId, ts)
			if saveErr != nil {

				return
			}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}


}
*/
