package ytchat

import (
	"bufio"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
)

var TIME_SPAN int = 30

type CountData struct {
	Time  int `json:"time"`
	Count int `json:"count"`
}

func getLiveChatFile(url string) error {
	_, err := exec.Command("yt-dlp", "--skip-download", "--write-sub", url, "-o", "result", "--no-warnings").CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func GetChatCount(url string) ([]CountData, error) {
	err := getLiveChatFile(url)
	if err != nil {
		return nil, err
	}
	fileName := "result.live_chat.json"
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var maxTimeSec int = 0
	chatCountMap := map[int]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`"videoOffsetTimeMsec": "[0-9]+"`)
		str := re.FindString(line)
		timeRe := regexp.MustCompile(`[0-9]+`)
		timeMsec, err := strconv.Atoi(timeRe.FindString(str))
		if err != nil {
			return nil, err
		}
		timeSec := timeMsec / (1000 * TIME_SPAN)
		oldCnt, found := chatCountMap[timeSec]
		if found {
			chatCountMap[timeSec] = oldCnt + 1
		} else {
			chatCountMap[timeSec] = 1
		}
		if maxTimeSec < timeSec {
			maxTimeSec = timeSec
		}
	}
	countList := make([]CountData, 0)
	for time, count := range chatCountMap {
		var countData CountData
		countData.Time = time * TIME_SPAN
		countData.Count = count
		countList = append(countList, countData)
	}
	sort.Slice(countList, func(i, j int) bool {
		return countList[i].Time < countList[j].Time
	})
	return countList, nil
}
