//    Copyright (C) 2016  Chen Yan
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
//    Author: Chen Yan <leochenlinux@gmail.com>

package main

import (
	"encoding/json"
	"fmt"
	"gowbem"
	"os"
)

type Client struct {
	conn *gowbem.WBEMConnection
}

type CliMeth func(*Client, string) ([]byte, error)

var MethMap map[string]CliMeth = map[string]CliMeth{
	"ei":  (*Client).EnumerateInstances,
	"ein": (*Client).EnumerateInstanceNames,
	"gc":  (*Client).GetClass,
	"gi":  (*Client).GetInstance,
	"di":  (*Client).DeleteInstance,
	"im":  (*Client).InvokeMethod,
}

func NewClient(url string) *Client {
	conn, err := gowbem.NewWBEMConn(url)
	if nil != err {
		return nil
	}
	return &Client{conn}
}

func (cli *Client) GetClass(className string) ([]byte, error) {
	var iClassName gowbem.ClassName = gowbem.ClassName{
		Name: className,
	}
	class, err := cli.conn.GetClass(&iClassName, true, true, false, nil)
	if nil != err {
		return nil, err
	}
	res, _ := json.MarshalIndent(&class, "", "    ")
	return res, nil
}

func (cli *Client) EnumerateInstances(className string) ([]byte, error) {
	var iClassName gowbem.ClassName = gowbem.ClassName{
		Name: className,
	}
	namedInstance, err := cli.conn.EnumerateInstances(&iClassName, true, false, nil)
	if nil != err {
		return nil, err
	}
	res, _ := json.MarshalIndent(&namedInstance, "", "    ")
	return res, nil
}

func (cli *Client) EnumerateInstanceNames(className string) ([]byte, error) {
	var iClassName gowbem.ClassName = gowbem.ClassName{
		Name: className,
	}
	instanceName, err := cli.conn.EnumerateInstanceNames(&iClassName)
	if nil != err {
		return nil, err
	}
	res, _ := json.MarshalIndent(&instanceName, "", "    ")
	return res, nil
}

func (cli *Client) GetInstance(className string) ([]byte, error) {
	var iClassName gowbem.ClassName = gowbem.ClassName{
		Name: className,
	}
	instanceNames, err := cli.conn.EnumerateInstanceNames(&iClassName)
	if nil != err {
		return nil, err
	}
	if 0 == len(instanceNames) {
		err = fmt.Errorf("No instance")
		return nil, err
	}
	for i, instName := range instanceNames {
		fmt.Printf("Instance [%d]:\n", i+1)
		res, _ := json.MarshalIndent(instName, "", "    ")
		fmt.Println(string(res))
	}
	fmt.Print("Choose instance:  ")
	idx := 0
	fmt.Scanf("%d\n", &idx)
	if 0 >= idx || len(instanceNames) < idx {
		err = fmt.Errorf("Invalid index")
		return nil, err
	}
	inst, err := cli.conn.GetInstance(&instanceNames[idx-1], false, nil)
	if nil != err {
		return nil, err
	}
	res, _ := json.MarshalIndent(&inst, "", "    ")
	return res, nil
}

func (cli *Client) DeleteInstance(className string) ([]byte, error) {
	var iClassName gowbem.ClassName = gowbem.ClassName{
		Name: className,
	}
	instanceNames, err := cli.conn.EnumerateInstanceNames(&iClassName)
	if nil != err {
		return nil, err
	}
	if 0 == len(instanceNames) {
		err = fmt.Errorf("No instance")
		return nil, err
	}
	for i, instName := range instanceNames {
		fmt.Printf("Instance [%d]:\n", i+1)
		res, _ := json.MarshalIndent(instName, "", "    ")
		fmt.Println(string(res))
	}
	fmt.Print("Choose instance:  ")
	idx := 0
	fmt.Scanf("%d\n", &idx)
	if 0 >= idx || len(instanceNames) < idx {
		err = fmt.Errorf("Invalid index")
		return nil, err
	}
	err = cli.conn.DeleteInstance(&instanceNames[idx-1])
	return nil, err
}

