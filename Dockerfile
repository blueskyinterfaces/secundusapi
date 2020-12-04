FROM golang:1.15.5

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

CMD ["go", "run", "cmd/api/main.go"]

EXPOSE 3000

## USAGE
## ## ## ## ## ## ## ## ## ## ## ## 
# BUILD CONTAINER
# docker build -t secundus .

# RUN (LOCAL) 
# docker run --rm -d -p 3000:3000/tcp secundus:latest
# RUN (NOSTART) 
# docker run -d -v secundus:latest -p 3000:3000/tcp secundus:latest tail -f /dev/null
# INTERACTIVE
# docker run --rm -it -p 3000:3000/tcp secundus:latest
# SHELL INTO RUNNING CONTAINER     
# docker exec -it {container_id} bash
# STOP
# docker container list
# docker container stop {containerid}
# docker container prune

## AZURE CONTAINERS (user/pass on azure portal)
## LOGIN         docker login secundus1.azurecr.io
## TAG           docker tag secundus:latest secundus1.azurecr.io/secundus
## PUSH          docker push secundus1.azurecr.io/secundus
## PULL          docker pull secundus1.azurecr.io/secundus
## DATA          docker cp {container_id}:/flight_review/data ./datacopy

## TO DEPLOY FROM VSCODE
## docker build -t secundus .
## docker tag secundus:latest secundus1.azurecr.io/secundus
## docker login secundus1.azurecr.io
## docker push secundus1.azurecr.io/secundus
## goto the Docker tab in open the azure registry, find the container, right-click and deploy 

## TO REDEPLOY
## docker build -t secundus .
## docker tag secundus:latest secundus1.azurecr.io/secundus
## docker login secundus1.azurecr.io
## docker push secundus1.azurecr.io/secundus
## goto the app in azure and save the Container settings page to trigger a refresh
