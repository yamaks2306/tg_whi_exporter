package collector

import (
	"github.com/yamaks2306/tg_whi_exporter/config"
	"github.com/yamaks2306/tg_whi_exporter/pg"
	"github.com/yamaks2306/tg_whi_exporter/telegram"

	"github.com/prometheus/client_golang/prometheus"
)

type tgInfoCollector struct {
	telegram_pending_updates *prometheus.Desc
	config                   config.Config
}

func NewTgInfoCollector(config config.Config) *tgInfoCollector {
	return &tgInfoCollector{
		telegram_pending_updates: prometheus.NewDesc(
			"pending_telegram_webhook_updates_count",
			"Telegram webhook pending updates count",
			[]string{"client", "mis_type"}, nil,
		),
		config: config,
	}
}

func (collector *tgInfoCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.telegram_pending_updates
}

func (collector *tgInfoCollector) Collect(ch chan<- prometheus.Metric) {

	telegram_webhook_info, err := telegram.GetWebhookInfo(collector.config.TgConfig.TelegramToken)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(
			prometheus.NewDesc("loyalmed_exporter_error",
				"Error getting Telegram webhook info", nil, nil),
			err)
	}
	pending_updates_count := telegram_webhook_info.Result.Pending_update_count

	ch <- prometheus.MustNewConstMetric(collector.telegram_pending_updates, prometheus.GaugeValue, float64(pending_updates_count), collector.config.TgConfig.ServerURL, collector.config.TgConfig.MisType)

}

type pgCollector struct {
	postgres_users_count *prometheus.Desc
	config               config.Config
}

func NewPgCollector(config config.Config) *pgCollector {
	return &pgCollector{
		postgres_users_count: prometheus.NewDesc(
			"postgres_users_count",
			"Postgres users count",
			[]string{"client", "mis_type"}, nil,
		),
		config: config,
	}
}

func (collector *pgCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.postgres_users_count
}

func (collector *pgCollector) Collect(ch chan<- prometheus.Metric) {

	pg_users_count, err := pg.GetDbUsersCount(collector.config.PgDockerConfig, collector.config.PgConfig)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(
			prometheus.NewDesc("loyalmed_exporter_error",
				"Error getting postgres metrics", nil, nil),
			err)
	}

	ch <- prometheus.MustNewConstMetric(collector.postgres_users_count, prometheus.GaugeValue, float64(pg_users_count), collector.config.TgConfig.ServerURL, collector.config.TgConfig.MisType)
}
