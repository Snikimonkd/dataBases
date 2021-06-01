package usecase

import forum_repository "github.com/Snikimonkd/dataBases/internal/forum/repository"

type ForumUseCase struct {
	Repository forum_repository.ForumRepository
}
