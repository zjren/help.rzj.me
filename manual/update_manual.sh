#/bin/sh
cd $GOPATH/src/help.rzj.me/manual
cd ./phpbook
git pull origin master
cd ../gobook
git pull origin master
cd ../gopkg
git pull origin master
cd ../golang-china
git pull origin master
cd ../build-web-application-with-golang
git pull origin master
cd ../the-way-to-go_ZH_CN
git pull origin master

