# GO Commands

## To install pkg: go.sum will create

go get -u github.com/gorilla/mux

go uses module pakage for our project initially from web then from go path cache folder

<!-- to see that go env command check go path in that will find cache -->

## module creation: go.mod file will create

go mod init guthub.com/authorname/modulename

1. go mod init: Initializes a new module, creating a go.mod file that lists the dependencies for your project.

2. go mod download: Downloads the required packages and dependencies for your project.

3. go mod edit: Opens the go.mod file in your default editor, allowing you to edit the dependencies for your project.

4. go mod graph: Displays a visual representation of the dependencies for your project, including the packages and versions required by each package.

5. go mod tidy: Cleans up your go.mod file and go.sum file by adding new dependencies, removing unused dependencies, and updating the versions of your dependencies as necessary.

6. go mod vendor: Copies all of your dependencies into a vendor directory, allowing you to vendor your dependencies and ensure that your project builds consistently even if the external packages change.

7. go mod verify: Verifies that the packages and dependencies for your project are up-to-date and have been downloaded correctly.

8. go mod why: Shows why a specific package is included as a dependency in your project.

## Run and build:

go build .

<!-- this looks pkgs directly from web or cache -->

go run filename

<!-- this looks pakages from locally vendor folder instead of web or cache -->

go run -mod=vendor filename

## GO list

<!-- list all module pkgs from our projet modules-->

go list -m all

<!-- list of all version of mux dependency  -->

go list -m -versions github.com/gorilla/mux
