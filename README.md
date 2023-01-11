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

- Semantic import versioning. If you trying to download go mod above v1. You will need to use a specific syntax
```
V0 - UNSTABLE
V1 - Stable minor
V2 - Stable major


go get rsc.io/quote@v1.5.2 - Will work since v1 is always stable
go get rsc.io/quote@v3.1.0 - Will not work
go get rsc.io/quote/v3 - Will work and will download latest v3 version

```
