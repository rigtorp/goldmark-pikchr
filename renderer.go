package pikchr

import (
	"bytes"
	"fmt"

	"github.com/rigtorp/go-pikchr"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type Renderer struct{}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindBlock, r.Render)
}

func (r *Renderer) Render(w util.BufWriter, src []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*Block)
	if !entering {
		return ast.WalkContinue, nil
	}

	b := bytes.Buffer{}
	lines := n.Lines()
	for i := 0; i < lines.Len(); i++ {
		line := lines.At(i)
		b.Write(line.Value(src))
	}

	if b.Len() == 0 {
		return ast.WalkContinue, nil
	}

	out := bytes.Buffer{}
	if err := pikchr.Render(&b, &out); err != nil {
		_, err := fmt.Fprintf(w, "<pre>\n%s</pre>\n", out.String())
		return ast.WalkContinue, err
	}

	_, err := out.WriteTo(w)

	return ast.WalkContinue, err
}
