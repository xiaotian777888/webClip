package main

import (
	"flag"
	. "webClip/api"
)

func main() {
	var ip = flag.String("web", "./public", "please input you web root")
	var c = flag.String("crt", "./webclip-crt.pem", "please input you  crt path")
	var k = flag.String("key", "./webclip-key.pem", "please input you  key path")
	var u = flag.String("ca", "./webclip-ca.pem", "please input you  ca path")
	flag.Parse()
	Run(*ip, *c, *k, *u)
}
