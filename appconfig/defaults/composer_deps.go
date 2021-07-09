package defaults

import "github.com/clevyr/scaffold/modulemap"

var ComposerDeps = modulemap.ModuleMap{
	"backpack/crud": {},
	"clevyr/backpack-page-builder": {},
	"laravel/nova": {},
	"nunomaduro/larastan": {Dev: true, Enabled: true, Hidden: true},
	"superbalist/laravel-google-cloud-storage": {},
}
