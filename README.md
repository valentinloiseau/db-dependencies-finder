# db-dependencies-finder

Guess dependencies accross non-relational DB tables and find broken ones.

## Get Started

```bash
$ docker build . -t db-dependencies-finder
$ docker run -it db-dependencies-finder /bin/bash
$ go run .
```

Put config parameters in ``.env`` file for not be prompted :

```env
FOREIGN_KEY_PATTERN=
DB_USER=
DB_PWD=
DB_HOST=
DB_PORT=
DB_NAME=
```
