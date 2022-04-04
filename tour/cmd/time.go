package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
	"tour/internal/timer"
)

// 这两个用来接收用户想要计算的参数
var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s,%d", nowTime.Format("2022-04-04 15:53:22"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2022-04-04 15:53:22"

		//如果用户没有给出时间，默认为现在时间
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			// 计算calculateTime中有多少个空格
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2022-04-04"
			}
			if space == 1 {
				layout = "2022-04-04 15:53:22"
			}

			// 如果这一步存在异常，直接使用时间戳的格式进行处理
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}

		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err:%v", err)
		}
		log.Printf("输出结果：%s,%d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd, calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或者已经格式化之后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间，有效单位为`ns`,`us`,`ms`,`s`,`m`,`h`")
}
