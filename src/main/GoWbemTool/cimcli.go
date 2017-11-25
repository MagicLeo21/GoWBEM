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
	"encoding/xml"
	"flag"
	"fmt"
	"gowbem"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	"SI":  (*Client).SubscribeIndications,
	"LI":  (*Client).ListenIndications,
	"CA":  (*Client).CleanAllSubscriptions,
	"CL":  (*Client).CancelLocalSubscription,
}

func NewClient(url string) *Client {
	conn, err := gowbem.NewWBEMConn(url)
	if nil != err {
		return nil
	}
	return &Client{conn}
}

func (cli *Client) ExecQuery(query string, queryLanguage string) ([]byte, error) {
	instance, err := cli.conn.ExecQuery( queryLanguage, query )
	if nil != err {
		return nil, err
	}
	res, _ := json.MarshalIndent( &instance, "", "    " )
	return res, nil
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
	if 0 < len(method.ParameterArray) {
		for _, paramArray := range method.ParameterArray {
			arry := gowbem.ValueArray{}
			for {
				str := ""
				fmt.Printf("ParameterArray '%s':  ", paramArray.Name)
				fmt.Scanf("%s\n", &str)
				if "" != str {
					arry.Value = append(arry.Value, gowbem.Value{Value: str})
				} else {
					if 0 < len(arry.Value) {
						paramVal = append(paramVal, gowbem.ParamValue{Name: paramArray.Name, ValueArray: &arry})
					}
					break
				}
			}
		}
	}
	if 0 < len(method.ParameterReference) {
		for _, paramRef := range method.ParameterReference {
			str := ""
			fmt.Printf("ParameterReference '%s' ref to class:  ", paramRef.Name)
			fmt.Scanf("%s\n", &str)
			if "" != str {
				var refClsName gowbem.ClassName = gowbem.ClassName{
					Name: str,
				}
				refInstNames, err := cli.conn.EnumerateInstanceNames(&refClsName)
				if nil != err {
					return nil, err
				}
				if 0 == len(refInstNames) {
					err = fmt.Errorf("No instance")
					return nil, err
				}
				for i, refInstName := range refInstNames {
					fmt.Printf("Instance [%d]:\n", i+1)
					res, _ := json.MarshalIndent(refInstName, "", "    ")
					fmt.Println(string(res))
				}
				fmt.Print("Choose instance:  ")
				idx := 0
				fmt.Scanf("%d\n", &idx)
				if 0 >= idx || len(refInstNames) < idx {
					err = fmt.Errorf("Invalid index")
					return nil, err
				}
				var valRef = gowbem.ValueReference{InstanceName: &refInstNames[idx-1]}
				paramVal = append(paramVal, gowbem.ParamValue{Name: paramRef.Name, ValueReference: &valRef})
			}
		}
	}
	objName := gowbem.ObjectName{InstanceName: &instName}
	ret, paramVal, err := cli.conn.InvokeMethod(&objName, method.Name, paramVal)
	if nil != err {
		return nil, err
	}
	fmt.Printf("Return code: %d\n", ret)
	res, _ := json.MarshalIndent(paramVal, "", "    ")
	return res, nil
}

func (cli *Client) SubscribeIndications(unused string) ([]byte, error) {
	localName, _ := GetHostName()
	localIP, _ := GetLocalIP(cli.conn.GetHostAddr())
	localPort := 59988

	inst := NewIndicationFilter(localName, "root/cimv2")
	fmt.Println("Creating CIM_IndicationFilter...")
	err := cli.conn.CreateInstance(inst)
	if nil != err && false == strings.Contains(err.Error(), "11 - CIM_ERR_ALREADY_EXISTS") {
		return nil, err
	}
	inst = NewListenerDestination(localName, localIP, localPort)
	fmt.Println("Creating CIM_ListenerDestinationCIMXML...")
	err = cli.conn.CreateInstance(inst)
	if nil != err && false == strings.Contains(err.Error(), "11 - CIM_ERR_ALREADY_EXISTS") {
		return nil, err
	}
	inst = NewIndicationSubscription(localName, cli.conn.GetNamespace(), localIP, localPort)
	fmt.Println("Creating CIM_IndicationSubscription...")
	err = cli.conn.CreateInstance(inst)
	if nil != err && false == strings.Contains(err.Error(), "11 - CIM_ERR_ALREADY_EXISTS") {
		return nil, err
	}
	return nil, nil
}

