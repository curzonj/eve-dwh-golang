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



# Observations

This was gather via linear regression:

`-59.48*HeadRadius + 1.51565 = Percentage actual of QtyPerCycle`

HeadRadius increases on average by 0.0000300676 per 15 minutes of Extraction cycle time. The stddev is 0.69% of the average.
