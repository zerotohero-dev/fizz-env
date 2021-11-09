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

type spireEnv struct {
	SocketPath      string
	ServerAddress   string
	MtlsTimeoutSecs time.Duration
}

func (e spireEnv) Sanitize() { sanitize(reflect.ValueOf(e)) }

func newSpireEnv() *spireEnv {
	const mtlsTimeoutSecsDefault = 300

	mtlsTimeoutSecs := os.Getenv("FIZZ_SPIRE_MTLS_TIMEOUT_SECS")
	mtlsTimeoutSecsNum := mtlsTimeoutSecsDefault
	if mtlsTimeoutSecs != "" {
		mtlsTimeoutSecsNum, _ = strconv.Atoi(mtlsTimeoutSecs)
	}

	return &spireEnv{
		SocketPath:      os.Getenv("FIZZ_SPIRE_SOCKET_PATH"),
		ServerAddress:   os.Getenv("FIZZ_SPIRE_SERVER_ADDRESS"),
		MtlsTimeoutSecs: time.Duration(mtlsTimeoutSecsNum) * time.Second,
	}
}
