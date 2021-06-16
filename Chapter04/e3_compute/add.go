package compute

type AddOperator interface {
	/*
		算术相加
	*/
	Add() interface{}
}

// 如果函数、类型、变量、常量位于不同的包下，则需要将它们名称的
// 首字母大写，表示它们是公开可访问的，对于结构体下的字段，如果
// 想要在包外访问，还需要将字段变量名首字母大写
type IntParams struct {
	P1 int // 加数
	P2 int // 被加数
}

func (params *IntParams) Add() interface{} {
	return params.P1 + params.P2
}

// 通过 package compute 将 add.go 归属于 compute 包下；定义了
// AddOperator 接口，接口中只有一个方法 Add，用于进行算术相加
// 操作；同时定义了 IntParams 结构体用于实现 AddOperator.Add，
// 实现方法是将 IntParams 内包含的两个整数进行相加并返回结果。
