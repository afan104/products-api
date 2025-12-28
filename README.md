File structure

- `go.mod` and `go.sum` - contains the dependency links, versions, and checksums to ensure dependency match
- `main.go` - contains all the endpoints, structures, and logic for creating the products API
- `Makefile` - automatically re-renders the page logic to new changes while editing while the `make dev` command is running
- `assets/` folder -
- `templates/` folder -
- `internal/controllers` folder - contains all the controllers for each of the 5 endpoints. The controller receives the request, decides what to do, and returns a response. In this case, the controllers all define the status code returned upon routing to the api and the response result.

Workflow

- run requirements from `` file
- run `go mod tidy` if you run into any dependency errors and this will fix them.
- start the service by either running `make dev` if you want to make changes and have the changes render as you go; or run `go run main.go` if you don't need to
- If you need to import another dependency not in `go.mod`, run `go get <library url without https:// prefix> `and the `go.mod` file will automatically update with the new dependency.

Database:

- uses sqlite3 to store data
- pasted `schema.sql` contents into `sqlite3 ./databadb` to create the products table (`.tables`) and its schema (`.schema`)
