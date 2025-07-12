// cmd/quantumnexus/main.go
package main

import (
"flag"
"log"
"os"

"quantumnexus/internal/quantumnexus"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := quantumnexus.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
