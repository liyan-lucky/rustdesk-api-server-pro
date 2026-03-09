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
	"sync"
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
	for current := startOfWeek; current.Lte(endOfWeek); current = current.AddDay() {
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

	type rawPieChartsData struct {
		Value int    `xorm:"value"`
		Name  string `xorm:"name"`
	}

	rawRows := make([]rawPieChartsData, 0)
	err := c.Db.SQL(`
select
  case
    when trim(ifnull(d.os, '')) <> '' then trim(d.os)
    when trim(ifnull(ato.device_os, '')) <> '' then trim(ato.device_os)
    when trim(ifnull(po.platform, '')) <> '' then trim(po.platform)
    else 'unknown'
  end as name,
  count(*) as value
from device d
left join (
  select t.rustdesk_id, t.device_os
  from auth_token t
  where ifnull(t.device_os, '') <> ''
    and t.id in (
      select max(id)
      from auth_token
      where ifnull(device_os, '') <> ''
      group by rustdesk_id
    )
) ato on ato.rustdesk_id = d.rustdesk_id
left join (
  select p.rustdesk_id, p.platform
  from peer p
  where ifnull(p.platform, '') <> ''
    and p.id in (
      select max(id)
      from peer
      where ifnull(platform, '') <> ''
      group by rustdesk_id
    )
) po on po.rustdesk_id = d.rustdesk_id
group by name
`).Find(&rawRows)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	merged := make(map[string]int)
	for _, r := range rawRows {
		key := normalizeDashboardOSName(r.Name)
		merged[key] += r.Value
	}

	pieMap := make([]pieChartsData, 0, len(merged))
	for name, value := range merged {
		pieMap = append(pieMap, pieChartsData{Name: name, Value: value})
	}

	return c.Success(pieMap, "ok")
}

