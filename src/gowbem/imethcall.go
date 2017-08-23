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

package gowbem

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func newIMechCall(name string) *IMethodCall {
	return &IMethodCall{Name: name}
}

func (iMethCall *IMethodCall) appendNamespace(namespace string) {
	var ns []Namespace = []Namespace{}
	for _, sub := range strings.Split(namespace, "/") {
		ns = append(ns, Namespace{Name: sub})
	}
	iMethCall.LocalNamespacePath = &LocalNamespacePath{ns}
}

func (iMethCall *IMethodCall) appendParamVal(paramName string, param interface{}) {
	switch param := param.(type) {
	case bool:
		iMethCall.IParamValue = append(
			iMethCall.IParamValue,
			IParamValue{
				Name: paramName,
				Value: &Value{
					strconv.FormatBool(param),
				},
			},
		)
	case string:
		iMethCall.IParamValue = append(
			iMethCall.IParamValue,
			IParamValue{
				Name: paramName,
				Value: &Value{
					param,
				},
			},
		)
	case []string:
		var arry ValueArray
		for _, sub := range param {
			arry.Value = append(
				arry.Value,
				Value{sub},
			)
		}
		iMethCall.IParamValue = append(
			iMethCall.IParamValue,
			IParamValue{
				Name:       paramName,
				ValueArray: &arry,
			},
		)
	case *ClassName:
		iMethCall.IParamValue = append(
			iMethCall.IParamValue,
			IParamValue{
				Name:      paramName,
				ClassName: param,
			},
		)
	case *Class:
		iMethCall.IParamValue = append(
			iMethCall.IParamValue,
			IParamValue{
				Name:  paramName,
				Class: param,
			},
		)
	case *InstanceName:
		iMethCall.IParamValue = append(
			iMethCall.IParamValue,
			IParamValue{
				Name:         paramName,
				InstanceName: param,
			},
		)
	case *Instance:
		iMethCall.IParamValue = append(
			iMethCall.IParamValue,
			IParamValue{
				Name:     paramName,
				Instance: param,
			},
		)
	case *ValueNamedInstance:
		iMethCall.IParamValue = append(
			iMethCall.IParamValue,
			IParamValue{
				Name:               paramName,
				ValueNamedInstance: param,
			},
		)
	default:
	}
}

func (conn *WBEMConnection) doPostIMethodCall(method string, content []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s:%d/%s", conn.scheme, conn.host, conn.port, DefaultRequestURI), bytes.NewReader(content))
	if nil != err {
		return nil, err
	}
	req.SetBasicAuth(conn.username, conn.password)
	req.Header.Add("Content-Type", "application/xml; charset=\"utf-8\"")
	req.Header.Add("Host", fmt.Sprintf("%s:%d", conn.host, conn.port))
	req.Header.Add("Accept-Encoding", "identity")
	req.Header["TE"] = append(req.Header["TE"], "trailers")
	req.Header[HttpHdrOperation] = append(req.Header[HttpHdrOperation], "MethodCall")
	req.Header[HttpHdrMethod] = append(req.Header[HttpHdrMethod], method)
	req.Header[HttpHdrObject] = append(req.Header[HttpHdrObject], conn.namespace)
	res, err := conn.httpc.Do(req)
	if nil != err {
		return nil, err
	}
	if 200 != res.StatusCode {
		err = fmt.Errorf("HTTP_ERR - %d - %s", res.StatusCode, res.Status)
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func (conn *WBEMConnection) iMethodCall(call *IMethodCall) (*IMethodResponse, error) {
	if nil == call {
		return nil, conn.oops(ErrFailed, "")
	}
	var cim CIM = CIM{
		CIMVersion: "2.0",
		DTDVersion: "2.0",
		Message: &Message{
			ID:              "1001",
			ProtocolVersion: "1.0",
			SimpleReq: &SimpleReq{
				IMethodCall: call,
			},
		},
	}
	raw, err := xml.Marshal(&cim)
	if nil != err {
		return nil, err
	}
	raw = append([]byte(xml.Header), raw...)
	raw, err = conn.doPostIMethodCall(call.Name, raw)
	if nil != err {
		return nil, err
	}
	cim = CIM{}
	err = xml.Unmarshal(raw, &cim)
	if nil != err {
		loggerPrint(string(raw))
		return nil, err
	}
	if nil == cim.Message || nil == cim.Message.SimpleRsp || nil == cim.Message.SimpleRsp.IMethodResponse {
		return nil, conn.oops(ErrFailed, "")
	}
	return cim.Message.SimpleRsp.IMethodResponse, nil
}
