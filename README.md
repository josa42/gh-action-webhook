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

```yaml
- uses: 'josa42/gh-action-webhook/latest@master'
  with:
    url:           'http://httpbin.org/post'
    authorization: 'Bearer ${{ secrets.TOKEN }}'
    data:          '{ "data": "TEST" }'
```

To use a specific version of the work flow add a step like this:

```yaml
- uses: 'josa42/gh-action-webhook/build@d111e4e'
  with:
    url: 'http://httpbin.org/post'
```

## Examples

`.github/workflows/master,yaml`

```yaml
name: 'Example'
on:
  push:
    branches: [ 'master' ]
jobs:
  test:
    name:    'Run Webhooks'
    runs-on: 'ubuntu-latest'
    steps:
      - name: 'POST'
        uses: 'josa42/gh-action-webhook/lastest@master'
        with:
          url: 'http://httpbin.org/post'
```

**Rebuild Github Page**

```yaml
name: 'Schedule'
on:
  schedule:
    - cron: '*/15 * * * *'
jobs:
  test:
    name:    'Publish master'
    runs-on: 'ubuntu-latest'
    steps:
      - uses: 'josa42/gh-action-webhook/latest@master'
        with:
          url:           'https://api.github.com/repos/josa42/josa42.github.io/pages/builds'
          authorization: 'Token ${{ secrets.GITHUB_TOKEN }}'
          accept:        'application/vnd.github.mister-fantastic-preview+json'

```


## License

[MIT © Josa Gesell](LICENSE)
