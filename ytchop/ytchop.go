package ytchop

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

type ChopData struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func Download(url string, downloadDir string, chopDataList []ChopData) error {
	result, err := exec.Command("yt-dlp", "--print", "id", url, "--no-warnings").CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return err
	}
	id := string(result)
	id = strings.ReplaceAll(id, "\n", "")
	id = strings.ReplaceAll(id, "\r\n", "")

	videoName := "video_" + id + ".mp4"
	downloadPath := path.Join(downloadDir, videoName)
	fmt.Printf("%s", downloadPath)
	exec.Command("yt-dlp", "-f", "bv*[ext=mp4]+ba[ext=m4a]/b[ext=mp4]", "-o", downloadPath, url).CombinedOutput()
	if err != nil {
		return err
	}
	for idx, chopData := range chopDataList {
		startSec := getSec(chopData.Start)
		endSec := getSec(chopData.End)
		duration := endSec - startSec
		chopVideoName := videoName + "_" + strconv.Itoa(idx+1) + ".mp4"
		chopVideoPath := path.Join(downloadDir, chopVideoName)
		exec.Command("ffmpeg", "-ss", strconv.Itoa(startSec), "-i", downloadPath, "-t", strconv.Itoa(duration), chopVideoPath).CombinedOutput()
	}
	os.Remove(downloadPath)
	return nil
}

func getSec(time string) int {
	splitedTimes := strings.Split(time, ":")
	timeCnt := len(splitedTimes)
	var sec int = 0
	for i, t := range splitedTimes {
		n, _ := strconv.Atoi(t)
		baseTime := int(math.Pow(60, float64(timeCnt-(i+1))))
		sec += baseTime * n
	}
	return sec
}
