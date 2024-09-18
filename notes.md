## Test

go test ./... -v -coverprofile cover.out
go tool cover -html cover.out
## Deploy script

```console
```

## Publish

Example with VERSION=v1.0.0

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

