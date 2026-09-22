Pet project to be a playground for golang and different integrations.

# Details

The basic flows are 

```shell
# start dependencies and server
task docker:up server 

# Create products (uses mongo)
curl --request POST \
  --url http://localhost:8080/api/products \
  --header 'Content-Type: application/json' \
  --data '{
  "name": "TV",
  "description": "LG TV",
  "price": 10,
  "availableQuantity": 100
}'

# Add items to a cart (users redis)
curl --request POST \
  --url http://localhost:8080/api/cart \
  --header 'Content-Type: application/json' \
  --data '{
  "productId": "test",
  "quantity": 10
}'

# Purchase - save to postgres, publish to kafka
curl --request POST \
  --url http://localhost:8080/api/orders \
  --header 'Content-Type: application/json' \
  --data '{
  "orderId": "128"
}'
```

Cli is also a client:

```shell

# Insert a product direct on mongo database
./bin/cli insert-product

# Insert a product calling service using GRPC
./bin/cli grpc-insert-product
```

# Tools

- [mise](https://mise.jdx.dev/) for dependency manager (similar to [asdf](https://asdf-vm.com/))
- [Taskfile](https://taskfile.dev/) for task/dependency chain
- Golang
- Docker

# Pending

- Graphql
- MCP server
- SQS / SNS
- Auth
- Dockerfile
- Tests
- K8S descriptors
- Swagger
- Performance Tests