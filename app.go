package main

import (
    . "adCloser/web"
)

func main() {
    router := Service()

    router.GET("/", Index)
    router.GET("/v1", Index)

    router.Run(":3000")
}
