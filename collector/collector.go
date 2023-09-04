package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yamaks2306/tg_whi_exporter/config"
	"github.com/yamaks2306/tg_whi_exporter/pg"
	"github.com/yamaks2306/tg_whi_exporter/telegram"
)

type tgInfoCollector struct {
	telegram_pending_updates *prometheus.Desc
	postgres_users_count     *prometheus.Desc
}

func NewTgInfoCollector() *tgInfoCollector {
	return &tgInfoCollector{
		telegram_pending_updates: prometheus.NewDesc(
			"pending_telegram_webhook_updates_count",
			"Telegram webhook pending updates count",
			[]string{"client", "mis_type"}, nil,
		),
		postgres_users_count: prometheus.NewDesc(
			"postgres_users_count",
			"Postgres users count",
			[]string{"client", "mis_type"}, nil,
		),
	}
}

func (collector *tgInfoCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.telegram_pending_updates
	ch <- collector.postgres_users_count
}

func (collector *tgInfoCollector) Collect(ch chan<- prometheus.Metric) {

	conf := config.New()

	pending_updates_count := telegram.GetWebhookInfo(conf.TgConfig.TelegramToken).Result.Pending_update_count

	pg_users_count := pg.GetDbUsersCount(conf.PgDockerConfig, conf.PgConfig)

	ch <- prometheus.MustNewConstMetric(collector.telegram_pending_updates, prometheus.GaugeValue, float64(pending_updates_count), conf.TgConfig.ServerURL, conf.TgConfig.MisType)
	ch <- prometheus.MustNewConstMetric(collector.postgres_users_count, prometheus.GaugeValue, float64(pg_users_count), conf.TgConfig.ServerURL, conf.TgConfig.MisType)
}
