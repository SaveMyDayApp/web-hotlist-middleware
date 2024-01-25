package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web-hotlist-middleware/pkg/statistics"
)

func NewRouter(handler *Handler) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/api/statistics", handler.GetStatistics)
	router.HandleFunc("/api/submit-statistics", handler.SubmitStatistics)

	return router
}

// 接口设计
// 获取统计数据API：

// URL: GET /api/statistics
// 请求参数：类型（type），时间范围（period），返回条目数量（limit，可选）。
// 返回：统计数据列表，包括标识符、计数和排名。
// 提交统计数据API：

// URL: POST /api/submit-statistics
// 请求体：统计类型（type），标识符（identifier），计数（count）。
// 返回：操作成功与否的状态和消息。
// 技术细节
// Redis数据结构：

// 使用散列表存储每个统计项的计数和最后更新时间。
// 采用日期为键的结构存储每天的统计数据，以支持时间范围查询。
// 性能和扩展性：

// 确保中间件能够处理大量并发请求。
// 设计易于扩展的代码结构，以便未来添加新的统计类型。

// URL: GET /api/statistics
// 请求参数：类型（type），时间范围（period），返回条目数量（limit，可选）。
// 返回：统计数据列表，包括标识符、计数和排名。
// 提交统计数据API：
// GetStatistics handles the GET /api/statistics endpoint
func (h *Handler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	// 从请求中提取参数
	query := r.URL.Query()
	typ := query.Get("type")
	period := query.Get("period")
	limitStr := query.Get("limit")

	// 将limit转换为整数
	var limit int
	var err error
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}

	// 调用统计服务
	stats, err := h.StatsService.GetStatistics(typ, period, limit)
	if err != nil {
		http.Error(w, "Error getting statistics", http.StatusInternalServerError)
		return
	}

	// 将结果写入响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Result{Code: 200, Data: stats})
}

// SubmitStatistics handles the POST /api/submit-statistics endpoint
func (h *Handler) SubmitStatistics(w http.ResponseWriter, r *http.Request) {
	// 从请求体中提取参数
	var stat statistics.Statistic
	err := json.NewDecoder(r.Body).Decode(&stat)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Result{Code: 400, Data: "Error parsing request body"})
		return
	}

	// 调用统计服务
	err = h.StatsService.SubmitStatistic(stat.Type, stat.Identifier, stat.Count)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Result{Code: 501, Data: err})
		return
	}

	// 将结果写入响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Result{Code: 200, Data: "ok"})
}
