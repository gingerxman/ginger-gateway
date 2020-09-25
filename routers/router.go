package routers

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/eel/handler/rest/console"
	"github.com/gingerxman/eel/handler/rest/op"
	"github.com/gingerxman/ginger-gateway/rest/dev"
	"github.com/gingerxman/ginger-gateway/rest/resource"
)

func init() {
	eel.RegisterResource(&console.Console{})
	eel.RegisterResource(&op.Health{})
	
	/*
	 material
	*/
	eel.RegisterResource(&resource.Image{})

	/*
	 dev
	 */
	eel.RegisterResource(&dev.BDDReset{})
}