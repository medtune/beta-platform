package initpkg

import (
	"fmt"
	"log"

	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
	"github.com/medtune/go-utils/random"

	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/config"
	"github.com/medtune/beta-platform/pkg/service/capsul"
	"github.com/medtune/beta-platform/pkg/service/secret"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/store"
	"github.com/medtune/beta-platform/pkg/store/db"
)

// InitFromFile from configuration file
func InitFromFile(file string) error {
	// Read configuration from file
	config, err := config.LoadConfigFromPath(file)
	if err != nil {
		return err
	}

	// Init from config
	if err = InitFromConfig(config); err != nil {
		return err
	}

	return nil
}

// InitFromConfig uration file
func InitFromConfig(c *config.StartupConfig) error {
	// Check version
	if c.Meta.Version != pkg.VERSION {
		return fmt.Errorf("Configuration meta version did'nt match\n\tpackage version: %v\n\tconfigs version: %v", pkg.VERSION, c.Meta.Version)
	}

	// init package pkg/session
	if err := initSession(c.Session); err != nil {
		return err
	}

	// init package pkg/store + pkg/store/db
	if err := initStore(c.Database, c.Meta.IsProd); err != nil {
		return err
	}

	// init package pkg/secret
	if err := initSecrets(c.Secrets); err != nil {
		return err
	}

	// init package pkg/service/capsul
	if err := initCapsulClients(c.Capsul); err != nil {
		return err
	}

	// init package pkg/service/capsul (custom capsules)
	if err := initCustomCapsulClients(c.CustomCapsul); err != nil {
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

	str := db.ConnStr{
		Host:         c.Creds.Host,
		Database:     database,
		User:         c.Creds.User,
		Password:     c.Creds.Password,
		Port:         c.Creds.Port,
		SslMode:      0,
		MaxIdleConns: c.MIC,
		MaxOpenConns: c.MOC,
	}
	// Create store engine
	engine, err := store.New(str)
	if err != nil {
		return err
	}

	store.Agent = engine
	return nil
}

func initSecrets(c *config.Secrets) error {
	for _, s := range c.Signup {
		secret.Register("signup", s)
		log.Printf("Registred signup secret: %s\n", s)
	}

	return nil
}

func initCapsulClients(c *config.Capsul) error {
	// init mnist client
	mnistClient, err := tfsclient.New(c.Mnist.Address)
	if err != nil {
		return err
	}
	capsul.MnistClient = mnistClient

	// init inception client
	inceptionClient, err := tfsclient.New(c.Inception.Address)
	if err != nil {
		return err
	}
	capsul.InceptionClient = inceptionClient

	// init mura client
	muraClient, err := tfsclient.New(c.Mura.Address)
	if err != nil {
		return err
	}
	capsul.MuraClient = muraClient

	// init chexray client
	chexrayClient, err := tfsclient.New(c.Chexray.Address)
	if err != nil {
		return err
	}
	capsul.ChexrayClient = chexrayClient

	return nil
}

func initCustomCapsulClients(c *config.CustomCapsul) error {
	// init mura-cam
	muraCustomCamClient, err := tfsclient.NewRest(c.MuraCam.Address, 5)
	if err != nil {
		return err
	}
	capsul.MuraCamClient = muraCustomCamClient

	return nil
}
