```
██████╗ ██████╗ ███╗   ██╗███████╗    ███████╗███████╗██████╗ ██╗   ██╗███████╗██████╗ 
██╔══██╗██╔══██╗████╗  ██║██╔════╝    ██╔════╝██╔════╝██╔══██╗██║   ██║██╔════╝██╔══██╗
██║  ██║██║  ██║██╔██╗ ██║███████╗    ███████╗█████╗  ██████╔╝██║   ██║█████╗  ██████╔╝
██║  ██║██║  ██║██║╚██╗██║╚════██║    ╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██╔══╝  ██╔══██╗
██████╔╝██████╔╝██║ ╚████║███████║    ███████║███████╗██║  ██║ ╚████╔╝ ███████╗██║  ██║
╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝    ╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝                                                                                      
```

[![GoDoc Badge]][godoc] [![GoReportCard Badge]][goreportcard]

Simple DDNS Server API, similar to the oracle's ddns service. 

The goal of this project to self host the ddns service with similar features to any ddns service but with your chosen dns provider. Currenly supports AWS Route53 for DNS Updates.

Any feedback and pull requests are welcome and highly appreciated. Report any bugs/issues or feature requests in issues. Feel free to start discussion of you need help or have huge feature request.

## Start Application

- Clone this repository
- Run the application to see available commands: `go run main.go`
- Run the application with command _serve_: `go run main.go serve`


[godoc]: https://godoc.org/go.pkg.dipak.io/ddns-server/
[godoc badge]: https://godoc.org/go.pkg.dipak.io/ddns-server?status.svg
[goreportcard]: https://goreportcard.com/report/go.pkg.dipak.io/ddns-server
[goreportcard badge]: https://goreportcard.com/badge/go.pkg.dipak.io/ddns-server
