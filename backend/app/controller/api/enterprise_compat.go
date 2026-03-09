package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/util"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// EnterpriseCompatController implements a minimal persistent subset of RustDesk Pro-style
// management APIs (device-groups, user-groups, strategies, devices admin ops, users admin ops).
// It is intentionally compatibility-first and does not yet implement the full upstream permission model.
type EnterpriseCompatController struct {
	basicController
}

func (c *EnterpriseCompatController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "device-groups", "HandleDeviceGroupsList")
	b.Handle("POST", "device-groups", "HandleDeviceGroupsCreate")
	b.Handle("GET", "device-groups/{guid:string}", "HandleDeviceGroupsGet")
	b.Handle("PUT", "device-groups/{guid:string}", "HandleDeviceGroupsUpdate")
	b.Handle("DELETE", "device-groups/{guid:string}", "HandleDeviceGroupsDelete")
	b.Handle("GET", "device-groups/{guid:string}/devices", "HandleDeviceGroupsDevicesList")
	b.Handle("POST", "device-groups/{guid:string}/devices", "HandleDeviceGroupsDevicesAssign")

	b.Handle("GET", "user-groups", "HandleUserGroupsList")
	b.Handle("POST", "user-groups", "HandleUserGroupsCreate")
	b.Handle("GET", "user-groups/{guid:string}", "HandleUserGroupsGet")
	b.Handle("PUT", "user-groups/{guid:string}", "HandleUserGroupsUpdate")
	b.Handle("DELETE", "user-groups/{guid:string}", "HandleUserGroupsDelete")

	b.Handle("GET", "strategies", "HandleStrategiesList")
	b.Handle("POST", "strategies", "HandleStrategiesCreate")
	b.Handle("GET", "strategies/{guid:string}", "HandleStrategiesGet")
	b.Handle("PUT", "strategies/{guid:string}", "HandleStrategiesUpdate")
	b.Handle("DELETE", "strategies/{guid:string}", "HandleStrategiesDelete")
	b.Handle("GET", "strategies/{guid:string}/status", "HandleStrategiesStatus")
	b.Handle("POST", "strategies/assign", "HandleStrategiesAssign")

	b.Handle("GET", "devices", "HandleDevicesList")
	b.Handle("GET", "devices/{guid:string}", "HandleDevicesGet")
	b.Handle("POST", "devices/{guid:string}/enable", "HandleDevicesEnable")
	b.Handle("POST", "devices/{guid:string}/disable", "HandleDevicesDisable")
	b.Handle("POST", "devices/{guid:string}/assign", "HandleDevicesAssign")

	b.Handle("GET", "users/{guid:string}", "HandleUsersGetByGuid")
	b.Handle("POST", "users/{guid:string}/enable", "HandleUsersEnable")
	b.Handle("POST", "users/{guid:string}/disable", "HandleUsersDisable")
	b.Handle("POST", "users/disable_login_verification", "HandleUsersDisableLoginVerification")
	b.Handle("POST", "users/force-logout", "HandleUsersForceLogout")
	b.Handle("POST", "users/invite", "HandleUsersInvite")
	b.Handle("POST", "users/tfa/totp/enforce", "HandleUsersTfaTotpEnforce")
}

func (c *EnterpriseCompatController) requireAdmin() *model.User {
	user := c.GetUser()
	if user == nil || !user.IsAdmin {
		return nil
	}
	return user
}

func (c *EnterpriseCompatController) readJSONMap() (map[string]any, error) {
	var m map[string]any
	if err := c.readJSONBody(&m); err != nil {
		return nil, err
	}
	return m, nil
}

func intFromAny(v any) int {
	switch x := v.(type) {
	case float64:
		return int(x)
	case int:
		return x
	case string:
		i, _ := strconv.Atoi(strings.TrimSpace(x))
		return i
	default:
		return 0
	}
}

func boolFromAny(v any) bool {
	switch x := v.(type) {
	case bool:
		return x
	case float64:
		return x != 0
	case string:
		x = strings.ToLower(strings.TrimSpace(x))
		return x == "1" || x == "true" || x == "yes" || x == "on"
	default:
		return false
	}
}

func stringFromAny(v any) string {
	if v == nil {
		return ""
	}
	switch x := v.(type) {
	case string:
		return strings.TrimSpace(x)
	default:
		return strings.TrimSpace(fmt.Sprint(v))
	}
}

