// Copyright 2021-2023 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package value_scan

import (
	"bytes"
	"time"

	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/logutil"
	"github.com/matrixorigin/matrixone/pkg/pb/plan"
	plan2 "github.com/matrixorigin/matrixone/pkg/pb/plan"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec"
	"github.com/matrixorigin/matrixone/pkg/vm"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

const opName = "value_scan"

func (valueScan *ValueScan) String(buf *bytes.Buffer) {
	buf.WriteString(opName)
	buf.WriteString(": value_scan ")
}

func (valueScan *ValueScan) OpType() vm.OpType {
	return vm.ValueScan
}

func evalRowsetData(proc *process.Process, rowsetExpr []*plan.RowsetExpr, vec *vector.Vector, exprExecs []colexec.ExpressionExecutor,
) error {
	bats := []*batch.Batch{batch.EmptyForConstFoldBatch}
	for i, expr := range exprExecs {
		val, err := expr.Eval(proc, bats, nil)
		if err != nil {
			return err
		}
		if err := vec.Copy(val, int64(rowsetExpr[i].RowPos), 0, proc.Mp()); err != nil {
			return err
		}
	}
	return nil
}

func (valueScan *ValueScan) makeValueScanBatch(proc *process.Process) (err error) {
	var exprList []colexec.ExpressionExecutor

	if valueScan.RowsetData == nil { // select 1,2
		bat := batch.NewWithSize(1)
		bat.Vecs[0] = vector.NewConstNull(types.T_int64.ToType(), 1, proc.Mp())
		bat.SetRowCount(1)
		valueScan.Batchs = append(valueScan.Batchs, bat)
		valueScan.Batchs = append(valueScan.Batchs, nil)
		return nil
	}

	start := time.Now()
	if valueScan.ExprExecLists == nil {
		if err := valueScan.InitExprExecList(proc); err != nil {
			return err
		}
	}
	if t := time.Since(start); t > time.Second {
		logutil.Infof("666666 InitExprExecList cost %v", t)
	}

	// select * from (values row(1,1), row(2,2), row(3,3)) a;
	bat := valueScan.Batchs[0]

	start = time.Now()
	for i := 0; i < valueScan.ColCount; i++ {
		exprList = valueScan.ExprExecLists[i]
		if len(exprList) == 0 {
			continue
		}
		vec := bat.Vecs[i]
		if err := evalRowsetData(proc, valueScan.RowsetData.Cols[i].Data, vec, exprList); err != nil {
			return err
		}
	}
	if t := time.Since(start); t > time.Second {
		logutil.Infof("666666 evalRowsetData cost %v", t)
	}

	return nil
}

func (valueScan *ValueScan) InitExprExecList(proc *process.Process) error {
	exprExecLists := make([][]colexec.ExpressionExecutor, len(valueScan.RowsetData.Cols))
	for i, col := range valueScan.RowsetData.Cols {
		var exprExecList []colexec.ExpressionExecutor
		for _, data := range col.Data {
			exprExecutor, err := colexec.NewExpressionExecutor(proc, data.Expr)
			if err != nil {
				return err
			}
			exprExecList = append(exprExecList, exprExecutor)
		}
		exprExecLists[i] = exprExecList
	}

	valueScan.ExprExecLists = exprExecLists
	return nil
}

func (valueScan *ValueScan) Prepare(proc *process.Process) (err error) {
	if valueScan.OpAnalyzer == nil {
		valueScan.OpAnalyzer = process.NewAnalyzer(valueScan.GetIdx(), valueScan.IsFirst, valueScan.IsLast, "value_scan")
	} else {
		valueScan.OpAnalyzer.Reset()
	}

	err = valueScan.PrepareProjection(proc)
	if err != nil {
		return err
	}
	if valueScan.NodeType == plan2.Node_VALUE_SCAN {
		err = valueScan.makeValueScanBatch(proc)
		if err != nil {
			return err
		}
	}

	return
}

func (valueScan *ValueScan) Call(proc *process.Process) (vm.CallResult, error) {
	if err, isCancel := vm.CancelCheck(proc); isCancel {
		return vm.CancelResult, err
	}

	analyzer := valueScan.OpAnalyzer
	analyzer.Start()
	defer analyzer.Stop()

	result := vm.NewCallResult()

	if valueScan.ctr.idx < len(valueScan.Batchs) {
		result.Batch = valueScan.Batchs[valueScan.ctr.idx]
		valueScan.ctr.idx += 1
	}
	var err error
	result.Batch, err = valueScan.EvalProjection(result.Batch, proc)

	analyzer.Input(result.Batch)
	analyzer.Output(result.Batch)
	return result, err

}
