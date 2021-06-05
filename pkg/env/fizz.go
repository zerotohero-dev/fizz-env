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
)

type envCrypto struct {
	PortSvcCrypto    string
	JwtKey           string
	RandomByteLength string
	BcryptHashRounds string
	AesPassphrase    string
}

type FizzEnv struct {
	Crypto envCrypto
}

func (e FizzEnv) SanitizeCrypto()  {
	v := reflect.ValueOf(e.Crypto)

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

func New() FizzEnv {
	res := FizzEnv{
		Crypto: envCrypto{
			PortSvcCrypto:    os.Getenv("FIZZ_PORT_SVC_CRYPTO"),
			JwtKey:           os.Getenv("FIZZ_JWT_KEY"),
			RandomByteLength: os.Getenv("FIZZ_RANDOM_BYTE_LENGTH"),
			BcryptHashRounds: os.Getenv("FIZZ_BCRYPT_HASH_ROUNDS"),
			AesPassphrase:    os.Getenv("FIZZ_AES_PASSPHRASE"),
		},
	}

	return res
}