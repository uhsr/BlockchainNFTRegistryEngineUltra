// cmd/blockchainnftregistryengineultra/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistryengineultra/internal/blockchainnftregistryengineultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistryengineultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
