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

// <!--
// **************************************************
// Top-level element
// **************************************************
// -->

// <!ELEMENT CIM (MESSAGE | DECLARATION)>
// <!ATTLIST CIM
//     CIMVERSION CDATA #REQUIRED
//     DTDVERSION CDATA #REQUIRED
// >
type CIM struct {
	CIMVersion  string       `xml:"CIMVERSION,attr" json:",omitempty"`
	DTDVersion  string       `xml:"DTDVERSION,attr" json:",omitempty"`
	Message     *Message     `xml:"MESSAGE" json:",omitempty"`
	Declaration *Declaration `xml:"DECLARATION" json:",omitempty"`
}

// <!--
// **************************************************
// Object declaration elements
// **************************************************
// -->

// <!ELEMENT DECLARATION (DECLGROUP | DECLGROUP.WITHNAME | DECLGROUP.WITHPATH)+>
type Declaration struct {
	DeclGroup         []DeclGroup         `xml:"DECLGROUP" json:",omitempty"`
	DeclGroupWithName []DeclGroupWithName `xml:"DECLGROUP.WITHNAME" json:",omitempty"`
	DeclGroupWithPath []DeclGroupWithPath `xml:"DECLGROUP.WITHPATH" json:",omitempty"`
}

// <!ELEMENT DECLGROUP ((LOCALNAMESPACEPATH | NAMESPACEPATH)?, QUALIFIER.DECLARATION*, VALUE.OBJECT*)>
type DeclGroup struct {
	LocalNamespacePath   *LocalNamespacePath    `xml:"LOCALNAMESPACEPATH" json:",omitempty"`
	NamespacePath        *NamespacePath         `xml:"NAMESPACEPATH" json:",omitempty"`
	QualifierDeclaration []QualifierDeclaration `xml:"QUALIFIER.DECLARATION" json:",omitempty"`
	ValueObject          []ValueObject          `xml:"VALUE.OBJECT" json:",omitempty"`
}

// <!ELEMENT DECLGROUP.WITHNAME ((LOCALNAMESPACEPATH | NAMESPACEPATH)?, QUALIFIER.DECLARATION*, VALUE.NAMEDOBJECT*)>
type DeclGroupWithName struct {
	LocalNamespacePath   *LocalNamespacePath    `xml:"LOCALNAMESPACEPATH" json:",omitempty"`
	NamespacePath        *NamespacePath         `xml:"NAMESPACEPATH" json:",omitempty"`
	QualifierDeclaration []QualifierDeclaration `xml:"QUALIFIER.DECLARATION" json:",omitempty"`
	ValueNamedObject     []ValueNamedObject     `xml:"VALUE.NAMEDOBJECT" json:",omitempty"`
}

// <!ELEMENT DECLGROUP.WITHPATH (VALUE.OBJECTWITHPATH | VALUE.OBJECTWITHLOCALPATH)*>
type DeclGroupWithPath struct {
	ValueObjectWithPath      *ValueObjectWithPath      `xml:"VALUE.OBJECTWITHPATH" json:",omitempty"`
	ValueObjectWithLocalPath *ValueObjectWithLocalPath `xml:"VALUE.OBJECTWITHLOCALPATH" json:",omitempty"`
}

// <!ELEMENT QUALIFIER.DECLARATION (SCOPE?, (VALUE | VALUE.ARRAY)?)>
// <!ATTLIST QUALIFIER.DECLARATION
//     %CIMName;
//     %CIMType; #REQUIRED
//     ISARRAY (true|false) #IMPLIED
//     %ArraySize;
//     %QualifierFlavor;
// >
type QualifierDeclaration struct {
	Name         string      `xml:"NAME,attr,omitempty" json:",omitempty"`
	Type         string      `xml:"TYPE,attr" json:",omitempty"`
	IsArray      string      `xml:"ISARRAY,attr,omitempty" json:",omitempty"`
	ArraySize    string      `xml:"ARRAYSIZE,attr,omitempty" json:",omitempty"`
	Overridable  string      `xml:"OVERRIDABLE,attr,omitempty" json:",omitempty"`
	ToSubClass   string      `xml:"TOSUBCLASS,attr,omitempty" json:",omitempty"`
	ToInstance   string      `xml:"TOINSTANCE,attr,omitempty" json:",omitempty"`
	Translatable string      `xml:"TRANSLATABLE,attr,omitempty" json:",omitempty"`
	Scope        *Scope      `xml:"SCOPE" json:",omitempty"`
	Value        *Value      `xml:"VALUE" json:",omitempty"`
	ValueArray   *ValueArray `xml:"VALUE.ARRAY" json:",omitempty"`
}

