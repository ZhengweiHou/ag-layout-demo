package internal

import (
	"ag-layout-demo/internal/adpgen"
	"ag-layout-demo/internal/svcgen"

	"go.uber.org/fx"
)

var FxInternalModule = fx.Module("fx-internal-module",
	svcgen.FxServiceWithProxyModule(),
	adpgen.FxAdapterModule(),
)
