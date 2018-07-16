# medtune-beta command line

- [1. Root](#root-command)
- [2. Version](#version)
- [3. Automigrate](#automigrate)
- [4. Generate views](#generate-views)
- [5. Debug server](#debug-server)
- [6. Run server](#run-server)

#### Root command
```
Usage:
  medtune-beta [command]

Available Commands:
  automigrate Auto migrate database
  capsules    Not implemented
  debug       debug server for UI dev
  gen-views   Generate views html files
  help        Help about any command
  start       Run Medtune beta server
  version     Medtune beta actual version

Flags:
  -h, --help   help for medtune-beta
```

#### Subcommands

###### Version
```
Print Medtune Beta version

Usage:
  medtune-beta version [flags]

Flags:
  -h, --help   help for version
```

###### Auto-migrate
```
Sync database models by updating/creating existing
database tables

Usage:
  medtune-beta automigrate [flags]

Aliases:
  automigrate, syncdb

Flags:
  -f, --file string   Configuration file name (default "config.yml")
  -h, --help          help for automigrate
```

###### Generate views
```
Generate views html files

Usage:
  medtune-beta gen-views [flags]

Aliases:
  gen-views, gen-tmpl, gen

Flags:
  -h, --help            help for gen-views
  -o, --output string   output directory (default generate-views) (default "generated-views")
  -v, --views string    views to generate (comma separated string) (default "...")
``` 

###### Debug server
```
Debug UI server for dev purposes

Usage:
  medtune-beta debug [flags]

Aliases:
  debug, debug-server

Flags:
  -h, --help            help for debug
  -p, --port int        port (default 8005)
  -s, --static string   Static files directory (default "./static")
```

###### Run server
```
Run Medtune beta server

Usage:
  medtune-beta start [flags]

Aliases:
  start, run, run-server

Flags:
  -f, --file string          Configuration file name (default "./config.yml")
  -g, --gin-mode int         Gin server mode [0 OR 1]
  -h, --help                 help for start
  -p, --port int             port (default 8005)
  -s, --static string        Static files directory (default "./static")
  -x, --syncdb               Sync database before start (default true)
  -w, --wait                 Wait all services to go up
  -c, --wait-attempts int    Wait max attempts (default 30)
  -t, --wait-timestamp int   Wait timestamp (default 1)
```
