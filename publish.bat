SET VERSION=v0.1.1

git tag %VERSION%
git push origin %VERSION%

SET GOPROXY=proxy.golang.org 
go list -m github.com/zeroboo/randomselector@%VERSION%
