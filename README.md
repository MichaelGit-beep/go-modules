[VIDEO](https://www.youtube.com/watch?v=Z1VhG7cf83M&ab_channel=SteveHook)
- Add the model to the project
`go get github.com/julienschmidt/httprouter`

- Add the model to the project by commit hash
`github.com/julienschmidt/httprouter@lkajsdfoip8asdfasdkuf7678lasdlfk `

- Add the model to the project by git tag
`github.com/julienschmidt/httprouter@v1.3`

- Get rid of unused modules in the project
`go mod tidy`

- Show who is using the module
`go mod why github.com/julienschmidt/httprouter`

- Semantic import versioning. If you trying to download go mod above v1. You will need to use a specific syntax to get and import
```
V0 - UNSTABLE
V1 - Stable minor
V2 - Stable major


go get rsc.io/quote@v1.5.2 - Will work since v1 is always stable
go get rsc.io/quote@v3.1.0 - Will not work
go get rsc.io/quote/v3 - Will work and will download latest v3 version

import (
    quoteV3 "rsc.io/quote/v3"
)

```

- Download all the dependencies of the project that is used in your code
```
go get ./...

# Or

go build
go run 
go test
```

- Clean all Gocache and module cache
```
go clean -cache -modcache -i -r
```

- vendore directory generating that compatible with all versions of go
```
go mod vendor
```

- check all available version of the module
```
go list -m -versions go.uber.org/zap
```

- fetching all dependencies of you go progect specified in go.mod
```
go mod download
```

- explore golang variables
```
go env
```


- when you downloading a go package from some source it will cached in outside repo to encrease redundancy and availability of your dependency, package checksum is also cached in external DB to compare if package you download is valid and wasn't compromised
```
go env | grep GOPROXY
GOPROXY=https://proxy.golang.org,direct 

go env | grep GOSUMDB
GOSUMDB=sum.golang.org

```

