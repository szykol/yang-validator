# Yang Validator

WIP

Simple backend for Yang Validator app.
The aim is to provide easy adhoc yang validation.

## Credits

This project uses [openconfig/goyang](https://github.com/openconfig/goyang) for the yang validation.

## Usage

For MVP implementation the yang validation is done via HTTP server to get the project up & running without need of creating my own protocol.
This will be adjusted to the needs of frontend and probably changed in the future.

The one and only endpoint is `/validate` that simply accepts bytes and returns `200 OK` if provided yang is valid. Otherwise `400 Bad Request` will be returned with validation error details sent as bytes.

> **NOTE**: User inputs are not yet sanitized or validated and are passed directly to the parsing lib.

```bash
# example of valid yang
curl localhost:1337/validate -d 'module small { namespace "http://example.com/small"; prefix small; }'
OK
# example of invalid yang
curl localhost:1337/validate -d 'module {'
Error: Invalid YANG: :2:0: missing 1 closing brace
```
