### File structure

- `go.mod` and `go.sum` - contains the dependency links, versions, and checksums to ensure dependency match
- `main.go` - contains all the endpoints, structures, and logic for creating the products API
- `Makefile` - automatically re-renders the page logic to new changes while editing while the `make dev` command is running
- `internal/controllers` folder - contains all the controllers for each of the 5 endpoints. The controller receives the request, optionally performs an action, and returns a response. In this case, the controllers are:
  - getProducts: returns all products from the database in json format
  - getProduct: returns specified single product in json format
  - postProduct: creates a new product in database with provided input json values
  - deleteProduct: removes specified single product from database
  - putProduct: updates specified single product in json format with provided input json values

### Workflow

- run `go mod tidy` to grab the dependencies
- start the service by either running `make dev` if you want to make changes and have the changes render as you go; or run `go run main.go` if you don't need to
- Use postman to run the get, post, put, delete requests
  - getProducts: http://localhost:3000/products
  - getProduct, postProduct, deleteProduct, and putProduct: http://localhost:3000/products/`<some guid>`
    - can get the guid by viewing results of getProducts or searching in sqlite
- If interested in further development that requires dependencies not in `go.mod`, run `go get <library url without https:// prefix> `and the `go.mod` file will automatically update with the new dependency.

### Database:

- uses sqlite3 to store data
- pasted `schema.sql` contents into `sqlite3 ./databadb` to create the products table (`.tables`) and its schema (`.schema`)