func stringSliceFromAny(v any) []string {
	if v == nil {
		return nil
	}
	var out []string
	switch x := v.(type) {
	case []any:
		for _, it := range x {
			s := stringFromAny(it)
			if s != "" {
				out = append(out, s)
			}
		}
	case []string:
		for _, s := range x {
			s = strings.TrimSpace(s)
			if s != "" {
				out = append(out, s)
			}
		}
	case string:
		for _, s := range strings.Split(x, ",") {
			s = strings.TrimSpace(s)
			if s != "" {
				out = append(out, s)
			}
		}
	}
	return out
}

func pagedResult(current, pageSize int, total int64, data any) mvc.Result {
	return mvc.Response{Object: iris.Map{
		"total":    total,
		"current":  current,
		"pageSize": pageSize,
		"data":     data,
	}}
}

func (c *EnterpriseCompatController) HandleDeviceGroupsList() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 100)
	if current < 1 {
		current = 1
	}
	if pageSize <= 0 {
		pageSize = 100
	}

	total, err := c.Db.Count(&model.DeviceGroup{})
	if err != nil {
		return c.fail(err)
	}
	var groups []model.DeviceGroup
	if err := c.Db.Desc("id").Limit(pageSize, (current-1)*pageSize).Find(&groups); err != nil {
		return c.fail(err)
	}
	data := make([]iris.Map, 0, len(groups))
	for _, g := range groups {
		cnt, _ := c.Db.Where("group_guid = ?", g.Guid).Count(&model.DeviceGroupDevice{})
		data = append(data, iris.Map{"guid": g.Guid, "name": g.Name, "device_count": cnt})
	}
	return pagedResult(current, pageSize, total, data)
}

func (c *EnterpriseCompatController) HandleDeviceGroupsCreate() mvc.Result {
	admin := c.requireAdmin()
	if admin == nil {
		return c.failMsg("Admin required!")
	}
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	name := stringFromAny(m["name"])
	if name == "" {
		return c.failMsg("name required")
	}
	g := model.DeviceGroup{Guid: util.GetUUID(), Name: name, OwnerId: admin.Id}
	if _, err := c.Db.Insert(&g); err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: iris.Map{"guid": g.Guid, "name": g.Name}}
}

func (c *EnterpriseCompatController) loadDeviceGroup(guid string) (*model.DeviceGroup, error) {
	var g model.DeviceGroup
	has, err := c.Db.Where("guid = ?", guid).Get(&g)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errAddressBookNotFound
	}
	return &g, nil
}

func (c *EnterpriseCompatController) HandleDeviceGroupsGet() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	g, err := c.loadDeviceGroup(guid)
	if err != nil {
		return c.failMsg("device group not found")
	}
	cnt, _ := c.Db.Where("group_guid = ?", g.Guid).Count(&model.DeviceGroupDevice{})
	return mvc.Response{Object: iris.Map{"guid": g.Guid, "name": g.Name, "device_count": cnt}}
}

func (c *EnterpriseCompatController) HandleDeviceGroupsUpdate() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	update := model.DeviceGroup{}
	if name := stringFromAny(m["name"]); name != "" {
		update.Name = name
	}
	if update.Name == "" {
		return c.ok()
	}
	if _, err := c.Db.Where("guid = ?", guid).Cols("name").Update(&update); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleDeviceGroupsDelete() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	session := c.Db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return c.fail(err)
	}
	if _, err := session.Where("group_guid = ?", guid).Delete(&model.DeviceGroupDevice{}); err != nil {
		_ = session.Rollback()
		return c.fail(err)
	}
	if _, err := session.Where("guid = ?", guid).Delete(&model.DeviceGroup{}); err != nil {
		_ = session.Rollback()
		return c.fail(err)
	}
	if err := session.Commit(); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleDeviceGroupsDevicesList() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 100)
	if current < 1 {
		current = 1
	}
	if pageSize <= 0 {
		pageSize = 100
	}
	var links []model.DeviceGroupDevice
	total, err := c.Db.Where("group_guid = ?", guid).Count(&model.DeviceGroupDevice{})
	if err != nil {
		return c.fail(err)
	}
	if err := c.Db.Where("group_guid = ?", guid).Limit(pageSize, (current-1)*pageSize).Find(&links); err != nil {
		return c.fail(err)
	}
	data := make([]iris.Map, 0, len(links))
	for _, link := range links {
		var d model.Device
		has, _ := c.Db.Where("rustdesk_id = ?", link.RustdeskId).Get(&d)
		row := iris.Map{"guid": link.RustdeskId, "id": link.RustdeskId}
		if has {
			row["hostname"] = d.Hostname
			row["username"] = d.Username
			row["os"] = d.Os
			row["disabled"] = d.Disabled
		}
		data = append(data, row)
	}
	return pagedResult(current, pageSize, total, data)
}

