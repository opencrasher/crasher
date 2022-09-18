package entities

import (
	"time"
)

type CoreDumpStatus int

const (
	ToDoCoreDumpStatus CoreDumpStatus = iota
	InProgressCoreDumpStatus
	SolvedCoreDumpStatus
	RejectedCoreDumpStatus
)

type CoreDump struct {
	ID         string
	OsInfo     *OSInfo
	AppInfo    *AppInfo
	Status     CoreDumpStatus
	Data       string
	Timestamp  time.Time
	Extensions []Extension
}

type Extension struct {
	Key   string
	Value string
}

func NewCoreDump() *CoreDump {
	return &CoreDump{}
}

func (c *CoreDump) GetOSInfo() *OSInfo {
	return c.OsInfo
}

func (c *CoreDump) GetAppInfo() *AppInfo {
	return c.AppInfo
}

func (c *CoreDump) GetStatus() CoreDumpStatus {
	return c.Status
}

func (c *CoreDump) GetData() string {
	return c.Data
}

func (c *CoreDump) GetTimestamp() time.Time {
	return c.Timestamp
}

func (c *CoreDump) GetExtension(index int) *Extension {
	return &c.Extensions[index]
}

func (c *CoreDump) GetExtensions() *[]Extension {
	return &c.Extensions
}

func (c *CoreDump) SetOSInfo(info *OSInfo) {
	c.OsInfo = info
}

func (c *CoreDump) SetAppInfo(info *AppInfo) {
	c.AppInfo = info
}

func (c *CoreDump) SetStatus(status CoreDumpStatus) {
	c.Status = status
}

func (c *CoreDump) SetData(data string) {
	c.Data = data
}

func (c *CoreDump) SetTimestamp(timestamp time.Time) {
	c.Timestamp = timestamp
}

func (c *CoreDump) SetExtensions(key, value string) {
	extension := &Extension{Key: key, Value: value}
	c.Extensions = append(c.Extensions, *extension)
}
