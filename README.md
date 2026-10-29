# Resume Theme Default

[![Build](https://github.com/ansidev/resume-theme-default/actions/workflows/deploy-demo.yml/badge.svg)](https://github.com/ansidev/resume-theme-default/actions/workflows/deploy-demo.yml)

## Introduction

This is the default theme for [Resume CLI](https://github.com/ansidev/resume-cli-ce).

## Getting started

### Installation

Download the latest `resume-cli` release from [ansidev/resume-cli-ce](https://github.com/ansidev/resume-cli-ce/releases).

### How to use?

Create a `resume.config.yaml` in your project. Choose `local` if you have cloned this theme, or `remote` to let Resume CLI fetch it automatically:

```yaml
theme:
  type: remote                                        # or: local
  location: https://github.com/ansidev/resume-theme-default.git  # or: path/to/theme
resume:
  basics:
    name: Your Name
    # ... see demo/resume.config.yaml for the full structure
```

Then generate your resume:

```bash
# HTML
resume-cli generate --format html --output dist/index.html

# PDF
resume-cli generate --format pdf --output dist/static/resume.pdf
```

See [`demo/resume.config.yaml`](demo/resume.config.yaml) for a complete working example. The live demo is published at https://resume-theme-default-demo.pages.dev.

> **Remote helpers:** when `type: remote` is set, Resume CLI refuses to run `helpers.go` unless you
> review the file and pass `--allow-remote-helpers`:
>
> ```bash
> resume-cli generate --allow-remote-helpers
> ```
>
> `local` themes run helpers automatically without the flag.

The `--theme` option overrides `theme.location`. The template engine and helpers are declared by this theme in [`theme.yaml`](/theme.yaml); they are not set in the Resume CLI config.

### Helper functions

[`helpers.go`](/helpers.go) is interpreted by Resume CLI at runtime (with [yaegi](https://github.com/traefik/yaegi)).
Its `Funcs() map[string]interface{}` function returns the template helpers: `trimPrefix`, `bootstrapIcon`, `joinArray`,
`markdown`, `displayUrl`, `toLowerCase` and `year`. The file must stay self-contained (single file, standard library,
`github.com/mailgun/raymond/v2` and `github.com/yuin/goldmark` only).

> **Security:** helper code runs with the full privileges of the CLI process. It runs automatically for `local` themes. For `remote` themes, pass `--allow-remote-helpers` as described above.

### Theme structure

```
|-- theme.yaml: theme manifest (name, templateEngine, helpers)
|-- resources: view files
|   |-- assets
|   |   |-- css
|   |   |   |-- styles.css
|   |   |-- js
|   |       |-- scripts.js
|   |-- partials: partial files
|   |-- resume.{template_extension}: file extension depends on the used template engine.
|-- helpers.go: helper functions (entry point: Funcs)
|-- helpers_test.go
|-- go.mod
|-- go.sum
```

## Demo deployment

The [demo configuration](demo/resume.config.yaml) is copied from Resume CLI's
sample configuration. GitHub Actions builds it with the latest published
[`resume-cli`](https://github.com/ansidev/resume-cli-ce) release and deploys
the resulting site to Cloudflare Pages on pushes to `main` (or when started
manually). Pull requests only build the demo and upload it as an artifact.

The demo config uses `theme.location: .`, so run `resume-cli generate
--config demo/resume.config.yaml` from the repository root.

Before enabling the workflow, configure these repository values in GitHub:

- `CLOUDFLARE_API_TOKEN` secret: an API token with **Account → Cloudflare
  Pages → Edit** permission.
- `CLOUDFLARE_ACCOUNT_ID` secret: the Cloudflare account ID that owns the
  Pages project.
- `CLOUDFLARE_PAGES_PROJECT_NAME` Actions variable: the existing Cloudflare
  Pages project name.

The workflow uses this local theme checkout and produces:

- `dist/index.html`
- `dist/static/resume.pdf`

## Contact

Le Minh Tri [@ansidev](https://ansidev.xyz/about).

## License

This source code is available under the [AGPL-3.0 LICENSE](/LICENSE).
