# helm-envlint
🚧WIP🚧 helm-envlint is lint tool for helm chart values.yaml.

## Required

* helm1.7+

## Installation

Install or upgrade helmenvlint with this command.

```bash
go get -u github.com/laqiiz/helm-envlint
```

## Usage

```bash
$ helmenvlint --help
Usage of helmenvlint:
  -d string
        helm directory (default ".")
  -dir string
        helm directory (default ".")
  -l string
        helm values file path
  -left string
        helm values file path
  -r string
        helm values file path
  -right string
        helm values file path
```

### Example(Diff only)


```bash
$ helmenvlint -d examples/nginx/. -r examples/nginx/values_left.yaml -l examples/nginx/values_right.yam
```
```json output
[
  {
    "data": {
      "index.html": [
        "@@ -30,12 +30,11 @@\n est \n-righ\n+lef\n t%3C/p\n",
        0,
        2
      ]
    }
  },
  {
    "spec": {
      "replicas": [
        1,
        3
      ],
      "template": {
        "spec": {
          "containers": {
            "0": {
              "image": [
                "nginx:amazonlinux",
                "nginx:alpine"
              ]
            },
            "_t": "a"
          }
        }
      }
    }
  }
]
```

## helm-envlint is using below great apps

* `helm tempalte` command
* bronze1man/yaml2json
* benjamine/jsondiffpatch

