package cobra

import (
	"github.com/gohade/hade/framework/contract"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func (c *Command) AddDistributedCronCommand(serviceName , spec string, cmd *Command, holdTime time.Duration)  {
	root := c.Root()

	//初始化cron
	if root.Cron == nil {
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor)))
		root.CronSpecs = []CronSpec{}
	}

	// cron 命令的注释，这里注意Type为distribute-cron,ServiceName需要填写
	appService := root.GetContainer().MustMake(contract.AppKey).(contract.App)
	distributeService := root.GetContainer().MustMake(contract.DistributedKey).(contract.Distributed)
	appID := appService.AppID()

	// 复制要执行的command为cronCmd，并且设置为rootCmd
	var cronCmd Command
	ctx := root.Context()
	cronCmd = *cmd
	cronCmd.args = []string{}
	cronCmd.SetParentNull()

	// cron增加匿名函数
	root.Cron.AddFunc(spec, func() {
		// 防止panic
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		//节点进行选举，返回选举结果
		selectAppID, err := distributeService.Select(serviceName, appID, holdTime)
		if err != nil {
			return
		}

		if selectAppID != appID {
			return
		}

		// 如果已经被选择到了，执行这个定时任务
		err = cronCmd.ExecuteContext(ctx)
		if err != nil {
			log.Println(err)
		}
	})
}