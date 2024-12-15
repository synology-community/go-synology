package core

import (
	"context"
	"fmt"
	"os"
	"path"
	"slices"
	"time"

	"github.com/synology-community/go-synology/pkg/api"
	"github.com/synology-community/go-synology/pkg/api/core/methods"
	"github.com/synology-community/go-synology/pkg/util/form"
)

type Client struct {
	client api.Api
}

// TaskFind implements Api.
func (c *Client) TaskFind(ctx context.Context, name string) (*TaskResult, error) {
	tasks, err := c.TaskList(ctx, ListTaskRequest{})
	if err != nil {
		return nil, err
	}
	i := slices.IndexFunc(tasks.Tasks, func(t TaskResult) bool {
		return t.Name == name
	})
	if i < 0 {
		return nil, TaskNotFoundError{}
	}
	return &tasks.Tasks[i], nil
}

// ContentLength implements Api.
func (c *Client) ContentLength(ctx context.Context, url string) (int64, error) {
	resp, err := c.client.Client().Head(url)
	if err != nil {
		return int64(0), err
	}
	return resp.ContentLength, nil
}

// SystemInfo implements CoreApi.
func (c Client) SystemInfo(ctx context.Context) (*SystemInfoResponse, error) {
	return api.Post[SystemInfoResponse, api.Request](c.client, ctx, nil, methods.SystemInfo)
}

func (c Client) PackageInstallCheck(ctx context.Context, req PackageInstallCheckRequest) (*PackageInstallCheckResponse, error) {
	return api.Get[PackageInstallCheckResponse](c.client, ctx, &req, methods.PackageInstallationCheck)
}

// PackageList implements CoreApi.
func (c Client) PackageList(ctx context.Context) (*PackageListResponse, error) {
	return api.Get[PackageListResponse](c.client, ctx, &PackageListRequest{
		Additional: []string{"description", "description_enu", "dependent_packages", "beta", "distributor", "distributor_url", "maintainer", "maintainer_url", "dsm_apps", "dsm_app_page", "dsm_app_launch_name", "report_beta_url", "support_center", "startable", "installed_info", "support_url", "is_uninstall_pages", "install_type", "autoupdate", "silent_upgrade", "installing_progress", "ctl_uninstall", "updated_at", "status", "url", "available_operation", "install_type"},
	}, methods.PackageList)
}

func (c Client) PackageFind(ctx context.Context, name string) (*Package, error) {
	for i := 0; i < 2; i++ {
		var req PackageServerListRequest
		if i == 0 {
			req.LoadOthers = false
		} else {
			req.LoadOthers = true
		}
		resp, err := c.PackageServerList(ctx, req)
		if err != nil {
			return nil, err
		}

		ii := slices.IndexFunc(resp.Packages, func(p Package) bool {
			return p.Package == name
		})
		if ii != -1 {
			return &resp.Packages[ii], nil
		}
	}
	return nil, fmt.Errorf("package %s not found", name)
}

func (c Client) PackageServerList(ctx context.Context, req PackageServerListRequest) (*PackageServerListResponse, error) {
	return api.Get[PackageServerListResponse](c.client, ctx, &req, methods.PackageServerList)
}

func (c Client) PackageGet(ctx context.Context, id string) (*PackageGetResponse, error) {
	return api.Get[PackageGetResponse](c.client, ctx, &PackageGetRequest{
		ID:         id,
		Additional: []string{"description", "description_enu", "dependent_packages", "beta", "distributor", "distributor_url", "maintainer", "maintainer_url", "dsm_apps", "dsm_app_page", "dsm_app_launch_name", "report_beta_url", "support_center", "startable", "installed_info", "support_url", "is_uninstall_pages", "install_type", "autoupdate", "silent_upgrade", "installing_progress", "ctl_uninstall", "updated_at", "status", "url", "available_operation", "install_type"},
	}, methods.PackageGet)
}

