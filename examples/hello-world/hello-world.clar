;; Our first smart-contract will return the classic
;; "Hello World" message.

;; Define a public function called "hello-world".
(define-public (hello-world)
;; Returns "Hello World" with a status `ok`.
	(ok "Hello World")
)