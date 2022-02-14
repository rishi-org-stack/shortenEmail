package context

import (
	"context"
	"time"
)

func ServiceContext(keyVal map[interface{}]interface{}) context.Context {
	parentContext := context.WithValue(context.Background(), "env", keyVal)
	ctx, _ := context.WithTimeout(parentContext, time.Second*10)
	// defer cancel()
	return ctx
}
