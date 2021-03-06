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

type cryptoEnv struct {
	Port              string
	PathPrefix        string
	HoneybadgerApiKey string
	JwtKey            string
	JwtExpiration     time.Duration
	RandomByteLength  int
	BcryptHashRounds  int
	AesPassphrase     string
	ServiceName       string
	MtlsServerAddress string
}

func (e cryptoEnv) Sanitize() { sanitize(reflect.ValueOf(e)) }

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

	randomByteLength := os.Getenv("FIZZ_CRYPTO_RANDOM_BYTE_LENGTH")
	randomByteLengthNum := 0
	if randomByteLength != "" {
		randomByteLengthNum, _ = strconv.Atoi(randomByteLength)
	}

	return &cryptoEnv{
		Port:              os.Getenv("FIZZ_CRYPTO_SVC_PORT"),
		JwtKey:            os.Getenv("FIZZ_CRYPTO_JWT_KEY"),
		JwtExpiration:     time.Duration(jwtExpiryHoursNum) * time.Hour,
		RandomByteLength:  randomByteLengthNum,
		BcryptHashRounds:  bcryptHashRoundsNum,
		AesPassphrase:     os.Getenv("FIZZ_CRYPTO_AES_PASSPHRASE"),
		HoneybadgerApiKey: os.Getenv("FIZZ_CRYPTO_HONEYBADGER_API_KEY"),
		PathPrefix:        os.Getenv("FIZZ_CRYPTO_PATH_PREFIX"),
		ServiceName:       os.Getenv("FIZZ_CRYPTO_SERVICE_NAME"),
		MtlsServerAddress: os.Getenv("FIZZ_CRYPTO_MTLS_SERVER_ADDRESS"),
	}
}
