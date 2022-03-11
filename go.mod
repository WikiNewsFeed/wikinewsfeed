module github.com/wikinewsfeed/wikinewsfeed

replace github.com/wikinewsfeed/wikinewsfeed/parser => ./parser

go 1.16

require (
	github.com/gorilla/feeds v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/kr/pretty v0.3.0 // indirect
	github.com/wikinewsfeed/wikinewsfeed/parser v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
)
