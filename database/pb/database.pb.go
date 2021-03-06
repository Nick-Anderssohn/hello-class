// Code generated by protoc-gen-go. DO NOT EDIT.
// source: database.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	database.proto

It has these top-level messages:
	NewClassReq
	SessionResp
	Submission
	Setting
	Problem
	EducatorHomeData
	EducatorHomeDataRequest
	LoginRequest
	LoginResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NewClassReq struct {
	ClassName   string `protobuf:"bytes,1,opt,name=className" json:"className,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	SessionGUID string `protobuf:"bytes,4,opt,name=sessionGUID" json:"sessionGUID,omitempty"`
}

func (m *NewClassReq) Reset()                    { *m = NewClassReq{} }
func (m *NewClassReq) String() string            { return proto.CompactTextString(m) }
func (*NewClassReq) ProtoMessage()               {}
func (*NewClassReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewClassReq) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *NewClassReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NewClassReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *NewClassReq) GetSessionGUID() string {
	if m != nil {
		return m.SessionGUID
	}
	return ""
}

type SessionResp struct {
	SessionGUID string `protobuf:"bytes,1,opt,name=sessionGUID" json:"sessionGUID,omitempty"`
	Success     bool   `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	Message     string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *SessionResp) Reset()                    { *m = SessionResp{} }
func (m *SessionResp) String() string            { return proto.CompactTextString(m) }
func (*SessionResp) ProtoMessage()               {}
func (*SessionResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SessionResp) GetSessionGUID() string {
	if m != nil {
		return m.SessionGUID
	}
	return ""
}

func (m *SessionResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SessionResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Submission struct {
	StudentName string `protobuf:"bytes,1,opt,name=studentName" json:"studentName,omitempty"`
	AnswerText  string `protobuf:"bytes,2,opt,name=answerText" json:"answerText,omitempty"`
	Graded      bool   `protobuf:"varint,3,opt,name=graded" json:"graded,omitempty"`
	Correct     bool   `protobuf:"varint,4,opt,name=correct" json:"correct,omitempty"`
	Success     bool   `protobuf:"varint,5,opt,name=success" json:"success,omitempty"`
	Message     string `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	Id          uint64 `protobuf:"varint,7,opt,name=id" json:"id,omitempty"`
	ProblemID   uint64 `protobuf:"varint,8,opt,name=problemID" json:"problemID,omitempty"`
}

func (m *Submission) Reset()                    { *m = Submission{} }
func (m *Submission) String() string            { return proto.CompactTextString(m) }
func (*Submission) ProtoMessage()               {}
func (*Submission) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Submission) GetStudentName() string {
	if m != nil {
		return m.StudentName
	}
	return ""
}

func (m *Submission) GetAnswerText() string {
	if m != nil {
		return m.AnswerText
	}
	return ""
}

func (m *Submission) GetGraded() bool {
	if m != nil {
		return m.Graded
	}
	return false
}

func (m *Submission) GetCorrect() bool {
	if m != nil {
		return m.Correct
	}
	return false
}

func (m *Submission) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Submission) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Submission) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Submission) GetProblemID() uint64 {
	if m != nil {
		return m.ProblemID
	}
	return 0
}

type Setting struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	IsSelected bool   `protobuf:"varint,2,opt,name=isSelected" json:"isSelected,omitempty"`
	Success    bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
	Message    string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	Id         uint64 `protobuf:"varint,5,opt,name=id" json:"id,omitempty"`
}

func (m *Setting) Reset()                    { *m = Setting{} }
func (m *Setting) String() string            { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()               {}
func (*Setting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Setting) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Setting) GetIsSelected() bool {
	if m != nil {
		return m.IsSelected
	}
	return false
}

func (m *Setting) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Setting) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Setting) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Problem struct {
	Id             uint64        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title          string        `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Prompt         string        `protobuf:"bytes,3,opt,name=prompt" json:"prompt,omitempty"`
	Submissions    []*Submission `protobuf:"bytes,4,rep,name=submissions" json:"submissions,omitempty"`
	Settings       []*Setting    `protobuf:"bytes,5,rep,name=settings" json:"settings,omitempty"`
	Success        bool          `protobuf:"varint,6,opt,name=success" json:"success,omitempty"`
	Message        string        `protobuf:"bytes,7,opt,name=message" json:"message,omitempty"`
	ClassName      string        `protobuf:"bytes,8,opt,name=className" json:"className,omitempty"`
	SessionGUID    string        `protobuf:"bytes,9,opt,name=sessionGUID" json:"sessionGUID,omitempty"`
	ExpectedOutput string        `protobuf:"bytes,10,opt,name=expectedOutput" json:"expectedOutput,omitempty"`
}

func (m *Problem) Reset()                    { *m = Problem{} }
func (m *Problem) String() string            { return proto.CompactTextString(m) }
func (*Problem) ProtoMessage()               {}
func (*Problem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Problem) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Problem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Problem) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

