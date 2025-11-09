package main

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
)

var (
	logger, _ = NewLogger(DefaultLogFilePath)
)

type TestCNI struct {
	CniVersion string `json:"cniVersion"`
}

func init() {
	runtime.LockOSThread()
}

func main() {
	cniFuncs := skel.CNIFuncs{
		Add:   CmdAdd,
		Del:   CmdDel,
		Check: CmdCheck,
	}
	skel.PluginMainFuncs(cniFuncs, version.All, "")
}

func CmdAdd(args *skel.CmdArgs) error {
	b, err := json.Marshal(*args)
	if err != nil {
		return fmt.Errorf("marshal CmdArgs: %w", err)
	}

	logger.Infof("CmdArgs: %s", string(b))

	sb := TestCNI{}
	err = json.Unmarshal(args.StdinData, &sb)
	if err != nil {
		return fmt.Errorf("CmdAdd() - unmarshal StdinData: %v", err)
	}

	b, err = json.Marshal(*args)
	if err != nil {
		return fmt.Errorf("marshal CmdArgs: %w", err)
	}

	logger.Infof("TestCNI: %s", string(b))

	return nil
	//return types.PrintResult(result, sb.CniVersion)

}

func CmdCheck(args *skel.CmdArgs) error {
	return nil
}

func CmdDel(args *skel.CmdArgs) error {
	//netns, err := ns.GetNS(args.Netns)
	//if err != nil {
	//	return err
	//}
	//defer netns.Close()
	//
	//err = netns.Do(func(nn ns.NetNS) error {
	//	NsNet, err := netlink.LinkByName(args.IfName)
	//	if err != nil {
	//		return err
	//	}
	//
	//	NsNetAddr, err := netlink.AddrList(NsNet, netlink.FAMILY_V4)
	//	if err != nil {
	//		return err
	//	}
	//
	//	for _, v := range NsNetAddr {
	//		if v.IPNet != nil {
	//			logger.WithFields(log.Fields{
	//				"IP":           v.IPNet.String(),
	//				"NS":           args.Netns,
	//				"Container ID": args.ContainerID,
	//			}).Info("Deleting ip from db")
	//			err := RemoveIP(v.IPNet)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//	}
	//
	//	err = netns.Close()
	//	return err
	//})

	return nil
}
