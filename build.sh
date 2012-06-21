#!/bin/bash

if [ $0 != "./build.sh" ]; then
   echo 'error! build.sh requires execution with the path of "./build.sh"'
   exit 0
fi

echo -n 'build...'

DIR=`pwd`

export GOPATH=$DIR:$GOPATH

go install ./...

echo '[done]'
echo '-------------------------------------------------------------'
echo -e "|  \033[40m\033[31mgolor \033[32mis \033[33min \033[34m\"bin/golor\"\033[0m                                  |"
echo -e "|  \033[40m\033[35mtry \033[36m\"cat \033[33msrc/hantuo.org/golor/golor.go \033[31m| \033[32mbin/golor\"\033[0m      |"
echo '-------------------------------------------------------------'
