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
	"strconv"
	"strings"
	"time"
)

type envCrypto struct {
	Port              string
	JwtKey            string
	JwtExpiryHours    time.Duration
	RandomByteLength  string
	BcryptHashRounds  int
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
		typeName := v.Field(i).Type().Name()

		if typeName == "string" {
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

			continue
		}

		if typeName == "bool" {
			continue
		}

		val := v.Field(i).Int()
		objName := v.Type().Name()
		fieldName := v.Type().Field(i).Name

		if val == 0 {
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

func (e FizzEnv) IsDevelopment() bool {
	return e.Deployment.Type == Development
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

	res := &FizzEnv{
		Crypto: envCrypto{
			Port:              os.Getenv("FIZZ_CRYPTO_SVC_PORT"),
			JwtKey:            os.Getenv("FIZZ_CRYPTO_JWT_KEY"),
			JwtExpiryHours:    time.Duration(jwtExpiryHoursNum),
			RandomByteLength:  os.Getenv("FIZZ_CRYPTO_RANDOM_BYTE_LENGTH"),
			BcryptHashRounds:  bcryptHashRoundsNum,
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