// <!ELEMENT SCOPE EMPTY>
// <!ATTLIST SCOPE
//     CLASS (true | false) "false"
//     ASSOCIATION (true | false) "false"
//     REFERENCE (true | false) "false"
//     PROPERTY (true | false) "false"
//     METHOD (true | false) "false"
//     PARAMETER (true | false) "false"
//     INDICATION (true | false) "false"
// >
type Scope struct {
	Class       string `xml:"CLASS,attr,omitempty" json:",omitempty"`
	Association string `xml:"ASSOCIATION,attr,omitempty" json:",omitempty"`
	Reference   string `xml:"REFERENCE,attr,omitempty" json:",omitempty"`
	Property    string `xml:"PROPERTY,attr,omitempty" json:",omitempty"`
	Method      string `xml:"METHOD,attr,omitempty" json:",omitempty"`
	Parameter   string `xml:"PARAMETER,attr,omitempty" json:",omitempty"`
	Indication  string `xml:"INDICATION,attr,omitempty" json:",omitempty"`
}

// <!--
// **************************************************
// Object Value elements
// **************************************************
// -->

// <!ELEMENT VALUE (#PCDATA)>
//type Value string
type Value struct {
	Value string `xml:",innerxml" json:",omitempty"`
}

// <!ELEMENT VALUE.ARRAY (VALUE | VALUE.NULL)*>
type ValueArray struct {
	Value     []Value     `xml:"VALUE" json:",omitempty"`
	ValueNull []ValueNull `xml:"VALUE.NULL" json:",omitempty"`
}

// <!ELEMENT VALUE.REFERENCE (CLASSPATH | LOCALCLASSPATH | CLASSNAME | INSTANCEPATH | LOCALINSTANCEPATH | INSTANCENAME)>
type ValueReference struct {
	ClassPath         *ClassPath         `xml:"CLASSPATH" json:",omitempty"`
	LocalClassPath    *LocalClassPath    `xml:"LOCALCLASSPATH" json:",omitempty"`
	ClassName         *ClassName         `xml:"CLASSNAME" json:",omitempty"`
	InstancePath      *InstancePath      `xml:"INSTANCEPATH" json:",omitempty"`
	LocalInstancePath *LocalInstancePath `xml:"LOCALINSTANCEPATH" json:",omitempty"`
	InstanceName      *InstanceName      `xml:"INSTANCENAME" json:",omitempty"`
}

// <!ELEMENT VALUE.REFARRAY (VALUE.REFERENCE | VALUE.NULL)*>
type ValueRefArray struct {
	ValueReference []ValueReference `xml:"ValueReference" json:",omitempty"`
	ValueNull      []ValueNull      `xml:"VALUE.NULL" json:",omitempty"`
}

// <!ELEMENT VALUE.OBJECT (CLASS | INSTANCE)>
type ValueObject struct {
	Class    *Class    `xml:"CLASS" json:",omitempty"`
	Instance *Instance `xml:"INSTANCE" json:",omitempty"`
}

// <!ELEMENT VALUE.NAMEDINSTANCE (INSTANCENAME, INSTANCE)>
type ValueNamedInstance struct {
	InstanceName *InstanceName `xml:"INSTANCENAME" json:",omitempty"`
	Instance     *Instance     `xml:"INSTANCE" json:",omitempty"`
}

// <!ELEMENT VALUE.NAMEDOBJECT (CLASS | (INSTANCENAME, INSTANCE))>
type ValueNamedObject struct {
	Class        *Class        `xml:"CLASS" json:",omitempty"`
	InstanceName *InstanceName `xml:"INSTANCENAME" json:",omitempty"`
	Instance     *Instance     `xml:"INSTANCE" json:",omitempty"`
}

// <!ELEMENT VALUE.OBJECTWITHPATH ((CLASSPATH, CLASS) | (INSTANCEPATH, INSTANCE))>
type ValueObjectWithPath struct {
	ClassPath    *ClassPath    `xml:"CLASSPATH" json:",omitempty"`
	Class        *Class        `xml:"CLASS" json:",omitempty"`
	InstancePath *InstancePath `xml:"INSTANCEPATH" json:",omitempty"`
	Instance     *Instance     `xml:"INSTANCE" json:",omitempty"`
}

// <!ELEMENT VALUE.OBJECTWITHLOCALPATH ((LOCALCLASSPATH, CLASS) | (LOCALINSTANCEPATH, INSTANCE))>
type ValueObjectWithLocalPath struct {
	LocalClassPath    *LocalClassPath    `xml:"LOCALCLASSPATH" json:",omitempty"`
	Class             *Class             `xml:"CLASS" json:",omitempty"`
	LocalInstancePath *LocalInstancePath `xml:"LOCALINSTANCEPATH" json:",omitempty"`
	Instance          *Instance          `xml:"INSTANCE" json:",omitempty"`
}

// <!ELEMENT VALUE.NULL EMPTY>
type ValueNull struct {
}

