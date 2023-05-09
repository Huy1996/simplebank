# HTTPS API

### Popular web frameworks

- Gin
- Beego
- Echo
- Revel
- Martini
- Fiber
- Buffalo

### Popular HTTP router

- FastHttp
- Gorilla Mux
- HttpRouter
- Chi

### Viper Package

- Find, load, unmarshal config file
  - JSON, TOML, YAML, ENV, INI
- Read config from environment variables or flags
  - Override existing values, set default values
- Read config from remote system
  - Etcd, Consul
- Live watching and writting config file
  - Reread changed file, save any modifications

### Mocking DB for testing

- Independent tests: Isolate tests data to avoid conflicts
- Faster tests: Reduce a lot of time talking to the database
- 100% coverage: Easily setup edge cases: unexpected errors

### How to mock DB

- Use fake DB: Memory - implement a fake version of DB store data in memory
- Use DB stubs: Gomock - Generate and build stubs that return hard-coded value
  - `go get github.com/golang/mock/mockgen@v1.6.0`