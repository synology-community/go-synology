package api

import (
	"encoding/json"
	"fmt"
	"maps"
	"strconv"

	"github.com/hashicorp/go-multierror"
)

type ErrorSummaries func() ErrorSummary

func (es ErrorSummaries) Combine(errs ErrorSummary) ErrorSummaries {
	return func() ErrorSummary {
		return es().Combine(errs)
	}
}

var GlobalErrors ErrorSummaries = func() ErrorSummary {
	return globalErrors
}

// GlobalErrors holds mapping of global errors not related to particular API endpoint.
var globalErrors ErrorSummary = ErrorSummary{
	0:    "common - commfail",
	101:  "No parameter of API, method or version",
	102:  "The requested API does not exist",
	103:  "The requested method does not exist",
	104:  "The requested version does not support the functionality",
	105:  "The logged in session does not have permission",
	106:  "Session timeout",
	107:  "Session interrupted by duplicate login",
	108:  "user - user_file_upload_fail",
	109:  "error - error_error_system",
	110:  "error - error_error_system",
	1101: "error - error_subject",
	1102: "firewall - firewall_restore_failed",
	1103: "firewall - firewall_block_admin_client",
	1104: "firewall - firewall_rule_exceed_max_number",
	1105: "firewall - firewall_rule_disable_fail",
	111:  "error - error_error_system",
	112:  "Stop Handling Compound Request",
	113:  "Invalid Compound Request",
	114:  "error - error_invalid",
	115:  "error - error_invalid",
	116:  "uicommon - error_demo",
	117:  "error - error_error_system",
	118:  "error - error_error_system",
	119:  "SID not found",
	1198: "common - version_not_support",
	1201: "error - error_subject",
	1202: "firewall - firewall_tc_ceil_exceed_system_upper_bound",
	1203: "firewall - firewall_tc_max_ceil_too_large",
	1204: "firewall - firewall_tc_restore_failed",
	122:  "error - error_privilege_not_enough",
	123:  "error - error_privilege_not_enough",
	124:  "error - error_privilege_not_enough",
	125:  "uicommon - error_timeout",
	126:  "error - error_privilege_not_enough",
	127:  "error - error_privilege_not_enough",
	1301: "error - error_subject",
	1302: "firewall - firewall_dos_restore_failed",
	1400: "Failed to extract files.",
	1401: "Cannot open the file as archive.",
	1402: "Failed to read archive data error",
	1403: "Wrong password.",
	1404: "Failed to get the file and dir list in an archive.",
	1405: "Failed to find the item ID in an archive file.",
	1410: "service - service_ddns_operation_fail",
	1500: "common - error_system",
	1501: "common - error_apply_occupied",
	1502: "routerconf - routerconf_external_ip_warning",
	1503: "routerconf - routerconf_require_gateway",
	1504: "routerconf - dns_setting_no_response",
	1510: "routerconf - routerconf_update_db_failed",
	1521: "routerconf - routerconf_exceed_singel_max_port",
	1522: "routerconf - routerconf_exceed_combo_max_port",
	1523: "routerconf - routerconf_exceed_singel_range_max_port",
	1524: "routerconf - routerconf_exceed_max_rule",
	1525: "routerconf - routerconf_port_conflict",
	1526: "routerconf - routerconf_add_port_failed",
	1527: "routerconf - routerconf_apply_failed",
	1528: "routerconf - protocol_on_router_not_enabled",
	1530: "routerconf - routerconf_syntax_version_error",
	160:  "error - error_privilege_not_enough",
	1600: "ups - operation_failed",
	1601: "ups - set_info_failed",
	1602: "ups - get_info_failed",
	164:  "error - error_load_system_settings",
	165:  "error - error_set_system_settings",
	1701: "error - error_port_conflict",
	1702: "error - error_file_exist",
	1703: "error - error_no_path",
	1704: "error - error_error_system",
	1706: "error - error_volume_ro",
	1800: "There is no Content-Length information in the HTTP header or the received size doesn't match the value of Content-Length information in the HTTP header.",
	1801: "Wait too long, no date can be received from client (Default maximum wait time is 3600 seconds).",
	1802: "No filename information in the last part of file content.",
	1803: "Upload connection is cancelled.",
	1804: "Failed to upload oversized file to FAT file system.",
	1805: "Can't overwrite or skip the existing file, if no   parameter is given.",
	1903: "error - error_port_conflict",
	1904: "error - error_port_conflict",
	1905: "ftp - ftp_anoymous_root_share_invalid",
	1951: "error - error_port_conflict",
	2001: "error - error_error_system",
	2002: "error - error_error_system",
	2101: "error - error_error_system",
	2102: "error - error_error_system",
	2201: "error - error_error_system",
	2202: "error - error_error_system",
	2301: "error - error_invalid",
	2303: "error - error_port_conflict",
	2331: "nfs - nfs_key_wrong_format",
	2332: "user - user_file_upload_fail",
	2371: "error - error_mount_point_nfs",
	2372: "error - error_hfs_plus_mount_point_nfs",
	2401: "error - error_error_system",
	2402: "error - error_error_system",
	2403: "error - error_port_conflict",
	2500: "error - error_unknown_desc",
	2502: "error - error_invalid",
	2503: "error - error_error_system",
	2504: "error - error_error_system",
	2505: "error - error_error_system",
	2601: "network - domain_name_err",
	2602: "network - domain_dns_name_err",
	2603: "network - domain_kdc_ip_error",
	2604: "network - error_badgname",
	2605: "network - domain_unreachserver_err",
	2606: "network - domain_port_unreachable_err",
	2607: "network - domain_password_err",
	2608: "network - domain_acc_revoked_ads",
	2609: "network - domain_acc_revoked_rpc",
	2610: "network - domain_acc_err",
	2611: "network - domain_notadminuser",
	2612: "network - domain_change_passwd",
	2613: "network - domain_check_kdcip",
	2614: "network - domain_error_misc_rpc",
	2615: "network - domain_join_err",
	2616: "directory_service - warr_enable_samba",
	2626: "directory_service - warr_db_not_ready",
	2628: "directory_service - warr_synoad_exists",
	2702: "network - status_connected",
	2703: "network - status_disconnected",
	2704: "common - error_occupied",
	2705: "common - error_system",
	2706: "ldap_error - ldap_invalid_credentials",
	2707: "ldap_error - ldap_operations_error",
	2708: "ldap_error - ldap_server_not_support",
	2709: "domain - domain_ldap_conflict",
	2710: "ldap_error - ldap_operations_error",
	2712: "ldap_error - ldap_no_such_object",
	2713: "ldap_error - ldap_protocol_error",
	2714: "ldap_error - ldap_invalid_dn_syntax",
	2715: "ldap_error - ldap_insufficient_access",
	2716: "ldap_error - ldap_insufficient_access",
	2717: "ldap_error - ldap_timelimit_exceeded",
	2718: "ldap_error - ldap_inappropriate_auth",
	2719: "ldap_error - ldap_smb2_enable_warning",
	2721: "ldap_error - ldap_confidentiality_required",
	2723: "ldap_error - ldap_weak_pwd",
	2799: "common - error_system",
	2800: "error - error_unknown_desc",
	2801: "error - error_unknown_desc",
	2900: "error - error_unknown_desc",
	2901: "error - error_unknown_desc",
	2902: "relayservice - relayservice_err_network",
	2903: "relayservice - error_alias_server_internal",
	2904: "relayservice - relayservice_err_alias_in_use",
	2905: "pkgmgr - myds_error_account",
	2906: "relayservice - error_alias_used_in_your_own",
	3001: "error - error_unknown_desc",
	3002: "relayservice - relayservice_err_resolv",
	3003: "relayservice - myds_server_internal_error",
	3004: "error - error_auth",
	3005: "relayservice - relayservice_err_alias_in_use",
	3006: "relayservice - myds_exceed_max_register_error",
	3009: "error - error_unknown_desc",
	3010: "myds - already_logged_in",
	3013: "myds - error_migrate_authen",
	3015: "myds - invalid_machine",
	3106: "user - no_such_user",
	3107: "user - error_nameused",
	3108: "user - error_nameused",
	3109: "user - error_disable_admin",
	3110: "user - error_too_much_user",
	3111: "user - homes_not_found",
	3112: "common - error_apply_occupied",
	3113: "common - error_occupied",
	3114: "user - error_nameused",
	3115: "user - user_cntrmvdefuser",
	3116: "user - user_set_fail",
	3117: "user - user_quota_set_fail",
	3118: "common - error_no_enough_space",
	3119: "user - error_home_is_moving",
	3121: "common - err_pass",
	3122: "login - password_in_history",
	3123: "login - password_too_common",
	3124: "common - err_pass",
	3130: "user - invalid_syntax_enclosed_trailing",
	3131: "user - invalid_syntax_double_quote_in_middle",
	3132: "user - invalid_syntax_not_double_quote_ending",
	3191: "user - user_file_open_fail",
	3192: "user - user_file_empty",
	3193: "user - user_file_not_utf8",
	3194: "user - user_upload_no_volume",
	3202: "common - error_occupied",
	3204: "group - failed_load_group",
	3205: "group - failed_load_group",
	3206: "group - error_nameused",
	3207: "group - error_nameused",
	3208: "group - error_badname",
	3209: "group - error_toomanygr",
	3210: "group - error_rmmember",
	3217: "group - error_too_many_dir_admin",
	3221: "share - error_too_many_acl_rules",
	3299: "common - error_system",
	3301: "share - share_already_exist",
	3302: "share - share_acl_volume_not_support",
	3303: "share - error_encrypt_reserve",
	3304: "share - error_volume_not_found",
	3305: "share - error_badname",
	3308: "share - encryption_wrong_key",
	3309: "share - error_toomanysh",
	3312: "share - share_normal_folder_exist",
	3313: "share - error_volume_not_found",
	3314: "share - error_volume_read_only",
	3319: "share - error_nameused",
	3320: "share - share_space_not_enough",
	3321: "share - error_too_many_acl_rules",
	3322: "share - mount_point_not_empty",
	3323: "error - error_mount_point_change_vol",
	3324: "error - error_mount_point_rename",
	3326: "share - error_key_file",
	3327: "share - share_conflict_on_new_volume",
	3328: "share - get_lock_failed",
	3329: "share - error_toomanysnapshot",
	3330: "share - share_snapshot_busy",
	3332: "backup - is_backing_up_restoring",
	3334: "share - error_mount_point_restore",
	3335: "share - share_cannot_move_fstype_not_support",
	3336: "share - share_cannot_move_replica_busy",
	3337: "snapmgr - snap_system_preserved",
	3338: "share - error_mounted_encrypt_restore",
	3340: "snapmgr - snap_restore_share_conf_err",
	3341: "snapmgr - err_quota_is_not_enough",
	3344: "keymanager - error_invalid_passphrase",
	3345: "keymanager - error_used_keystore",
	3347: "share - umount_fail",
	3350: "share - error_worm_snapshot",
	3400: "error - error_error_system",
	3401: "error - error_error_system",
	3402: "error - error_error_system",
	3403: "app_privilege - error_no_such_user_or_group",
	3404: "error - error_privilege_not_enough",
	3405: "app_privilege - error_wrong_data_format",
	3500: "error - error_invalid",
	3501: "common - error_badport",
	3502: "error - error_port_conflict",
	3503: "error - error_port_conflict",
	3504: "error - error_port_conflict",
	3505: "app_port_alias - err_fqdn_duplicated",
	3506: "error - err_bad_server_header",
	3510: "error - error_invalid",
	3511: "app_port_alias - err_port_dup",
	3550: "volume - volume_no_volumes",
	3551: "error - error_no_shared_folder",
	3552: fmt.Sprint("volume", "volume_crashed_service_disable"),
	3553: "volume - volume_expanding_waiting",
	3554: "error - error_port_conflict",
	3555: "common - error_badport",
	3603: "volume - volume_share_volumeno",
	3604: "error - error_space_not_enough",
	3605: "usb - usb_printer_driver_fail",
	3606: "login - error_cantlogin",
	3607: "common - error_badip",
	3608: "usb - net_prntr_ip_exist_error",
	3609: "usb - net_prntr_ip_exist_unknown",
	3610: "common - error_demo",
	3611: "usb - net_prntr_name_exist_error",
	3700: "error - error_invalid",
	3701: "status - status_not_available",
	3702: "error - error_invalid",
	3710: "status - status_not_available",
	3711: "error - error_invalid",
	3712: "cms - fan_mode_not_supported",
	3720: "status - status_not_available",
	3721: "error - error_invalid",
	3730: "status - status_not_available",
	3731: "error - error_invalid",
	3740: "status - status_not_available",
	3741: "error - error_invalid",
	3750: "status - status_not_available",
	3751: "error - error_invalid",
	3760: "status - status_not_available",
	3761: "error - error_invalid",
	3795: "error - error_port_conflict",
	3800: "error - error_invalid",
	3801: "error - error_invalid",
	400:  "Invalid parameter of file operation",
	4001: "error - error_error_system",
	4002: "dsmoption - error_format",
	4003: "dsmoption - error_size",
	4005: "dsmoption - error_logo_size",
	4006: "dsmoption - error_background_size",
	401:  "Unknown error of file operation",
	402:  "System is too busy",
	403:  "Invalid user does this file operation",
	404:  "Invalid group does this file operation",
	405:  "Invalid user and group does this file operation",
	406:  "Can't get user/group information from the account server",
	407:  "Operation not permitted",
	408:  "No such file or directory",
	409:  "Non-supported file system",
	410:  "Failed to connect internet-based file system (e.g., CIFS)",
	4100: "error - error_invalid",
	4101: "error - error_invalid",
	4102: "app_port_alias - err_alias_refused",
	4103: "app_port_alias - err_alias_used",
	4104: "app_port_alias - err_port_used",
	4105: "app_port_alias - err_port_used",
	4106: "app_port_alias - err_port_used",
	4107: "app_port_alias - err_fqdn_duplicated",
	411:  "Read-only file system",
	412:  "Filename too long in the non-encrypted file system",
	413:  "Filename too long in the encrypted file system",
	414:  "File already exists",
	415:  "Disk quota exceeded",
	4154: "app_port_alias - err_fqdn_duplicated",
	4155: "app_port_alias - err_port_used",
	4156: "app_port_alias - err_invalid_backend_host",
	416:  "No space left on device",
	4164: "app_port_alias - err_invalid_header_name",
	4165: "app_port_alias - err_invalid_header_value",
	4166: "app_port_alias - err_header_name_duplicated",
	4168: "app_port_alias - err_proxy_timeout",
	4169: "app_port_alias - err_proxy_timeout",
	417:  "Input/output error",
	4170: "app_port_alias - err_proxy_timeout",
	418:  "Illegal name or path",
	419:  "Illegal file name",
	420:  "Illegal file name on FAT file system",
	421:  "Device or resource busy",
	4300: "error - error_error_system",
	4301: "error - error_error_system",
	4302: "error - error_error_system",
	4303: "error - error_invalid",
	4304: "error - error_error_system",
	4305: "error - error_error_system",
	4306: "error - error_error_system",
	4307: "error - error_error_system",
	4308: "error - error_error_system",
	4309: "error - error_invalid",
	4310: "error - error_error_system",
	4311: "network - interface_not_found",
	4312: "tcpip - tcpip_ip_used",
	4313: "tcpip - ipv6_ip_used",
	4314: "tunnel - tunnel_conn_fail",
	4315: "tcpip - ipv6_err_link_local",
	4316: "network - error_applying_network_setting",
	4317: "common - error_notmatch",
	4319: "error - error_error_system",
	4320: "vpnc - name_conflict",
	4321: "vpnc - err_illegal_ca",
	4322: "service - service_illegel_key",
	4323: "service - service_ca_not_utf8",
	4324: "service - service_unknown_cipher",
	4325: "vpnc - l2tp_conflict",
	4326: "vpnc - vpns_conflict",
	4327: "vpnc - ovpnfile_invalid_format",
	4328: "vpnc - ovpn_private_key_not_support",
	4340: "background_task - task_processing",
	4350: "tcpip - ipv6_invalid_config",
	4351: "tcpip - ipv6_router_bad_lan_req",
	4352: "tcpip - ipv6_router_err_enable",
	4353: "tcpip - ipv6_router_err_disable",
	4354: "tcpip - ipv6_no_public_ip",
	4370: "ovs - ovs_not_support_bonding",
	4371: "ovs - ovs_not_support_vlan",
	4372: "ovs - ovs_not_support_bridge",
	4373: "network - linkaggr_mode_inconsistent_err",
	4380: "router_networktools - ping_target_invalid",
	4381: "router_networktools - ping_timeout",
	4382: "router_networktools - traceroute_target_invalid",
	4500: "error - error_error_system",
	4501: "error - error_error_system",
	4502: "pkgmgr - pkgmgr_space_not_ready",
	4503: "error - volume_creating",
	4504: "pkgmgr - error_sys_no_space",
	4506: "pkgmgr - noncancellable",
	4520: "error - error_space_not_enough",
	4521: "pkgmgr - pkgmgr_file_not_package",
	4522: "pkgmgr - broken_package",
	4529: "pkgmgr - pkgmgr_pkg_cannot_upgrade",
	4530: "pkgmgr - error_occupied",
	4531: "pkgmgr - pkgmgr_not_syno_publish",
	4532: "pkgmgr - pkgmgr_unknown_publisher",
	4533: "pkgmgr - pkgmgr_cert_expired",
	4534: "pkgmgr - pkgmgr_cert_revoked",
	4535: "pkgmgr - broken_package",
	4540: "pkgmgr - pkgmgr_file_install_failed",
	4541: "pkgmgr - upgrade_fail",
	4542: "error - error_error_system",
	4543: "pkgmgr - pkgmgr_file_not_package",
	4544: "pkgmgr - pkgmgr_pkg_install_already",
	4545: "pkgmgr - pkgmgr_pkg_not_available",
	4548: "pkgmgr - install_version_less_than_limit",
	4549: "pkgmgr - depend_cycle",
	4570: "common - error_invalid_serial",
	4580: "pkgmgr - pkgmgr_pkg_start_failed",
	4581: "pkgmgr - pkgmgr_pkg_stop_failed",
	4590: "pkgmgr - invalid_feed",
	4591: "pkgmgr - duplicate_feed",
	4592: "pkgmgr - duplicate_certificate",
	4593: "pkgmgr - duplicate_certificate_sys",
	4594: "pkgmgr - revoke_certificate",
	4595: "service - service_illegel_crt",
	4600: "error - error_error_system",
	4601: "error - error_error_system",
	4602: "notification - google_auth_failed",
	4603: "notification - outlook_auth_failed",
	4631: "error - error_error_system",
	4632: "error - error_error_system",
	4633: "error - sms_provider_not_found",
	4634: "error - sms_provider_exist",
	4635: "error - error_error_system",
	4661: "pushservice - error_update_ds_info",
	4681: "error - error_error_system",
	4682: "error - error_error_system",
	4683: "error - webhook_provider_not_found",
	4684: "error - webhook_provider_exist",
	4685: "error - error_error_system",
	4800: "schedule - error_unknown",
	4801: "schedule - error_load_failed",
	4802: "schedule - error_delete_failed",
	4803: "schedule - error_run_failed",
	4804: "schedule - error_save_failed",
	4900: "error - error_invalid",
	4901: "error - error_error_system",
	4902: "user - no_such_user",
	4903: "report - err_dest_share_not_exist",
	4904: "error - error_file_exist",
	4905: "error - error_space_not_enough",
	5001: "error - error_invalid",
	5002: "error - error_invalid",
	5003: "error - error_invalid",
	5004: "error - error_invalid",
	5005: "syslog - err_server_disconnected",
	5006: "syslog - service_ca_copy_failed",
	5007: "syslog - service_ca_copy_failed",
	5008: "log - no_active_log",
	5009: "error - error_port_conflict",
	5010: "error - error_invalid",
	5011: "error - error_invalid",
	5012: "syslog - err_name_conflict",
	5016: fmt.Sprint("error", "error_load_system_settings"),
	5017: fmt.Sprint("error", "error_load_system_settings"),
	5018: fmt.Sprint("error", "error_unexpected_load_settings"),
	5019: fmt.Sprint("error", "error_load_system_settings"),
	5020: fmt.Sprint("error", "error_set_system_settings"),
	5021: fmt.Sprint("error", "error_load_system_settings"),
	5022: "syslog - error_quota_not_enough",
	5023: "error - error_space_not_enough",
	5024: "syslog - error_sys_no_space",
	5100: "error - error_invalid",
	5101: "error - error_invalid",
	5102: "error - error_invalid",
	5103: "error - error_invalid",
	5104: "error - error_invalid",
	5105: "error - error_invalid",
	5106: "error - error_invalid",
	5202: "update - error_apply_lock",
	5203: "volume - volume_busy_waiting",
	5205: "update - error_bad_dsm_version",
	5206: "update - error_downgrade",
	5207: "update - error_model",
	5208: "update - error_apply_lock",
	5209: "update - error_patch",
	5210: "update - error_passive_patch",
	5211: "update - upload_err_no_space",
	5213: "pkgmgr - error_occupied",
	5214: "update - check_new_dsm_err",
	5215: "error - error_space_not_enough",
	5216: "error - error_fs_ro",
	5217: "error - error_dest_no_path",
	5219: "update - autoupdate_cancel_failed_running",
	5220: "update - autoupdate_cancel_failed_no_task",
	5221: "update - autoupdate_cancel_failed",
	5222: "update - error_verify_patch",
	5223: "update - error_updater_prehook_failed",
	5224: "update - error_hybrid_ha_patch_version_inconsistent",
	5225: "update - error_hybrid_ha_passive_bad_model",
	5300: "error - error_invalid",
	5301: "user - no_such_user",
	5510: "service - service_illegel_crt",
	5511: "service - service_illegel_key",
	5512: "service - service_illegal_inter_crt",
	5513: "service - service_unknown_cypher",
	5514: "service - service_key_not_match",
	5515: "service - service_ca_copy_failed",
	5516: "service - service_ca_not_utf8",
	5517: "certificate - inter_and_crt_verify_error",
	5518: "certificate - not_support_dsa",
	5519: "service - service_illegal_csr",
	5520: "backup - general_backup_destination_no_response",
	5521: "certificate - err_connection",
	5522: "certificate - err_server_not_match",
	5523: "certificate - err_too_many_reg",
	5524: "certificate - err_too_many_req",
	5525: "certificate - err_mail",
	5526: "s2s - err_invalid_param_value",
	5527: "certificate - err_le_server_busy",
	5528: "certificate - err_not_synoddns",
	5529: "certificate - err_invalid_domain",
	5530: "certificate - err_challenge_unauthorized",
	5531: "certificate - err_no_realname_verified",
	5532: "certificate - err_free_limit_exceeded",
	5533: "certificate - err_illigal_access_token",
	5534: "certificate - err_insecure_certificate",
	5600: "error - error_no_path",
	5601: "file - error_bad_file_content",
	5602: "error - error_error_system",
	5603: "texteditor - LoadFileFail",
	5604: "texteditor - SaveFileFail",
	5605: "error - error_privilege_not_enough",
	5606: "texteditor - CodepageConvertFail",
	5607: "texteditor - AskForceSave",
	5608: "error - error_encryption_long_path",
	5609: "error - error_long_path",
	5610: "error - error_quota_not_enough",
	5611: "error - error_space_not_enough",
	5612: "error - error_io",
	5613: "error - error_privilege_not_enough",
	5614: "error - error_fs_ro",
	5615: "error - error_file_exist",
	5616: "error - error_no_path",
	5617: "error - error_dest_no_path",
	5618: "error - error_testjoin",
	5619: "error - error_reserved_name",
	5620: "error - error_fat_reserved_name",
	5621: "texteditor - exceed_load_max",
	5623: "texteditor - file_locked",
	5703: "time - ntp_service_disable_warning",
	5800: "error - error_invalid",
	5801: "share - no_such_share",
	5901: "error - error_subject",
	5902: "firewall - firewall_vpnpassthrough_restore_failed",
	5903: "firewall - firewall_vpnpassthrough_specific_platform",
	599:  "No such task of the file operation",
	6001: "error - error_error_system",
	6002: "error - error_error_system",
	6003: "error - error_error_system",
	6004: "common - loadsetting_fail",
	6005: "error - error_subject",
	6006: "error - error_service_start_failed",
	6007: "error - error_service_stop_failed",
	6008: "error - error_service_start_failed",
	6009: "firewall - firewall_save_failed",
	6010: "common - error_badip",
	6011: "common - error_badip",
	6012: "common - error_badip",
	6013: "share - no_such_share",
	6014: "cms - cms_no_volumes",
	6015: "ftp - tftp_no_privilege_restart_service",
	6016: "ftp - tftp_invalid_root_folder",
	6200: "error - error_error_system",
	6201: "error - error_acl_volume_not_support",
	6202: "error - error_fat_privilege",
	6203: "error - error_remote_privilege",
	6204: "error - error_fs_ro",
	6205: "error - error_privilege_not_enough",
	6206: "error - error_no_path",
	6207: "error - error_no_path",
	6208: "error - error_testjoin",
	6209: "error - error_privilege_not_enough",
	6210: "acl_editor - admin_cannot_set_acl_perm",
	6211: "acl_editor - error_invalid_user_or_group",
	6212: "error - error_acl_mp_not_support",
	6213: "acl_editor - quota_exceeded",
	6215: "acl_editor - acl_rules_reach_limit",
	6216: "acl_editor - error_xattr_exceeded",
	6703: "error - error_port_conflict",
	6704: "error - error_port_conflict",
	6705: "user - no_such_user",
	6706: "user - error_nameused",
	6708: "share - error_volume_not_found",
	6709: "netbackup - err_create_service_share",
	7100: "connections - error_disable_admin_name",
	8001: "network - net_daemon_not_ready",
	8002: "network - usbmodem_daemon_not_ready",
	8003: "wireless_ap - ap_ssid_limit_alert",
	8010: "router_topology - get_topology_fail",
	8011: "network - net_get_fail",
	8012: "network - net_get_setting_fail",
	8013: "router_wireless - wifi_setting_get_fail",
	8020: "router_topology - set_topology_fail",
	8021: "network - net_set_fail",
	8022: "network - net_set_setting_fail",
	8023: "router_wireless - wifi_setting_set_fail",
	8030: "network - get_redirect_info_fail",
	8031: "router_common - dhcp_range_conflict_err",
	8100: "router_wireless - guest_network_get_count_down_fail",
	8101: "router_wireless - guest_network_set_count_down_fail",
	8130: "pppoe - pppoe_get_setting_fail",
	8131: "pppoe - pppoe_set_setting_fail",
	8132: "pppoe - pppoe_no_interface_available",
	8133: "pppoe - pppoe_service_start_fail",
	8134: "pppoe - pppoe_service_stop_fail",
	8135: "pppoe - pppoe_connection_failed",
	8136: "pppoe - pppoe_disconnect_fail",
	8150: "wireless_ap - country_code_get_fail",
	8151: "wireless_ap - country_code_set_fail",
	8152: "wireless_ap - country_code_read_list_fail",
	8153: "wireless_ap - country_code_region_not_support",
	8170: "routerconf - routerconf_exceed_max_rule",
	8175: "smartwan - sw_too_many_rules",
	8180: "routerconf - routerconf_exceed_max_rule",
	8190: "router_parental - err_device_reach_max",
	8200: "router_parental - err_domain_name_reach_max",
	8230: "routerconf - routerconf_exceed_max_reservation",
	8231: "routerconf - routerconf_exceed_max_reservation_v6",
	9060: "disk_info - disk_upload_db_error_verify",
	9061: "error - error_space_not_enough",
	9062: "disk_info - disk_upload_db_error_version_too_old",
	9063: "disk_info - disk_upload_db_error_model_not_match",
	100:  "Unknown error",
}