// <!ELEMENT VALUE.INSTANCEWITHPATH (INSTANCEPATH, INSTANCE)>
type ValueInstanceWithPath struct {
	InstancePath *InstancePath `xml:"INSTANCEPATH" json:",omitempty"`
	Instance     *Instance     `xml:"INSTANCE" json:",omitempty"`
}

// <!--
// **************************************************
// Object naming and locating elements
// **************************************************
// -->

// <!ELEMENT NAMESPACEPATH (HOST, LOCALNAMESPACEPATH)>
type NamespacePath struct {
	Host               *Host               `xml:"HOST" json:",omitempty"`
	LocalNamespacePath *LocalNamespacePath `xml:"LOCALNAMESPACEPATH" json:",omitempty"`
}

// <!ELEMENT LOCALNAMESPACEPATH (NAMESPACE+)>
type LocalNamespacePath struct {
	Namespace []Namespace `xml:"NAMESPACE" json:",omitempty"`
}

// <!ELEMENT HOST (#PCDATA)>
//type Host string
type Host struct {
	Host string `xml:",innerxml" json:",omitempty"`
}

// <!ELEMENT NAMESPACE EMPTY>
// <!ATTLIST NAMESPACE
//     %CIMName;
// >
type Namespace struct {
	Name string `xml:"NAME,attr,omitempty" json:",omitempty"`
}

// <!ELEMENT CLASSPATH (NAMESPACEPATH, CLASSNAME)>
type ClassPath struct {
	NamespacePath *NamespacePath `xml:"NAMESPACEPATH" json:",omitempty"`
	ClassName     *ClassName     `xml:"CLASSNAME" json:",omitempty"`
}

// <!ELEMENT LOCALCLASSPATH (LOCALNAMESPACEPATH, CLASSNAME)>
type LocalClassPath struct {
	LocalNamespacePath *LocalNamespacePath `xml:"LOCALNAMESPACEPATH" json:",omitempty"`
	ClassName          *ClassName          `xml:"CLASSNAME" json:",omitempty"`
}

// <!ELEMENT CLASSNAME EMPTY>
// <!ATTLIST CLASSNAME
//     %CIMName;
// >
type ClassName struct {
	Name string `xml:"NAME,attr,omitempty" json:",omitempty"`
}

// <!ELEMENT INSTANCEPATH (NAMESPACEPATH, INSTANCENAME)>
type InstancePath struct {
	NamespacePath *NamespacePath `xml:"NAMESPACEPATH" json:",omitempty"`
	InstanceName  *InstanceName  `xml:"INSTANCENAME" json:",omitempty"`
}

// <!ELEMENT LOCALINSTANCEPATH (LOCALNAMESPACEPATH, INSTANCENAME)>
type LocalInstancePath struct {
	LocalNamespacePath *LocalNamespacePath `xml:"LOCALNAMESPACEPATH" json:",omitempty"`
	InstanceName       *InstanceName       `xml:"INSTANCENAME" json:",omitempty"`
}

// <!ELEMENT INSTANCENAME (KEYBINDING* | KEYVALUE? | VALUE.REFERENCE?)>
// <!ATTLIST INSTANCENAME
//     %ClassName;
// >
type InstanceName struct {
	ClassName      string          `xml:"CLASSNAME,attr,omitempty",json:",omitempty" json:",omitempty"`
	KeyBinding     []KeyBinding    `xml:"KEYBINDING",json:",omitempty" json:",omitempty"`
	KeyValue       *KeyValue       `xml:"KEYVALUE",json:",omitempty" json:",omitempty"`
	ValueReference *ValueReference `xml:"VALUE.REFERENCE",json:",omitempty" json:",omitempty"`
}

// <!ELEMENT OBJECTPATH (INSTANCEPATH | CLASSPATH)>
type ObjectPath struct {
	InstancePath *InstancePath `xml:"INSTANCEPATH" json:",omitempty"`
	ClassPath    *ClassPath    `xml:"CLASSPATH" json:",omitempty"`
}

// <!ELEMENT KEYBINDING (KEYVALUE | VALUE.REFERENCE)>
// <!ATTLIST KEYBINDING
//     %CIMName;
// >
type KeyBinding struct {
	Name           string          `xml:"NAME,attr,omitempty",json:",omitempty" json:",omitempty"`
	KeyValue       *KeyValue       `xml:"KEYVALUE",json:",omitempty" json:",omitempty"`
	ValueReference *ValueReference `xml:"VALUE.REFERENCE",json:",omitempty" json:",omitempty"`
}

