/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package env

import (
	"os"
	"reflect"
	"strconv"
	"time"
)

type idmEnv struct {
	Port                 string
	HoneybadgerApiKey    string
	DbName               string
	DbConnectionString   string
	UsersTableName       string
	VerifiedUrl          string
	CryptoEndpointUrl    string
	MailerEndpointUrl    string
	PathPrefix           string
	ServiceName          string
	JwtCookieExpiryHours time.Duration
	LaunchState        string  
}

func (e idmEnv) Sanitize() { sanitize(reflect.ValueOf(e)) }

func newIdmEnv() *idmEnv {
	jwtCookieExpiryHours := os.Getenv("FIZZ_IDM_JWT_COOKIE_EXPIRY_HOURS")
	jwtCookieExpiryHoursNum := 0
	if jwtCookieExpiryHours != "" {
		jwtCookieExpiryHoursNum, _ = strconv.Atoi(jwtCookieExpiryHours)
	}

	return &idmEnv{
		Port:                 os.Getenv("FIZZ_IDM_SVC_PORT"),
		UsersTableName:       os.Getenv("FIZZ_IDM_USERS_TABLE_NAME"),
		VerifiedUrl:          os.Getenv("FIZZ_IDM_VERIFIED_URL"),
		HoneybadgerApiKey:    os.Getenv("FIZZ_IDM_HONEYBADGER_API_KEY"),
		DbName:               os.Getenv("FIZZ_IDM_DB_NAME"),
		DbConnectionString:   os.Getenv("FIZZ_IDM_DB_CONNECTION_STRING"),
		CryptoEndpointUrl:    os.Getenv("FIZZ_IDM_CRYPTO_ENDPOINT_URL"),
		MailerEndpointUrl:    os.Getenv("FIZZ_IDM_MAILER_ENDPOINT_URL"),
		PathPrefix:           os.Getenv("FIZZ_IDM_PATH_PREFIX"),
		ServiceName:          os.Getenv("FIZZ_IDM_SERVICE_NAME"),
		JwtCookieExpiryHours: time.Duration(jwtCookieExpiryHoursNum),
		LaunchState:          os.Getenv("FIZZ_IDM_LAUNCH_STATE"),
	}
}
