// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

type ContainerStatus string

// Enum values for ContainerStatus
const (
	ContainerStatusActive   ContainerStatus = "ACTIVE"
	ContainerStatusCreating ContainerStatus = "CREATING"
	ContainerStatusDeleting ContainerStatus = "DELETING"
)

func (enum ContainerStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContainerStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MethodName string

// Enum values for MethodName
const (
	MethodNamePut    MethodName = "PUT"
	MethodNameGet    MethodName = "GET"
	MethodNameDelete MethodName = "DELETE"
	MethodNameHead   MethodName = "HEAD"
)

func (enum MethodName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MethodName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