// <!ELEMENT KEYVALUE (#PCDATA)>
// <!ATTLIST KEYVALUE
//     VALUETYPE (string | boolean | numeric) "string"
//     %CIMType; #REQUIRED
// >
type KeyValue struct {
	ValueType string `xml:"VALUETYPE,attr,omitempty" json:",omitempty"`
	Type      string `xml:"TYPE,attr,omitempty" json:",omitempty"`
	KeyValue  string `xml:",innerxml" json:",omitempty"`
}

// <!--
// **************************************************
// Object definition elements
// **************************************************
// -->
// <!ELEMENT CLASS (QUALIFIER*, (PROPERTY | PROPERTY.ARRAY | PROPERTY.REFERENCE)*, METHOD*)>
// <!ATTLIST CLASS
//     %CIMName;
//     %SuperClass;
// >
type Class struct {
	Name              string              `xml:"NAME,attr,omitempty" json:",omitempty"`
	SuperClass        string              `xml:"SUPERCLASS,attr,omitempty" json:",omitempty"`
	Qualifier         []Qualifier         `xml:"QUALIFIER" json:",omitempty"`
	Property          []Property          `xml:"PROPERTY" json:",omitempty"`
	PropertyArray     []PropertyArray     `xml:"PROPERTY.ARRAY" json:",omitempty"`
	PropertyReference []PropertyReference `xml:"PROPERTY.REFERENCE" json:",omitempty"`
	Method            []Method            `xml:"METHOD" json:",omitempty"`
}

// <!ELEMENT INSTANCE (QUALIFIER*, (PROPERTY | PROPERTY.ARRAY | PROPERTY.REFERENCE)*)>
// <!ATTLIST INSTANCE
//     %ClassName;
//     xml:lang NMTOKEN #IMPLIED
// >
type Instance struct {
	ClassName         string              `xml:"CLASSNAME,attr,omitempty" json:",omitempty"`
	Qualifier         []Qualifier         `xml:"QUALIFIER" json:",omitempty"`
	Property          []Property          `xml:"PROPERTY" json:",omitempty"`
	PropertyArray     []PropertyArray     `xml:"PROPERTY.ARRAY" json:",omitempty"`
	PropertyReference []PropertyReference `xml:"PROPERTY.REFERENCE" json:",omitempty"`
}

// <!ELEMENT QUALIFIER ((VALUE | VALUE.ARRAY)?)>
// <!ATTLIST QUALIFIER
//     %CIMName;
//     %CIMType; #REQUIRED
//     %Propagated;
//     %QualifierFlavor;
//     xml:lang NMTOKEN #IMPLIED
// >
type Qualifier struct {
	Name         string      `xml:"NAME,attr,omitempty" json:",omitempty"`
	Type         string      `xml:"TYPE,attr" json:",omitempty"`
	Propagated   string      `xml:"PROPAGATED,attr,omitempty" json:",omitempty"`
	Overridable  string      `xml:"OVERRIDABLE,attr,omitempty" json:",omitempty"`
	ToSubClass   string      `xml:"TOSUBCLASS,attr,omitempty" json:",omitempty"`
	ToInstance   string      `xml:"TOINSTANCE,attr,omitempty" json:",omitempty"`
	Translatable string      `xml:"TRANSLATABLE,attr,omitempty" json:",omitempty"`
	Value        *Value      `xml:"VALUE" json:",omitempty"`
	ValueArray   *ValueArray `xml:"VALUE.ARRAY" json:",omitempty"`
}

// <!ELEMENT PROPERTY (QUALIFIER*, VALUE?)>
// <!ATTLIST PROPERTY
//     %CIMName;
//     %CIMType; #REQUIRED
//     %ClassOrigin;
//     %Propagated;
//     %EmbeddedObject;
//     xml:lang NMTOKEN #IMPLIED
// >
type Property struct {
	Name           string      `xml:"NAME,attr,omitempty" json:",omitempty"`
	Type           string      `xml:"TYPE,attr" json:",omitempty"`
	ClassOrigin    string      `xml:"CLASSORIGIN,attr,omitempty" json:",omitempty"`
	Propagated     string      `xml:"PROPAGATED,attr,omitempty" json:",omitempty"`
	EmbeddedObject string      `xml:"EmbeddedObject,attr,omitempty" json:",omitempty"`
	Qualifier      []Qualifier `xml:"QUALIFIER" json:",omitempty"`
	Value          *Value      `xml:"VALUE" json:",omitempty"`
}

