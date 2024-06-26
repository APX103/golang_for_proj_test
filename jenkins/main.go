package main

import (
	"context"
	"fmt"

	"github.com/bndr/gojenkins"
)

func main() {
	ctx := context.Background()
	jenkins, err := gojenkins.CreateJenkins(nil, "https://ci.staging.openxlab.org.cn/", "apx103", "oooo").Init(ctx)
	if err != nil {
		fmt.Println(err)
	}
	// nodes, _ := jenkins.GetAllNodes(ctx)

	// for _, node := range nodes {
	// 	// Fetch Node Data
	// 	node.Poll(ctx)
	// 	if online, _ := node.IsOnline(ctx); online {
	// 		fmt.Println(node.Raw.DisplayName + " Node is Online")
	// 	}
	// }

	tasks, _ := jenkins.GetQueue(ctx)

	for _, task := range tasks.Raw.Items {
		fmt.Println(task.Why)
	}

	// jobs, err := jenkins.GetAllJobs(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// time_1 := time.Now()
	// jobNames := GetAllJobs(ctx, jenkins, jobs, "")
	// time_2 := time.Now()
	// // for _, job := range jobs {
	// // 	if job.Raw.Class == "com.cloudbees.hudson.plugins.folder.Folder" {
	// // 		fmt.Println(job.Raw.Name + " is folder")
	// // 		continue
	// // 	}
	// // 	fmt.Println(job.Raw.Name + " is job")
	// // }

	// fmt.Printf("===================%s==================\n", time_1.GoString())
	// fmt.Printf("===================%s==================\n", time_2.GoString())
	// for _, name := range jobNames {
	// 	fmt.Println(name)
	// }
	fmt.Println("===========")
	theJob, err := jenkins.GetJob(ctx, "release_test", "openmmlab_algo_test", "mmengine")
	if err != nil {
		fmt.Println(err)
	}
	pl, _ := theJob.GetParameters(ctx)
	for _, p := range pl {
		fmt.Print(p.Name + " : ")
		fmt.Println(p.Description)
	}
}

type JobStruct struct {
	Job     string
	Parents string
}

func GetAllJobs(ctx context.Context, jenkins *gojenkins.Jenkins, jobs []*gojenkins.Job, parents string) []*JobStruct {
	jobNames := []*JobStruct{}
	for _, job := range jobs {
		_parents := parents
		if job.Raw.Class == "com.cloudbees.hudson.plugins.folder.Folder" {
			if _parents == "" {
				_parents = job.Raw.Name
			} else {
				_parents += "/" + job.Raw.Name
			}
			subJobs, _ := job.GetInnerJobs(ctx)
			subJobNames := GetAllJobs(ctx, jenkins, subJobs, _parents)
			jobNames = append(jobNames, subJobNames...)
		} else {
			jobNames = append(jobNames, &JobStruct{
				Job:     job.GetName(),
				Parents: _parents,
			})
		}
	}
	return jobNames
}
