//对于go.mod文件：一个vscode中打开的文件夹中应该只有一个go.mod文件（没有启动多工作区之前），因此我们需要保证每次打开的文件夹里面只有一份go.mod

package main

import (
	winniepooh "b/icomefromalaska" //文件名和包名不同时，这样引入
	"b/stringutil"                 //当要用的包和文件名相同时，这样引入

	//包路径的规则：从当前go.mod文件中的module参数开始，一直到包含了要引用的函数的go文件所在的目录下面
	//比如这个例子：当前go.mod文件中module是“b”，则b就代表了go.mod文件所处的文件夹(02_package)，然后从这里一直写到name.go那几个文件所处的文件夹位置(stringutil)
	//go mod init "filename"里面的filename可以随便写，但是最好是当前路径文件夹的名字，方便理解。比如当前pwd的结果是/d/CODING/GOPROGRAM/GoBasicLearning/02_package，那么gomod命令最好是：go mod init 02_package

	"fmt"
)

func main() {
	fmt.Println(stringutil.Reverse("!!!!!oG,olleH"))
	fmt.Println(stringutil.MyName) //在前面import了相应的包之后就可以调用包暴露出来的函数了
	fmt.Println(winniepooh.BearName)
}
