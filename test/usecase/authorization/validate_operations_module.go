package authorization

type OperationsModuleParams struct {
	ModuleID    int64
	OperationID int64
}

func ValidateOperationPerModule(opMod OperationsModuleParams) {
	// search in the persistence layer if the module searched exists

}
