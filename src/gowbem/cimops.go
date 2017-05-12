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
	"strconv"
	"strings"
)

const (
	// The CIMOperation header shall be present in all CIM Operation Request and CIM Operation Response messages. It identifies the HTTP message as carrying a CIM operation request or response.
	//      CIMOperation = "CIMOperation" ":" ("MethodCall" | "MethodResponse")
	HttpHdrOperation = "CIMOperation"

	// The CIMExport header shall be present in all CIM export request and response messages. It identifies the HTTP message as carrying a CIM export method request or response.
	//      CIMExport = "CIMExport" ":" ("MethodRequest" | "MethodResponse")
	HttpHdrExport = "CIMExport"

	// The CIMProtocolVersion header may be present in any CIM message. The header identifies the version of the CIM operations over the HTTP specification in use by the sending entity.
	//      CIMProtocolVersion = "CIMProtocolVersion" ":" 1*DIGIT "." 1*DIGIT
	HttpHdrProtocolVersion = "CIMProtocolVersion"

	// The CIMMethod header shall be present in any CIM Operation Request message that contains a Simple Operation Request.
	// The name of the CIM method within a simple operation request is the value of the NAME attribute of the <METHODCALL> or <IMETHODCALL> element.
	//      CIMMethod = "CIMMethod" ":" MethodName
	//      MethodName = CIMIdentifier
	HttpHdrMethod = "CIMMethod"

	// The CIMObject header shall be present in any CIM Operation Request message that contains a Simple Operation Request.
	//      CIMObject = "CIMObject" ":" ObjectPath
	//      ObjectPath = CIMObjectPath
	HttpHdrObject = "CIMObject"

	// The CIMExportMethod header shall be present in any CIM export request message that contains a simple export request.
	//      CIMExportMethod = "CIMExportMethod" ":" ExportMethodName
	//      ExportMethodName = CIMIdentifier
	HttpHdrExportMethod = "CIMExportMethod"

	// The CIMBatch header shall be present in any CIM Operation Request message that contains a Multiple Operation Request.
	//      CIMBatch = "CIMBatch" ":"
	HttpHdrBatch = "CIMBatch"

	// The CIMExportBatch header shall be present in any CIM export request message that contains a multiple export request.
	//      CIMExportBatch = "CIMExportBatch" ":"
	HttpHdrExportBatch = "CIMExportBatch"

	// The CIMError header may be present in any HTTP response to a CIM message request that is not a CIM message response.
	//      CIMError = "CIMError" ":" cim-error
	//      cim-error = "unsupported-protocol-version" |
	//                  "multiple-requests-unsupported" |
	//                  "unsupported-cim-version" |
	//                  "unsupported-dtd-version" |
	//                  "request-not-valid" |
	//                  "request-not-well-formed" |
	//                  "request-not-loosely-valid" |
	//                  "header-mismatch" |
	//                  "unsupported-operation"
	HttpHdrError = "CIMError"

	// A CIM server may return a CIMRoleAuthenticate header as part of the 401 Unauthorized response along with the WWW-Authenticate header. The CIMRoleAuthenticate header must meet the challenge of indicating the CIM server policy on role credentials.
	//      challenge = "credentialrequired" | "credentialoptional" | "credentialnotrequired"
	HttpHdrAuthenticate = "CIMRoleAuthenticate"

	// The CIMRoleAuthorization header is supplied along with the normal authorization header that the CIM client populates to perform user authentication. If the CIM client needs to perform role assumption and the server challenge is credentialrequired, the CIMRoleAuthorization header must be supplied with the appropriate credentials. The credentials supplied as part of the CIMRoleAuthorization header must use the same scheme as those specified for the authorization header, as specified in RFC 2617. Therefore, both Basic and Digest authentication are possible for the role credential.
	HttpHdrRoleAuthorization = "CIMRoleAuthorization"

	// If a CIM product includes the CIMStatusCode trailer, it may also include the CIMStatusCodeDescription trailer.
	HttpHdrStatusCodeDescription = "CIMStatusCodeDescription"

	// The WBEMServerResponseTime header may be present in any CIM response message.
	HttpHdrServerResponseTime = "WBEMServerResponseTime"
)

const (
	DefaultRequestURI string = "cimom"
)

// <object>: VALUE.OBJECT|VALUE.OBJECTWITHLOCALPATH|VALUE.OBJECTWITHPATH
type Object struct {
	ValueObject              []ValueObject
	ValueObjectWithLocalPath []ValueObjectWithLocalPath
	ValueObjectWithPath      []ValueObjectWithPath
}

