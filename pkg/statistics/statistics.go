package statistics

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Statistic struct {
	Type        string
	Identifier  string
	Count       int
	LastUpdated time.Time
}

type StatisticsService struct {
	// This would be your connection to Redis or any other database you're using
	db *redis.Client
}

func NewStatisticsService(db *redis.Client) *StatisticsService {
	return &StatisticsService{db: db}
}

func (s *StatisticsService) GetStatistics(statType string, period string, limit int) ([]Statistic, error) {
	// 根据类型和时间范围确定要查询的键
	var key string
	switch period {
	case "day":
		key = statType + ":" + time.Now().Format("2006-01-02")
	case "week":
		key = statType + ":" + time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	case "month":
		key = statType + ":" + time.Now().AddDate(0, -1, 0).Format("2006-01-02")
	case "all":
		key = statType
	default:
		return nil, errors.New("Invalid period parameter")
	}

	// 从Redis中获取最高分的成员
	result, err := s.db.ZRevRangeWithScores(context.Background(), key, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	// 将结果转换为Statistic结构体的切片
	stats := make([]Statistic, len(result))
	for i, member := range result {
		// 从散列表中获取统计项的最后更新时间
		lastUpdatedStr, err := s.db.HGet(context.Background(), statType+":"+member.Member.(string), "lastUpdated").Result()
		if err != nil {
			return nil, err
		}

		lastUpdated, err := strconv.ParseInt(lastUpdatedStr, 10, 64)
		if err != nil {
			return nil, err
		}

		// 将Unix时间戳转换为time.Time
		lastUpdatedTime := time.Unix(lastUpdated, 0)

		stats[i] = Statistic{
			Type:        statType,
			Identifier:  member.Member.(string),
			Count:       int(member.Score),
			LastUpdated: lastUpdatedTime,
		}
	}

	return stats, nil
}

func (s *StatisticsService) SubmitStatistic(statType string, identifier string, count int) error {
	// 获取当前日期
	today := time.Now().Format("2006-01-02")

	// 使用事务来确保所有操作都成功
	err := s.db.Watch(context.Background(), func(tx *redis.Tx) error {
		// 在散列表中增加统计项的计数
		_, err := tx.HIncrBy(context.Background(), statType+":"+identifier, "count", int64(count)).Result()
		if err != nil {
			return err
		}

		// 更新统计项的最后更新时间
		_, err = tx.HSet(context.Background(), statType+":"+identifier, "lastUpdated", time.Now().Unix()).Result()
		if err != nil {
			return err
		}

		// 在日期键的sorted set中增加统计项的计数
		_, err = tx.ZIncrBy(context.Background(), statType+":"+today, float64(count), identifier).Result()
		if err != nil {
			return err
		}

		return nil
	}, statType+":"+identifier, statType+":"+today)

	return err
}
