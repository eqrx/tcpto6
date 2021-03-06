// Copyright (C) 2021 Alexander Sowitzki
//
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
// warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more
// details.
//
// You should have received a copy of the GNU Affero General Public License along with this program.
// If not, see <https://www.gnu.org/licenses/>.

// Package main provides an entry point to tcp4to6.Run.
package main

import (
	"context"
	stdlog "log"
	"os"
	"os/signal"

	"dev.eqrx.net/tcpto6"
	"github.com/go-logr/stdr"
	"golang.org/x/sys/unix"
)

func main() {
	log := stdr.New(stdlog.New(os.Stderr, "", 0))

	// Ensure that following signal is always canceled by putting
	// os.Exit call in the first defer. Error output is also put here
	// for brevity.
	var err error
	defer func() {
		if err != nil {
			log.Error(err, "program error")
			os.Exit(1)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
	defer cancel()

	err = tcpto6.Run(ctx, log)
}