func (cli *Client) CancelLocalSubscription(unused string) ([]byte, error) {
	localName, _ := GetHostName()
	localIP, _ := GetLocalIP(cli.conn.GetHostAddr())

	instanceName := NewIndicationSubscriptionInstName(localName, cli.conn.GetNamespace(), localIP)
	fmt.Println("Deleting IndicationSubscription...")
	err := cli.conn.DeleteInstance(instanceName)
	if nil != err && false == strings.Contains(err.Error(), "6 - CIM_ERR_NOT_FOUND") {
		return nil, err
	}
	instanceName = NewListenerDestinationInstName(localName, localIP)
	fmt.Println("Deleting ListenerDestination...")
	err = cli.conn.DeleteInstance(instanceName)
	if nil != err && false == strings.Contains(err.Error(), "6 - CIM_ERR_NOT_FOUND") {
		return nil, err
	}
	instanceName = NewIndicationFilterInstName(localName)
	fmt.Println("Deleting IndicationFilter...")
	err = cli.conn.DeleteInstance(instanceName)
	if nil != err && false == strings.Contains(err.Error(), "6 - CIM_ERR_NOT_FOUND") {
		return nil, err
	}
	return nil, nil
}

func (cli *Client) CleanAllSubscriptions(unused string) ([]byte, error) {
	iClassName := gowbem.ClassName{
		Name: "CIM_IndicationSubscription",
	}
	fmt.Println("Enumerating IndicationSubscription...")
	instanceNames, err := cli.conn.EnumerateInstanceNames(&iClassName)
	if nil != err {
		return nil, err
	}
	if 0 != len(instanceNames) {
		fmt.Println("Deleting IndicationSubscription...")
		for _, instName := range instanceNames {
			err = cli.conn.DeleteInstance(&instName)
			if nil != err {
				return nil, err
			}
		}
	}
	iClassName = gowbem.ClassName{
		Name: "CIM_ListenerDestinationCIMXML",
	}
	fmt.Println("Enumerating ListenerDestination...")
	instanceNames, err = cli.conn.EnumerateInstanceNames(&iClassName)
	if nil != err {
		return nil, err
	}
	if 0 != len(instanceNames) {
		fmt.Println("Deleting ListenerDestination...")
		for _, instName := range instanceNames {
			err = cli.conn.DeleteInstance(&instName)
			if nil != err {
				return nil, err
			}
		}
	}
	iClassName = gowbem.ClassName{
		Name: "CIM_IndicationFilter",
	}
	fmt.Println("Enumerating IndicationFilter...")
	instanceNames, err = cli.conn.EnumerateInstanceNames(&iClassName)
	if nil != err {
		return nil, err
	}
	if 0 != len(instanceNames) {
		fmt.Println("Deleting IndicationFilter...")
		for _, instName := range instanceNames {
			err = cli.conn.DeleteInstance(&instName)
			if nil != err {
				return nil, err
			}
		}
	}
	return nil, nil
}