var loginErrors ErrorSummaries = func() ErrorSummary {
	return ErrorSummary{
		400: "No such account or incorrect password",
		401: "Disabled account",
		402: "Denied permission",
		403: "2-factor authentication code required",
		404: "Failed to authenticate 2-factor authentication code (invalid or reused OTP)", // https://datatracker.ietf.org/doc/html/rfc6238#autoid-11
		406: "Enforce to authenticate with 2-factor authentication code",
		407: "Blocked IP source",
		408: "Expired password cannot change",
		409: "Expired password",
		410: "Password must be changed",
	}
}

func (es ErrorSummary) Combine(params ...ErrorSummary) ErrorSummary {
	for _, p := range params {
		maps.Copy(es, p)
	}

	return es
}

// ErrorDescriber defines interface to report all known errors to particular object.
type ErrorDescriber interface {
	// ErrorSummaries returns information about all known errors.
	ErrorSummaries() []ErrorSummary
}

// ApiError defines a structure for error object returned by Synology API.
// It is a high-level error for a particular API family.
type ApiError struct {
	Code    int           `json:"code,omitempty"`
	Summary string        `json:"-"`
	Errors  []ErrorFields `json:"errors,omitempty"`

	underlying error `json:"-"`
}

// ErrorSummary is a simple mapping of code->text to translate error codes to readable text.
type ErrorSummary map[int]string

