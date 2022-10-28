# aixc
使用golang实现的cli

```bash
[root@centos ~]# aixc -h
use show to show everything you want form cpe
Usage:
  aixc [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  conn        Use this function to connect CPE or node
  help        Help about any command
  list        Print multiple rows of data in tabular form
  show        Print single line data in tabular form
  version     Print the version number of xc

Flags:
  -h, --help   help for aixc
```
```bash
[root@centos ~]# aixc list -h
use list to list everything you want form cpe
Usage:
  aixc list [<SN>, <SN>,....] [flags]

Flags:
  -h, --help          help for list
  -m, --mode string   ucpe mode, Please use it
      --seven         if ucpe belongs to SDWAN6 platform, Please use it
```
```bash
[root@centos ~]# aixc show -h
use show to show everything you want form cpe
Usage:
  aixc show <SN> [flags]
Flags:
  -h, --help    help for show
      --seven   If ucpe belongs to SDWAN6 platform, Please use it
```