func (c *EnterpriseCompatController) HandleDeviceGroupsDevicesAssign() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	ids := stringSliceFromAny(m["device_guids"])
	if len(ids) == 0 {
		ids = stringSliceFromAny(m["device_ids"])
	}
	if len(ids) == 0 {
		ids = stringSliceFromAny(m["ids"])
	}
	session := c.Db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return c.fail(err)
	}
	if _, err := session.Where("group_guid = ?", guid).Delete(&model.DeviceGroupDevice{}); err != nil {
		_ = session.Rollback()
		return c.fail(err)
	}
	for _, id := range ids {
		if _, err := session.Insert(&model.DeviceGroupDevice{GroupGuid: guid, RustdeskId: id}); err != nil {
			_ = session.Rollback()
			return c.fail(err)
		}
	}
	if err := session.Commit(); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleUserGroupsList() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 100)
	total, err := c.Db.Count(&model.UserGroup{})
	if err != nil {
		return c.fail(err)
	}
	var groups []model.UserGroup
	if err := c.Db.Desc("id").Limit(pageSize, (current-1)*pageSize).Find(&groups); err != nil {
		return c.fail(err)
	}
	data := make([]iris.Map, 0, len(groups))
	for _, g := range groups {
		cnt, _ := c.Db.Where("group_guid = ?", g.Guid).Count(&model.UserGroupMember{})
		data = append(data, iris.Map{"guid": g.Guid, "name": g.Name, "user_count": cnt})
	}
	return pagedResult(current, pageSize, total, data)
}

func (c *EnterpriseCompatController) HandleUserGroupsCreate() mvc.Result {
	admin := c.requireAdmin()
	if admin == nil {
		return c.failMsg("Admin required!")
	}
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	name := stringFromAny(m["name"])
	if name == "" {
		return c.failMsg("name required")
	}
	g := model.UserGroup{Guid: util.GetUUID(), Name: name, OwnerId: admin.Id}
	if _, err := c.Db.Insert(&g); err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: iris.Map{"guid": g.Guid, "name": g.Name}}
}

func (c *EnterpriseCompatController) HandleUserGroupsGet() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	var g model.UserGroup
	has, err := c.Db.Where("guid = ?", guid).Get(&g)
	if err != nil {
		return c.fail(err)
	}
	if !has {
		return c.failMsg("user group not found")
	}
	cnt, _ := c.Db.Where("group_guid = ?", g.Guid).Count(&model.UserGroupMember{})
	return mvc.Response{Object: iris.Map{"guid": g.Guid, "name": g.Name, "user_count": cnt}}
}

