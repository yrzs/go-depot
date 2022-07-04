package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

/**
jaeger client
*/
func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{ //jaeger client 的配置项
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{ //固定采样、对所有数据都进行采样
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{ //是否启用LoggingReporter、刷新缓冲区的频率、上报的Agent地址
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer) //设置全局的Tracer对象，根据实际情况设置即可。因为通常会统一使用一套追踪系统，因此该语句常常会被使用。
	return tracer, closer, nil
}
