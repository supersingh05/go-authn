package config

import (
	"flag"
	"os"
)

type Config struct {
	Addr      string
	StaticDir string
	Dsn       string
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
	flag.StringVar(&c.Addr, "dsn", "", "HTTP network address")
	flag.Parse()
}

func (c *Config) parseEnv() {
	if isConfigBlank(c.StaticDir) {
		c.StaticDir = os.Getenv("QS_STATICDIR")
	}
	if isConfigBlank(c.Addr) {
		c.Addr = os.Getenv("QS_ADDR")
	}
	if isConfigBlank(c.Dsn) {
		c.Dsn = os.Getenv("QS_DSN")
	}
}

func (c *Config) defaults() {
	if isConfigBlank(c.StaticDir) {
		c.StaticDir = "./ui/static"
	}
	if isConfigBlank(c.Addr) {
		c.Addr = ":4000"
	}
	if isConfigBlank(c.Dsn) {
		c.Dsn = "web:pass@/snippetbox?parseTime=true"
	}

}

func isConfigBlank(s string) bool {
	return len(s) == 0
}