func (c *EnterpriseCompatController) HandleUserGroupsUpdate() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	if name := stringFromAny(m["name"]); name != "" {
		if _, err := c.Db.Where("guid = ?", guid).Cols("name").Update(&model.UserGroup{Name: name}); err != nil {
			return c.fail(err)
		}
	}
	// Optional full member replace.
	userRefs := stringSliceFromAny(m["user_guids"])
	if len(userRefs) == 0 {
		userRefs = stringSliceFromAny(m["usernames"])
	}
	if len(userRefs) > 0 {
		session := c.Db.NewSession()
		defer session.Close()
		if err := session.Begin(); err != nil {
			return c.fail(err)
		}
		if _, err := session.Where("group_guid = ?", guid).Delete(&model.UserGroupMember{}); err != nil {
			_ = session.Rollback()
			return c.fail(err)
		}
		for _, ref := range userRefs {
			var u model.User
			has, err := session.Where("username = ? OR id = ?", ref, intFromAny(ref)).Get(&u)
			if err != nil {
				_ = session.Rollback()
				return c.fail(err)
			}
			if has {
				if _, err := session.Insert(&model.UserGroupMember{GroupGuid: guid, UserId: u.Id}); err != nil {
					_ = session.Rollback()
					return c.fail(err)
				}
			}
		}
		if err := session.Commit(); err != nil {
			return c.fail(err)
		}
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleUserGroupsDelete() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	session := c.Db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return c.fail(err)
	}
	if _, err := session.Where("group_guid = ?", guid).Delete(&model.UserGroupMember{}); err != nil {
		_ = session.Rollback()
		return c.fail(err)
	}
	if _, err := session.Where("guid = ?", guid).Delete(&model.UserGroup{}); err != nil {
		_ = session.Rollback()
		return c.fail(err)
	}
	if err := session.Commit(); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleStrategiesList() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 100)
	total, err := c.Db.Count(&model.Strategy{})
	if err != nil {
		return c.fail(err)
	}
	var rows []model.Strategy
	if err := c.Db.Desc("id").Limit(pageSize, (current-1)*pageSize).Find(&rows); err != nil {
		return c.fail(err)
	}
	data := make([]iris.Map, 0, len(rows))
	for _, s := range rows {
		data = append(data, iris.Map{
			"guid":    s.Guid,
			"name":    s.Name,
			"enabled": s.Enabled,
		})
	}
	return pagedResult(current, pageSize, total, data)
}

func (c *EnterpriseCompatController) HandleStrategiesCreate() mvc.Result {
	admin := c.requireAdmin()
	if admin == nil {
		return c.failMsg("Admin required!")
	}
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	name := stringFromAny(m["name"])
	if name == "" {
		return c.failMsg("name required")
	}
	contentBytes, _ := json.Marshal(m["content"])
	s := model.Strategy{
		Guid:    util.GetUUID(),
		Name:    name,
		Content: string(contentBytes),
		Enabled: boolFromAny(m["enabled"]),
		OwnerId: admin.Id,
	}
	if _, err := c.Db.Insert(&s); err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: iris.Map{"guid": s.Guid, "name": s.Name, "enabled": s.Enabled}}
}

func (c *EnterpriseCompatController) HandleStrategiesGet() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	var s model.Strategy
	has, err := c.Db.Where("guid = ?", guid).Get(&s)
	if err != nil {
		return c.fail(err)
	}
	if !has {
		return c.failMsg("strategy not found")
	}
	var content any
	_ = json.Unmarshal([]byte(s.Content), &content)
	return mvc.Response{Object: iris.Map{
		"guid":    s.Guid,
		"name":    s.Name,
		"enabled": s.Enabled,
		"content": content,
	}}
}

func (c *EnterpriseCompatController) HandleStrategiesUpdate() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	cols := make([]string, 0, 3)
	update := model.Strategy{}
	if name := stringFromAny(m["name"]); name != "" {
		update.Name = name
		cols = append(cols, "name")
	}
	if _, ok := m["enabled"]; ok {
		update.Enabled = boolFromAny(m["enabled"])
		cols = append(cols, "enabled")
	}
	if content, ok := m["content"]; ok {
		b, _ := json.Marshal(content)
		update.Content = string(b)
		cols = append(cols, "content")
	}
	if len(cols) == 0 {
		return c.ok()
	}
	if _, err := c.Db.Where("guid = ?", guid).Cols(cols...).Update(&update); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleStrategiesDelete() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	session := c.Db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return c.fail(err)
	}
	if _, err := session.Where("strategy_guid = ?", guid).Delete(&model.StrategyAssignment{}); err != nil {
		_ = session.Rollback()
		return c.fail(err)
	}
	if _, err := session.Where("guid = ?", guid).Delete(&model.Strategy{}); err != nil {
		_ = session.Rollback()
		return c.fail(err)
	}
	if err := session.Commit(); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleStrategiesStatus() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	var s model.Strategy
	has, err := c.Db.Where("guid = ?", guid).Get(&s)
	if err != nil {
		return c.fail(err)
	}
	if !has {
		return c.failMsg("strategy not found")
	}
	cnt, _ := c.Db.Where("strategy_guid = ?", guid).Count(&model.StrategyAssignment{})
	return mvc.Response{Object: iris.Map{
		"guid":         s.Guid,
		"enabled":      s.Enabled,
		"assigned_cnt": cnt,
	}}
}