// <!ELEMENT PROPERTY.ARRAY (QUALIFIER*, VALUE.ARRAY?)>
// <!ATTLIST PROPERTY.ARRAY
//     %CIMName;
//     %CIMType; #REQUIRED
//     %ArraySize;
//     %ClassOrigin;
//     %Propagated;
//     %EmbeddedObject;
//     xml:lang NMTOKEN #IMPLIED
// >
type PropertyArray struct {
	Name           string      `xml:"NAME,attr,omitempty" json:",omitempty"`
	Type           string      `xml:"TYPE,attr" json:",omitempty"`
	ArraySize      string      `xml:"ARRAYSIZE,attr,omitempty" json:",omitempty"`
	ClassOrigin    string      `xml:"CLASSORIGIN,attr,omitempty" json:",omitempty"`
	Propagated     string      `xml:"PROPAGATED,attr,omitempty" json:",omitempty"`
	EmbeddedObject string      `xml:"EmbeddedObject,attr,omitempty" json:",omitempty"`
	Qualifier      []Qualifier `xml:"QUALIFIER" json:",omitempty"`
	ValueArray     *ValueArray `xml:"VALUE.ARRAY" json:",omitempty"`
}

// <!ELEMENT PROPERTY.REFERENCE (QUALIFIER*, VALUE.REFERENCE?)>
// <!ATTLIST PROPERTY.REFERENCE
//     %CIMName;
//     %ReferenceClass;
//     %ClassOrigin;
//     %Propagated;
// >
type PropertyReference struct {
	Name           string          `xml:"NAME,attr,omitempty" json:",omitempty"`
	ReferenceClass string          `xml:"REFERENCECLASS,attr,omitempty" json:",omitempty"`
	ClassOrigin    string          `xml:"CLASSORIGIN,attr,omitempty" json:",omitempty"`
	Propagated     string          `xml:"PROPAGATED,attr,omitempty" json:",omitempty"`
	Qualifier      []Qualifier     `xml:"QUALIFIER" json:",omitempty"`
	ValueReference *ValueReference `xml:"VALUE.REFERENCE" json:",omitempty"`
}

// <!ELEMENT METHOD (QUALIFIER*, (PARAMETER | PARAMETER.REFERENCE | PARAMETER.ARRAY | PARAMETER.REFARRAY)*)>
// <!ATTLIST METHOD
//     %CIMName;
//     %CIMType; #IMPLIED
//     %ClassOrigin;
//     %Propagated;
// >
type Method struct {
	Name               string               `xml:"NAME,attr,omitempty" json:",omitempty"`
	Type               string               `xml:"TYPE,attr,omitempty" json:",omitempty"`
	ReferenceClass     string               `xml:"REFERENCECLASS,attr,omitempty" json:",omitempty"`
	ClassOrigin        string               `xml:"CLASSORIGIN,attr,omitempty" json:",omitempty"`
	Propagated         string               `xml:"PROPAGATED,attr,omitempty" json:",omitempty"`
	Qualifier          []Qualifier          `xml:"QUALIFIER" json:",omitempty"`
	Parameter          []Parameter          `xml:"PARAMETER" json:",omitempty"`
	ParameterReference []ParameterReference `xml:"PARAMETER.REFERENCE" json:",omitempty"`
	ParameterArray     []ParameterArray     `xml:"PARAMETER.ARRAY" json:",omitempty"`
	ParameterRefArray  []ParameterRefArray  `xml:"PARAMETER.REFARRAY" json:",omitempty"`
}

// <!ELEMENT PARAMETER (QUALIFIER*)>
// <!ATTLIST PARAMETER
//     %CIMName;
//     %CIMType; #REQUIRED
// >
type Parameter struct {
	Name      string      `xml:"NAME,attr,omitempty" json:",omitempty"`
	Type      string      `xml:"TYPE,attr" json:",omitempty"`
	Qualifier []Qualifier `xml:"QUALIFIER" json:",omitempty"`
}

// <!ELEMENT PARAMETER.REFERENCE (QUALIFIER*)>
// <!ATTLIST PARAMETER.REFERENCE
//     %CIMName;
//     %ReferenceClass;
// >
type ParameterReference struct {
	Name           string      `xml:"NAME,attr,omitempty" json:",omitempty"`
	ReferenceClass string      `xml:"REFERENCECLASS,attr,omitempty" json:",omitempty"`
	Qualifier      []Qualifier `xml:"QUALIFIER" json:",omitempty"`
}

// <!ELEMENT PARAMETER.ARRAY (QUALIFIER*)>
// <!ATTLIST PARAMETER.ARRAY
//     %CIMName;
//     %CIMType; #REQUIRED
//     %ArraySize;
// >
type ParameterArray struct {
	Name      string      `xml:"NAME,attr,omitempty" json:",omitempty"`
	Type      string      `xml:"TYPE,attr" json:",omitempty"`
	ArraySize string      `xml:"ARRAYSIZE,attr,omitempty" json:",omitempty"`
	Qualifier []Qualifier `xml:"QUALIFIER" json:",omitempty"`
}

