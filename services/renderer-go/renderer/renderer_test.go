package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Render_Header(t *testing.T) {
	src := `# h1
## h2
### h3
#### h4`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, `<h1>h1</h1>
<h2>h2</h2>
<h3>h3</h3>
<h4>h4</h4>
`, html)
}

func Test_Render_List(t *testing.T) {
	src := `- list1
- list2`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, `<ul>
<li>list1</li>
<li>list2</li>
</ul>
`, html)
}

func Test_Render_Link(t *testing.T) {
	src := `[google.com](https:google.com)`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, `<p><a href="https:google.com">google.com</a></p>
`, html)
}