func (c *EnterpriseCompatController) HandleStrategiesAssign() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	strategyGuid := stringFromAny(m["strategy_guid"])
	if strategyGuid == "" {
		strategyGuid = stringFromAny(m["guid"])
	}
	if strategyGuid == "" {
		return c.failMsg("strategy_guid required")
	}
	targetType := stringFromAny(m["target_type"])
	if targetType == "" {
		targetType = "device"
	}
	targets := stringSliceFromAny(m["target_guids"])
	if len(targets) == 0 {
		targets = stringSliceFromAny(m["targets"])
	}
	session := c.Db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return c.fail(err)
	}
	if _, err := session.Where("strategy_guid = ? AND target_type = ?", strategyGuid, targetType).Delete(&model.StrategyAssignment{}); err != nil {
		_ = session.Rollback()
		return c.fail(err)
	}
	for _, t := range targets {
		if _, err := session.Insert(&model.StrategyAssignment{
			StrategyGuid: strategyGuid,
			TargetType:   targetType,
			TargetGuid:   t,
		}); err != nil {
			_ = session.Rollback()
			return c.fail(err)
		}
	}
	if err := session.Commit(); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) resolveDevice(guid string) (*model.Device, error) {
	var d model.Device
	has, err := c.Db.Where("rustdesk_id = ? OR uuid = ?", guid, guid).Get(&d)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errAddressBookNotFound
	}
	return &d, nil
}

func (c *EnterpriseCompatController) HandleDevicesList() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 100)
	keyword := c.Ctx.URLParamDefault("keyword", "")
	countQ := c.Db.Table(&model.Device{})
	listQ := c.Db.Table(&model.Device{})
	if keyword != "" {
		like := "%" + keyword + "%"
		countQ = countQ.Where("rustdesk_id like ? OR hostname like ? OR username like ?", like, like, like)
		listQ = listQ.Where("rustdesk_id like ? OR hostname like ? OR username like ?", like, like, like)
	}
	total, err := countQ.Count(&model.Device{})
	if err != nil {
		return c.fail(err)
	}
	var rows []model.Device
	if err := listQ.Desc("id").Limit(pageSize, (current-1)*pageSize).Find(&rows); err != nil {
		return c.fail(err)
	}
	data := make([]iris.Map, 0, len(rows))
	for _, d := range rows {
		data = append(data, iris.Map{
			"guid":      d.RustdeskId,
			"id":        d.RustdeskId,
			"hostname":  d.Hostname,
			"username":  d.Username,
			"os":        d.Os,
			"version":   d.Version,
			"uuid":      d.Uuid,
			"disabled":  d.Disabled,
			"is_online": d.IsOnline,
		})
	}
	return pagedResult(current, pageSize, total, data)
}

func (c *EnterpriseCompatController) HandleDevicesGet() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	guid := c.Ctx.Params().Get("guid")
	d, err := c.resolveDevice(guid)
	if err != nil {
		return c.failMsg("device not found")
	}
	return mvc.Response{Object: iris.Map{
		"guid":      d.RustdeskId,
		"id":        d.RustdeskId,
		"hostname":  d.Hostname,
		"username":  d.Username,
		"os":        d.Os,
		"version":   d.Version,
		"uuid":      d.Uuid,
		"disabled":  d.Disabled,
		"is_online": d.IsOnline,
	}}
}

func (c *EnterpriseCompatController) setDeviceDisabled(guid string, disabled bool) mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	if _, err := c.Db.Where("rustdesk_id = ? OR uuid = ?", guid, guid).Cols("disabled").Update(&model.Device{Disabled: disabled}); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleDevicesEnable() mvc.Result {
	return c.setDeviceDisabled(c.Ctx.Params().Get("guid"), false)
}

func (c *EnterpriseCompatController) HandleDevicesDisable() mvc.Result {
	return c.setDeviceDisabled(c.Ctx.Params().Get("guid"), true)
}

