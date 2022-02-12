# Create our contract.
$ clarinet contract new SIP-10
Updated Clarinet.toml with contract SIP-10

# Now add the source to the contract 
# located at `contracts/SIP-10.clar`.

# Enter the [REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop#Overview).
$ clarinet poke
clarity-repl v0.21.0
Enter "::help" for usage hints.
Connected to a transient in-memory database.
# List of contracts and callable functions.
Contracts
+---------------------+----------------------------------+
| Contract identifier | Public functions                 |
+---------------------+----------------------------------+
| address.SIP-10      | (get-balance(address principal)) |
|                     | (get-decimals)                   |
|                     | (get-name)                       |
|                     | (get-symbol)                     |
|                     | (get-token-uri)                  |
|                     | (get-total-supply)               |
|                     | (transfer                        |
|                     |     (amount uint)                |
|                     |     (from principal)             |
|                     |     (to principal)               |
|                     |     (memo (optional (buff 34)))) |
+---------------------+----------------------------------+

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

# Now we can call the public functions.

# Get the total supply of example.
$ (contract-call? .SIP-10 get-total-supply)
(ok u100000)

# Transfer 100 of our uExample to a random user.\
# *Note: Principales are prefixed by '*.
$ (contract-call? .SIP-10 transfer u100 tx-sender 'SP3TZ3N
$ Y4GB3E3Y1K1D40BHE07P20KMS4A8YC4QRJ none)
Events emitted
{
  type: "ft_transfer_event",
  ft_transfer_event: {
    asset_identifier: "address.SIP-10::example",
    sender: "address",
    recipient: "address",
    amount: "100"
  }
}
(ok true)