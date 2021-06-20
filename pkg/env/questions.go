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

import "os"

type questionsEnv struct {
	Port              string
	HoneybadgerApiKey string
	DataPath          string
}

func newQuestionsEnv() *questionsEnv {
	return &questionsEnv{
		Port:              os.Getenv("FIZZ_QUESTIONS_SVC_PORT"),
		HoneybadgerApiKey: os.Getenv("FIZZ_QUESTIONS_HONEYBADGER_API_KEY"),
		DataPath:          os.Getenv("FIZZ_QUESTIONS_DATA_PATH"),
	}
}
