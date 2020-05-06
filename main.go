// Generates the requested number of credit card numbers.
// Pass in the desired count using the -count flag.
// Defaults to count=10 if not specified.
// Prints a json representation of the data that can be
// sent to the Vault Transform Secret Engine as a batch_input.
// Reference:
// https://www.vaultproject.io/docs/secrets/transform
// https://www.vaultproject.io/api/secret/transform#batch_input

package main

import (
  "flag"
  "fmt"
  "os"
  "github.com/brianvoe/gofakeit"
)

func usage() {
  fmt.Fprintf(os.Stderr, "Usage: %s\n", os.Args[0])
  flag.PrintDefaults()
}

func main() {

  // Parse command line options with default values defined below
	flag.Usage = usage
  var count = flag.Int("count", 10, "Number of credit card numbers to generate")
  flag.Parse()

  fmt.Printf("{ \n  \"batch_input\": [\n")

  gofakeit.Seed(0)

  for i := 1; i <= *count; i++ {
    cc := gofakeit.CreditCardNumber()
    fmt.Printf("    { \"value\": \"%d\"}", cc)
    if i != *count {
      fmt.Printf(",")
    }
    fmt.Printf("\n")
  }

  fmt.Printf("  ]\n}")
}
