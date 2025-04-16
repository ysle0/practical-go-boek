module github.com/ysle0/chpt6/middleware-chaining

go 1.24

replace (
	github.com/ysle0/chpt6/handlerWrap => ./handlerWrap
	github.com/ysle0/chpt6/middleware => ./middleware
)

require (
	github.com/ysle0/chpt6/handlerWrap v0.0.0-00010101000000-000000000000
	github.com/ysle0/chpt6/middleware v0.0.0-00010101000000-000000000000
)
