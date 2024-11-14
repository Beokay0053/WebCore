package main

import "web_core/framework"

func registerRoute(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
