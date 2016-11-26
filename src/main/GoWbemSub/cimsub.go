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
	"encoding/xml"
	"fmt"
	"gowbem"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetHostName() (string, error) {
	name, err := os.Hostname()
	if nil != err {
		return "", err
	}
	return name, nil
}

func GetDestAddr(urlstr string) (string, error) {
	res, err := url.Parse(urlstr)
	if nil != err {
		return "", err
	}
	host := strings.Split(res.Host, ":")
	return host[0], nil
}

func GetLocalIP(dest string) (string, error) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:80", dest))
	if nil != err {
		return "", err
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0], nil
}

func ListenerHandler(writer http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	cim := gowbem.CIM{}
	err := xml.Unmarshal(body, &cim)
	if nil == err {
		if nil != cim.Message &&
			nil != cim.Message.SimpleExpReq &&
			nil != cim.Message.SimpleExpReq.ExpMethodCall &&
			0 != len(cim.Message.SimpleExpReq.ExpMethodCall.ExpParamValue) &&
			nil != cim.Message.SimpleExpReq.ExpMethodCall.ExpParamValue[0].Instance {
			msg := ""
			ts := ""
			for _, prop := range cim.Message.SimpleExpReq.ExpMethodCall.ExpParamValue[0].Instance.Property {
				if "Message" == prop.Name && nil != prop.Value {
					msg = prop.Value.Value
				} else if "IndicationTime" == prop.Name && nil != prop.Value {
					ts = prop.Value.Value
				}
			}
			fmt.Printf("%s - %s\n", ts, msg)
		}
	}
}

func SubscriptionThread(destURL, localName, srcNS, localIP string, localPort int) {
	conn, err := gowbem.NewWBEMConn(destURL)
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	iInstance := NewIndicationFilter(localName, srcNS)
	fmt.Println("Creating IndicationFilter...")
	err = conn.CreateInstance(iInstance)
	if nil != err && "11 - CIM_ERR_ALREADY_EXISTS - Object already exists" != err.Error() {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	iInstance = NewListenerDestination(localName, localIP, localPort)
	fmt.Println("Creating ListenerDestination...")
	err = conn.CreateInstance(iInstance)
	if nil != err && "11 - CIM_ERR_ALREADY_EXISTS - Object already exists" != err.Error() {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	iInstance = NewIndicationSubscription(localName, conn.GetNamespace(), localIP, localPort)
	fmt.Println("Creating IndicationSubscription...")
	err = conn.CreateInstance(iInstance)
	if nil != err && "11 - CIM_ERR_ALREADY_EXISTS - Object already exists" != err.Error() {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Done!")
}

func ListenerThread(localIP string, localPort int) {
	fmt.Printf("Start listening on %s:%d...\n", localIP, localPort)
	http.HandleFunc("/", ListenerHandler)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", localIP, localPort), nil)
	if nil != err {
		fmt.Println("ERR - Failed to start listener:", err.Error())
		os.Exit(1)
	}
}

func CleanSubscription(destURL string) {
	conn, err := gowbem.NewWBEMConn(destURL)
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var iClassName gowbem.ClassName = gowbem.ClassName{
		Name: "CIM_IndicationSubscription",
	}
	fmt.Println("Enumerating IndicationSubscription...")
	instanceName, err := conn.EnumerateInstanceNames(&iClassName)
	if nil != err {
		fmt.Println("ERR - Failed to enumerate subscription:", err.Error())
		os.Exit(1)
	}
	if 0 != len(instanceName) {
		fmt.Println("Deleting IndicationSubscription...")
		for _, inst := range instanceName {
			err = conn.DeleteInstance(&inst)
			if nil != err {
				fmt.Printf("ERR - Failed to clean subscription:", err.Error())
				os.Exit(1)
			}
		}
	}

	iClassName = gowbem.ClassName{
		Name: "CIM_ListenerDestinationCIMXML",
	}
	fmt.Println("Enumerating ListenerDestination...")
	instanceName, err = conn.EnumerateInstanceNames(&iClassName)
	if nil != err {
		fmt.Println("ERR - Failed to enumerate destination:", err.Error())
		os.Exit(1)
	}
	if 0 != len(instanceName) {
		fmt.Println("Deleting ListenerDestination...")
		for _, inst := range instanceName {
			err = conn.DeleteInstance(&inst)
			if nil != err {
				fmt.Println("ERR - Failed to clean destination:", err.Error())
				os.Exit(1)
			}
		}
	}

	iClassName = gowbem.ClassName{
		Name: "CIM_IndicationFilter",
	}
	fmt.Println("Enumerating IndicationFilter...")
	instanceName, err = conn.EnumerateInstanceNames(&iClassName)
	if nil != err {
		fmt.Println("ERR - Failed to enumerate filter:", err.Error())
		os.Exit(1)
	}
	if 0 != len(instanceName) {
		fmt.Println("Deleting IndicationFilter...")
		for _, inst := range instanceName {
			err = conn.DeleteInstance(&inst)
			if nil != err {
				fmt.Println("ERR - Failed to clean filter:", err.Error())
				os.Exit(1)
			}
		}
	}
	fmt.Println("Done!")
}

func DoSubscribeAndListen(urls string) {
	iDestAddr, _ := GetDestAddr(urls)
	iLocalName, _ := GetHostName()
	iLocalIP, _ := GetLocalIP(iDestAddr)
	iSourceNS := "root/cimv2"
	iLocalPort := 59988
	go SubscriptionThread(urls, iLocalName, iSourceNS, iLocalIP, iLocalPort)
	ListenerThread(iLocalIP, iLocalPort)
}

func DoCleanSubscription(urls string) {
	CleanSubscription(urls)
}

func DoListen(urls string) {
	iDestAddr, _ := GetDestAddr(urls)
	iLocalIP, _ := GetLocalIP(iDestAddr)
	iLocalPort := 59988
	ListenerThread(iLocalIP, iLocalPort)
}

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("  ", os.Args[0], "<-S|-L|-C> <scheme>://[<username>[:<passwd>]@]<host>[/<namespace>][:<port>]")
	fmt.Println("Examples:")
	fmt.Println("  Subscribe and listen indications: ", os.Args[0], "-S http://USER:PASSWD@127.0.0.1/root/interop")
	fmt.Println("  Just listen indications: ", os.Args[0], "-L https://USER:PASSWD@127.0.0.1/root/interop")
	fmt.Println("  Clean all subscriptions: ", os.Args[0], "-C https://USER:PASSWD@127.0.0.1/root/interop")
	fmt.Println("")
}

func main() {
	if 3 != len(os.Args) {
		Usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "-L":
		DoListen(os.Args[2])
	case "-C":
		DoCleanSubscription(os.Args[2])
	case "-S":
		DoSubscribeAndListen(os.Args[2])
	default:
		Usage()
		os.Exit(1)
	}
	os.Exit(0)
}
