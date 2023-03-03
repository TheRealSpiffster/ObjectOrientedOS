package emulator

type Method struct {
	Name  uint32 // points to string
	Start uint32 // points to function start
}

type Link struct {
	ClassName  uint32
	MethodName uint32
}

// library format -- this structure is read by the OS when it activates the class
type Library struct {
	Name               uint32 // 32 bit pointer to name of class
	NameSize           uint32 //
	CodeSize           uint32
	GlobalSize         uint32
	DataSize           uint32
	MethodsIDSize      uint32 // size of methods data
	LinksSize          uint32 // size of link data
	Links              []byte // external calls into other classes
	Strings            []byte // string data
	Methods            []byte // methods array data
	CodeSegment        []byte // code segment data
	GlobalSegment      []byte // global variable data
	DefaultDataSegment []byte // data segment (instance) start data
}