func (cli *Client) ListenIndications(unused string) ([]byte, error) {
	localPort := 59988
	fmt.Printf("Start listening on port %d...\n", localPort)
	http.HandleFunc("/", ListenerHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", localPort), nil)
	return nil, err
}

func GetHostName() (string, error) {
	name, err := os.Hostname()
	if nil != err {
		return "", err
	}
	return name, nil
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
	req.Body.Close()
	cim := gowbem.CIM{}
	err := xml.Unmarshal(body, &cim)
	if nil == err {
		if nil != cim.Message &&
			nil != cim.Message.SimpleExpReq &&
			nil != cim.Message.SimpleExpReq.ExpMethodCall &&
			0 != len(cim.Message.SimpleExpReq.ExpMethodCall.ExpParamValue) &&
			nil != cim.Message.SimpleExpReq.ExpMethodCall.ExpParamValue[0].Instance {
			msg := ""
			ts := "Unknown"
			uuid := "Unknown"
			ipaddr := req.RemoteAddr
			for _, prop := range cim.Message.SimpleExpReq.ExpMethodCall.ExpParamValue[0].Instance.Property {
				if "Message" == prop.Name && nil != prop.Value {
					msg = prop.Value.Value
				} else if "IndicationTime" == prop.Name && nil != prop.Value {
					ts = prop.Value.Value
				} else if "SystemUUID" == prop.Name && nil != prop.Value {
					uuid = prop.Value.Value
				}
			}
			fmt.Printf("%s | %s | %s | %s\n", ipaddr, ts, uuid, msg)
		}
	}
}

func usage() {
	base := filepath.Base(os.Args[0])
	fmt.Println("Usage:")
	fmt.Printf("    %s -o <action> [-u <url>] [-c <class>] [-t <timeout>]\n", base)
	fmt.Printf("    %s -o exq -q <WqlQuery> [-ql <QueryLang>] [-u <url>] [-t <timeout>]\n", base)
	fmt.Printf("<url>:\n")
	fmt.Printf("    <scheme>://[<username>[:<passwd>]@]<host>[:<port>][/<namespace>]\n")
	fmt.Printf("<act>:\n")
	fmt.Printf("    ei  - EnumerateInstances\n")
	fmt.Printf("    ein - EnumerateInstanceNames\n")
	fmt.Printf("    gc  - GetClass\n")
	fmt.Printf("    gi  - GetInstance\n")
	fmt.Printf("    di  - DeleteInstance\n")
	fmt.Printf("    im  - InvokeMethod\n")
	fmt.Printf("    SI  - SubscribeIndications\n")
	fmt.Printf("    LI  - ListenForIndications\n")
	fmt.Printf("    CA  - CleanAllSubscriptions\n")
	fmt.Printf("    CL  - CancelLocalSubscription\n")
	fmt.Printf("Examples:\n")
	fmt.Printf("    %s -u http://USER:PASSWD@127.0.0.1 -o ei -c CIM_ComputerSystem\n", base)
	fmt.Printf("    %s -u https://USER:PASSWD@127.0.0.1:5989/root/cimv2 -o ein -c CIM_ComputerSystem\n", base)
	fmt.Printf("    %s -u https://USER:PASSWD@127.0.0.1:5989/root/interop -o SI\n", base)
	fmt.Printf("    %s -o LI\n", base)
}

func main() {
	flag.Usage = usage
	url := flag.String("u", "", "")
	cls := flag.String("c", "", "")
	opt := flag.String("o", "", "")
	to  := flag.Int("t", 120, "")
	qlang := flag.String("ql", "WQL", "")
	query := flag.String("q", "", "")

	flag.Parse()
	cli := NewClient(*url)
	if nil == cli {
	    usage()
	} else if "exq" == *opt && "" != *query {
		cli.conn.SetHttpTimeout(time.Second * time.Duration(*to))
		res, err := (*cli).ExecQuery(*query, *qlang)
		if nil != err {
			log.Println("Error:", err.Error())
		} else if nil != res {
			fmt.Println(string(res))
		}
	} else if nil == MethMap[*opt] {
		// not in the method map, which is only for actions that have a class parameter
		usage()
	} else {
		// the action is in the method map (for actions that take a class parameter)
		cli.conn.SetHttpTimeout(time.Second * time.Duration(*to))
		res, err := MethMap[*opt](cli, *cls)
		if nil != err {
			log.Println("Error:", err.Error())
		} else if nil != res {
			fmt.Println(string(res))
		}
	}
	os.Exit(0)
}
