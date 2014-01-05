package main

import (
    "github.com/hoisie/web"
)

func hello(ctx *web.Context, val string) { 
    ctx.WriteString("hello " + val)
} 

func main() {
    web.Get("/(.*)", hello)
    web.Run("0.0.0.0:9999")
}
