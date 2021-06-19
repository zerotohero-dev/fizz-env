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

type idmDeps struct {
	Crypto *cryptoEnv
}

type idmEnv struct {
	Port           string
	UsersTableName string
	VerifiedUrl    string
	Dependencies   idmDeps
}

func (e idmEnv) Sanitize() {
	sanitize(reflect.ValueOf(e))
	// IDM depends on crypto, so sanitize it too.
	sanitize(reflect.ValueOf(e.Dependencies.Crypto))
}

func newIdmEnv(deps idmDeps) *idmEnv {
	return &idmEnv{
		Port:           os.Getenv("FIZZ_IDM_SVC_PORT"),
		UsersTableName: os.Getenv("FIZZ_IDM_USERS_TABLE_NAME"),
		VerifiedUrl:    os.Getenv("FIZZ_IDM_VERIFIED_URL"),
		Dependencies:   deps,
	}
}
