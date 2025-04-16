module github.com/ysle0/chpt6/middleware-chaining

go 1.24

replace (
	github.com/ysle0/chpt6/middleware-chaining/handlerWrap => ./handlerWrap
	github.com/ysle0/chpt6/middleware-chaining/middleware => ./middleware
	github.com/ysle0/chpt6/middleware-chaining/handler => ./handler
)

require (
	github.com/ysle0/chpt6/middleware-chaining/handler v0.0.0-00010101000000-000000000000
	github.com/ysle0/chpt6/middleware-chaining/handlerWrap v0.0.0-00010101000000-000000000000
	github.com/ysle0/chpt6/middleware-chaining/middleware v0.0.0-00010101000000-000000000000
)
