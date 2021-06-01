package usecase

import thread_repository "github.com/Snikimonkd/dataBases/internal/thread/repository"

type ThreadUseCase struct {
	Repository thread_repository.ThreadRepository
}