func (c Client) PackageUninstallCompound(ctx context.Context, name string) error {
	p, err := c.PackageGet(ctx, name)
	if err != nil {
		return err
	}

	dsmApps := p.Additional.DsmApps

	_, err = c.PackageUninstall(ctx, PackageUninstallRequest{
		ID:      name,
		DSMApps: dsmApps,
		ExtraValues: UninstallExtra{
			KeepData:   true,
			DeleteData: false,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func readFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c Client) PackageInstallCompound(ctx context.Context, req PackageInstallCompoundRequest) error {

	if req.File != "" {
		b, err := readFile(req.File)
		if err != nil {
			return err
		}

		fileName := path.Base(req.File)

		ctx, cancel := context.WithTimeout(ctx, 120*time.Minute)
		defer cancel()

		got, err := c.PackageInstallUpload(ctx, form.File{
			Name:    fileName,
			Content: b,
		})

		if err != nil {
			return err
		}

		req.Name = got.Name
	}

	pkgSetting, err := c.PackageSettingGet(ctx, PackageSettingGetRequest{})

	if err != nil {
		return err
	}

	defaultVol := pkgSetting.DefaultVol

	if defaultVol == "" {
		return fmt.Errorf("Default volume empty")
	}

	dlRes, err := c.PackageInstall(ctx, PackageInstallRequest{
		Name:       req.Name,
		URL:        req.URL,
		Type:       0,
		BigInstall: false,
		FileSize:   req.Size,
	})
	if err != nil {
		return err
	}

	if dlRes.TaskID == "" {
		return fmt.Errorf("Task ID empty")
	}

	status := new(PackageInstallStatusResponse)

	for retry := 0; !status.Finished; retry++ {
		status, err = c.PackageInstallStatus(ctx, PackageInstallStatusRequest{
			TaskID: dlRes.TaskID,
		})

		if err != nil {
			return err
		}

		if status.Finished {
			break
		}

		if retry > 10 {
			return fmt.Errorf("Maximum retries exceeded: Package install status - compound")
		}

		if !status.Finished {
			time.Sleep(2 * time.Second)
		}
	}

	path := fmt.Sprintf("%s/%s", status.TmpFolder, status.Taskid)

	_, err = c.PackageInstallCheck(ctx, PackageInstallCheckRequest{
		ID:                   req.Name,
		InstallType:          "",
		InstallOnColdStorage: false,
		BreakPkgs:            "",
		BlCheckDep:           false,
		ReplacePkgs:          "",
	})

	if err != nil {
		return err
	}

	instRes, err := c.PackageInstall(ctx, PackageInstallRequest{
		// Name:              status.Name,
		Path:              path,
		InstallRunPackage: false,
		Force:             true,
		CheckCodesign:     false,
		Type:              0,
		ExtraValues:       ExtraValues(req.ExtraValues),
		VolumePath:        defaultVol,
	})

	if err != nil {
		return err
	}

	if instRes.PackageName == "" {
		return fmt.Errorf("Installation package name response empty")
	}

	return nil
}

func (c Client) PackageInstallStatus(ctx context.Context, req PackageInstallStatusRequest) (*PackageInstallStatusResponse, error) {
	return api.Get[PackageInstallStatusResponse](c.client, ctx, &req, methods.PackageInstallationStatus)
}

func (c Client) PackageSettingGet(ctx context.Context, req PackageSettingGetRequest) (*PackageSettingGetResponse, error) {
	return api.Get[PackageSettingGetResponse](c.client, ctx, &req, methods.PackageSettingGet)
}

func (c Client) PackageInstall(ctx context.Context, req PackageInstallRequest) (*PackageInstallResponse, error) {
	if req.URL != "" && req.FileSize < 1 {
		size, err := c.ContentLength(ctx, req.URL)
		if err != nil {
			return nil, err
		}
		req.FileSize = size
	}

	req.Operation = []string{"install"}

	return api.Get[PackageInstallResponse](c.client, ctx, &req, methods.PackageInstallationInstall)
}

func (c Client) PackageInstallDelete(ctx context.Context, req PackageInstallDeleteRequest) error {
	return api.Void[PackageInstallDeleteRequest](c.client, ctx, &req, methods.PackageInstallationDelete)
}

func (c Client) PackageUninstall(ctx context.Context, req PackageUninstallRequest) (*PackageUninstallResponse, error) {
	return api.Get[PackageUninstallResponse](c.client, ctx, &req, methods.PackageUnistallationUninstall)
}

func (c Client) PackageInstallUpload(ctx context.Context, req form.File) (*PackageInstallUploadResponse, error) {
	return api.PostFileWithQuery[PackageInstallUploadResponse](c.client, ctx, &PackageInstallUploadRequest{
		Additional: []string{"description", "maintainer", "distributor", "startable", "dsm_apps", "status", "install_reboot", "install_type", "install_on_cold_storage", "break_pkgs", "replace_pkgs"},
		File:       req,
	}, methods.PackageInstallationUpload)
}

func (c Client) PackageFeedList(ctx context.Context) (*PackageFeedListResponse, error) {
	return api.List[PackageFeedListResponse](c.client, ctx, methods.PackageFeedList)
}

func (c Client) PackageFeedAdd(ctx context.Context, req PackageFeedAddRequest) error {
	return api.Void(c.client, ctx, &req, methods.PackageFeedAdd)
}

func (c Client) PackageFeedDelete(ctx context.Context, req PackageFeedDeleteRequest) error {
	return api.Void(c.client, ctx, &req, methods.PackageFeedDelete)
}

var additionalShareInfo = []string{"name", "hidden", "encryption", "is_aclmode", "unite_permission", "is_support_acl", "is_sync_share", "is_force_readonly", "force_readonly_reason", "recyclebin", "is_share_moving", "is_cluster_share", "is_exfat_share", "is_c2_share", "is_cold_storage_share", "is_missing_share", "is_offline_share", "support_snapshot", "share_quota", "enable_share_compress", "enable_share_cow", "enable_share_tiering", "load_worm_attr", "include_cold_storage_share", "is_cold_storage_share", "include_missing_share", "is_missing_share", "include_offline_share", "is_offline_share", "include_worm_share"}

func (c Client) ShareList(ctx context.Context) (*ShareListResponse, error) {
	return api.Get[ShareListResponse](c.client, ctx, &ShareListRequest{
		ShareType:  "all",
		Additional: additionalShareInfo,
	}, methods.ShareList)
}

func (c Client) ShareGet(ctx context.Context, name string) (*ShareGetResponse, error) {
	return api.Get[ShareGetResponse](c.client, ctx, &ShareGetRequest{
		Name:       name,
		Additional: additionalShareInfo,
	}, methods.ShareGet)
}

func (c Client) ShareGetByID(ctx context.Context, id string) (*Share, error) {
	resp, err := c.ShareList(ctx)
	if err != nil {
		return nil, err
	}
	i := slices.IndexFunc(resp.Shares, func(s Share) bool {
		return s.UUID == id
	})
	if i < 0 {
		return nil, fmt.Errorf("share not found")
	}
	return &resp.Shares[i], nil
}

func (c Client) ShareCreate(ctx context.Context, share ShareInfo) error {
	return api.Void(c.client, ctx, &ShareCreateRequest{
		Name:      share.Name,
		ShareInfo: share,
	}, methods.ShareCreate)
}

func (c Client) ShareDelete(ctx context.Context, name string) error {
	return api.Void(c.client, ctx, &ShareDeleteRequest{
		Name: name,
	}, methods.ShareDelete)
}

func (c Client) VolumeList(ctx context.Context) (*VolumeListResponse, error) {
	return api.Post[VolumeListResponse](c.client, ctx, &VolumeListRequest{
		Limit:    -1,
		Offset:   0,
		Location: "local",
	}, methods.VolumeList)
}

func (c Client) PasswordConfirm(ctx context.Context, password string) (*PasswordConfirmResponse, error) {
	return api.Post[PasswordConfirmResponse](c.client, ctx, &PasswordConfirmRequest{
		Password: password,
	}, methods.PasswordConfirm)
}

func (c *Client) RootTaskCreate(ctx context.Context, req TaskRequest) (*TaskResult, error) {
	pwtoken := req.SynoConfirmPWToken
	if pwtoken == "" {
		res, err := c.PasswordConfirm(ctx, c.client.Password())
		if err != nil {
			return nil, err
		}
		pwtoken = res.SynoConfirmPWToken
	}

	req.SynoConfirmPWToken = pwtoken

	return api.Post[TaskResult](c.client, ctx, &req, methods.RootTaskCreate)
}

func (c *Client) RootTaskUpdate(ctx context.Context, req TaskRequest) (*TaskResult, error) {
	pwtoken := req.SynoConfirmPWToken
	if pwtoken == "" {
		res, err := c.PasswordConfirm(ctx, c.client.Password())
		if err != nil {
			return nil, err
		}
		pwtoken = res.SynoConfirmPWToken
	}

	req.SynoConfirmPWToken = pwtoken

	return api.Post[TaskResult](c.client, ctx, &req, methods.RootTaskUpdate)
}

func (c *Client) RootEventCreate(ctx context.Context, req EventRequest) (*EventResult, error) {
	pwtoken := req.SynoConfirmPWToken
	if pwtoken == "" {
		res, err := c.PasswordConfirm(ctx, c.client.Password())
		if err != nil {
			return nil, err
		}
		pwtoken = res.SynoConfirmPWToken
	}

	req.SynoConfirmPWToken = pwtoken

	return api.Post[EventResult](c.client, ctx, &req, methods.RootEventCreate)
}

func (c *Client) RootEventUpdate(ctx context.Context, req EventRequest) (*EventResult, error) {
	pwtoken := req.SynoConfirmPWToken
	if pwtoken == "" {
		res, err := c.PasswordConfirm(ctx, c.client.Password())
		if err != nil {
			return nil, err
		}
		pwtoken = res.SynoConfirmPWToken
	}

	req.SynoConfirmPWToken = pwtoken

	return api.Post[EventResult](c.client, ctx, &req, methods.RootEventUpdate)
}

func (c *Client) RootEventDelete(ctx context.Context, req EventRequest) error {
	return api.Void(c.client, ctx, &req, methods.RootEventDelete)
}

func (c *Client) EventCreate(ctx context.Context, req EventRequest) (*EventResult, error) {
	return api.Post[EventResult](c.client, ctx, &req, methods.EventCreate)
}

func (c *Client) EventUpdate(ctx context.Context, req EventRequest) (*EventResult, error) {
	return api.Post[EventResult](c.client, ctx, &req, methods.EventUpdate)
}

func (c *Client) EventDelete(ctx context.Context, req EventRequest) error {
	return api.Void(c.client, ctx, &req, methods.EventDelete)
}

func (c *Client) EventRun(ctx context.Context, name string) error {
	return api.Void(c.client, ctx, &EventRequest{
		Name: name,
	}, methods.EventRun)
}

func (c *Client) EventGet(ctx context.Context, name string) (*EventRequest, error) {
	return api.Get[EventRequest](c.client, ctx, &EventRequest{
		Name: name,
	}, methods.EventGet)
}

func (c *Client) TaskCreate(ctx context.Context, req TaskRequest) (*TaskResult, error) {
	return api.Post[TaskResult](c.client, ctx, &req, methods.TaskCreate)
}

func (c *Client) TaskUpdate(ctx context.Context, req TaskRequest) (*TaskResult, error) {
	return api.Post[TaskResult](c.client, ctx, &req, methods.TaskUpdate)
}

func (c *Client) TaskDelete(ctx context.Context, ids ...int64) error {
	tasks := make([]TaskRef, len(ids))
	for i, id := range ids {
		tasks[i] = TaskRef{ID: id}
	}
	return api.Void(c.client, ctx, &TaskDeleteRequest{Tasks: tasks}, methods.TaskDelete)
}

func (c *Client) TaskGet(ctx context.Context, id int64) (*TaskResult, error) {
	return api.Get[TaskResult](c.client, ctx, &TaskGetRequest{
		ID: id,
	}, methods.TaskGet)
}

func (c *Client) TaskList(ctx context.Context, req ListTaskRequest) (*ListTaskResponse, error) {
	return api.Get[ListTaskResponse](c.client, ctx, &req, methods.TaskList)
}

func (c *Client) TaskRun(ctx context.Context, ids ...int64) error {
	tasks := make([]TaskRef, len(ids))
	for i, id := range ids {
		tasks[i] = TaskRef{ID: id}
	}
	return api.Void(c.client, ctx, &TaskRunRequest{
		Tasks: tasks,
	}, methods.TaskRun)
}

func (c *Client) UserList(ctx context.Context) (*UserListResponse, error) {
	return api.Get[UserListResponse](c.client, ctx, &UserListRequest{
		Additional: []string{"uid"},
	}, methods.UserList)
}

func New(client api.Api) Api {
	return &Client{client: client}
}
