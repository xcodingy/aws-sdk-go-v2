// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// An address book with attributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/AddressBook
type AddressBook struct {
	_ struct{} `type:"structure"`

	// The ARN of the address book.
	AddressBookArn *string `type:"string"`

	// The description of the address book.
	Description *string `min:"1" type:"string"`

	// The name of the address book.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s AddressBook) String() string {
	return awsutil.Prettify(s)
}

// Information related to an address book.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/AddressBookData
type AddressBookData struct {
	_ struct{} `type:"structure"`

	// The ARN of the address book.
	AddressBookArn *string `type:"string"`

	// The description of the address book.
	Description *string `min:"1" type:"string"`

	// The name of the address book.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s AddressBookData) String() string {
	return awsutil.Prettify(s)
}

// Usage report with specified parameters.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/BusinessReport
type BusinessReport struct {
	_ struct{} `type:"structure"`

	// The time of report delivery.
	DeliveryTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The download link where a user can download the report.
	DownloadUrl *string `type:"string"`

	// The failure code.
	FailureCode BusinessReportFailureCode `type:"string" enum:"true"`

	// The S3 location of the output reports.
	S3Location *BusinessReportS3Location `type:"structure"`

	// The status of the report generation execution (RUNNING, SUCCEEDED, or FAILED).
	Status BusinessReportStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s BusinessReport) String() string {
	return awsutil.Prettify(s)
}

// The content range of the report.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/BusinessReportContentRange
type BusinessReportContentRange struct {
	_ struct{} `type:"structure"`

	// The interval of the content range.
	Interval BusinessReportInterval `type:"string" enum:"true"`
}

// String returns the string representation
func (s BusinessReportContentRange) String() string {
	return awsutil.Prettify(s)
}

// The recurrence of the reports.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/BusinessReportRecurrence
type BusinessReportRecurrence struct {
	_ struct{} `type:"structure"`

	// The start date.
	StartDate *string `type:"string"`
}

// String returns the string representation
func (s BusinessReportRecurrence) String() string {
	return awsutil.Prettify(s)
}

// The S3 location of the output reports.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/BusinessReportS3Location
type BusinessReportS3Location struct {
	_ struct{} `type:"structure"`

	// The S3 bucket name of the output reports.
	BucketName *string `type:"string"`

	// The path of the business report.
	Path *string `type:"string"`
}

// String returns the string representation
func (s BusinessReportS3Location) String() string {
	return awsutil.Prettify(s)
}

// The schedule of the usage report.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/BusinessReportSchedule
type BusinessReportSchedule struct {
	_ struct{} `type:"structure"`

	// The content range of the reports.
	ContentRange *BusinessReportContentRange `type:"structure"`

	// The format of the generated report (individual CSV files or zipped files
	// of individual files).
	Format BusinessReportFormat `type:"string" enum:"true"`

	// The details of the last business report delivery for a specified time interval.
	LastBusinessReport *BusinessReport `type:"structure"`

	// The recurrence of the reports.
	Recurrence *BusinessReportRecurrence `type:"structure"`

	// The S3 bucket name of the output reports.
	S3BucketName *string `type:"string"`

	// The S3 key where the report is delivered.
	S3KeyPrefix *string `type:"string"`

	// The ARN of the business report schedule.
	ScheduleArn *string `type:"string"`

	// The name identifier of the schedule.
	ScheduleName *string `type:"string"`
}

// String returns the string representation
func (s BusinessReportSchedule) String() string {
	return awsutil.Prettify(s)
}

// The skill store category that is shown. Alexa skills are assigned a specific
// skill category during creation, such as News, Social, and Sports.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/Category
type Category struct {
	_ struct{} `type:"structure"`

	// The ID of the skill store category.
	CategoryId *int64 `min:"1" type:"long"`

	// The name of the skill store category.
	CategoryName *string `type:"string"`
}

// String returns the string representation
func (s Category) String() string {
	return awsutil.Prettify(s)
}

// The default conference provider that is used if no other scheduled meetings
// are detected.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ConferencePreference
type ConferencePreference struct {
	_ struct{} `type:"structure"`

	// The ARN of the default conference provider.
	DefaultConferenceProviderArn *string `type:"string"`
}

