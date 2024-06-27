module client

go 1.20


require (
	watch v0.0.0
	server v0.0.0
)

replace (
	watch => ../watch
	server => ../Server
)