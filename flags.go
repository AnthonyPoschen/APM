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
	dbpath       string
	clean        int
	nodeps       int
	groups       int
	info         int
	list         string
	print        bool
	quiet        bool
	root         string
	search       bool
	sysupgrade   int
	verbose      bool
	downloadOnly bool
	refresh      int
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
