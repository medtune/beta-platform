# medtune-beta command line

- [1. Root](#root-command)
- [2. Version](#version)
- [3. Automigrate](#automigrate)
- [4. Generate config](#generate-startup-config)
- [5. Generate views](#generate-views)
- [6. Debug server](#debug-server)
- [7. Run server](#run-server)

#### Root command
```shell
Usage:
  medtune-beta [command]

Available Commands:
  syncdb Auto migrate database
  debug       debug server for UI dev
  gen-config  Generate empty startup config file
  gen-views   Generate views html files
  help        Help about any command
  start       Run Medtune beta server
  version     Medtune beta actual version

Flags:
  -h, --help   help for medtune-beta
```

#### Subcommands

###### Version
```shell
Print Medtune Beta version

Usage:
  medtune-beta version [flags]

Flags:
  -h, --help   help for version
```

###### Syncdb
```shell
Sync database models by updating/creating existing
database tables

Usage:
  medtune-beta automigrate [flags]

Aliases:
  automigrate, syncdb
)
Flags:
  -y, --create-users   Create default users before start (default true)
  -f, --file string    Configuration file name (default "config.yml")
  -h, --help           help for automigrate
```

###### Generate startup config

```shell
Generate empty startup config file

Usage:
  medtune-beta gen-config [flags]

Aliases:
  gen-config, gen-cfg

Flags:
  -h, --help            help for gen-config
  -o, --output string   output directory (default "gen.config.yml")
```

###### Generate views
```shell
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
```shell
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
```shell
Run Medtune beta server

Usage:
  medtune-beta start [flags]

Aliases:
  start, run, run-server

Flags:
  -y, --create-users         Create default users before start (default true)
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
