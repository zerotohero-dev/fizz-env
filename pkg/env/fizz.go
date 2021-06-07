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
	"fmt"
	"os"
	"reflect"
	"strings"
)

type envCrypto struct {
	PortSvcCrypto    string
	JwtKey           string
	RandomByteLength string
	BcryptHashRounds string
	AesPassphrase    string
}

type envLogging struct {
	Destination string
	UseStdOut   bool
}

type FizzEnv struct {
	Crypto envCrypto
	Log    envLogging
}

func sanitize(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).String()
		name := v.Type().Field(i).Name

		if val == "" {
			panic(
				fmt.Sprintf(
					"The environment variable that corresponds to '%s' is not defined.", name,
				),
			)
		}
	}
}

func (e FizzEnv) SanitizeCrypto() {
	sanitize(reflect.ValueOf(e.Crypto))
}

func (e FizzEnv) SanitizeLog() {
	sanitize(reflect.ValueOf(e.Log))
}

func New() *FizzEnv {
	res := &FizzEnv{
		Crypto: envCrypto{
			PortSvcCrypto:    os.Getenv("FIZZ_PORT_SVC_CRYPTO"),
			JwtKey:           os.Getenv("FIZZ_JWT_KEY"),
			RandomByteLength: os.Getenv("FIZZ_RANDOM_BYTE_LENGTH"),
			BcryptHashRounds: os.Getenv("FIZZ_BCRYPT_HASH_ROUNDS"),
			AesPassphrase:    os.Getenv("FIZZ_AES_PASSPHRASE"),
		},
		Log: envLogging{
			Destination: os.Getenv("FIZZ_LOG_DESTINATION"),
			UseStdOut: strings.ToLower(os.Getenv("FIZZ_LOG_USE_STD_OUT")) == "true" ||
				strings.ToLower(os.Getenv("FIZZ_LOG_USE_STD_OUT")) == "yes",
		},
	}

	return res
}
