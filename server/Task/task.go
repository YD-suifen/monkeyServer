package Task

import (
	"fmt"
	"monkeyServer/server/Calculation"
	"monkeyServer/utils"
)

func Task()  {
	fmt.Println("Task init")
	Crontab := utils.TasksInit()

	spec := utils.Config.Tasks.Trend

	Crontab.AddFunc(spec, func() {
		fmt.Println("Task start Algorithm")
		Calculation.Algorithm()
	})
	Crontab.Start()
}
