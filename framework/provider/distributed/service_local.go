package distributed

import (
	"errors"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// LocalDistributedService 代表hade框架的App实现
type LocalDistributedService struct {
	container framework.Container // 服务容器
}

func NewLocalDistributedService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	return &LocalDistributedService{container:container}, nil
}

func (s LocalDistributedService) Select(serviceName , appID string, holdTime time.Duration ) (selectAppID string, err error)  {
	appService := s.container.MustMake(contract.AppKey).(contract.App)
	runtimeFolder := appService.RuntimeFolder()
	lockFile := filepath.Join(runtimeFolder, "distribute_" + serviceName )

	// 打开文件锁
	lock, err := os.OpenFile(lockFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	// 尝试独占文件锁
	err = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	// 抢不到文件锁
	if err != nil {
		// 被选择的appid
		selectAppIDByt, err := ioutil.ReadAll(lock)
		if err != nil {
			return "", err
		}
		return string(selectAppIDByt), err
	}

	// 在一段时间内， 选举有效，其他节点在这段时间不能再进行抢占
	go func() {
		defer func() {
			// 释放文件锁
			syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
			// 释放文件
			lock.Close()
			// 删除文件锁对应的文件
			os.Remove(lockFile)
		}()
		timer := time.NewTimer(holdTime)
		// 等待计数器结束
		<-timer.C
	}()
	// 这里已经是抢占到了，将抢占到的appID写入文件
	if _, err := lock.WriteString(appID); err != nil {
		return "", err
	}
	return appID, nil
}