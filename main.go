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

var chainCodeName string = "first_chaincode1"

//var chainCodeName string = "myacc"

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
		ChaincodeID: chainCodeName,         //链码名称
		Fcn:         "query",               // 函数名称
		Args:        [][]byte{[]byte("cmbc")}, // 参数
	}, channel.WithTargetEndpoints("peer0.org1.example.com"))

	fmt.Printf("respQuery Payload: %v\n", string(respQuery.Payload))
	fmt.Printf("respQuery response: %v\n", respQuery.Responses)
	//增删改

	/*
		targets := []fab.Peer{
				mocks.NewMockPeer("Org1MSP", "peer0.org1.example.com"),
				mocks.NewMockPeer("Org2MSP", "peer0.org2.example.com"),
			}
		respExecute, err := cc.Execute(channel.Request{
			ChaincodeID: chainCodeName,                                                            //链码名称
			Fcn:         "invoke",                                                                 // 函数名称
			Args:        [][]byte{[]byte("invoke"), []byte("a"), []byte("b"), []byte("10")}, // 参数
		}, channel.WithTargets(targets...))
	*/

	respExecute, err := cc.Execute(channel.Request{
		ChaincodeID: chainCodeName,                                          //链码名称
		Fcn:         "invoke",                                               // 函数名称
		Args:        [][]byte{[]byte("cmbc"), []byte("icba"), []byte("10")}, // 参数
	}, channel.WithTargetEndpoints("peer0.org1.example.com"))
	if err != nil {
		fmt.Println("----------------\n")
		fmt.Println("respExecute err", err.Error())
		fmt.Println("----------------\n")
	}
	fmt.Println("respExecute", respExecute.Payload)
}
