# Before we can run the smart-contract, we need to set our
# [Clarinet](https://github.com/hirosystems/clarinet) workspace.

# First create our workspace.
$ clarinet new example
Creating directory example
Created file example/contracts
Created file example/settings
Created file example/tests
Created file example/Clarinet.toml
Created file example/settings/Mainnet.toml
Created file example/settings/Testnet.toml
Created file example/settings/Devnet.toml
Created file example/.vscode
Created file example/.vscode/settings.json
Created file example/.vscode/tasks.json
Created file example/.gitignore

# Then create our contract.
$ clarinet contract new hello-world
Updated Clarinet.toml with contract hello-world

# Put the source code in the contract (located at 
# `contracts/hello-world.clar`).

# Now we have our [Clarinet](https://github.com/hirosystems/clarinet) workspace set up 
# we can load our contract into the [REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop#Overview).
$ clarinet poke
clarity-repl v0.21.0
Enter "::help" for usage hints.
Connected to a transient in-memory database.
# List of contracts and callable functions.
Contracts
+---------------------+------------------+
| Contract identifier | Public functions |
+---------------------+------------------+
| address.hello-world | (hello-world)    |
+---------------------+------------------+

# Wallets loaded in from `settings/Devnet.toml`, 
# balances listed in uSTX.
Initialized balances
+--------------------+-----------------+
| Address            | STX             |
+--------------------+-----------------+
| address (deployer) | 100000000000000 |
+--------------------+-----------------+
| address (wallet_1) | 100000000000000 |
+--------------------+-----------------+
| address (wallet_2) | 100000000000000 |
+--------------------+-----------------+
| address (wallet_3) | 100000000000000 |
+--------------------+-----------------+
| address (wallet_4) | 100000000000000 |
+--------------------+-----------------+
| address (wallet_5) | 100000000000000 |
+--------------------+-----------------+
| address (wallet_6) | 100000000000000 |
+--------------------+-----------------+
| address (wallet_8) | 100000000000000 |
+--------------------+-----------------+
| address (wallet_7) | 100000000000000 |
+--------------------+-----------------+
| address (wallet_9) | 100000000000000 |
+--------------------+-----------------+

# Now that we're in the [REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop#Overview), we can call the 
# public functions/contracts listed above. 

# For example here's how we can call the `hello-world` 
# function from our `.hello-world` contract.
$ (contract-call? .hello-world hello-world)
(ok "Hello World")

# The process for setting up our [Clarinet](https://github.com/hirosystems/clarinet) workspace
# will be the same for the rest of our smart-contracts.
