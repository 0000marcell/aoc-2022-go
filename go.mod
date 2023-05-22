module example/hello

go 1.19

require example.com/set v0.0.0-00010101000000-000000000000

replace example.com/set => ./set

replace marcell.com/queue => ../queue
