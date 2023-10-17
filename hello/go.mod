module example/hello

go 1.21.1

require example/greetings v0.0.0-00010101000000-000000000000

replace example.com/greetings => ../greetings

replace example/greetings => ../greetings
