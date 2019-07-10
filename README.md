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

### Load variables

```
$ gitlabenv load egegunes/gitlabenv gitlabenv.json
```
