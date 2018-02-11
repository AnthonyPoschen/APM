package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type flagList struct {
	SyncCommand syncCommand
	Options     optionFlags
	Version     *bool
}

type optionFlags struct {
	dbpath          string
	clean           int
	nodeps          int
	groups          int
	info            int
	list            string
	print           bool
	quiet           bool
	root            string
	search          bool
	sysupgrade      int
	verbose         bool
	downloadOnly    bool
	refresh         int
	arch            string
	asdeps          bool
	asexplicit      bool
	assumeInstalled string
	cachedir        string
	color           string
	config          string
	confirm         bool
	dbonly          bool
	debug           bool
	force           bool
	gpgdir          string
	hookdir         string
	ignoreList      []string
	ignoreGroupList []string
	logfile         string
	needed          bool
	noconfirm       bool
	noProgressBar   bool
	noScriptLet     bool
}

type syncCommand struct {
	Sync    bool
	SyncAur bool

	// options
	Upgrade bool
}

func (s *syncCommand) syncPacman(c *kingpin.ParseContext) error {
	fmt.Printf("Sync: %t\nUpgrade: %t\n", s.Sync, s.Upgrade)
	return nil
}

func (s *syncCommand) syncAUR(c *kingpin.ParseContext) error {
	fmt.Printf("Aur Sync: %t\nUpgrade: %t\n", s.SyncAur, s.Upgrade)
	return nil
}