func (m *Problem) GetSubmissions() []*Submission {
	if m != nil {
		return m.Submissions
	}
	return nil
}

func (m *Problem) GetSettings() []*Setting {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Problem) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Problem) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Problem) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *Problem) GetSessionGUID() string {
	if m != nil {
		return m.SessionGUID
	}
	return ""
}

func (m *Problem) GetExpectedOutput() string {
	if m != nil {
		return m.ExpectedOutput
	}
	return ""
}

type EducatorHomeData struct {
	ClassName      string   `protobuf:"bytes,1,opt,name=className" json:"className,omitempty"`
	ProblemTitles  []string `protobuf:"bytes,2,rep,name=problemTitles" json:"problemTitles,omitempty"`
	ProblemIDs     []uint64 `protobuf:"varint,3,rep,packed,name=problemIDs" json:"problemIDs,omitempty"`
	CurrentProblem *Problem `protobuf:"bytes,4,opt,name=currentProblem" json:"currentProblem,omitempty"`
	Success        bool     `protobuf:"varint,5,opt,name=success" json:"success,omitempty"`
	Message        string   `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
}

func (m *EducatorHomeData) Reset()                    { *m = EducatorHomeData{} }
func (m *EducatorHomeData) String() string            { return proto.CompactTextString(m) }
func (*EducatorHomeData) ProtoMessage()               {}
func (*EducatorHomeData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EducatorHomeData) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *EducatorHomeData) GetProblemTitles() []string {
	if m != nil {
		return m.ProblemTitles
	}
	return nil
}

func (m *EducatorHomeData) GetProblemIDs() []uint64 {
	if m != nil {
		return m.ProblemIDs
	}
	return nil
}

func (m *EducatorHomeData) GetCurrentProblem() *Problem {
	if m != nil {
		return m.CurrentProblem
	}
	return nil
}

func (m *EducatorHomeData) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *EducatorHomeData) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EducatorHomeDataRequest struct {
	ClassName   string `protobuf:"bytes,1,opt,name=className" json:"className,omitempty"`
	SessionGUID string `protobuf:"bytes,2,opt,name=sessionGUID" json:"sessionGUID,omitempty"`
}

func (m *EducatorHomeDataRequest) Reset()                    { *m = EducatorHomeDataRequest{} }
func (m *EducatorHomeDataRequest) String() string            { return proto.CompactTextString(m) }
func (*EducatorHomeDataRequest) ProtoMessage()               {}
func (*EducatorHomeDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *EducatorHomeDataRequest) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *EducatorHomeDataRequest) GetSessionGUID() string {
	if m != nil {
		return m.SessionGUID
	}
	return ""
}

type LoginRequest struct {
	ClassName   string `protobuf:"bytes,1,opt,name=className" json:"className,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	SessionGUID string `protobuf:"bytes,3,opt,name=sessionGUID" json:"sessionGUID,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LoginRequest) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetSessionGUID() string {
	if m != nil {
		return m.SessionGUID
	}
	return ""
}

