package ziface

//服务器接口
type IServer interface {
	Start() //启动
	Stop()  //停止
	Serve() //运行
}
