package cobra

import (
	"fmt"
	"github.com/gohade/hade/framework"
	"github.com/robfig/cron/v3"
	"log"
)

// SetContainer 设置服务容器
func (c *Command) SetContainer(container framework.Container) {
	fmt.Println("SetContainer", container)
	c.container = container
}

// GetContainer 获取服务容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}

// CronSpec 保存Cron命令的信息，用于展示
type CronSpec struct {
	Type        string
	Cmd         *Command
	Spec        string
	ServiceName string
}

func (c *Command) SetParentNull()  {
	c.parent = nil
}

func (c *Command) AddCronCommand(spec string, cmd *Command)  {
	// cron结构挂载在根Command上的
	root := c.Root()
	if root. Cron == nil {
		//初始化cron
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dow | cron.Descriptor)))
		root.CronSpecs = []CronSpec{}
	}
	//增加说明信息
	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Type:"normal-cron",
		Cmd: cmd,
		Spec:spec,
	})

	//制作一个rootCommand
	var cronCmd Command
	ctx := root.Context()
	cronCmd = *cmd
	cronCmd.args = []string{}
	cronCmd.SetParentNull()
	cronCmd.SetContainer(root.GetContainer())

	//增加调用函数
	root.Cron.AddFunc(spec, func() {
		// 如果后续的command出现panic,这里捕获
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		err := cronCmd.ExecuteContext(ctx)
		if err != nil {
			// 打印出err信息
			log.Println(err)
		}
	})
}
