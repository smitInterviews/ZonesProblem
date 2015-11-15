mkdir -p gowork/src/github.com
mkdir -p gowork/bin
mkdir -p gowork/pkg
export GOPATH=$HOME/gowork
cd $GOPATH/src/github.com
git clone git@github.com:smitInterviews/ZonesProblem.git
pushd ZonesProblem/server
go get ...
go build
./server
popd
pushd tests
go get ...
popd 