// <objectName>: (CLASSNAME|INSTANCENAME)
type ObjectName struct {
	ClassName    *ClassName
	InstanceName *InstanceName
}

// <propertyValue>: (VALUE|VALUE.ARRAY|VALUE.REFERENCE)
type PropertyValue struct {
	Value          *Value
	ValueArray     *ValueArray
	ValueReference *ValueReference
}

// <namedInstance>: VALUE.NAMEDINSTANCE
type NamedInstance ValueNamedInstance

// <instanceWithPath>: VALUE.INSTANCEWITHPATH
type InstanceWithPath ValueInstanceWithPath

// The GetClass operation returns a single CIM class from the target namespace:
//      GetClass <class>GetClass (
//           [IN] <className> ClassName,
//           [IN,OPTIONAL] boolean LocalOnly = true,
//           [IN,OPTIONAL] boolean IncludeQualifiers = true,
//           [IN,OPTIONAL] boolean IncludeClassOrigin = false,
//           [IN,OPTIONAL,NULL] string PropertyList [] = NULL
//      )
func (conn *WBEMConnection) GetClass(className *ClassName, localOnly bool, includeQualifiers bool, includeClassOrigin bool, propertyList []string) ([]Class, error) {
	if nil == className {
		return nil, conn.oops(ErrFailed)
	}
	iMethCall := newIMechCall("GetClass")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("ClassName", className)
	if true != localOnly {
		iMethCall.appendParamVal("LocalOnly", localOnly)
	}
	if true != includeQualifiers {
		iMethCall.appendParamVal("IncludeQualifiers", includeQualifiers)
	}
	if false != includeClassOrigin {
		iMethCall.appendParamVal("IncludeClassOrigin", includeQualifiers)
	}
	if nil != propertyList {
		iMethCall.appendParamVal("PropertyList", propertyList)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.Class, err
}

// The GetInstance operation returns a single CIM instance from the target namespace:
//      GetInstance <instance>GetInstance (
//           [IN] <instanceName> InstanceName,
//           [IN,OPTIONAL] boolean LocalOnly = true, (DEPRECATED)
//           [IN,OPTIONAL] boolean IncludeQualifiers = false, (DEPRECATED)
//           [IN,OPTIONAL] boolean IncludeClassOrigin = false,
//           [IN,OPTIONAL,NULL] string PropertyList [] = NULL
//      )
func (conn *WBEMConnection) GetInstance(instanceName *InstanceName, includeClassOrigin bool, propertyList []string) ([]Instance, error) {
	if nil == instanceName {
		return nil, conn.oops(ErrFailed)
	}
	iMethCall := newIMechCall("GetInstance")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("InstanceName", instanceName)
	if false != includeClassOrigin {
		iMethCall.appendParamVal("IncludeClassOrigin", includeClassOrigin)
	}
	if nil != propertyList {
		iMethCall.appendParamVal("PropertyList", propertyList)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.Instance, err
}

// The DeleteClass operation deletes a single CIM class from the target namespace:
//      void DeleteClass (
//           [IN] <className> ClassName
//      )
func (conn *WBEMConnection) DeleteClass(className *ClassName) error {
	if nil == className {
		return conn.oops(ErrFailed)
	}
	iMethCall := newIMechCall("DeleteClass")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("ClassName", className)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return conn.oops(i)
	}
	return nil
}

// The DeleteInstance operation deletes a single CIM instance from the target namespace.
//      void DeleteInstance (
//           [IN] <instanceName> InstanceName
//      )
func (conn *WBEMConnection) DeleteInstance(instanceName *InstanceName) error {
	if nil == instanceName {
		return conn.oops(ErrFailed)
	}
	iMethCall := newIMechCall("DeleteInstance")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("InstanceName", instanceName)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return conn.oops(i)
	}
	return nil
}

// The CreateClass operation creates a single CIM class in the target namespace. The class shall not already exist:
//      void CreateClass (
//           [IN] <class> NewClass
//      )
func (conn *WBEMConnection) CreateClass(newClass *Class) error {
	if nil == newClass {
		return conn.oops(ErrFailed)
	}
	iMethCall := newIMechCall("CreateClass")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("NewClass", newClass)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return conn.oops(i)
	}
	return nil
}

