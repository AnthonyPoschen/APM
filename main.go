package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func configureSync(app *kingpin.Application, f *flagList) {
	pacSync := app.Flag("sync", "sync package(s)").Short('S').Action(f.SyncCommand.syncPacman)
	pacSync.BoolVar(&f.SyncCommand.Sync)
	aurSync := app.Flag("syncaur", "sync aur package(s)").Short('y').Action(f.SyncCommand.syncAUR)
	aurSync.BoolVar(&f.SyncCommand.SyncAur)
	//app.Flag("upgrade", "upgrade the full system").Short('u').BoolVar(&f.SyncCommand.Upgrade)
}

func configureOptions(app *kingpin.Application, f *flagList) {
	app.Flag("dbpath", "set an alternate database location").Short('b').StringVar(&f.Options.dbpath)
	app.Flag("clean", "remove old packages from cache directory (-cc for all)").Short('c').CounterVar(&f.Options.clean)
	app.Flag("nodeps", "skip dependency version checks (-dd to skip all checks)").Short('d').CounterVar(&f.Options.nodeps)
	app.Flag("groups", "view all member of a package group\n(-gg to view all groups and members)").Short('g').CounterVar(&f.Options.groups)
	app.Flag("info", "view package information (-ii for extended information)").Short('i').CounterVar(&f.Options.info)
	app.Flag("list", "view a list of packages in a repo").Short('l').StringVar(&f.Options.list)
	app.Flag("print", "print the targets instead of performing the operation").Short('p').BoolVar(&f.Options.print)
	app.Flag("quiet", "show less informatuon for query and search").Short('q').BoolVar(&f.Options.quiet)
	app.Flag("root", "set an alternate installation root").Short('r').StringVar(&f.Options.root)
	app.Flag("search", "search remote repositories for matching strings").Short('s').BoolVar(&f.Options.search)
	app.Flag("sysupgrade", "upgrade installed packages (-uu enables downgrades)").Short('u').CounterVar(&f.Options.sysupgrade)
	app.Flag("verbose", "be verbose").Short('v').BoolVar(&f.Options.verbose)
	app.Flag("downloadonly", "download packages but do not install/upgrade anything").Short('w').BoolVar(&f.Options.downloadOnly)
	//app.Flag("refresh", "donload fresh package databases from the server\n(-yy to force a refresh even if up to date)").Short('y').CounterVar(&f.Options.refresh)
}

func main() {
	app := kingpin.New("apm", "Arch/AUR package manager")
	flags := &flagList{}
	app.HelpFlag.Short('h')
	configureSync(app, flags)
	configureOptions(app, flags)

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Printf("Command: %s\n", cmd)

}
