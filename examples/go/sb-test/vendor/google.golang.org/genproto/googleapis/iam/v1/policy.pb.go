// Code generated by protoc-gen-go.
// source: google/iam/v1/policy.proto
// DO NOT EDIT!

package iam

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Defines an Identity and Access Management (IAM) policy. It is used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` consists of a list of `bindings`. A `Binding` binds a list of
// `members` to a `role`, where the members can be user accounts, Google groups,
// Google domains, and service accounts. A `role` is a named list of permissions
// defined by IAM.
//
// **Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/owner",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//             "serviceAccount:my-other-app@appspot.gserviceaccount.com",
//           ]
//         },
//         {
//           "role": "roles/viewer",
//           "members": ["user:sean@example.com"]
//         }
//       ]
//     }
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam).
type Policy struct {
	// Version of the `Policy`. The default version is 0.
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Associates a list of `members` to a `role`.
	// Multiple `bindings` must not be specified for the same `role`.
	// `bindings` with no members will result in an error.
	Bindings []*Binding `protobuf:"bytes,4,rep,name=bindings" json:"bindings,omitempty"`
	// `etag` is used for optimistic concurrency control as a way to help
	// prevent simultaneous updates of a policy from overwriting each other.
	// It is strongly suggested that systems make use of the `etag` in the
	// read-modify-write cycle to perform policy updates in order to avoid race
	// conditions: An `etag` is returned in the response to `getIamPolicy`, and
	// systems are expected to put that etag in the request to `setIamPolicy` to
	// ensure that their change will be applied to the same version of the policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the existing
	// policy is overwritten blindly.
	Etag []byte `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Policy) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Policy) GetBindings() []*Binding {
	if m != nil {
		return m.Bindings
	}
	return nil
}

func (m *Policy) GetEtag() []byte {
	if m != nil {
		return m.Etag
	}
	return nil
}

// Associates `members` with a `role`.
type Binding struct {
	// Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	// Required
	Role string `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	// Specifies the identities requesting access for a Cloud Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents anyone
	//    who is authenticated with a Google account or a service account.
	//
	// * `user:{emailid}`: An email address that represents a specific Google
	//    account. For example, `alice@gmail.com` or `joe@example.com`.
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a service
	//    account. For example, `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google group.
	//    For example, `admins@example.com`.
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all the
	//    users of that domain. For example, `google.com` or `example.com`.
	//
	//
	Members []string `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *Binding) Reset()                    { *m = Binding{} }
func (m *Binding) String() string            { return proto.CompactTextString(m) }
func (*Binding) ProtoMessage()               {}
func (*Binding) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Binding) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Binding) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "google.iam.v1.Policy")
	proto.RegisterType((*Binding)(nil), "google.iam.v1.Binding")
}

func init() { proto.RegisterFile("google/iam/v1/policy.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xe5, 0xb6, 0xb4, 0xd4, 0x85, 0x81, 0x20, 0xa1, 0xa8, 0x62, 0x88, 0x3a, 0x65, 0xb2,
	0x49, 0x19, 0x18, 0xd8, 0xc2, 0x80, 0xd8, 0xa2, 0x0c, 0x0c, 0x6c, 0x97, 0xd6, 0xb2, 0x0e, 0xc5,
	0xbe, 0xc8, 0x0e, 0x91, 0xf8, 0x4b, 0xfc, 0x42, 0x46, 0x14, 0x3b, 0x45, 0xea, 0x76, 0xa7, 0xef,
	0x3d, 0xbd, 0x7b, 0xc7, 0xb7, 0x9a, 0x48, 0xb7, 0x4a, 0x22, 0x18, 0x39, 0x14, 0xb2, 0xa3, 0x16,
	0x0f, 0xdf, 0xa2, 0x73, 0xd4, 0x53, 0x72, 0x1d, 0x99, 0x40, 0x30, 0x62, 0x28, 0xb6, 0xf7, 0x93,
	0x14, 0x3a, 0x94, 0x60, 0x2d, 0xf5, 0xd0, 0x23, 0x59, 0x1f, 0xc5, 0xbb, 0x4f, 0xbe, 0xac, 0x82,
	0x39, 0x49, 0xf9, 0x6a, 0x50, 0xce, 0x23, 0xd9, 0x94, 0x65, 0x2c, 0xbf, 0xa8, 0x4f, 0x6b, 0xb2,
	0xe7, 0x97, 0x0d, 0xda, 0x23, 0x5a, 0xed, 0xd3, 0x45, 0x36, 0xcf, 0x37, 0xfb, 0x3b, 0x71, 0x96,
	0x21, 0xca, 0x88, 0xeb, 0x7f, 0x5d, 0x92, 0xf0, 0x85, 0xea, 0x41, 0xa7, 0xf3, 0x8c, 0xe5, 0x57,
	0x75, 0x98, 0x77, 0x4f, 0x7c, 0x35, 0x09, 0x47, 0xec, 0xa8, 0x55, 0x21, 0x69, 0x5d, 0x87, 0x79,
	0x3c, 0xc0, 0x28, 0xd3, 0x28, 0xe7, 0xd3, 0x59, 0x36, 0xcf, 0xd7, 0xf5, 0x69, 0x2d, 0x0d, 0xbf,
	0x39, 0x90, 0x39, 0xcf, 0x2c, 0x37, 0xf1, 0xee, 0x6a, 0xac, 0x51, 0xb1, 0x8f, 0x87, 0x89, 0x6a,
	0x6a, 0xc1, 0x6a, 0x41, 0x4e, 0x4b, 0xad, 0x6c, 0x28, 0x29, 0x23, 0x82, 0x0e, 0xfd, 0xf4, 0xb0,
	0x67, 0x04, 0xf3, 0xcb, 0xd8, 0xcf, 0xec, 0xf6, 0x35, 0xba, 0x5e, 0x5a, 0xfa, 0x3a, 0x8a, 0x37,
	0x30, 0xe2, 0xbd, 0x68, 0x96, 0xc1, 0xf5, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x02, 0xb8,
	0xca, 0x65, 0x01, 0x00, 0x00,
}