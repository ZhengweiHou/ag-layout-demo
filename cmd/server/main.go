package main

import (
	"ag-core/ag/ag_app"
	"ag-layout-demo/internal"

	"ag-core/fxs"

	"ag-core/ag/ag_service"
	"ag-core/contribute/agdb"
	"ag-core/contribute/agdb/gormdb"
	hserver "ag-core/contribute/aghertz/server"
	kserver "ag-core/contribute/agkitex/server"
	"ag-core/contribute/agnacos"

	"go.uber.org/fx"

	// _ "ag-layout-demo/api/logs/logserver"
	_ "go.uber.org/automaxprocs"
)

func main() {
	var fxopts []fx.Option

	fxopts = append(
		fxopts,
		mainFx,
		fx.Invoke(func(s *ag_app.App) {}),
	)

	fxapp := fx.New(
		fxopts...,
	)

	fxapp.Run()
}

var mainFx = fx.Module("main",
	/** conf **/
	// 初始化配置
	fxs.FxAgConfModule,

	// nacosconf
	agnacos.FxNacosConfigMode,
	agnacos.FxEnableNacosRemoteConfigMode,

	// 根APP
	fxs.FxAppMode,

	// http服务
	hserver.FxAgHertzServerModule,
	// grpc服务
	kserver.FxKitexServerBaseModule,

	// 日志模块
	fxs.FxAgSlogMode,
	fxs.FxAgSlogZapMode,
	// agslogzap.FxAgSlogZapMode,

	// 服务基础模块
	ag_service.FxAgServiceMode,

	// 数据库模块
	agdb.FxAgDbModule,
	gormdb.FxAicGromdbModule,

	/** 项目内部组件 **/

	// 内部组件
	internal.FxInternalModule,
)
