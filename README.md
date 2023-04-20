# tt

A utility that retrieves the latest SemVer tag from a repository and parses it into component parts that can be used as transient tags.

```sh
$ tt
v1.2.3,v1,v1.2
```

The latest tag `v1.2.3` would be parsed into: `{{Full}},{{Major}},{{Minor}}`.

## Explicit SemVer

If explicit SemVer versioning is required, set environment variable `TT_SEMVER` to `1` while running `tt`:

```sh
$ TT_SEMVER=1 tt
1.2.3,1,1.2
```
