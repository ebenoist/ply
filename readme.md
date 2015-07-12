PLY
---
> ply is a dependency-less docker deploy tool.

Installation
---

### OSX
`brew tap somecask/ebenoist/ply`
`brew install ply`

Usage
---
`ply deploy -e staging -c ply.yml -var "Env=staging,DockerTag=your-docker-tag"`

`ply remote -e staging -c ply.yml bin/rake db:migrate`

`ply remote -e staging -c ply.yml script/console`

Config
---

Configuration is done with a [templated](http://golang.org/pkg/text/template/) [yml](http://yaml.org) file.
Any value can be overridden with either the -var flag or the vars hash under each DeployEnv.

```yaml
Version: 1

AppName: my-new-app
DeployUser: deploy
Container: ebenoist/my-app
BeforeRun:
  - docker exec {{.AppName}} rm /tmp/heartbeat.txt
RunCommand: >
  docker run -d
    -e "ENV={{.Env}}
    -e "JRUBY_OPTS={{.JRubyOpts}}"
    -v /var/{{.AppName}/log:/var/log/nginx
    -p 80:80
    --name {{.AppName}} {{.Container}}:{{.DockerTag}}
AfterRun:
    - docker exec {{.AppName}} touch /tmp/heartbeat.txt
DeployEnv:
  Staging:
    Hosts:
      - my-app1.com
      - my-app2.com
    Vars:
      JRubyOpts: -J-Xmn5376m -J-Xms7168m -J-Xmx7168m
Plugins:
  Hipchat:
    Token: lasdkfj-lasdkfj
    Rooms:
      - 20348
      - 53458

```

Plugins
---
Plugins will likely be separate binaries that talk via RPC?

`ply plugin path/to/plugin` -- pull the binary? build it?

Or should they just be built in?
