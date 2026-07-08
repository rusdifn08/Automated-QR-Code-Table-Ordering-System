package monitor

import (
	"container/ring"
	"runtime"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type SystemMetrics struct {
	Uptime        string      `json:"uptime"`
	TotalRequests uint64      `json:"total_requests"`
	AverageLatency string      `json:"average_latency"`
	MemoryAllocated uint64    `json:"memory_allocated"`
	MemorySys      uint64      `json:"memory_sys"`
	RecentLogs    []string    `json:"recent_logs"`
}

var (
	startTime     = time.Now()
	totalRequests uint64
	totalLatency  time.Duration
	logsRing      *ring.Ring
	mu            sync.Mutex
)

func init() {
	logsRing = ring.New(50) // Store last 50 logs
}

func AddLog(msg string) {
	mu.Lock()
	defer mu.Unlock()
	logsRing.Value = time.Now().Format("15:04:05") + " | " + msg
	logsRing = logsRing.Next()
}

func MetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		
		err := c.Next()
		
		duration := time.Since(start)
		
		mu.Lock()
		totalRequests++
		totalLatency += duration
		mu.Unlock()

		if c.Path() != "/api/admin/system-metrics" && c.Path() != "/metrics" {
			AddLog(c.Method() + " " + c.Path() + " - " + duration.String())
		}
		return err
	}
}

func GetMetrics() SystemMetrics {
	mu.Lock()
	defer mu.Unlock()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	avgLat := "0ms"
	if totalRequests > 0 {
		avgLat = (totalLatency / time.Duration(totalRequests)).String()
	}

	var logs []string
	logsRing.Do(func(p interface{}) {
		if p != nil {
			logs = append(logs, p.(string))
		}
	})

	return SystemMetrics{
		Uptime:         time.Since(startTime).String(),
		TotalRequests:  totalRequests,
		AverageLatency: avgLat,
		MemoryAllocated: m.Alloc / 1024 / 1024, // in MB
		MemorySys:      m.Sys / 1024 / 1024,   // in MB
		RecentLogs:     logs,
	}
}
