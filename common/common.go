package common

import (
	"context"
)

var BackCtx, BackCtxCancel = context.WithCancel(context.Background());