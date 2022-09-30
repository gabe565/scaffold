package appconfig

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/clevyr/scaffold/config/modules"
	"github.com/clevyr/scaffold/internal/iexec"
	"github.com/clevyr/scaffold/internal/module"
)

type AppConfig struct {
	InitLaravel    bool `json:"-"`
	AppName        string
	AppSlug        string `json:"-"`
	AppKey         string
	Database       string
	PhpModules     module.Map
	JetstreamTeams bool
	GitCommit      bool
	ComposerDeps   module.ComposerMap
	NpmDeps        module.NpmMap
	MaxUploadSize  string
}

var Defaults = AppConfig{
	Database:       "PostgreSQL",
	PhpModules:     modules.Php(),
	JetstreamTeams: false,
	GitCommit: true,
	ComposerDeps:   modules.Composer(),
	NpmDeps:        modules.Npm(),
	MaxUploadSize:  "64m",
}

const AppKeyPrefix = "base64:"
const AppKeyBytes = 32

func (appConfig *AppConfig) GenerateAppKey() (err error) {
	if appConfig.AppKey != "" {
		return
	}
	randomBytes := make([]byte, AppKeyBytes)

	if _, err = rand.Read(randomBytes); err != nil {
		return
	}

	appConfig.AppKey = fmt.Sprintf("%s%s", AppKeyPrefix, base64.StdEncoding.EncodeToString(randomBytes))
	return
}

func (appConfig *AppConfig) EnableSelectedDatabase() {
	var name string
	switch appConfig.Database {
	case "PostgreSQL":
		name = "pgsql"
	case "MariaDB":
		name = "mysql"
	default:
		panic(fmt.Errorf("unknown database: %s", name))
	}
	if m, ok := appConfig.PhpModules[name]; ok {
		m.Enabled = true
	}
}

func (appConfig *AppConfig) EnableJetstreamTeams() {
	if appConfig.JetstreamTeams {
		// Append Jetstream's post install command with the '--teams' modifier
		jetstream, ok := appConfig.ComposerDeps.Map["laravel/jetstream"]
		if !ok {
			panic(fmt.Errorf("%v: %s", module.ErrInvalidModule, "laravel/jetstream"))
		}
		for _, action := range jetstream.Then {
			if action.Run != nil {
				run := *action.Run
				if len(run) > 3 && run[2] == "jetstream:install" {
					*action.Run = append(*action.Run, "--teams")
					break
				}
			}
		}
	}
}

func (appConfig AppConfig) SetPackageName() error {
	return iexec.NewBuilder("npm", "pkg", "set", "name="+appConfig.AppSlug).Run()
}

func (appConfig AppConfig) CreateGitCommit() error {
	if err := iexec.NewBuilder("git", "init", "--initial-branch=main").Run(); err != nil {
		return err
	}

	if err := iexec.NewBuilder("git", "add", "--all").Run(); err != nil {
		return err
	}

	return iexec.NewBuilder("git", "commit", "-m", "Initial Commit").Run()
}
