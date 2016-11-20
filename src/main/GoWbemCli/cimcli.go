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

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  ", os.Args[0], "<scheme>://[<username>[:<passwd>]@]<host>[/<namespace>][:<port>] <act> <param>")
	fmt.Println("Examples:")
	fmt.Println("  ", os.Args[0], "http://USER:PASSWD@127.0.0.1 ei CIM_SystemComputer")
	fmt.Println("  ", os.Args[0], "https://USER:PASSWD@127.0.0.1 ein CIM_SystemComputer")
	fmt.Println("")
}

func main() {
	fmt.Println("")

	if 4 != len(os.Args) {
		fmt.Println("Error: Invalid parameter(s)")
		fmt.Println("---")
		usage()
		os.Exit(1)
	}
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
	if nil == err {
		fmt.Println(string(res))
	} else {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println("")
	os.Exit(0)
}
