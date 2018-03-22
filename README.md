# Golang based IP echo service

This is replacing the shitty php script that was previously my IP echo service.

It's written/designed to run behind a reverse proxy on the system it's installed (like `nginx`).

## Build instructions
```
# env GOOS=$TARGET-OS GOARCH=$TARGET-ARCH go build -v ip_echo.go
env GOOS=linux GOARCH=amd64 go build -v ip_echo.go
```

## Example Nginx proxy config:

Replace `$name` as needed
```
server {
        listen          80;
        listen          [::]:80;

        server_name     $name;

        access_log      /var/log/nginx/$name/access.log;
        error_log       /var/log/nginx/$name/error.log;

        location / {
                proxy_pass http://localhost:8080;
        }
}
```

## Example systemd service file to start the process:
```
[Unit]
Description=Golang webserver to echo the client's IP address
After=network.target network-online.target nss-lookup.target
Before=nginx.service

[Service]
Type=simple
SyslogLevel=err

ExecStart=/path/to/ip_echo
KillSignal=SIGINT
KillMode=mixed

[Install]
WantedBy=multi-user.target
```

## License
This project is jokingly licensed under the [WTFPL](http://www.wtfpl.net/). Go nuts.

Full license can be found in the 'LICENSE.md' file.
