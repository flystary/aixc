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
  search      Print multiple or single rows of data in table form
  show        Print your filteCyan data in tabular form

Flags:
  -h, --help   help for aixc

Use "aixc [command] --help" for more information about a command.
```
```bash
[root@centos ~]# aixc list -h
Use list to display single or multiple ucpe data

Usage:
  aixc list [<SN>, <SN>,....] [flags]

Flags:
  -h, --help           help for list
  -m, --mode string    Appoint the UCPE Mode, Option is valor|tassadar|watsons|watsonsha
      --seven          Appoint that the ucpe Mode belongs to SDWAN6
  -w, --write string   Write current data to a file
```
```bash
[root@centos ~]# aixc show -h
Use show to list the UCPE of the specified options according to your filter

Usage:
  aixc show [flags]

Flags:
  -e, --enterprise string   Appoint that the filtering object of UCPE is the enterprise number
  -h, --help                help for show
  -m, --mode string         Appoint the UCPE Mode, Option is valor|tassadar|watsons|watsonsha (default "valor")
  -s, --select string       Appoint the filter object of UCPE, Option is model|version|pop|enterprise|port|update
  -w, --write string        Write current data to a file
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
```bash
[root@centos ~]# aixc search -h
Use search to display single or multiple ucpe data

Usage:
  aixc search [<SN>, <SN>,....] [flags]

Flags:
  -h, --help           help for search
  -w, --write string   Write current data to a file
```
