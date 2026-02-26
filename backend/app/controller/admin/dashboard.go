package admin

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	v2service "rustdesk-api-server-pro/internal/service"
	"strings"
	"time"

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

	idServer, idServerSource := resolveConfigValue(os.Getenv("RUSTDESK_ID_SERVER"), hostOnly)
	relayServer, relayServerSource := resolveConfigValue(os.Getenv("RUSTDESK_RELAY_SERVER"), hostOnly)
	apiServer, apiServerSource := resolveConfigValue(
		os.Getenv("RUSTDESK_API_SERVER"),
		fmt.Sprintf("%s://%s", scheme, hostWithPort),
	)
	key, keySource := resolveConfigValue(os.Getenv("RUSTDESK_KEY"), "")

	return c.Success(iris.Map{
		"idServer":    idServer,
		"relayServer": relayServer,
		"apiServer":   apiServer,
		"key":         key,
		"sources": iris.Map{
			"idServer":    idServerSource,
			"relayServer": relayServerSource,
			"apiServer":   apiServerSource,
			"key":         keySource,
		},
	}, "ok")
}

func (c *DashboardController) GetDashboardServerConnectivity() mvc.Result {
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

	idServer, _ := resolveConfigValue(os.Getenv("RUSTDESK_ID_SERVER"), hostOnly)
	relayServer, _ := resolveConfigValue(os.Getenv("RUSTDESK_RELAY_SERVER"), hostOnly)
	apiServer, _ := resolveConfigValue(os.Getenv("RUSTDESK_API_SERVER"), fmt.Sprintf("%s://%s", scheme, hostWithPort))
	key, _ := resolveConfigValue(os.Getenv("RUSTDESK_KEY"), "")

	return c.Success(iris.Map{
		"idServer":    probeTCPServer(idServer, "21116"),
		"relayServer": probeTCPServer(relayServer, "21117"),
		"apiServer":   probeHTTPServer(apiServer),
		"key":         probeKeyConfig(key),
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

func resolveConfigValue(envValue, fallbackValue string) (string, string) {
	envValue = strings.TrimSpace(envValue)
	if envValue != "" {
		return envValue, "env"
	}

	fallbackValue = strings.TrimSpace(fallbackValue)
	if fallbackValue != "" {
		return fallbackValue, "inferred"
	}

	return "", "empty"
}

func probeTCPServer(value, defaultPort string) iris.Map {
	value = strings.TrimSpace(value)
	if value == "" {
		return iris.Map{"status": "skip", "message": "empty", "target": ""}
	}

	target, err := tcpDialTarget(value, defaultPort)
	if err != nil {
		return iris.Map{"status": "error", "message": err.Error(), "target": value}
	}

	conn, err := net.DialTimeout("tcp", target, 2*time.Second)
	if err != nil {
		return iris.Map{"status": "error", "message": err.Error(), "target": target}
	}
	_ = conn.Close()

	return iris.Map{"status": "ok", "message": "connected", "target": target}
}

func tcpDialTarget(value, defaultPort string) (string, error) {
	raw := strings.TrimSpace(value)
	if raw == "" {
		return "", fmt.Errorf("empty target")
	}

	if strings.Contains(raw, "://") {
		u, err := url.Parse(raw)
		if err != nil {
			return "", err
		}
		raw = u.Host
	}
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", fmt.Errorf("empty host")
	}

	if _, _, err := net.SplitHostPort(raw); err == nil {
		return raw, nil
	}

	host := strings.Trim(raw, "[]")
	if host == "" {
		return "", fmt.Errorf("empty host")
	}
	return net.JoinHostPort(host, defaultPort), nil
}

func probeHTTPServer(value string) iris.Map {
	target := strings.TrimSpace(value)
	if target == "" {
		return iris.Map{"status": "skip", "message": "empty", "target": ""}
	}

	if !strings.Contains(target, "://") {
		target = "http://" + target
	}
	u, err := url.Parse(target)
	if err != nil || u.Scheme == "" || u.Host == "" {
		if err == nil {
			err = fmt.Errorf("invalid url")
		}
		return iris.Map{"status": "error", "message": err.Error(), "target": target}
	}

	client := &http.Client{Timeout: 3 * time.Second}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	resp, err := client.Do(req)
	if err != nil {
		return iris.Map{"status": "error", "message": err.Error(), "target": target}
	}
	_ = resp.Body.Close()

	return iris.Map{
		"status":  "ok",
		"message": fmt.Sprintf("http %d", resp.StatusCode),
		"target":  target,
	}
}

func probeKeyConfig(value string) iris.Map {
	if strings.TrimSpace(value) == "" {
		return iris.Map{"status": "error", "message": "key is empty", "target": ""}
	}
	return iris.Map{"status": "ok", "message": "configured", "target": ""}
}