// String returns the string representation
func (s ConferencePreference) String() string {
	return awsutil.Prettify(s)
}

// An entity that provides a conferencing solution. Alexa for Business acts
// as the voice interface and mediator that connects users to their preferred
// conference provider. Examples of conference providers include Amazon Chime,
// Zoom, Cisco, and Polycom.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ConferenceProvider
type ConferenceProvider struct {
	_ struct{} `type:"structure"`

	// The ARN of the newly created conference provider.
	Arn *string `type:"string"`

	// The IP endpoint and protocol for calling.
	IPDialIn *IPDialIn `type:"structure"`

	// The meeting settings for the conference provider.
	MeetingSetting *MeetingSetting `type:"structure"`

	// The name of the conference provider.
	Name *string `min:"1" type:"string"`

	// The information for PSTN conferencing.
	PSTNDialIn *PSTNDialIn `type:"structure"`

	// The type of conference providers.
	Type ConferenceProviderType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ConferenceProvider) String() string {
	return awsutil.Prettify(s)
}

// A contact with attributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/Contact
type Contact struct {
	_ struct{} `type:"structure"`

	// The ARN of the contact.
	ContactArn *string `type:"string"`

	// The name of the contact to display on the console.
	DisplayName *string `min:"1" type:"string"`

	// The first name of the contact, used to call the contact on the device.
	FirstName *string `min:"1" type:"string"`

	// The last name of the contact, used to call the contact on the device.
	LastName *string `min:"1" type:"string"`

	// The phone number of the contact.
	PhoneNumber *string `type:"string"`
}

// String returns the string representation
func (s Contact) String() string {
	return awsutil.Prettify(s)
}

// Information related to a contact.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ContactData
type ContactData struct {
	_ struct{} `type:"structure"`

	// The ARN of the contact.
	ContactArn *string `type:"string"`

	// The name of the contact to display on the console.
	DisplayName *string `min:"1" type:"string"`

	// The first name of the contact, used to call the contact on the device.
	FirstName *string `min:"1" type:"string"`

	// The last name of the contact, used to call the contact on the device.
	LastName *string `min:"1" type:"string"`

	// The phone number of the contact.
	PhoneNumber *string `type:"string"`
}

// String returns the string representation
func (s ContactData) String() string {
	return awsutil.Prettify(s)
}

// The details about the developer that published the skill.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeveloperInfo
type DeveloperInfo struct {
	_ struct{} `type:"structure"`

	// The name of the developer.
	DeveloperName *string `type:"string"`

	// The email of the developer.
	Email *string `min:"1" type:"string"`

	// The URL of the privacy policy.
	PrivacyPolicy *string `type:"string"`

	// The website of the developer.
	Url *string `type:"string"`
}

// String returns the string representation
func (s DeveloperInfo) String() string {
	return awsutil.Prettify(s)
}

// A device with attributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/Device
type Device struct {
	_ struct{} `type:"structure"`

	// The ARN of a device.
	DeviceArn *string `type:"string"`

	// The name of a device.
	DeviceName *string `min:"2" type:"string"`

	// The serial number of a device.
	DeviceSerialNumber *string `type:"string"`

	// The status of a device. If the status is not READY, check the DeviceStatusInfo
	// value for details.
	DeviceStatus DeviceStatus `type:"string" enum:"true"`

	// Detailed information about a device's status.
	DeviceStatusInfo *DeviceStatusInfo `type:"structure"`

	// The type of a device.
	DeviceType *string `type:"string"`

	// The MAC address of a device.
	MacAddress *string `type:"string"`

	// The room ARN of a device.
	RoomArn *string `type:"string"`

	// The software version of a device.
	SoftwareVersion *string `type:"string"`
}

// String returns the string representation
func (s Device) String() string {
	return awsutil.Prettify(s)
}

