module aixc

go 1.18

replace aixc/struct/cpe => ../cpe

replace aixc/struct/pop => ../pop

replace aixc/struct/dvc => ../dvc

replace aixc/struct/opt => ../opt

replace aixc/conf => ../conf

replace aixc/cmd => ../cmd

require (
	github.com/olekukonko/tablewriter v0.0.5
	github.com/spf13/cobra v1.4.0 // direct
	gopkg.in/yaml.v3 v3.0.1 // direct
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
