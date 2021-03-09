package main

import (
	"github.com/clevyr/scaffold/iexec"
)

func npmInstall() (err error) {
	err = iexec.Command("npm", "install")
	return
}
