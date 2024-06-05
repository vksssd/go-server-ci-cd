version: 2.1

jobs:
  build:
    docker:
      - image: circleci/golang:1.16
    working_directory: /go/src/github.com/vksssd/go-server-ci-cd
    steps:
      - checkout
      - run: go mod download
      - run: go test ./...
      - run: go build -o server .

workflows:
  version: 2
  build:
    jobs:
      - build
