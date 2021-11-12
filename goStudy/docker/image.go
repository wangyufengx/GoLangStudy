package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"goStudy/pkg/utils"
	"io"
	"os"
)

//NewClient 建立客户端连接
func NewClient() (*client.Client, error) {
	//建立本地客户端连接
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	/*
		//建立远程客户端连接
		dockerClient, err := client.NewClientWithOpts(client.WithHost("tcp://172.16.12.1:2375"), client.WithAPIVersionNegotiation())
		if err != nil {
			return nil,err
		}
	*/
	return dockerClient, err
}

//ImagePull 拉去镜像
func ImagePull(dockerClient *client.Client, imageName string) error {
	ctx := context.Background()

	//拉去镜像
	pull, err := dockerClient.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	//标准输出
	io.Copy(os.Stdout, pull)
	return nil
}

//PushImage 推送镜像
func PushImage(dockerClient *client.Client, username, password, serverAddress, imageName string) error {
	ctx := context.Background()
	authConfig := types.AuthConfig{
		Username:      username,
		Password:      password,
		ServerAddress: serverAddress,
	}
	bytes, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}

	push, err := dockerClient.ImagePush(ctx, imageName, types.ImagePushOptions{
		RegistryAuth: base64.URLEncoding.EncodeToString(bytes),
	})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, push)
	return nil
}

func ListImages(dockerClient *client.Client, imageName string) ([]types.ImageSummary, error) {
	ctx := context.Background()
	//筛选条件
	args := filters.NewArgs()
	args.Add("reference", imageName)
	//查询
	list, err := dockerClient.ImageList(ctx, types.ImageListOptions{
		All:     false,
		Filters: args,
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func ImageSave(dockerClient *client.Client, imageIDs []string, tarName string) error {
	ctx := context.Background()
	//导出镜像流
	save, err := dockerClient.ImageSave(ctx, imageIDs)
	if err != nil {
		return err
	}
	//存储
	if err = utils.GenerateFile(tarName, save); err != nil {
		return err
	}
	return nil
}

func ImageOperation(dockerClient *client.Client) error {
	ctx := context.Background()

	//拉去镜像
	pull, err := dockerClient.ImagePull(ctx, "alpine:3.13.6", types.ImagePullOptions{})
	if err != nil {
		return err
	}
	//标准输出
	io.Copy(os.Stdout, pull)

	//筛选条件
	args := filters.NewArgs()
	args.Add("reference", "alpine:3.13.6")
	//查询
	list, err := dockerClient.ImageList(ctx, types.ImageListOptions{
		All:     false,
		Filters: args,
	})
	if err != nil {
		return err
	}
	fmt.Println(list)

	authConfig := types.AuthConfig{
		Username:      "用户名",
		Password:      "密码",
		ServerAddress: "私库",
	}
	bytes, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}

	push, err := dockerClient.ImagePush(ctx, "alpine:3.13.6", types.ImagePushOptions{
		RegistryAuth: base64.URLEncoding.EncodeToString(bytes),
	})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, push)

	//导出镜像流
	save, err := dockerClient.ImageSave(ctx, []string{list[0].ID})
	if err != nil {
		return err
	}
	//存储
	if err = utils.GenerateFile("./alpine.tar.gz", save); err != nil {
		return err
	}

	return nil
}
