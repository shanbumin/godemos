package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

/*

todo 条件更新:只有满足条件时，才能对数据表中的数据进行更新；当不满足条件时，更新失败。
todo 前提条件:已创建数据表并写入数据。
todo 使用方法:
      1.在通过PutRow、UpdateRow、DeleteRow或BatchWriteRow接口更新数据时，可以使用条件更新检查行存在性条件和列条件，只有满足条件时才能更新成功。
      2.条件更新包括行存在性条件和列条件。
        a.行存在性条件：包括IGNORE、EXPECT_EXIST和EXPECT_NOT_EXIST，分别代表忽略、期望存在和期望不存在。
          对数据表进行更改操作时，系统会先检查行存在性条件，如果不满足行存在性条件，则更改失败并给用户报错。
        b.列条件：包括SingleColumnCondition和CompositeColumnCondition，是基于某一列或者某些列的列值进行条件判断。
          SingleColumnCondition支持一列（可以是主键列）和一个常量比较。不支持两列或者两个常量相比较。
          CompositeColumnCondition的内节点为逻辑运算，子条件可以是SingleColumnCondition或CompositeColumnCondition。
      3.条件更新可以实现乐观锁功能，即在更新某行时，先获取某列的值，假设为列A，值为1，然后设置条件列A＝1，更新行使列A＝2。如果更新失败，表示有其他客户端已成功更新该行。



updateRowChange.SetCondition( tablestore.RowExistenceExpectation_IGNORE)    :忽略
updateRowChange.SetCondition( tablestore.RowExistenceExpectation_EXPECT_EXIST)   :期望存在
updateRowChange.SetCondition( tablestore.RowExistenceExpectation_EXPECT_NOT_EXIST)   :期望不存在

*/

func main() {

	sample.CreateTableConditionSample(start.Client,sample.TableConditionName)
	sample.PutRowWithConditionSample(start.Client,sample.TableConditionName)
	sample.ConditionRowUpdateSample(start.Client,sample.TableConditionName)
	sample.ConditionColUpdateSample(start.Client,sample.TableConditionName)

}
