package main

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	vsm "github.com/iso22468/vsm-go-sdk"
)

// Resource Handlers

func handleSymbolsResource(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	// Get all VSM symbols
	symbols := []map[string]interface{}{}

	// Add process symbols
	processSymbols := vsm.GetSymbolsByCategory(vsm.CategoryProcess)
	for _, symbol := range processSymbols {
		symbols = append(symbols, map[string]interface{}{
			"code":        symbol.Code,
			"name":        symbol.Name,
			"description": symbol.Description,
			"category":    symbol.Category,
			"position":    symbol.Position,
		})
	}

	// Add control symbols (similar to inventory)
	controlSymbols := vsm.GetSymbolsByCategory(vsm.CategoryControl)
	for _, symbol := range controlSymbols {
		symbols = append(symbols, map[string]interface{}{
			"code":        symbol.Code,
			"name":        symbol.Name,
			"description": symbol.Description,
			"category":    symbol.Category,
			"position":    symbol.Position,
		})
	}

	// Add flow symbols
	flowSymbols := vsm.GetSymbolsByCategory(vsm.CategoryFlow)
	for _, symbol := range flowSymbols {
		symbols = append(symbols, map[string]interface{}{
			"code":        symbol.Code,
			"name":        symbol.Name,
			"description": symbol.Description,
			"category":    symbol.Category,
			"position":    symbol.Position,
		})
	}


	symbolsJSON, err := json.MarshalIndent(symbols, "", "  ")
	if err != nil {
		return nil, err
	}

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      req.Params.URI,
			MIMEType: "application/json",
			Text:     string(symbolsJSON),
		},
	}, nil
}

func handleStandardsResource(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	standards := map[string]interface{}{
		"standard": "ISO 22468:2020",
		"title": "Value stream management (VSM)",
		"compliance_matrix": map[string]interface{}{
			"core_methodology": map[string]interface{}{
				"status": "Fully Implemented",
				"description": "Complete VSM 4-phase methodology (Analysis, Design, Planning, Assessment)",
			},
			"annex_a_calculations": map[string]interface{}{
				"status": "Fully Implemented",
				"description": "All calculation procedures for KPIs, process parameters, and performance metrics",
			},
			"annex_b_data_boxes": map[string]interface{}{
				"status": "Fully Implemented",
				"description": "Standard data box formats for process information and inventory",
			},
			"annex_c_symbols": map[string]interface{}{
				"status": "Fully Implemented",
				"description": "Complete set of VSM symbols for mapping and visualization",
			},
			"seven_types_of_waste": map[string]interface{}{
				"status": "Fully Implemented",
				"description": "Transportation, Inventory, Motion, Waiting, Over-processing, Over-production, Defects",
			},
			"eight_design_principles": map[string]interface{}{
				"status": "Fully Implemented",
				"description": "Takt time, Supermarket, Continuous flow, Pull systems, Pacemaker, Levelling, Release, Continuous improvement",
			},
		},
		"roth_schook_design_principles": []string{
			"Takt Time: Calculate and use customer takt time to pace production",
			"Supermarket: Implement controlled inventory points for pull systems",
			"Continuous Flow: Design processes for smooth, uninterrupted flow",
			"Pull Systems: Use demand signals to control production",
			"Pacemaker: Establish a pacemaker process to control the entire value stream",
			"Levelling (Heijunka): Level production mix and volume to reduce variability",
			"Release: Control work release to prevent overburdening",
			"Continuous Improvement: Establish mechanisms for ongoing kaizen activities",
		},
		"calculation_accuracy": "All calculations follow the exact formulas specified in ISO 22468:2020 Annex A",
		"validation_testing": []string{
			"Parameter range checking",
			"Formula accuracy verification",
			"Data integrity validation",
			"Cross-reference validation",
			"Standards compliance testing",
		},
	}

	standardsJSON, err := json.MarshalIndent(standards, "", "  ")
	if err != nil {
		return nil, err
	}

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      req.Params.URI,
			MIMEType: "application/json",
			Text:     string(standardsJSON),
		},
	}, nil
}

func handleTemplatesResource(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	templates := []map[string]interface{}{
		{
			"name":        "Manufacturing Line Template",
			"description": "Template for a basic manufacturing production line",
			"category":    "Manufacturing",
			"processes": []map[string]interface{}{
				{
					"id":          "receiving",
					"name":        "Receiving",
					"type":        "material",
					"description": "Receive and inspect incoming materials",
				},
				{
					"id":          "production",
					"name":        "Production",
					"type":        "material",
					"description": "Main production process",
				},
				{
					"id":          "quality_check",
					"name":        "Quality Check",
					"type":        "material",
					"description": "Final quality inspection",
				},
				{
					"id":          "shipping",
					"name":        "Shipping",
					"type":        "material",
					"description": "Package and ship finished products",
				},
			},
			"inventory_points": []string{"raw_materials", "finished_goods"},
		},
		{
			"name":        "Service Process Template",
			"description": "Template for a service-based value stream",
			"category":    "Service",
			"processes": []map[string]interface{}{
				{
					"id":          "customer_request",
					"name":        "Customer Request",
					"type":        "data",
					"description": "Receive and validate customer requests",
				},
				{
					"id":          "processing",
					"name":        "Service Processing",
					"type":        "data",
					"description": "Process the customer service request",
				},
				{
					"id":          "verification",
					"name":        "Verification",
					"type":        "data",
					"description": "Verify service completion and quality",
				},
				{
					"id":          "completion",
					"name":        "Service Completion",
					"type":        "data",
					"description": "Complete service and notify customer",
				},
			},
			"inventory_points": []string{"pending_requests", "completed_services"},
		},
		{
			"name":        "Order Fulfillment Template",
			"description": "Template for order processing and fulfillment",
			"category":    "E-commerce",
			"processes": []map[string]interface{}{
				{
					"id":          "order_entry",
					"name":        "Order Entry",
					"type":        "data",
					"description": "Enter and validate customer orders",
				},
				{
					"id":          "picking",
					"name":        "Order Picking",
					"type":        "material",
					"description": "Pick items from inventory",
				},
				{
					"id":          "packing",
					"name":        "Packing",
					"type":        "material",
					"description": "Pack orders for shipping",
				},
				{
					"id":          "shipping",
					"name":        "Shipping",
					"type":        "material",
					"description": "Ship orders to customers",
				},
			},
			"inventory_points": []string{"order_queue", "picked_orders", "packed_orders"},
		},
	}

	templatesJSON, err := json.MarshalIndent(templates, "", "  ")
	if err != nil {
		return nil, err
	}

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      req.Params.URI,
			MIMEType: "application/json",
			Text:     string(templatesJSON),
		},
	}, nil
}