func normalizeDashboardOSName(raw string) string {
	s := strings.ToLower(strings.TrimSpace(raw))
	if s == "" || s == "unknown" || s == "unkn" || s == "n/a" || s == "na" || s == "none" || s == "null" {
		return "unknown"
	}
	if strings.Contains(s, "mac") || strings.Contains(s, "darwin") || strings.Contains(s, "osx") {
		return "macOS"
	}
	if strings.Contains(s, "chrome") {
		return "ChromeOS"
	}
	if strings.Contains(s, "win") {
		return "Windows"
	}
	if strings.Contains(s, "linux") ||
		strings.Contains(s, "ubuntu") ||
		strings.Contains(s, "debian") ||
		strings.Contains(s, "centos") ||
		strings.Contains(s, "fedora") ||
		strings.Contains(s, "arch") ||
		strings.Contains(s, "openwrt") ||
		strings.Contains(s, "alpine") ||
		strings.Contains(s, "manjaro") {
		return "Linux"
	}
	if strings.Contains(s, "android") {
		return "Android"
	}
	if strings.Contains(s, "ios") || strings.Contains(s, "iphone") || strings.Contains(s, "ipad") {
		return "iOS"
	}
	return strings.TrimSpace(raw)
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

	hostOnly := extractHostOnly(hostWithPort)
	inferredAPIServer := inferAPIServerURL(scheme, hostWithPort, c.getConfiguredHTTPPort())

	idServer, idServerSource := resolveConfigValue(os.Getenv("RUSTDESK_ID_SERVER"), hostOnly)
	relayServer, relayServerSource := resolveConfigValue(os.Getenv("RUSTDESK_RELAY_SERVER"), hostOnly)
	apiServer, apiServerSource := resolveConfigValue(
		os.Getenv("RUSTDESK_API_SERVER"),
		inferredAPIServer,
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

	hostOnly := extractHostOnly(hostWithPort)
	inferredAPIServer := inferAPIServerURL(scheme, hostWithPort, c.getConfiguredHTTPPort())

	idServer, _ := resolveConfigValue(os.Getenv("RUSTDESK_ID_SERVER"), hostOnly)
	relayServer, _ := resolveConfigValue(os.Getenv("RUSTDESK_RELAY_SERVER"), hostOnly)
	apiServer, _ := resolveConfigValue(os.Getenv("RUSTDESK_API_SERVER"), inferredAPIServer)
	key, _ := resolveConfigValue(os.Getenv("RUSTDESK_KEY"), "")
	targetKey := strings.TrimSpace(c.Ctx.URLParamDefault("target", ""))

	if targetKey != "" {
		switch targetKey {
		case "idServer":
			return c.Success(iris.Map{"idServer": probeTCPServer(idServer, "21116")}, "ok")
		case "relayServer":
			return c.Success(iris.Map{"relayServer": probeTCPServer(relayServer, "21117")}, "ok")
		case "apiServer":
			return c.Success(iris.Map{"apiServer": probeHTTPServer(apiServer)}, "ok")
		case "key":
			return c.Success(iris.Map{"key": probeKeyConfig(key)}, "ok")
		default:
			return c.Error(nil, "invalid connectivity target")
		}
	}

	type probeResult struct {
		idServer    iris.Map
		relayServer iris.Map
		apiServer   iris.Map
		key         iris.Map
	}
	result := &probeResult{}
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		result.idServer = probeTCPServer(idServer, "21116")
	}()
	go func() {
		defer wg.Done()
		result.relayServer = probeTCPServer(relayServer, "21117")
	}()
	go func() {
		defer wg.Done()
		result.apiServer = probeHTTPServer(apiServer)
	}()
	go func() {
		defer wg.Done()
		result.key = probeKeyConfig(key)
	}()
	wg.Wait()

	return c.Success(iris.Map{
		"idServer":    result.idServer,
		"relayServer": result.relayServer,
		"apiServer":   result.apiServer,
		"key":         result.key,
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

func (c *DashboardController) getConfiguredHTTPPort() string {
	if c == nil || c.Cfg == nil || c.Cfg.HttpConfig == nil {
		return ""
	}
	return strings.TrimSpace(c.Cfg.HttpConfig.Port)
}

func extractHostOnly(hostWithPort string) string {
	hostWithPort = strings.TrimSpace(hostWithPort)
	if hostWithPort == "" {
		return ""
	}

	if h, _, err := net.SplitHostPort(hostWithPort); err == nil {
		return strings.Trim(h, "[]")
	}

	if strings.HasPrefix(hostWithPort, "[") && strings.Contains(hostWithPort, "]") {
		end := strings.Index(hostWithPort, "]")
		return strings.Trim(hostWithPort[1:end], "[]")
	}

	if idx := strings.LastIndex(hostWithPort, ":"); idx > -1 && !strings.Contains(hostWithPort, "]") {
		portPart := strings.TrimSpace(hostWithPort[idx+1:])
		if isDigits(portPart) {
			return strings.Trim(hostWithPort[:idx], "[]")
		}
	}

	return strings.Trim(hostWithPort, "[]")
}

func extractPort(hostWithPort string) string {
	hostWithPort = strings.TrimSpace(hostWithPort)
	if hostWithPort == "" {
		return ""
	}

	if _, p, err := net.SplitHostPort(hostWithPort); err == nil {
		return p
	}

	if strings.HasPrefix(hostWithPort, "[") && strings.Contains(hostWithPort, "]:") {
		parts := strings.SplitN(hostWithPort, "]:", 2)
		if len(parts) == 2 {
			p := strings.TrimSpace(parts[1])
			if isDigits(p) {
				return p
			}
		}
	}

	if idx := strings.LastIndex(hostWithPort, ":"); idx > -1 && !strings.Contains(hostWithPort, "]") {
		p := strings.TrimSpace(hostWithPort[idx+1:])
		if isDigits(p) {
			return p
		}
	}

	return ""
}

func normalizePort(port string) string {
	port = strings.TrimSpace(port)
	if port == "" {
		return ""
	}

	if strings.HasPrefix(port, ":") {
		p := strings.TrimPrefix(port, ":")
		if isDigits(p) {
			return p
		}
	}

	if isDigits(port) {
		return port
	}

	if _, p, err := net.SplitHostPort(port); err == nil && isDigits(p) {
		return p
	}

	if idx := strings.LastIndex(port, ":"); idx > -1 {
		p := strings.TrimSpace(port[idx+1:])
		if isDigits(p) {
			return p
		}
	}

	return ""
}

func inferAPIServerURL(scheme, hostWithPort, configuredPort string) string {
	scheme = strings.TrimSpace(scheme)
	if scheme == "" {
		scheme = "http"
	}

	hostOnly := extractHostOnly(hostWithPort)
	if hostOnly == "" {
		return ""
	}

	port := firstNonEmpty(extractPort(hostWithPort), normalizePort(configuredPort))
	if port != "" {
		return fmt.Sprintf("%s://%s", scheme, net.JoinHostPort(hostOnly, port))
	}
	return fmt.Sprintf("%s://%s", scheme, hostOnly)
}

func isDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
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
	start := time.Now()
	value = strings.TrimSpace(value)
	if value == "" {
		return iris.Map{"status": "skip", "message": "empty", "target": "", "durationMs": 0}
	}

	target, err := tcpDialTarget(value, defaultPort)
	if err != nil {
		return iris.Map{"status": "error", "message": err.Error(), "target": value, "durationMs": time.Since(start).Milliseconds()}
	}

	conn, err := net.DialTimeout("tcp", target, 2*time.Second)
	if err != nil {
		return iris.Map{"status": "error", "message": err.Error(), "target": target, "durationMs": time.Since(start).Milliseconds()}
	}
	_ = conn.Close()

	return iris.Map{"status": "ok", "message": "connected", "target": target, "durationMs": time.Since(start).Milliseconds()}
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
	start := time.Now()
	target := strings.TrimSpace(value)
	if target == "" {
		return iris.Map{"status": "skip", "message": "empty", "target": "", "durationMs": 0}
	}

	if !strings.Contains(target, "://") {
		target = "http://" + target
	}
	u, err := url.Parse(target)
	if err != nil || u.Scheme == "" || u.Host == "" {
		if err == nil {
			err = fmt.Errorf("invalid url")
		}
		return iris.Map{"status": "error", "message": err.Error(), "target": target, "durationMs": time.Since(start).Milliseconds()}
	}

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := doHTTPProbeRequest(client, target, http.MethodHead)
	if err == nil && (resp.StatusCode == http.StatusMethodNotAllowed || resp.StatusCode == http.StatusNotImplemented) {
		_ = resp.Body.Close()
		resp, err = doHTTPProbeRequest(client, target, http.MethodGet)
	}
	if err != nil {
		return iris.Map{"status": "error", "message": err.Error(), "target": target, "durationMs": time.Since(start).Milliseconds()}
	}
	_ = resp.Body.Close()

	return iris.Map{
		"status":  "ok",
		"message": fmt.Sprintf("http %d", resp.StatusCode),
		"target":  target,
		"durationMs": time.Since(start).Milliseconds(),
	}
}

func doHTTPProbeRequest(client *http.Client, target, method string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, method, target, nil)
	return client.Do(req)
}

func probeKeyConfig(value string) iris.Map {
	if strings.TrimSpace(value) == "" {
		return iris.Map{"status": "error", "message": "key is empty", "target": "", "durationMs": 0}
	}
	return iris.Map{"status": "ok", "message": "configured", "target": "", "durationMs": 0}
}
