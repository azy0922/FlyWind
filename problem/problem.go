package problem

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/azy0922/flywind/model"
	"github.com/robertkrimen/otto"
)

// 一个题目由分类、描述、解决算法组成，定义一个接口
type Problem interface {
	CateClass() string
	Question() string
	Generator() []interface{}
	GoSolver(...interface{}) interface{}
}

type ProblemRet struct {
	Category string
	Question string
	Answer   interface{}
}

// 题目仓库
type ProblemDepos struct {
	Prob []Problem
}

// TODO: 使用map实现
var depository ProblemDepos
var VM *otto.Otto
var ChanRebuild = make(chan bool)

// 注册题目
func RegisterQuest(p Problem) {
	depository.Prob = append(depository.Prob, p)
}

// TODO:使用字典实现
func RegQuest(p Problem) {
	// Depository[p.CateClass()] = p
}

func init() {
	// 以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
	VM = otto.New()
	RC := RandCase{}

	VM.Set("Rand", &RC)

	// 从数据库中取题并注册,构造题目仓库结构
	// 建立一个信号机制，当有web端有新题加入时，从数据库中读新题进内存
	go func() {
		for {
			// TODO: 改为Select通道模式，不同的通道触发修改map
			var questions []QuestEntity
			depository.Prob = nil
			err := model.DB.Table("problems").
				Select("category, description, generator_script, solver_script").
				Scan(&questions).Error
			if err != nil {
				log.Panicln("从数据库加载题目失败:", err.Error())
			}

			if len(questions) == 0 {
				log.Println("题库中没有试题，请在web端添加试题！")
				<-ChanRebuild
				continue
			}

			for _, q := range questions {
				RegisterQuest(&QuestEntity{q.Category,
					q.Description,
					q.GeneratorScript,
					q.SolverScript})
			}
			<-ChanRebuild
		}
	}()
}

// 随机生成一道题目，返回给用户
func Generate(category string) *ProblemRet {
	// step0: 从题目仓库中随机抽取一个类型题目
	var ans interface{}
	n := rand.Intn(len(depository.Prob))

	for i, v := range depository.Prob {
		if v.CateClass() == category {
			n = i
		}
	}
	// step1: 生成题目用例和输入参数
	problem := depository.Prob[n]
	args := problem.Generator()

	if args == nil {
		return nil
	}

	// step2: 代入题目解法,计算答案
	solution := problem.GoSolver(args...)

	if solution == nil {
		return nil
	}

	// step3: 构造完整题目,返回前端
	switch solution.(type) {
	case int:
		ans = strconv.Itoa(solution.(int))
	case int64:
		ans = strconv.Itoa(int(solution.(int64)))
	case float64:
		ans = strconv.FormatFloat(solution.(float64), 'f', -1, 64)
	case string:
		ans = solution.(string)
	case []string:
		ans = solution.([]string)
	default:

	}

	// 将输入参数转成Sprintf参数字符串
	argConv := ConvertArgs(args)

	prob := &ProblemRet{
		Category: problem.CateClass(),
		Question: fmt.Sprintf(problem.Question(), argConv...),
		Answer:   ans,
	}
	return prob
}
