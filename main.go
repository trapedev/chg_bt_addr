package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
        addr := flag.String("addr", "", "bluetooth addr, format: 00:00:00:00:00:00")
        flag.Parse()
        addrArr := make([]string, 6)
	if *addr == "00:00:00:00:00:00" || len([]rune(*addr)) != 17 {
		fmt.Println("Invalid address")
		return
	} else if addr == nil {
		fmt.Println("Address is NIL")
		return
	} 
        prsAddr := strings.Split(*addr, ":")
        for i := range prsAddr {
                addrArr[len(prsAddr) - i - 1] = prsAddr[i]
        }
        cmd1 := exec.Command("sudo", "hcitool", "cmd", "0x04", "0x009")
        output, err := cmd1.CombinedOutput()
        if err != nil {
                panic(err)
        }

        cmd2 := exec.Command("sudo", "hcitool", "cmd", "0x3f", "0x001", fmt.Sprintf("0x%s", addrArr[0]), fmt.Sprintf("0x%s", addrArr[1]), fmt.Sprintf("0x%s", addrArr[2]), fmt.Sprintf("0x%s", addrArr[3]), fmt.Sprintf("0x%s", addrArr[4]), fmt.Sprintf("0x%s", addrArr[5]))
        output, err = cmd2.CombinedOutput()
        if err != nil {
                panic(err)
        }
        fmt.Println(string(output))

        cmd3 := exec.Command("sudo", "hciconfig", "reset")
        output, err = cmd3.CombinedOutput()
        if err != nil {
                panic(err)
        }
        fmt.Println(string(output))

        cmd4 := exec.Command("sudo", "systemctl", "restart", "bluetooth")
        output, err = cmd4.CombinedOutput()
        if err != nil {
                panic(err)
        }
        fmt.Println(string(output))
}
