module client

go 1.20

require hello v0.0.0

replace (
	hello => ../hello
	main => ../main
)
