package main

import (
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server with full capabilities
	s := server.NewMCPServer(
		"ISO 22468:2020 VSM MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithLogging(),
		server.WithRecovery(),
		server.WithInstructions("A comprehensive MCP server for ISO 22468:2020 Value Stream Management (VSM) operations including analysis, design, planning, assessment, and serialization of value streams."),
	)

	// Initialize VSM components
	setupVSMTools(s)

	// Initialize VSM resources
	setupVSMResources(s)

	// Initialize VSM prompts
	setupVSMPrompts(s)

	// Start the server using stdio transport
	log.Println("Starting ISO 22468:2020 VSM MCP Server...")
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func setupVSMTools(s *server.MCPServer) {
	// VSM Analysis Tools
	setupVSMAnalysisTools(s)

	// VSM Design Tools
	setupVSMDesignTools(s)

	// VSM Planning Tools
	setupVSMPlanningTools(s)

	// VSM Assessment Tools
	setupVSMAssessmentTools(s)

	// VSM Serialization Tools
	setupVSMSerializationTools(s)

	// VSM Utility Tools
	setupVSMUtilityTools(s)
}

func setupVSMAnalysisTools(s *server.MCPServer) {
	// Create Value Stream Analysis Tool
	analyzeTool := mcp.NewTool("vsm_analyze_current_state",
		mcp.WithDescription("Analyze the current state of a value stream to identify bottlenecks, waste, and improvement opportunities"),
		mcp.WithString("value_stream_json",
			mcp.Required(),
			mcp.Description("JSON representation of the ValueStream to analyze"),
		),
	)

	s.AddTool(analyzeTool, handleVSMAnalysis)

	// Detect Bottlenecks Tool
	bottleneckTool := mcp.NewTool("vsm_detect_bottlenecks",
		mcp.WithDescription("Detect bottlenecks in a value stream based on process cycle times and customer demand"),
		mcp.WithString("value_stream_json",
			mcp.Required(),
			mcp.Description("JSON representation of the ValueStream to analyze for bottlenecks"),
		),
	)

	s.AddTool(bottleneckTool, handleBottleneckDetection)

	// Waste Analysis Tool
	wasteTool := mcp.NewTool("vsm_analyze_waste",
		mcp.WithDescription("Perform detailed waste analysis on a value stream to identify the 8 types of waste (Muda)"),
		mcp.WithString("value_stream_json",
			mcp.Required(),
			mcp.Description("JSON representation of the ValueStream to analyze for waste"),
		),
	)

	s.AddTool(wasteTool, handleWasteAnalysis)
}

func setupVSMDesignTools(s *server.MCPServer) {
	// Future State Design Tool
	designTool := mcp.NewTool("vsm_design_future_state",
		mcp.WithDescription("Design future state improvements for a value stream based on analysis results"),
		mcp.WithString("current_state_json",
			mcp.Required(),
			mcp.Description("JSON representation of the current ValueStream state"),
		),
		mcp.WithString("analysis_result_json",
			mcp.Required(),
			mcp.Description("JSON representation of the analysis results"),
		),
		mcp.WithString("project_json",
			mcp.Required(),
			mcp.Description("JSON representation of the VSM project details"),
		),
	)

	s.AddTool(designTool, handleFutureStateDesign)

	// Ideal State Generation Tool
	idealTool := mcp.NewTool("vsm_generate_ideal_state",
		mcp.WithDescription("Generate an ideal state value stream design with minimal waste and optimal flow"),
		mcp.WithString("current_state_json",
			mcp.Required(),
			mcp.Description("JSON representation of the current ValueStream to base ideal state on"),
		),
	)

	s.AddTool(idealTool, handleIdealStateGeneration)
}

func setupVSMPlanningTools(s *server.MCPServer) {
	// Implementation Planning Tool
	planTool := mcp.NewTool("vsm_plan_implementation",
		mcp.WithDescription("Create implementation plans with SMART measures and timelines"),
		mcp.WithString("design_result_json",
			mcp.Required(),
			mcp.Description("JSON representation of the design results"),
		),
		mcp.WithString("project_json",
			mcp.Required(),
			mcp.Description("JSON representation of the VSM project details"),
		),
	)

	s.AddTool(planTool, handleImplementationPlanning)


}

func setupVSMAssessmentTools(s *server.MCPServer) {
	// Performance Assessment Tool
	assessTool := mcp.NewTool("vsm_assess_performance",
		mcp.WithDescription("Assess value stream performance against targets and generate action items"),
		mcp.WithString("value_stream_json",
			mcp.Required(),
			mcp.Description("JSON representation of the ValueStream to assess"),
		),
		mcp.WithString("targets_json",
			mcp.Required(),
			mcp.Description("JSON object with KPI target values"),
		),
		mcp.WithString("history_json",
			mcp.Description("JSON array of historical AssessmentResult (optional)"),
		),
	)

	s.AddTool(assessTool, handlePerformanceAssessment)


}

func setupVSMSerializationTools(s *server.MCPServer) {
	// Serialize Value Stream Tool
	serializeTool := mcp.NewTool("vsm_serialize_value_stream",
		mcp.WithDescription("Serialize a value stream to JSON format"),
		mcp.WithString("value_stream_json",
			mcp.Required(),
			mcp.Description("JSON representation of the ValueStream to serialize"),
		),
		mcp.WithBoolean("pretty_print",
			mcp.DefaultBool(false),
			mcp.Description("Enable pretty-printed JSON output"),
		),
		mcp.WithBoolean("include_metadata",
			mcp.DefaultBool(true),
			mcp.Description("Include timestamps and version metadata"),
		),
	)

	s.AddTool(serializeTool, handleValueStreamSerialization)

	// Deserialize Value Stream Tool
	deserializeTool := mcp.NewTool("vsm_deserialize_value_stream",
		mcp.WithDescription("Deserialize a value stream from JSON format"),
		mcp.WithString("json_data",
			mcp.Required(),
			mcp.Description("JSON string containing serialized ValueStream data"),
		),
	)

	s.AddTool(deserializeTool, handleValueStreamDeserialization)

	// Save to File Tool
	saveTool := mcp.NewTool("vsm_save_to_file",
		mcp.WithDescription("Save a value stream to a JSON file"),
		mcp.WithString("value_stream_json",
			mcp.Required(),
			mcp.Description("JSON representation of the ValueStream to save"),
		),
		mcp.WithString("filename",
			mcp.Required(),
			mcp.Description("Filename to save the value stream to"),
		),
	)

	s.AddTool(saveTool, handleSaveToFile)

	// Load from File Tool
	loadTool := mcp.NewTool("vsm_load_from_file",
		mcp.WithDescription("Load a value stream from a JSON file"),
		mcp.WithString("filename",
			mcp.Required(),
			mcp.Description("Filename to load the value stream from"),
		),
	)

	s.AddTool(loadTool, handleLoadFromFile)
}

func setupVSMUtilityTools(s *server.MCPServer) {
	// Calculate KPIs Tool
	kpiTool := mcp.NewTool("vsm_calculate_kpis",
		mcp.WithDescription("Calculate key performance indicators for a value stream"),
		mcp.WithString("value_stream_json",
			mcp.Required(),
			mcp.Description("JSON representation of the ValueStream to calculate KPIs for"),
		),
	)

	s.AddTool(kpiTool, handleKPICalculation)

	// Create Value Stream Tool
	createTool := mcp.NewTool("vsm_create_value_stream",
		mcp.WithDescription("Create a new value stream with basic structure"),
		mcp.WithString("id",
			mcp.Required(),
			mcp.Description("Unique identifier for the value stream"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Display name for the value stream"),
		),
		mcp.WithString("product_family",
			mcp.Required(),
			mcp.Description("Product family this value stream serves"),
		),
		mcp.WithString("customer_json",
			mcp.Description("JSON representation of the customer (optional)"),
		),
	)

	s.AddTool(createTool, handleCreateValueStream)
}

func setupVSMResources(s *server.MCPServer) {
	// VSM Symbols Resource
	symbolsResource := mcp.NewResource(
		"vsm://symbols",
		"VSM Symbols Reference",
		mcp.WithResourceDescription("Complete reference of ISO 22468:2020 VSM symbols"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(symbolsResource, handleSymbolsResource)

	// VSM Standards Resource
	standardsResource := mcp.NewResource(
		"vsm://standards",
		"VSM Standards Information",
		mcp.WithResourceDescription("ISO 22468:2020 standards compliance information"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(standardsResource, handleStandardsResource)

	// VSM Templates Resource
	templatesResource := mcp.NewResource(
		"vsm://templates",
		"VSM Templates",
		mcp.WithResourceDescription("Pre-defined value stream templates for common scenarios"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(templatesResource, handleTemplatesResource)
}

func setupVSMPrompts(s *server.MCPServer) {
	// VSM Analysis Prompt
	analysisPrompt := mcp.NewPrompt("vsm_analysis_workflow",
		mcp.WithPromptDescription("Guide through a complete VSM current state analysis"),
		mcp.WithArgument("industry",
			mcp.ArgumentDescription("Industry sector for context-specific guidance"),
		),
		mcp.WithArgument("complexity",
			mcp.ArgumentDescription("Value stream complexity level (simple, medium, complex)"),
		),
	)

	s.AddPrompt(analysisPrompt, handleVSMAnalysisPrompt)

	// VSM Implementation Prompt
	implementationPrompt := mcp.NewPrompt("vsm_implementation_guide",
		mcp.WithPromptDescription("Guide through VSM implementation planning and execution"),
		mcp.WithArgument("timeline",
			mcp.ArgumentDescription("Available timeline for implementation"),
		),
		mcp.WithArgument("resources",
			mcp.ArgumentDescription("Available resources and budget"),
		),
	)

	s.AddPrompt(implementationPrompt, handleVSMImplementationPrompt)
}
