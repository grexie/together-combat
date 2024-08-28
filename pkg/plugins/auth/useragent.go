package auth

type UserAgentBrowser interface {
	Family() *string
	Version() *string
	Major() *string
	Minor() *string
}

type UserAgentDevice interface {
	Brand() *string
	Family() *string
	Model() *string
}

type UserAgentOS interface {
	Family() *string
	Version() *string
	Major() *string
	Minor() *string
	Patch() *string
	PatchMinor() *string
}

type UserAgent interface {
	Raw() string
	Browser() UserAgentBrowser
	Device() UserAgentDevice
	OS() UserAgentOS
}