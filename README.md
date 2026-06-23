# e-backend-boilerplate

`e-backend-boilerplate` is a backend for all the projects.

![cat](./modules/doc/data/public/android-chrome-192x192.png)

```
▗▄▄▄▖▗▄▄▖  ▗▄▖  ▗▄▄▖▗▖ ▗▖▗▄▄▄▖▗▖  ▗▖▗▄▄▄ 
▐▌   ▐▌ ▐▌▐▌ ▐▌▐▌   ▐▌▗▞▘▐▌   ▐▛▚▖▐▌▐▌  █
▐▛▀▀▘▐▛▀▚▖▐▛▀▜▌▐▌   ▐▛▚▖ ▐▛▀▀▘▐▌ ▝▜▌▐▌  █
▐▙▄▄▖▐▙▄▞▘▐▌ ▐▌▝▚▄▄▖▐▌ ▐▌▐▙▄▄▖▐▌  ▐▌▐▙▄▄▀
```

## Create database

```bash
sudo -iu postgres
createdb ebackend
```

## Prepare

```bash
cp .config.example .config
```

And then edit `.config` file.

## Run HTTP-server

```bash
make run
```

## Building

Build binary:

```bash
make build
```

Clean:

```bash
make clean
```

Run binary:

```bash
./build/e-backend-boilerplate serve
```

## Run as service (POSIX systems with systemd)

```bash
sudo mkdir /opt/e-backend-boilerplate
sudo cp ./build/* /opt/e-backend-boilerplate -r
sudo cp /opt/e-backend-boilerplate/.e-backend-boilerplate.example /opt/e-backend-boilerplate/.e-backend-boilerplate
sudo cp ./systemd/e-backend-boilerplate.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now e-backend-boilerplate.service
```

## Module generation

```bash
go run -tags="all dev" main.go module create [name] -t crud
```

Where `name` is name of module is `lowerCamelCase`, `-t` is template name
(simple, crud).

# Run with docker

Build image:

```bash
docker build -f dockerfiles/Dockerfile -t e-backend-boilerplate .
```

Run:

```bash
docker run --name e-backend-boilerplate --rm --network host \
-v "$(pwd)/.e-backend-boilerplate:/app/.e-backend-boilerplate" \
e-backend-boilerplate
```
