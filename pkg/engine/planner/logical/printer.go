package logical

import (
	"fmt"
	"io"

	"example.com/acme/logstore/v3/pkg/engine/planner/internal/tree"
)

func WriteMermaidFormat(w io.Writer, p *Plan) {
	var t treeFormatter
	for _, inst := range p.Instructions {
		switch inst := inst.(type) {
		case *Return:
			node := t.convert(inst.Value)
			printer := tree.NewMermaid(w)
			_ = printer.Write(node)

			fmt.Fprint(w, "\n\n")
		}
	}
}