func (c *EnterpriseCompatController) HandleDevicesAssign() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	deviceGuid := c.Ctx.Params().Get("guid")
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	groupGuid := stringFromAny(m["device_group_guid"])
	if groupGuid == "" {
		groupGuid = stringFromAny(m["group_guid"])
	}
	if groupGuid != "" {
		session := c.Db.NewSession()
		defer session.Close()
		if err := session.Begin(); err != nil {
			return c.fail(err)
		}
		if _, err := session.Where("rustdesk_id = ?", deviceGuid).Delete(&model.DeviceGroupDevice{}); err != nil {
			_ = session.Rollback()
			return c.fail(err)
		}
		if _, err := session.Insert(&model.DeviceGroupDevice{GroupGuid: groupGuid, RustdeskId: deviceGuid}); err != nil {
			_ = session.Rollback()
			return c.fail(err)
		}
		if err := session.Commit(); err != nil {
			return c.fail(err)
		}
	}
	return c.ok()
}

func (c *EnterpriseCompatController) resolveUserByRef(ref string) (*model.User, error) {
	var u model.User
	if id, err := strconv.Atoi(ref); err == nil {
		if has, err := c.Db.Where("id = ?", id).Get(&u); err != nil {
			return nil, err
		} else if has {
			return &u, nil
		}
	}
	has, err := c.Db.Where("username = ?", ref).Get(&u)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errAddressBookNotFound
	}
	return &u, nil
}

func (c *EnterpriseCompatController) HandleUsersGetByGuid() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	ref := c.Ctx.Params().Get("guid")
	u, err := c.resolveUserByRef(ref)
	if err != nil {
		return c.failMsg("user not found")
	}
	return mvc.Response{Object: iris.Map{
		"guid":         u.Username,
		"name":         u.Name,
		"display_name": u.Name,
		"username":     u.Username,
		"email":        u.Email,
		"note":         u.Note,
		"status":       u.Status,
		"is_admin":     u.IsAdmin,
	}}
}

func (c *EnterpriseCompatController) setUserStatus(ref string, status int) mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	u, err := c.resolveUserByRef(ref)
	if err != nil {
		return c.failMsg("user not found")
	}
	if _, err := c.Db.Where("id = ?", u.Id).Cols("status").Update(&model.User{Status: status}); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleUsersEnable() mvc.Result {
	return c.setUserStatus(c.Ctx.Params().Get("guid"), 1)
}
func (c *EnterpriseCompatController) HandleUsersDisable() mvc.Result {
	return c.setUserStatus(c.Ctx.Params().Get("guid"), 0)
}

func (c *EnterpriseCompatController) HandleUsersDisableLoginVerification() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	ref := stringFromAny(m["user_guid"])
	if ref == "" {
		ref = stringFromAny(m["username"])
	}
	if ref == "" {
		return c.failMsg("user required")
	}
	u, err := c.resolveUserByRef(ref)
	if err != nil {
		return c.failMsg("user not found")
	}
	if _, err := c.Db.Where("id = ?", u.Id).Cols("login_verify").Update(&model.User{LoginVerify: ""}); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleUsersForceLogout() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	ref := stringFromAny(m["user_guid"])
	if ref == "" {
		ref = stringFromAny(m["username"])
	}
	if ref == "" {
		return c.ok()
	}
	u, err := c.resolveUserByRef(ref)
	if err != nil {
		return c.failMsg("user not found")
	}
	if _, err := c.Db.Where("user_id = ?", u.Id).Delete(&model.AuthToken{}); err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *EnterpriseCompatController) HandleUsersInvite() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	// Compatibility response shape; email delivery workflow is project-specific.
	return mvc.Response{Object: iris.Map{
		"ok":      true,
		"message": "invite accepted",
	}}
}

func (c *EnterpriseCompatController) HandleUsersTfaTotpEnforce() mvc.Result {
	if c.requireAdmin() == nil {
		return c.failMsg("Admin required!")
	}
	m, err := c.readJSONMap()
	if err != nil {
		return c.fail(err)
	}
	enabled := boolFromAny(m["enabled"])
	targets := stringSliceFromAny(m["user_guids"])
	if len(targets) == 0 {
		targets = stringSliceFromAny(m["usernames"])
	}
	for _, ref := range targets {
		u, err := c.resolveUserByRef(ref)
		if err != nil {
			continue
		}
		lv := ""
		if enabled {
			lv = model.LOGIN_TFA_CHECK
		}
		_, _ = c.Db.Where("id = ?", u.Id).Cols("login_verify").Update(&model.User{LoginVerify: lv})
	}
	return c.ok()
}
