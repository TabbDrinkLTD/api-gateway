cd ~/go/src/github.com/micro/kubernetes/cmd/micro
#git pull
#go get -u

echo "Copying custom plugins to micro project \n"
cp ~/go/src/github.com/TabbDrinkLTD/api-gateway/plugins.go ./plugins.go

echo "Compiling... \n"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i -o micro ./main.go ./plugins.go

echo "Building docker image \n"
docker build -t "eu.gcr.io/tabb-168314/tabb-gateway:latest" .
gcloud docker -- push "eu.gcr.io/tabb-168314/tabb-gateway:latest"

echo "Moving binary back to api-gateway project"
cp micro ~/go/src/github.com/TabbDrinkLTD/api-gateway
cd ~/go/src/github.com/TabbDrinkLTD/api-gateway

echo "Done!"
ls -la
