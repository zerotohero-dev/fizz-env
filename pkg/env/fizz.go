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

type FizzEnv struct {
	Spire      spireEnv
	Crypto     cryptoEnv
	Log        loggingEnv
	Deployment deploymentEnv
	Idm        idmEnv
	Mailer     mailerEnv
	Questions  questionsEnv
	Store      storeEnv
}

type FizzIdmEnv struct {
	Spire  spireEnv
	Crypto cryptoEnv
	Log    loggingEnv
	Idm    idmEnv
}

type FizzLogEnv struct {
}

func (e FizzEnv) IsDevelopment() bool {
	return e.Deployment.Type == Development
}

func New() *FizzEnv {
	return &FizzEnv{
		Spire:      *newSpireEnv(),
		Crypto:     *newCryptoEnv(),
		Log:        *newLoggingEnv(),
		Deployment: *newDeploymentEnv(),
		Idm:        *newIdmEnv(),
		Mailer:     *newMailerEnv(),
		Questions:  *newQuestionsEnv(),
		Store:      *newStoreEnv(),
	}
}
