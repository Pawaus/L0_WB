package usecase

type Repo interface {
	ProcessOrder(uid, data string)
	LoadCache()
}

type Usecase struct {
	repo Repo
}

func NewUseCase(r Repo) *Usecase {
	return &Usecase{repo: r}
}

func (u *Usecase) ProcessOrder(uid, data string) {
	u.repo.ProcessOrder(uid, data)
}

func (u *Usecase) LoadCache() {
	u.repo.LoadCache()
}
