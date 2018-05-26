package service

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/qiniu/logkit/transforms"
	. "github.com/qiniu/logkit/utils/models"
)

const (
	K8sDirType = "k8sdir"
	PodName    = "pod_name"
	Namespace  = "namespace"
)

type K8sDir struct {
	SourceFileKey string `json:"sourcefilefield"`
	stats         StatsInfo
}

func (g *K8sDir) RawTransform(datas []string) ([]string, error) {
	return datas, errors.New("k8sdir transformer not support rawTransform")
}

func (g *K8sDir) Transform(datas []Data) ([]Data, error) {
	var err, ferr error
	errnums := 0
	for i := range datas {
		val, ok := datas[i][g.SourceFileKey]
		if !ok {
			errnums++
			err = fmt.Errorf("transform key %v not exist in data", g.SourceFileKey)
			continue
		}
		strval, ok := val.(string)
		if !ok {
			errnums++
			err = fmt.Errorf("transform key %v data type is not string", g.SourceFileKey)
			continue
		}
		strval = filepath.Dir(strval)
		splits := strings.Split(strval, "/")
		parseString := splits[len(splits)-1]
		splits = strings.Split(parseString, "_")

		datas[i][Namespace] = splits[0]
		datas[i][PodName] = splits[1]
	}
	if err != nil {
		g.stats.LastError = err.Error()
		ferr = fmt.Errorf("find total %v erorrs in transform k8sdir, last error info is %v", errnums, err)
	}
	g.stats.Errors += int64(errnums)
	g.stats.Success += int64(len(datas) - errnums)
	return datas, ferr
}

func (g *K8sDir) Description() string {
	//return "k8stag will get kubernetes tags from sourcefile name"
	return "从目录中获取 kubernetes namespace，pod_name 信息"
}

func (g *K8sDir) Type() string {
	return K8sDirType
}

func (g *K8sDir) SampleConfig() string {
	return `{
		"type":"k8sdir",
		"sourcefilefield":"datasource"
	}`
}

func (g *K8sDir) ConfigOptions() []Option {
	return []Option{
		transforms.KeyFieldName,
	}
}

func (g *K8sDir) Stage() string {
	return transforms.StageAfterParser
}

func (g *K8sDir) Stats() StatsInfo {
	return g.stats
}

func init() {
	transforms.Add(K8sDirType, func() transforms.Transformer {
		return &K8sDir{}
	})
}
