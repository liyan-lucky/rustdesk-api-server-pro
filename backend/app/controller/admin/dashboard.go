package admin

import (
	"fmt"
	"net"
	"os"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	v2service "rustdesk-api-server-pro/internal/service"
	"strings"

	"github.com/golang-module/carbon/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type DashboardController struct {
	basicController
	Cfg *config.ServerConfig
}

func (c *DashboardController) GetDashboardStat() mvc.Result {

	userCount, err := c.Db.Count(&model.User{})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	deviceCount, err := c.Db.Count(&model.Device{})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	onlineCount, err := c.Db.Where("is_online = 1").Count(&model.Device{})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	visitsCount, err := c.Db.Count(&model.Audit{})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(iris.Map{
		"userCount":     userCount,
		"deviceCount":   deviceCount,
		"onlineCount":   onlineCount,
		"visitsCount":   visitsCount,
		"compatVersion": v2service.CompatSysinfoVersion,
	}, "ok")
}

func (c *DashboardController) GetDashboardLineCharts() mvc.Result {
	now := carbon.Now()
	startOfWeek := now.SetWeekStartsAt(carbon.Monday).StartOfWeek()
	endOfWeek := now.SetWeekStartsAt(carbon.Monday).EndOfWeek()

	startOfWeekString := startOfWeek.ToDateTimeString()
	endOfWeekString := endOfWeek.ToDateTimeString()

	type lineChartsData struct {
		Value int    `json:"value"`
		Date  string `json:"date"`
	}

	var userList = make([]lineChartsData, 0)
	err := c.Db.Table(&model.User{}).Select("count(*) as `value`, date(created_at) as `date`").Where("created_at between ? and ?", startOfWeekString, endOfWeekString).GroupBy("`date`").Find(&userList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	var userData = make(map[string]int)
	for _, r := range userList {
		userData[r.Date] = r.Value
	}

	var peerList = make([]lineChartsData, 0)
	err = c.Db.Table(&model.Peer{}).Select("count(*) as `value`, date(created_at) as `date`").Where("created_at between ? and ?", startOfWeekString, endOfWeekString).GroupBy("`date`").Find(&peerList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	var peerData = make(map[string]int)
	for _, r := range peerList {
		peerData[r.Date] = r.Value
	}

	var xDateLine []string
	var seriesUser []int
	var seriesPeer []int
	for i := startOfWeek.Day(); i <= endOfWeek.Day(); i++ {
		current := carbon.Parse(fmt.Sprintf("%d-%d-%d", startOfWeek.Year(), startOfWeek.Month(), i))
		s := current.ToDateString()
		xDateLine = append(xDateLine, s)
		seriesUser = append(seriesUser, userData[s])
		seriesPeer = append(seriesPeer, peerData[s])
	}

	return c.Success(iris.Map{
		"xAxis": xDateLine,
		"users": seriesUser,
		"peer":  seriesPeer,
	}, "ok")
}

func (c *DashboardController) GetDashboardPieCharts() mvc.Result {
	type pieChartsData struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	}
	pieMap := make([]pieChartsData, 0)
	err := c.Db.Table(&model.Peer{}).GroupBy("platform").Select("case when `platform` = '' then 'unknown' else `platform` end as `name`, count(*) as `value`").Find(&pieMap)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(pieMap, "ok")
}

func (c *DashboardController) GetDashboardServerConfig() mvc.Result {
	req := c.Ctx.Request()
	hostWithPort := req.Host
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}
	if forwardedProto := c.Ctx.GetHeader("X-Forwarded-Proto"); forwardedProto != "" {
		scheme = strings.TrimSpace(strings.Split(forwardedProto, ",")[0])
	}

	hostOnly := hostWithPort
	if h, _, err := net.SplitHostPort(hostWithPort); err == nil {
		hostOnly = h
	} else if idx := strings.LastIndex(hostWithPort, ":"); idx > -1 && !strings.Contains(hostWithPort, "]") {
		hostOnly = hostWithPort[:idx]
	}
	hostOnly = strings.Trim(hostOnly, "[]")

	apiServer := firstNonEmpty(
		os.Getenv("RUSTDESK_API_SERVER"),
		fmt.Sprintf("%s://%s", scheme, hostWithPort),
	)

	return c.Success(iris.Map{
		"idServer":    firstNonEmpty(os.Getenv("RUSTDESK_ID_SERVER"), hostOnly),
		"relayServer": firstNonEmpty(os.Getenv("RUSTDESK_RELAY_SERVER"), hostOnly),
		"apiServer":   apiServer,
		"key":         os.Getenv("RUSTDESK_KEY"),
	}, "ok")
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}
