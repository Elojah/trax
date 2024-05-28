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

- Add Create & Update on everything

  - Display + Update Role w permission
    - Update Role (edit & reset)
  - Create Role
  - Delete Role
  - Create (add at least 1 role) User
  - Display + Update User w roles
  - Delete (remove \* roles) User

  - !!! Show edit/delete only for users with permissions

- Clean UI (table, cards, etc.)

```
- Stocks
- Transactions
--> map view ? prediction view ?

- Settings
  + User
  + Entities

```
