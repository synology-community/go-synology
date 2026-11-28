package core

// SYNO.Core.Share.Permission -- share-level ACL rows for local users and
// groups.
//
// `list` accepts a direct call. `set` does not: it returns 403 regardless of
// parameters (including none at all) unless wrapped in a SYNO.Entry.Request
// compound envelope, exactly as DSM's own Control Panel dispatches it. See
// PLAT-705 and Client.SharePermissionSet.

const (
	ShareUserGroupTypeLocalUser  = "local_user"
	ShareUserGroupTypeLocalGroup = "local_group"
)

// SharePermission is one row as returned by SYNO.Core.Share.Permission list.
type SharePermission struct {
	Name       string `json:"name"`
	IsAdmin    bool   `json:"is_admin"`
	IsCustom   bool   `json:"is_custom"`
	IsDeny     bool   `json:"is_deny"`
	IsReadonly bool   `json:"is_readonly"`
	IsWritable bool   `json:"is_writable"`
}

// SharePermissionListRequest lists permission rows for one share.
type SharePermissionListRequest struct {
	Name          string `url:"name"`
	UserGroupType string `url:"user_group_type"`
	Offset        int    `url:"offset,omitempty"`
	Limit         int    `url:"limit,omitempty"`
}

// SharePermissionListResponse is the SYNO.Core.Share.Permission list payload.
type SharePermissionListResponse struct {
	Items []SharePermission `json:"items"`
	Total int               `json:"total,omitempty"`
}

// SharePermissionSetEntry is one row sent to SYNO.Core.Share.Permission set.
// This is exactly the shape verified against the live NAS.
//
// DSM's own Control Panel sends only the rows an operator actually changed
// (store.getModifiedRecords()), never the whole permission list -- callers
// of Client.SharePermissionSet should do the same, passing a subset of
// SharePermission rows rather than the full list. Within a row that IS
// sent, every field is required, including an explicit false: these are
// plain bools with no `omitempty`, because the verified payload always
// sends all four, and DSM was verified to 403 the bare `set` call, not
// tested with a partial row -- don't assume a missing field defaults
// safely.
type SharePermissionSetEntry struct {
	Name       string `json:"name"`
	IsReadonly bool   `json:"is_readonly"`
	IsWritable bool   `json:"is_writable"`
	IsDeny     bool   `json:"is_deny"`
	IsCustom   bool   `json:"is_custom"`
}

// sharePermissionSetParams is the flat body merged onto the compound
// element's api/method/version -- see api.CompoundElement.
type sharePermissionSetParams struct {
	Name          string                    `json:"name"`
	UserGroupType string                    `json:"user_group_type"`
	Permissions   []SharePermissionSetEntry `json:"permissions"`
}
