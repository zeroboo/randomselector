Deploy script:
```console
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
