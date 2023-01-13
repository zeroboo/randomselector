SET VERSION=v0.1.2

git tag %VERSION%
git push origin %VERSION%

SET GOPROXY=proxy.golang.org 
go list -m github.com/zeroboo/randomselector@%VERSION%
