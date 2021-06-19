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

type FizzEnv struct {
	Crypto     cryptoEnv
	Log        loggingEnv
	Deployment deploymentEnv
	Idm        idmEnv
}

func (e FizzEnv) IsDevelopment() bool {
	return e.Deployment.Type == Development
}

func New() *FizzEnv {
	cEnv := newCryptoEnv()
	return &FizzEnv{
		Crypto:     *cEnv,
		Log:        *newLoggingEnv(),
		Deployment: *newDeploymentEnv(),
		Idm:        *newIdmEnv(idmDeps{Crypto: cEnv}),
	}
}
