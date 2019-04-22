package code

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"tools"
)

func LogInit2(logName string, logPath string, maxRemainCnt uint) {
	logHook := tools.NewHook(logName, logPath, maxRemainCnt)
	logrus.AddHook(logHook)
}

func ExecCommand(command string) (data string, err error) {
	defer tools.MRecover()
	cmd := exec.Command("/bin/bash", "-c", command)
	// 创建获取命令输出的管道
	stdout, e := cmd.StdoutPipe()
	if e != nil {
		var logger = logrus.WithFields(logrus.Fields{
			"command": command,
			"err":     string(e.Error()),
		})
		logger.Error("StdoutPipe command error : " + e.Error())
		return "nil", errors.New("StdoutPipe command error : " + e.Error())
	}
	// 执行命令
	if err := cmd.Start(); err != nil {
		var logger = logrus.WithFields(logrus.Fields{
			"command": command,
			"err":     string(e.Error()),
		})
		logger.Error("exec command error : " + err.Error())
		return "", errors.New("exec command error : " + err.Error())
	}
	// 读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		var logger = logrus.WithFields(logrus.Fields{
			"command": command,
			"err":     string(e.Error()),
		})
		logger.Error("exec command error : " + err.Error())
		return "", errors.New("read stdout error : " + err.Error())
	}

	if err := cmd.Wait(); err != nil {
		var logger = logrus.WithFields(logrus.Fields{
			"command": command,
			"err":     string(e.Error()),
		})
		logger.Error("exec wait error : " + err.Error())
		return "", errors.New("exec wait error : " + err.Error())
	}
	return string(bytes), nil
}
