// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacemetering

type UsageRecordResultStatus string

// Enum values for UsageRecordResultStatus
const (
	UsageRecordResultStatusSuccess               UsageRecordResultStatus = "Success"
	UsageRecordResultStatusCustomerNotSubscribed UsageRecordResultStatus = "CustomerNotSubscribed"
	UsageRecordResultStatusDuplicateRecord       UsageRecordResultStatus = "DuplicateRecord"
)

func (enum UsageRecordResultStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UsageRecordResultStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
