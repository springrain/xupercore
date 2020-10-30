package cmd

import (
	"fmt"

	"github.com/xuperchain/xupercore/kernel/engines"
	econf "github.com/xuperchain/xupercore/kernel/engines/config"
	"github.com/xuperchain/xupercore/kernel/engines/xuperos"
	"github.com/xuperchain/xupercore/kernel/engines/xuperos/def"
	"github.com/xuperchain/xupercore/lib/logs"
	sconf "github.com/xuperchain/xupercore/server/config"
	// import要使用的领域组件驱动
)

type StartupCmd struct {
	BaseCmd
}

func GetStartupCmd() *StartupCmd {
	startupCmdIns = new(StartupCmd)

	// 定义命令行参数变量
	var envCfgPath string

	startupCmdIns.cmd = &cobra.Command{
		Use:           "xchain startup [flag]",
		Short:         "Start up the blockchain node.",
		Example:       "xchain startup --conf /home/rd/xchain/conf/env.yaml",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return StartupXchain(envCfgPath)
		},
	}

	// 设置命令行参数并绑定变量
	startupCmdIns.cmd.Flags().StringVarP(&envCfgPath, "conf", "c", "",
		"engine environment config file path")

	return httpServerCmdIns
}

// 启动节点
func StartupXchain(envCfgPath string) error {
	// 加载基础配置
	envConf, servConf, err := loadConf(envCfgPath)
	if err != nil {
		return err
	}

	// 初始化日志
	logs.InitLog(envConf.GenConfFilePath(envCfg.LogConf), envCfg.GenDirAbsPath(envCfg.LogDir))

	// 实例化区块链引擎
	engine, err := engines.CreateBCEngine(def.BCEngineName, envConf)
	if err != nil {
		return err
	}
	// 实例化rpc server
	rpcServ, err := rpc.NewRpcServMG(servConf, engine)
	if err != nil {
		return err
	}

	// 启动服务和区块链引擎
	wg := &sync.WaitGroup{}
	wg.Add(2)
	engChan := runEngine(engine)
	rpcChan := runRpcServ(rpcServ)

	// 阻塞等待进程退出指令
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for {
			select {
			case err := <-engChan:
				wg.Done()
				rpcServ.Exit()
			case err := <-runRpcServ():
				wg.Done()
				engine.Stop()
			case <-sigChan:
				rpcServ.Exit()
				engine.Stop()
			}
		}
	}()

	// 等待异步任务全部退出
	wg.Wait()
	return nil
}

func loadConf(envCfgPath string) (econf.EnvConf, sconf.ServConf, error) {
	// 加载环境配置
	envConf, err := econf.LoadEnvConf(envCfgPath)
	if err != nil {
		return nil, nil, err
	}

	// 加载服务配置
	servConf, err := sconf.LoadServConf(envConf.GenConfFilePath(envConf.ServConf))
	if err != nil {
		return nil, nil, err
	}

	return envConf, servConf, nil
}

func runEngine(engine engines.BCEngine) <-chan error {
	exitCh := make(chan error)

	// 启动引擎，监听退出信号
	go func() {
		err := <-engine.Run()
		exitCh <- err
	}()

	return exitCh
}

func runRpcServ(servMG *rpc.RpcServMG) <-chan error {
	exitCh := make(chan error)

	// 启动服务，监听退出信号
	go func() {
		err := servMG.Run()
		exitCh <- err
	}()

	return exitCh
}
