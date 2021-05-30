FROM golang:1.16.3-alpine

WORKDIR $GOPATH/src/github.com/simpleWebApp

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 80

CMD ["simpleWebApp", "80"]
