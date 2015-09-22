package lib

import (
	"qiniupkg.com/api.v7/kodo"
	"golang.org/x/net/context"
	"fmt"
)

const(
	AK string = "j2A01V8ZS0AVdxU1gRv4Pep6ImSUR10-riuvpCWI"
	SK string = "JmfvJ9nfTRU6V38rK0XZ15RYxQyS2pxHm274-Lg8"
	OutUrl string = "http://7u2nnm.com1.z0.glb.clouddn.com/"
)

func UpdateQiniu(fileName string, fileData string) bool{
	kodo.SetMac(AK, SK)
	zone := 0 // 您空间(Bucket)所在的区域
	c := kodo.New(zone, nil) // 用默认配置创建 Client
	bucket := c.Bucket("imgdewei")
	ctx := context.Background()
	err := bucket.PutFile(ctx, nil, fileName, fileData, nil)
	
	if err != nil{
		fmt.Println(err)
		return false
	}else{
		return true
	}
	
	
}

