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
	"strconv"
	"time"
)

type cryptoEnv struct {
	Port              string
	JwtKey            string
	JwtExpiryHours    time.Duration
	RandomByteLength  string
	BcryptHashRounds  int
	AesPassphrase     string
	HoneybadgerApiKey string
}

func (e cryptoEnv) Sanitize() {
	sanitize(reflect.ValueOf(e))
}

func newCryptoEnv() *cryptoEnv {
	jwtExpiryHours := os.Getenv("FIZZ_CRYPTO_JWT_EXPIRY_HOURS")
	jwtExpiryHoursNum := 0
	if jwtExpiryHours != "" {
		jwtExpiryHoursNum, _ = strconv.Atoi(jwtExpiryHours)
	}

	bcryptHashRounds := os.Getenv("FIZZ_CRYPTO_BCRYPT_HASH_ROUNDS")
	bcryptHashRoundsNum := 0
	if bcryptHashRounds != "" {
		bcryptHashRoundsNum, _ = strconv.Atoi(bcryptHashRounds)
	}

	return &cryptoEnv{
		Port:              os.Getenv("FIZZ_CRYPTO_SVC_PORT"),
		JwtKey:            os.Getenv("FIZZ_CRYPTO_JWT_KEY"),
		JwtExpiryHours:    time.Duration(jwtExpiryHoursNum),
		RandomByteLength:  os.Getenv("FIZZ_CRYPTO_RANDOM_BYTE_LENGTH"),
		BcryptHashRounds:  bcryptHashRoundsNum,
		AesPassphrase:     os.Getenv("FIZZ_CRYPTO_AES_PASSPHRASE"),
		HoneybadgerApiKey: os.Getenv("FIZZ_CRYPTO_HONEYBADGER_API_KEY"),
	}
}
