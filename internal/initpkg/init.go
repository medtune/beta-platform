package initpkg

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/sessions/redis"
	"github.com/go-xorm/core"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
	"github.com/medtune/go-utils/random"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/config"
	"github.com/medtune/beta-platform/internal/service/capsul"
	"github.com/medtune/beta-platform/internal/service/dashboard"
	"github.com/medtune/beta-platform/internal/service/secret"
	"github.com/medtune/beta-platform/internal/session"
	"github.com/medtune/beta-platform/internal/store"
	"github.com/medtune/beta-platform/internal/store/db"
)

// InitFromFile from configuration file
func InitFromFile(file string) error {
	// Read configuration from file
	config, err := config.LoadConfigFromPath(file)
	if err != nil {
		return fmt.Errorf("load configuration: %v", err)
	}

	// Init from config
	if err = InitFromConfig(config); err != nil {
		return fmt.Errorf("initialize package: %v", err)
	}

	return nil
}

// InitFromConfig uration file
func InitFromConfig(c *config.StartupConfig) error {
	// Validate config
	if _, err := c.Validate(); err != nil {
		return fmt.Errorf("non valid configuration: %v", err)
	}

	// Check version
	if c.Meta.Version != internal.VERSION {
		return fmt.Errorf("configuration meta version don't match\n\tpackage version: %v\n\tconfigs version: %v", internal.VERSION, c.Meta.Version)
	}

	// init package internal/store + internal/store/db
	if err := initStore(c.Database, c.Meta.IsProd); err != nil {
		return fmt.Errorf("init package store failed: %v", err)
	}

	// init package internal/session
	if err := initSession(c.Session); err != nil {
		return fmt.Errorf("init package session failed: %v", err)
	}

	// init package internal/secret
	if err := initSecrets(c.Secrets); err != nil {
		return fmt.Errorf("init service secrets failed: %v", err)
	}

	// init package internal/service/capsul
	if err := initCapsulClients(c.Capsul); err != nil {
		return fmt.Errorf("init service capsul clients failed: %v", err)
	}

	// init package internal/service/capsul (custom capsules)
	if err := initCustomCapsulClients(c.CustomCapsul); err != nil {
		return fmt.Errorf("init service capsul custom clients failed: %v", err)
	}

	// inint internal/service/dashboard
	if err := initDashboard(c); err != nil {
		return fmt.Errorf("init service dashboard failed: %v", err)
	}

	return nil
}

func initSession(c *config.Session) error {
	if c.Type == "cookie" {
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

	} else if c.Type == "redis" {
		store, err := redis.NewStore(10, "tcp", c.Address, c.Password, []byte(c.Secret))
		if err != nil {
			return err
		}
		session.Store = store
	}
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

	if c.LTC == 0.0 {
		fmt.Println("using 0.0 (ms) as connection max lifetime")
		c.LTC = 0.001
	}

	str := db.ConnStr{
		Host:             c.Creds.Host,
		Database:         database,
		User:             c.Creds.User,
		Password:         c.Creds.Password,
		Port:             c.Creds.Port,
		SslMode:          0,
		MaxIdleConns:     c.MIC,
		MaxOpenConns:     c.MOC,
		MaxLifetimeConns: c.LTC,
	}

	// Create store engine
	engine, err := store.New(str)
	if err != nil {
		return err
	}

	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	//engine.SetLogger(xorm.NewSimpleLogger(os.Stdout))

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
	muraMNV2Client, err := tfsclient.New(c.MuraMNV2.Address)
	if err != nil {
		return err
	}
	capsul.MuraMNV2Client = muraMNV2Client

	// init mura client

	muraIRNV2Client, err := tfsclient.New(c.MuraIRNV2.Address)
	if err != nil {
		return err
	}
	capsul.MuraIRNV2Client = muraIRNV2Client

	// init chexray client
	chexrayMNV2Client, err := tfsclient.New(c.ChexrayMNV2.Address)
	if err != nil {
		return err
	}
	capsul.ChexrayMNV2Client = chexrayMNV2Client

	// init chexray dn121 client
	chexrayDN121Client, err := tfsclient.New(c.ChexrayDN121.Address)
	if err != nil {
		return err
	}
	capsul.ChexrayDN121Client = chexrayDN121Client

	return nil
}

func initCustomCapsulClients(c *config.CustomCapsul) error {
	// init mura cam
	muraMNV2CamClient, err := tfsclient.NewRest(c.MuraMNV2Cam.Address, 5)
	if err != nil {
		return err
	}
	capsul.MuraMNV2CamClient = muraMNV2CamClient

	// init chexray cam
	chexrayMNV2CamClient, err := tfsclient.NewRest(c.ChexrayMNV2Cam.Address, 5)
	if err != nil {
		return err
	}
	capsul.ChexrayMNV2CamClient = chexrayMNV2CamClient

	// init chexray preprocessing helper
	chexrayPPHelper, err := tfsclient.NewRest(c.ChexrayPP.Address, 5)
	if err != nil {
		return err
	}
	capsul.ChexrayPPHelper = chexrayPPHelper

	return nil
}

func initDashboard(c *config.StartupConfig) error {
	dashboard.StartupConfig = c
	dashboard.StartupDate = time.Now()
	return nil
}
