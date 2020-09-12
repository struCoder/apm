<div align="center">
<a>
   <img src="https://i.loli.net/2018/12/06/5c08b9a294c29.png">
</a>
<br/>
<b>PMGO</b>
<br/><br/>
<a href="https://circleci.com/gh/struCoder/pmgo">
<img src="https://circleci.com/gh/struCoder/pmgo.svg?&style=shield&circle-token=0fa8ccfc85928edc54a0d7d848cbc784e31813ff" alt="Build Status">
</a>

<a href="http://commitizen.github.io/cz-cli">
  <img src="https://img.shields.io/badge/commitizen-friendly-brightgreen.svg" alt="Commitizen friendly" />
</a>

<a href="https://gitter.im/getpmgo/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge">
  <img src="https://badges.gitter.im/getpmgo/Lobby.svg" alt="Join the chat" />
</a>

<a href="https://goreportcard.com/report/github.com/struCoder/pmgo">
  <img src="https://goreportcard.com/badge/github.com/struCoder/pmgo" alt="Go Report Card" />
</a>

<a href="https://godoc.org/github.com/struCoder/pmgo">
  <img src="https://godoc.org/github.com/struCoder/pmgo?status.svg" alt="GoDoc" />
</a>
<br/><br/>
</div>


# PMGO 
PMGO is a lightweight process manager written in Golang for Golang applications. It helps you keep your applications alive forever, reload and start them from the source code.



## Change log(2020-06-27 updated)

[Change log](./changelog.md)


## Install pmgo

```bash
$ go get github.com/struCoder/pmgo
$ mv $GOPATH/bin/pmgo /usr/local/bin
```

Or
```bash
git clone https://github.com/struCoder/pmgo.git
cd path/to/struCoder/pmgo
go build -v pmgo.go
mv pmgo /usr/local/bin
```


## Starting a new application
If it's the first time you are starting a new golang application, you need to tell pmgo to first build its binary. Then you need to first run:
```bash
$ pmgo start path/to/source-directory app-name
```

This will automatically compile, start and daemonize your application. If you need to later on, stop, restart or delete your app from PMGO, you can just run normal commands using the app-name you specified. Example:
```bash
$ pmgo stop app-name
$ pmgo restart app-name
$ pmgo delete app-name
```

## Main features

### Commands overview

```bash
$ pmgo kill                                                  # kill pmgo daemon process

$ pmgo start source app-name                                 # Compile, start, daemonize and auto  restart application.
$ pmgo restart app-name                                      # Restart a previously saved process
$ pmgo stop app-name                                         # Stop application.
$ pmgo delete app-name                                       # Delete application forever.

$ pmgo save                                                  # Save current process list

$ pmgo list                                                  # Display status for each app.
$ pmgo info app-name                                         # describe importance parameters of a process name
```

#### Start your GO-application with parameters
```bash
pmgo start tmp/ test --args "arg1 arg2 arg3"

# In your application
fmt.Println(os.Args[1:])
# Output: [arg1, arg2, arg3]
```


### Start application from user input compiled binary

```bash
# true means use user input compiled binary path
pmgo start /Users/strucoder/personalPro/goplace/main awesome_name true --args="arg1 arg2 arg3"
```

### Demo
![demo](https://i.loli.net/2018/12/06/5c08bbd407b35.png)

### I Love This. How do I Help?

- Simply star this repository :-)
- Help us spread the world on Facebook and Twitter
- Contribute Code!
- I'll be very grateful if you'd like to donate to encourage me to continue maintaining pmgo.


### Contributors
<a href="https://github.com/struCoder/pmgo/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=struCoder/pmgo" />
</a>

### Donate

|      **Paypal**       |        **Alipay**         |
| :------------------------: | :------------------------: |
| [![paypal](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.me/strucoder) | [![alipay](https://img.shields.io/badge/Donate-alipay-blue.svg)](https://i.loli.net/2018/11/29/5bff95e2d29df.png) |

### By The Way
In China Mainland, maybe you can't download some packages in golang.org, thus just click [here](https://goproxy.io/zh/) to set `GOPROXY`
### LICENSE

[MIT](https://github.com/struCoder/pmgo/blob/master/LICENSE)