// The CreateInstance operation creates a single CIM Instance in the target namespace. The instance shall not already exist:
//      <instanceName>CreateInstance (
//           [IN] <instance> NewInstance
//      )
func (conn *WBEMConnection) CreateInstance(newInstance *Instance) error {
	if nil == newInstance {
		return conn.oops(ErrFailed)
	}
	iMethCall := newIMechCall("CreateInstance")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("NewInstance", newInstance)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return conn.oops(i)
	}
	return nil
}

// The ModifyClass operation modifies an existing CIM class in the target namespace. The class shall already exist:
//      void ModifyClass (
//           [IN] <class> ModifiedClass
//      )
func (conn *WBEMConnection) ModifyClass(modifiedClass *Class) error {
	if nil == modifiedClass {
		return conn.oops(ErrFailed)
	}
	iMethCall := newIMechCall("ModifyClass")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("ModifiedClass", modifiedClass)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return conn.oops(i)
	}
	return nil
}

// ModifyInstance
//      void ModifyInstance (
//           [IN] <namedInstance> ModifiedInstance,
//           [IN, OPTIONAL] boolean IncludeQualifiers = true, (DEPRECATED)
//           [IN, OPTIONAL, NULL] string propertyList[] = NULL
//      )
func (conn *WBEMConnection) ModifyInstance(modifiedInstance *ValueNamedInstance, propertyList []string) error {
	iMethCall := newIMechCall("ModifyInstance")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("ModifiedInstance", modifiedInstance)
	if nil != propertyList {
		iMethCall.appendParamVal("PropertyList", propertyList)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return conn.oops(i)
	}
	return nil
}

