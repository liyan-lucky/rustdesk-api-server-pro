package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/form/admin"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/errcode"
	"strconv"
	"xorm.io/xorm"
)

type MailTemplateController struct {
	basicController
}

func (c *MailTemplateController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/mail/templates/list", "HandleList")
	b.Handle("POST", "/mail/templates/add", "HandleAdd")
	b.Handle("POST", "/mail/templates/edit", "HandleEdit")
}

func (c *MailTemplateController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	id := c.Ctx.URLParamDefault("id", "")
	name := c.Ctx.URLParamDefault("name", "")
	subject := c.Ctx.URLParamDefault("subject", "")
	_type := c.Ctx.URLParamDefault("type", "")
	created_at_0 := c.Ctx.URLParamDefault("created_at[0]", "")
	created_at_1 := c.Ctx.URLParamDefault("created_at[1]", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.MailTemplate{})
		if id != "" {
			q.Where("id = ?", id)
		}
		if name != "" {
			name = "%" + name + "%"
			q.Where("name like ?", name)
		}
		if subject != "" {
			subject = "%" + subject + "%"
			q.Where("subject like ?", subject)
		}
		if _type != "" {
			q.Where("type = ?", _type)
		}
		if created_at_0 != "" && created_at_1 != "" {
			q.Where("created_at BETWEEN ? AND ?", created_at_0, created_at_1)
		}
		q.Desc("id")
		return q
	}

	pagination := db.NewPagination(currentPage, pageSize)
	templateList := make([]model.MailTemplate, 0)
	err := pagination.Paginate(query, &model.MailTemplate{}, &templateList)
	if err != nil {
		return c.dbError(err)
	}

	list := make([]iris.Map, 0)
	for _, u := range templateList {
		list = append(list, iris.Map{
			"id":         u.Id,
			"name":       u.Name,
			"type":       u.Type,
			"subject":    u.Subject,
			"contents":   u.Contents,
			"created_at": u.CreatedAt.Format(config.TimeFormat),
		})
	}
	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *MailTemplateController) HandleAdd() mvc.Result {
	user := c.GetUser()
	if user == nil || !user.IsAdmin {
		return c.Error(nil, errcode.New(errcode.ERR1010.Code, errcode.ERR1010.Message).Error())
	}
	var form admin.MailTemplateForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		c.recordMailTemplateOperationAudit("admin_mail_template_add", "", nil, nil, "failure", err.Error())
		return c.dbError(err)
	}

	if form.Name == "" {
		c.recordMailTemplateOperationAudit("admin_mail_template_add", "", nil, sanitizeMailTemplateFormForAudit(form), "failure", "MailTemplateNameEmpty")
		return c.Error(nil, errcode.New(errcode.ERR8001.Code, errcode.ERR8001.Message).Error())
	}
	if form.Subject == "" {
		c.recordMailTemplateOperationAudit("admin_mail_template_add", "", nil, sanitizeMailTemplateFormForAudit(form), "failure", "MailTemplateSubjectEmpty")
		return c.Error(nil, errcode.New(errcode.ERR8002.Code, errcode.ERR8002.Message).Error())
	}
	if form.Contents == "" {
		c.recordMailTemplateOperationAudit("admin_mail_template_add", "", nil, sanitizeMailTemplateFormForAudit(form), "failure", "MailTemplateContentsEmpty")
		return c.Error(nil, errcode.New(errcode.ERR8003.Code, errcode.ERR8003.Message).Error())
	}

	template := &model.MailTemplate{
		Name:     form.Name,
		Type:     form.Type,
		Subject:  form.Subject,
		Contents: form.Contents,
	}

	_, err = c.Db.Insert(template)
	if err != nil {
		c.recordMailTemplateOperationAudit("admin_mail_template_add", "", nil, sanitizeMailTemplateFormForAudit(form), "failure", err.Error())
		return c.dbError(err)
	}

	c.recordMailTemplateOperationAudit("admin_mail_template_add", strconv.Itoa(template.Id), nil, sanitizeMailTemplateForAudit(template), "success", "")
	return c.Success(nil, "MailTemplateAddSuccess")
}