// <!ELEMENT PARAMETER.REFARRAY (QUALIFIER*)>
// <!ATTLIST PARAMETER.REFARRAY
//     %CIMName;
//     %ReferenceClass;
//     %ArraySize;
// >
type ParameterRefArray struct {
	Name           string      `xml:"NAME,attr,omitempty" json:",omitempty"`
	ReferenceClass string      `xml:"REFERENCECLASS,attr,omitempty" json:",omitempty"`
	ArraySize      string      `xml:"ARRAYSIZE,attr,omitempty" json:",omitempty"`
	Qualifier      []Qualifier `xml:"QUALIFIER" json:",omitempty"`
}

// <!--
// **************************************************
// Message elements
// **************************************************
// -->

// <!ELEMENT MESSAGE (SIMPLEREQ | MULTIREQ | SIMPLERSP | MULTIRSP | SIMPLEEXPREQ | MULTIEXPREQ | SIMPLEEXPRSP | MULTIEXPRSP)>
// <!ATTLIST MESSAGE
//     ID CDATA #REQUIRED
//     PROTOCOLVERSION CDATA #REQUIRED
// >
type Message struct {
	ID              string        `xml:"ID,attr" json:",omitempty"`
	ProtocolVersion string        `xml:"PROTOCOLVERSION,attr" json:",omitempty"`
	SimpleReq       *SimpleReq    `xml:"SIMPLEREQ" json:",omitempty"`
	MultiReq        *MultiReq     `xml:"MULTIREQ" json:",omitempty"`
	SimpleRsp       *SimpleRsp    `xml:"SIMPLERSP" json:",omitempty"`
	MultiRsp        *MultiRsp     `xml:"MULTIRSP" json:",omitempty"`
	SimpleExpReq    *SimpleExpReq `xml:"SIMPLEEXPREQ" json:",omitempty"`
	MultiExpReq     *MultiExpReq  `xml:"MULTIEXPREQ" json:",omitempty"`
	SimpleExpRsp    *SimpleExpRsp `xml:"SIMPLEEXPRSP" json:",omitempty"`
	MultiExpRsp     *MultiExpRsp  `xml:"MULTIEXPRSP" json:",omitempty"`
}

// <!ELEMENT MULTIREQ (SIMPLEREQ, SIMPLEREQ+)>
type MultiReq struct {
	SimpleReq []SimpleReq `xml:"SIMPLEREQ" json:",omitempty"`
}

// <!ELEMENT SIMPLEREQ (CORRELATOR*, (METHODCALL | IMETHODCALL))>
type SimpleReq struct {
	Correlator  []Correlator `xml:"CORRELATOR" json:",omitempty"`
	MethodCall  *MethodCall  `xml:"METHODCALL" json:",omitempty"`
	IMethodCall *IMethodCall `xml:"IMETHODCALL" json:",omitempty"`
}

// <!ELEMENT METHODCALL ((LOCALCLASSPATH | LOCALINSTANCEPATH), PARAMVALUE*)>
// <!ATTLIST METHODCALL
//     %CIMName;
// >
type MethodCall struct {
	Name              string             `xml:"NAME,attr,omitempty" json:",omitempty"`
	LocalClassPath    *LocalClassPath    `xml:"LOCALCLASSPATH" json:",omitempty"`
	LocalInstancePath *LocalInstancePath `xml:"LOCALINSTANCEPATH" json:",omitempty"`
	ParamValue        []ParamValue       `xml:"PARAMVALUE" json:",omitempty"`
}

// <!ELEMENT PARAMVALUE (VALUE | VALUE.REFERENCE | VALUE.ARRAY | VALUE.REFARRAY | CLASSNAME | INSTANCENAME | CLASS | INSTANCE | VALUE.NAMEDINSTANCE)?>
// <!ATTLIST PARAMVALUE
//     %CIMName;
//     %ParamType; #IMPLIED
//     %EmbeddedObject;
// >
type ParamValue struct {
	Value              *Value              `xml:"VALUE" json:",omitempty"`
	ValueReference     *ValueReference     `xml:"VALUE.REFERENCE" json:",omitempty"`
	ValueArray         *ValueArray         `xml:"VALUE.ARRAY" json:",omitempty"`
	ValueRefArray      *ValueRefArray      `xml:"VALUE.REFARRAY" json:",omitempty"`
	ClassName          *ClassName          `xml:"CLASSNAME" json:",omitempty"`
	InstanceName       *InstanceName       `xml:"INSTANCENAME" json:",omitempty"`
	Class              *Class              `xml:"CLASS" json:",omitempty"`
	Instance           *Instance           `xml:"INSTANCE" json:",omitempty"`
	ValueNamedInstance *ValueNamedInstance `xml:"VALUE.NAMEDINSTANCE" json:",omitempty"`
}

