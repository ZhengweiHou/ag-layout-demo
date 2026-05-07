package internal

import (
	"ag-layout-demo/internal/svcgen"

	"gitlab.allinfinance.com/aifgo/ag-core/ag/ag_common/agmetadata"
	"gitlab.allinfinance.com/aifgo/ag-core/contribute/agdb"
)

func init() {
	// 初始化headdev
	agmetadata.RegMdKey("ORG")
	agmetadata.RegMdKey("Tradeid")

	// 注册事务
	// agdb.TransactionRegistry.RegTxForCall(svcgen.StudentServiceCreateStudentCallInfo)
	svcgen.StudentServiceCreateStudentCallInfo.AddTag(agdb.TransactionTag, true)
}
