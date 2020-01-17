package strategy

// ”策略模式就是用来封装算法的，但在实践中，我们发现可以用它来封装几乎任何类型的规则，
//只要在分析过程中听到需要在不同时间应用不同的业务规则，就可以考虑使用策略模式处理这种变化的可能性(by程杰. 大话设计模式)
//优化代码switch判断，要用到反射
// 实现此接口，则为一个策略
type IStrategy interface {
	do(int, int) int
}

// 具体策略1实现：加
type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

// 具体策略2实现：减
type reduce struct{}

func (*reduce) do(a, b int) int {
	return a - b
}

// 具体策略的执行者，封装了策略接口
type Operator struct {
	strategy IStrategy
}

// 设置策略-基础版
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 设置策略-策略与简单工厂模式结合
var (
	addIStrategy    *add
	reduceIStrategy *reduce
)

//客户端只用传入想要策略的名称即可，不用暴露具体的策略接口,进一步降低耦合
func (operator *Operator) setStrategy2(strategyType string) {
	switch strategyType {
	case "add":
		operator.strategy = addIStrategy
	case "reduce":
		operator.strategy = reduceIStrategy
	}
}

// 调用策略中的方法
func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}
