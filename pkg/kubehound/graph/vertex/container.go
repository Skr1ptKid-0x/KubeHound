package vertex

import (
	"github.com/DataDog/KubeHound/pkg/kubehound/graph/types"
	"github.com/DataDog/KubeHound/pkg/kubehound/models/graph"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
)

const (
	ContainerLabel = "Container"
)

var _ Builder = (*Container)(nil)

type Container struct {
}

func (v Container) Label() string {
	return ContainerLabel
}

func (v Container) BatchSize() int {
	return DefaultBatchSize
}

func (v Container) Traversal() Traversal {
	return func(source *gremlingo.GraphTraversalSource, inserts []types.TraversalInput) *gremlingo.GraphTraversal {
		g := source.GetGraphTraversal()

		for _, insert := range inserts {
			i := insert.(*graph.Container)
			g = g.AddV(v.Label()).
				Property("class", v.Label()). // labels are not indexed - use a mirror property
				Property("storeID", i.StoreID).
				Property("name", i.Name).
				Property("image", i.Image).
				Property("privileged", i.Privileged).
				Property("privesc", i.PrivEsc).
				Property("hostPid", i.HostPID).
				Property("hostPath", i.HostPath).
				Property("hostNetwork", i.HostNetwork).
				Property("runAsUser", i.RunAsUser).
				Property("pod", i.Pod).
				Property("node", i.Node).
				Property("compromised", int(i.Compromised)).
				Property("critical", i.Critical)

			for _, cmd := range i.Command {
				g = g.Property(gremlingo.Cardinality.Set, "command", cmd)
			}
			for _, arg := range i.Args {
				g = g.Property(gremlingo.Cardinality.Set, "args", arg)
			}
			for _, cap := range i.Capabilities {
				g = g.Property(gremlingo.Cardinality.Set, "capabilities", cap)
			}
			for _, port := range i.Ports {
				g = g.Property(gremlingo.Cardinality.Set, "ports", port)
			}
		}
		return g
	}
}
