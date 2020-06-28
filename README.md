# Synthesis


Synthesis is a lightweight url shortening server written in go.


## API


Api documentation is reachable at the route `{{host}}/swagger/index.html#`

## Architecture

The software is build using plain `net/http` go lib and `github.com/gorilla/mux` for routing.

To store data PostgresQL has been used for semplicity reasons, but plan to support Redis are being considered.

## Build instructions

Just clone the repo and run `go build` or, to build a docker image, just run `docker build -t <your tag> .`

## Configuration

The configuration file must be place in the root path of the project and should be named `configuration.yaml`.
When none or an incomplete one provided, default parameters will be loaded as described:

```yaml
Host:        "127.0.0.1"
Port:        "3000" # must be a string for semplicity reasons
PostgresUri: "postgresql://localhost:5432/synthesis"
KeyTimer:    600
```

`KeyTimer` is the the number of seconds a key will be available on the platform.

## Running

You can simply run the compiled executable or docker container loading your configuration file.

### Local build

```bash
go build .
vim configuration.yaml # write your configuration here and save. To exit vim :q 😬
./synthesis
```

### Docker container

```bash
docker build -t synthesis:latest .
vim configuration.yaml # write your configuration here and save. To exit vim :q 😬
docker run -p 3000:3000 -v "$(pwd)/configuration.yaml:/app/configuration.yaml" synthesis
```

Enjoy!