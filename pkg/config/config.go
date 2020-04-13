package config

import (
	"flag"
	"os"
)

type Config struct {
	Addr      string
	StaticDir string
}

/*
   Parse Config from command line or env.
   Prescedence is Commandline, then Env
   TODO: file
*/
func ParseConfig() Config {
	cfg := new(Config)
	cfg.parseFlags()
	cfg.parseEnv()
	cfg.defaults()
	return *cfg
}

func (c *Config) parseFlags() {
	flag.StringVar(&c.StaticDir, "static-dir", "", "Path to static assets")
	flag.StringVar(&c.Addr, "addr", "", "HTTP network address, default is :4000")
	flag.Parse()
}

func (c *Config) parseEnv() {
	if isConfigBlank(c.StaticDir) {
		c.StaticDir = os.Getenv("QS_STATICDIR")
	}
	if isConfigBlank(c.Addr) {
		c.StaticDir = os.Getenv("QS_ADDR")
	}
}

func (c *Config) defaults() {
	if isConfigBlank(c.StaticDir) {
		c.StaticDir = "./ui/static"
	}
	if isConfigBlank(c.Addr) {
		c.Addr = ":4000"
	}
}

func isConfigBlank(s string) bool {
	return len(s) == 0
}
