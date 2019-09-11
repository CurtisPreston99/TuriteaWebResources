



mkdir ../src
mkdir ../src/TuriteaWebResources
cp -r ./* ../src/TuriteaWebResources



function buildAndRun {
  cd ..
  local x=$(pwd)
  echo $x
  export GOPATH=$x
  cd src/TuriteaWebResources
  go get github.com/lib/pq
  go get github.com/ChenXingyuChina/asynchronousIO
  echo "compiling"
  go build ./server/main.go
  echo "running"
  ./main
}

function finish() {
  pwd
  echo "cleaning up"
  rm -rf ../../src
  echo "removed SRC"
  rm -rf ../../pkg
  echo "removed PKG"
}

trap finish EXIT


buildAndRun
