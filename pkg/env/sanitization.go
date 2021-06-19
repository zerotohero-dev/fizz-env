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
	"reflect"
)

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
