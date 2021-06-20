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
)

type mailerEnv struct {
	Port                     string
	HoneybadgerApiKey        string
	EmailVerificationBaseUrl string
	WebAppHostBaseUrl        string
	MailgunDomain            string
	MailgunApiKey            string
}

func (e mailerEnv) Sanitize() {
	sanitize(reflect.ValueOf(e))
}

func newMailerEnv() *mailerEnv {
	return &mailerEnv{
		Port:                     os.Getenv("FIZZ_MAILER_SVC_PORT"),
		HoneybadgerApiKey:        os.Getenv("FIZZ_MAILER_HONEYBADGER_API_KEY"),
		EmailVerificationBaseUrl: os.Getenv("FIZZ_MAILER_EMAIL_VERIFICATION_BASE_URL"),
		WebAppHostBaseUrl:        os.Getenv("FIZZ_MAILER_WEB_APP_HOST_BASE_URL"),
		MailgunDomain:            os.Getenv("FIZZ_MAILER_MAILGUN_DOMAIN"),
		MailgunApiKey:            os.Getenv("FIZZ_MAILER_MAILGUN_API_KEY"),
	}
}
