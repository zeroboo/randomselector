$VERSION = "v1.0.0"
Write-Host Version ${VERSION}
git tag ${VERSION}
git push origin ${VERSION}

$env:GOPROXY="proxy.golang.org"
go list -m github.com/zeroboo/randomselector@${VERSION}