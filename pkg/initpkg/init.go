package initpkg

import (
	"log"

	"github.com/medtune/beta-platform/pkg/config"
	"github.com/medtune/beta-platform/pkg/secret"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/store"
	"github.com/medtune/beta-platform/pkg/store/db"
	"github.com/medtune/go-utils/random"
)

// Init packages from configuration file
func InitFromFile(file string) error {
	config, err := config.LoadConfigFromPath(file)
	if err != nil {
		return err
	}
	err = InitFromConfig(config)
	if err != nil {
		return err
	}
	return nil
}

func InitFromConfig(c *config.StartupConfig) error {
	if err := initSession(c.Session); err != nil {
		return err
	}
	if err := initStore(c.Database, c.Meta.IsProd); err != nil {
		return err
	}
	if err := initSecrets(c.Secrets); err != nil {
		return err
	}

	return nil
}

func initSession(c *config.Session) error {
	var s string
	if !c.Random {
		session.SessID = c.Name
		s = c.Secret
	} else {
		session.SessID = random.Alpha(10)
		s = random.Alpha(10)
	}
	log.Printf("Session secrets:\n\t-> %v\n\t-> %v\n", session.SessID, s)
	session.Store = session.NewStore(s)
	return nil
}

func initStore(c *config.Database, prod bool) error {
	// Choose database
	var database string
	if prod {
		database = c.Prod
	} else {
		database = c.Test
	}

	// Create store engine
	engine, err := store.New(db.ConnStr{
		Host:     c.Creds.Host,
		Database: database,
		User:     c.Creds.User,
		Password: c.Creds.Password,
		Port:     c.Creds.Port,
		SslMode:  0,
		Maxconn:  2,
	})
	if err != nil {
		return err
	}

	store.Agent = engine
	return nil
}

func initSecrets(c *config.Secrets) error {
	if c.Signup == nil {
		// Need warn
		return nil
	}

	for _, s := range c.Signup {
		secret.Register("signup", s)
		//log.Printf("Registred signup key: %s", s)
	}

	return nil
}

func initCapsules() error {
	return nil
}
