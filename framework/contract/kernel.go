package contract

import "net/http"

const KernelKey = "hade:kernel"

type Kernel interface {
	HttpEngine() http.Handler
}