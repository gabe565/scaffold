package defaults

import "github.com/clevyr/scaffold/module"

var PhpModules = module.ModuleMap{
	"bcmath":    {},
	"calendar":  {},
	"exif":      {},
	"gd":        {},
	"imagick":   {},
	"mosquitto": {},
	"mysql":     {},
	"opcache":   {Enabled: true},
	"pgsql":     {},
	"redis":     {Enabled: true},
	"sqlsrv":    {},
	"xdebug":    {Hidden: true},
	"zip":       {},
}
