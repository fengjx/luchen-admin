package transport

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/luchen-admin/transport/http"
)

func Start(_ context.Context) {
	luchen.Start(http.GetServer())
}

func Stop(_ context.Context) {
	luchen.Stop()
}
