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

	nodes, _ := jenkins.GetAllNodes(ctx)

	for _, node := range nodes {
		// Fetch Node Data
		node.Poll(ctx)
		if online, _ := node.IsOnline(ctx); online {
			fmt.Println(node.Raw.DisplayName + " Node is Online")
		}
	}

	tasks, _ := jenkins.GetQueue(ctx)

	for _, task := range tasks.Raw.Items {
		fmt.Println(task.Why)
	}
}
