package usecase

import post_repository "github.com/Snikimonkd/dataBases/internal/post/repository"

type PostUseCase struct {
	Repository post_repository.PostRepository
}
