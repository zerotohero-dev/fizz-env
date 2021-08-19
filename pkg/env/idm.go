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
)

type idmEnv struct {
	Port               string
	HoneybadgerApiKey  string
	DbName             string
	DbConnectionString string
	UsersTableName     string
	VerifiedUrl        string
	CryptoEndpointUrl  string
	MailerEndpointUrl  string
	PathPrefix         string
	ServiceName        string
}

func (e idmEnv) Sanitize() { sanitize(reflect.ValueOf(e)) }

func newIdmEnv() *idmEnv {
	return &idmEnv{
		Port:               os.Getenv("FIZZ_IDM_SVC_PORT"),
		UsersTableName:     os.Getenv("FIZZ_IDM_USERS_TABLE_NAME"),
		VerifiedUrl:        os.Getenv("FIZZ_IDM_VERIFIED_URL"),
		HoneybadgerApiKey:  os.Getenv("FIZZ_IDM_HONEYBADGER_API_KEY"),
		DbName:             os.Getenv("FIZZ_IDM_DB_NAME"),
		DbConnectionString: os.Getenv("FIZZ_IDM_DB_CONNECTION_STRING"),
		CryptoEndpointUrl:  os.Getenv("FIZZ_IDM_CRYPTO_ENDPOINT_URL"),
		MailerEndpointUrl:  os.Getenv("FIZZ_IDM_MAILER_ENDPOINT_URL"),
		PathPrefix:         os.Getenv("FIZZ_IDM_PATH_PREFIX"),
		ServiceName:        os.Getenv("FIZZ_SERVICE_NAME"),
	}
}
