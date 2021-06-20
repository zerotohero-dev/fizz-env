/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
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
	Port              string
	HoneybadgerApiKey string
	UsersTableName    string
	VerifiedUrl       string
	CryptoEndpointUrl string
}

func (e idmEnv) Sanitize() {
	sanitize(reflect.ValueOf(e))
	// IDM depends on crypto, so sanitize it too.

	//// TODO: this looks smelly. But it’s reflection after all. Maybe reconsider?
	//// I’d prefer to manually type a bunch of key instead of the code
	//// becoming even more cryptic.
	//ce := cryptoEnv{
	//	Port:              e.Dependencies.Crypto.Port,
	//	JwtKey:            e.Dependencies.Crypto.JwtKey,
	//	JwtExpiryHours:    e.Dependencies.Crypto.JwtExpiryHours,
	//	RandomByteLength:  e.Dependencies.Crypto.RandomByteLength,
	//	BcryptHashRounds:  e.Dependencies.Crypto.BcryptHashRounds,
	//	AesPassphrase:     e.Dependencies.Crypto.AesPassphrase,
	//	HoneybadgerApiKey: e.Dependencies.Crypto.HoneybadgerApiKey,
	//}
	//
	//sanitize(reflect.ValueOf(ce))
}

func newIdmEnv() *idmEnv {
	return &idmEnv{
		Port:              os.Getenv("FIZZ_IDM_SVC_PORT"),
		UsersTableName:    os.Getenv("FIZZ_IDM_USERS_TABLE_NAME"),
		VerifiedUrl:       os.Getenv("FIZZ_IDM_VERIFIED_URL"),
		HoneybadgerApiKey: os.Getenv("FIZZ_IDM_HONEYBADGER_API_KEY"),
		CryptoEndpointUrl: os.Getenv("FIZZ_IDM_CRYPTO_ENDPOINT_URL"),
	}
}
