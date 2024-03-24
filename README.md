# trax

Trading platform network

Local development setup:

Once:

```sh
$ cd cmd/client && npm install
$ GO111MODULE=off go get github.com/gogo/protobuf/proto
$ go install github.com/gogo/protobuf/protoc-gen-gogoslick
$ cat scripts/add_localhost.hosts | sudo tee -a /etc/hosts > /dev/null
```

Start:

```sh
$ docker-compose up -d
$ make admin && ./bin/game_03_admin config/admin/local.json
$ make api && ./bin/game_03_api config/api/local.json
$ make auth && ./bin/game_03_auth config/auth/local.json
$ make client && ./bin/game_03_web_client config/web_client/local.json
```

TODO:

- Fetch firstname lastname from google token

- Create profile page

- Redirect in profile when already logged in

- Check refresh and redirect when already logged in

- Leaflet ?
