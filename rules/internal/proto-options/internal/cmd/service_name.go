package cmd

import (
	"regexp"
)

var (
	serviceNameExp   = regexp.MustCompile(`^.*?(?P<ServiceName>[a-zA-Z][\w]*)(\.(?P<Version>v\d+(_alpha|_beta)?))?$`)
	serviceNameIndex = serviceNameExp.SubexpIndex("ServiceName")
	versionIndex     = serviceNameExp.SubexpIndex("Version")
)

type serviceName struct {
	Name    string
	Version string
}

func parseServiceName(packageName string) *serviceName {
	match := serviceNameExp.FindStringSubmatch(packageName)
	if len(match) == 0 {
		return nil
	}

	name := match[serviceNameIndex]
	version := match[versionIndex]
	return &serviceName{
		Name:    name,
		Version: version,
	}
}
