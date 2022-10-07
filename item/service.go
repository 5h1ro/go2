package item

type Service interface {
	Get() ([]Item, error)
	GetByFK(OrderID int) ([]Item, error)
	Find(ItemID int) (Item, error)
	Create(req ItemRequest) (Item, error)
	Update(ItemID int, req ItemRequest) (Item, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get() ([]Item, error) {
	items, err := s.repository.Get()
	return items, err
}

func (s *service) Find(ItemID int) (Item, error) {
	item, err := s.repository.Find(ItemID)
	return item, err
}

func (s *service) GetByFK(OrderID int) ([]Item, error) {
	items, err := s.repository.GetByFK(OrderID)
	return items, err
}

func (s *service) Create(req ItemRequest) (Item, error) {
	item := Item{
		ItemCode:    req.ItemCode,
		Description: req.Description,
		Quantity:    req.Quantity,
		OrderID:     req.OrderID,
	}
	NewItem, err := s.repository.Create(item)
	return NewItem, err
}

func (s *service) Update(ItemID int, req ItemRequest) (Item, error) {

	item, err := s.repository.Find(ItemID)

	item.ItemCode = req.ItemCode
	item.Description = req.Description
	item.Quantity = req.Quantity

	newItem, err := s.repository.Update(item)

	return newItem, err
}
