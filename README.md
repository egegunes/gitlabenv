# gitlabenv

Managing CI variables in Gitlab is so hard, I had to write this.

## Config

`gitlabenv` looks for its configuration file in three places:

1. $HOME/.gitlab.json
2. $HOME/.config/gitlabenv/.gitlab.json
3. $PWD/.gitlab.json

## Usage

### Help

```
$ gitlabenv help
```

### List variables

```
$ gitlabenv list egegunes/gitlabenv
```

### Dump variables

```
$ gitlabenv dump egegunes/gitlabenv > gitlabenv.json
```

### Load multiple variables

```
$ gitlabenv load egegunes/gitlabenv gitlabenv.json
```

### Add variable

```
$ gitlabenv add egegunes/gitlabenv GO_VERSION 1.12
```

### Update variable

```
$ gitlabenv update egegunes/gitlabenv GO_VERSION 1.13
```

### Delete variable

```
$ gitlabenv delete egegunes/gitlabenv GO_VERSION
```
