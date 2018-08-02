# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

(See dev branch)

## 0.1.0 - 2018-08-01

Release 0.1.0 base on 0.0.4

## 0.0.4 - 2018-07-30

### Added
General:
- Bump v0.0.4
- add ldflag `cicd` to controll compiling gRPC Clients (Capsul dependencies used to cause problems while compiling on CircleCI/CD)

Command:
- Added debug mode for UI dev (called with flag `debug`) 
- cmd flags `start` and `syndb` now support `--create-users` subflag ( This will precreate users in database before starting the service)  

Application:
- Add mnist - inception - pr - chexray - mura demos
- Add their backed capsules
- Add package `service` (central buisness logic)
- Add filesystem package `fs` 
- Add gRPC capsul clients (see project capsul)
- Support upload/drop image to inception demo

Front:
- Add home menu
- Add demos menu
- Add small datahub (for inception demo only)
- Add mnist - Inception - Polynomial regression - chexray - mura Demos UI
- tmpl engine automatically render version

Deployment:
- Add swarm support

### Changed
General:
- Align badges
- Add home menu
- Fix Broken CI/CD tests

Command:
- Fix gen views bug (generated views have now one global data see `pkg/tmpl/data.Gen()`)

Front
- Footer 'Made with <3 by SII RESEARCH'

Deployment:
- Fix kubernetes deploy config
- Updated makefiles and dockerfiles
- Move docker-compose.yml to `./deploy`

### Removed 

Command:
- Capsules flag (will be medtune/capsul project responsability)

## 0.0.3 - 2018-07-16

### Added

General:
- Add changelog file
- Add license file
- Add Maintainers list file
- Add circle CI/CD tests config

Commands:
- Add subcommands: `run-server` / `gen-views` / `debug` / `syncdb` / `version`
- Subcommand `run-server` now waits database to get set before serving
- Add `index/home/login/signup` responsive templates

Application:
- Add `/pkg/initpkg` for initalizing packages
- Add user DAO model
- Add agent.user auth/sign methods
- Add json utilities API responses/messages
- Add custom session methods
- Split handlers to public/hidden/debug/api
- Add session logout
- Handlers have a no route method

Front:
- Views use jquery.js
- Views types.js
- Support forms post with js
- Add social media links in mega footer 
- Add error page 

Deployement
- Docker image build (Dockerfile)
- Docker swarm build script (docker-compose.yml)
- Add makefiles
- Add deploy kubernetes config

### Changed
- Improved medtune-beta command line 
- Config from env to yaml format (Yet Another Markup Language)
- Session moved to `/pkg/session`
- Session keys are randomly generated.
- Database abstraction went to `/pkg/agent`
- Footer to megafooter

### Security
- Start using SHA256 to hash passwords
- Handlers checks session before rendering views
- Handlers redirect to index if user isn't connected


## 0.0.2 - Unreleased - 2018-07-10
### Added
- basic server structure
- configuration from env
- main.go simple cmd entrypoint
- simple database connection
- template engine
- static directory
- Init middlewares
- Init session
- Add delphine.yml
- Public handlers
- Basic documentation
- Gitignore

## 0.0.1 - Unreleased - 2018-07-05
### Init
- Init project

[Unreleased]: https://github.com/medtune/beta-platform/compare/v0.0.4...HEAD