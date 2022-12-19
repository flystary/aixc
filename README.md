# aixc
使用golang实现的cli

```bash
[root@centos ~]# aixc -h
Usage:
  aixc [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  conn        Use this option to connect CPE or POP remotely
  exec        Use this option to remotely execute commands on UCPE or POP
  help        Help about any command
  list        Print multiple or single rows of data in table form
  show        Print your filtered data in tabular form

Flags:
  -h, --help   help for aixc
```
```bash
[root@centos ~]# aixc list -h
Use list to display single or multiple ucpe data

Usage:
  aixc list [<SN>, <SN>,....] [flags]

Flags:
  -h, --help          help for list
  -m, --mode string   Appoint the UCPE Mode
      --seven         Appoint that the ucpe Mode belongs to SDWAN6
```
```bash
[root@centos ~]# aixc show -h
Use show to list the UCPE of the specified options according to your filter

Usage:
  aixc show [flags]

Flags:
  -e, --enterprise string   Appoint that the filtering object of UCPE is the enterprise number (default "null")
  -h, --help                help for show
  -m, --mode string         Appoint the UCPE Mode
```
```bash
[root@centos ~]# aixc conn -h
Use conn to remotely connect the UCPE and POP you need

Usage:
  aixc conn <SN> [flags]

Flags:
  -h, --help    help for conn
      --seven   Appoint that the UCPE Mode belongs to SDWAN6
```