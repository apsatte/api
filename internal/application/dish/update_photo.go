package dish_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) UpdatePhoto(c context.Context, dto *UpdatePhotoInput) (*UpdatePhotoOutput, error) {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return nil, err
	}

	cstmr, err := u.customerRepo.GetOneByID(c, session.CustomerID)
	if err != nil {
		return nil, err
	}

	//

	dish, err := u.dishRepo.GetOneByID(c, dto.DishID)
	if err != nil {
		return nil, err
	}

	category, err := u.categoryRepo.GetOneByID(c, dish.CategoryID)
	if err != nil {
		return nil, err
	}

	prjct, err := u.projectRepo.GetOneByID(c, category.ProjectID)
	if err != nil {
		return nil, err
	}

	if !prjct.IsOwner(cstmr.ID) {
		return nil, domain.ErrForbidden
	}

	// upload photo
	photoURL, err := u.storage.PutImage(c, dto.Photo, 512)
	if err != nil {
		return nil, err
	}

	photoMiniURL, err := u.storage.PutImage(c, dto.Photo, 128)
	if err != nil {
		u.storage.Remove(c, photoURL)
		return nil, err
	}

	u.storage.Remove(c, dish.PhotoURL)
	u.storage.Remove(c, dish.PhotoMiniURL)

	dish.UpdatePhoto(photoURL, photoMiniURL)

	if err := u.dishRepo.Save(c, dish); err != nil {
		return nil, err
	}

	return &UpdatePhotoOutput{
		PhotoURL:     photoURL,
		PhotoMiniURL: photoMiniURL,
	}, nil
}