func (cli *Client) InvokeMethod(className string) ([]byte, error) {
	var iClassName gowbem.ClassName = gowbem.ClassName{
		Name: className,
	}
	instanceNames, err := cli.conn.EnumerateInstanceNames(&iClassName)
	if nil != err {
		return nil, err
	}
	class, err := cli.conn.GetClass(&iClassName, true, true, false, nil)
	if nil != err {
		return nil, err
	}
	if 0 == len(class[0].Method) {
		err = fmt.Errorf("No method")
		return nil, err
	}
	if 0 == len(instanceNames) {
		err = fmt.Errorf("No instance")
		return nil, err
	}
	for i, instName := range instanceNames {
		fmt.Printf("Instance [%d]:\n", i+1)
		res, _ := json.MarshalIndent(instName, "", "    ")
		fmt.Println(string(res))
	}
	fmt.Print("Choose instance:  ")
	idx := 0
	fmt.Scanf("%d\n", &idx)
	if 0 >= idx || len(instanceNames) < idx {
		err = fmt.Errorf("Invalid index")
		return nil, err
	}
	instName := instanceNames[idx-1]
	//fmt.Println("")
	for i, method := range class[0].Method {
		fmt.Printf("Method [%d]: %s\n", i+1, method.Name)
	}
	fmt.Print("Choose method:  ")
	idx = 0
	fmt.Scanf("%d\n", &idx)
	if 0 >= idx || len(class[0].Method) < idx {
		err = fmt.Errorf("Invalid index")
		return nil, err
	}
	method := class[0].Method[idx-1]
	var paramVal []gowbem.ParamValue = []gowbem.ParamValue{}
	if 0 < len(method.Parameter) {
		for _, param := range method.Parameter {
			str := ""
			fmt.Printf("Parameter '%s':  ", param.Name)
			fmt.Scanf("%s\n", &str)
			if "" != str {
				paramVal = append(paramVal, gowbem.ParamValue{Name: param.Name, Value: &gowbem.Value{str}})
			}
		}
	}
	objName := gowbem.ObjectName{InstanceName: &instName}
	ret, paramVal, err := cli.conn.InvokeMethod(&objName, method.Name, paramVal)
	if nil != err {
		return nil, err
	}
	//fmt.Println("")
	fmt.Printf("Return code: %d\n", ret)
	res, _ := json.MarshalIndent(paramVal, "", "    ")
	return res, nil
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  ", os.Args[0], "<scheme>://[<username>[:<passwd>]@]<host>[/<namespace>][:<port>] <act> <param>")
	fmt.Println("<act>:")
	fmt.Println("  ", "ei  - EnumerateInstances")
	fmt.Println("  ", "ein - EnumerateInstanceNames")
	fmt.Println("  ", "gc  - GetClass")
	fmt.Println("  ", "gi  - GetInstance")
	fmt.Println("  ", "di  - DeleteInstance")
	fmt.Println("  ", "im  - InvokeMethod")
	fmt.Println("Examples:")
	fmt.Println("  ", os.Args[0], "http://USER:PASSWD@127.0.0.1 ei CIM_ComputerSystem")
	fmt.Println("  ", os.Args[0], "https://USER:PASSWD@127.0.0.1:5989/root/cimv2 ein CIM_ComputerSystem")
	//fmt.Println("")
}

func main() {
	//fmt.Println("")
	if 4 != len(os.Args) {
		fmt.Println("Error: Invalid parameter(s)")
		//fmt.Println("---")
		usage()
		os.Exit(0)
	}
	gowbem.SetLoggerEnabled(true)
	cli := NewClient(os.Args[1])
	if nil == cli {
		fmt.Println("Error: Invalid URL")
		os.Exit(1)
	}
	if nil == MethMap[os.Args[2]] {
		fmt.Println("Error: Invalid action")
		os.Exit(1)
	}
	res, err := MethMap[os.Args[2]](cli, os.Args[3])
	if nil != err {
		fmt.Println("Error:", err.Error())
	} else if nil != res {
		fmt.Println(string(res))
	}
	//fmt.Println("")
	os.Exit(0)
}
