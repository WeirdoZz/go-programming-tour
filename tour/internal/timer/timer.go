package timer

import "time"

// GetNowTime 获取现在时间
func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

// GetCalculateTime 根据给定时间和duration计算得到的时间
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	// 解析传入的duration，比如1h,10us等
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}
