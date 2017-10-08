# eve-dwh-golang

### migrations

```bash
forego run scripts/dbmigrate [up|down]
```

### Testing

```bash
heroku local:start dev
```

## SDE

You can download the SDE from [Fuzzworks](https://www.fuzzwork.co.uk/dump/postgres-latest.dmp.bz2)

### Updating/Initializing the SDE in heroku

```bash
$ h config -s
# Copy the DATABASE_URL
$ env DATABASE_URL=$(value_from_heroku) ./scripts/sde-load ./dumps/postgres-latest.dmp
```

### Updating the SDE locally

```bash
$ forego run ./scripts/sde-load ./dumps/postgres-latest.dmp
```
