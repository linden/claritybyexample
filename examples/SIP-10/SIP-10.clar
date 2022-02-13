;; A simple SIP-10 token.

;; Defines the token with the identifier `example`.
(define-fungible-token example)
;; When a user lacks funds.
(define-constant ERR_INSUFFICIENT_BALANCE u1000)
;; When a user acts on behalf of another user.
(define-constant ERR_NOT_AUTHORIZED u1001)

;; Mint 100,000 uExample and send it to the 
;; contract deployer.
(ft-mint? example u100000 tx-sender)

;; User to user transfering.
(define-public (transfer 
;; Amount of tokens being sent.
		(amount uint) 
;; Sender of tokens.
		(from principal)
;; Receiver of tokens.
		(to principal) 
;; Optionally stores arbitrary data.
		(memo (optional (buff 34)))
	)
	
	(begin
;; Makes sure sender is the contract caller.
		(asserts! 
			(is-eq tx-sender from) 
			(err ERR_NOT_AUTHORIZED)
		)
		
;; Makes sure the sender has the funds.
		(asserts! 
			(>= (ft-get-balance example from) amount) 
			(err ERR_INSUFFICIENT_BALANCE)
		)
		
;; Send the tokens or panic.
		(unwrap-panic 
			(ft-transfer? example amount from to)
		)

		(ok true)
	)
)

;; The human-readable name of the token.
(define-read-only (get-name)
	(ok "Example")
)

;; The ticker of the token.
(define-read-only (get-symbol)
	(ok "EXP")
)

;; Where the decimal point is stored.
(define-read-only (get-decimals)
	(ok u4)
)

;; URL of the manifest (or none).
(define-public (get-token-uri)
;; Follows [this specification](https://bit.ly/35X3pex).
	(ok (some "https://example.com/manifest.json"))
)

;; The total supply of tokens.
(define-read-only (get-total-supply)
    (ok (ft-get-supply example))
)

;; The balance of a specific user.
(define-read-only (get-balance (address principal))
    (ok (ft-get-balance example address))
)