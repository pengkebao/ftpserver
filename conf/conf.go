package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"strconv"

	"github.com/Unknwon/goconfig"
)

var (
	Port      int
	PermOwner string
	PermGroup string
	RootPath  string
	Users     map[string]UserInfo
	PasvPort  string
)

type UserInfo struct {
	Password string
	Dir      string
}

func init() {
	//读取配置文件：
	cfg, err := goconfig.LoadConfigFile("conf/config.ini")
	if err != nil {
		fmt.Println("Fail to load config.ini file: " + err.Error())
		os.Exit(1)
	}
	//端口
	Port, _ = cfg.Int("server", "port")
	//权限：
	PermOwner, _ = cfg.GetValue("perm", "owner")
	PermGroup, _ = cfg.GetValue("perm", "group")
	//上传目录dir
	RootPath, _ = cfg.GetValue("server", "defaultDir")

	//被动端口
	pasvMinPort, _ := cfg.GetValue("server", "pasv_min_port")
	pasvMaxPort, _ := cfg.GetValue("server", "pasv_max_port")

	if len(pasvMinPort) > 0 && len(pasvMaxPort) > 0 {
		minPort, _ := strconv.Atoi(pasvMinPort)
		maxPort, _ := strconv.Atoi(pasvMaxPort)
		if minPort >= maxPort {
			fmt.Println("最小端口不能大于或等于最大端口相同")
			os.Exit(1)
		}
		PasvPort = fmt.Sprintf("%d", minPort) + "-" + fmt.Sprintf("%d", maxPort)
	}

	//ReadFile函数会读取文件的全部内容7，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile("conf/users.json")
	if err != nil {
		if err != nil {
			fmt.Println("Fail to load users.json file: " + err.Error())
			os.Exit(1)
		}
	}
	var users []map[string]string
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Fail to load users.json file: " + err.Error())
		os.Exit(1)
	}
	Users = make(map[string]UserInfo)
	for _, v := range users {
		Users[v["user"]] = UserInfo{v["password"], v["dir"]}
	}
}
