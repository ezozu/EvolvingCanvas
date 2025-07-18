// cmd/evolvingcanvas/main.go
package main

import (
"flag"
"log"
"os"

"evolvingcanvas/internal/evolvingcanvas"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := evolvingcanvas.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
