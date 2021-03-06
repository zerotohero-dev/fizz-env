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
	SocketPath         string
	ServerAddress      string
	MtlsTimeout        time.Duration
	AppPrefixFizz      string
	AppNameFizzDefault string
	AppTrustDomainFizz string
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
		SocketPath:         os.Getenv("SPIFFE_ENDPOINT_SOCKET"),
		ServerAddress:      os.Getenv("FIZZ_SPIRE_SERVER_ADDRESS"),
		AppPrefixFizz:      os.Getenv("FIZZ_SPIRE_APP_PREFIX"),
		AppNameFizzDefault: os.Getenv("FIZZ_SPIRE_DEFAULT_APP_NAME"),
		AppTrustDomainFizz: os.Getenv("FIZZ_SPIRE_APP_TRUST_DOMAIN"),
		MtlsTimeout:        time.Duration(mtlsTimeoutSecsNum) * time.Second,
	}
}
