package main

import "net/http"

func prepareDefaultContext(request *http.Request, context map[string]interface{}) {
	context["namespaces"] = retrieveAvailableNamespaces()
	context["version"] = "0.0.3-dev"
}
