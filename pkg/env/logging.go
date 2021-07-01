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

type loggingEnv struct {
	Destination string
}

func (e loggingEnv) Sanitize() {
	sanitize(reflect.ValueOf(e))
}

func newLoggingEnv() *loggingEnv {
	return &loggingEnv{
		Destination: os.Getenv("FIZZ_LOG_DESTINATION"),
	}
}
