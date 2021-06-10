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
	Port              string
	JwtKey            string
	RandomByteLength  string
	BcryptHashRounds  string
	AesPassphrase     string
	HoneybadgerApiKey string
}

type envLogging struct {
	Destination string
}

type DeploymentType string

const Production DeploymentType = "production"
const Development DeploymentType = "development"
const Staging DeploymentType = "staging"

type envDeployment struct {
	Type DeploymentType
}

type FizzEnv struct {
	Crypto     envCrypto
	Log        envLogging
	Deployment envDeployment
}

func sanitize(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).String()
		objName := v.Type().Name()
		fieldName := v.Type().Field(i).Name

		if val == "" {
			panic(
				fmt.Sprintf(
					"The environment variable that corresponds to "+
						"'%s.%s' is not defined.", objName, fieldName,
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

func deploymentTypeFromEnv(val string) DeploymentType {
	if strings.ToLower(val) == string(Production) {
		return Production
	}
	if strings.ToLower(val) == string(Staging) {
		return Staging
	}
	if strings.ToLower(val) == string(Development) {
		return Development
	}
	panic("FIZZ_DEPLOYMENT_TYPE must be a value " +
		"from {'production','staging','development'}.")
}

func New() *FizzEnv {
	deploymentType := deploymentTypeFromEnv(os.Getenv("FIZZ_DEPLOYMENT_TYPE"))

	res := &FizzEnv{
		Crypto: envCrypto{
			Port:              os.Getenv("FIZZ_CRYPTO_SVC_PORT"),
			JwtKey:            os.Getenv("FIZZ_CRYPTO_JWT_KEY"),
			RandomByteLength:  os.Getenv("FIZZ_CRYPTO_RANDOM_BYTE_LENGTH"),
			BcryptHashRounds:  os.Getenv("FIZZ_CRYPTO_BCRYPT_HASH_ROUNDS"),
			AesPassphrase:     os.Getenv("FIZZ_CRYPTO_AES_PASSPHRASE"),
			HoneybadgerApiKey: os.Getenv("FIZZ_CRYPTO_HONEYBADGER_API_KEY"),
		},
		Log: envLogging{
			Destination: os.Getenv("FIZZ_LOG_DESTINATION"),
		},
		Deployment: envDeployment{Type: deploymentType},
	}

	return res
}
