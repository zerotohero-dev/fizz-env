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

type spireEnv struct {
	SocketPath    string
	ServerAddress string
}

func (e spireEnv) Sanitize() { sanitize(reflect.ValueOf(e)) }

func newSpireEnv() *spireEnv {
	return &spireEnv{
		SocketPath:    os.Getenv("FIZZ_SPIRE_SOCKET_PATH"),
		ServerAddress: os.Getenv("FIZZ_SPIRE_SERVER_ADDRESS"),
	}
}
