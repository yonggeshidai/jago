#!/bin/bash

export JAGO_HOME=`pwd`

export GOPATH=`pwd`
go get github.com/urfave/cli
go get github.com/orcaman/concurrent-map
go get github.com/fatih/color
go build -o bin/jago
#export GOBIN=$JAGO_HOME/bin

if [ "$JAGO_HOME" == '' ]; then
    sudo mkdir -p /usr/local/jago
    sudo chown -R $USER /usr/local/jago
    mkdir -p /usr/local/jago/bin
    mkdir -p /usr/local/jago/log
    export JAGO_HOME=/usr/local/jago
    echo "export JAGO_HOME=$JAGO_HOME" >> ~/.bash_profile
    source ~/.bash_profile
fi

sudo cp bin/* $JAGO_HOME/bin/
sudo ln -sf $JAGO_HOME/bin/jago /usr/local/bin/jago

sudo cp -rf jdk $JAGO_HOME/

#go build -buildmode=c-shared -o build/libgvm.dylib launcher.go
#go tool cgo -srcdir src -objdir build/gni -exportheader build/gni.h -- -Isrc/gvm  gvm/gni.go
