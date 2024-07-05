package parserx

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

func Test_NewSpecParser(t *testing.T) {
	spa, err := NewSpecParser().ParseApi("./_testdata/test.api")
	assert.Equal(t, nil, err)

	t.Log(jsonconv.ObjectToJsonIndent(spa.Service))

	swa, err := NewSwaggerParser().ParseApi("./_testdata/test.json")
	assert.Equal(t, nil, err)

	t.Log(jsonconv.ObjectToJsonIndent(swa.Service))
}
