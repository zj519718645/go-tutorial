package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/zj519718645/go-tutorial/projects/timeformatter/internal/timer"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		//nowTime := timer.GetNowTime()
		//log.Printf("输出结果: %s %d", nowTime.Format("2016-10-02 11:09:12"), nowTime.Unix())
		var currentTimer time.Time
		var layout = "2001-12-25 15:05:19"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 1 {
				layout = "2001-12-25 15:05:19"
			}
			if space == 0 {
				layout = "2001-12-25"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCaculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer GetCalculateTimer err: %v", err)
		}
		log.Printf("输出结果：%s, %d", t.Format(layout), t.Unix())
	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s, %d", nowTime.Format("2001-02-21 15:10:13"), nowTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculateTime", "c", "", "需要计算的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间，单位可以是ms, s, m, h")
}
