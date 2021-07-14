# Changelog

## 7.5.0 (2021/04/26)

* Add `disableMLSD` ftp option (#176)
* Fix Dockerfile

## 7.4.0 (2021/04/25)

* Add `darwin/arm64` artifact (#175)
* Bump github.com/go-playground/validator/v10 from 10.4.1 to 10.5.0 (#171)
* Use logger `PartsExclude` (#174)
* MkDocs Materials 7.1.3
* Bump github.com/rs/zerolog from 1.20.0 to 1.21.0 (#166)
* Go 1.16 (#167)
* Deploy docs on workflow dispatch or tag
* Bump github.com/alecthomas/kong from 0.2.15 to 0.2.16 (#165)
* Bump github.com/pkg/sftp from 1.12.0 to 1.13.0 (#164)
* Switch to goreleaser-xx (#163)

## 7.3.0 (2021/02/19)

* Refactor CI and dev workflow with buildx bake (#161)
    * Add `image-local` target
    * Single job for artifacts and image
    * Add `armv5`, `ppc64le` and `s390x` artifacts
    * Upload artifacts
    * Validate
* Remove `linux/s390x` Docker platform support for now
* Bump github.com/stretchr/testify from 1.6.1 to 1.7.0 (#154)
  Bump github.com/alecthomas/kong from 0.2.12 to 0.2.15 (#160)
* MkDocs Materials 6.2.8

## 7.2.0 (2020/11/29)

* Allow to download files to a temp dir first (#149)
* Allow to disable log timestamp (#148)
* Add script notification (#147)
* Bump github.com/crazy-max/gonfig from 0.3.0 to 0.4.0 (#140)

## 7.1.1 (2020/11/02)

* Use embedded tzdata package
* Remove `--timezone` flag
* Docker image also available on [GitHub Container Registry](https://github.com/users/crazy-max/packages/container/package/ftpgrab)
* Use Docker meta action to handle tags and labels

## 7.1.0 (2020/10/04)

* Allow to disable `OPTS UTF8 ON` command
* Refactor to start working on #48
* Switch to Docker actions
* Go 1.15
* Update `GOPROXY` setting
* Update deps

## 7.0.1 (2020/08/04)

* Fix SFTP not taken into account

## 7.0.0 (2020/07/18)

:warning: See **Migration notes** in the documentation for breaking changes.

* Repository moved to [crazy-max/ftpgrab](https://github.com/crazy-max/ftpgrab)
* DockerHub repository moved to [crazymax/ftpgrab](https://hub.docker.com/r/crazymax/ftpgrab)
* Configuration transposed into environment variables (#90)
* `FTPGRAB_DB` env var renamed `FTPGRAB_DB_PATH`
* `key` field for SFTP authentication has been renamed `keyFile`
* Add `keyPassphrase` to provide a passphrase linked to `keyFile`
* Improve configuration validation
* All fields in configuration now _camelCased_
* Add tests and coverage
* Seek configuration file from default places
* Configuration file not required anymore
* Switch to [gonfig](https://github.com/crazy-max/gonfig)
* Add fields to load sensitive values from file
* Update deps

## 6.5.0 (2020/07/07)

* Docs website with mkdocs
* Move documentation to main repository
* Update deps

## 6.4.0 (2020/05/17)

* Use kong command-line parser
* Switch to Open Container Specification labels as label-schema.org ones are deprecated
* Update deps

## 6.3.0 (2020/01/19)

* Only accept duration as timeout value for FTP, SFTP and Webhook notif config (#69)
* Update [pkg/sftp](https://github.com/pkg/sftp) module

## 6.2.0 (2019/12/19)

* Add Slack notifier
* Update deps
* Go 1.13.5
* Seconds field optional for schedule

## 6.1.0 (2019/10/13)

* Multi-platform Docker image
* Move [ftpgrab/docker](https://github.com/ftpgrab/docker) repo here
* Go 1.12.10
* Use GOPROXY
* Stop publishing Docker image on Quay
* Switch to GitHub Actions
* Add instructions to create a Linux service
* Remove `--docker` flag
* Allow to override database path through `FTPGRAB_DB` env var
* Allow to override download output path through `FTPGRAB_DOWNLOAD_OUTPUT` env var

## 6.0.2 (2019/07/27)

* Use `io.Copy` to avoid crash due to insufficient memory

## 6.0.1 (2019/07/24)

* Fix cron stopped after first trigger

## 6.0.0 (2019/07/21)

:warning: See **Migration notes** in the documentation for breaking changes.

* Log skip status
* Set ServerName field if implicit TLS
* Switch to [jlaffaye/ftp](https://github.com/jlaffaye/ftp) module
    * Fix race condition
    * Performance improvement

## 5.5.0 (2019/07/20)

* Switch to [crazy-max/goftp](https://github.com/crazy-max/goftp) (#55)

## 5.4.1 (2019/07/18)

* Fix durafmt runtime error

## 5.4.0 (2019/07/18)

* Improve logging
* Display next execution time
* Use v3 robfig/cron
* Always run on startup
* Go 1.12.4

## 5.3.0 (2019/05/04)

* Escape all regexp metacharacters on read dir (#49)
* Remove unused field
* Go 1.12
* Update deps

## 5.2.0 (2019/03/29)

:warning: See **Migration notes** in the documentation for breaking changes.

* Add webhook notification method
* Remove unnecessary `connections_per_host` field (#48)
* Fix log folder creation

## 5.1.1 (2019/02/18)

* Blackfriday module fixed through hermes v2.0.2 (matcornic/hermes#51)

## 5.1.0 (2019/02/14)

:warning: See **Migration notes** in the documentation for breaking changes.

* Add SFTP support (#42)

## 5.0.1 (2019/02/13)

* Fix high CPU load on schedule
* Add support for FreeBSD

## 5.0.0 (2019/02/12)

:warning: See **Migration notes** in the documentation for breaking changes.

* BIG rewrite (#36)
* Multiplatform : Linux, macOS and Windows on architectures like amd64, 386, ARM and others
* Modern CLI interactions
* Yaml Configuration file
* Detect and merge configuration
* Handle defaults
* Add [Goreleaser](https://goreleaser.com/)
* [Bolt](https://github.com/etcd-io/bbolt) db to audit files already downloaded
* Native FTP client
* Logging with [zerolog](https://github.com/rs/zerolog)
* Send reports through email
* Generate responsive and beautiful email reports through [hermes](https://github.com/matcornic/hermes/)
* Lightweight Docker image (~6MB)
* Docker image moved to a dedicated organization on [Docker Hub](https://hub.docker.com/u/ftpgrab) and [Quay](https://quay.io/organization/ftpgrab).
* [Embedded cron](https://github.com/crazy-max/cron) using go routines
* Manage base dir
* Set original modtime
* Include/exclude based on regexp
* Ignore files by date (#39)
* Handle mutex

## 4.3.5 (2019/02/04)

* Switch to Travis CI (com)

## 4.3.4 (2018/08/15)

* Empty folder leeds to spinlock (#33)

## 4.3.3 (2018/05/14)

* nawk and gawk not required anymore (#38)

## 4.3.2 (2018/04/20)

* Detect if file size is currently changing and hold for download (#37)

## 4.3.1 (2018/01/15)

* Fix issue while checking source hash (#35)

## 4.3.0 (2017/12/26)

* Add an exclude filter for files through `DL_EXCLUDE_REGEX` (#27)

## 4.2.4 (2017/11/01)

* Do not exit if connection failed

## 4.2.3 (2017/10/30)

* Fix files download again (#32)

## 4.2.2 (2017/10/29)

* Rebuild PATH

## 4.2.1 (2017/10/16)

* Add ssmtp on Docker image to send emails
* Use sendmail instead of mail command

## 4.2.0 (2017/10/15)

:warning: See **Migration notes** in the documentation for breaking changes.

* Add Docker image (more info on [docker repository](https://github.com/ftpgrab/docker))
* Remove init script
* Fix issue while resuming downloads
* Move script to `/usr/bin`
* Coding style

## 4.1.1 (2017/04/26)

* Add tests (#30)
* Use type instead of which (#29)
* Fix error prone and performance issues
* Coding style
* Add default config
* Add Codacy

## 4.1 (2017/03/15)

:warning: See **Migration notes** in the documentation for breaking changes.

* Rename the project ftpgrab ! (#28)

## 4.0 (2017/03/14)

:warning: See **Migration notes** in the documentation for breaking changes.

* Shuffle file/folder listing (#25)
* Allow multiple instances (#22)

## 3.2 (2016/06/20)

* Add messages for permission issue (#19)
* Move some instructions to Wiki (#18)
* Update `ISSUE_TEMPLATE.md`
* Add [.editorconfig](http://editorconfig.org/)
* MIT License

## 3.1 (2016/03/27)

**You have to edit the config file `ftp-sync.conf` if you upgrade from a previous release!**

* Add multiple ftp sources paths (#18)
* Sed not escaping & char (#17)
* Add `DL_CREATE_BASEDIR` option to create basename of a ftp source path in the destination folder.

## 3.0 (2016/03/20)

**You have to edit the config file `ftp-sync.conf` if you upgrade from a previous release!**

* MD5 file not created with text mode (#16)
* Implement FTPS support for Curl (#15)
* Implement resume downloads support (#14)
* Add DEBUG option
* Full Curl implementation when selected for file size and list files
* Bug with ftpsyncGetHumanSize function
* Display download regex
* Add sha1 hash type
* Bug with special chars for curl method
* Bug with bash condition
* Add `Found a bug?` section in README.md
* Add `ISSUE_TEMPLATE.md`

## 2.03 (2015/03/22)

* Change location of MD5 file

## 2.02 (2015/03/21)

* Bug checking MD5 (#11)

## 2.01 (2015/03/20)

* Bug download with sqlite3 (#10)

## 2.00 (2015/03/19)

* Add SQLite method to store MD5 hash (#8)

## 1.95 (2014/08/09)

* Bug trailing slash  (#6)

## 1.94 (2014/05/22)

* Bug replacing destination folder

## 1.93 (2014/02/16)

* Update README.md and .gitignore
* New year!
* Adding hide progress option

## 1.92 (2013/12/01)

* Bug with the config file

## 1.91 (2013/12/01)

* Adding curl download method

## 1.9 (2013/10/30)

* Remove progress filter on wget

## 1.8 (2013/10/12)

* Update README.md
* Bug with empty folders

## 1.7 (2013/10/06)

* Adding external config file
* Add gawk as required package
* Update README.md with awk problem
* Change perms recursively when downloads are finished

## 1.6 (2013/07/10)

* Misspelling
* Decoding wget problem
* Alternative to kill old and sub process

## 1.5 (2013/06/10)

* Update README.md
* Add synology example

## 1.4 (2013/06/05)

* Check process already running

## 1.3 (2013/06/02)

* Use wget instead of curlftpfs

## 1.2 (2013/06/01)

* Adding email var to receive logs

## 1.1 (2013/05/31)

* Remove dualEcho
* Improvement of the error log with exec and tail
* Change MD5 filter
* Filter bug and add grep search for hash

## 1.0 (2013/05/24)

* Update README.md
* Initial version
