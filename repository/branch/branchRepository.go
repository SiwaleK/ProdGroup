package repo

type BranchRepository interface {
}

type branchRepository struct{}

func NewBranchRepository() BranchRepository {
	return &branchRepository{}
}
