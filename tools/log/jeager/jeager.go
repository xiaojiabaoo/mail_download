/**
  @author: $(USER)
  @data:$(DATE)
  @note:
**/
package jeager

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

//分布式追踪系统
type Jeager struct {
	name   string
	addr   string
	closer io.Closer
}

func NewJeager(name, addr string) *Jeager {
	return &Jeager{
		name: name,
		addr: addr,
	}
}

func (j *Jeager) Open() error {
	// 根据配置初始化Tracer 返回Closer
	tracer, closer, err := (&config.Configuration{
		ServiceName: j.name,
		Disabled:    false,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			// param的值在0到1之间，设置为1则将所有的Operation输出到Reporter
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			// 请注意，如果logSpans没开启就一定没有结果
			LogSpans: true,
			// 请一定要注意，LocalAgentHostPort填错了就一定没有结果
			LocalAgentHostPort: j.addr,
		},
	}).NewTracer()
	if err != nil {
		return err
	}
	j.closer = closer
	// 设置全局Tracer - 如果不设置将会导致上下文无法生成正确的Span
	opentracing.SetGlobalTracer(tracer)
	return nil
}

func (j *Jeager) Close() error {
	if j.closer != nil {
		return j.closer.Close()
	}
	return nil
}