// <!ELEMENT IMETHODCALL (LOCALNAMESPACEPATH, IPARAMVALUE*)>
// <!ATTLIST IMETHODCALL
//     %CIMName;
// >
type IMethodCall struct {
	Name               string              `xml:"NAME,attr,omitempty" json:",omitempty"`
	LocalNamespacePath *LocalNamespacePath `xml:"LOCALNAMESPACEPATH" json:",omitempty"`
	IParamValue        []IParamValue       `xml:"IPARAMVALUE" json:",omitempty"`
}

// <!ELEMENT IPARAMVALUE (VALUE | VALUE.ARRAY | VALUE.REFERENCE | CLASSNAME | INSTANCENAME | QUALIFIER.DECLARATION | CLASS | INSTANCE | VALUE.NAMEDINSTANCE)?>
// <!ATTLIST IPARAMVALUE
//     %CIMName;
// >
type IParamValue struct {
	Name                 string                `xml:"NAME,attr,omitempty" json:",omitempty"`
	Value                *Value                `xml:"VALUE" json:",omitempty"`
	ValueArray           *ValueArray           `xml:"VALUE.ARRAY" json:",omitempty"`
	ValueReference       *ValueReference       `xml:"VALUE.REFERENCE" json:",omitempty"`
	ClassName            *ClassName            `xml:"CLASSNAME" json:",omitempty"`
	InstanceName         *InstanceName         `xml:"INSTANCENAME" json:",omitempty"`
	QualifierDeclaration *QualifierDeclaration `xml:"QUALIFIER.DECLARATION" json:",omitempty"`
	Class                *Class                `xml:"CLASS" json:",omitempty"`
	Instance             *Instance             `xml:"INSTANCE" json:",omitempty"`
	ValueNamedInstance   *ValueNamedInstance   `xml:"VALUE.NAMEDINSTANCE" json:",omitempty"`
}

// <!ELEMENT MULTIRSP (SIMPLERSP, SIMPLERSP+)>
type MultiRsp struct {
	SimpleRsp []SimpleRsp `xml:"SIMPLERSP" json:",omitempty"`
}

// <!ELEMENT SIMPLERSP (METHODRESPONSE | IMETHODRESPONSE)>
type SimpleRsp struct {
	MethodResponse  *MethodResponse  `xml:"METHODRESPONSE" json:",omitempty"`
	IMethodResponse *IMethodResponse `xml:"IMETHODRESPONSE" json:",omitempty"`
}

// <!ELEMENT METHODRESPONSE (ERROR | (RETURNVALUE?, PARAMVALUE*))>
// <!ATTLIST METHODRESPONSE
//     %CIMName;
// >
type MethodResponse struct {
	Name        string       `xml:"NAME,attr,omitempty" json:",omitempty"`
	Error       *Error       `xml:"ERROR" json:",omitempty"`
	ReturnValue *ReturnValue `xml:"RETURNVALUE" json:",omitempty"`
	ParamValue  []ParamValue `xml:"PARAMVALUE" json:",omitempty"`
}

// <!ELEMENT IMETHODRESPONSE (ERROR | (IRETURNVALUE?, PARAMVALUE*))>
// <!ATTLIST IMETHODRESPONSE
//     %CIMName;
// >
type IMethodResponse struct {
	Name         string        `xml:"NAME,attr,omitempty" json:",omitempty"`
	Error        *Error        `xml:"ERROR" json:",omitempty"`
	IReturnValue *IReturnValue `xml:"IRETURNVALUE" json:",omitempty"`
	ParamValue   []ParamValue  `xml:"PARAMVALUE" json:",omitempty"`
}

// <!ELEMENT ERROR (INSTANCE*)>
// <!ATTLIST ERROR
//     CODE CDATA #REQUIRED
//     DESCRIPTION CDATA #IMPLIED
// >
type Error struct {
	Code        string `xml:"CODE,attr" json:",omitempty"`
	Description string `xml:"DESCRIPTION,attr" json:",omitempty"`
}

// <!ELEMENT RETURNVALUE (VALUE | VALUE.REFERENCE)?>
// <!ATTLIST RETURNVALUE
//     %EmbeddedObject;
//     %ParamType; #IMPLIED
// >
type ReturnValue struct {
	EmbeddedObject string          `xml:"EmbeddedObject,attr,omitempty" json:",omitempty"`
	ParamType      string          `xml:"PARAMTYPE,attr,omitempty" json:",omitempty"`
	Value          *Value          `xml:"VALUE" json:",omitempty"`
	ValueReference *ValueReference `xml:"VALUE.REFERENCE" json:",omitempty"`
}

