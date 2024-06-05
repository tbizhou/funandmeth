package deploy

import "fmt"

type Deploy struct {
	Command  string
	Args     []string
	FilePath string
}

type DeployFileCopy struct {
	Command string
	Args    []string
	SRCPath string
	DSTPath string
}

type Deployer interface {
	Run() string
}

func (d Deploy) Run() string {
	return fmt.Sprint(d.Command, " ", d.Args, " ", d.FilePath)
}

func (d DeployFileCopy) Run() string {
	return fmt.Sprint(d.Command, " ", d.Args, " ", d.SRCPath, " ", d.DSTPath)
}

func RunCmd(d Deployer) string {
	return d.Run()
}
func AllDeploy() {
	deployMysql := Deploy{
		Command:  "mysql",
		Args:     []string{"-uroot", "-p123456"},
		FilePath: "/data/mysql/data",
	}
	deployRedis := Deploy{
		Command:  "redis-server",
		Args:     []string{"-p 6379"},
		FilePath: "/data/redis/data",
	}
	fmt.Println(RunCmd(deployMysql))
	fmt.Println(RunCmd(deployRedis))
}

func AllFileCopy() {
	deployMysqlFileCp := DeployFileCopy{
		Command: "cp",
		Args:    []string{"-r"},
		SRCPath: "/data/mysql/data",
		DSTPath: "/data/mysql/data",
	}
	deployRedisFileCp := DeployFileCopy{
		Command: "cp",
		Args:    []string{"-r"},
		SRCPath: "/data/redis/data",
		DSTPath: "/data/redis/data",
	}
	fmt.Println(RunCmd(deployMysqlFileCp))
	fmt.Println(RunCmd(deployRedisFileCp))

}
