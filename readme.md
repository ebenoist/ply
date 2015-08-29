PLY
---
> ply is a dependency-less remote runner

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
Any value can be overridden with either the -vars flag or the vars hash under each DeployEnv.

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

### Setup
`go get github.com/constabulary/gb/...`
