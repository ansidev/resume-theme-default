# Resume Theme Default

## Introduction

This is the default theme for [Resume CLI](https://github.com/ansidev/resume-cli).

## Getting started

### How to use?

Please checkout an [example project](https://github.com/ansidev/resume-cli-example) to understand how to use this tool.

### Theme config

Configure the theme in the Resume CLI config file (`resume.config.yaml`):

```yaml
theme:
  type: local # local | remote
  location: ./themes/resume-theme-default # local: path to the theme folder; remote: git URL
```

- `local`: `location` is the path to this theme folder.
- `remote`: `location` is the git URL of this theme, e.g. `https://github.com/ansidev/resume-theme-default`.
- `--theme` overrides `theme.location`.

The template engine and helpers are declared by this theme in [`theme.yaml`](/theme.yaml); they are not set in the
Resume CLI config.

### Helper functions

[`helpers.go`](/helpers.go) is interpreted by Resume CLI at runtime (with [yaegi](https://github.com/traefik/yaegi)).
Its `Funcs() map[string]interface{}` function returns the template helpers: `trimPrefix`, `bootstrapIcon`, `joinArray`,
`markdown`, `displayUrl`, `toLowerCase` and `year`. The file must stay self-contained (single file, standard library,
`github.com/mailgun/raymond/v2` and `github.com/yuin/goldmark` only).

> **Security:** helper code runs with the full privileges of the CLI process. It runs automatically for `local` themes.
> For a `remote` theme, Resume CLI refuses to run it unless you review `helpers.go` and pass `--allow-remote-helpers`:
>
> ```
> resume generate --allow-remote-helpers
> ```

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

## Contact

Le Minh Tri [@ansidev](https://ansidev.xyz/about).

## License

This source code is available under the [AGPL-3.0 LICENSE](/LICENSE).
