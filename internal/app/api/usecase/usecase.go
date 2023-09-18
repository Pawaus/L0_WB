package usecase

type Repo interface {
	GetOrderByID(uid string) (string, error)
}

type Usecase struct {
	repo Repo
}

func NewUsecase(repo Repo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) GetOrderByID(uid string) (string, error) {
	return u.repo.GetOrderByID(uid)
}
