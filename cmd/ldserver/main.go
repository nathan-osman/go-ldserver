package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/nathan-osman/go-ldserver/manager"
)

func main() {
	if err := func() error {
		if len(os.Args) != 2 {
			return fmt.Errorf("Usage: %s <filename>", path.Base(os.Args[0]))
		}
		c, err := LoadConfig(os.Args[1])
		if err != nil {
			return err
		}
		m, err := manager.NewManager(c.Manager)
		if err != nil {
			return err
		}
		defer m.Close()
		sigCh := make(chan os.Signal)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		return nil
	}(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
