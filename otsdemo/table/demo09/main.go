package main


//局部事务
//todo 使用局部事务功能，创建数据范围在一个分区键值内的局部事务。对局部事务中的数据进行读写操作后，可以根据实际提交或者丢弃局部事务。
//todo 目前局部事务功能处于邀测中，默认关闭。如果需要使用该功能，请提交工单进行申请或者加入钉钉群23307953（表格存储技术交流群-2）进行咨询。
//todo 使用局部事务可以指定某个分区键值内的操作是原子的，对分区键值内的数据进行的操作要么全部成功要么全部失败，并且所提供的隔离级别为串行化。

func main() {

}
