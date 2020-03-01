package engine

type ConcurrentEngine struct {
	// schedule 处理器
	Scheduler Scheduler
	// 开启work的数量（为什么需要数量）
	WorkerCount int

	ItemChan chan interface{}
}

// schedule处理器接口 定义
type Scheduler interface {
	// 处理器接受导的 请求
	Submit(Request)
	//
	//ConfigureMasterWorkerChan(chan Request)

	WorkerChan() chan Request

	ReadyNotifyer

	Run()
}

type ReadyNotifyer interface {
	WorkerReady(chan Request)
}

/**
@desc 引擎run入口
第一种实现方式 schedule实现I：所有worker公用一个输入schedule
	1 接收到请求处理：
		schedule是什么东西，能干什么？
		schedule怎么来的？
		把请求交给schedule，怎么交？
	2 接收到的请求，交给schedule：
		schedule接收到的请求，怎么处理？处理完再处理下一个
		实现一个channel，接受worker的结果
		循环 开启goroutine 交给worker
		worker的channel结果，循环给schedule，不断处理
	3 开启 worker 的 goroutine
		下载页面内容
		解析内容
		返回结果
第二种实现方式 schedule实现II：
	1 接收到请求处理：
		schedule是什么东西，能干什么？
		schedule怎么来的？
		把请求交给schedule，怎么交？
	2 接收到的请求，交给schedule：
		schedule接收到的请求，怎么处理？开一个协程goroutine，不断的接受请求
		实现一个channel，接受worker的结果
		循环 开启goroutine 交给worker
		worker的channel结果，循环给schedule，不断处理，接受请求goroutine处理
	3 开启 worker 的 goroutine
		下载页面内容
		解析内容
		返回结果
 */
func (e *ConcurrentEngine) Run(seeds ...Request)  {
	// 请求channel
	//in := make(chan Request)

	// 响应中的解析结果 channel
	out := make(chan ParseResult)

	// 干什么用？？？
	//e.Scheduler.ConfigureMasterWorkerChan(in)

	e.Scheduler.Run()

	// 开启worker，处理请求channel，并返回解析结果channel
	for i := 0;i < e.WorkerCount; i++ {
		//createWorker(in, out)
		//createWorker(out, e.Scheduler)
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 对接收到的请求不断交给schedule处理
	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	// 对解析结果channel，不断的接受 且 不断的 重新提交给schedule来循环处理
	//itemCount := 0
	for {
		result := <- out

		for _, item := range result.Items {

			go func() {
				e.ItemChan <- item
			}()
			//if _,ok := item.(model.Profile);ok {
			//	fmt.Printf("Got profile $%d: %v", itemCount, item)
			//	itemCount++
			//}

			//fmt.Printf("Got profile $%d: %v", itemCount, item)
			//itemCount++

			//save(item)

		}

		for _,request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// goroutine里为什么需要for 循环？？？
// goroutine里有channel，那么channel在没有时需要等待，那么为了开启goroutine不需要等到便开启goroutine

// 开启goroutine 并发处理： 对request channel处理，并返回解析结果 channel
//func createWorker(in chan Request, out chan ParseResult)  {
//	// 开启goroutine
//	go func() {
//		// for循环处理channel
//		for {
//			// tell scheduler i'm ready
//			request := <- in
//
//			// 真正下载页面，解析结果返回
//			result, err := work(request)
//			if err != nil {
//				continue
//			}
//
//			out <- result
//		}
//	}()
//}

// 队列版
//func createWorker(in chan Request, out chan ParseResult, s Scheduler)  {
	func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifyer)  {
	//in := make(chan Request) // 自己造

	// 开启goroutine
	go func() {
		// for循环处理channel
		for {
			// tell scheduler i'm ready
			//s.WorkerReady(in)
			ready.WorkerReady(in)

			request := <- in

			// 真正下载页面，解析结果返回
			result, err := work(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}