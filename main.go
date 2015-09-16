package main

import (
	"github.com/mitchellh/packer/packer/plugin"
  "github.com/artn/packer-post-processor-ovaexport/post-processor/ovaexport"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(ovaexport.PostProcessor))
	server.Serve()
}
