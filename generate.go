package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	useCLI := true

	if useCLI {
		generateWithCLI()
	} else {
		structName := "value"
		receiverName := "val"
		modelPackage := "valmodels"
		fields := []Field{
			NewField("ID", "bson.ObjectId"),
			NewField("Name", "string"),
			NewField("Value", "xjson.Value"),
			NewField("Private", "bool"),
		}

		// output the write
		writeFiles(NewModel(structName, receiverName, modelPackage, "model.go", fModel, fields))
	}
}

func generateWithCLI() {
	prompt := promptui.Prompt{Label: "Enter Model/Struct Name "}
	structName, err := prompt.Run()
	check(err)

	prompt = promptui.Prompt{Label: "Enter Model Package Name "}
	modelPackage, err := prompt.Run()
	check(err)

	prompt = promptui.Prompt{Label: "Enter Model Receiver Name "}
	receiverName, err := prompt.Run()
	check(err)

	var fields []Field
	for {
		prompt = promptui.Prompt{Label: "Enter Field Name or 'done' to finish "}
		fieldName, err := prompt.Run()
		check(err)

		if fieldName == "done" {
			break
		}

		prompt = promptui.Prompt{Label: "Enter Field Type "}
		fieldType, err := prompt.Run()
		check(err)

		fields = append(fields, NewField(fieldName, fieldType))
	}
	writeFiles(NewModel(structName, receiverName, modelPackage, "model.go", fModel, fields))
}

func writeFiles(model Model) {
	// Open three files
	modelBytes := bytes.NewBuffer([]byte{})
	bsonBytes := bytes.NewBuffer([]byte{})
	jsonBytes := bytes.NewBuffer([]byte{})

	// write the model
	model.f = modelBytes
	writeModel(model)

	// write the bson model
	model.externalStructName = model.interfaceName
	model.gettersName = model.interfaceName
	model.interfaceName = model.packageName + "." + model.interfaceName
	model.f = bsonBytes
	model.fileName = "model_bson.go"
	model.fType = fBSON
	writeModel(model)

	model.f = jsonBytes
	model.fileName = "model_json.go"
	model.fType = fJSON
	writeModel(model)

}

func writeLine(file *bytes.Buffer, str string, a ...interface{}) {
	// fmt.Println("*")
	file.WriteString(fmt.Sprintf(str+"\n", a...))
	// fmt.Println(string(file.Bytes()))
	// fmt.Println("*")
}

func writeModel(model Model) {
	switch model.fType {
	case fModel:
		writeLine(model.f, "package %s", model.packageName)
		writeLine(model.f, "import (\n\t\"errors\"\n\t\"fmt\")")
	case fJSON:
		writeLine(model.f, "package apimodelsv3")
		writeLine(model.f, "import (\n\t\"errors\"\n\t\"encoding/json\"\n\n\t\"github.com/10gen/stitch/utils/xjson\"\n\n\t\"gopkg.in/mgo.v2/bson\")")
	case fBSON:
		writeLine(model.f, "package INSERTPACKAGENAMEHERE")
		writeLine(model.f, "import (\n\t\"errors\"\n\n\t\"github.com/10gen/stitch/utils/xjson\"\n\n\t\"gopkg.in/mgo.v2/bson\")")
	}

	if model.fType == fModel {
		writeInterfaceStub(model)
	}

	writeExternalStruct(model)
	writeStructStub(model)
	writeNewFunc(model)

	writeInterfaceMethods(model)

	if model.fType == fModel {
		writeSetStubs(model)
	} else {
		writeMarshallingStubs(model)
		writeToStructStub(model)
	}

	fileBytes, err := format.Source(model.f.Bytes())
	check(err)
	err = ioutil.WriteFile(model.fileName, fileBytes, 0644)
	check(err)
}

func writeInterfaceStub(model Model) {
	writeLine(model.f, "// %s represents ... TODO", model.interfaceName)
	writeLine(model.f, "type %s interface {", model.interfaceName)

	for _, field := range model.fields {
		writeLine(model.f, "%s() %s", field.upperName, field.ftype)
	}

	writeLine(model.f, "Builder() *%s", model.externalStructName)

	writeLine(model.f, "}")
	writeLine(model.f, "")

}

func writeExternalStruct(model Model) {
	if model.fType == fModel {
		writeLine(model.f, "// %s builds and validates %ss", model.externalStructName, model.internalStructName)

	} else {
		writeLine(model.f, "// %s is the representation of a %s interface", model.externalStructName, model.interfaceName)
	}
	writeLine(model.f, "type %s struct {", model.externalStructName)
	writeLine(model.f, "data %s", model.internalStructName)
	writeLine(model.f, "}\n")
}

