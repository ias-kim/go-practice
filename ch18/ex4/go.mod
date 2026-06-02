module ch18/ex4

go 1.26.3

require (
	github.com/tuckersGo/musthaveGo2/ch18/fedex v0.0.0
	github.com/tuckersGo/musthaveGo2/ch18/koreaPost v0.0.0
)

replace (
	github.com/tuckersGo/musthaveGo2/ch18/fedex => ../fedex
	github.com/tuckersGo/musthaveGo2/ch18/koreaPost => ../koreaPost
)
