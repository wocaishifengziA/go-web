package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// var m sync.Map
	// m.Store("a", 1)
	// m.Store("b", 2)
	// m.Store("c", 3)
	// fmt.Println(m.Load("a"))

	// mx := make(map[string]int)
	// m.Range(func(k, v interface{}) bool {
	// 	mx[k.(string)] = v.(int)
	// 	return true
	// })
	// fmt.Println(mx)

	// r := gin.Default()
	// g := r.Group("/cnm")
	// g.GET("/bb", func(ctx *gin.Context) {
	// 	ctx.JSON(200, "sdsd")
	// })
	// r.GET("/download", func(c *gin.Context) {
	// 	c.File("./README.pdf")
	// })
	// r.Run(":8080")

	// pflag test
	// pflagx.PlagDo()

	c := CreateApedgeAppReq{
		Name:     "pcinfer",
		Describe: "鹏城推理",
		Project:  "pc-1",
		Models: Model{
			Name:         "m1",
			Version:      "v1",
			File:         "modes-a",
			DownloadLink: "http://pc.com/m1",
		},
		ModelMountPath:  "/data/models",
		HarborUser:      "",
		HarborPass:      "",
		ContainerImage:  "ubuntu:20.04",
		ContainerRunCmd: []string{"bash", "-c", "sleep 100000"},
		ContainerPorts: ContainerPorts{
			Ports: []AppPort{
				{
					ContainerPort: 8080,
					HostPort:      8080,
					Name:          "infer",
					Describe:      "推理端口",
				},
			},
		},
		NodeIds:      []string{"aaa"},
		Envs:         make(map[string]string),
		Cpu:          1,
		Mem:          1000000000,
		DeviceType:   "CPU",
		DeviceNum:    1,
		IsReturnData: false,
	}
	s, _ := json.Marshal(c)
	fmt.Println(string(s))
}

type CreateApedgeAppReq struct {
	Name            string            `json:"name" comment:"服务名称" validate:"required,name-checker"`
	Describe        string            `json:"describe" comment:"描述"`
	Project         string            `json:"project" comment:"所属项目"`
	Models          Model             `json:"models" comment:"模型"`
	ModelMountPath  string            `json:"modelMountPath" comment:"模型挂在路径" validate:"required"`
	HarborUser      string            `json:"harborUser" comment:"harbor用户"`
	HarborPass      string            `json:"harborPass" comment:"harbor密码"`
	ContainerImage  string            `json:"containerImage" comment:"推理镜像" validate:"required"`
	ContainerRunCmd []string          `json:"containerRunCmd" comment:"容器运行命令"`
	ContainerPorts  ContainerPorts    `json:"containerPorts" comment:"容器网络端口"`
	NodeIds         []string          `json:"nodeIds" comment:"描述" validate:"required"`
	Envs            map[string]string `json:"envs" comment:"环境变量"`
	Cpu             float64           `json:"cpu" comment:"cpu"`
	Mem             int64             `json:"mem" comment:"mem"`
	DeviceType      string            `json:"deviceType" comment:"deviceType"` // CPU | GPU | NPU
	DeviceNum       float64           `json:"deviceNum" comment:"deviceNum"`
	IsReturnData    bool              `json:"isReturnData" comment:"是否数据回传"`
	DataChannel     DataChannelDetail `json:"dataChannel" comment:"数据通道信息"`
}

type Model struct {
	Name         string `json:"name" validate:"required"`
	Version      string `json:"version" validate:"required"`
	File         string `json:"file" validate:"required"`
	DownloadLink string `json:"FileServerUri"`
}

type ContainerPorts struct {
	Ports []AppPort `json:"ports"`
}

type AppPort struct {
	ContainerPort int    `json:"port"`       // 容器端口
	HostPort      int    `json:"targetPort"` // 主机端口
	Name          string `json:"name"`
	Describe      string `json:"describe"`
}

type DataChannelDetail struct {
	DataPoolId      int64  `json:"dataPoolId"`
	DataPoolName    string `json:"dataPoolName"`
	DataChannelId   int64  `json:"dataChannelId"`
	DataChannelName string `json:"dataChannelName"`
	// DataChannelInfo protocol.DataChannelInfoRsp `json:"dataChannelInfo"`
}
