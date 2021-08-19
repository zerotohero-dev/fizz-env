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

type storeEnv struct {
	Port                string
	HoneybadgerApiKey   string
	CryptoEndpointUrl   string
	IdmEndpointUrl      string
	UsersTableName      string
	PlansTableName      string
	StripePrivateKey    string
	StripePublicKey     string
	SubscribeApiUrl     string
	SubscribeSuccessUrl string
	SubscribeErrorUrl   string
	PathPrefix          string
	ServiceName         string
}

func (e storeEnv) Sanitize() { sanitize(reflect.ValueOf(e)) }

func newStoreEnv() *storeEnv {
	return &storeEnv{
		Port:                os.Getenv("FIZZ_STORE_SVC_PORT"),
		HoneybadgerApiKey:   os.Getenv("FIZZ_STORE_HONEYBADGER_API_KEY"),
		CryptoEndpointUrl:   os.Getenv("FIZZ_STORE_CRYPTO_ENDPOINT_URL"),
		IdmEndpointUrl:      os.Getenv("FIZZ_STORE_IDM_ENDPOINT_URL"),
		UsersTableName:      os.Getenv("FIZZ_STORE_USERS_TABLE_NAME"),
		PlansTableName:      os.Getenv("FIZZ_STORE_PLANS_TABLE_NAME"),
		StripePrivateKey:    os.Getenv("FIZZ_STORE_STRIPE_PRIVATE_KEY"),
		StripePublicKey:     os.Getenv("FIZZ_STORE_STRIPE_PUBLIC_KEY"),
		SubscribeApiUrl:     os.Getenv("FIZZ_STORE_SUBSCRIBE_API_URL"),
		SubscribeSuccessUrl: os.Getenv("FIZZ_STORE_SUBSCRIBE_SUCCESS_URL"),
		SubscribeErrorUrl:   os.Getenv("FIZZ_STORE_SUBSCRIBE_ERROR_URL"),
		PathPrefix:          os.Getenv("FIZZ_STORE_PATH_PREFIX"),
		ServiceName:         os.Getenv("FIZZ_STORE_SERVICE_NAME"),
	}
}
