package defaults

import "github.com/clevyr/scaffold/module"

var ComposerDeps = module.ModuleSlice{
	{
		Name:    "laravel/jetstream",
		Enabled: true,
		Version: "2.3.11",
		PostInstallCmds: [][]string{
			{"php", "artisan", "jetstream:install", "inertia", "--teams"},
		},
	}, {
		Name:    "laravel/telescope",
		Enabled: true,
		Version: "4.6.0",
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
		Version: "3.1.3",
		PostInstallCmds: [][]string{
			{"php", "artisan", "socialstream:install"},
		},
	}, {
		Name:    "laravel/nova",
		Version: "3.27.0",
		PostInstallCmds: [][]string{
			{"mkdir", "-p", "nova-components"},
			{"php", "artisan", "nova:install"},
		},
	}, {
		Name:    "laravel/spark-paddle",
		Enabled: true,
		Version: "1.1.5",
		PostInstallCmds: [][]string{
			{"php", "artisan", "spark:install"},
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
