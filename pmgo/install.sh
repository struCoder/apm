#!/bin/sh

version="0.5.1"

UNAME=$(uname)
PROJECT_PATH="$GOPATH/src/github.com/struCoder"
OLD_PROJECT_PATH="$GOPATH/src/github.com/struCoder/pmgo"
BACKUP=false

# Check to see if it starts with MINGW.
if [ "$UNAME" ">" "MINGW" -a "$UNAME" "<" "MINGX" ] ; then
    echo "current release do not support windows"
    exit 1
fi
if [ "$UNAME" != "Linux" -a "$UNAME" != "Darwin" ] ; then
    echo "Sorry, this OS is not supported yet via this installer."
    exit 1
fi

# Check GOPATH exist
if [ ! -n "$GOPATH" ] ; then
    echo "please make sure that GOPATH exist"
    exit 1
fi

# create source dir
if [ ! -d "$PROJECT_PATH" ] ; then
  mkdir -p $PROJECT_PATH
fi

# backup if pmgo dir exited
if [ -d "$OLD_PROJECT_PATH" ] ; then
  mv $OLD_PROJECT_PATH "$PROJECT_PATH/pmgo_backup"
  BACKUP=true
fi


cd $PROJECT_PATH

wget https://github.com/struCoder/pmgo/archive/v${version}.tar.gz
tar -zxvf v${version}.tar.gz

mv pmgo-${version} pmgo
cd pmgo

echo "build pmgo..."
go build -o /usr/local/bin/pmgo pmgo.go

echo "build done"

cd ..
echo "rm build files"

rm -rf pmgo-${version}
rm -rf pmgo
rm v${version}.tar.gz

if [ "$BACKUP" = true ] ; then
  mv "$PROJECT_PATH/pmgo_backup" pmgo
fi

echo "all down."
