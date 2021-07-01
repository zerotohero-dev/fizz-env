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

type notesEnv struct {
	Port              string
	HoneybadgerApiKey string
	DataPath          string
}

func (e notesEnv) Sanitize() {
	sanitize(reflect.ValueOf(e))
}

func newNotesEnv() *notesEnv {
	return &notesEnv{
		Port:              os.Getenv("FIZZ_NOTES_SVC_PORT"),
		HoneybadgerApiKey: os.Getenv("FIZZ_NOTES_HONEYBADGER_API_KEY"),
		DataPath:          os.Getenv("FIZZ_NOTES_DATA_PATH"),
	}
}
