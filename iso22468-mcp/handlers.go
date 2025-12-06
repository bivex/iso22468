package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	vsm "github.com/iso22468/vsm-go-sdk"
)

// VSM Analysis Handlers

func handleVSMAnalysis(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	valueStreamJSON, err := req.RequireString("value_stream_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing value_stream_json: %v", err)), nil
	}

	var vs vsm.ValueStream
	if err := json.Unmarshal([]byte(valueStreamJSON), &vs); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid value stream JSON: %v", err)), nil
	}

	analyzer := vsm.NewAnalyzer()
	analysisResult, err := analyzer.AnalyzeCurrentState(&vs)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Analysis failed: %v", err)), nil
	}

	resultJSON, err := json.MarshalIndent(analysisResult, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Analysis completed successfully:\n%s", string(resultJSON))), nil
}

func handleBottleneckDetection(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	valueStreamJSON, err := req.RequireString("value_stream_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing value_stream_json: %v", err)), nil
	}

	var vs vsm.ValueStream
	if err := json.Unmarshal([]byte(valueStreamJSON), &vs); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid value stream JSON: %v", err)), nil
	}

	analyzer := vsm.NewAnalyzer()
	bottlenecks, err := analyzer.IdentifyBottlenecks(&vs)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Bottleneck detection failed: %v", err)), nil
	}

	resultJSON, err := json.MarshalIndent(bottlenecks, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Found %d bottlenecks:\n%s", len(bottlenecks), string(resultJSON))), nil
}

func handleWasteAnalysis(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	valueStreamJSON, err := req.RequireString("value_stream_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing value_stream_json: %v", err)), nil
	}

	var vs vsm.ValueStream
	if err := json.Unmarshal([]byte(valueStreamJSON), &vs); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid value stream JSON: %v", err)), nil
	}

	analyzer := vsm.NewAnalyzer()
	wasteAnalysis, err := analyzer.AnalyzeWaste(&vs)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Waste analysis failed: %v", err)), nil
	}

	resultJSON, err := json.MarshalIndent(wasteAnalysis, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Waste analysis completed:\n%s", string(resultJSON))), nil
}

// VSM Design Handlers

func handleFutureStateDesign(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	currentStateJSON, err := req.RequireString("current_state_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing current_state_json: %v", err)), nil
	}

	analysisResultJSON, err := req.RequireString("analysis_result_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing analysis_result_json: %v", err)), nil
	}

	var currentState vsm.ValueStream
	if err := json.Unmarshal([]byte(currentStateJSON), &currentState); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid current state JSON: %v", err)), nil
	}

	var analysisResult vsm.AnalysisResult
	if err := json.Unmarshal([]byte(analysisResultJSON), &analysisResult); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid analysis result JSON: %v", err)), nil
	}

	designer := vsm.NewDesigner()
	designResult, err := designer.DesignFutureState(&currentState, &analysisResult)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Future state design failed: %v", err)), nil
	}

	resultJSON, err := json.MarshalIndent(designResult, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Future state design completed:\n%s", string(resultJSON))), nil
}

func handleIdealStateGeneration(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	currentStateJSON, err := req.RequireString("current_state_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing current_state_json: %v", err)), nil
	}

	var currentState vsm.ValueStream
	if err := json.Unmarshal([]byte(currentStateJSON), &currentState); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid current state JSON: %v", err)), nil
	}

	designer := vsm.NewDesigner()
	idealState, err := designer.CreateIdealState(&currentState)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Ideal state generation failed: %v", err)), nil
	}

	resultJSON, err := json.MarshalIndent(idealState, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Ideal state generated:\n%s", string(resultJSON))), nil
}

// VSM Planning Handlers

func handleImplementationPlanning(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	designResultJSON, err := req.RequireString("design_result_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing design_result_json: %v", err)), nil
	}

	projectJSON, err := req.RequireString("project_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing project_json: %v", err)), nil
	}

	var designResult vsm.DesignResult
	if err := json.Unmarshal([]byte(designResultJSON), &designResult); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid design result JSON: %v", err)), nil
	}

	var project vsm.VSMProject
	if err := json.Unmarshal([]byte(projectJSON), &project); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid project JSON: %v", err)), nil
	}

	planner := vsm.NewPlanner()
	planningResult, err := planner.PlanImplementation(&designResult, &project)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Implementation planning failed: %v", err)), nil
	}

	resultJSON, err := json.MarshalIndent(planningResult, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Implementation plan created:\n%s", string(resultJSON))), nil
}


// VSM Assessment Handlers

