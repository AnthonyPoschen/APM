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
	app.Flag("arch", "set an alternate architecture").StringVar(&f.Options.arch)
	app.Flag("asdeps", "install packages as non-explicitly installed").BoolVar(&f.Options.asdeps)
	app.Flag("asexplicit", "install packages as explicitly installed").BoolVar(&f.Options.asexplicit)
	app.Flag("assume-installed", "<package=version>\nadd a virtual package to satisfy dependencies").StringVar(&f.Options.assumeInstalled)
	app.Flag("cachedir", "set an alternate package cache location").StringVar(&f.Options.cachedir)
	app.Flag("color", "colorise the output").StringVar(&f.Options.color)
	app.Flag("config", "set an alternate configuration file").StringVar(&f.Options.config)
	app.Flag("confirm", "always ask for confirmation").BoolVar(&f.Options.confirm)
	app.Flag("dbonly", "only modify database entries, not package files").BoolVar(&f.Options.dbonly)
	app.Flag("debug", "display debug messages").BoolVar(&f.Options.debug)
	app.Flag("force", "force install, overwrites conflicting files").BoolVar(&f.Options.force)
	app.Flag("gpgdir", "set an alternate home directory for GnuPG").StringVar(&f.Options.gpgdir)
	app.Flag("hookdir", "set an alternate hook location").StringVar(&f.Options.hookdir)
	app.Flag("ignore", "ignore a package upgrade (can be used more than once").StringsVar(&f.Options.ignoreList)
	app.Flag("ignoregroup", "ignore a group upgrade ( can be used more than once)").StringsVar(&f.Options.ignoreGroupList)
	app.Flag("logfile", "set an alternate log file").StringVar(&f.Options.logfile)
	app.Flag("needed", "do not reinstall up to date packages").BoolVar(&f.Options.needed)
	app.Flag("noconfirm", "do not ask for any confirmation").BoolVar(&f.Options.noconfirm)
	app.Flag("noprogressbar", "do not show a progress bar when downloading files").BoolVar(&f.Options.noProgressBar)
	app.Flag("noscriptlet", "do not execute the install scriptlet if one exists").BoolVar(&f.Options.noScriptLet)
	app.Flag("print-format", "specify how the targets should be printed")

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
