package mapreduce

import (
	"fmt"
	"sync"
)

//
// schedule() starts and waits for all tasks in the given phase (mapPhase
// or reducePhase). the mapFiles argument holds the names of the files that
// are the inputs to the map phase, one per map task. nReduce is the
// number of reduce tasks. the registerChan argument yields a stream
// of registered workers; each item is the worker's RPC address,
// suitable for passing to call(). registerChan will yield all
// existing registered workers (if any) and new ones as they register.


//schedule() tells a worker to execute a task by sending a Worker.DoTask RPC to the worker.
//	This RPC's arguments are defined by DoTaskArgs in mapreduce/common_rpc.go.
//	The File element is only used by Map tasks, and is the name of the file to read;
//  schedule() can find these file names in mapFiles.
//
//Use the call() function in mapreduce/common_rpc.go
//to send an RPC to a worker. The first argument is the the worker's address
//, as read from registerChan.
//	The second argument should be "Worker.DoTask".
//	The third argument should be the DoTaskArgs structure, and the last argument should be nil.
//
func schedule(jobName string, mapFiles []string, nReduce int, phase jobPhase, registerChan chan string) {
	var ntasks int
	var n_other int // number of inputs (for reduce) or outputs (for map)

	// check是map阶段还是reduce阶段
	switch phase {
	case mapPhase:
		ntasks = len(mapFiles)
		n_other = nReduce
	case reducePhase:
		ntasks = nReduce
		n_other = len(mapFiles)
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, n_other)

	// All ntasks tasks have to be scheduled on workers. Once all tasks
	// have completed successfully, schedule() should return.
	//
	// Your code here (Part III, Part IV).
	//

	// 给每个worker发rpc要求执行任务
	// 因为task远比worker多 所以我们需要先让每个worker工作，再看哪个完成了派发新的任务
	var wg sync.WaitGroup

	for i := 0; i < ntasks; i++ {

		taskArg := DoTaskArgs{}
		taskArg.JobName = jobName
		taskArg.Phase = phase
		taskArg.TaskNumber = i
		taskArg.NumOtherPhase = n_other

		if phase == mapPhase {
			taskArg.File = mapFiles[i]
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			// 堵塞 获取worker
			worker := <- registerChan

			for {
				status := call(worker, "Worker.DoTask", taskArg, nil)
				if status {
					// 如果执行成功
					break
				} else {
					// 尝试获取另一个能用的在线worker
					worker = <- registerChan
				}
			}

			// 必须单独开一个协程来还  因为可能会出现两个worker同时写registerChan的情况，
			// 此时就死锁了
			go func() {
				// 放回该worker
				registerChan <- worker
			}()
		}()
	}

	// 等待所有on the fly任务完成
	wg.Wait()
	fmt.Printf("Schedule: %v done\n", phase)
}
