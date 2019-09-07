# Github Action: Webhook

A [Github Action](https://github.com/features/actions) to trigger a http request.

## Arguments

| Key             | Required | Default | Description                           |
|:----------------|:---------|:--------|:--------------------------------------|
| `url`           | **Yes**  |         |                                       |
| `method`        | No       | `POST`  |                                       |
| `data`          | No       |         |                                       |
| `authorization` |          |         | Set `Authorization` header.           |
| `accept`        | No       |         | Set `Accept` header.                  |

## Usage

To use the latest build add this step to your workflow:

```
- uses: josa42/gh-action-webhook@master
  with:
    url: http://httpbin.org/post
    data: '{ "data": "TEST" }'
    authorization: 'Bearer ${{ secrets.TOKEN }}'
```

To use a specific version of the work flow add a step like this:

```
- uses: josa42/gh-action-webhook/build@d111e4e
  with:
    url: http://httpbin.org/post
```

## Examples

`.github/workflows/master,yaml`

```yml
name: Example
on:
  push:
    branches: [ master ]
jobs:
  test:
    name: Run Webhooks
    runs-on: ubuntu-latest
    steps:
      - name: POST (lastest)
        uses: josa42/gh-action-webhook@master
        with:
          url: http://httpbin.org/post
      - name: POST (commit)
        uses: josa42/gh-action-webhook/build@d111e4e
        with:
          url: http://httpbin.org/post
```

## License

[MIT © Josa Gesell](LICENSE)
