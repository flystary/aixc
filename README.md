# aixc
使用golang实现的cli

```bash
xc 2.0.0
flyZer0 <flyoney@163.com>

USAGE:
    xc [SUBCOMMAND]

FLAGS:
    -h, --help       Prints help information
    -V, --version    Prints version information

SUBCOMMANDS:
    connet    Connect can be used to remotely connect CPE and display the process on the terminal.
    help      Prints this message or the help of the given subcommand(s)
    show      Use show to obtain CPE information and display it on the current terminal
    update    Use update to update local CPE information

Use show to obtain CPE information and display it on the current terminal

USAGE:
    xc show [OPTIONS] <sn>

FLAGS:
    -h, --help       Prints help information
    -V, --version    Prints version information

OPTIONS:
    -m <mode>        Use show to select the CPE, the default version is nexus. [possible values: nexus, valor, watsons,
                     watsons_ha, tassadar]

ARGS:
    <sn>    cpe serial number
```