# Github Action: Webhook

[![Github Action](https://github.com/josa42/gh-action-webhook/workflows/Master/badge.svg)](https://github.com/josa42/gh-action-webhook/actions)

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

```yaml
- uses: 'josa42/gh-action-webhook/release@master'
  with:
    url:           'http://httpbin.org/post'
    authorization: 'Bearer ${{ secrets.TOKEN }}'
    data:          '{ "data": "TEST" }'
```

You can choose one of these versions:

- **Latest Release:** `josa42/gh-action-webhook/release@master`
- **Latest Build** `josa42/gh-action-webhook/latest@master`
- **A specific Release:** `josa42/gh-action-webhook/release@0.1.0`
- **A specific commit:** `josa42/gh-action-webhook/build@54e5a00`  
  (This will be slow since, the workflow needs to be build on every run)

## Examples

**Post a Slack Message**

```yaml
name: 'New Tag'
on:
  tag: [ '*' ]
jobs:
  test:
    name:    'Publish master'
    runs-on: 'ubuntu-latest'
    steps:
      - uses: 'josa42/gh-action-webhook/latest@master'
        with:
          url:  'https://hooks.slack.com/services/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'
          data: '{ "text": "New Tag: ${{github.ref}}!" }'

```

See [Slack Documentation](https://api.slack.com/incoming-webhooks#enable_webhooks) for more infos.

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