func (c *MailTemplateController) HandleEdit() mvc.Result {
	user := c.GetUser()
	if user == nil || !user.IsAdmin {
		return c.Error(nil, errcode.New(errcode.ERR1010.Code, errcode.ERR1010.Message).Error())
	}
	var form admin.MailTemplateForm
	err := c.Ctx.ReadJSON(&form)
	if err != nil {
		c.recordMailTemplateOperationAudit("admin_mail_template_edit", "", nil, nil, "failure", err.Error())
		return c.dbError(err)
	}

	if form.Id <= 0 {
		c.recordMailTemplateOperationAudit("admin_mail_template_edit", "", nil, sanitizeMailTemplateFormForAudit(form), "failure", "DataError")
		return c.Error(nil, errcode.New(errcode.ERR8004.Code, errcode.ERR8004.Message).Error())
	}

	var before model.MailTemplate
	has, err := c.Db.Where("id = ?", form.Id).Get(&before)
	if err != nil {
		c.recordMailTemplateOperationAudit("admin_mail_template_edit", strconv.Itoa(form.Id), nil, sanitizeMailTemplateFormForAudit(form), "failure", err.Error())
		return c.dbError(err)
	}
	if !has {
		c.recordMailTemplateOperationAudit("admin_mail_template_edit", strconv.Itoa(form.Id), nil, sanitizeMailTemplateFormForAudit(form), "failure", "MailTemplateNotFound")
		return c.Error(nil, errcode.New(errcode.ERR8005.Code, errcode.ERR8005.Message).Error())
	}

	template := &model.MailTemplate{
		Name:     form.Name,
		Type:     form.Type,
		Subject:  form.Subject,
		Contents: form.Contents,
	}

	_, err = c.Db.Where("id = ?", form.Id).Update(template)
	if err != nil {
		c.recordMailTemplateOperationAudit("admin_mail_template_edit", strconv.Itoa(form.Id), sanitizeMailTemplateForAudit(&before), sanitizeMailTemplateFormForAudit(form), "failure", err.Error())
		return c.dbError(err)
	}

	var after model.MailTemplate
	_, _ = c.Db.Where("id = ?", form.Id).Get(&after)
	c.recordMailTemplateOperationAudit("admin_mail_template_edit", strconv.Itoa(form.Id), sanitizeMailTemplateForAudit(&before), sanitizeMailTemplateForAudit(&after), "success", "")
	return c.Success(nil, "MailTemplateUpdateSuccess")
}

func (c *MailTemplateController) recordMailTemplateOperationAudit(action string, resourceID string, beforeData interface{}, afterData interface{}, result string, errorMessage string) {
	actor := c.GetUser()
	cmd := core.OperationAuditCreateCommand{
		Action:       action,
		ResourceType: "mail_template",
		ResourceID:   resourceID,
		BeforeData:   auditJSON(beforeData),
		AfterData:    auditJSON(afterData),
		IP:           c.Ctx.RemoteAddr(),
		UserAgent:    c.Ctx.GetHeader("User-Agent"),
		Result:       result,
		ErrorMessage: errorMessage,
	}
	if actor != nil {
		cmd.ActorUserID = actor.Id
		cmd.ActorUsername = actor.Username
	}
	_ = c.auditService().CreateOperationAudit(cmd)
}

func sanitizeMailTemplateForAudit(template *model.MailTemplate) iris.Map {
	if template == nil {
		return nil
	}
	return iris.Map{
		"id":               template.Id,
		"name":             template.Name,
		"type":             template.Type,
		"subject":          template.Subject,
		"contents_length":  len(template.Contents),
		"contents_preview": truncateForAudit(template.Contents, 120),
	}
}

func sanitizeMailTemplateFormForAudit(form admin.MailTemplateForm) iris.Map {
	return iris.Map{
		"id":               form.Id,
		"name":             form.Name,
		"type":             form.Type,
		"subject":          form.Subject,
		"contents_length":  len(form.Contents),
		"contents_preview": truncateForAudit(form.Contents, 120),
	}
}

func truncateForAudit(value string, limit int) string {
	if limit <= 0 || len(value) <= limit {
		return value
	}
	return value[:limit]
}