// ErrorFields defines extra fields for particular detailed error.
// All ErrorFields entries will always include a "code" field with an integer value.
type ErrorFields struct {
	Code   int            `json:"code"`
	Fields map[string]any `json:",inline"`
}

func (ef ErrorFields) WithSummaries(knownErrors ErrorSummaries) error {
	var err error

	// Handle the code field
	err = multierror.Append(err, fmt.Errorf("code: %d", ef.Code))

	// Handle additional fields
	for k, v := range ef.Fields {
		b, e := json.Marshal(v)
		if e != nil {
			err = multierror.Append(err, fmt.Errorf("%s: %v", k, e))
		} else {
			err = multierror.Append(err, fmt.Errorf("%s: %v", k, string(b)))
		}
	}
	return err
}

func (s ApiError) WithSummaries(errorSummaries ErrorSummaries) error {
	s.Summary = DescribeError(s.Code, errorSummaries())
	return s
}

func (ef ErrorFields) Underlying() error {
	var err error

	// Handle the code field
	err = multierror.Append(err, fmt.Errorf("code: %d", ef.Code))

	// Handle additional fields
	for k, v := range ef.Fields {
		b, e := json.Marshal(v)
		if e != nil {
			err = multierror.Append(err, fmt.Errorf("%s: %v", k, e))
		} else {
			err = multierror.Append(err, fmt.Errorf("%s: %v", k, string(b)))
		}
	}
	return err
}

