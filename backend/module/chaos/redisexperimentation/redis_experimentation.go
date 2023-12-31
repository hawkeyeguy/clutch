package redisexperimentation

// <!-- START clutchdoc -->
// description: Chaos Experimentation Framework - Supports Redis specific experiments.
// <!-- END clutchdoc -->

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"

	experimentationv1 "github.com/lyft/clutch/backend/api/chaos/experimentation/v1"
	redisexperimentationv1 "github.com/lyft/clutch/backend/api/chaos/redisexperimentation/v1"
	configv1 "github.com/lyft/clutch/backend/api/config/module/chaos/redisexperimentation/v1"
	"github.com/lyft/clutch/backend/module"
	"github.com/lyft/clutch/backend/module/chaos/experimentation/xds"
	redisexperimentationxds "github.com/lyft/clutch/backend/module/chaos/redisexperimentation/xds"
	"github.com/lyft/clutch/backend/service"
	"github.com/lyft/clutch/backend/service/chaos/experimentation/experimentstore"
)

const (
	Name = "clutch.module.chaos.redisexperimentation"
)

type Service struct {
	storer experimentstore.Storer
}

// New instantiates a Service object.
func New(untypedConfig *any.Any, logger *zap.Logger, scope tally.Scope) (module.Module, error) {
	config := &configv1.Config{}
	if err := untypedConfig.UnmarshalTo(config); err != nil {
		return nil, err
	}

	store, ok := service.Registry[experimentstore.Name]
	if !ok {
		return nil, errors.New("could not find experiment store service")
	}

	storer, ok := store.(experimentstore.Storer)
	if !ok {
		return nil, errors.New("service was not the correct type")
	}

	xds.RTDSGeneratorsByTypeUrl[xds.TypeUrl(&redisexperimentationv1.FaultConfig{})] = &redisexperimentationxds.RTDSFaultsGenerator{
		FaultRuntimePrefix: config.FaultRuntimePrefix,
	}

	return &Service{
		storer: storer,
	}, nil
}

func (s *Service) Register(r module.Registrar) error {
	transformation := experimentstore.Transformation{ConfigTypeUrl: "type.googleapis.com/clutch.chaos.redisexperimentation.v1.FaultConfig", RunTransform: s.transform}
	return s.storer.RegisterTransformation(transformation)
}

func (s *Service) transform(_ *experimentstore.ExperimentRun, config *experimentstore.ExperimentConfig) ([]*experimentationv1.Property, error) {
	var experimentConfig = redisexperimentationv1.FaultConfig{}
	if err := config.Config.UnmarshalTo(&experimentConfig); err != nil {
		return []*experimentationv1.Property{}, err
	}

	faultsDescription, err := experimentConfigToFaultString(&experimentConfig)
	if err != nil {
		return nil, err
	}

	var downstream, upstream string
	downstream = experimentConfig.GetFaultTargeting().GetDownstreamCluster().GetName()
	upstream = experimentConfig.GetFaultTargeting().GetUpstreamCluster().GetName()

	return []*experimentationv1.Property{
		{
			Id:    "type",
			Label: "Type",
			Value: &experimentationv1.Property_StringValue{StringValue: "Redis"},
		},
		{
			Id:    "target",
			Label: "Target",
			Value: &experimentationv1.Property_StringValue{StringValue: fmt.Sprintf("%s ➡️ %s", downstream, upstream)},
		},
		{
			Id:    "fault_types",
			Label: "Fault Types",
			Value: &experimentationv1.Property_StringValue{StringValue: faultsDescription},
		},
	}, nil
}

func experimentConfigToFaultString(experiment *redisexperimentationv1.FaultConfig) (string, error) {
	if experiment == nil {
		return "", errors.New("experiment is nil")
	}

	switch experiment.GetFault().(type) {
	case *redisexperimentationv1.FaultConfig_ErrorFault:
		return "Error", nil
	case *redisexperimentationv1.FaultConfig_LatencyFault:
		return "Delay", nil
	default:
		return "", fmt.Errorf("unexpected fault type %v", experiment.GetFault())
	}
}
