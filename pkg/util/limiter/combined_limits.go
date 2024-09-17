package limiter

import (
	bloombuilder "github.com/agardiman/loki/v3/pkg/bloombuild/builder"
	bloomplanner "github.com/agardiman/loki/v3/pkg/bloombuild/planner"
	"github.com/agardiman/loki/v3/pkg/bloomgateway"
	"github.com/agardiman/loki/v3/pkg/compactor"
	"github.com/agardiman/loki/v3/pkg/distributor"
	"github.com/agardiman/loki/v3/pkg/indexgateway"
	"github.com/agardiman/loki/v3/pkg/ingester"
	querier_limits "github.com/agardiman/loki/v3/pkg/querier/limits"
	queryrange_limits "github.com/agardiman/loki/v3/pkg/querier/queryrange/limits"
	"github.com/agardiman/loki/v3/pkg/ruler"
	scheduler_limits "github.com/agardiman/loki/v3/pkg/scheduler/limits"
	"github.com/agardiman/loki/v3/pkg/storage"
)

type CombinedLimits interface {
	compactor.Limits
	distributor.Limits
	ingester.Limits
	querier_limits.Limits
	queryrange_limits.Limits
	ruler.RulesLimits
	scheduler_limits.Limits
	storage.StoreLimits
	indexgateway.Limits
	bloomgateway.Limits
	bloomplanner.Limits
	bloombuilder.Limits
}
