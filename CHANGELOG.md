# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/medtune/beta-platform/compare/v0.0.3...HEAD
