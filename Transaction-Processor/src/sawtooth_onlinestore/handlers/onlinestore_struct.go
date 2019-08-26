/**
 * Author：Robert WYQ
 * Date：2019-7-22 ——  2019-8-31
 * Description: 根据设计需求设计数据模型，设计一个handler类型，定义其数据结构，注意这部分的类型要与cli实现的功能及输入的参数相关联
 * 核心是context和request，context包含当前state的信息，request包含操作的信息
 * ------------------------------------------------------------------------------
 */

package handlers

import "github.com/hyperledger/sawtooth-sdk-go/processor"

type OnlinestoreHandler struct {
	// stores the information about current state
	context   *processor.Context
	operation string
	amount    int
	goods     string
	AddressFrom  string
	AddressTo    string
}

func (self OnlinestoreHandler) getOperation() string {
	return self.operation
}

func (self OnlinestoreHandler) getAmount() int {
	return self.amount
}

func (self OnlinestoreHandler) getAddressFrom() string {
	return self.AddressFrom
}

func (self OnlinestoreHandler) getAddressTo() string {
	return self.AddressTo
}

func (self OnlinestoreHandler) getContext() *processor.Context {
	return self.context
}

func (self OnlinestoreHandler) getGoods() string {
	return self.goods
}