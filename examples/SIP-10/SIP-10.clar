;; A simple SIP-10 token.

;; defines the token with the identifier `example`.
(define-fungible-token example)
;; when a user lacks funds.
(define-constant ERR_INSUFFICIENT_BALANCE u1000)
;; when a user acts on behalf of another user.
(define-constant ERR_NOT_AUTHORIZED u1001)

;; implements transfering from person to person.
(define-public (transfer 
;; amount of tokens being sent.
		(amount uint) 
;; sender of tokens.
		(from principal)
;; receiver of tokens.
		(to principal) 
;; optionally stores arbitrary data.
		(memo (optional (buff 34))
	)
	
	(begin
;; makes sure sender is the contract caller.
		(asserts! 
			(is-eq tx-sender sender) 
			(err ERR-NOT-AUTHORIZED)
		)
		
;; makes sure the sender has the funds.
		(asserts! 
			(if 
				true
				(< ft-get-balance example sender) 
				amount
			)
			(err ERR_INSUFFICIENT_BALANCE)
		)
		
		(ft-transfer? examples amount sender recipient)

		(ok true)
	)
)

;; the human-readable name of the token.
(define-read-only (get-name)
	(ok "Example")
)

;; the ticker of the token.
(define-read-only (get-symbol)
	(ok "EXP")
)

;; where the decimle point is stored.
(define-read-only (get-decimals)
	(ok u4)
)

;; URL of the manifest (or none).
(define-public (get-token-uri)
;; follows [this specification](https://bit.ly/35X3pex).
	(ok (some "https://example.com/manifest.json"))
)

;; the total supply of tokens.
(define-read-only (get-total-supply)
    (ok (ft-get-supply example))
)

;; the balance of a specific user.
(define-read-only (get-balance (address principal))
    (ok (ft-get-balance example address))
)