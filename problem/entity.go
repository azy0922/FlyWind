package problem

import (
	"log"
)

// 问题实体结构
type QuestEntity struct {
	Category        string
	Description     string
	GeneratorScript string
	SolverScript    string
}

// 随机生成题目用例,一般返回数组切片及其对应的string

func (q *QuestEntity) CateClass() string {
	return q.Category
}

func (q *QuestEntity) Question() string {
	return q.Description
}

// func (q *QuestEntity) GetGenScript() string {
// 	return q.GeneratorScript
// }

// func (q *QuestEntity) GetSolverScript() string {
// 	return q.SolverScript
// }

// 所有题目必须实现JsGenerator函数，生成随机测试用例
// 测试用例可能包括以下类型：
// 1. 单个或多个整型变量，int/int-int
// 2. 单个字符串或字符串数组,string/[]string
// 3. 整型数组/[]int
// 4. 私有结构/struct
// 5. 单个整型数组+单个整型变量[]int+int

// Test Example
func retJsGernerator() string {
	// 这里到时从数据库中取即可
	return `
	function JsGenerator() {
		// 这里是自定义的自写函数
		return Rand.GenRandTest();
		//return Case.GenRandSliceString(8, 20, 6);
	}`
}

// 随机用例生成函数, 加载题目的用例生成函数JsGenerator
func (q *QuestEntity) Generator() []interface{} {

	if _, err := VM.Run(q.GeneratorScript); err != nil {
		log.Panicln(err)
	}

	if main, err := VM.Get("JsGenerator"); err != nil || !main.IsFunction() {
		log.Panicln(err)
	}

	value, err := VM.Call("JsGenerator", nil, nil)
	if err != nil {
		log.Panicln(err)
	}
	result, err := value.Export()
	if err != nil {
		log.Panicln(err)
	}
	return result.([]interface{})
}

// 题目的解法函数
// 1. 用switch解析参数的类型
// 2. 加载题目的解法js脚本
// 3. 执行JsSolver函数获取结果
// 4. 结果返回golang

func (q *QuestEntity) GoSolver(params ...interface{}) interface{} {

	_, err := VM.Run(q.SolverScript)
	if err != nil {
		log.Panicln(err)
	}

	if main, err := VM.Get("JsSolver"); err != nil || !main.IsFunction() {
		log.Panicln(err)
	}

	// 注意这里的参数可以是任何类型的
	value, err := VM.Call("JsSolver", nil, params...)
	if err != nil {
		log.Panicln(err)
	}
	//fmt.Println(value)

	v, _ := value.Export()
	return v
}

// Test Example.
// 所有添加题目均要实现JsSolver入口函数
func retJsSolver() string {
	// 这里到时从数据库中取即可
	// 如何处理传进来的struct呢？
	return `
	function JsSolver(params) {
		var profit = 0;
  		for (var i = 0; i < params.length - 1; i++) {
   	 	if (params[i + 1] > params[i]) profit += params[i + 1] - params[i];
		  }
		//console.log(typeof profit)
		return profit*100;
		
		  //console.log(params.Matcher)
		  //console.log(params.Target)
		  //return ["123","456","789"]
	}`
}
