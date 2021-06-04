package usecase

import (
	"errors"
	"log"
	"strconv"

	"github.com/Snikimonkd/dataBases/internal/models"
	thread_repository "github.com/Snikimonkd/dataBases/internal/thread/repository"
)

type ThreadUseCase struct {
	Repository thread_repository.ThreadRepository
}

func (u *ThreadUseCase) ThreadCreate(newThread models.Thread) (interface{}, int, error) {
	forums, err := u.Repository.CheckForum(newThread)
	if err != nil {
		log.Println(err)
		return nil, 500, nil
	}
	if len(forums) == 0 {
		return nil, 404, errors.New("can`t find forum")
	}

	newThread.Forum = forums[0].Slug

	users, err := u.Repository.CheckUser(newThread)
	if err != nil {
		log.Println(err)
		return nil, 500, nil
	}
	if len(users) == 0 {
		return nil, 404, errors.New("can`t find user")
	}

	threads, err := u.Repository.CheckThread(newThread)
	if err != nil {
		log.Println(err)
		return nil, 500, nil
	}
	if len(threads) != 0 {
		return threads[0], 409, nil
	}

	newThread.Id, err = u.Repository.CreateThread(newThread)
	if err != nil {
		log.Println(err)
		return nil, 500, nil
	}

	return newThread, 201, nil
}

func (u *ThreadUseCase) ThreadGetOne(slug_or_id string) (interface{}, int, error) {
	id, err := strconv.Atoi(slug_or_id)
	if err == nil {
		threads, err := u.Repository.ThreadGetOneId(id)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}

		if len(threads) == 0 {
			return nil, 404, errors.New("cant find thread")
		}

		return threads[0], 200, nil
	}

	threads, err := u.Repository.ThreadGetOneSlug(slug_or_id)
	if err != nil {
		log.Println(err)
		return nil, 500, nil
	}

	if len(threads) == 0 {
		return nil, 404, errors.New("cant find thread")
	}

	return threads[0], 200, nil
}
