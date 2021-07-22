package defaults

import "github.com/clevyr/scaffold/module"

var ComposerDeps = module.ModuleSlice{
	{
		Name:    "laravel/jetstream",
		Enabled: true,
		PostInstallCmds: [][]string{
			{"php", "artisan", "jetstream:install", "inertia", "--teams"},
		},
	}, {
		Name:    "laravel/telescope",
		Enabled: true,
		PostInstallCmds: [][]string{
			{
				"cp",
				"vendor/laravel/telescope/stubs/TelescopeServiceProvider.stub",
				"app/Providers/TelescopeServiceProvider.php",
			},
			{"php", "artisan", "telescope:install"},
			{"mkdir", "-p", "public/vendor/telescope"},
			{"sh", "-c", "cp vendor/laravel/telescope/public/* public/vendor/telescope"},
		},
	}, {
		Name:    "joelbutcher/socialstream",
		Enabled: true,
		PostInstallCmds: [][]string{
			{"php", "artisan", "socialstream:install"},
		},
	}, {
		Name: "laravel/nova",
		PostInstallCmds: [][]string{
			{"mkdir", "-p", "nova-components"},
			{"php", "artisan", "nova:install"},
		},
	}, {
		Name: "backpack/crud",
	}, {
		Name: "clevyr/backpack-page-builder",
	}, {
		Name: "superbalist/laravel-google-cloud-storage",
	}, {
		Name:    "nunomaduro/larastan",
		Dev:     true,
		Enabled: true,
		Hidden:  true,
	},
}
