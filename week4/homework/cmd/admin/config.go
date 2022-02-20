package main

type web struct {
	Address string `toml:"address"`
}

type mysql struct {
	DataSourceName string `toml:"dataSourceName"`
}

type config struct {
	Web   web   `toml:"web"`
	MySQL mysql `toml:"mysql"`
}
