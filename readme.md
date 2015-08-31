PLY
---
> ply is a dependency-less remote runner

[![Build Status](https://travis-ci.org/ebenoist/ply.svg)](https://travis-ci.org/ebenoist/ply)

Installation
---

### OSX
- `brew tap ebenoist/ply`
- `brew install ply`

Usage
---
`ply -e staging -c ply.yml -vars "Env=staging,DockerTag=your-docker-tag" deploy`

Config
---

Configuration is done with a [templated](http://golang.org/pkg/text/template/) [yml](http://yaml.org) file.
Any value can be overridden with either the -vars flag, the vars key under each DeployEnv, or at the root level.

```yaml
Version: 1

DeployUser: deploy
Vars:
  AppName: my-new-app
  Container: my-container
Tasks:
  WriteHeartBeat:
    - docker exec {{.AppName}} touch /tmp/heartbeat.txt
  Stop:
    - docker exec {{.AppName}} rm /tmp/heartbeat.txt
    - docker stop
  Run:
    - docker run -d -e "ENV={{.Env}} -e "JRUBY_OPTS={{.JRubyOpts}}" -v /var/{{.AppName}/log:/var/log/nginx -p 80:80 --name {{.AppName}} {{.Container}}:{{.DockerTag}}
  Deploy:
    - Stop
    - Run
    - WriteHeartBeat
DeployEnv:
  Staging:
    Hosts:
      - my-app1.com
      - my-app2.com
    Vars:
      JRubyOpts: -J-Xmn5376m -J-Xms7168m -J-Xmx7168m
```

Plugins
---
Will probably be special predefined tasks

Development
---
Ply is a command line tool developed in [Go](http://golang.org). Ply uses [gb](https://getgb.io/) for dependency management.

### Setup
- [Go 1.5](http://golang.org/doc/install)
- [gb](https://getgb.io) `go get github.com/constabulary/gb/...`
- `make` will run tests
- `make build` will produce a binary in your ./bin folder

Releases
---
Releases are managed by GitHub. Creating a new release will kick off a deployment build which will upload the resulting binaries to the release on a successful build.

After a release, to update the [homebrew](http://brew.sh/) formula open a pull request against [hombrew-ply](https://github.com/ebenoist/homebrew-ply) with the updated version number.

License
---
#### The MIT License (MIT)
Copyright © 2015 Erik Benoist, http://erikbenoist.com

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