// <!ELEMENT IRETURNVALUE (CLASSNAME* | INSTANCENAME* | VALUE* | VALUE.OBJECTWITHPATH* | VALUE.OBJECTWITHLOCALPATH* | VALUE.OBJECT* | OBJECTPATH* | QUALIFIER.DECLARATION* | VALUE.ARRAY? | VALUE.REFERENCE? | CLASS* | INSTANCE* | INSTANCEPATH* | VALUE.NAMEDINSTANCE* | VALUE.INSTANCEWITHPATH*)>
type IReturnValue struct {
	ClassName                []ClassName
	InstanceName             []InstanceName             `xml:"INSTANCENAME" json:",omitempty"`
	Value                    []Value                    `xml:"VALUE" json:",omitempty"`
	ValueObjectWithPath      []ValueObjectWithPath      `xml:"VALUE.OBJECTWITHPATH" json:",omitempty"`
	ValueObjectWithLocalPath []ValueObjectWithLocalPath `xml:"VALUE.OBJECTWITHLOCALPATH" json:",omitempty"`
	ValueObject              []ValueObject              `xml:"VALUE.OBJECT" json:",omitempty"`
	ObjectPath               []ObjectPath               `xml:"OBJECTPATH" json:",omitempty"`
	QualifierDeclaration     []QualifierDeclaration     `xml:"QUALIFIER.DECLARATION" json:",omitempty"`
	ValueArray               *ValueArray                `xml:"VALUE.ARRAY" json:",omitempty"`
	ValueReference           *ValueReference            `xml:"VALUE.REFERENCE" json:",omitempty"`
	Class                    []Class                    `xml:"CLASS" json:",omitempty"`
	Instance                 []Instance                 `xml:"INSTANCE" json:",omitempty"`
	InstancePath             []InstancePath             `xml:"INSTANCEPATH" json:",omitempty"`
	ValueNamedInstance       []ValueNamedInstance       `xml:"VALUE.NAMEDINSTANCE" json:",omitempty"`
	ValueInstanceWithPath    []ValueInstanceWithPath    `xml:"VALUE.INSTANCEWITHPATH" json:",omitempty"`
}

// <!ELEMENT MULTIEXPREQ (SIMPLEEXPREQ, SIMPLEEXPREQ+)>
type MultiExpReq struct {
	SimpleExpReq []SimpleExpReq `xml:"SIMPLEEXPREQ" json:",omitempty"`
}

// <!ELEMENT SIMPLEEXPREQ (CORRELATOR*, EXPMETHODCALL)>
type SimpleExpReq struct {
	Correlator    []Correlator   `xml:"CORRELATOR" json:",omitempty"`
	ExpMethodCall *ExpMethodCall `xml:"EXPMETHODCALL" json:",omitempty"`
}

// <!ELEMENT EXPMETHODCALL (EXPPARAMVALUE*)>
// <!ATTLIST EXPMETHODCALL
//     %CIMName;
// >
type ExpMethodCall struct {
	Name          string          `xml:"NAME,attr,omitempty" json:",omitempty"`
	ExpParamValue []ExpParamValue `xml:"EXPPARAMVALUE" json:",omitempty"`
}

// <!ELEMENT MULTIEXPRSP (SIMPLEEXPRSP, SIMPLEEXPRSP+)>
type MultiExpRsp struct {
	SimpleExpRsp []SimpleExpRsp `xml:"SIMPLEEXPRSP" json:",omitempty"`
}

// <!ELEMENT SIMPLEEXPRSP (EXPMETHODRESPONSE)>
type SimpleExpRsp struct {
	ExpMethodResponse *ExpMethodResponse `xml:"EXPMETHODRESPONSE" json:",omitempty"`
}

// <!ELEMENT EXPMETHODRESPONSE (ERROR | IRETURNVALUE?)>
// <!ATTLIST EXPMETHODRESPONSE
//     %CIMName;
// >
type ExpMethodResponse struct {
	Name         string        `xml:"NAME,attr,omitempty" json:",omitempty"`
	Error        *Error        `xml:"ERROR" json:",omitempty"`
	IReturnValue *IReturnValue `xml:"IRETURNVALUE" json:",omitempty"`
}

// <!ELEMENT EXPPARAMVALUE (INSTANCE?)>
// <!ATTLIST EXPPARAMVALUE
//     %CIMName;
// >
type ExpParamValue struct {
	Name     string    `xml:"NAME,attr,omitempty" json:",omitempty"`
	Instance *Instance `xml:"INSTANCE" json:",omitempty"`
}

// <!--
// **************************************************
// CHANGE NOTE: The ENUMERATIONCONTEXT element was
// removed in version 2.4.0 of this document.
// **************************************************
// -->

// <!ELEMENT CORRELATOR (VALUE)>
// <!ATTLIST CORRELATOR
//     %CIMName;
//     %CIMType; #REQUIRED
// >
type Correlator struct {
	Name  string `xml:"NAME,attr,omitempty" json:",omitempty"`
	Type  string `xml:"TYPE,attr" json:",omitempty"`
	Value *Value `xml:"VALUE" json:",omitempty"`
}
