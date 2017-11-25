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

You'll need the url for the ZIP file SDE package from https://developers.eveonline.com/resource/resources

### Updating/Initializing the SDE in heroku

```bash
h run scripts/sde-yaml-etl <url>
```

### Updating the SDE locally

```bash
$ forego run ./scripts/sde-yaml-etl <url>
```
