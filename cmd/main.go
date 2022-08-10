package main

import (
	"github.com/kimxuanhong/go-campaign-no-02/pkg/router"
)

func main() {
	Router := router.RouterImpl{}

	Router.Start()
}
