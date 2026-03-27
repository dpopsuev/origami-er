// Package er provides the Entity Resolution schematic for Origami.
//
// Entity Resolution matches records across disconnected data sources
// (RP failures, Jira tickets, GitHub PRs) using LLM-driven semantic
// matching. Used by consumers to build and maintain ground truth datasets.
//
// This is a schematic — it provides circuit definitions, transformers,
// and hooks. The engine walks the circuit. The consumer declares the
// topology in circuits/collect-ground-truth.yaml.
package er
