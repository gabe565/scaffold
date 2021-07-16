package defaults

import "github.com/clevyr/scaffold/modulemap"

var ComposerDeps = modulemap.ModuleMap{
	"laravel/jetstream": {
		Enabled: true,
		PostInstallCmds: [][]string{
			{"php", "artisan", "jetstream:install", "inertia", "--teams"},
		},
	},
	"laravel/telescope": {
		Enabled: true,
		PostInstallCmds: [][]string{
			{
				"cp",
				"vendor/laravel/telescope/stubs/TelescopeServiceProvider.stub",
				"app/Providers/TelescopeServiceProvider.php",
			},
			{"php", "artisan", "telescope:install"},
		},
	},
	"joelbutcher/socialstream": {
		Enabled: true,
		PostInstallCmds: [][]string{
			{"php", "artisan", "socialstream:install"},
		},
	},
	"laravel/nova": {
		PostInstallCmds: [][]string{
			{"mkdir", "-p", "nova-components"},
			{"php", "artisan", "nova:install"},
		},
	},
	"backpack/crud":                            {},
	"clevyr/backpack-page-builder":             {},
	"superbalist/laravel-google-cloud-storage": {},
	"nunomaduro/larastan":                      {Dev: true, Enabled: true, Hidden: true},
}
