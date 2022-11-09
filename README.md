# randomselector
Randomly select objects in golang
Current version: v0.0.2
## Install 
```console
go get  github.com/zeroboo/randomselector
```

## Test
```console
go test -timeout 60s github.com/zeroboo/randomselector -v
```

## Publish
Example with VERSION=v0.0.2

- Tag on git

```console
git tag $VERSION
git push $VERSION
```

- Publish go

```console
SET GOPROXY=proxy.golang.org 
go list -m github.com/zeroboo/randomselector@$VERSION
```
