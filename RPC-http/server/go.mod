module main

go 1.20

require hello v0.0.0

replace (
	client => ../client
	hello => ../hello
)
