package core

type SystemInfoResponse struct {
	CPUClockSpeed       int    `json:"cpu_clock_speed"`
	CPUCores            string `json:"cpu_cores"`
	CPUFamily           string `json:"cpu_family"`
	CPUSeries           string `json:"cpu_series"`
	CPUVendor           string `json:"cpu_vendor"`
	EnabledNtp          bool   `json:"enabled_ntp"`
	ExternalPciSlotInfo []struct {
		Occupied   string `json:"Occupied"`
		Recognized string `json:"Recognized"`
		CardName   string `json:"cardName"`
		Slot       string `json:"slot"`
	} `json:"external_pci_slot_info"`
	FirmwareDate       string `json:"firmware_date"`
	FirmwareVer        string `json:"firmware_ver"`
	Model              string `json:"model"`
	NtpServer          string `json:"ntp_server"`
	RAMSize            int    `json:"ram_size"`
	SataDev            []any  `json:"sata_dev"`
	Serial             string `json:"serial"`
	SupportEsata       string `json:"support_esata"`
	SysTemp            int    `json:"sys_temp"`
	SysTempwarn        bool   `json:"sys_tempwarn"`
	Systempwarn        bool   `json:"systempwarn"`
	TemperatureWarning bool   `json:"temperature_warning"`
	Time               string `json:"time"`
	TimeZone           string `json:"time_zone"`
	TimeZoneDesc       string `json:"time_zone_desc"`
	UpTime             string `json:"up_time"`
	UsbDev             []struct {
		Cls      string `json:"cls"`
		Pid      string `json:"pid"`
		Producer string `json:"producer"`
		Product  string `json:"product"`
		Rev      string `json:"rev"`
		Vid      string `json:"vid"`
	} `json:"usb_dev"`
}
