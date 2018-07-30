# faker
造数据

文件|功能
-|-
faker.go|定义母体接口
util.go|工具函数
composite.go|组合母体，能够控制各个子母体频率
constant.go|产生固定值的母体
int.go|产生某个区间内数值的母体
hex.go|int.go的变体，输出不一样
string.go|基于基础字库，随机产生固定长度的字串，或者产生随机长度的字串
boolean.go|本质是组合母体，组合母体结合固定值母体的一种实现



