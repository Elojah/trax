# trax

Trading platform network

Local development setup:

Once:

```sh
# Protobuf setup
$ sudo apt install -y protobuf-compiler
# Protobuf golang with gogoproto
$ export PATH=$PATH:$GOPATH/bin
$ go get github.com/gogo/protobuf/proto
$ go install github.com/gogo/protobuf/protoc-gen-gogoslick
# Protobuf TS (install node + npm latest versions with nvm)
$ sudo npm i -g npx
$ cd cmd/client && npm install
# Host aliases for local development
$ cat scripts/add_localhost.hosts | sudo tee -a /etc/hosts > /dev/null
```

Start:

```sh
docker-compose up -d postgres redis
make admin && ./bin/trax_admin config/admin/local.json
make api && ./bin/trax_api config/api/local.json
make auth && ./bin/trax_auth config/auth/local.json
make client && ./bin/trax_web_client config/web_client/local.json
```

TODO:

- [x] Reactivity of tables and details after ALL actions
- [x] Small UI pas on tag/status

- [x] Multi roles user bug when listing gorup users
  - group icons still bug
- [x] Invite User
  - Finish invite_user route to associate existing user
  - Add invitations tab under group vue

- [Map]
  - Add clean bounding and centering/zooming on click/move
  - Add reverse Geocoding to show location info in panel
  - Allow panel to upload video/images/live video on webrtc
  - Allow panel to share Marker with some groups only

- Write test flows

- Find a way to have expire map for pinia stores

```
- Settings
  + User

```
