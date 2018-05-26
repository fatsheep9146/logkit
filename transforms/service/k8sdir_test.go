package service

import (
	"testing"

	"github.com/qiniu/logkit/transforms"
	. "github.com/qiniu/logkit/utils/models"

	"github.com/stretchr/testify/assert"
)

func TestK8sDirTransformer(t *testing.T) {
	ktag := &K8sDir{
		SourceFileKey: "sourcetag",
	}
	data, err := ktag.Transform([]Data{{"sourcetag": "/applog/kube-system_pod1/abc.log", "abc": "x1 y2"}})
	assert.NoError(t, err)
	exp := []Data{
		{"sourcetag": "/applog/kube-system_pod1/abc.log", PodName: "pod1", Namespace: "kube-system", "abc": "x1 y2"}}
	assert.Equal(t, exp, data)

	assert.Equal(t, ktag.Stage(), transforms.StageAfterParser)
}