// Device attributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeviceData
type DeviceData struct {
	_ struct{} `type:"structure"`

	// The ARN of a device.
	DeviceArn *string `type:"string"`

	// The name of a device.
	DeviceName *string `min:"2" type:"string"`

	// The serial number of a device.
	DeviceSerialNumber *string `type:"string"`

	// The status of a device.
	DeviceStatus DeviceStatus `type:"string" enum:"true"`

	// Detailed information about a device's status.
	DeviceStatusInfo *DeviceStatusInfo `type:"structure"`

	// The type of a device.
	DeviceType *string `type:"string"`

	// The MAC address of a device.
	MacAddress *string `type:"string"`

	// The room ARN associated with a device.
	RoomArn *string `type:"string"`

	// The name of the room associated with a device.
	RoomName *string `min:"1" type:"string"`

	// The software version of a device.
	SoftwareVersion *string `type:"string"`
}

// String returns the string representation
func (s DeviceData) String() string {
	return awsutil.Prettify(s)
}

// The list of device events.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeviceEvent
type DeviceEvent struct {
	_ struct{} `type:"structure"`

	// The time (in epoch) when the event occurred.
	Timestamp *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The type of device event.
	Type DeviceEventType `type:"string" enum:"true"`

	// The value of the event.
	Value *string `type:"string"`
}

// String returns the string representation
func (s DeviceEvent) String() string {
	return awsutil.Prettify(s)
}

// Details of a device’s status.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeviceStatusDetail
type DeviceStatusDetail struct {
	_ struct{} `type:"structure"`

	// The device status detail code.
	Code DeviceStatusDetailCode `type:"string" enum:"true"`
}

// String returns the string representation
func (s DeviceStatusDetail) String() string {
	return awsutil.Prettify(s)
}

// Detailed information about a device's status.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeviceStatusInfo
type DeviceStatusInfo struct {
	_ struct{} `type:"structure"`

	// The latest available information about the connection status of a device.
	ConnectionStatus ConnectionStatus `type:"string" enum:"true"`

	// One or more device status detail descriptions.
	DeviceStatusDetails []DeviceStatusDetail `type:"list"`
}

// String returns the string representation
func (s DeviceStatusInfo) String() string {
	return awsutil.Prettify(s)
}

// A filter name and value pair that is used to return a more specific list
// of results. Filters can be used to match a set of resources by various criteria.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/Filter
type Filter struct {
	_ struct{} `type:"structure"`

	// The key of a filter.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The values of a filter.
	//
	// Values is a required field
	Values []string `type:"list" required:"true"`
}

// String returns the string representation
func (s Filter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Filter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Filter"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Values == nil {
		invalidParams.Add(aws.NewErrParamRequired("Values"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The details of the gateway.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/Gateway
type Gateway struct {
	_ struct{} `type:"structure"`

	// The ARN of the gateway.
	Arn *string `type:"string"`

	// The description of the gateway.
	Description *string `type:"string"`

	// The ARN of the gateway group that the gateway is associated to.
	GatewayGroupArn *string `type:"string"`

	// The name of the gateway.
	Name *string `min:"1" type:"string"`

	// The software version of the gateway. The gateway automatically updates its
	// software version during normal operation.
	SoftwareVersion *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Gateway) String() string {
	return awsutil.Prettify(s)
}

// The details of the gateway group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GatewayGroup
type GatewayGroup struct {
	_ struct{} `type:"structure"`

	// The ARN of the gateway group.
	Arn *string `type:"string"`

	// The description of the gateway group.
	Description *string `type:"string"`

	// The name of the gateway group.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GatewayGroup) String() string {
	return awsutil.Prettify(s)
}

// The summary of a gateway group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GatewayGroupSummary
type GatewayGroupSummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the gateway group.
	Arn *string `type:"string"`

	// The description of the gateway group.
	Description *string `type:"string"`

	// The name of the gateway group.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GatewayGroupSummary) String() string {
	return awsutil.Prettify(s)
}

// The summary of a gateway.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GatewaySummary
type GatewaySummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the gateway.
	Arn *string `type:"string"`

	// The description of the gateway.
	Description *string `type:"string"`

	// The ARN of the gateway group that the gateway is associated to.
	GatewayGroupArn *string `type:"string"`

	// The name of the gateway.
	Name *string `min:"1" type:"string"`

	// The software version of the gateway. The gateway automatically updates its
	// software version during normal operation.
	SoftwareVersion *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GatewaySummary) String() string {
	return awsutil.Prettify(s)
}

// The IP endpoint and protocol for calling.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/IPDialIn
type IPDialIn struct {
	_ struct{} `type:"structure"`

	// The protocol, including SIP, SIPS, and H323.
	//
	// CommsProtocol is a required field
	CommsProtocol CommsProtocol `type:"string" required:"true" enum:"true"`

	// The IP address.
	//
	// Endpoint is a required field
	Endpoint *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s IPDialIn) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *IPDialIn) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "IPDialIn"}
	if len(s.CommsProtocol) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("CommsProtocol"))
	}

	if s.Endpoint == nil {
		invalidParams.Add(aws.NewErrParamRequired("Endpoint"))
	}
	if s.Endpoint != nil && len(*s.Endpoint) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Endpoint", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The values that indicate whether a pin is always required (YES), never required
