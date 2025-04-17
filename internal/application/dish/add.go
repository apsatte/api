package dish_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"api/internal/domain/menu"
	"context"
)

func (u *useCase) Add(c context.Context, dto *AddInput) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	cstmr, err := u.customerRepo.GetOneByID(c, session.CustomerID)
	if err != nil {
		return err
	}

	// check permission
	category, err := u.categoryRepo.GetOneByID(c, dto.CategoryID)
	if err != nil {
		return err
	}

	prjct, err := u.projectRepo.GetOneByID(c, category.ProjectID)
	if err != nil {
		return err
	}

	if !prjct.IsOwner(cstmr.ID) {
		return domain.ErrForbidden
	}

	// upload photos
	photoURL, err := u.storage.PutImage(c, dto.Photo, 512)
	if err != nil {
		return err
	}

	photoMiniURL, err := u.storage.PutImage(c, dto.Photo, 128)
	if err != nil {
		u.storage.Remove(c, photoURL)
		return err
	}

	// add the dish
	dish, err := menu.NewDish(
		dto.CategoryID, dto.Position,
		dto.Price, photoURL, photoMiniURL,
		dto.IsAvailable, dto.Translations)
	if err != nil {
		u.storage.Remove(c, photoURL)
		u.storage.Remove(c, photoMiniURL)
		return err
	}

	for _, opt := range dto.Options {
		items := []*menu.DishOptionItem{}

		for _, itm := range opt.Items {
			item, err := menu.NewDishOptionItem(itm.MaxQuantity, itm.PriceModifier, itm.Position, itm.Translations)
			if err != nil {
				u.storage.Remove(c, photoURL)
				u.storage.Remove(c, photoMiniURL)
				return err
			}
			items = append(items, item)
		}

		option, err := menu.NewDishOption(opt.MinQuantity, opt.MaxQuantity, opt.Position, opt.Translations, items)
		if err != nil {
			u.storage.Remove(c, photoURL)
			u.storage.Remove(c, photoMiniURL)
			return err
		}

		dish.AddOption(option)
	}

	if err := u.dishRepo.Save(c, dish); err != nil {
		u.storage.Remove(c, photoURL)
		u.storage.Remove(c, photoMiniURL)
		return err
	}

	return nil
}
