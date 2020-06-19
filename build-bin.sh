#/bin/sh

VERSION=$(./git-semver.sh)

cat << EOF > version.go
package main

var Version = "$VERSION"
EOF

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/deployf .