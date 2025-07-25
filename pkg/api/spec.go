package api

import "net/url"

type Method struct {
	API            string         `form:"api"     url:"api"`
	Version        int            `form:"version" url:"version"`
	Method         string         `form:"method"  url:"method"`
	ErrorSummaries ErrorSummaries `form:"-"       url:"-"       json:"-"`
}

func (m Method) AsApiParams() ApiParams {
	return ApiParams{
		Version: m.Version,
		API:     m.API,
		Method:  m.Method,
	}
}

func (l ApiParams) EncodeValues(_ string, _ *url.Values) error {
	return nil
}

type APIMethodLookup map[string]Method

var Spec = map[string]Method{
	"Login": {
		API:     "SYNO.API.Auth",
		Version: 7,
		Method:  "login",
	},
	"API.Auth.Key": {API: "SYNO.API.Auth.Key", Version: 7},
	"API.Auth.Key.Code": {
		API:     "SYNO.API.Auth.Key.Code",
		Version: 7,
	},
	"API.Auth.RedirectURI": {
		API:     "SYNO.API.Auth.RedirectURI",
		Version: 1,
	},
	"API.Auth.Type": {API: "SYNO.API.Auth.Type", Version: 1},
	"API.Auth.UIConfig": {
		API:     "SYNO.API.Auth.UIConfig",
		Version: 1,
	},
	"API.Encryption": {API: "SYNO.API.Encryption", Version: 1},
	"API.Info":       {API: "SYNO.API.Info", Version: 1},
	"API.OTP":        {API: "SYNO.API.OTP", Version: 1},
	"ActiveInsight.CompleteCollection": {
		API:     "SYNO.ActiveInsight.CompleteCollection",
		Version: 2,
	},
	"ActiveInsight.License": {
		API:     "SYNO.ActiveInsight.License",
		Version: 2,
	},
	"ActiveInsight.Package": {
		API:     "SYNO.ActiveInsight.Package",
		Version: 2,
	},
	"ActiveInsight.Setting": {
		API:     "SYNO.ActiveInsight.Setting",
		Version: 2,
	},
	"AudioPlayer": {API: "SYNO.AudioPlayer", Version: 2},
	"AudioPlayer.Stream": {
		API:     "SYNO.AudioPlayer.Stream",
		Version: 2,
	},
	"Auth.ForgotPwd":   {API: "SYNO.Auth.ForgotPwd", Version: 1},
	"Auth.RescueEmail": {API: "SYNO.Auth.RescueEmail", Version: 1},
	"Backup.App":       {API: "SYNO.Backup.App", Version: 1},
	"Backup.App.Backup": {
		API:     "SYNO.Backup.App.Backup",
		Version: 1,
	},
	"Backup.App.Restore": {
		API:     "SYNO.Backup.App.Restore",
		Version: 1,
	},
	"Backup.Config.AutoBackup": {
		API:     "SYNO.Backup.Config.AutoBackup",
		Version: 2,
	},
	"Backup.Config.Backup": {
		API:     "SYNO.Backup.Config.Backup",
		Version: 2,
	},
	"Backup.Config.Restore": {
		API:     "SYNO.Backup.Config.Restore",
		Version: 2,
	},
	"Backup.ED.Restore": {
		API:     "SYNO.Backup.ED.Restore",
		Version: 2,
	},
	"Backup.ED.Restore.Disks": {
		API:     "SYNO.Backup.ED.Restore.Disks",
		Version: 2,
	},
	"Backup.ED.Restore.Space": {
		API:     "SYNO.Backup.ED.Restore.Space",
		Version: 2,
	},
	"Backup.ED.Restore.Target": {
		API:     "SYNO.Backup.ED.Restore.Target",
		Version: 2,
	},
	"Backup.ED.Restore.Version": {
		API:     "SYNO.Backup.ED.Restore.Version",
		Version: 2,
	},
	"Backup.ED.Target.Config": {
		API:     "SYNO.Backup.ED.Target.Config",
		Version: 2,
	},
	"Backup.Service.NetworkBackup": {
		API:     "SYNO.Backup.Service.NetworkBackup",
		Version: 1,
	},
	"Btrfs.Replica.Core": {
		API:     "SYNO.Btrfs.Replica.Core",
		Version: 1,
	},
	"C2FS.Account": {API: "SYNO.C2FS.Account", Version: 1},
	"C2FS.Bucket":  {API: "SYNO.C2FS.Bucket", Version: 1},
	"C2FS.Conf":    {API: "SYNO.C2FS.Conf", Version: 1},
	"C2FS.File":    {API: "SYNO.C2FS.File", Version: 2},
	"C2FS.Share":   {API: "SYNO.C2FS.Share", Version: 1},
	"C2FS.Share.Status": {
		API:     "SYNO.C2FS.Share.Status",
		Version: 1,
	},
	"C2FS.Transform": {API: "SYNO.C2FS.Transform", Version: 1},
	"C2FS.Utils":     {API: "SYNO.C2FS.Utils", Version: 1},
	"CCC.CLog":       {API: "SYNO.CCC.CLog", Version: 1},
	"CCC.Cluster":    {API: "SYNO.CCC.Cluster", Version: 1},
	"CCC.Cluster.Member": {
		API:     "SYNO.CCC.Cluster.Member",
		Version: 2,
	},
	"CCC.Guest":        {API: "SYNO.CCC.Guest", Version: 1},
	"CCC.Guest.Image":  {API: "SYNO.CCC.Guest.Image", Version: 1},
	"CCC.Guest.VNC":    {API: "SYNO.CCC.Guest.VNC", Version: 1},
	"CCC.GuestReplica": {API: "SYNO.CCC.GuestReplica", Version: 1},
	"CCC.GuestSnap":    {API: "SYNO.CCC.GuestSnap", Version: 1},
	"CCC.GuestSnap.Policy": {
		API:     "SYNO.CCC.GuestSnap.Policy",
		Version: 1,
	},
	"CCC.HA":           {API: "SYNO.CCC.HA", Version: 1},
	"CCC.License":      {API: "SYNO.CCC.License", Version: 1},
	"CCC.License.Pro":  {API: "SYNO.CCC.License.Pro", Version: 1},
	"CCC.License.VDSM": {API: "SYNO.CCC.License.VDSM", Version: 1},
	"CCC.Network":      {API: "SYNO.CCC.Network", Version: 1},
	"CCC.Network.SRIOV": {
		API:     "SYNO.CCC.Network.SRIOV",
		Version: 1,
	},
	"CCC.Notify":      {API: "SYNO.CCC.Notify", Version: 1},
	"CCC.Package":     {API: "SYNO.CCC.Package", Version: 1},
	"CCC.Repo":        {API: "SYNO.CCC.Repo", Version: 1},
	"CCC.Resource":    {API: "SYNO.CCC.Resource", Version: 1},
	"CCC.Utils":       {API: "SYNO.CCC.Utils", Version: 1},
	"CCC.VDisk":       {API: "SYNO.CCC.VDisk", Version: 1},
	"CCC.VNic":        {API: "SYNO.CCC.VNic", Version: 1},
	"Core.ACL":        {API: "SYNO.Core.ACL", Version: 2},
	"Core.ActionPriv": {API: "SYNO.Core.ActionPriv", Version: 1},
	"Core.ActionPriv.Role": {
		API:     "SYNO.Core.ActionPriv.Role",
		Version: 1,
	},
	"Core.AppNotify": {API: "SYNO.Core.AppNotify", Version: 1},
	"Core.AppPortal": {API: "SYNO.Core.AppPortal", Version: 2},
	"Core.AppPortal.AccessControl": {
		API:     "SYNO.Core.AppPortal.AccessControl",
		Version: 1,
	},
	"Core.AppPortal.Config": {
		API:     "SYNO.Core.AppPortal.Config",
		Version: 1,
	},
	"Core.AppPortal.ReverseProxy": {
		API:     "SYNO.Core.AppPortal.ReverseProxy",
		Version: 1,
	},
	"Core.AppPriv":     {API: "SYNO.Core.AppPriv", Version: 2},
	"Core.AppPriv.App": {API: "SYNO.Core.AppPriv.App", Version: 3},
	"Core.AppPriv.Rule": {
		API:     "SYNO.Core.AppPriv.Rule",
		Version: 1,
	},
	"Core.BackgroundTask": {
		API:     "SYNO.Core.BackgroundTask",
		Version: 1,
	},
	"Core.Backup.ED": {API: "SYNO.Core.Backup.ED", Version: 1},
	"Core.BandwidthControl": {
		API:     "SYNO.Core.BandwidthControl",
		Version: 2,
	},
	"Core.BandwidthControl.Protocol": {
		API:     "SYNO.Core.BandwidthControl.Protocol",
		Version: 1,
	},
	"Core.BandwidthControl.Status": {
		API:     "SYNO.Core.BandwidthControl.Status",
		Version: 1,
	},
	"Core.CMS":       {API: "SYNO.Core.CMS", Version: 2},
	"Core.CMS.Cache": {API: "SYNO.Core.CMS.Cache", Version: 2},
	"Core.CMS.Identity": {
		API:     "SYNO.Core.CMS.Identity",
		Version: 1,
	},
	"Core.CMS.Info":   {API: "SYNO.Core.CMS.Info", Version: 1},
	"Core.CMS.Policy": {API: "SYNO.Core.CMS.Policy", Version: 1},
	"Core.CMS.ServerInfo": {
		API:     "SYNO.Core.CMS.ServerInfo",
		Version: 2,
	},
	"Core.CMS.Task":    {API: "SYNO.Core.CMS.Task", Version: 1},
	"Core.CMS.Token":   {API: "SYNO.Core.CMS.Token", Version: 1},
	"Core.Certificate": {API: "SYNO.Core.Certificate", Version: 1},
	"Core.Certificate.CRT": {
		API:     "SYNO.Core.Certificate.CRT",
		Version: 1,
	},
	"Core.Certificate.CSR": {
		API:     "SYNO.Core.Certificate.CSR",
		Version: 1,
	},
	"Core.Certificate.LetsEncrypt": {
		API:     "SYNO.Core.Certificate.LetsEncrypt",
		Version: 1,
	},
	"Core.Certificate.LetsEncrypt.Account": {
		API:     "SYNO.Core.Certificate.LetsEncrypt.Account",
		Version: 1,
	},
	"Core.Certificate.Service": {
		API:     "SYNO.Core.Certificate.Service",
		Version: 1,
	},
	"Core.Certificate.Tencent": {
		API:     "SYNO.Core.Certificate.Tencent",
		Version: 1,
	},
	"Core.CurrentConnection": {
		API:     "SYNO.Core.CurrentConnection",
		Version: 1,
	},
	"Core.DDNS.Ethernet": {
		API:     "SYNO.Core.DDNS.Ethernet",
		Version: 1,
	},
	"Core.DDNS.ExtIP": {API: "SYNO.Core.DDNS.ExtIP", Version: 2},
	"Core.DDNS.Provider": {
		API:     "SYNO.Core.DDNS.Provider",
		Version: 1,
	},
	"Core.DDNS.Record": {API: "SYNO.Core.DDNS.Record", Version: 1},
	"Core.DDNS.Synology": {
		API:     "SYNO.Core.DDNS.Synology",
		Version: 1,
	},
	"Core.DDNS.TWNIC": {API: "SYNO.Core.DDNS.TWNIC", Version: 1},
	"Core.DSMNotify":  {API: "SYNO.Core.DSMNotify", Version: 1},
	"Core.DSMNotify.MailContent": {
		API:     "SYNO.Core.DSMNotify.MailContent",
		Version: 1,
	},
	"Core.DSMNotify.Strings": {
		API:     "SYNO.Core.DSMNotify.Strings",
		Version: 1,
	},
	"Core.DataCollect": {API: "SYNO.Core.DataCollect", Version: 1},
	"Core.DataCollect.Application": {
		API:     "SYNO.Core.DataCollect.Application",
		Version: 1,
	},
	"Core.Desktop.Defs": {
		API:     "SYNO.Core.Desktop.Defs",
		Version: 1,
	},
	"Core.Desktop.Initdata": {
		API:     "SYNO.Core.Desktop.Initdata",
		Version: 1,
	},
	"Core.Desktop.JSUIString": {
		API:     "SYNO.Core.Desktop.JSUIString",
		Version: 1,
	},
	"Core.Desktop.PersonalUpdater": {
		API:     "SYNO.Core.Desktop.PersonalUpdater",
		Version: 1,
	},
	"Core.Desktop.SessionData": {
		API:     "SYNO.Core.Desktop.SessionData",
		Version: 1,
	},
	"Core.Desktop.Timeout": {
		API:     "SYNO.Core.Desktop.Timeout",
		Version: 1,
	},
	"Core.Desktop.UIString": {
		API:     "SYNO.Core.Desktop.UIString",
		Version: 1,
	},
	"Core.Desktop.Upgrade": {
		API:     "SYNO.Core.Desktop.Upgrade",
		Version: 1,
	},
	"Core.Directory.Azure.SSO": {
		API:     "SYNO.Core.Directory.Azure.SSO",
		Version: 1,
	},
	"Core.Directory.Domain": {
		API:     "SYNO.Core.Directory.Domain",
		Version: 3,
	},
	"Core.Directory.Domain.Conf": {
		API:     "SYNO.Core.Directory.Domain.Conf",
		Version: 3,
	},
	"Core.Directory.Domain.Schedule": {
		API:     "SYNO.Core.Directory.Domain.Schedule",
		Version: 1,
	},
	"Core.Directory.Domain.Trust": {
		API:     "SYNO.Core.Directory.Domain.Trust",
		Version: 1,
	},
	"Core.Directory.LDAP": {
		API:     "SYNO.Core.Directory.LDAP",
		Version: 2,
	},
	"Core.Directory.LDAP.BaseDN": {
		API:     "SYNO.Core.Directory.LDAP.BaseDN",
		Version: 2,
	},
	"Core.Directory.LDAP.Login.Notify": {
		API:     "SYNO.Core.Directory.LDAP.Login.Notify",
		Version: 1,
	},
	"Core.Directory.LDAP.Profile": {
		API:     "SYNO.Core.Directory.LDAP.Profile",
		Version: 2,
	},
	"Core.Directory.LDAP.Refresh": {
		API:     "SYNO.Core.Directory.LDAP.Refresh",
		Version: 1,
	},
	"Core.Directory.LDAP.User": {
		API:     "SYNO.Core.Directory.LDAP.User",
		Version: 1,
	},
	"Core.Directory.OIDC.SSO": {
		API:     "SYNO.Core.Directory.OIDC.SSO",
		Version: 1,
	},
	"Core.Directory.SSO": {
		API:     "SYNO.Core.Directory.SSO",
		Version: 2,
	},
	"Core.Directory.SSO.CAS": {
		API:     "SYNO.Core.Directory.SSO.CAS",
		Version: 1,
	},
	"Core.Directory.SSO.IWA": {
		API:     "SYNO.Core.Directory.SSO.IWA",
		Version: 1,
	},
	"Core.Directory.SSO.Profile": {
		API:     "SYNO.Core.Directory.SSO.Profile",
		Version: 1,
	},
	"Core.Directory.SSO.SAML": {
		API:     "SYNO.Core.Directory.SSO.SAML",
		Version: 1,
	},
	"Core.Directory.SSO.SAML.Metadata": {
		API:     "SYNO.Core.Directory.SSO.SAML.Metadata",
		Version: 1,
	},
	"Core.Directory.SSO.SAML.Status": {
		API:     "SYNO.Core.Directory.SSO.SAML.Status",
		Version: 1,
	},
	"Core.Directory.SSO.Setting": {
		API:     "SYNO.Core.Directory.SSO.Setting",
		Version: 1,
	},
	"Core.Directory.SSO.Status": {
		API:     "SYNO.Core.Directory.SSO.Status",
		Version: 1,
	},
	"Core.Directory.SSO.utils": {
		API:     "SYNO.Core.Directory.SSO.utils",
		Version: 1,
	},
	"Core.Directory.WebSphere.SSO": {
		API:     "SYNO.Core.Directory.WebSphere.SSO",
		Version: 1,
	},
	"Core.DirectoryServiceCheck.Common": {
		API:     "SYNO.Core.DirectoryServiceCheck.Common",
		Version: 2,
	},
	"Core.DirectoryServiceCheck.Debug": {
		API:     "SYNO.Core.DirectoryServiceCheck.Debug",
		Version: 1,
	},
	"Core.DirectoryServiceCheck.Domain": {
		API:     "SYNO.Core.DirectoryServiceCheck.Domain",
		Version: 1,
	},
	"Core.DirectoryServiceCheck.DomainJoin": {
		API:     "SYNO.Core.DirectoryServiceCheck.DomainJoin",
		Version: 1,
	},
	"Core.DirectoryServiceCheck.DomainService": {
		API:     "SYNO.Core.DirectoryServiceCheck.DomainService",
		Version: 1,
	},
	"Core.DirectoryServiceCheck.DomainValidation": {
		API:     "SYNO.Core.DirectoryServiceCheck.DomainValidation",
		Version: 1,
	},
	"Core.DirectoryServiceCheck.LDAP": {
		API:     "SYNO.Core.DirectoryServiceCheck.LDAP",
		Version: 1,
	},
	"Core.DirectoryServiceCheck.Progress": {
		API:     "SYNO.Core.DirectoryServiceCheck.Progress",
		Version: 2,
	},
	"Core.DisableAdmin": {
		API:     "SYNO.Core.DisableAdmin",
		Version: 1,
	},
	"Core.EW.Info": {API: "SYNO.Core.EW.Info", Version: 1},
	"Core.EventScheduler": {
		API:     "SYNO.Core.EventScheduler",
		Version: 1,
	},
	"Core.EventScheduler.Root": {
		API:     "SYNO.Core.EventScheduler.Root",
		Version: 1,
	},
	"Core.ExternalDevice.Bluetooth": {
		API:     "SYNO.Core.ExternalDevice.Bluetooth",
		Version: 2,
	},
	"Core.ExternalDevice.Bluetooth.Device": {
		API:     "SYNO.Core.ExternalDevice.Bluetooth.Device",
		Version: 1,
	},
	"Core.ExternalDevice.Bluetooth.Settings": {
		API:     "SYNO.Core.ExternalDevice.Bluetooth.Settings",
		Version: 1,
	},
	"Core.ExternalDevice.DefaultPermission": {
		API:     "SYNO.Core.ExternalDevice.DefaultPermission",
		Version: 1,
	},
	"Core.ExternalDevice.Printer": {
		API:     "SYNO.Core.ExternalDevice.Printer",
		Version: 1,
	},
	"Core.ExternalDevice.Printer.BonjourSharing": {
		API:     "SYNO.Core.ExternalDevice.Printer.BonjourSharing",
		Version: 1,
	},
	"Core.ExternalDevice.Printer.Driver": {
		API:     "SYNO.Core.ExternalDevice.Printer.Driver",
		Version: 1,
	},
	"Core.ExternalDevice.Printer.Network": {
		API:     "SYNO.Core.ExternalDevice.Printer.Network",
		Version: 1,
	},
	"Core.ExternalDevice.Printer.Network.Host": {
		API:     "SYNO.Core.ExternalDevice.Printer.Network.Host",
		Version: 1,
	},
	"Core.ExternalDevice.Printer.OAuth": {
		API:     "SYNO.Core.ExternalDevice.Printer.OAuth",
		Version: 1,
	},
	"Core.ExternalDevice.Printer.USB": {
		API:     "SYNO.Core.ExternalDevice.Printer.USB",
		Version: 1,
	},
	"Core.ExternalDevice.Storage.EUnit": {
		API:     "SYNO.Core.ExternalDevice.Storage.EUnit",
		Version: 1,
	},
	"Core.ExternalDevice.Storage.Setting": {
		API:     "SYNO.Core.ExternalDevice.Storage.Setting",
		Version: 1,
	},
	"Core.ExternalDevice.Storage.USB": {
		API:     "SYNO.Core.ExternalDevice.Storage.USB",
		Version: 1,
	},
	"Core.ExternalDevice.Storage.eSATA": {
		API:     "SYNO.Core.ExternalDevice.Storage.eSATA",
		Version: 1,
	},
	"Core.ExternalDevice.UPS": {
		API:     "SYNO.Core.ExternalDevice.UPS",
		Version: 1,
	},
	"Core.Factory.Config": {
		API:     "SYNO.Core.Factory.Config",
		Version: 1,
	},
	"Core.Factory.Manutild": {
		API:     "SYNO.Core.Factory.Manutild",
		Version: 1,
	},
	"Core.File": {API: "SYNO.Core.File", Version: 2},
	"Core.File.Thumbnail": {
		API:     "SYNO.Core.File.Thumbnail",
		Version: 1,
	},
	"Core.FileHandle": {API: "SYNO.Core.FileHandle", Version: 1},
	"Core.FileServ.AFP": {
		API:     "SYNO.Core.FileServ.AFP",
		Version: 2,
	},
	"Core.FileServ.FTP": {
		API:     "SYNO.Core.FileServ.FTP",
		Version: 3,
	},
	"Core.FileServ.FTP.ChrootUser": {
		API:     "SYNO.Core.FileServ.FTP.ChrootUser",
		Version: 2,
	},
	"Core.FileServ.FTP.SFTP": {
		API:     "SYNO.Core.FileServ.FTP.SFTP",
		Version: 1,
	},
	"Core.FileServ.FTP.Security": {
		API:     "SYNO.Core.FileServ.FTP.Security",
		Version: 1,
	},
	"Core.FileServ.NFS": {
		API:     "SYNO.Core.FileServ.NFS",
		Version: 3,
	},
	"Core.FileServ.NFS.AdvancedSetting": {
		API:     "SYNO.Core.FileServ.NFS.AdvancedSetting",
		Version: 1,
	},
	"Core.FileServ.NFS.ConfBackup": {
		API:     "SYNO.Core.FileServ.NFS.ConfBackup",
		Version: 1,
	},
	"Core.FileServ.NFS.IDMap": {
		API:     "SYNO.Core.FileServ.NFS.IDMap",
		Version: 1,
	},
	"Core.FileServ.NFS.Kerberos": {
		API:     "SYNO.Core.FileServ.NFS.Kerberos",
		Version: 1,
	},
	"Core.FileServ.NFS.SharePrivilege": {
		API:     "SYNO.Core.FileServ.NFS.SharePrivilege",
		Version: 1,
	},
	"Core.FileServ.ReflinkCopy": {
		API:     "SYNO.Core.FileServ.ReflinkCopy",
		Version: 1,
	},
	"Core.FileServ.Rsync.Account": {
		API:     "SYNO.Core.FileServ.Rsync.Account",
		Version: 1,
	},
	"Core.FileServ.SMB": {
		API:     "SYNO.Core.FileServ.SMB",
		Version: 3,
	},
	"Core.FileServ.SMB.ConfBackup": {
		API:     "SYNO.Core.FileServ.SMB.ConfBackup",
		Version: 1,
	},
	"Core.FileServ.SMB.Control": {
		API:     "SYNO.Core.FileServ.SMB.Control",
		Version: 1,
	},
	"Core.FileServ.SMB.MSDFS": {
		API:     "SYNO.Core.FileServ.SMB.MSDFS",
		Version: 1,
	},
	"Core.FileServ.ServiceDiscovery": {
		API:     "SYNO.Core.FileServ.ServiceDiscovery",
		Version: 1,
	},
	"Core.FileServ.ServiceDiscovery.WSTransfer": {
		API:     "SYNO.Core.FileServ.ServiceDiscovery.WSTransfer",
		Version: 1,
	},
	"Core.Findhost": {API: "SYNO.Core.Findhost", Version: 1},
	"Core.Group":    {API: "SYNO.Core.Group", Version: 1},
	"Core.Group.ExtraAdmin": {
		API:     "SYNO.Core.Group.ExtraAdmin",
		Version: 1,
	},
	"Core.Group.Member": {
		API:     "SYNO.Core.Group.Member",
		Version: 1,
	},
	"Core.Group.ValidLocalAdmin": {
		API:     "SYNO.Core.Group.ValidLocalAdmin",
		Version: 1,
	},
	"Core.GroupSettings": {
		API:     "SYNO.Core.GroupSettings",
		Version: 1,
	},
	"Core.Hardware.BeepControl": {
		API:     "SYNO.Core.Hardware.BeepControl",
		Version: 1,
	},
	"Core.Hardware.FanSpeed": {
		API:     "SYNO.Core.Hardware.FanSpeed",
		Version: 1,
	},
	"Core.Hardware.Hibernation": {
		API:     "SYNO.Core.Hardware.Hibernation",
		Version: 1,
	},
	"Core.Hardware.LCM": {
		API:     "SYNO.Core.Hardware.LCM",
		Version: 1,
	},
	"Core.Hardware.Led.Brightness": {
		API:     "SYNO.Core.Hardware.Led.Brightness",
		Version: 1,
	},
	"Core.Hardware.MemoryLayout": {
		API:     "SYNO.Core.Hardware.MemoryLayout",
		Version: 1,
	},
	"Core.Hardware.NeedReboot": {
		API:     "SYNO.Core.Hardware.NeedReboot",
		Version: 1,
	},
	"Core.Hardware.OOBManagement": {
		API:     "SYNO.Core.Hardware.OOBManagement",
		Version: 1,
	},
	"Core.Hardware.PowerRecovery": {
		API:     "SYNO.Core.Hardware.PowerRecovery",
		Version: 1,
	},
	"Core.Hardware.PowerSchedule": {
		API:     "SYNO.Core.Hardware.PowerSchedule",
		Version: 1,
	},
	"Core.Hardware.RemoteFanStatus": {
		API:     "SYNO.Core.Hardware.RemoteFanStatus",
		Version: 1,
	},
	"Core.Hardware.SpectreMeltdown": {
		API:     "SYNO.Core.Hardware.SpectreMeltdown",
		Version: 1,
	},
	"Core.Hardware.VideoTranscoding": {
		API:     "SYNO.Core.Hardware.VideoTranscoding",
		Version: 1,
	},
	"Core.Hardware.ZRAM": {
		API:     "SYNO.Core.Hardware.ZRAM",
		Version: 1,
	},
	"Core.Help": {API: "SYNO.Core.Help", Version: 1},
	"Core.ISCSI.FCTarget": {
		API:     "SYNO.Core.ISCSI.FCTarget",
		Version: 1,
	},
	"Core.ISCSI.Host": {API: "SYNO.Core.ISCSI.Host", Version: 1},
	"Core.ISCSI.LUN":  {API: "SYNO.Core.ISCSI.LUN", Version: 1},
	"Core.ISCSI.Lunbkp": {
		API:     "SYNO.Core.ISCSI.Lunbkp",
		Version: 1,
	},
	"Core.ISCSI.Node": {API: "SYNO.Core.ISCSI.Node", Version: 1},
	"Core.ISCSI.Replication": {
		API:     "SYNO.Core.ISCSI.Replication",
		Version: 1,
	},
	"Core.ISCSI.Target": {
		API:     "SYNO.Core.ISCSI.Target",
		Version: 1,
	},
	"Core.ISCSI.VMware": {
		API:     "SYNO.Core.ISCSI.VMware",
		Version: 1,
	},
	"Core.MediaIndexing": {
		API:     "SYNO.Core.MediaIndexing",
		Version: 1,
	},
	"Core.MediaIndexing.IndexFolder": {
		API:     "SYNO.Core.MediaIndexing.IndexFolder",
		Version: 1,
	},
	"Core.MediaIndexing.MediaConverter": {
		API:     "SYNO.Core.MediaIndexing.MediaConverter",
		Version: 2,
	},
	"Core.MediaIndexing.Scheduler": {
		API:     "SYNO.Core.MediaIndexing.Scheduler",
		Version: 1,
	},
	"Core.MediaIndexing.ThumbnailQuality": {
		API:     "SYNO.Core.MediaIndexing.ThumbnailQuality",
		Version: 1,
	},
	"Core.MyDSCenter": {API: "SYNO.Core.MyDSCenter", Version: 2},
	"Core.MyDSCenter.Account": {
		API:     "SYNO.Core.MyDSCenter.Account",
		Version: 1,
	},
	"Core.MyDSCenter.Login": {
		API:     "SYNO.Core.MyDSCenter.Login",
		Version: 1,
	},
	"Core.MyDSCenter.Logout": {
		API:     "SYNO.Core.MyDSCenter.Logout",
		Version: 1,
	},
	"Core.MyDSCenter.Purchase": {
		API:     "SYNO.Core.MyDSCenter.Purchase",
		Version: 1,
	},
	"Core.Network": {API: "SYNO.Core.Network", Version: 2},
	"Core.Network.Authentication": {
		API:     "SYNO.Core.Network.Authentication",
		Version: 1,
	},
	"Core.Network.Authentication.Cert": {
		API:     "SYNO.Core.Network.Authentication.Cert",
		Version: 1,
	},
	"Core.Network.Bond": {
		API:     "SYNO.Core.Network.Bond",
		Version: 2,
	},
	"Core.Network.Ethernet": {
		API:     "SYNO.Core.Network.Ethernet",
		Version: 2,
	},
	"Core.Network.Ethernet.External": {
		API:     "SYNO.Core.Network.Ethernet.External",
		Version: 1,
	},
	"Core.Network.IPv6": {
		API:     "SYNO.Core.Network.IPv6",
		Version: 1,
	},
	"Core.Network.IPv6.Router": {
		API:     "SYNO.Core.Network.IPv6.Router",
		Version: 1,
	},
	"Core.Network.IPv6.Router.Prefix": {
		API:     "SYNO.Core.Network.IPv6.Router.Prefix",
		Version: 1,
	},
	"Core.Network.Interface": {
		API:     "SYNO.Core.Network.Interface",
		Version: 1,
	},
	"Core.Network.MACClone": {
		API:     "SYNO.Core.Network.MACClone",
		Version: 1,
	},
	"Core.Network.OVS": {API: "SYNO.Core.Network.OVS", Version: 1},
	"Core.Network.PPPoE": {
		API:     "SYNO.Core.Network.PPPoE",
		Version: 1,
	},
	"Core.Network.PPPoE.Relay": {
		API:     "SYNO.Core.Network.PPPoE.Relay",
		Version: 1,
	},
	"Core.Network.Proxy": {
		API:     "SYNO.Core.Network.Proxy",
		Version: 1,
	},
	"Core.Network.Router.Gateway.List": {
		API:     "SYNO.Core.Network.Router.Gateway.List",
		Version: 1,
	},
	"Core.Network.Router.Static.Route": {
		API:     "SYNO.Core.Network.Router.Static.Route",
		Version: 1,
	},
	"Core.Network.TrafficControl.RouterRules": {
		API:     "SYNO.Core.Network.TrafficControl.RouterRules",
		Version: 1,
	},
	"Core.Network.TrafficControl.Rules": {
		API:     "SYNO.Core.Network.TrafficControl.Rules",
		Version: 1,
	},
	"Core.Network.UPnPServer": {
		API:     "SYNO.Core.Network.UPnPServer",
		Version: 1,
	},
	"Core.Network.VPN": {API: "SYNO.Core.Network.VPN", Version: 1},
	"Core.Network.VPN.L2TP": {
		API:     "SYNO.Core.Network.VPN.L2TP",
		Version: 1,
	},
	"Core.Network.VPN.OpenVPN": {
		API:     "SYNO.Core.Network.VPN.OpenVPN",
		Version: 1,
	},
	"Core.Network.VPN.OpenVPN.CA": {
		API:     "SYNO.Core.Network.VPN.OpenVPN.CA",
		Version: 1,
	},
	"Core.Network.VPN.OpenVPNWithConf": {
		API:     "SYNO.Core.Network.VPN.OpenVPNWithConf",
		Version: 1,
	},
	"Core.Network.VPN.OpenVPNWithConf.Certs": {
		API:     "SYNO.Core.Network.VPN.OpenVPNWithConf.Certs",
		Version: 1,
	},
	"Core.Network.VPN.PPTP": {
		API:     "SYNO.Core.Network.VPN.PPTP",
		Version: 1,
	},
	"Core.Network.WOL": {API: "SYNO.Core.Network.WOL", Version: 1},
	"Core.NormalUser":  {API: "SYNO.Core.NormalUser", Version: 2},
	"Core.NormalUser.LoginNotify": {
		API:     "SYNO.Core.NormalUser.LoginNotify",
		Version: 1,
	},
	"Core.Notification.Advance.CustomizedData": {
		API:     "SYNO.Core.Notification.Advance.CustomizedData",
		Version: 1,
	},
	"Core.Notification.Advance.FilterSettings": {
		API:     "SYNO.Core.Notification.Advance.FilterSettings",
		Version: 2,
	},
	"Core.Notification.Advance.FilterSettings.Profile": {
		API:     "SYNO.Core.Notification.Advance.FilterSettings.Profile",
		Version: 1,
	},
	"Core.Notification.Advance.FilterSettings.Template": {
		API:     "SYNO.Core.Notification.Advance.FilterSettings.Template",
		Version: 1,
	},
	"Core.Notification.Advance.Variables": {
		API:     "SYNO.Core.Notification.Advance.Variables",
		Version: 1,
	},
	"Core.Notification.Advance.WarningPercentage": {
		API:     "SYNO.Core.Notification.Advance.WarningPercentage",
		Version: 1,
	},
	"Core.Notification.CMS": {
		API:     "SYNO.Core.Notification.CMS",
		Version: 1,
	},
	"Core.Notification.CMS.Conf": {
		API:     "SYNO.Core.Notification.CMS.Conf",
		Version: 2,
	},
	"Core.Notification.Line": {
		API:     "SYNO.Core.Notification.Line",
		Version: 1,
	},
	"Core.Notification.Mail": {
		API:     "SYNO.Core.Notification.Mail",
		Version: 1,
	},
	"Core.Notification.Mail.Auth": {
		API:     "SYNO.Core.Notification.Mail.Auth",
		Version: 1,
	},
	"Core.Notification.Mail.Conf": {
		API:     "SYNO.Core.Notification.Mail.Conf",
		Version: 2,
	},
	"Core.Notification.Mail.Oauth": {
		API:     "SYNO.Core.Notification.Mail.Oauth",
		Version: 1,
	},
	"Core.Notification.Mail.Profile.Conf": {
		API:     "SYNO.Core.Notification.Mail.Profile.Conf",
		Version: 1,
	},
	"Core.Notification.Push": {
		API:     "SYNO.Core.Notification.Push",
		Version: 1,
	},
	"Core.Notification.Push.AuthToken": {
		API:     "SYNO.Core.Notification.Push.AuthToken",
		Version: 1,
	},
	"Core.Notification.Push.Conf": {
		API:     "SYNO.Core.Notification.Push.Conf",
		Version: 1,
	},
	"Core.Notification.Push.Mail": {
		API:     "SYNO.Core.Notification.Push.Mail",
		Version: 2,
	},
	"Core.Notification.Push.Mobile": {
		API:     "SYNO.Core.Notification.Push.Mobile",
		Version: 2,
	},
	"Core.Notification.Push.Webhook.Provider": {
		API:     "SYNO.Core.Notification.Push.Webhook.Provider",
		Version: 2,
	},
	"Core.Notification.SMS": {
		API:     "SYNO.Core.Notification.SMS",
		Version: 2,
	},
	"Core.Notification.SMS.Conf": {
		API:     "SYNO.Core.Notification.SMS.Conf",
		Version: 2,
	},
	"Core.Notification.SMS.Provider": {
		API:     "SYNO.Core.Notification.SMS.Provider",
		Version: 2,
	},
	"Core.Notification.Sysnotify": {
		API:     "SYNO.Core.Notification.Sysnotify",
		Version: 1,
	},
	"Core.OAuth.Scope": {API: "SYNO.Core.OAuth.Scope", Version: 1},
	"Core.OAuth.Server": {
		API:     "SYNO.Core.OAuth.Server",
		Version: 1,
	},
	"Core.OTP":       {API: "SYNO.Core.OTP", Version: 3},
	"Core.OTP.Admin": {API: "SYNO.Core.OTP.Admin", Version: 1},
	"Core.OTP.EnforcePolicy": {
		API:     "SYNO.Core.OTP.EnforcePolicy",
		Version: 1,
	},
	"Core.OTP.Ex":   {API: "SYNO.Core.OTP.Ex", Version: 1},
	"Core.OTP.Mail": {API: "SYNO.Core.OTP.Mail", Version: 1},
	"Core.Package":  {API: "SYNO.Core.Package", Version: 2},
	"Core.Package.AutoUpgrade.Progress": {
		API:     "SYNO.Core.Package.AutoUpgrade.Progress",
		Version: 1,
	},
	"Core.Package.Control": {
		API:     "SYNO.Core.Package.Control",
		Version: 1,
	},
	"Core.Package.FakeIFrame": {
		API:     "SYNO.Core.Package.FakeIFrame",
		Version: 1,
	},
	"Core.Package.Feed": {
		API:     "SYNO.Core.Package.Feed",
		Version: 1,
	},
	"Core.Package.Info": {
		API:     "SYNO.Core.Package.Info",
		Version: 1,
	},
	"Core.Package.Installation": {
		API:     "SYNO.Core.Package.Installation",
		Version: 2,
	},
	"Core.Package.Installation.Download": {
		API:     "SYNO.Core.Package.Installation.Download",
		Version: 1,
	},
	"Core.Package.Legal.PreRelease": {
		API:     "SYNO.Core.Package.Legal.PreRelease",
		Version: 1,
	},
	"Core.Package.Log": {API: "SYNO.Core.Package.Log", Version: 1},
	"Core.Package.MyDS": {
		API:     "SYNO.Core.Package.MyDS",
		Version: 1,
	},
	"Core.Package.MyDS.Purchase": {
		API:     "SYNO.Core.Package.MyDS.Purchase",
		Version: 1,
	},
	"Core.Package.Progress": {
		API:     "SYNO.Core.Package.Progress",
		Version: 1,
	},
	"Core.Package.Screenshot": {
		API:     "SYNO.Core.Package.Screenshot",
		Version: 1,
	},
	"Core.Package.Screenshot.Server": {
		API:     "SYNO.Core.Package.Screenshot.Server",
		Version: 1,
	},
	"Core.Package.Server": {
		API:     "SYNO.Core.Package.Server",
		Version: 2,
	},
	"Core.Package.Setting": {
		API:     "SYNO.Core.Package.Setting",
		Version: 1,
	},
	"Core.Package.Setting.Update": {
		API:     "SYNO.Core.Package.Setting.Update",
		Version: 1,
	},
	"Core.Package.Setting.Volume": {
		API:     "SYNO.Core.Package.Setting.Volume",
		Version: 1,
	},
	"Core.Package.Thumb": {
		API:     "SYNO.Core.Package.Thumb",
		Version: 1,
	},
	"Core.Package.Thumb.Server": {
		API:     "SYNO.Core.Package.Thumb.Server",
		Version: 1,
	},
	"Core.Package.Uninstallation": {
		API:     "SYNO.Core.Package.Uninstallation",
		Version: 1,
	},
	"Core.PersonalNotification.Device": {
		API:     "SYNO.Core.PersonalNotification.Device",
		Version: 1,
	},
	"Core.PersonalNotification.Event": {
		API:     "SYNO.Core.PersonalNotification.Event",
		Version: 1,
	},
	"Core.PersonalNotification.Filter": {
		API:     "SYNO.Core.PersonalNotification.Filter",
		Version: 1,
	},
	"Core.PersonalNotification.Mobile": {
		API:     "SYNO.Core.PersonalNotification.Mobile",
		Version: 1,
	},
	"Core.PersonalNotification.Settings": {
		API:     "SYNO.Core.PersonalNotification.Settings",
		Version: 1,
	},
	"Core.PersonalSettings": {
		API:     "SYNO.Core.PersonalSettings",
		Version: 1,
	},
	"Core.PhotoViewer": {API: "SYNO.Core.PhotoViewer", Version: 1},
	"Core.Polling.Data": {
		API:     "SYNO.Core.Polling.Data",
		Version: 1,
	},
	"Core.PortForwarding": {
		API:     "SYNO.Core.PortForwarding",
		Version: 1,
	},
	"Core.PortForwarding.Compatibility": {
		API:     "SYNO.Core.PortForwarding.Compatibility",
		Version: 1,
	},
	"Core.PortForwarding.RouterConf": {
		API:     "SYNO.Core.PortForwarding.RouterConf",
		Version: 1,
	},
	"Core.PortForwarding.RouterInfo": {
		API:     "SYNO.Core.PortForwarding.RouterInfo",
		Version: 1,
	},
	"Core.PortForwarding.RouterList": {
		API:     "SYNO.Core.PortForwarding.RouterList",
		Version: 1,
	},
	"Core.PortForwarding.Rules": {
		API:     "SYNO.Core.PortForwarding.Rules",
		Version: 1,
	},
	"Core.PortForwarding.Rules.Serv": {
		API:     "SYNO.Core.PortForwarding.Rules.Serv",
		Version: 1,
	},
	"Core.Promotion.Info": {
		API:     "SYNO.Core.Promotion.Info",
		Version: 1,
	},
	"Core.Promotion.PreInstall": {
		API:     "SYNO.Core.Promotion.PreInstall",
		Version: 1,
	},
	"Core.QuickConnect": {
		API:     "SYNO.Core.QuickConnect",
		Version: 3,
	},
	"Core.QuickConnect.Hostname": {
		API:     "SYNO.Core.QuickConnect.Hostname",
		Version: 1,
	},
	"Core.QuickConnect.Permission": {
		API:     "SYNO.Core.QuickConnect.Permission",
		Version: 1,
	},
	"Core.QuickConnect.RegisterSite": {
		API:     "SYNO.Core.QuickConnect.RegisterSite",
		Version: 1,
	},
	"Core.QuickConnect.SNI": {
		API:     "SYNO.Core.QuickConnect.SNI",
		Version: 1,
	},
	"Core.QuickConnect.Upnp": {
		API:     "SYNO.Core.QuickConnect.Upnp",
		Version: 1,
	},
	"Core.QuickStart.Info": {
		API:     "SYNO.Core.QuickStart.Info",
		Version: 3,
	},
	"Core.QuickStart.Install": {
		API:     "SYNO.Core.QuickStart.Install",
		Version: 1,
	},
	"Core.Quota":      {API: "SYNO.Core.Quota", Version: 1},
	"Core.RecycleBin": {API: "SYNO.Core.RecycleBin", Version: 1},
	"Core.RecycleBin.User": {
		API:     "SYNO.Core.RecycleBin.User",
		Version: 1,
	},
	"Core.Region.Language": {
		API:     "SYNO.Core.Region.Language",
		Version: 1,
	},
	"Core.Region.NTP": {API: "SYNO.Core.Region.NTP", Version: 3},
	"Core.Region.NTP.DateTimeFormat": {
		API:     "SYNO.Core.Region.NTP.DateTimeFormat",
		Version: 1,
	},
	"Core.Region.NTP.Server": {
		API:     "SYNO.Core.Region.NTP.Server",
		Version: 1,
	},
	"Core.ResetAdmin": {API: "SYNO.Core.ResetAdmin", Version: 1},
	"Core.SNMP":       {API: "SYNO.Core.SNMP", Version: 1},
	"Core.Security.AutoBlock": {
		API:     "SYNO.Core.Security.AutoBlock",
		Version: 1,
	},
	"Core.Security.AutoBlock.Rules": {
		API:     "SYNO.Core.Security.AutoBlock.Rules",
		Version: 1,
	},
	"Core.Security.DSM": {
		API:     "SYNO.Core.Security.DSM",
		Version: 5,
	},
	"Core.Security.DSM.Embed": {
		API:     "SYNO.Core.Security.DSM.Embed",
		Version: 1,
	},
	"Core.Security.DSM.Proxy": {
		API:     "SYNO.Core.Security.DSM.Proxy",
		Version: 1,
	},
	"Core.Security.DoS": {
		API:     "SYNO.Core.Security.DoS",
		Version: 2,
	},
	"Core.Security.Firewall": {
		API:     "SYNO.Core.Security.Firewall",
		Version: 1,
	},
	"Core.Security.Firewall.Adapter": {
		API:     "SYNO.Core.Security.Firewall.Adapter",
		Version: 1,
	},
	"Core.Security.Firewall.Conf": {
		API:     "SYNO.Core.Security.Firewall.Conf",
		Version: 1,
	},
	"Core.Security.Firewall.Geoip": {
		API:     "SYNO.Core.Security.Firewall.Geoip",
		Version: 1,
	},
	"Core.Security.Firewall.Profile": {
		API:     "SYNO.Core.Security.Firewall.Profile",
		Version: 1,
	},
	"Core.Security.Firewall.Profile.Apply": {
		API:     "SYNO.Core.Security.Firewall.Profile.Apply",
		Version: 1,
	},
	"Core.Security.Firewall.Rules": {
		API:     "SYNO.Core.Security.Firewall.Rules",
		Version: 1,
	},
	"Core.Security.Firewall.Rules.Serv": {
		API:     "SYNO.Core.Security.Firewall.Rules.Serv",
		Version: 1,
	},
	"Core.SecurityScan.Conf": {
		API:     "SYNO.Core.SecurityScan.Conf",
		Version: 1,
	},
	"Core.SecurityScan.Operation": {
		API:     "SYNO.Core.SecurityScan.Operation",
		Version: 1,
	},
	"Core.SecurityScan.Status": {
		API:     "SYNO.Core.SecurityScan.Status",
		Version: 1,
	},
	"Core.Service": {API: "SYNO.Core.Service", Version: 3},
	"Core.Service.Conf": {
		API:     "SYNO.Core.Service.Conf",
		Version: 1,
	},
	"Core.Service.PortInfo": {
		API:     "SYNO.Core.Service.PortInfo",
		Version: 1,
	},
	"Core.Share": {API: "SYNO.Core.Share", Version: 1},
	"Core.Share.Crypto": {
		API:     "SYNO.Core.Share.Crypto",
		Version: 1,
	},
	"Core.Share.Crypto.Key": {
		API:     "SYNO.Core.Share.Crypto.Key",
		Version: 1,
	},
	"Core.Share.CryptoFile": {
		API:     "SYNO.Core.Share.CryptoFile",
		Version: 1,
	},
	"Core.Share.KeyManager.AutoKey": {
		API:     "SYNO.Core.Share.KeyManager.AutoKey",
		Version: 1,
	},
	"Core.Share.KeyManager.Key": {
		API:     "SYNO.Core.Share.KeyManager.Key",
		Version: 1,
	},
	"Core.Share.KeyManager.MachineKey": {
		API:     "SYNO.Core.Share.KeyManager.MachineKey",
		Version: 1,
	},
	"Core.Share.KeyManager.Store": {
		API:     "SYNO.Core.Share.KeyManager.Store",
		Version: 2,
	},
	"Core.Share.Migration": {
		API:     "SYNO.Core.Share.Migration",
		Version: 1,
	},
	"Core.Share.Migration.Task": {
		API:     "SYNO.Core.Share.Migration.Task",
		Version: 1,
	},
	"Core.Share.Permission": {
		API:     "SYNO.Core.Share.Permission",
		Version: 1,
	},
	"Core.Share.PermissionReport": {
		API:     "SYNO.Core.Share.PermissionReport",
		Version: 1,
	},
	"Core.Share.Snapshot": {
		API:     "SYNO.Core.Share.Snapshot",
		Version: 2,
	},
	"Core.Sharing": {API: "SYNO.Core.Sharing", Version: 1},
	"Core.Sharing.Initdata": {
		API:     "SYNO.Core.Sharing.Initdata",
		Version: 1,
	},
	"Core.Sharing.Login": {
		API:     "SYNO.Core.Sharing.Login",
		Version: 1,
	},
	"Core.Sharing.Session": {
		API:     "SYNO.Core.Sharing.Session",
		Version: 1,
	},
	"Core.SmartBlock": {API: "SYNO.Core.SmartBlock", Version: 1},
	"Core.SmartBlock.Device": {
		API:     "SYNO.Core.SmartBlock.Device",
		Version: 1,
	},
	"Core.SmartBlock.Trusted": {
		API:     "SYNO.Core.SmartBlock.Trusted",
		Version: 1,
	},
	"Core.SmartBlock.Untrusted": {
		API:     "SYNO.Core.SmartBlock.Untrusted",
		Version: 1,
	},
	"Core.SmartBlock.User": {
		API:     "SYNO.Core.SmartBlock.User",
		Version: 1,
	},
	"Core.Storage.Disk": {
		API:     "SYNO.Core.Storage.Disk",
		Version: 1,
	},
	"Core.Storage.Disk.FWUpgrade": {
		API:     "SYNO.Core.Storage.Disk.FWUpgrade",
		Version: 1,
	},
	"Core.Storage.Pool": {
		API:     "SYNO.Core.Storage.Pool",
		Version: 1,
	},
	"Core.Storage.Volume": {
		API:     "SYNO.Core.Storage.Volume",
		Version: 1,
	},
	"Core.Storage.iSCSILUN": {
		API:     "SYNO.Core.Storage.iSCSILUN",
		Version: 1,
	},
	"Core.SupportForm.Form": {
		API:     "SYNO.Core.SupportForm.Form",
		Version: 1,
	},
	"Core.SupportForm.Log": {
		API:     "SYNO.Core.SupportForm.Log",
		Version: 1,
	},
	"Core.SupportForm.Service": {
		API:     "SYNO.Core.SupportForm.Service",
		Version: 1,
	},
	"Core.Synohdpack": {API: "SYNO.Core.Synohdpack", Version: 1},
	"Core.SyslogClient.FileTransfer": {
		API:     "SYNO.Core.SyslogClient.FileTransfer",
		Version: 1,
	},
	"Core.SyslogClient.Log": {
		API:     "SYNO.Core.SyslogClient.Log",
		Version: 1,
	},
	"Core.SyslogClient.PersonalActivity": {
		API:     "SYNO.Core.SyslogClient.PersonalActivity",
		Version: 1,
	},
	"Core.SyslogClient.Setting.Notify": {
		API:     "SYNO.Core.SyslogClient.Setting.Notify",
		Version: 1,
	},
	"Core.SyslogClient.Status": {
		API:     "SYNO.Core.SyslogClient.Status",
		Version: 1,
	},
	"Core.System": {API: "SYNO.Core.System", Version: 3},
	"Core.System.Process": {
		API:     "SYNO.Core.System.Process",
		Version: 1,
	},
	"Core.System.ProcessGroup": {
		API:     "SYNO.Core.System.ProcessGroup",
		Version: 1,
	},
	"Core.System.ResetButton": {
		API:     "SYNO.Core.System.ResetButton",
		Version: 1,
	},
	"Core.System.Status": {
		API:     "SYNO.Core.System.Status",
		Version: 1,
	},
	"Core.System.SystemHealth": {
		API:     "SYNO.Core.System.SystemHealth",
		Version: 1,
	},
	"Core.System.Utilization": {
		API:     "SYNO.Core.System.Utilization",
		Version: 1,
	},
	"Core.TFTP": {API: "SYNO.Core.TFTP", Version: 1},
	"Core.TaskScheduler": {
		API:     "SYNO.Core.TaskScheduler",
		Version: 4,
	},
	"Core.TaskScheduler.Root": {
		API:     "SYNO.Core.TaskScheduler.Root",
		Version: 4,
	},
	"Core.Terminal": {API: "SYNO.Core.Terminal", Version: 3},
	"Core.Theme.AppPortalLogin": {
		API:     "SYNO.Core.Theme.AppPortalLogin",
		Version: 1,
	},
	"Core.Theme.Desktop": {
		API:     "SYNO.Core.Theme.Desktop",
		Version: 1,
	},
	"Core.Theme.FileSharingLogin": {
		API:     "SYNO.Core.Theme.FileSharingLogin",
		Version: 1,
	},
	"Core.Theme.Image": {API: "SYNO.Core.Theme.Image", Version: 1},
	"Core.Theme.Login": {API: "SYNO.Core.Theme.Login", Version: 1},
	"Core.TrustDevice": {API: "SYNO.Core.TrustDevice", Version: 1},
	"Core.Tuned":       {API: "SYNO.Core.Tuned", Version: 1},
	"Core.UISearch":    {API: "SYNO.Core.UISearch", Version: 1},
	"Core.Upgrade":     {API: "SYNO.Core.Upgrade", Version: 2},
	"Core.Upgrade.AutoUpgrade": {
		API:     "SYNO.Core.Upgrade.AutoUpgrade",
		Version: 1,
	},
	"Core.Upgrade.Cluster.Patch": {
		API:     "SYNO.Core.Upgrade.Cluster.Patch",
		Version: 1,
	},
	"Core.Upgrade.Cluster.Server": {
		API:     "SYNO.Core.Upgrade.Cluster.Server",
		Version: 1,
	},
	"Core.Upgrade.Cluster.Server.Download": {
		API:     "SYNO.Core.Upgrade.Cluster.Server.Download",
		Version: 1,
	},
	"Core.Upgrade.Group": {
		API:     "SYNO.Core.Upgrade.Group",
		Version: 1,
	},
	"Core.Upgrade.Group.Download": {
		API:     "SYNO.Core.Upgrade.Group.Download",
		Version: 1,
	},
	"Core.Upgrade.Group.Setting": {
		API:     "SYNO.Core.Upgrade.Group.Setting",
		Version: 1,
	},
	"Core.Upgrade.GroupInstall": {
		API:     "SYNO.Core.Upgrade.GroupInstall",
		Version: 2,
	},
	"Core.Upgrade.GroupInstall.Network": {
		API:     "SYNO.Core.Upgrade.GroupInstall.Network",
		Version: 1,
	},
	"Core.Upgrade.JuniorModeData": {
		API:     "SYNO.Core.Upgrade.JuniorModeData",
		Version: 1,
	},
	"Core.Upgrade.Patch": {
		API:     "SYNO.Core.Upgrade.Patch",
		Version: 2,
	},
	"Core.Upgrade.PreCheck": {
		API:     "SYNO.Core.Upgrade.PreCheck",
		Version: 2,
	},
	"Core.Upgrade.RemoteAction": {
		API:     "SYNO.Core.Upgrade.RemoteAction",
		Version: 1,
	},
	"Core.Upgrade.Server": {
		API:     "SYNO.Core.Upgrade.Server",
		Version: 4,
	},
	"Core.Upgrade.Server.Download": {
		API:     "SYNO.Core.Upgrade.Server.Download",
		Version: 2,
	},
	"Core.Upgrade.Setting": {
		API:     "SYNO.Core.Upgrade.Setting",
		Version: 4,
	},
	"Core.User":       {API: "SYNO.Core.User", Version: 1},
	"Core.User.Group": {API: "SYNO.Core.User.Group", Version: 1},
	"Core.User.Home":  {API: "SYNO.Core.User.Home", Version: 1},
	"Core.User.PasswordConfirm": {
		API:     "SYNO.Core.User.PasswordConfirm",
		Version: 2,
	},
	"Core.User.PasswordExpiry": {
		API:     "SYNO.Core.User.PasswordExpiry",
		Version: 1,
	},
	"Core.User.PasswordMeter": {
		API:     "SYNO.Core.User.PasswordMeter",
		Version: 1,
	},
	"Core.User.PasswordPolicy": {
		API:     "SYNO.Core.User.PasswordPolicy",
		Version: 1,
	},
	"Core.User.UsernamePolicy": {
		API:     "SYNO.Core.User.UsernamePolicy",
		Version: 1,
	},
	"Core.UserSettings": {
		API:     "SYNO.Core.UserSettings",
		Version: 1,
	},
	"Core.Virtualization.Host.Capability": {
		API:     "SYNO.Core.Virtualization.Host.Capability",
		Version: 1,
	},
	"Core.Web.DSM": {API: "SYNO.Core.Web.DSM", Version: 2},
	"Core.Web.DSM.External": {
		API:     "SYNO.Core.Web.DSM.External",
		Version: 1,
	},
	"Core.Web.Security.HTTPCompression": {
		API:     "SYNO.Core.Web.Security.HTTPCompression",
		Version: 1,
	},
	"Core.Web.Security.TLSProfile": {
		API:     "SYNO.Core.Web.Security.TLSProfile",
		Version: 1,
	},
	"DR.Node": {API: "SYNO.DR.Node", Version: 1},
	"DR.Node.Credential": {
		API:     "SYNO.DR.Node.Credential",
		Version: 1,
	},
	"DR.Node.Session": {API: "SYNO.DR.Node.Session", Version: 2},
	"DSM.FindMe":      {API: "SYNO.DSM.FindMe", Version: 2},
	"DSM.Info":        {API: "SYNO.DSM.Info", Version: 2},
	"DSM.Network":     {API: "SYNO.DSM.Network", Version: 2},
	"DSM.PortEnable":  {API: "SYNO.DSM.PortEnable", Version: 1},
	"DSM.PushNotification": {
		API:     "SYNO.DSM.PushNotification",
		Version: 2,
	},
	"Default.API": {API: "SYNO.Default.API", Version: 1},
	"DisasterRecovery.Log": {
		API:     "SYNO.DisasterRecovery.Log",
		Version: 1,
	},
	"DisasterRecovery.Retention": {
		API:     "SYNO.DisasterRecovery.Retention",
		Version: 1,
	},
	"Entry.Request": {API: "SYNO.Entry.Request", Version: 2},
	"Entry.Request.Polling": {
		API:     "SYNO.Entry.Request.Polling",
		Version: 1,
	},
	"Entry.SocketIo": {API: "SYNO.Entry.SocketIo", Version: 1},
	"FileStation.BackgroundTask": {
		API:     "SYNO.FileStation.BackgroundTask",
		Version: 3,
	},
	"FileStation.CheckExist": {
		API:     "SYNO.FileStation.CheckExist",
		Version: 2,
	},
	"FileStation.CheckPermission": {
		API:     "SYNO.FileStation.CheckPermission",
		Version: 3,
	},
	"FileStation.Compress": {
		API:     "SYNO.FileStation.Compress",
		Version: 3,
	},
	"FileStation.CopyMove": {
		API:     "SYNO.FileStation.CopyMove",
		Version: 3,
	},
	"FileStation.CreateFolder": {
		API:     "SYNO.FileStation.CreateFolder",
		Version: 2,
	},
	"FileStation.Delete": {
		API:     "SYNO.FileStation.Delete",
		Version: 2,
	},
	"FileStation.DirSize": {
		API:     "SYNO.FileStation.DirSize",
		Version: 2,
	},
	"FileStation.Download": {
		API:     "SYNO.FileStation.Download",
		Version: 2,
	},
	"FileStation.External.GoogleDrive": {
		API:     "SYNO.FileStation.External.GoogleDrive",
		Version: 2,
	},
	"FileStation.Extract": {
		API:     "SYNO.FileStation.Extract",
		Version: 2,
	},
	"FileStation.Favorite": {
		API:     "SYNO.FileStation.Favorite",
		Version: 2,
	},
	"FileStation.FormUpload": {
		API:     "SYNO.FileStation.FormUpload",
		Version: 2,
	},
	"FileStation.Info": {API: "SYNO.FileStation.Info", Version: 2},
	"FileStation.List": {API: "SYNO.FileStation.List", Version: 2},
	"FileStation.MD5":  {API: "SYNO.FileStation.MD5", Version: 2},
	"FileStation.Mount": {
		API:     "SYNO.FileStation.Mount",
		Version: 1,
	},
	"FileStation.Mount.List": {
		API:     "SYNO.FileStation.Mount.List",
		Version: 1,
	},
	"FileStation.Notify": {
		API:     "SYNO.FileStation.Notify",
		Version: 1,
	},
	"FileStation.PhotoUpload": {
		API:     "SYNO.FileStation.PhotoUpload",
		Version: 3,
	},
	"FileStation.Property": {
		API:     "SYNO.FileStation.Property",
		Version: 1,
	},
	"FileStation.Property.ACLOwner": {
		API:     "SYNO.FileStation.Property.ACLOwner",
		Version: 1,
	},
	"FileStation.Property.CompressSize": {
		API:     "SYNO.FileStation.Property.CompressSize",
		Version: 1,
	},
	"FileStation.Property.Mtime": {
		API:     "SYNO.FileStation.Property.Mtime",
		Version: 1,
	},
	"FileStation.Rename": {
		API:     "SYNO.FileStation.Rename",
		Version: 2,
	},
	"FileStation.Search": {
		API:     "SYNO.FileStation.Search",
		Version: 2,
	},
	"FileStation.Search.History": {
		API:     "SYNO.FileStation.Search.History",
		Version: 1,
	},
	"FileStation.Settings": {
		API:     "SYNO.FileStation.Settings",
		Version: 1,
	},
	"FileStation.Sharing": {
		API:     "SYNO.FileStation.Sharing",
		Version: 3,
	},
	"FileStation.Sharing.Download": {
		API:     "SYNO.FileStation.Sharing.Download",
		Version: 1,
	},
	"FileStation.Snapshot": {
		API:     "SYNO.FileStation.Snapshot",
		Version: 2,
	},
	"FileStation.Thumb": {
		API:     "SYNO.FileStation.Thumb",
		Version: 3,
	},
	"FileStation.Timeout": {
		API:     "SYNO.FileStation.Timeout",
		Version: 1,
	},
	"FileStation.UIString": {
		API:     "SYNO.FileStation.UIString",
		Version: 1,
	},
	"FileStation.Upload": {
		API:     "SYNO.FileStation.Upload",
		Version: 3,
	},
	"FileStation.UserGrp": {
		API:     "SYNO.FileStation.UserGrp",
		Version: 1,
	},
	"FileStation.VFS.Connection": {
		API:     "SYNO.FileStation.VFS.Connection",
		Version: 1,
	},
	"FileStation.VFS.File": {
		API:     "SYNO.FileStation.VFS.File",
		Version: 1,
	},
	"FileStation.VFS.GDrive": {
		API:     "SYNO.FileStation.VFS.GDrive",
		Version: 1,
	},
	"FileStation.VFS.Profile": {
		API:     "SYNO.FileStation.VFS.Profile",
		Version: 1,
	},
	"FileStation.VFS.Protocol": {
		API:     "SYNO.FileStation.VFS.Protocol",
		Version: 1,
	},
	"FileStation.VFS.User": {
		API:     "SYNO.FileStation.VFS.User",
		Version: 1,
	},
	"FileStation.VirtualFolder": {
		API:     "SYNO.FileStation.VirtualFolder",
		Version: 2,
	},
	"FileStation.Worm": {API: "SYNO.FileStation.Worm", Version: 2},
	"FileStation.Worm.Lock": {
		API:     "SYNO.FileStation.Worm.Lock",
		Version: 2,
	},
	"Finder.AppIndexing.Search": {
		API:     "SYNO.Finder.AppIndexing.Search",
		Version: 1,
	},
	"Finder.Bookmark": {API: "SYNO.Finder.Bookmark", Version: 1},
	"Finder.Elastic.SearchHistory": {
		API:     "SYNO.Finder.Elastic.SearchHistory",
		Version: 1,
	},
	"Finder.Elastic.Spotlight": {
		API:     "SYNO.Finder.Elastic.Spotlight",
		Version: 1,
	},
	"Finder.Elastic.Term": {
		API:     "SYNO.Finder.Elastic.Term",
		Version: 1,
	},
	"Finder.File": {API: "SYNO.Finder.File", Version: 1},
	"Finder.File.Cover": {
		API:     "SYNO.Finder.File.Cover",
		Version: 1,
	},
	"Finder.File.Thumbnail": {
		API:     "SYNO.Finder.File.Thumbnail",
		Version: 1,
	},
	"Finder.FileIndexing.Folder": {
		API:     "SYNO.Finder.FileIndexing.Folder",
		Version: 1,
	},
	"Finder.FileIndexing.Highlight": {
		API:     "SYNO.Finder.FileIndexing.Highlight",
		Version: 1,
	},
	"Finder.FileIndexing.Indicate": {
		API:     "SYNO.Finder.FileIndexing.Indicate",
		Version: 1,
	},
	"Finder.FileIndexing.Search": {
		API:     "SYNO.Finder.FileIndexing.Search",
		Version: 1,
	},
	"Finder.FileIndexing.Status": {
		API:     "SYNO.Finder.FileIndexing.Status",
		Version: 1,
	},
	"Finder.FileIndexing.Term": {
		API:     "SYNO.Finder.FileIndexing.Term",
		Version: 1,
	},
	"Finder.Preference": {
		API:     "SYNO.Finder.Preference",
		Version: 1,
	},
	"Finder.Settings": {API: "SYNO.Finder.Settings", Version: 1},
	"Finder.UserGrp":  {API: "SYNO.Finder.UserGrp", Version: 1},
	"FolderSharing.Download": {
		API:     "SYNO.FolderSharing.Download",
		Version: 2,
	},
	"FolderSharing.List": {
		API:     "SYNO.FolderSharing.List",
		Version: 2,
	},
	"FolderSharing.Thumb": {
		API:     "SYNO.FolderSharing.Thumb",
		Version: 2,
	},
	"HCI.Cluster":  {API: "SYNO.HCI.Cluster", Version: 1},
	"HCI.Guest":    {API: "SYNO.HCI.Guest", Version: 1},
	"HCI.Host":     {API: "SYNO.HCI.Host", Version: 1},
	"HCI.Utils":    {API: "SYNO.HCI.Utils", Version: 1},
	"License.HA":   {API: "SYNO.License.HA", Version: 1},
	"Lunbackup":    {API: "SYNO.Lunbackup", Version: 1},
	"OAUTH.Client": {API: "SYNO.OAUTH.Client", Version: 1},
	"OAUTH.Common": {API: "SYNO.OAUTH.Common", Version: 1},
	"OAUTH.Log":    {API: "SYNO.OAUTH.Log", Version: 1},
	"OAUTH.Token":  {API: "SYNO.OAUTH.Token", Version: 1},
	"Package":      {API: "SYNO.Package", Version: 1},
	"PersonMailAccount": {
		API:     "SYNO.PersonMailAccount",
		Version: 1,
	},
	"PersonMailAccount.Contacts": {
		API:     "SYNO.PersonMailAccount.Contacts",
		Version: 1,
	},
	"PersonMailAccount.Mail": {
		API:     "SYNO.PersonMailAccount.Mail",
		Version: 1,
	},
	"PersonMailAccount.Mail.Oauth": {
		API:     "SYNO.PersonMailAccount.Mail.Oauth",
		Version: 1,
	},
	"Personal.Application.Info": {
		API:     "SYNO.Personal.Application.Info",
		Version: 1,
	},
	"Personal.MailAccount": {
		API:     "SYNO.Personal.MailAccount",
		Version: 1,
	},
	"Personal.MailAccount.Contacts": {
		API:     "SYNO.Personal.MailAccount.Contacts",
		Version: 1,
	},
	"Personal.MailAccount.Mail": {
		API:     "SYNO.Personal.MailAccount.Mail",
		Version: 1,
	},
	"Personal.Notification.Conf": {
		API:     "SYNO.Personal.Notification.Conf",
		Version: 2,
	},
	"Personal.Notification.Device": {
		API:     "SYNO.Personal.Notification.Device",
		Version: 2,
	},
	"Personal.Notification.Event": {
		API:     "SYNO.Personal.Notification.Event",
		Version: 2,
	},
	"Personal.Notification.Filter": {
		API:     "SYNO.Personal.Notification.Filter",
		Version: 1,
	},
	"Personal.Notification.GDPR": {
		API:     "SYNO.Personal.Notification.GDPR",
		Version: 1,
	},
	"Personal.Notification.Identifier": {
		API:     "SYNO.Personal.Notification.Identifier",
		Version: 1,
	},
	"Personal.Notification.Mobile": {
		API:     "SYNO.Personal.Notification.Mobile",
		Version: 3,
	},
	"Personal.Notification.Settings": {
		API:     "SYNO.Personal.Notification.Settings",
		Version: 2,
	},
	"Personal.Notification.Token": {
		API:     "SYNO.Personal.Notification.Token",
		Version: 1,
	},
	"Personal.Notification.VapidPublicKey": {
		API:     "SYNO.Personal.Notification.VapidPublicKey",
		Version: 1,
	},
	"Personal.Profile": {API: "SYNO.Personal.Profile", Version: 2},
	"Personal.Profile.Photo": {
		API:     "SYNO.Personal.Profile.Photo",
		Version: 1,
	},
	"Remote.Credential": {
		API:     "SYNO.Remote.Credential",
		Version: 1,
	},
	"Remote.Credential.Challenge": {
		API:     "SYNO.Remote.Credential.Challenge",
		Version: 1,
	},
	"Remote.Credential.Info": {
		API:     "SYNO.Remote.Credential.Info",
		Version: 1,
	},
	"Remote.Credential.Verifier": {
		API:     "SYNO.Remote.Credential.Verifier",
		Version: 1,
	},
	"ResourceMonitor.EventRule": {
		API:     "SYNO.ResourceMonitor.EventRule",
		Version: 1,
	},
	"ResourceMonitor.Log": {
		API:     "SYNO.ResourceMonitor.Log",
		Version: 1,
	},
	"ResourceMonitor.Setting": {
		API:     "SYNO.ResourceMonitor.Setting",
		Version: 1,
	},
	"S2S.Client":      {API: "SYNO.S2S.Client", Version: 1},
	"S2S.Client.Job":  {API: "SYNO.S2S.Client.Job", Version: 1},
	"S2S.Server":      {API: "SYNO.S2S.Server", Version: 1},
	"S2S.Server.Pair": {API: "SYNO.S2S.Server.Pair", Version: 1},
	"SAS.APIRunner":   {API: "SYNO.SAS.APIRunner", Version: 1},
	"SAS.APIRunner.Chatbot": {
		API:     "SYNO.SAS.APIRunner.Chatbot",
		Version: 1,
	},
	"SAS.Encryption": {API: "SYNO.SAS.Encryption", Version: 1},
	"SAS.Group":      {API: "SYNO.SAS.Group", Version: 1},
	"SAS.Group.Members": {
		API:     "SYNO.SAS.Group.Members",
		Version: 1,
	},
	"SAS.Guest": {API: "SYNO.SAS.Guest", Version: 1},
	"SecureSignIn.AMFA.Mail.Ex": {
		API:     "SYNO.SecureSignIn.AMFA.Mail.Ex",
		Version: 1,
	},
	"SecureSignIn.AMFA.Policy": {
		API:     "SYNO.SecureSignIn.AMFA.Policy",
		Version: 1,
	},
	"SecureSignIn.AMFA.SuggestConn": {
		API:     "SYNO.SecureSignIn.AMFA.SuggestConn",
		Version: 1,
	},
	"SecureSignIn.Authenticator": {
		API:     "SYNO.SecureSignIn.Authenticator",
		Version: 1,
	},
	"SecureSignIn.Authenticator.Ex": {
		API:     "SYNO.SecureSignIn.Authenticator.Ex",
		Version: 1,
	},
	"SecureSignIn.Authenticator.Info": {
		API:     "SYNO.SecureSignIn.Authenticator.Info",
		Version: 2,
	},
	"SecureSignIn.Authenticator.Registration": {
		API:     "SYNO.SecureSignIn.Authenticator.Registration",
		Version: 2,
	},
	"SecureSignIn.Authenticator.Registration.Ex": {
		API:     "SYNO.SecureSignIn.Authenticator.Registration.Ex",
		Version: 2,
	},
	"SecureSignIn.Authenticator.Request": {
		API:     "SYNO.SecureSignIn.Authenticator.Request",
		Version: 1,
	},
	"SecureSignIn.Authenticator.Verdict": {
		API:     "SYNO.SecureSignIn.Authenticator.Verdict",
		Version: 2,
	},
	"SecureSignIn.Fido.Manage": {
		API:     "SYNO.SecureSignIn.Fido.Manage",
		Version: 1,
	},
	"SecureSignIn.Fido.Manage.Ex": {
		API:     "SYNO.SecureSignIn.Fido.Manage.Ex",
		Version: 1,
	},
	"SecureSignIn.Fido.Register": {
		API:     "SYNO.SecureSignIn.Fido.Register",
		Version: 1,
	},
	"SecureSignIn.Fido.Register.Ex": {
		API:     "SYNO.SecureSignIn.Fido.Register.Ex",
		Version: 1,
	},
	"SecureSignIn.Method": {
		API:     "SYNO.SecureSignIn.Method",
		Version: 1,
	},
	"SecureSignIn.Method.Admin": {
		API:     "SYNO.SecureSignIn.Method.Admin",
		Version: 1,
	},
	"SecureSignIn.Method.Ex": {
		API:     "SYNO.SecureSignIn.Method.Ex",
		Version: 1,
	},
	"SecureSignIn.Package": {
		API:     "SYNO.SecureSignIn.Package",
		Version: 1,
	},
	"SecureSignIn.Package.Ex": {
		API:     "SYNO.SecureSignIn.Package.Ex",
		Version: 1,
	},
	"SecureSignIn.Package.Request": {
		API:     "SYNO.SecureSignIn.Package.Request",
		Version: 1,
	},
	"SecurityAdvisor.Conf": {
		API:     "SYNO.SecurityAdvisor.Conf",
		Version: 1,
	},
	"SecurityAdvisor.Conf.Checklist": {
		API:     "SYNO.SecurityAdvisor.Conf.Checklist",
		Version: 1,
	},
	"SecurityAdvisor.Conf.Checklist.Alert": {
		API:     "SYNO.SecurityAdvisor.Conf.Checklist.Alert",
		Version: 1,
	},
	"SecurityAdvisor.Conf.Location": {
		API:     "SYNO.SecurityAdvisor.Conf.Location",
		Version: 1,
	},
	"SecurityAdvisor.LoginActivity": {
		API:     "SYNO.SecurityAdvisor.LoginActivity",
		Version: 1,
	},
	"SecurityAdvisor.LoginActivity.User": {
		API:     "SYNO.SecurityAdvisor.LoginActivity.User",
		Version: 1,
	},
	"SecurityAdvisor.Report": {
		API:     "SYNO.SecurityAdvisor.Report",
		Version: 1,
	},
	"SecurityAdvisor.Report.HTML": {
		API:     "SYNO.SecurityAdvisor.Report.HTML",
		Version: 1,
	},
	"Snap.Usage.Share": {API: "SYNO.Snap.Usage.Share", Version: 1},
	"Storage.CGI.BtrfsDedupe": {
		API:     "SYNO.Storage.CGI.BtrfsDedupe",
		Version: 1,
	},
	"Storage.CGI.Cache.Protection": {
		API:     "SYNO.Storage.CGI.Cache.Protection",
		Version: 1,
	},
	"Storage.CGI.Check": {
		API:     "SYNO.Storage.CGI.Check",
		Version: 1,
	},
	"Storage.CGI.DetectedPool": {
		API:     "SYNO.Storage.CGI.DetectedPool",
		Version: 1,
	},
	"Storage.CGI.DualEnclosure": {
		API:     "SYNO.Storage.CGI.DualEnclosure",
		Version: 1,
	},
	"Storage.CGI.Enclosure": {
		API:     "SYNO.Storage.CGI.Enclosure",
		Version: 1,
	},
	"Storage.CGI.EncryptionKeyVault": {
		API:     "SYNO.Storage.CGI.EncryptionKeyVault",
		Version: 1,
	},
	"Storage.CGI.EncryptionKeyVault.UnlockMode": {
		API:     "SYNO.Storage.CGI.EncryptionKeyVault.UnlockMode",
		Version: 1,
	},
	"Storage.CGI.Flashcache": {
		API:     "SYNO.Storage.CGI.Flashcache",
		Version: 1,
	},
	"Storage.CGI.HddMan": {
		API:     "SYNO.Storage.CGI.HddMan",
		Version: 1,
	},
	"Storage.CGI.KMIP": {API: "SYNO.Storage.CGI.KMIP", Version: 1},
	"Storage.CGI.Pool": {API: "SYNO.Storage.CGI.Pool", Version: 1},
	"Storage.CGI.Scrubbing": {
		API:     "SYNO.Storage.CGI.Scrubbing",
		Version: 1,
	},
	"Storage.CGI.Smart": {
		API:     "SYNO.Storage.CGI.Smart",
		Version: 1,
	},
	"Storage.CGI.Smart.Scheduler": {
		API:     "SYNO.Storage.CGI.Smart.Scheduler",
		Version: 1,
	},
	"Storage.CGI.Spare": {
		API:     "SYNO.Storage.CGI.Spare",
		Version: 1,
	},
	"Storage.CGI.Spare.Conf": {
		API:     "SYNO.Storage.CGI.Spare.Conf",
		Version: 1,
	},
	"Storage.CGI.Storage": {
		API:     "SYNO.Storage.CGI.Storage",
		Version: 1,
	},
	"Storage.CGI.TaipeiEnclosure": {
		API:     "SYNO.Storage.CGI.TaipeiEnclosure",
		Version: 1,
	},
	"Storage.CGI.Volume": {
		API:     "SYNO.Storage.CGI.Volume",
		Version: 1,
	},
	"Storage.CGI.Volume.Installer": {
		API:     "SYNO.Storage.CGI.Volume.Installer",
		Version: 1,
	},
	"Storage.CGI.Volume.OfflineOp": {
		API:     "SYNO.Storage.CGI.Volume.OfflineOp",
		Version: 1,
	},
	"SupportService.Setting": {
		API:     "SYNO.SupportService.Setting",
		Version: 2,
	},
	"TextEditor": {API: "SYNO.TextEditor", Version: 1},
	"TextEditor.Preference": {
		API:     "SYNO.TextEditor.Preference",
		Version: 1,
	},
	"Utils": {API: "SYNO.Utils", Version: 1},
	"VMMDR.Btrfs.Replica": {
		API:     "SYNO.VMMDR.Btrfs.Replica",
		Version: 1,
	},
	"VMMDR.Credential": {API: "SYNO.VMMDR.Credential", Version: 1},
	"VMMDR.Plan":       {API: "SYNO.VMMDR.Plan", Version: 3},
	"VMMDR.Plan.DRSite": {
		API:     "SYNO.VMMDR.Plan.DRSite",
		Version: 1,
	},
	"VMMDR.Plan.MainSite": {
		API:     "SYNO.VMMDR.Plan.MainSite",
		Version: 1,
	},
	"VMMDR.Plan.Site": {API: "SYNO.VMMDR.Plan.Site", Version: 1},
	"VMMDR.Topology":  {API: "SYNO.VMMDR.Topology", Version: 1},
	"VideoPlayer.Subtitle": {
		API:     "SYNO.VideoPlayer.Subtitle",
		Version: 1,
	},
	"VideoPlayer.SynologyDrive.Subtitle": {
		API:     "SYNO.VideoPlayer.SynologyDrive.Subtitle",
		Version: 1,
	},
	"Virtualization.GetGuest": {
		API:     "SYNO.Virtualization.API.Guest",
		Version: 1,
		Method:  "get",
	},
	"Virtualization.API.Guest.Action": {
		API:     "SYNO.Virtualization.API.Guest.Action",
		Version: 1,
	},
	"Virtualization.API.Guest.Image": {
		API:     "SYNO.Virtualization.API.Guest.Image",
		Version: 1,
	},
	"Virtualization.API.Host": {
		API:     "SYNO.Virtualization.API.Host",
		Version: 1,
	},
	"Virtualization.API.Network": {
		API:     "SYNO.Virtualization.API.Network",
		Version: 1,
	},
	"Virtualization.API.Storage": {
		API:     "SYNO.Virtualization.API.Storage",
		Version: 1,
	},
	"Virtualization.API.Task.Info": {
		API:     "SYNO.Virtualization.API.Task.Info",
		Version: 1,
	},
	"Virtualization.Cluster": {
		API:     "SYNO.Virtualization.Cluster",
		Version: 2,
	},
	"Virtualization.Guest": {
		API:     "SYNO.Virtualization.Guest",
		Version: 2,
	},
	"Virtualization.Guest.Action": {
		API:     "SYNO.Virtualization.Guest.Action",
		Version: 1,
	},
	"Virtualization.Guest.Image": {
		API:     "SYNO.Virtualization.Guest.Image",
		Version: 2,
	},
	"Virtualization.Guest.P2V": {
		API:     "SYNO.Virtualization.Guest.P2V",
		Version: 1,
	},
	"Virtualization.GuestProtect.Plan": {
		API:     "SYNO.Virtualization.GuestProtect.Plan",
		Version: 2,
	},
	"Virtualization.GuestProtect.Policy": {
		API:     "SYNO.Virtualization.GuestProtect.Policy",
		Version: 1,
	},
	"Virtualization.GuestProtect.Snap": {
		API:     "SYNO.Virtualization.GuestProtect.Snap",
		Version: 1,
	},
	"Virtualization.HA": {
		API:     "SYNO.Virtualization.HA",
		Version: 1,
	},
	"Virtualization.HA.Setting": {
		API:     "SYNO.Virtualization.HA.Setting",
		Version: 1,
	},
	"Virtualization.Host": {
		API:     "SYNO.Virtualization.Host",
		Version: 2,
	},
	"Virtualization.License": {
		API:     "SYNO.Virtualization.License",
		Version: 1,
	},
	"Virtualization.License.Pro": {
		API:     "SYNO.Virtualization.License.Pro",
		Version: 1,
	},
	"Virtualization.License.VDSM": {
		API:     "SYNO.Virtualization.License.VDSM",
		Version: 1,
	},
	"Virtualization.Log": {
		API:     "SYNO.Virtualization.Log",
		Version: 1,
	},
	"Virtualization.Network": {
		API:     "SYNO.Virtualization.Network",
		Version: 2,
	},
	"Virtualization.Network.SRIOV": {
		API:     "SYNO.Virtualization.Network.SRIOV",
		Version: 1,
	},
	"Virtualization.Package": {
		API:     "SYNO.Virtualization.Package",
		Version: 1,
	},
	"Virtualization.Repo": {
		API:     "SYNO.Virtualization.Repo",
		Version: 2,
	},
	"Virtualization.Setting.General": {
		API:     "SYNO.Virtualization.Setting.General",
		Version: 1,
	},
	"Virtualization.Setting.Notify": {
		API:     "SYNO.Virtualization.Setting.Notify",
		Version: 1,
	},
	"Virtualization.Setting.UI": {
		API:     "SYNO.Virtualization.Setting.UI",
		Version: 1,
	},
	"Virtualization.Sharing.VNC": {
		API:     "SYNO.Virtualization.Sharing.VNC",
		Version: 1,
	},
	"Virtualization.User": {
		API:     "SYNO.Virtualization.User",
		Version: 1,
	},
	"Virtualization.Utils": {
		API:     "SYNO.Virtualization.Utils",
		Version: 1,
	},
	"WebRTC.Proxy": {API: "SYNO.WebRTC.Proxy", Version: 1},
	"WebRTC.Proxy.Sharing": {
		API:     "SYNO.WebRTC.Proxy.Sharing",
		Version: 1,
	},
	"WebStation.Backup": {
		API:     "SYNO.WebStation.Backup",
		Version: 1,
	},
	"WebStation.Default": {
		API:     "SYNO.WebStation.Default",
		Version: 1,
	},
	"WebStation.Docker": {
		API:     "SYNO.WebStation.Docker",
		Version: 1,
	},
	"WebStation.ErrorPage": {
		API:     "SYNO.WebStation.ErrorPage",
		Version: 1,
	},
	"WebStation.HTTP.VHost": {
		API:     "SYNO.WebStation.HTTP.VHost",
		Version: 1,
	},
	"WebStation.PHP": {API: "SYNO.WebStation.PHP", Version: 1},
	"WebStation.PHP.Profile": {
		API:     "SYNO.WebStation.PHP.Profile",
		Version: 1,
	},
	"WebStation.Package": {
		API:     "SYNO.WebStation.Package",
		Version: 1,
	},
	"WebStation.Python": {
		API:     "SYNO.WebStation.Python",
		Version: 1,
	},
	"WebStation.Python.Profile": {
		API:     "SYNO.WebStation.Python.Profile",
		Version: 1,
	},
	"WebStation.ScriptLanguage": {
		API:     "SYNO.WebStation.ScriptLanguage",
		Version: 1,
	},
	"WebStation.ScriptLanguage.Utils": {
		API:     "SYNO.WebStation.ScriptLanguage.Utils",
		Version: 1,
	},
	"WebStation.Shortcut": {
		API:     "SYNO.WebStation.Shortcut",
		Version: 1,
	},
	"WebStation.Status": {
		API:     "SYNO.WebStation.Status",
		Version: 1,
	},
	"WebStation.Task": {API: "SYNO.WebStation.Task", Version: 1},
	"WebStation.WebService.Portal": {
		API:     "SYNO.WebStation.WebService.Portal",
		Version: 1,
	},
	"WebStation.WebService.Portal.Log": {
		API:     "SYNO.WebStation.WebService.Portal.Log",
		Version: 1,
	},
	"WebStation.WebService.Service": {
		API:     "SYNO.WebStation.WebService.Service",
		Version: 1,
	},
}
