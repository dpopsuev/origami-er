package er

import (
	"context"
	_ "embed"
	"fmt"
	"io/fs"

	"github.com/dpopsuev/origami/circuit"
	"github.com/dpopsuev/origami/engine"
)

//go:embed circuit.yaml
var defaultCircuitYAML []byte

// DefaultCircuitYAML returns the embedded base ER circuit definition.
func DefaultCircuitYAML() []byte { return defaultCircuitYAML }

// SchematicResolver returns an AssetResolver that resolves "er" to the
// embedded base circuit.
func SchematicResolver() circuit.AssetResolver {
	return func(name string) ([]byte, error) {
		if name == "er" {
			return defaultCircuitYAML, nil
		}
		return nil, fmt.Errorf("unknown schematic %q", name)
	}
}

// erSessionFactory implements engine.SessionFactory for the ER schematic.
type erSessionFactory struct{}

func (f *erSessionFactory) CreateSession(ctx context.Context, params *engine.SessionParams) (*engine.SessionConfig, error) {
	return createSession(ctx, params)
}

func (f *erSessionFactory) FormatReport(result any) (string, any, error) {
	return fmt.Sprintf("%v", result), result, nil
}

// Factory returns the SessionFactory for the ER schematic.
func Factory() engine.SessionFactory {
	return &erSessionFactory{}
}

func createSession(_ context.Context, params *engine.SessionParams) (*engine.SessionConfig, error) {
	circuitData := defaultCircuitYAML

	// Load consumer overlay if available in DomainFS.
	if params.DomainFS != nil {
		if overlay, err := fs.ReadFile(params.DomainFS, "circuits/collect-ground-truth.yaml"); err == nil {
			circuitData = overlay
		}
	}

	def, err := circuit.LoadCircuitWithOverlay(circuitData, SchematicResolver())
	if err != nil {
		return nil, fmt.Errorf("load ER circuit: %w", err)
	}

	// The ER circuit uses LLM dispatch for semantic matching.
	// Gather data (failures, tickets, PRs) comes via walker context
	// injected by the consumer's orchestrator.

	return &engine.SessionConfig{
		CircuitDef: def,
		Meta: engine.SessionMeta{
			TotalCases: 1,
			Scenario:   "er",
		},
	}, nil
}
