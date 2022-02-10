# To run the smart-contract, put the code in 
# `hello-world.clar` and pipe it into the `clarity-repl`.
$ cat hello-world.clar | clarity-repl
clarity-repl v0.21.0
Enter "::help" for usage hints.
Connected to a transient in-memory database.

Events emitted {
  contract_event: {
    contract_identifier: "ST000.contract-2",
    topic: "print",
    value: '"hello world"'
  },
  type: "contract_event"
}
"hello world"

# Now that we can run a basic Clarity
# smart-contract, let's learn more about the language.
