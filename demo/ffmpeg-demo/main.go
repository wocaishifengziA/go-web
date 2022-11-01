package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
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
	viedos, err := GetAllVideoTimeSecend("video")
	if err != nil {
		log.Println(err)
	}
	log.Println("viedos total time:", viedos)
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
	videos, err := GetDirVideoFiles("video")
	if err != nil {
		return 0, nil
	}
	t := 0
	log.Println("viedos:", videos)
	for _, v := range videos {
		sec, err := GetVideoTimeSecend(v)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%s: %d", v, sec)
		t += sec
	}
	log.Println("viedos total time:", t)
	return t, nil
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
			if VideoFormat[s[len(s)-1]] {
				videos = append(videos, fmt.Sprintf("%s/%s", path, file.Name()))
			}
		}
	}
	return videos, nil
}
