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

// version: 2.1

// jobs:
//   build:
//     docker:
//       - image: circleci/golang:1.16
//     working_directory: /go/src/github.com/<your_username>/go-server-ci-cd
//     steps:
//       - checkout
//       - run: go mod download
//       - run: go test ./...
//       - run: go build -o server .
//       - run:
//           name: Build and push Docker image
//           command: |
//             echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
//             docker build -t <your_dockerhub_username>/go-server-ci-cd .
//             docker push <your_dockerhub_username>/go-server-ci-cd

// workflows:
//   version: 2
//   build:
//     jobs:
//       - build