type LoginResponse struct {
	ClassName   string `protobuf:"bytes,1,opt,name=className" json:"className,omitempty"`
	Success     bool   `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	SessionGUID string `protobuf:"bytes,3,opt,name=sessionGUID" json:"sessionGUID,omitempty"`
	Message     string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *LoginResponse) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *LoginResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LoginResponse) GetSessionGUID() string {
	if m != nil {
		return m.SessionGUID
	}
	return ""
}

func (m *LoginResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*NewClassReq)(nil), "pb.NewClassReq")
	proto.RegisterType((*SessionResp)(nil), "pb.SessionResp")
	proto.RegisterType((*Submission)(nil), "pb.Submission")
	proto.RegisterType((*Setting)(nil), "pb.Setting")
	proto.RegisterType((*Problem)(nil), "pb.Problem")
	proto.RegisterType((*EducatorHomeData)(nil), "pb.EducatorHomeData")
	proto.RegisterType((*EducatorHomeDataRequest)(nil), "pb.EducatorHomeDataRequest")
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
}

func init() { proto.RegisterFile("database.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x8b, 0xd3, 0x4c,
	0x14, 0x26, 0x1f, 0x6d, 0xd3, 0x93, 0x77, 0xcb, 0xcb, 0xb0, 0x68, 0x10, 0x91, 0x12, 0x44, 0x7b,
	0x55, 0xc4, 0xfd, 0x09, 0x56, 0x74, 0x41, 0x56, 0x99, 0xae, 0x17, 0x5e, 0x4e, 0x92, 0x43, 0x89,
	0x34, 0x1f, 0x9b, 0x33, 0xa1, 0x7b, 0x27, 0xe2, 0xff, 0xf2, 0xaf, 0x08, 0xfe, 0x12, 0x99, 0xc9,
	0xb4, 0x9d, 0xa6, 0xd6, 0x8a, 0x77, 0x39, 0xcf, 0x39, 0x99, 0xf3, 0x3c, 0xe7, 0x0b, 0x26, 0x99,
	0x90, 0x22, 0x11, 0x84, 0xf3, 0xba, 0xa9, 0x64, 0xc5, 0xdc, 0x3a, 0x89, 0xbf, 0x40, 0x78, 0x83,
	0x9b, 0x57, 0x6b, 0x41, 0xc4, 0xf1, 0x8e, 0x3d, 0x86, 0x71, 0xaa, 0xbe, 0x6f, 0x44, 0x81, 0x91,
	0x33, 0x75, 0x66, 0x63, 0xbe, 0x07, 0xd8, 0x25, 0x0c, 0xb0, 0x10, 0xf9, 0x3a, 0x72, 0xb5, 0xa7,
	0x33, 0xd8, 0x23, 0x08, 0x6a, 0x41, 0xb4, 0xa9, 0x9a, 0x2c, 0xf2, 0xb4, 0x63, 0x67, 0xb3, 0x29,
	0x84, 0x84, 0x44, 0x79, 0x55, 0xbe, 0xf9, 0x78, 0xbd, 0x88, 0x7c, 0xed, 0xb6, 0xa1, 0x38, 0x85,
	0x70, 0xd9, 0x99, 0x1c, 0xa9, 0xee, 0xff, 0xe0, 0x1c, 0xfd, 0xc0, 0x22, 0x18, 0x51, 0x9b, 0xa6,
	0x48, 0xa4, 0x69, 0x04, 0x7c, 0x6b, 0x2a, 0x4f, 0x81, 0x44, 0x62, 0x85, 0x86, 0xc7, 0xd6, 0x8c,
	0x7f, 0x3a, 0x00, 0xcb, 0x36, 0x29, 0x72, 0xfd, 0x8c, 0x4e, 0x22, 0xdb, 0x0c, 0x4b, 0x69, 0xe9,
	0xb4, 0x21, 0xf6, 0x04, 0x40, 0x94, 0xb4, 0xc1, 0xe6, 0x16, 0xef, 0xa5, 0x91, 0x6b, 0x21, 0xec,
	0x01, 0x0c, 0x57, 0x8d, 0xc8, 0xb0, 0x53, 0x1c, 0x70, 0x63, 0x29, 0x0a, 0x69, 0xd5, 0x34, 0x98,
	0x4a, 0xad, 0x35, 0xe0, 0x5b, 0xd3, 0xa6, 0x3d, 0x38, 0x49, 0x7b, 0x78, 0x40, 0x9b, 0x4d, 0xc0,
	0xcd, 0xb3, 0x68, 0x34, 0x75, 0x66, 0x3e, 0x77, 0xf3, 0x4c, 0x75, 0xa7, 0x6e, 0xaa, 0x64, 0x8d,
	0xc5, 0xf5, 0x22, 0x0a, 0x34, 0xbc, 0x07, 0xe2, 0xaf, 0x0e, 0x8c, 0x96, 0x28, 0x65, 0x5e, 0xae,
	0x18, 0x03, 0xbf, 0xdc, 0x4b, 0xd3, 0xdf, 0x4a, 0x53, 0x4e, 0x4b, 0x5c, 0x63, 0x2a, 0x31, 0x33,
	0xb5, 0xb3, 0x10, 0x9b, 0xa1, 0x77, 0x92, 0xa1, 0xff, 0x3b, 0x86, 0x83, 0x2d, 0xc3, 0xf8, 0xbb,
	0x0b, 0xa3, 0x0f, 0x1d, 0x23, 0xe3, 0x73, 0x76, 0xec, 0x2f, 0x61, 0x20, 0x73, 0xb9, 0xc6, 0xed,
	0xf4, 0x68, 0x43, 0x55, 0xb2, 0x6e, 0xaa, 0xa2, 0x96, 0xa6, 0x67, 0xc6, 0x62, 0x2f, 0x20, 0xa4,
	0x5d, 0xc7, 0x28, 0xf2, 0xa7, 0xde, 0x2c, 0x7c, 0x39, 0x99, 0xd7, 0xc9, 0x7c, 0xdf, 0x48, 0x6e,
	0x87, 0xb0, 0xe7, 0x10, 0x50, 0x27, 0x5f, 0x95, 0x58, 0x85, 0x87, 0x3a, 0xbc, 0xc3, 0xf8, 0xce,
	0x69, 0x0b, 0x1d, 0x9e, 0x14, 0x3a, 0x3a, 0x14, 0x7a, 0xb0, 0x18, 0x41, 0x7f, 0x31, 0x7a, 0x53,
	0x3b, 0x3e, 0x9e, 0xda, 0x67, 0x30, 0xc1, 0xfb, 0x5a, 0x17, 0xfa, 0x7d, 0x2b, 0xeb, 0x56, 0x46,
	0xa0, 0x83, 0x7a, 0x68, 0xfc, 0xc3, 0x81, 0xff, 0x5f, 0x67, 0x6d, 0x2a, 0x64, 0xd5, 0xbc, 0xad,
	0x0a, 0x5c, 0x08, 0x29, 0xce, 0x6c, 0xe5, 0x53, 0xb8, 0x30, 0x43, 0x70, 0xab, 0x2a, 0xaa, 0xd6,
	0xc2, 0x9b, 0x8d, 0xf9, 0x21, 0xa8, 0xba, 0xbf, 0x1b, 0x15, 0xd5, 0x60, 0x6f, 0xe6, 0x73, 0x0b,
	0x61, 0x57, 0x30, 0x49, 0xdb, 0xa6, 0xc1, 0x52, 0x9a, 0xfe, 0xe9, 0x56, 0x9b, 0x1a, 0x1a, 0x88,
	0xf7, 0x42, 0xfe, 0x65, 0xa8, 0xe3, 0x4f, 0xf0, 0xb0, 0x2f, 0x90, 0xe3, 0x5d, 0x8b, 0x24, 0xcf,
	0xe8, 0xec, 0x15, 0xd9, 0x3d, 0xbe, 0x25, 0x9f, 0xe1, 0xbf, 0x77, 0xd5, 0x2a, 0x2f, 0xff, 0xee,
	0x3d, 0xfb, 0x6e, 0xb9, 0x7f, 0xbe, 0x5b, 0xde, 0x71, 0xae, 0x6f, 0x0e, 0x5c, 0x98, 0x64, 0x54,
	0x57, 0x25, 0xe1, 0x99, 0x6c, 0xa7, 0xcf, 0xd6, 0xd9, 0x5c, 0xa7, 0xf7, 0x2f, 0x19, 0xea, 0x4b,
	0x7e, 0xf5, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x93, 0xcb, 0xfd, 0x8a, 0xdb, 0x05, 0x00, 0x00,
}
