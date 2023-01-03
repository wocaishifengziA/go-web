package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	// s := "ffmpeg -i video/aaa.mp4 -vf fps=fps=1 -q:v 2 -f image2 " + "/tmp/aaa-%08d.jpg"
	// // s := fmt.Sprintf("for video in `ls %s`; do ffmpeg -i %s/${video} -vf fps=fps=%d/%d -q:v 2 -f image2 "+"/video-images/${video}%s.jpg", "video", "video", 1, 10, "%08d") + "; done;"
	// fmt.Println(s)
	// cmd := exec.Command("bash", "-c", s)
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println("result =>", string(output))

	// sec, err := GetVideoTimeSecend("video/aaa.mp4")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("sec:", sec)
	// GetDirFiles()
	viedos, err := GetAllVideoTimeSecend("videos")
	if err != nil {
		log.Println(err)
	}
	log.Println("viedos total time:", viedos)

	// n, err := GetVideoInfoPart("/workspace/mycode/go-web/demo/ffmpeg-demo/video/bbb.mp4")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println("format:", n)
}

var ErrGetVideoTime = errors.New("get video time failed")

func GetVideoTimeSecend(v string) (int, error) {
	base := "ffmpeg -i %s 2>&1 | grep 'Duration' | cut -d ' ' -f 4 | sed s/,//"
	c := fmt.Sprintf(base, v)
	log.Println("cmd: ", c)
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("get video time failed, err: %s", err.Error())
	}
	output := strings.Replace(string(out), "\n", "", -1)
	log.Println("cmd result: ", output)
	res := strings.Split(output, ".")
	t := strings.Split(res[0], ":")
	if len(t) != 3 {
		return 0, ErrGetVideoTime
	}
	s0, err := strconv.Atoi(t[0])
	if err != nil {
		return 0, ErrGetVideoTime
	}
	s1, err := strconv.Atoi(t[1])
	if err != nil {
		return 0, ErrGetVideoTime
	}
	s2, err := strconv.Atoi(t[2])
	if err != nil {
		return 0, ErrGetVideoTime
	}
	return s0*3600 + s1*60 + s2, nil
}

func GetAllVideoTimeSecend(path string) (int, error) {
	_, err := GetDirVideoFiles(path)
	if err != nil {
		return 0, nil
	}
	// t := 0
	// log.Println("viedos:", videos)
	// for _, v := range videos {
	// 	sec, err := GetVideoTimeSecend(v)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	log.Printf("%s: %d", v, sec)
	// 	t += sec
	// }
	// log.Println("viedos total time:", t)
	return 0, nil
}

var VideoFormat = map[string]bool{
	"mp4": true,
	"flv": true,
	"avi": true,
	"mov": true,
}

func GetDirVideoFiles(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	videos := []string{}
	for _, file := range files {
		if !file.IsDir() {
			s := strings.Split(file.Name(), ".")
			name := file.Name()
			xx, err := GetDecodeEncodingString(name, strings.NewReader(name))
			if err != nil {
				fmt.Println("-----------------------")
				return nil, err
			}
			fmt.Println("xxxxxx:", xx)
			if VideoFormat[s[len(s)-1]] {
				videos = append(videos, fmt.Sprintf("%s/%s", path, file.Name()))
			}
		}
	}
	log.Println("video", videos)
	return videos, nil
}

func isGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			//编码小于等于127,只有一个字节的编码，兼容ASCII吗
			i++
			continue
		} else {
			//大于127的使用双字节编码
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func Utf8ToGbkIfGbk(data string) string {
	b := []byte(data)
	if isGBK(b) {
		utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(b)
		return string(utf8Data)
	}
	return data
}

func preNUm(data byte) int {
	str := fmt.Sprintf("%b", data)
	var i int = 0
	for i < len(str) {
		if str[i] != '1' {
			break
		}
		i++
	}
	return i
}
func isUtf8(data []byte) bool {
	for i := 0; i < len(data); {
		if data[i]&0x80 == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if data[i]&0xc0 != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

func GetVideoInfo(file string) (map[string]interface{}, error) {
	base := "ffprobe -show_streams -show_format -v quiet -print_format json %s"
	c := fmt.Sprintf(base, file)
	log.Println("cmd: ", c)
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("get video info failed, err: %s", err.Error())
	}
	outStr := string(out)
	log.Println("cmd result: ", outStr)
	res := make(map[string]interface{})
	err = json.Unmarshal(out, &res)
	if err != nil {
		return nil, fmt.Errorf("get video info failed, err: %s", err.Error())
	}
	fmt.Println(res)
	return res, nil
}

type VideoInfoPart struct {
	CodecName string  `json:"codecName"`
	Size      int64   `json:"size"`
	Duration  float64 `json:"duration"`
}

func GetVideoInfoPart(file string) (*VideoInfoPart, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	info, err := GetVideoInfo(file)
	if err != nil {
		return nil, fmt.Errorf("get video format failed, err: %s", err.Error())
	}
	streams := info["streams"].([]interface{})
	info2 := streams[1].(map[string]interface{})
	format := info["format"].(map[string]interface{})

	codec_name := info2["codec_name"].(string)
	size := format["size"].(string)
	duration := format["duration"].(string)

	sizeInt64, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("get video format failed, err: %s", err.Error())
	}
	durationFloat, err := strconv.ParseFloat(duration, 64)
	if err != nil {
		return nil, fmt.Errorf("get video format failed, err: %s", err.Error())
	}
	r := VideoInfoPart{
		CodecName: codec_name,
		Size:      sizeInt64,
		Duration:  durationFloat,
	}
	return &r, nil
}

func GetDecodeEncodingString(input string, f io.Reader) (string, error) {
	r, err := DecodeEncoding(input, bufio.NewReader(f))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func DecodeEncoding(input string, r io.Reader) (*transform.Reader, error) {
	// 将对应格式文本转换成utf-8
	// 判断传输来的文本的字符集格式是什么
	peek, err := bufio.NewReader(r).Peek(1024)
	if !errors.Is(err, io.EOF) {
		return nil, err
	}
	e, name, certain := charset.DetermineEncoding(peek, "")
	fmt.Println(name, certain)
	return transform.NewReader(r, e.NewDecoder()), nil
}
