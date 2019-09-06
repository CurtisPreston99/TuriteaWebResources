

rm -rf ../src

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
  go build ./server/main.go
  ./main
}



buildAndRun
