# golang-clean-architecture-ent-gqlgen
Clean Architecture with ent and gqlgen

## Run Docker

```
$ cd docker
$ docker comopse up
```

## Install

```
$ make install
```

## Set up database

```
$ make setup_db
$ make migrate_schema
```

## Start server

```
$ make start
```

## Testing

```
$ make setup_test_db
$ make test_repository
```

## E2E

```
$ make setup_e2e_db
$ make e2e
```