// (NO), or OPTIONAL.
//
//    * If YES, Alexa will always ask for a meeting pin.
//
//    * If NO, Alexa will never ask for a meeting pin.
//
//    * If OPTIONAL, Alexa will ask if you have a meeting pin and if the customer
//    responds with yes, it will ask for the meeting pin.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/MeetingSetting
type MeetingSetting struct {
	_ struct{} `type:"structure"`

	// The values that indicate whether the pin is always required.
	//
	// RequirePin is a required field
	RequirePin RequirePin `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s MeetingSetting) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MeetingSetting) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MeetingSetting"}
	if len(s.RequirePin) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RequirePin"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The information for public switched telephone network (PSTN) conferencing.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/PSTNDialIn
type PSTNDialIn struct {
	_ struct{} `type:"structure"`

	// The zip code.
	//
	// CountryCode is a required field
	CountryCode *string `type:"string" required:"true"`

	// The delay duration before Alexa enters the conference ID with dual-tone multi-frequency
	// (DTMF). Each number on the dial pad corresponds to a DTMF tone, which is
	// how we send data over the telephone network.
	//
	// OneClickIdDelay is a required field
	OneClickIdDelay *string `min:"1" type:"string" required:"true"`

	// The delay duration before Alexa enters the conference pin with dual-tone
	// multi-frequency (DTMF). Each number on the dial pad corresponds to a DTMF
	// tone, which is how we send data over the telephone network.
	//
	// OneClickPinDelay is a required field
	OneClickPinDelay *string `min:"1" type:"string" required:"true"`

	// The phone number to call to join the conference.
	//
	// PhoneNumber is a required field
	PhoneNumber *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PSTNDialIn) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PSTNDialIn) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PSTNDialIn"}

	if s.CountryCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("CountryCode"))
	}

	if s.OneClickIdDelay == nil {
		invalidParams.Add(aws.NewErrParamRequired("OneClickIdDelay"))
	}
	if s.OneClickIdDelay != nil && len(*s.OneClickIdDelay) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OneClickIdDelay", 1))
	}

	if s.OneClickPinDelay == nil {
		invalidParams.Add(aws.NewErrParamRequired("OneClickPinDelay"))
	}
	if s.OneClickPinDelay != nil && len(*s.OneClickPinDelay) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OneClickPinDelay", 1))
	}

	if s.PhoneNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("PhoneNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A room profile with attributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/Profile
type Profile struct {
	_ struct{} `type:"structure"`

	// The address of a room profile.
	Address *string `min:"1" type:"string"`

	// The ARN of the address book.
	AddressBookArn *string `type:"string"`

	// The distance unit of a room profile.
	DistanceUnit DistanceUnit `type:"string" enum:"true"`

	// Retrieves if the profile is default or not.
	IsDefault *bool `type:"boolean"`

	// The max volume limit of a room profile.
	MaxVolumeLimit *int64 `type:"integer"`

	// The PSTN setting of a room profile.
	PSTNEnabled *bool `type:"boolean"`

	// The ARN of a room profile.
	ProfileArn *string `type:"string"`

	// The name of a room profile.
	ProfileName *string `min:"1" type:"string"`

	// The setup mode of a room profile.
	SetupModeDisabled *bool `type:"boolean"`

	// The temperature unit of a room profile.
	TemperatureUnit TemperatureUnit `type:"string" enum:"true"`

	// The time zone of a room profile.
	Timezone *string `min:"1" type:"string"`

	// The wake word of a room profile.
	WakeWord WakeWord `type:"string" enum:"true"`
}

// String returns the string representation
func (s Profile) String() string {
	return awsutil.Prettify(s)
}

// The data of a room profile.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ProfileData
type ProfileData struct {
	_ struct{} `type:"structure"`

	// The address of a room profile.
	Address *string `min:"1" type:"string"`

	// The distance unit of a room profile.
	DistanceUnit DistanceUnit `type:"string" enum:"true"`

	// Retrieves if the profile data is default or not.
	IsDefault *bool `type:"boolean"`

	// The ARN of a room profile.
	ProfileArn *string `type:"string"`

	// The name of a room profile.
	ProfileName *string `min:"1" type:"string"`

	// The temperature unit of a room profile.
	TemperatureUnit TemperatureUnit `type:"string" enum:"true"`

	// The timezone of a room profile.
	Timezone *string `min:"1" type:"string"`

	// The wake word of a room profile.
	WakeWord WakeWord `type:"string" enum:"true"`
}

// String returns the string representation
func (s ProfileData) String() string {
	return awsutil.Prettify(s)
}

// A room with attributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/Room
type Room struct {
	_ struct{} `type:"structure"`

	// The description of a room.
	Description *string `min:"1" type:"string"`

	// The profile ARN of a room.
	ProfileArn *string `type:"string"`

	// The provider calendar ARN of a room.
	ProviderCalendarId *string `type:"string"`

	// The ARN of a room.
	RoomArn *string `type:"string"`

	// The name of a room.
	RoomName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Room) String() string {
	return awsutil.Prettify(s)
}

// The data of a room.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/RoomData
type RoomData struct {
	_ struct{} `type:"structure"`

	// The description of a room.
	Description *string `min:"1" type:"string"`

	// The profile ARN of a room.
	ProfileArn *string `type:"string"`

	// The profile name of a room.
	ProfileName *string `min:"1" type:"string"`

	// The provider calendar ARN of a room.
	ProviderCalendarId *string `type:"string"`

	// The ARN of a room.
	RoomArn *string `type:"string"`

	// The name of a room.
	RoomName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s RoomData) String() string {
	return awsutil.Prettify(s)
}

// A skill parameter associated with a room.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/RoomSkillParameter
type RoomSkillParameter struct {
	_ struct{} `type:"structure"`

	// The parameter key of a room skill parameter. ParameterKey is an enumerated
	// type that only takes “DEFAULT” or “SCOPE” as valid values.
	//
	// ParameterKey is a required field
	ParameterKey *string `min:"1" type:"string" required:"true"`

	// The parameter value of a room skill parameter.
	//
	// ParameterValue is a required field
	ParameterValue *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RoomSkillParameter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RoomSkillParameter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RoomSkillParameter"}

	if s.ParameterKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterKey"))
	}
	if s.ParameterKey != nil && len(*s.ParameterKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ParameterKey", 1))
	}

	if s.ParameterValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterValue"))
	}
	if s.ParameterValue != nil && len(*s.ParameterValue) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ParameterValue", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Granular information about the skill.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SkillDetails
type SkillDetails struct {
	_ struct{} `type:"structure"`

	// The details about what the skill supports organized as bullet points.
	BulletPoints []string `type:"list"`

	// The details about the developer that published the skill.
	DeveloperInfo *DeveloperInfo `type:"structure"`

	// The URL of the end user license agreement.
	EndUserLicenseAgreement *string `type:"string"`

	// The generic keywords associated with the skill that can be used to find a
	// skill.
	GenericKeywords []string `type:"list"`

	// The phrase used to trigger the skill.
	InvocationPhrase *string `type:"string"`

	// The updates added in bullet points.
	NewInThisVersionBulletPoints []string `type:"list"`

	// The description of the product.
	ProductDescription *string `type:"string"`

	// The date when the skill was released.
	ReleaseDate *string `type:"string"`

	// The list of reviews for the skill, including Key and Value pair.
	Reviews map[string]string `type:"map"`

	// The types of skills.
	SkillTypes []string `type:"list"`
}

// String returns the string representation
func (s SkillDetails) String() string {
	return awsutil.Prettify(s)
}

// A skill group with attributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SkillGroup
type SkillGroup struct {
	_ struct{} `type:"structure"`

	// The description of a skill group.
	Description *string `min:"1" type:"string"`

	// The ARN of a skill group.
	SkillGroupArn *string `type:"string"`

	// The name of a skill group.
	SkillGroupName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s SkillGroup) String() string {
	return awsutil.Prettify(s)
}

// The attributes of a skill group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SkillGroupData
type SkillGroupData struct {
	_ struct{} `type:"structure"`

	// The description of a skill group.
	Description *string `min:"1" type:"string"`

	// The skill group ARN of a skill group.
	SkillGroupArn *string `type:"string"`

	// The skill group name of a skill group.
	SkillGroupName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s SkillGroupData) String() string {
	return awsutil.Prettify(s)
}

// The summary of skills.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SkillSummary
type SkillSummary struct {
	_ struct{} `type:"structure"`

	// Whether the skill is enabled under the user's account, or if it requires
	// linking to be used.
	EnablementType EnablementType `type:"string" enum:"true"`

	// The ARN of the skill summary.
	SkillId *string `type:"string"`

	// The name of the skill.
	SkillName *string `min:"1" type:"string"`

	// Whether the skill is publicly available or is a private skill.
	SkillType SkillType `min:"1" type:"string" enum:"true"`

	// Linking support for a skill.
	SupportsLinking *bool `type:"boolean"`
}

// String returns the string representation
func (s SkillSummary) String() string {
	return awsutil.Prettify(s)
}

// The detailed information about an Alexa skill.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SkillsStoreSkill
type SkillsStoreSkill struct {
	_ struct{} `type:"structure"`

	// The URL where the skill icon resides.
	IconUrl *string `type:"string"`

	// Sample utterances that interact with the skill.
	SampleUtterances []string `type:"list"`

	// Short description about the skill.
	ShortDescription *string `type:"string"`

	// Information about the skill.
	SkillDetails *SkillDetails `type:"structure"`

	// The ARN of the skill.
	SkillId *string `type:"string"`

	// The name of the skill.
	SkillName *string `min:"1" type:"string"`

	// Linking support for a skill.
	SupportsLinking *bool `type:"boolean"`
}

// String returns the string representation
func (s SkillsStoreSkill) String() string {
	return awsutil.Prettify(s)
}

// A smart home appliance that can connect to a central system. Any domestic
// device can be a smart appliance.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SmartHomeAppliance
type SmartHomeAppliance struct {
	_ struct{} `type:"structure"`

	// The description of the smart home appliance.
	Description *string `type:"string"`

	// The friendly name of the smart home appliance.
	FriendlyName *string `type:"string"`

	// The name of the manufacturer of the smart home appliance.
	ManufacturerName *string `type:"string"`
}

// String returns the string representation
func (s SmartHomeAppliance) String() string {
	return awsutil.Prettify(s)
}

// An object representing a sort criteria.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/Sort
type Sort struct {
	_ struct{} `type:"structure"`

	// The sort key of a sort object.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The sort value of a sort object.
	//
	// Value is a required field
	Value SortValue `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s Sort) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Sort) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Sort"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}
	if len(s.Value) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A key-value pair that can be associated with a resource.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// The key of a tag. Tag keys are case-sensitive.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The value of a tag. Tag values are case-sensitive and can be null.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information related to a user.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UserData
type UserData struct {
	_ struct{} `type:"structure"`

	// The email of a user.
	Email *string `min:"1" type:"string"`

	// The enrollment ARN of a user.
	EnrollmentId *string `type:"string"`

	// The enrollment status of a user.
	EnrollmentStatus EnrollmentStatus `type:"string" enum:"true"`

	// The first name of a user.
	FirstName *string `type:"string"`

	// The last name of a user.
	LastName *string `type:"string"`

	// The ARN of a user.
	UserArn *string `type:"string"`
}

// String returns the string representation
func (s UserData) String() string {
	return awsutil.Prettify(s)
}
