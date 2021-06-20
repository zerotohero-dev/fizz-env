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
	Mailer     mailerEnv
	Notes      notesEnv
	Questions  questionsEnv
	Store      storeEnv
}

func (e FizzEnv) IsDevelopment() bool {
	return e.Deployment.Type == Development
}

func New() *FizzEnv {
	return &FizzEnv{
		Crypto:     *newCryptoEnv(),
		Log:        *newLoggingEnv(),
		Deployment: *newDeploymentEnv(),
		Idm:        *newIdmEnv(),
		Mailer:     *newMailerEnv(),
		Notes:      *newNotesEnv(),
		Questions:  *newQuestionsEnv(),
		Store:      *newStoreEnv(),
	}
}
