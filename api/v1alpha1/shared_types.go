package v1alpha1

// poorman's golang enum
type CrCondition string

var (
	CrSubmitted CrCondition = ""
	CrPending   CrCondition = "pending"
	CrForbidden CrCondition = "forbidden"
	CrFailed    CrCondition = "failed"
	CrOK        CrCondition = "created"
)

func (i CrCondition) String() string {
	return string(i)
}