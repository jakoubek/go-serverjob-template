# go-serverjob-template

My template for programs that are run periodically on a server. Just clone to get a starting point.

These programs typically:

- connect to a database
- retrieve records
- render templates with that data
- send emails

It is checked for a configuration file _config.json_ in the same directory as the binary.

## Configuration

```json
{
  "database": {
    "dsn": ""
  },
  "email": {
    "host": "localhost",
    "port": 25,
    "auth": false,
    "username": "",
    "password": "",
    "sender_name": "",
    "sender_address": ""
  }
}
```

## Magefile

It comes with a [Magefile](https://magefile.org/):

- `mage build`
- `mage deploy`
- `mage clean`

## External depedencies

- [jakoubek/emaillib](https://github.com/jakoubek/emaillib) (and therefore [jordan-wright/email](https://github.com/jordan-wright/email))
- [alexbrainman/odbc](https://github.com/alexbrainman/odbc)
- [magefile/mage](https://github.com/magefile/mage)