func writeStructStub(model Model) {
	// writeLine(model.f, "// %s represents ... TODO", model.interfaceName)
	writeLine(model.f, "type %s struct {", model.internalStructName)

	for _, field := range model.fields {
		if model.fType == fModel {
			writeLine(model.f, "%s %s", field.lowerName, field.ftype)
			continue
		}
		str := fmt.Sprintf("%s %s", field.upperName, field.ftype)
		if model.fType == fBSON {
			str += fmt.Sprintf(" `bson:\"%s\"`", field.bsonName)
		} else if model.fType == fJSON {
			str += fmt.Sprintf(" `json:\"%s\"`", field.bsonName)
		}
		writeLine(model.f, str)
	}

	writeLine(model.f, "}\n")
}

func writeNewFunc(model Model) {
	writeLine(model.f, "// New%s returns a new %s", model.externalStructName, model.externalStructName)
	writeLine(model.f, "func New%s() *%s {", model.externalStructName, model.externalStructName)
	writeLine(model.f, "return &%s{}", model.externalStructName)
	// str := "func New" + model.interfaceName + "("
	// for _, field := range model.fields {
	// 	str += field.lowerName + " " + field.ftype + ","
	// }
	// str = strings.TrimSuffix(str, ",") + ") " + model.interfaceName + "{"
	// writeLine(model.f, str)

	writeLine(model.f, "}\n")
}

func writeInterfaceMethods(model Model) {
	for _, field := range model.fields {
		writeLine(model.f, "// %s returns the %s of this %s", field.upperName, field.lowerName, model.gettersName)
		writeLine(model.f, "func (%s *%s) %s() %s {", model.receiverName, model.gettersName, field.upperName, field.ftype)
		if model.fType == fModel {
			writeLine(model.f, "return %s.%s", model.receiverName, field.lowerName)

		} else {
			writeLine(model.f, "return %s.data.%s", model.receiverName, field.upperName)

		}
		writeLine(model.f, "}\n")
	}

	builder := model.externalStructName
	if model.fType != fModel {
		builder = model.interfaceName + "Builder"
	}
	writeLine(model.f, "// Builder creates a shallow copy of the %s and returns it as a %s", model.gettersName, builder)
	writeLine(model.f, "func (%s *%s) Builder() *%s {", model.receiverName, model.gettersName, builder)
	str := ""
	if model.fType == fModel {
		str = fmt.Sprintf("builder := New%s().\n", model.externalStructName)
		for _, field := range model.fields {
			str += fmt.Sprintf("\tWith%s(%s.%s()).\n", field.upperName, model.receiverName, field.upperName)
		}
	} else {
		str = fmt.Sprintf("builder := %s.New%sBuilder().\n", model.packageName, model.externalStructName)
		for _, field := range model.fields {
			str += fmt.Sprintf("\tWith%s(%s.%s()).\n", field.upperName, model.receiverName, field.upperName)
		}
	}
	str = str[:len(str)-2]
	writeLine(model.f, str+"\n")
	writeLine(model.f, "// perform any necessary checks")
	writeLine(model.f, "// if ....\n")
	writeLine(model.f, "return builder")
	writeLine(model.f, "}\n")
}

func writeSetStubs(model Model) {
	for _, field := range model.fields {
		writeLine(model.f, "// With%s sets the %s for the %s", field.upperName, field.upperName, model.externalStructName)
		writeLine(model.f, "func (builder *%s) With%s(%s %s) *%s {", model.externalStructName, field.upperName, field.lowerName, field.ftype, model.externalStructName)
		writeLine(model.f, "builder.data.%s = %s", field.lowerName, field.lowerName)
		writeLine(model.f, "return builder")
		writeLine(model.f, "}")
	}

	writeLine(model.f, "// Build builds a new %s if it is validated", model.interfaceName)
	writeLine(model.f, "func (builder *%s) Build() (%s, error) {", model.externalStructName, model.interfaceName)
	str := fmt.Sprintf("built := &%s{\n", model.internalStructName)
	for _, field := range model.fields {
		str += fmt.Sprintf("%s: builder.data.%s,\n", field.lowerName, field.lowerName)
	}
	writeLine(model.f, str+"}\n")
	writeLine(model.f, "// Do relevant checks here")
	writeLine(model.f, "// if %s.data.%s == \"\" {", model.receiverName, model.fields[0].lowerName)
	writeLine(model.f, "// 		return nil, errors.New(\"ERROR MESSAGE\")")
	writeLine(model.f, "// }\n")
	writeLine(model.f, "return built, nil")
	writeLine(model.f, "}")

	writeLine(model.f, "// MustBuild calls Build() but panics if there is an error")
	writeLine(model.f, "func (builder *%s) MustBuild() %s {", model.externalStructName, model.interfaceName)
	writeLine(model.f, "data, err := builder.Build()")
	writeLine(model.f, "if err != nil {")
	writeLine(model.f, "panic(fmt.Errorf(\"failed to build %s: %%v\", err))", model.internalStructName)
	writeLine(model.f, "}")
	writeLine(model.f, "return data")
	writeLine(model.f, "}")
}