func handlePerformanceAssessment(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	valueStreamJSON, err := req.RequireString("value_stream_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing value_stream_json: %v", err)), nil
	}

	targetsJSON, err := req.RequireString("targets_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing targets_json: %v", err)), nil
	}

	var vs vsm.ValueStream
	if err := json.Unmarshal([]byte(valueStreamJSON), &vs); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid value stream JSON: %v", err)), nil
	}

	var targets map[string]float64
	if err := json.Unmarshal([]byte(targetsJSON), &targets); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid targets JSON: %v", err)), nil
	}

	var history []vsm.AssessmentResult
	if historyJSON := req.GetString("history_json", ""); historyJSON != "" {
		if err := json.Unmarshal([]byte(historyJSON), &history); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Invalid history JSON: %v", err)), nil
		}
	}

	assessor := vsm.NewAssessor()
	assessment, err := assessor.AssessValueStream(&vs, targets, history)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Performance assessment failed: %v", err)), nil
	}

	resultJSON, err := json.MarshalIndent(assessment, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Performance assessment completed:\n%s", string(resultJSON))), nil
}


// VSM Serialization Handlers

func handleValueStreamSerialization(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	valueStreamJSON, err := req.RequireString("value_stream_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing value_stream_json: %v", err)), nil
	}

	prettyPrint := req.GetBool("pretty_print", false)

	var vs vsm.ValueStream
	if err := json.Unmarshal([]byte(valueStreamJSON), &vs); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid value stream JSON: %v", err)), nil
	}

	serializer := vsm.NewSerializer()

	serializedData, err := serializer.SerializeValueStream(&vs)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Serialization failed: %v", err)), nil
	}

	var result interface{}
	if prettyPrint {
		if err := json.Unmarshal(serializedData, &result); err == nil {
			prettyData, err := json.MarshalIndent(result, "", "  ")
			if err == nil {
				serializedData = prettyData
			}
		}
	}

	return mcp.NewToolResultText(fmt.Sprintf("Value stream serialized successfully (%d bytes):\n%s", len(serializedData), string(serializedData))), nil
}

func handleValueStreamDeserialization(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	jsonData, err := req.RequireString("json_data")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing json_data: %v", err)), nil
	}

	serializer := vsm.NewSerializer()
	vs, err := serializer.DeserializeValueStream([]byte(jsonData))
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Deserialization failed: %v", err)), nil
	}

	resultJSON, err := json.MarshalIndent(vs, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Value stream deserialized successfully:\n%s", string(resultJSON))), nil
}

func handleSaveToFile(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	valueStreamJSON, err := req.RequireString("value_stream_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing value_stream_json: %v", err)), nil
	}

	filename, err := req.RequireString("filename")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing filename: %v", err)), nil
	}

	var vs vsm.ValueStream
	if err := json.Unmarshal([]byte(valueStreamJSON), &vs); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid value stream JSON: %v", err)), nil
	}

	serializer := vsm.NewSerializer()
	err = serializer.SaveToFile(&vs, filename)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Save to file failed: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Value stream saved to file '%s' successfully", filename)), nil
}

func handleLoadFromFile(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	filename, err := req.RequireString("filename")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing filename: %v", err)), nil
	}

	serializer := vsm.NewSerializer()
	var vs vsm.ValueStream
	err = serializer.LoadFromFile(filename, &vs)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Load from file failed: %v", err)), nil
	}

	resultJSON, err := json.MarshalIndent(vs, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Value stream loaded from file '%s':\n%s", filename, string(resultJSON))), nil
}

// VSM Utility Handlers

func handleKPICalculation(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	valueStreamJSON, err := req.RequireString("value_stream_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing value_stream_json: %v", err)), nil
	}

	var vs vsm.ValueStream
	if err := json.Unmarshal([]byte(valueStreamJSON), &vs); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid value stream JSON: %v", err)), nil
	}

	calculator := vsm.NewCalculator()
	kpis := calculator.CalculateKPIs(&vs)

	resultJSON, err := json.MarshalIndent(kpis, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("KPIs calculated successfully:\n%s", string(resultJSON))), nil
}

func handleCreateValueStream(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing id: %v", err)), nil
	}

	name, err := req.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing name: %v", err)), nil
	}

	productFamily, err := req.RequireString("product_family")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing product_family: %v", err)), nil
	}

	var customer *vsm.Customer
	if customerJSON := req.GetString("customer_json", ""); customerJSON != "" {
		if err := json.Unmarshal([]byte(customerJSON), &customer); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Invalid customer JSON: %v", err)), nil
		}
	}

	vs := &vsm.ValueStream{
		ID:            id,
		Name:          name,
		State:         vsm.StateCurrent,
		ProductFamily: productFamily,
		Customer:      customer,
		Processes:     []vsm.Process{},
		Inventory:     []vsm.Inventory{},
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Version:       "1.0",
	}

	resultJSON, err := json.MarshalIndent(vs, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to serialize result: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Value stream created successfully:\n%s", string(resultJSON))), nil
}
