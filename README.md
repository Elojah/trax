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
$ docker-compose up -d postgres redis
$ make admin && ./bin/trax_admin config/admin/local.json
$ make api && ./bin/trax_api config/api/local.json
$ make auth && ./bin/trax_auth config/auth/local.json
$ make client && ./bin/trax_web_client config/web_client/local.json
```

TODO:

- Add Create & Update & delete on everything

  - Create/Invite (add at least 1 role) User
  - Edit role name
  - !!! Show edit/delete only for users with permissions -> Manage permissions globally with v-if to hide impossible action
    + Check user table w/ permission

- [S] Search bar on right panel entity fill everything until new+ button
- [S] Table footer in nested table must be smaller than table footer in container


- ~ Clean UI (table, cards, etc.)

- Find a way to have expire map for pinia stores

```
- Stocks
- Transactions
--> map view ? prediction view ?

- Settings
  + User
  + Entities

```
