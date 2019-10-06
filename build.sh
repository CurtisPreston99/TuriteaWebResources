

rm -rf ../src

mkdir ../src
mkdir ../src/TuriteaWebResources
cp -r ./* ../src/TuriteaWebResources



function build {
  cd ..

  local x=$(pwd)

  echo $x

  export GOPATH=$x

  cd src/TuriteaWebResources
  go get github.com/lib/pq
  go get github.com/ChenXingyuChina/asynchronousIO v0.0.0-20190821022857-384d90b77e26
  go build ./server/main.go
}



build
