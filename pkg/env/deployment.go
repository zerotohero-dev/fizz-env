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
	"strings"
)

type DeploymentType string

const Production DeploymentType = "production"
const Development DeploymentType = "development"
const Staging DeploymentType = "staging"

type deploymentEnv struct {
	Type DeploymentType
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

func newDeploymentEnv() *deploymentEnv {
	deploymentType := deploymentTypeFromEnv(os.Getenv("FIZZ_DEPLOYMENT_TYPE"))
	return &deploymentEnv{Type: deploymentType}
}
