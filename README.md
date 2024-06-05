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
$ make admin && ./bin/trax_admin config/admin/local.json
$ make api && ./bin/trax_api config/api/local.json
$ make auth && ./bin/trax_auth config/auth/local.json
$ make client && ./bin/trax_web_client config/web_client/local.json
```

TODO:

- Add Create & Update & delete on everything

  - Delete Role
  - Create (add at least 1 role) User
  - Display + Update User w roles
  - Delete (remove \* roles) User

  - !!! Show edit/delete only for users with permissions
  - Add pop up for every success/fail actions

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
