# aixc
使用golang实现的cli

```bash
[root@centos ~]# aixc show -h
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
  -m, --Mseven   If ucpe belongs to 6.X platform, Please use it
  -h, --help     help for list
```
```bash
[root@centos ~]# aixc show -h
use show to show everything you want form cpe

Usage:
  aixc show <SN> [flags]

Flags:
  -m, --Mseven   If ucpe belongs to 6.X platform, Please use it
  -h, --help     help for show
```