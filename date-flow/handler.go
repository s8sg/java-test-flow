package function

import (
	//	"fmt"
	faasflow "github.com/s8sg/faas-flow"
	"os"
)

func getDateFunc() string {
	function := os.Getenv("date-function")
	if function == "" {
		function = "date"
	}
	return function
}

// Define provide definiton of the workflow
func Define(flow *faasflow.Workflow, context *faasflow.Context) (err error) {
	function := getDateFunc()
	flow.SyncNode().Apply(function).Apply(function)
	// flow.SyncNode().Apply("date-node").Apply("date-python").Apply("date-go").Apply("date-java")
	// flow.SyncNode().Apply("date-python").Apply("date-node").Apply("date-go")
	// flow.SyncNode().Apply("date-java").Apply("date-python").Apply("date-node").Apply("date-go")
	return
}

// DefineStateStore provides the override of the default StateStore
func DefineStateStore() (faasflow.StateStore, error) {
	return nil, nil
}

// ProvideDataStore provides the override of the default DataStore
func DefineDataStore() (faasflow.DataStore, error) {
	return nil, nil
}
