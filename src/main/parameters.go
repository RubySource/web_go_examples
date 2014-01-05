package main

import (
    "github.com/hoisie/web"

    "fmt"
)

func logName(ctx *web.Context, val string) { 
  name := ctx.Params["name"]
  fmt.Println(name)
  ctx.WriteString("Hi, " + name)
} 

func main() {
    web.Get("/log_name?(.*)", logName)
    web.Run("0.0.0.0:9999")
}