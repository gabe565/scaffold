package defaults

import "github.com/clevyr/scaffold/modulemap"

var ComposerDeps = modulemap.ModuleMap{
	"backpack/crud": {Hidden: true},
	"clevyr/backpack-page-builder": {},
	"laravel/nova": {Hidden: true},
	"nunomaduro/larastan": {Enabled: true},
	"superbalist/laravel-google-cloud-storage": {Enabled: true},
}
