package maintenance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ImpactType enumerates the values for impact type.
type ImpactType string

const (
	// Freeze ...
	Freeze ImpactType = "Freeze"
	// None ...
	None ImpactType = "None"
	// Redeploy ...
	Redeploy ImpactType = "Redeploy"
	// Restart ...
	Restart ImpactType = "Restart"
)

// PossibleImpactTypeValues returns an array of possible values for the ImpactType const type.
func PossibleImpactTypeValues() []ImpactType {
	return []ImpactType{Freeze, None, Redeploy, Restart}
}

// Scope enumerates the values for scope.
type Scope string

const (
	// ScopeAll ...
	ScopeAll Scope = "All"
	// ScopeHost ...
	ScopeHost Scope = "Host"
	// ScopeInResource ...
	ScopeInResource Scope = "InResource"
	// ScopeResource ...
	ScopeResource Scope = "Resource"
)

// PossibleScopeValues returns an array of possible values for the Scope const type.
func PossibleScopeValues() []Scope {
	return []Scope{ScopeAll, ScopeHost, ScopeInResource, ScopeResource}
}

// UpdateStatus enumerates the values for update status.
type UpdateStatus string

const (
	// Completed ...
	Completed UpdateStatus = "Completed"
	// InProgress ...
	InProgress UpdateStatus = "InProgress"
	// Pending ...
	Pending UpdateStatus = "Pending"
	// RetryLater ...
	RetryLater UpdateStatus = "RetryLater"
	// RetryNow ...
	RetryNow UpdateStatus = "RetryNow"
)

// PossibleUpdateStatusValues returns an array of possible values for the UpdateStatus const type.
func PossibleUpdateStatusValues() []UpdateStatus {
	return []UpdateStatus{Completed, InProgress, Pending, RetryLater, RetryNow}
}
