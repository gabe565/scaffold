package defaults

import "github.com/clevyr/scaffold/module"

var ComposerDeps = module.ModuleSlice{
	{
		Name:    "laravel/jetstream",
		Enabled: true,
		Version: "2.6.6",
		PostInstallCmds: [][]string{
			{"php", "artisan", "jetstream:install", "inertia"},
		},
	}, {
		Name:    "laravel/telescope",
		Enabled: true,
		Version: "4.7.3",
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
		Version: "3.4.0",
		PostInstallCmds: [][]string{
			{"php", "artisan", "socialstream:install"},
		},
	}, {
		Name:    "laravel/nova",
		Version: "3.31.0",
		PostInstallCmds: [][]string{
			{"mkdir", "-p", "nova-components"},
			{"php", "artisan", "nova:install"},
		},
	}, {
		Name:    "laravel/spark-paddle",
		Enabled: true,
		Version: "1.2.2",
		PostInstallCmds: [][]string{
			{"php", "artisan", "spark:install"},
		},
	}, {
		Name:    "clevyr/nova-page-builder",
		Enabled: false,
		PostInstallCmds: [][]string{
			{"php", "artisan", "vendor:publish", "--tag=clevyr-nova-page-builder"},
		},
	}, {
		Name:    "ukfast/laravel-health-check",
		Enabled: false,
		Version: "1.13.0",
	}, {
		Name: "superbalist/laravel-google-cloud-storage",
	}, {
		Name:    "nunomaduro/larastan",
		Dev:     true,
		Enabled: true,
		Hidden:  true,
	},
}
