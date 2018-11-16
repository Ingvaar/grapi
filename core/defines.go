package core

import "net/http"

// Handlers : define the handlers map type
type Handlers map[string]http.HandlerFunc