func writeMarshallingStubs(model Model) {
	sonType := "BSON"
	if model.fType == fJSON {
		sonType = "JSON"
	}
	writeLine(model.f, "// %s marshals the %s to %s", sonType, model.externalStructName, sonType)
	writeLine(model.f, "func (%s %s) Marshal%s() ([]byte, error) {\nreturn %s.Marshal(%s)\n}", model.receiverName, model.gettersName, sonType, strings.ToLower(sonType), model.receiverName)

	writeLine(model.f, "// Unmarshal%s unmarshals the %s from %s", sonType, model.externalStructName, sonType)
	writeLine(model.f, "func (%s *%s) Unmarshal%s(data []byte) error {\n return %s.Unmarshal(data, %s)\n}", model.receiverName, model.gettersName, sonType, strings.ToLower(sonType), model.receiverName)

	writeLine(model.f, "// Get%s returns the inner data for %s marshaling", sonType, sonType)
	writeLine(model.f, "func (%s %s) Get%s() (interface{}, error) {\nreturn %s.data, nil\n}", model.receiverName, model.gettersName, sonType, model.receiverName)

	writeLine(model.f, "// Set%s unmarshals %s onto the %s", sonType, sonType, model.externalStructName)
	writeLine(model.f, "func (%s *%s) Set%s(raw %s.Raw) error {\n return raw.Unmarshal(&%s.data) \n}", model.receiverName, model.gettersName, sonType, strings.ToLower(sonType), model.receiverName)
}

func writeToStructStub(model Model) {
	writeLine(model.f, "// To%s converts a %s to a %s", model.externalStructName, model.interfaceName, model.externalStructName)
	writeLine(model.f, "func To%s(%s %s) *%s {", model.externalStructName, model.receiverName, model.interfaceName, model.externalStructName)
	str := fmt.Sprintf("data := %s{\n", model.internalStructName)
	for _, field := range model.fields {
		str += fmt.Sprintf("%s: %s.%s(),\n", field.upperName, model.receiverName, field.upperName)
	}
	writeLine(model.f, str+"}\n")
	writeLine(model.f, "// Perform some checks\n// if data.... == \n")
	writeLine(model.f, "return &%s{\ndata: data,\n}", model.externalStructName)
	writeLine(model.f, "}\n")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type FileType int

const (
	fModel FileType = 0
	fBSON  FileType = 1
	fJSON  FileType = 2
)

type Field struct {
	lowerName string
	upperName string
	bsonName  string
	ftype     string
}

func NewField(name, ftype string) Field {
	f := Field{ftype: ftype, upperName: strings.Title(name)}
	f.lowerName = strings.ToLower(string(name[0])) + name[1:]
	f.bsonName = f.lowerName
	if name == "ID" || name == "id" {
		f.lowerName = "id"
		f.upperName = "ID"
		f.bsonName = "_id"
	}
	return f
}

type Model struct {
	interfaceName      string
	internalStructName string
	externalStructName string
	gettersName        string
	receiverName       string
	packageName        string
	fileName           string
	fType              FileType
	fields             []Field
	f                  *bytes.Buffer
}

func NewModel(name, receiverName, packageName, fileName string, fileType FileType, fields []Field) Model {
	return Model{
		interfaceName:      strings.Title(name),
		internalStructName: strings.ToLower(string(name[0])) + name[1:],
		externalStructName: strings.Title(name) + "Builder",
		gettersName:        strings.ToLower(string(name[0])) + name[1:],
		receiverName:       receiverName,
		packageName:        packageName,
		fType:              fileType,
		fileName:           fileName,
		fields:             fields,
	}
}
