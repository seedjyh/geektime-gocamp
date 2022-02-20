package main

type web struct {
	Address string `toml:"address"`
}

type config struct {
	Web web `toml:"web"`
}