func (ef ErrorFields) Prefix(prefix string) error {
	return multierror.Prefix(ef.Underlying(), prefix)
}

func (ef ErrorFields) Error() string {
	return ef.Underlying().Error()
}

// Error satisfies error interface for SynologyError type.
func (se ApiError) Error() string {
	se.underlying = fmt.Errorf("[%d] %s", se.Code, se.Summary)

	if len(se.Errors) > 0 {
		for _, errorFields := range se.Errors {
			se.underlying = multierror.Append(se.underlying, errorFields)
		}
	}

	return se.underlying.Error()
}

// DescribeError translates error code to human-readable summary text.
// It accepts error code and number of summary maps to look in.
// First summary with this code wins.
func DescribeError(code int, summaries ...ErrorSummary) string {
	for _, summaryMap := range summaries {
		if summary, ok := summaryMap[code]; ok {
			return summary
		}
	}

	return "Unknown error code " + strconv.Itoa(code)
}

// UnmarshalJSON implements custom JSON unmarshalling for ApiError.
// This method automatically populates the Summary field based on the error Code
// using global error mappings.
func (ae *ApiError) UnmarshalJSON(data []byte) error {
	// Create a temporary struct to unmarshal the basic fields
	type Alias ApiError
	temp := &struct {
		*Alias
	}{
		Alias: (*Alias)(ae),
	}

	// Unmarshal into the temporary struct
	if err := json.Unmarshal(data, temp); err != nil {
		return fmt.Errorf("failed to unmarshal ApiError: %w", err)
	}

	// Automatically populate the Summary field based on the Code
	if ae.Code != 0 {
		ae.Summary = DescribeError(ae.Code, GlobalErrors())
	}

	return nil
}

// UnmarshalJSON implements custom JSON unmarshaling for ErrorFields.
// This ensures that the code field is properly extracted while preserving other fields.
func (ef *ErrorFields) UnmarshalJSON(data []byte) error {
	// First unmarshal into a generic map
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("failed to unmarshal ErrorFields: %w", err)
	}

	// Extract the code field
	if codeVal, exists := raw["code"]; exists {
		switch v := codeVal.(type) {
		case float64:
			ef.Code = int(v)
		case int:
			ef.Code = v
		default:
			return fmt.Errorf("code field must be an integer, got %T", v)
		}
		delete(raw, "code") // Remove code from the raw map
	}

	// Store remaining fields
	ef.Fields = raw

	return nil
}
