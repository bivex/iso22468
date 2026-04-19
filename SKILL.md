# SKILL: Value Stream Management (VSM) Analysis
## Based on ISO 22468:2020 

---

## WHEN TO USE THIS SKILL

Use this skill when the user asks to:
- Analyse a manufacturing or service value stream
- Calculate VSM KPIs (lead time, cycle time, efficiency, takt time, etc.)
- Identify bottlenecks or waste in a process
- Design a future state for a production line
- Build an implementation plan for lean improvements
- Assess performance against VSM targets
- Serialize / save / load value stream data

---

## STEP 0 — GATHER CONTEXT BEFORE WRITING CODE

Before generating any code, ask (or infer from context) the following:

| Question | Why it matters |
|----------|---------------|
| What processes exist in the value stream? | Need IDs, names, types |
| What are the time parameters per process? (PT, CT, setup, wait) | Core KPI inputs |
| What is customer demand and takt time? | Required for bottleneck and efficiency calcs |
| What inventory points exist? | Needed for inventory turns |
| What is the goal? (analysis / design / planning / assessment) | Determines which SDK component to use |
| What time unit is used? (minutes / hours / seconds) | Must be consistent across all processes |
| Is yield / scrap / rework data available? | Required for YF and quality KPIs |

If any critical input is missing — ask the user before writing code.

---

## SDK COMPONENT MAP

Choose the right component based on what the user needs:

| User Goal | SDK Component | Key Method |
|-----------|--------------|------------|
| Understand current state, find waste | Analyzer | AnalyzeCurrentState |
| Find bottlenecks only | Analyzer | DetectBottlenecks |
| Design improved future state | Designer | DesignFutureState |
| Generate improvement actions | Planner | PlanImplementation |
| Track KPIs over time | Assessor | AssessValueStream |
| Calculate KPIs manually | Calculator | CalculateKPIs |
| Save/load value stream | Serializer | SaveToFile / LoadFromFile |

---

## CORE DATA TO ALWAYS POPULATE

### ValueStream — required fields
- Unique ID and human-readable name
- State: Current, Future, or Ideal
- Product family name
- Customer: ID, name, demand (units per period), takt time (time per unit)
- List of processes

### Process — required fields per process
- Unique ID and name
- Type: Material, Energy, or Data
- Process Time — total time to complete one cycle (must be greater than 0)
- Cycle Time — time between completions (must be greater than 0)
- Yield Factor — value between 0.0 and 1.0 (default 1.0 if no scrap data)
- Setup Time, Waiting Time, Handling Time, Admin Time — set to 0 if unknown
- Time Unit — must be the same across ALL processes
- Availability — value between 0.0 and 1.0 (default 1.0)

WARNING: Process Time, Cycle Time, and Customer Demand must be greater than zero — otherwise analysis returns empty results.

---

## KPI FORMULAS (for validation and explanation to user)

| KPI | Formula |
|-----|---------|
| Process Time | PT = Processing Time + Setup Time + Handling Time + Waiting Time + Admin Time |
| Lead Time | LT = Value-Added Time + Non-Value-Added Time + Necessary Non-Value-Added Time |
| Value-Added Ratio | VAR = VAT divided by LT, multiplied by 100% |
| Process Efficiency | PE = VAT divided by PT, multiplied by 100% |
| Customer Takt Time | CTT = Operating Time divided by Customer Demand |
| Inventory Turns | IT = Customer Demand divided by Average Inventory Quantity |
| Yield Factor | YF = 1 minus Scrap Rate minus Rework Rate |
| Bottleneck Index | BDI = (Cycle Time × Demand) divided by Capacity — value above 1.0 signals a bottleneck |
| OTIF | (Orders On Time × Orders In Full) divided by Total Orders, multiplied by 100% |

---

## STANDARD WORKFLOW PATTERNS

### Pattern A — Current State Analysis
Use the Analyzer component. Optional tuning: sensitivity threshold (0.0–1.0, higher = fewer bottlenecks reported) and whether to include detailed waste analysis.
Results give: list of bottleneck processes, waste reduction potential as a percentage.

### Pattern B — KPI Calculation
Use the Calculator component. Pass the fully populated ValueStream.
Results give: a map of KPI names to numeric values. Multiply ratio values by 100 to get percentages.

### Pattern C — Future State Design
Requires completed current state analysis as input. Also requires a Project definition with objectives, timeframe, and budget.
Use the Designer component.
Results give: list of improvement potentials with estimated impact.

### Pattern D — Performance Assessment
Requires a targets map (KPI name to target value) and historical performance data with timestamps.
Use the Assessor component.
Results give: performance index (0.0–1.0), list of recommended action items.

### Pattern E — Save and Load
Use the Serializer component. Enable pretty print and validation before saving.
Saves to and loads from JSON files. For large datasets, enable GZIP compression.

---

## CHECKLIST BEFORE RETURNING CODE TO USER

- All Process Time and Cycle Time values are greater than zero
- Customer Demand is set and greater than zero
- All processes use the same Time Unit
- Yield Factor is between 0.0 and 1.0
- Availability is between 0.0 and 1.0 (default 1.0)
- No circular process input/output references
- Error handling is present on all SDK calls
- If saving data — validation is enabled on the Serializer

---

## TROUBLESHOOTING QUICK REFERENCE

| Symptom | Likely cause | Fix |
|---------|-------------|-----|
| Empty bottleneck list | Customer Demand = 0 or processes not linked | Set demand, check process connections |
| KPIs all zero | Process Time or Cycle Time = 0 | Populate all required parameters |
| Serialization error | Circular references or invalid data | Enable validation on Serializer |
| Incorrect ratio values | Mixed time units across processes | Standardize all to one Time Unit |
| Memory grows with large streams | Waste analysis enabled on big dataset | Disable waste analysis or process in chunks |

---

## WASTE TYPES TO REFERENCE WHEN EXPLAINING RESULTS (ISO 22468 Muda)

1. Transportation — unnecessary movement of materials
2. Inventory — excess stock between processes
3. Motion — unnecessary movement of people
4. Waiting — idle time between process steps
5. Over-processing — more processing than the customer requires
6. Over-production — producing more than demand
7. Defects — scrap and rework

---

## OUTPUT FORMAT GUIDANCE

When presenting analysis results to the user, always include:
1. Summary table of key KPIs with values and benchmark comparison
2. Bottleneck identification — which process has BDI above 1.0 and why
3. Top 3 improvement opportunities from waste analysis
4. Next recommended action — which SDK phase to run next

---
