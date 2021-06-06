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

func (u *ThreadUseCase) ThreadVote(slug_or_id string, newVote models.Vote) (interface{}, int, error) {
	var threads []models.Thread
	id, err := strconv.Atoi(slug_or_id)
	if err == nil {
		threads, err = u.Repository.ThreadGetOneId(id)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
		if len(threads) == 0 {
			return nil, 404, errors.New("cant find thread 1")
		}
	} else {
		threads, err = u.Repository.ThreadGetOneSlug(slug_or_id)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
		if len(threads) == 0 {
			return nil, 404, errors.New("cant find thread 2")
		}
	}

	thread := threads[0]

	lastVotes, err := u.Repository.GetVote(newVote, thread.Id)
	if err != nil {
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
	}
	if len(lastVotes) != 0 {
		if lastVotes[0].Voice == newVote.Voice {
			return thread, 200, nil
		}
		err = u.Repository.UpdateVote(newVote, thread.Id)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}

		thread.Votes += newVote.Voice * 2
		return thread, 200, nil
	}

	err = u.Repository.InsertNewVote(newVote, thread.Id)
	if err != nil {
		log.Println(err)
		return nil, 500, nil
	}

	thread.Votes += newVote.Voice
	return thread, 200, nil
}

func (u *ThreadUseCase) ThreadGetPosts(slug_or_id string, limitInt int, descBool bool, since string, sort string) (interface{}, int, error) {
	var threads []models.Thread
	id, err := strconv.Atoi(slug_or_id)
	if err == nil {
		threads, err = u.Repository.ThreadGetOneId(id)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
		if len(threads) == 0 {
			return nil, 404, errors.New("cant find thread 1")
		}
	} else {
		threads, err = u.Repository.ThreadGetOneSlug(slug_or_id)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
		if len(threads) == 0 {
			return nil, 404, errors.New("cant find thread 2")
		}
	}

	thread := threads[0]

	var posts []models.Post
	switch sort {
	case "flat":
		posts, err = u.Repository.ThreadGetPostsFlat(limitInt, descBool, since, thread)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
	case "tree":
		posts, err = u.Repository.ThreadGetPostsTree(limitInt, descBool, since, thread)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
	case "parent_tree":
		posts, err = u.Repository.ThreadGetPostsParentTree(limitInt, descBool, since, thread)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
	}

	return posts, 200, nil
}

func (u *ThreadUseCase) ThreadUpdate(slug_or_id string, newThread models.Thread) (interface{}, int, error) {
	var threads []models.Thread
	id, err := strconv.Atoi(slug_or_id)
	if err == nil {
		threads, err = u.Repository.ThreadGetOneId(id)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
		if len(threads) == 0 {
			return nil, 404, errors.New("cant find thread 1")
		}
	} else {
		threads, err = u.Repository.ThreadGetOneSlug(slug_or_id)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
		if len(threads) == 0 {
			return nil, 404, errors.New("cant find thread 2")
		}
	}

	thread := threads[0]

	thread = u.ChangeThread(newThread, thread)

	err = u.Repository.ThreadUpdate(thread)
	if err != nil {
		log.Println(err)
		return nil, 500, nil
	}

	return thread, 200, nil
}

func (u *ThreadUseCase) ChangeThread(newThread models.Thread, oldThread models.Thread) models.Thread {
	if newThread.Message != "" {
		oldThread.Message = newThread.Message
	}

	if newThread.Title != "" {
		oldThread.Title = newThread.Title
	}

	return oldThread
}
