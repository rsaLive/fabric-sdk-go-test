package main

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var configPath string = "./conf/config.yaml"

var channelId string = "zlktchannel"

// peer chaincode query -C zlktchannel -n myacc -c '{"Args":["query","a"]}'

// var chainCodeId string = "first_chaincode1"
var chainCodeName string = "myacc"

func main() {
	sdk, err := fabsdk.New(config.FromFile(configPath))
	defer sdk.Close()
	if err != nil {
		panic(err.Error())
	}

	//指定通道 组织 和用户 管理员才有权限，所有必须是Admin
	//ctx := sdk.ChannelContext(channelId, fabsdk.WithOrg("org1"), fabsdk.WithUser("Admin"))
	ctx := sdk.ChannelContext(channelId, fabsdk.WithOrg("Org1"), fabsdk.WithUser("Admin"))

	// 创建通道客户端
	cc, _err := channel.New(ctx)
	if _err != nil {
		panic(_err.Error())
	}

	// first_chaincode1  icba  cmbc
	//查询链码
	respQuery, err := cc.Query(channel.Request{
		ChaincodeID:     chainCodeName,         //链码名称
		Fcn:             "query",               // 函数名称
		Args:            [][]byte{[]byte("a")}, // 参数
		TransientMap:    nil,
		InvocationChain: nil,
		IsInit:          false,
	}, channel.WithTargetEndpoints("peer0.org1.example.com"))

	fmt.Println("resp", respQuery)
	//增删改
	/*respExecute, err := cc.Execute(channel.Request{
		ChaincodeID:     chainCodeName,            //链码名称
		Fcn:             "query",                  // 函数名称
		Args:            [][]byte{[]byte("cmbc")}, // 参数
		TransientMap:    nil,
		InvocationChain: nil,
		IsInit:          false,
	}, channel.WithTargetEndpoints("peer0.org1.example.com"))
	fmt.Println("respExecute", respExecute)*/
}