// The EnumerateClasses operation enumerates subclasses of a CIM class in the target namespace:
//      EnumerateClasses <class>*EnumerateClasses (
//           [IN,OPTIONAL,NULL] <className> ClassName=NULL,
//           [IN,OPTIONAL] boolean DeepInheritance = false,
//           [IN,OPTIONAL] boolean LocalOnly = true,
//           [IN,OPTIONAL] boolean IncludeQualifiers = true,
//           [IN,OPTIONAL] boolean IncludeClassOrigin = false
//      )
func (conn *WBEMConnection) EnumerateClasses(className *ClassName, deepInheritance bool, localOnly bool, includeQualifiers bool, includeClassOrigin bool) ([]Class, error) {
	iMethCall := newIMechCall("EnumerateClasses")
	iMethCall.appendNamespace(conn.namespace)
	if nil != className {
		iMethCall.appendParamVal("ClassName", className)
	}
	if false != deepInheritance {
		iMethCall.appendParamVal("DeepInheritance", deepInheritance)
	}
	if true != localOnly {
		iMethCall.appendParamVal("LocalOnly", localOnly)
	}
	if true != includeQualifiers {
		iMethCall.appendParamVal("IncludeQualifiers", includeQualifiers)
	}
	if false != includeClassOrigin {
		iMethCall.appendParamVal("IncludeClassOrigin", includeClassOrigin)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.Class, err
}

// The EnumerateClassNames operation enumerates the names of subclasses of a CIM class in the target namespace:
//      <className>*EnumerateClassNames (
//           [IN,OPTIONAL,NULL] <className> ClassName = NULL,
//           [IN,OPTIONAL] boolean DeepInheritance = false
//      )
func (conn *WBEMConnection) EnumerateClassNames(className *ClassName, deepInheritance bool) ([]Class, error) {
	iMethCall := newIMechCall("EnumerateInstances")
	iMethCall.appendNamespace(conn.namespace)
	if nil != className {
		iMethCall.appendParamVal("ClassName", className)
	}
	if false != deepInheritance {
		iMethCall.appendParamVal("DeepInheritance", deepInheritance)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.Class, err
}

// The EnumerateInstances operation enumerates instances of a CIM class in the target namespace, including instances in the class and any subclasses in accordance with the polymorphic nature of CIM objects:
//      <namedInstance>*EnumerateInstances (
//           [IN] <className> ClassName,
//           [IN,OPTIONAL] boolean LocalOnly = true, (DEPRECATED)
//           [IN,OPTIONAL] boolean DeepInheritance = true,
//           [IN,OPTIONAL] boolean IncludeQualifiers = false, (DEPRECATED)
//           [IN,OPTIONAL] boolean IncludeClassOrigin = false,
//           [IN,OPTIONAL,NULL] string PropertyList [] = NULL
//      )
func (conn *WBEMConnection) EnumerateInstances(className *ClassName, deepInheritance bool, includeClassOrigin bool, propertyList []string) ([]ValueNamedInstance, error) {
	if nil == className {
		return nil, conn.oops(ErrFailed)
	}
	iMethCall := newIMechCall("EnumerateInstances")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("ClassName", className)
	if true != deepInheritance {
		iMethCall.appendParamVal("DeepInheritance", deepInheritance)
	}
	if false != includeClassOrigin {
		iMethCall.appendParamVal("IncludeClassOrigin", includeClassOrigin)
	}
	if nil != propertyList {
		iMethCall.appendParamVal("PropertyList", propertyList)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.ValueNamedInstance, err
}

// The EnumerateInstanceNames operation enumerates the names (model paths) of the instances of a CIM class in the target namespace, including instances in the class and any subclasses in accordance with the polymorphic nature of CIM objects:
//      <instanceName>*EnumerateInstanceNames (
//           [IN] <className> ClassName
//      )
func (conn *WBEMConnection) EnumerateInstanceNames(className *ClassName) ([]InstanceName, error) {
	if nil == className {
		return nil, conn.oops(ErrFailed)
	}
	iMethCall := newIMechCall("EnumerateInstanceNames")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("ClassName", className)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.InstanceName, nil
}

// The ExecQuery operation executes a query against the target namespace:
//      <object>*ExecQuery (
//           [IN] string QueryLanguage,
//           [IN] string Query
//      )
func (conn *WBEMConnection) ExecQuery(queryLanguage string, query string) (*Object, error) {
	iMethCall := newIMechCall("ExecQuery")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("QueryLanguage", queryLanguage)
	iMethCall.appendParamVal("Query", query)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	var obj Object = Object{
		ValueObject:              iMethRes.IReturnValue.ValueObject,
		ValueObjectWithLocalPath: iMethRes.IReturnValue.ValueObjectWithLocalPath,
		ValueObjectWithPath:      iMethRes.IReturnValue.ValueObjectWithPath,
	}
	return &obj, nil
}

// The Associators operation enumerates CIM objects (classes or instances) associated with a particular source CIM object:
//      <objectWithPath>*Associators (
//           [IN] <objectName> ObjectName,
//           [IN,OPTIONAL,NULL] <className> AssocClass = NULL,
//           [IN,OPTIONAL,NULL] <className> ResultClass = NULL,
//           [IN,OPTIONAL,NULL] string Role = NULL,
//           [IN,OPTIONAL,NULL] string ResultRole = NULL,
//           [IN,OPTIONAL] boolean IncludeQualifiers = false, (DEPRECATED)
//           [IN,OPTIONAL] boolean IncludeClassOrigin = false,
//           [IN,OPTIONAL,NULL] string PropertyList [] = NULL
//      )
func (conn *WBEMConnection) Associators(objectName *ObjectName, assocClass *ClassName, resultClass *ClassName, role, resultRole *string, includeClassOrigin bool, propertyList []string) ([]ValueObjectWithPath, error) {
	iMethCall := newIMechCall("Associators")
	iMethCall.appendNamespace(conn.namespace)
	if nil != objectName.ClassName {
		iMethCall.appendParamVal("ObjectName", objectName.ClassName)
	} else if nil != objectName.InstanceName {
		iMethCall.appendParamVal("ObjectName", objectName.InstanceName)
	}
	if nil != assocClass {
		iMethCall.appendParamVal("AssocClass", assocClass)
	}
	if nil != resultClass {
		iMethCall.appendParamVal("ResultClass", resultClass)
	}
	if nil != role {
		iMethCall.appendParamVal("Role", *role)
	}
	if nil != resultRole {
		iMethCall.appendParamVal("ResultRole", *resultRole)
	}
	if false != includeClassOrigin {
		iMethCall.appendParamVal("IncludeClassOrigin", includeClassOrigin)
	}
	if nil != propertyList {
		iMethCall.appendParamVal("PropertyList", propertyList)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.ValueObjectWithPath, nil
}

// The AssociatorNames operation enumerates the names of CIM Objects (classes or instances) that are associated with a particular source CIM object:
//      <objectPath>*AssociatorNames (
//           [IN] <objectName> ObjectName,
//           [IN,OPTIONAL,NULL] <className> AssocClass = NULL,
//           [IN,OPTIONAL,NULL] <className> ResultClass = NULL,
//           [IN,OPTIONAL,NULL] string Role = NULL,
//           [IN,OPTIONAL,NULL] string ResultRole = NULL
//      )
func (conn *WBEMConnection) AssociatorNames(objectName *ObjectName, assocClass *ClassName, resultClass *ClassName, role, resultRole *string) ([]ObjectPath, error) {
	iMethCall := newIMechCall("AssociatorNames")
	iMethCall.appendNamespace(conn.namespace)
	if nil != objectName.ClassName {
		iMethCall.appendParamVal("ObjectName", objectName.ClassName)
	} else if nil != objectName.InstanceName {
		iMethCall.appendParamVal("ObjectName", objectName.InstanceName)
	}
	if nil != assocClass {
		iMethCall.appendParamVal("AssocClass", assocClass)
	}
	if nil != resultClass {
		iMethCall.appendParamVal("ResultClass", resultClass)
	}
	if nil != role {
		iMethCall.appendParamVal("Role", *role)
	}
	if nil != resultRole {
		iMethCall.appendParamVal("ResultRole", *resultRole)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.ObjectPath, nil
}

// The References operation enumerates the association objects that refer to a particular target CIM object (class or instance).
//      <objectWithPath>*References (
//           [IN] <objectName> ObjectName,
//           [IN,OPTIONAL,NULL] <className> ResultClass = NULL,
//           [IN,OPTIONAL,NULL] string Role = NULL,
//           [IN,OPTIONAL] boolean IncludeQualifiers = false, (DEPRECATED)
//           [IN,OPTIONAL] boolean IncludeClassOrigin = false,
//           [IN,OPTIONAL,NULL] string PropertyList [] = NULL
//      )
func (conn *WBEMConnection) References(objectName *ObjectName, resultClass *ClassName, role *string, includeClassOrigin bool, propertyList []string) ([]ValueObjectWithPath, error) {
	iMethCall := newIMechCall("References")
	iMethCall.appendNamespace(conn.namespace)
	if nil != objectName.ClassName {
		iMethCall.appendParamVal("ObjectName", objectName.ClassName)
	} else if nil != objectName.InstanceName {
		iMethCall.appendParamVal("ObjectName", objectName.InstanceName)
	}
	if nil != resultClass {
		iMethCall.appendParamVal("ResultClass", resultClass)
	}
	if nil != role {
		iMethCall.appendParamVal("Role", *role)
	}
	if false != includeClassOrigin {
		iMethCall.appendParamVal("IncludeClassOrigin", includeClassOrigin)
	}
	if nil != propertyList {
		iMethCall.appendParamVal("PropertyList", propertyList)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.ValueObjectWithPath, nil
}

// The ReferenceNames operation enumerates the association objects that refer to a particular target CIM object (class or instance):
//      <objectPath>*ReferenceNames (
//           [IN] <objectName> ObjectName,
//           [IN,OPTIONAL,NULL] <className> ResultClass = NULL,
//           [IN,OPTIONAL,NULL] string Role = NULL
//      )
func (conn *WBEMConnection) ReferenceNames(objectName *ObjectName, assocClass *ClassName, role *string) ([]ObjectPath, error) {
	iMethCall := newIMechCall("ReferenceNames")
	iMethCall.appendNamespace(conn.namespace)
	if nil != objectName.ClassName {
		iMethCall.appendParamVal("ObjectName", objectName.ClassName)
	} else if nil != objectName.InstanceName {
		iMethCall.appendParamVal("ObjectName", objectName.InstanceName)
	}
	if nil != assocClass {
		iMethCall.appendParamVal("AssocClass", assocClass)
	}
	if nil != role {
		iMethCall.appendParamVal("Role", *role)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.ObjectPath, nil
}

// The InvokeMethod operation enumerates the association objects that refer to a particular target CIM object (class or instance):
//      <int []ParamValues>*InvokeMethod (
//           [IN] <objectName> ObjectName,
//           [IN] <methodName> string,
//           [IN] <paramValues> IParamValues
//      )
func (conn *WBEMConnection) InvokeMethod(objectName *ObjectName, methodName string, paramValues []IParamValue) (int, []ParamValue, error) {
	methCall := &MethodCall{Name: methodName}

	var ns []Namespace = []Namespace{}
	for _, sub := range strings.Split(conn.namespace, "/") {
		ns = append(ns, Namespace{Name: sub})
	}
	namespacePath := &LocalNamespacePath{ns}

	if nil != objectName.ClassName {
		methCall.LocalClassPath = &LocalClassPath{LocalNamespacePath: namespacePath, ClassName: objectName.ClassName}
	} else if nil != objectName.InstanceName {
		methCall.LocalInstancePath = &LocalInstancePath{LocalNamespacePath: namespacePath, InstanceName: objectName.InstanceName}
	}
	methCall.ParamValue = paramValues

	methRes, err := conn.methodCall(methCall)
	if nil != err {
		return -1, nil, err
	}
	if nil != methRes.Error {
		i, _ := strconv.Atoi(methRes.Error.Code)
		return i, nil, conn.oops(i)
	}
	if nil == methRes.ReturnValue {
		return -1, nil, conn.oops(0)
	}
	retCode, err2 := strconv.Atoi(methRes.ReturnValue.Value.Value)
	return retCode, methRes.ParamValue, err2
}

// The GetProperty operation retrieves a single property value from a CIM instance in the target namespace:
//      <propertyValue>GetProperty (
//           [IN] <instanceName> InstanceName,
//           [IN] string PropertyName
//      )
func (conn *WBEMConnection) GetProperty(instanceName *InstanceName, propertyName string) (*PropertyValue, error) {
	iMethCall := newIMechCall("GetProperty")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("InstanceName", instanceName)
	iMethCall.appendParamVal("PropertyName", propertyName)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	propVal := PropertyValue{
		ValueArray:     iMethRes.IReturnValue.ValueArray,
		ValueReference: iMethRes.IReturnValue.ValueReference,
	}
	if 0 < len(iMethRes.IReturnValue.Value) {
		propVal.Value = &iMethRes.IReturnValue.Value[0]
	}
	return &propVal, nil
}

// The SetProperty operation sets a single property value in a CIM instance in the target namespace:
//      void SetProperty (
//           [IN] <instanceName> InstanceName,
//           [IN] string PropertyName,
//           [IN,OPTIONAL,NULL] <propertyValue> NewValue = NULL
//      )
func (conn *WBEMConnection) SetProperty(instanceName *InstanceName, propertyName string, newValue *PropertyValue) error {
	iMethCall := newIMechCall("SetProperty")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("InstanceName", instanceName)
	iMethCall.appendParamVal("PropertyName", propertyName)
	if nil != newValue {
		iMethCall.appendParamVal("NewValue", newValue)
	}
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return conn.oops(i)
	}
	return nil
}

// The GetQualifier operation retrieves a single qualifier declaration from the target namespace.
//      <qualifierDecl>GetQualifier (
//           [IN] string QualifierName
//      )
func (conn *WBEMConnection) GetQualifier(qualifierName string) (*QualifierDeclaration, error) {
	iMethCall := newIMechCall("GetQualifier")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("QualifierName", qualifierName)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	if 0 == len(iMethRes.IReturnValue.QualifierDeclaration) {
		return nil, nil
	}
	return &iMethRes.IReturnValue.QualifierDeclaration[0], nil
}

// The SetQualifier operation creates or updates a single qualifier declaration in the target namespace. If the qualifier declaration already exists, it is overwritten:
//      void SetQualifier (
//           [IN] <qualifierDecl> QualifierDeclaration
//      )
func (conn *WBEMConnection) SetQualifier(qualifierDecl *QualifierDeclaration) error {
	iMethCall := newIMechCall("SetQualifier")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("QualifierDeclaration", qualifierDecl)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return conn.oops(i)
	}
	return nil
}

// The DeleteQualifier operation deletes a single qualifier declaration from the target namespace.
//      void DeleteQualifier (
//           [IN] string QualifierName
//      )
func (conn *WBEMConnection) DeleteQualifier(qualifierName string) error {
	iMethCall := newIMechCall("DeleteQualifier")
	iMethCall.appendNamespace(conn.namespace)
	iMethCall.appendParamVal("QualifierName", qualifierName)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return conn.oops(i)
	}
	return nil
}

// The EnumerateQualifiers operation enumerates qualifier declarations from the target namespace.
//      <qualifierDecl>*EnumerateQualifiers (
//      )
func (conn *WBEMConnection) EnumerateQualifiers() ([]QualifierDeclaration, error) {
	iMethCall := newIMechCall("EnumerateQualifiers")
	iMethCall.appendNamespace(conn.namespace)
	iMethRes, err := conn.iMethodCall(iMethCall)
	if nil != err {
		return nil, err
	}
	if nil != iMethRes.Error {
		i, _ := strconv.Atoi(iMethRes.Error.Code)
		return nil, conn.oops(i)
	}
	if nil == iMethRes.IReturnValue {
		return nil, nil
	}
	return iMethRes.IReturnValue.QualifierDeclaration, nil
}
