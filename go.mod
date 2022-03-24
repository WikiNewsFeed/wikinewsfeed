module github.com/wikinewsfeed/wikinewsfeed

replace github.com/wikinewsfeed/wikinewsfeed/parser => ./parser

replace github.com/wikinewsfeed/wikinewsfeed/client => ./client

replace github.com/wikinewsfeed/wikinewsfeed/metrics => ./metrics

go 1.16

require (
	github.com/gobuffalo/envy v1.10.1
	github.com/gorilla/feeds v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/kr/pretty v0.3.0 // indirect
	github.com/prometheus/client_golang v1.12.1
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/wikinewsfeed/wikinewsfeed/client v0.0.0-00010101000000-000000000000
	github.com/wikinewsfeed/wikinewsfeed/parser v0.0.0-20220317213711-4b1d1315140c
	go.etcd.io/bbolt v1.3.6
)
