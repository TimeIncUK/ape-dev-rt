#!/usr/bin/env bash
#
# This script builds the application from source for multiple platforms.

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# Get the Terraform git commit
TF_PATH="vendor/github.com/hashicorp/terraform"
TF_GIT_COMMIT=$(git -C $TF_PATH rev-parse HEAD)
TF_GIT_DIRTY=$(test -n "`git -C $TF_PATH status --porcelain`" && echo "+CHANGES" || true)

# Delete the old dir
echo "==> Removing old directory..."
rm -f bin/*
rm -rf pkg/*
mkdir -p bin/

# Determine the arch/os combos we're building for
XC_OS=${XC_OS:-linux darwin windows}
XC_ARCH="amd64"

# If its dev mode, only build for ourself
if [ "${RT_DEV}x" != "x" ]; then
    XC_OS=$(go env GOOS)
    XC_ARCH=$(go env GOARCH)
fi

# Build!
echo "==> Building..."
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -ldflags "-X github.com/TimeIncOSS/ape-dev-rt/rt.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/TimeIncOSS/ape-dev-rt/rt.TerraformCommit=${TF_GIT_COMMIT}${TF_GIT_DIRTY}" \
    -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}" \
    $(go list ./... | grep -v /vendor/)

# Move all the compiled things to the $GOPATH/bin
GOPATH=${GOPATH:-$(go env GOPATH)}
case $(uname) in
    CYGWIN*)
        GOPATH="$(cygpath $GOPATH)"
        ;;
esac
OLDIFS=$IFS
IFS=: MAIN_GOPATH=($GOPATH)
IFS=$OLDIFS

# Create GOPATH/bin if it's doesn't exists
if [ ! -d $MAIN_GOPATH/bin ]; then
    echo "==> Creating GOPATH/bin directory..."
    mkdir -p $MAIN_GOPATH/bin
fi

# Copy our OS/Arch to the bin/ directory
DEV_PLATFORM="./pkg/$(go env GOOS)_$(go env GOARCH)"
for F in $(find ${DEV_PLATFORM} -mindepth 1 -maxdepth 1 -type f); do
    cp ${F} bin/
    cp ${F} ${MAIN_GOPATH}/bin/
done
